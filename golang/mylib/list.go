package mylib

import "container/list"

// A Java-like list struct
// just a demo, can't ensure performance,
// using slice instead of list.List might be better
// TODO : 1. A slice implementation of ArrayList. 2. Benchmark
type ArrayList struct {
	list.List
}

func (a *ArrayList) Add(v interface{}) {
	a.List.PushBack(v)
}
func (a *ArrayList) RemoveAt(index int) interface{} {
	l := a.List.Len()
	if index < 0 || index >= l {
		return nil
	}
	if index < l/2 {
		e := a.List.Front()
		for i := 0; i < index; i++ {
			e = e.Next()
		}
		return a.List.Remove(e)

	} else {
		e := a.List.Back()
		t := l - index - 1
		for i := 0; i < t; i++ {
			e = e.Prev()
		}
		return a.List.Remove(e)
	}
}
func (a *ArrayList) Get(index int) interface{} {
	l := a.List.Len()
	if index < 0 || index >= l {
		return nil
	}
	if index < l/2 {
		e := a.List.Front()
		for i := 0; i < index; i++ {
			e = e.Next()
		}
		return e.Value

	} else {
		e := a.List.Back()
		t := l - index - 1
		for i := 0; i < t; i++ {
			e = e.Prev()
		}
		return e.Value
	}
}
func (a *ArrayList) IndexOf(e *list.Element) int {
	l := a.Len()
	if l == 0 {
		return -1
	}
	tmp := a.Front()
	for i := 0; i < l; i++ {
		if tmp == e {
			return i
		}
		tmp = tmp.Next()

	}
	return -1
}
func (a *ArrayList) LastIndexOf(e *list.Element) int {
	l := a.Len()
	if l == 0 {
		return -1
	}
	tmp := a.Back()
	for i := l - 1; i >= 0; i-- {
		if tmp == e {
			return i
		}
		tmp = tmp.Prev()
	}
	return -1
}
func (a *ArrayList) ToSlice() []interface{} {
	l := a.Len()
	if l == 0 {
		return make([]interface{}, 0)
	}
	slc := make([]interface{}, l)
	tmp := a.Front()
	for i := 0; i < l; i++ {
		slc[i] = tmp.Value
		tmp = tmp.Next()
	}
	return slc
}
