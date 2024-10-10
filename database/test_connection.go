package main

import (
    "fmt"

    "doperecon/database"
)

func main() {
    fmt.Println("Testing database connection...")

    // Initialize the database connection
    database.InitDB()
    defer database.CloseDB()

    fmt.Println("Database connection test successful!")
}
