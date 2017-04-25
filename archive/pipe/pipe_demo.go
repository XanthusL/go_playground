package pipe

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Pipe in memory provided in standard library.
func ramPipeDemo() {
	pr, pw := io.Pipe()
	go func() {
		rr := bufio.NewReader(pr)
		for {
			var b = make([]byte, 50)
			rr.Read(b)
			fmt.Println(string(b))
		}
		pr.Close()

	}()
	for {
		var i string
		fmt.Scanln(&i)
		if i == "" {
			pw.Close()
			break
		}
		pw.Write([]byte(i))
	}
}

// System level pipe-file based pipe
// eg.
// $ ps -aux | grep nginx
// $ mkfifo myPipe
func osPipeDemo() {
	r, w, e := os.Pipe()
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	go func() {
		rr := bufio.NewReader(r)
		for {
			var b = make([]byte, 50)
			rr.Read(b)
			fmt.Println(string(b))
		}
		r.Close()

	}()
	for {
		var i string
		fmt.Scanln(&i)
		if i == "" {
			r.Close()
			break
		}
		w.Write([]byte(i))
	}
}
