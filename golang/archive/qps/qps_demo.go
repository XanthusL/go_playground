package qps

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync/atomic"
	"time"
)

//---------------------------------------------
//		Server
//---------------------------------------------
var (
	intFlag   int32 // 标记统计的开始和结束状态
	totalCnt  int32 // 请求数
	startTime int64 // 开始时间
	endTime   int64 // 结束时间
)

func main() {
	http.HandleFunc("/test", handler4Test)
	http.ListenAndServe(":1234", nil)
}

func handler4Test(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	f := req.Form.Get("flag")
	if f == "1" {
		atomic.StoreInt32(&intFlag, 1)
		atomic.StoreInt64(&startTime, time.Now().UnixNano())
	} else if f == "0" {
		ok := atomic.CompareAndSwapInt32(&intFlag, 1, 0)
		if ok {
			atomic.StoreInt64(&endTime, time.Now().UnixNano())
			fmt.Fprintf(w, "Total served: %d\n", totalCnt)
			d := endTime - startTime
			fmt.Fprintf(w, "Use Time: %dns\n", d)
			if totalCnt != 0 {
				fmt.Fprintf(w, "%dns/op\n", d/int64(totalCnt))
				fmt.Fprintf(w, "%.2fop/s\n", float32(totalCnt)*1e9/(float32(d)))
			}
		}
	}

	// do something
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)*10))

	if atomic.LoadInt32(&intFlag) == 1 {
		atomic.AddInt32(&totalCnt, 1)
	}
}

//---------------------------------------------
//		Client
//---------------------------------------------
/*
package main

import (
	"sync"
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"
)

var wg sync.WaitGroup

func main() {
	for i := 1; i < 2000; i++ {
		go qpsTest(strconv.Itoa(i))
	}
	wg.Wait()
	qpsTest(strconv.Itoa(0))
}

var qps chan struct{} = make(chan struct{}, 500)

func qpsTest(f string) {
	qps <- struct{}{}
	wg.Add(1)
	defer func() {
		<-qps
		wg.Add(-1)
	}()

	resp, err := http.Get("http://localhost:1234/test?flag=" + f)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(string(data))
		}

	}
}
*/
