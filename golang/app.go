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
		return
	}
	if *showHelp {
		flag.Usage()
		return
	}
}
