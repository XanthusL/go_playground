package signal_handle

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	pid := os.Getpid()
	fmt.Println("Pid is", pid)

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			sig := <-sigs
			fmt.Println(sig)
			done <- true
		}
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
