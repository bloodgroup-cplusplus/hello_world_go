package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Printf("Command line arguments: %v\n", args[1:])
}
