package file_server

import (
	"net/http"

	"fmt"
)

func main() {
	fmt.Println("[Static file server] start")
	http.Handle("/", http.FileServer(http.Dir("./")))
	fmt.Println(http.ListenAndServe(":1080", nil))
	fmt.Println("[Static file server] stop")

}
