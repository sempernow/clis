// http.FileServer : List and serve files of the current working directory.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	addr := "8080"
	if len(os.Args) > 1 {
		addr = strings.Join(os.Args[1:2], "")
	}
	addr = "localhost:" + addr
	fmt.Println("@ http://" + addr)
	log.Fatal(http.ListenAndServe(addr, http.FileServer(http.Dir("."))))
}
