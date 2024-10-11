package database

import (
    "testing"
)

func TestConnection(t *testing.T) {
    InitDB()
    defer CloseDB()
}
