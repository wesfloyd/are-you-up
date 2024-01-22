// main.go

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: are-you-up <ip-address-or-domain>")
		os.Exit(1)
	}

	target := os.Args[1]
	fmt.Printf("Pinging %s...\n", target)

	err := runPing(target)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Ping completed.")
}

func runPing(target string) error {
	cmd := exec.Command("ping", "-c", "4", target)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("Failed to execute ping: %s", err)
	}

	fmt.Println("Ping output:")
	fmt.Println(strings.TrimSpace(string(output)))

	return nil
}
