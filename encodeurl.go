package main

import (
	"fmt"
	"net/url"
	"os"
	"io"
	"strings"
)

func main() {
	var input string

	// Read from stdin (piped input) or arguments
	if len(os.Args) > 1 {
		input = strings.Join(os.Args[1:], " ")
	} else {
		fmt.Fprint(os.Stderr, "Reading from stdin...\n")
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
			os.Exit(1)
		}
		input = string(data)
	}

	// URL-encode the input
	encoded := url.QueryEscape(input)
	fmt.Println(encoded)
}
