package utils

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func LoadDomains(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening domain list file: %v", err)
	}
	defer file.Close()

	var domains []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domains = append(domains, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading domain list file: %v", err)
	}

	return domains, nil
}

// SortAndUniq reads a file, removes duplicates, sorts the lines, and writes the result to a new file.
func SortAndUniq(inputFile, outputFile string) error {
	// Open the input file
	file, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("error opening file %s: %v", inputFile, err)
	}
	defer file.Close()

	// Use a map to filter out duplicates
	uniqueLines := make(map[string]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		uniqueLines[line] = true
	}

	// Check for scanner error
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file %s: %v", inputFile, err)
	}

	// Convert map keys to a slice for sorting
	lines := make([]string, 0, len(uniqueLines))
	for line := range uniqueLines {
		lines = append(lines, line)
	}

	// Sort the slice
	sort.Strings(lines)

	// Write the sorted unique lines to the output file
	output, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("error creating file %s: %v", outputFile, err)
	}
	defer output.Close()

	for _, line := range lines {
		_, err := output.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("error writing to file %s: %v", outputFile, err)
		}
	}

	return nil
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}