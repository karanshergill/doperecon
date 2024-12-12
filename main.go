package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Parse command-line flags
	var filePath string
	flag.StringVar(&filePath, "list", "", "Path to the input file")
	flag.StringVar(&filePath, "l", "", "Path to the input file (short option)")
	flag.Parse()

	if filePath == "" {
		fmt.Println("Usage: go run script.go -list <file_path>")
		os.Exit(1)
	}

	// Run dsieve to detect the root domain
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

	// Create the directory named after the root domain
	if err := os.MkdirAll(rootDomain, 0755); err != nil {
		fmt.Printf("Error creating directory '%s': %v\n", rootDomain, err)
		os.Exit(1)
	}

	// Open the input file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Create the all-segments.txt file
	allSegmentsPath := rootDomain + "/all-segments.txt"
	allSegmentsFile, err := os.Create(allSegmentsPath)
	if err != nil {
		fmt.Printf("Error creating output file '%s': %v\n", allSegmentsPath, err)
		os.Exit(1)
	}
	defer allSegmentsFile.Close()

	// Write stripped hostnames to all-segments.txt
	scanner := bufio.NewScanner(file)
	allSegmentsWriter := bufio.NewWriter(allSegmentsFile)
	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSuffix(line, "."+rootDomain)
		if _, err := allSegmentsWriter.WriteString(trimmed + "\n"); err != nil {
			fmt.Printf("Error writing to output file: %v\n", err)
			os.Exit(1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	if err := allSegmentsWriter.Flush(); err != nil {
		fmt.Printf("Error flushing data to output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Hostnames have been successfully written to '%s'\n", allSegmentsPath)

	// Now read back the all-segments.txt file
	allSegmentsForReading, err := os.Open(allSegmentsPath)
	if err != nil {
		fmt.Printf("Error opening '%s' for reading: %v\n", allSegmentsPath, err)
		os.Exit(1)
	}
	defer allSegmentsForReading.Close()

	// Create single-segments.txt for writing unique segments
	singleSegmentsPath := rootDomain + "/single-segments.txt"
	singleSegmentsFile, err := os.Create(singleSegmentsPath)
	if err != nil {
		fmt.Printf("Error creating output file '%s': %v\n", singleSegmentsPath, err)
		os.Exit(1)
	}
	defer singleSegmentsFile.Close()

	singleSegmentsWriter := bufio.NewWriter(singleSegmentsFile)

	// Create combinations.txt for writing all contiguous segment combinations
	combinationsPath := rootDomain + "/combinations.txt"
	combinationsFile, err := os.Create(combinationsPath)
	if err != nil {
		fmt.Printf("Error creating output file '%s': %v\n", combinationsPath, err)
		os.Exit(1)
	}
	defer combinationsFile.Close()

	combinationsWriter := bufio.NewWriter(combinationsFile)

	// Track unique segments
	uniqueSegments := make(map[string]struct{})

	segmentScanner := bufio.NewScanner(allSegmentsForReading)
	for segmentScanner.Scan() {
		line := segmentScanner.Text()
		segments := strings.Split(line, ".")

		// Write unique segments
		for _, segment := range segments {
			segment = strings.TrimSpace(segment)
			if segment == "" {
				continue
			}
			if _, found := uniqueSegments[segment]; !found {
				if _, err := singleSegmentsWriter.WriteString(segment + "\n"); err != nil {
					fmt.Printf("Error writing segment to file: %v\n", err)
					os.Exit(1)
				}
				uniqueSegments[segment] = struct{}{}
			}
		}
		for start := 0; start < len(segments); start++ {
			for end := start + 1; end < len(segments)-1; end++ {
				combo := strings.Join(segments[start:end+1], ".")
				if combo == "" {
					continue
				}
				if _, err := combinationsWriter.WriteString(combo + "\n"); err != nil {
					fmt.Printf("Error writing combination to file: %v\n", err)
					os.Exit(1)
				}
			}
		}
	}

	if err := segmentScanner.Err(); err != nil {
		fmt.Printf("Error reading '%s': %v\n", allSegmentsPath, err)
		os.Exit(1)
	}

	// Flush the buffer for single-segments.txt
	if err := singleSegmentsWriter.Flush(); err != nil {
		fmt.Printf("Error flushing data to '%s': %v\n", singleSegmentsPath, err)
		os.Exit(1)
	}

	if err := combinationsWriter.Flush(); err != nil {
		fmt.Printf("Error flushing data to '%s': %v\n", combinationsPath, err)
		os.Exit(1)
	}

	fmt.Printf("Combinations have been successfully written to '%s'\n", combinationsPath)
	fmt.Printf("Unique segments have been successfully written to '%s'\n", singleSegmentsPath)
}
