// url.QueryEscape : URL-encode a string arg or pipe (stdin).
package main

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"
)

func main() {
	var input string

	if len(os.Args) > 1 {
		input = strings.Join(os.Args[1:], " ")
	} else {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
			os.Exit(1)
		}
		input = string(data)
	}

	encoded := url.QueryEscape(input)
	fmt.Println(encoded)
}
