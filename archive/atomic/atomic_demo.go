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
	// Result 0
	// ++ directly: 9113 time cost(nano): 7572762
	// atomic: 10000 time cost(nano): 1814050
	// Lock: 10000 time cost(nano): 2309760

	// Result 1
	// ++ directly: 8879 time cost(nano): 2337648
	// atomic: 10000 time cost(nano): 1960095
	// Lock: 10000 time cost(nano): 2103779

	// Result 2
	// ++ directly: 8716 time cost(nano): 2314949
	// atomic: 10000 time cost(nano): 5774471
	// Lock: 10000 time cost(nano): 2461061

	// Result 3
	// ++ directly: 9155 time cost(nano): 2231489
	// atomic: 10000 time cost(nano): 1949282
	// Lock: 10000 time cost(nano): 3570333

	// Result 4
	// ++ directly: 9999 time cost(nano): 1496126
	// atomic: 10000 time cost(nano): 6082791
	// Lock: 10000 time cost(nano): 6115087

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
