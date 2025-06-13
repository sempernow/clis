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
	u := "8080"
	if len(os.Args) > 1 {
		u = strings.Join(os.Args[1:2], "")
	}
	u = "localhost:" + u
	fmt.Println("@ http://"+u)
	log.Fatal(http.ListenAndServe(u, http.FileServer(http.Dir("."))))
}
