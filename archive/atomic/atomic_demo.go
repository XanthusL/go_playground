package atomic

import (
	"fmt"
	"sync/atomic"
	"time"
)

//var wg sync.WaitGroup

func RunTheTest() {
	demo1()
	demo2()

}

// 直接使用 ++
func demo1() {
	i := 0
	for j := 0; j < 1000; j++ {
		//wg.Add(1)
		go func() {
			i++
			//wg.Done()
		}()
	}
	//wg.Wait()
	time.Sleep(time.Second)
	fmt.Println("demo1:", i)
}

// 原子操作
func demo2() {
	var i int32 = 0
	for j := 0; j < 1000; j++ {
		go func() {
			atomic.AddInt32(&i, 1)
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("demo2:", i)
}
