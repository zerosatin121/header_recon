package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/fatih/color"
)

//	func printBanner() {
//		c := color.New(color.FgCyan).Add(color.Bold)
//		c.Println(`
//	  _   _            _               ____
//	 | | | | ___  _ __| |_ ___ _ __   |  _ \ ___  ___ ___  _ __ ___
//	 | |_| |/ _ \| '__| __/ _ \ '__|  | |_) / _ \/ __/ _ \| '_ ' _ \
//	 |  _  | (_) | |  | ||  __/ |     |  _ <  __/ (_| (_) | | | | | |
//	 |_| |_|\___/|_|   \__\___|_|     |_| \_\___|\___\___/|_| |_| |_|
//	                          Header Recon Tool - by ZeroSatin121
//		`)
//	}
func printBanner() {
	cyan := color.New(color.FgCyan).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	boldWhite := color.New(color.FgWhite, color.Bold).SprintFunc()

	fmt.Println(cyan("  _   _            _               ____                         "))
	fmt.Println(magenta(" | | | | ___  _ __| |_ ___ _ __   |  _ \\ ___  ___ ___  _ __ ___ "))
	fmt.Println(green(" | |_| |/ _ \\| '__| __/ _ \\ '__|  | |_) / _ \\/ __/ _ \\| '_ ` _ \\"))
	fmt.Println(magenta(" |  _  | (_) | |  | ||  __/ |     |  _ <  __/ (_| (_) | | | | | |"))
	fmt.Println(cyan(" |_| |_|\\___/|_|   \\__\\___|_|     |_| \\_\\___|\\___\\___/|_| |_| |_|"))
	fmt.Println(boldWhite("\n                  Header Recon Tool - by ZeroSatin121\n"))
}
func main() {
	printBanner()
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
