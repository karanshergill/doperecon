// Uses a goroutine to read the file line by line and emits each line of 
// the file as a string through the channel

package files

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// ReadFile returns a channel that emits each line from the file as a string
func ReadFile(filePath string) (<-chan string, error) {

	// Check file exists
	if !FileExists(filePath) {
        return nil, fmt.Errorf("File does not exist: %s", filePath)
    }

    // Open the file
    file, err := os.Open(filePath)
    if err != nil {
        return nil, fmt.Errorf("Error opening file: %v", err)
    }

    // Create a channel to send data
    dataChan := make(chan string)

    // Start a goroutine to read the file and send data to the channel
    go func() {
        defer file.Close()
        defer close(dataChan)

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            line := scanner.Text()
            dataChan <- line
        }

        // Handle any potential scanner error
        if err := scanner.Err(); err != nil {
            log.Printf("Error reading file: %v", err)
        }
    }()

    return dataChan, nil
}