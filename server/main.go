package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// create a new router

	args := os.Args
	mux := http.NewServeMux()

	// Add a route for the get method
	mux.HandleFunc("/words", func(w http.ResponseWriter, r *http.Request) {
		// create a new json object .
		data := map[string]string{"page": "words", "input": args[1], "words": args[1]}

		// Encode the json object into a byte array
		encodedData, err := json.Marshal(data)

		if err != nil {
			fmt.Println(err)
			return
		}

		// Write the JSON object to the response writer
		// Respond with a JSON object with the status code 200 ok
		w.WriteHeader((http.StatusOK))
		w.Write(encodedData)

	})

	// Listen for requests on port 8080.
	http.ListenAndServe(":8080", nil)
}
