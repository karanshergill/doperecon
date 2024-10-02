package utils

import (
	"fmt"
	"os"
	"os/exec"
)

// CheckCommand checks if the given command is available in the system's PATH.
func CheckCommand(cmd string, name string) {
	_, err := exec.LookPath(cmd)
	if err != nil {
		fmt.Printf("Error: %s is not installed or not available in the system's PATH.\n", name)
		fmt.Printf("Please install %s and ensure it is accessible.\n", name)
		os.Exit(1)
	}
}
