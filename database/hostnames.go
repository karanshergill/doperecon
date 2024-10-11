package database

import (
    "context"
    "fmt"
    "log"

    "doperecon/utils/files"
)

// InsertHostnames reads hostnames from a file and inserts them into the database
func PushHostnames(filePath string) error {
    // Initialize the database connection
    InitDB()
    defer CloseDB()

    // Get the data channel from the files module
    dataChan, err := files.ReadFile(filePath)
    if err != nil {
        return fmt.Errorf("error reading file: %v", err)
    }

    // Consume the data channel and write each item to the database
    for line := range dataChan {
        if line != "" {
            _, err := Pool.Exec(context.Background(), "INSERT INTO hostnames (label) VALUES ($1)", line)
            if err != nil {
                log.Printf("Error inserting label '%s': %v", line, err)
            } else {
                fmt.Printf("Inserted label: %s\n", line)
            }
        }
    }

    fmt.Println("All hostnames processed successfully!")
    return nil
}
