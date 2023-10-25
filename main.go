package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage ./hello-world <argument>\n")
		os.Exit(1)
	}
	fmt.Printf("Command line arguments: %v\n", args[1:])
}
