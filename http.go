// http servev files from PWD.
// $ http
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
	fmt.Println("@ ", u)
	log.Fatal(http.ListenAndServe(u, http.FileServer(http.Dir("."))))
}
