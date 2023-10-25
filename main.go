package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage ./http-get <url>\n")
		os.Exit(1)
	}
	fmt.Printf("Command line arguments: %v\n", args[1:])

	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Printf("URL is in invalid format: %s\n", err)
		os.Exit(1)
	}

}
