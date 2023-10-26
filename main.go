package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

// parse javascript object to golang struct
// we need exported fileds so change lowercase to uppercase
// so that fields are exported

type Words struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"words"`
}

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
	response, err := http.Get(args[1])
	if err != nil {
		log.Fatal(err)
	}

	// once this is done it automatically closes it
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != 200 {
		fmt.Printf("Invalid output (HTTP code %d):%s\n", response.StatusCode, body)
		os.Exit(1)
	}

	var words Words

	// use json to unmarshal json to struct
	err = json.Unmarshal(body, &words)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("JSON Parsed \nPage: %s \nWords :%s", words.Page, words.Words)
}
