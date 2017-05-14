package main

import (
	"flag"
	"fmt"
)

var version string

func main() {
	showVersion := flag.Bool("version", false, "show version info")
	showHelp := flag.Bool("help", false, "show usage")
	flag.Parse()
	if *showVersion {
		fmt.Println(version)
	}
	if *showHelp {
		flag.Usage()
	}
	aa := make([]int, 0,20)
	bb := make([]int, 0,2)
	aa = append(aa, 1)
	aa = append(aa, 2, 3, 4)
	if len(aa) < 20 {
		fmt.Println(append(bb,aa[:20]...))
	}
}
