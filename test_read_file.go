package main

import (
    "fmt"
    "log"

    "doperecon/utils/files"
)

func main() {
    // Define the path to a test file
    filePath := "test_domains.txt"

    // Call the ReadFile function from reader.go
    dataChan, err := files.ReadFile(filePath)
    if err != nil {
        log.Fatalf("Error reading file: %v", err)
    }

    // Iterate over the data channel and print each line
    for line := range dataChan {
        fmt.Printf("Read line: %s\n", line)
    }

    fmt.Println("Finished reading all lines.")
}
