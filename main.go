// main.go

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: my-web-pinger <url>")
		os.Exit(1)
	}

	url := os.Args[1]
	fmt.Printf("Pinging %s...\n", url)

	err := pingURL(url)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("URL is live!")
}

func pingURL(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Received non-OK status code: %d", resp.StatusCode)
	}

	return nil
}
