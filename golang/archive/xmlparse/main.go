package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"flag"

	"gopkg.in/cheggaaa/pb.v1"
)

type Osm struct {
	Name  xml.Name `xml:"osm"`
	Nodes []*Node  `xml:"node"`
	Ways  []*Way   `xml:"way"`
}

type Node struct {
	Name xml.Name `xml:"node"`
	Id   int64    `xml:"id,attr"`
}
type Way struct {
	Name xml.Name `xml:"way"`
	Id   int64    `xml:"id,attr"`
	Nds  []*Nd    `xml:"nd"`
}
type Nd struct {
	Name xml.Name `xml:"nd"`
	Ref  int64    `xml:"ref,attr"`
}

type sortableNodes []*Node

func (s sortableNodes) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}
func (s sortableNodes) Len() int {
	return len(s)
}
func (s sortableNodes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	fin := flag.String("i", "", "XML data source file")
	fout := flag.String("o", "result.dat", "Output file")
	help := flag.Bool("h", false, "Show usage")
	flag.Parse()
	if *fin == "" || *help {
		flag.Usage()
		return
	}
	f, err := os.Open(*fin)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	dat, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	result := &Osm{}
	err = xml.Unmarshal(dat, result)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	size := len(result.Nodes)

	graph := make([][]int8, size)
	for i := range graph {
		graph[i] = make([]int8, size)
	}
	sort.Sort(sortableNodes(result.Nodes))
	id2Index := make(map[int64]int)
	for i, v := range result.Nodes {
		id2Index[v.Id] = i
	}
	for _, v := range result.Ways {
		steps := len(v.Nds)
		for j := steps - 2; j >= 0; j-- {
			id_a := v.Nds[j].Ref
			id_b := v.Nds[j+1].Ref
			_a := id2Index[id_a]
			_b := id2Index[id_b]
			graph[_a][_b] = 1
			graph[_b][_a] = 1
		}
	}
	fmt.Println("Done ! Saving...")
	fff, err := os.Create(*fout)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	w := bufio.NewWriter(fff)
	totalLine := len(graph)
	progressBar := pb.StartNew(totalLine)
	for _, v := range graph {
		progressBar.Increment()
		s := strings.TrimLeft(fmt.Sprintf("%v", v), "[")
		s = strings.TrimRight(s, "]")
		w.WriteString(s)
	}
	w.Flush()
	fmt.Println("Everything is done")
}
