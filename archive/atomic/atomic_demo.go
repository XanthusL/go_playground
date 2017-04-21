package atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup

func RunTheTest() {
	demo1() // Wrong result
	demo2() // Correct
	demo3() // Correct
	// Result
	//++ directly: 9113 time cost(nano): 7572762
	//atomic: 10000 time cost(nano): 1814050
	//Lock: 10000 time cost(nano): 2309760

}

// 直接使用 ++
func demo1() {
	i := 0
	nano := time.Now().UnixNano()
	for j := 0; j < 10000; j++ {
		wg.Add(1)
		go func() {
			i++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("++ directly:", i, "time cost(nano):", time.Now().UnixNano()-nano)
}

// 原子操作
func demo2() {
	var i int32 = 0
	nano := time.Now().UnixNano()
	for j := 0; j < 10000; j++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&i, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("atomic:", i, "time cost(nano):", time.Now().UnixNano()-nano)
}

// 锁
func demo3() {
	lck := sync.Mutex{}
	var i int32 = 0
	nano := time.Now().UnixNano()
	for j := 0; j < 10000; j++ {
		wg.Add(1)
		go func() {
			lck.Lock()
			i++
			lck.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Lock:", i, "time cost(nano):", time.Now().UnixNano()-nano)
}
