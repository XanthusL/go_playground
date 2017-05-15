package main

import "fmt"

// http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#blocked_goroutines

func main() {
	//sliceCorruption()
	fmt.Println(tmpFN())
}
func tmpFN() int {
	done := make(chan struct{})
	ret := make(chan int)
	defer close(done)
	f := func(i int) {
		select {
		case ret <- i:
			fmt.Println("hit")
		case <-done:
			fmt.Println("done")
		}
	}
	ints := []int{1, 2, 3, 4, 5, 6}
	for _, i := range ints {
		go f(i)
	}
	return <-ret
}

// Slice Data "Corruption"
// http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#slice_data_corruption
func sliceCorruption() {
	aa := make([]int, 0)
	aa = append(aa, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	/*
		This problem can fixed by allocating
		new slices and copying the data you need.
		Another option is to use the full slice expression.
	*/
	b2 := aa[:3:9]
	b1 := aa[:3]                  // Maybe unexpected result
	fmt.Println(len(b2), cap(b2)) // 3 9
	fmt.Println(len(b1), cap(b1)) // 3 10
	_ = append(b1, -1)            // makes dirty data in aa
	fmt.Println(aa)
}

// Multiple slices can reference the same data.
// This can happen when you create a new slice from an existing slice,
// for example. If your application relies on this behavior to function properly
// then you'll need to worry about "stale" slices.
//
// At some point adding data to one of the slices will result in a new array allocation
// when the original array can't hold any more new data.
// Now other slices will point to the old array (with old data).
func sliceStale() {
	s1 := []int{1, 2, 3}
	fmt.Println(len(s1), cap(s1), s1) //prints 3 3 [1 2 3]

	s2 := s1[1:]
	fmt.Println(len(s2), cap(s2), s2) //prints 2 2 [2 3]

	for i := range s2 {
		s2[i] += 20
	}

	//still referencing the same array
	fmt.Println(s1) //prints [1 22 23]
	fmt.Println(s2) //prints [22 23]

	s2 = append(s2, 4)

	for i := range s2 {
		s2[i] += 10
	}

	//s1 is now "stale"
	fmt.Println(s1) //prints [1 22 23]
	fmt.Println(s2) //prints [32 33 14]
}

// Arguments for a deferred function call are evaluated when
// the defer statement is evaluated (not when the function is actually executing).
func deferArgs() {
	var i int = 1

	defer fmt.Println("result =>", func() int { return i * 2 }())
	i++
	//prints: result => 2 (not ok if you expected 4)
}
