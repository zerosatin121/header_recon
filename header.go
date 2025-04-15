package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Define a command-line flag for the URL
	url := flag.String("url", "", "URL to check headers for")
	flag.Parse()

	// Check if the URL is provided
	if *url == "" {
		fmt.Println("Please provide a URL using the -url flag.")
		os.Exit(1)
	}

	// Make a HEAD request to the URL
	resp, err := http.Head(*url)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	// Print the response headers
	fmt.Println("Response Headers:")
	for key, values := range resp.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
}
