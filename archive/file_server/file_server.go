package file_server

import (
	"net/http"

	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./")))
	s := &http.Server{Addr: ":1080", Handler: mux}
	go func() {
		fmt.Println("[Static file server] start to listen port 1080")
		if err := s.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Second)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	fmt.Printf("[Static file server] stopped by signal:%v\n", <-c)
	s.Shutdown(nil)
}
