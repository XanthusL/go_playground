package main

import (
	"os"
	"math"
	"io"
	"fmt"
	"crypto/sha1"
	"path/filepath"
	"flag"
	"regexp"
	"sync"
	"bufio"
	"encoding/json"
)

const cBUFFER_SIZE = 1024 * 8

type HashInfo struct {
	Name string `json:"name"`
	Sha1 string `json:"sha1"`
	Size int64 `json:"size"`
	Err  error `json:"err"`
}
// 用来处理获取到的信息
type HashInfoHandler func(*HashInfo)

var wait sync.WaitGroup
var bufLck sync.Mutex // todo 无锁数据结构
func main() {
	// 获取参数
	outArg := flag.String("out", "", "The file to save to. Prints to standard if not assigned.")
	pathArg := flag.String("path", ".", "The path to scan. Current path as default.")
	helpArg := flag.Bool("help", false, "Show this usage.")
	fmtArg := flag.String("format", "line", "Out put format. \"line\" and \"json\" supported.")
	filterArg := flag.String("filter", "", "Regex of files to skip")
	flag.Parse()
	// 显示帮助信息
	if *helpArg {
		flag.Usage()
		return
	}
	// 过滤参数的正则
	var rgxp *regexp.Regexp
	if *filterArg != "" {
		var err error
		rgxp, err = regexp.Compile(*filterArg)
		// 正则有误，显示帮助信息
		if err != nil {
			fmt.Println(err.Error())
			flag.Usage()
			return
		}
	}
	if *fmtArg != "line"  && *fmtArg != "json" {
		flag.Usage()
		return
	}
	isJson := *fmtArg == "json"
	//fmt.Println(*outArg, *pathArg)
	var w *bufio.Writer
	if *outArg == "" {
		// 未指定输出，直接打印结果
		WalkFiles(*pathArg, func(info *HashInfo) {
			s := toString(isJson, info)
			fmt.Println(s)
		}, rgxp)
	} else {
		f, err := os.Create(*outArg)
		if err != nil {
			fmt.Println(err.Error())
			flag.Usage()
			return
		}
		defer f.Close()
		// 使用缓冲，减少IO操作次数
		w = bufio.NewWriter(f)
		WalkFiles(*pathArg, func(info *HashInfo) {
			s := toString(isJson, info)
			bufLck.Lock()
			defer bufLck.Unlock()
			w.WriteString(s + "\n")
			i := w.Buffered()
			if i >= cBUFFER_SIZE {
				w.Flush()
			}
		}, rgxp)

	}
	wait.Wait()

	if w != nil && w.Buffered() > 0 {
		w.Flush()
	}
}
func toString(isJson bool, info *HashInfo) string {
	var s string
	if isJson {
		data, _ := json.Marshal(info)
		s = string(data)
	} else {
		s = fmt.Sprintf("%s,%s,%d", info.Name, info.Sha1, info.Size)
		if info.Err != nil {
			s = fmt.Sprintf("%s,%s", s, info.Err.Error())
		}
	}
	return s
}
// 遍历
func WalkFiles(filename string, handler HashInfoHandler, rgxp *regexp.Regexp) {
	filepath.Walk(filename, func(path string, info os.FileInfo, err error) error {
		if ( info == nil ) {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if rgxp != nil && rgxp.MatchString(path) {
			return nil
		}
		// 瓶颈: IO? 计算?
		go GetInfo(path, handler)
		return nil
	})
}

// 获取文件的大小, sha1等信息
func GetInfo(filename string, handler HashInfoHandler) HashInfo {
	wait.Add(1)
	defer wait.Done()

	h := HashInfo{Name:filename}
	defer handler(&h)
	f, err := os.Open(filename)
	if err != nil {
		h.Err = err
		return h
	}
	defer f.Close()

	hash := sha1.New()
	info, err := f.Stat()
	if err != nil {
		h.Err = err
		return h
	}
	h.Size = info.Size()
	blocks := uint64(math.Ceil(float64(h.Size) / float64(cBUFFER_SIZE)))

	for i := uint64(0); i < blocks; i++ {
		blockSize := int(math.Min(cBUFFER_SIZE, float64(h.Size - int64(i * cBUFFER_SIZE))))
		buf := make([]byte, blockSize)
		f.Read(buf)
		io.WriteString(hash, string(buf))
	}
	h.Sha1 = fmt.Sprintf("%x", hash.Sum(nil))
	return h
}
