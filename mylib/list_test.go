package mylib

import (
	"fmt"
	"testing"
)

func TestArrayList_RemoveAt(t *testing.T) {
	l := &ArrayList{}

	l.Init()
	l.PushBack("a")
	l.PushBack("b")
	l.PushBack("c")
	l.PushBack("d")

	length := l.Len()
	e := l.List.Front()
	for i := 0; i < length; i++ {
		fmt.Println(e.Value)
		e = e.Next()
	}
	fmt.Println(l.RemoveAt(-1))
	fmt.Println(l.RemoveAt(100))
	fmt.Println(l.RemoveAt(2))
	length = l.Len()
	e = l.List.Front()
	for i := 0; i < length; i++ {
		fmt.Println(e.Value)
		e = e.Next()
	}
}
