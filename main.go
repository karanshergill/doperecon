package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run script.go <file_path>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	// Run the external tool to detect the root domain
	cmd := exec.Command("dsieve", "-if", filePath, "-f", "2")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error detecting root domain: %v\n", err)
		os.Exit(1)
	}

	rootDomain := strings.TrimSpace(string(output))
	if rootDomain == "" {
		fmt.Println("Root domain could not be detected.")
		os.Exit(1)
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSuffix(line, "."+rootDomain)
		fmt.Println(trimmed)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
}
