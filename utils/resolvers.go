package utils

import (
	"fmt"
	"os"
	"os/exec"
)

// generate the resolver and trusted resolver files using publicresolvers
func GenerateResolvers() {
	resolversFile := "resolvers.txt"
	trustedResolversFile := "trusted_resolvers.txt"

	fmt.Println("Generating resolvers...")

	// generate resolvers.txt
	err := runCommand("publicresolvers", "--resolvers", resolversFile)
	if err != nil {
		fmt.Printf("Error generating resolvers: %v\n", err)
		os.Exit(1)
	}

	// generate trusted_resolvers.txt
	err = runCommand("publicresolvers", "--trusted", trustedResolversFile)
	if err != nil {
		fmt.Printf("Error generating trusted resolvers: %v\n", err)
		os.Exit(1)
	}

	// check if files are non-empty
	if isFileNonEmpty(resolversFile) && isFileNonEmpty(trustedResolversFile) {
		fmt.Println("Resolvers generated successfully.")
	} else {
		fmt.Println("Error: Failed to generate resolvers!")
		os.Exit(1)
	}
}

// Helper function to run a command and redirect output to a file.
func runCommand(command string, arg string, outputFile string) error {
	cmd := exec.Command(command, arg)
	outfile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", outputFile, err)
	}
	defer outfile.Close()

	cmd.Stdout = outfile
	return cmd.Run()
}

// Helper function to check if a file is non-empty.
func isFileNonEmpty(filePath string) bool {
	info, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("Error checking file: %v\n", err)
		return false
	}
	return info.Size() > 0
}
