package main

import "fmt"

func aaa(a int32) int32 {
	return a << 2
}

func main() {
	if 8 >= aaa(2) {
		fmt.Println("asdf")
	}

}
