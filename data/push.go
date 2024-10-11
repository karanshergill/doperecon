package data

import (
    "log"

    "doperecon/database"
    "github.com/spf13/cobra"
)

var (
    hostnames bool
    inputFile string
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
    Use:   "push",
    Short: "Push data to the database",
    Long:  `The push command allows you to push various types of data to the database.`,
    Run: func(cmd *cobra.Command, args []string) {
        // Validate flags
        if !hostnames {
            log.Fatal("You must specify the type of data with --hostnames (currently only hostnames are supported)")
        }

        if inputFile == "" {
            log.Fatal("You must provide an input file with --input-file")
        }
        
        // Call the InsertHostnames function from the database package
        err := database.PushHostnames(inputFile)
        if err != nil {
            log.Fatalf("Error inserting hostnames: %v", err)
        }

        log.Println("All hostnames processed successfully!")
    },
}

func init() {
    // Add the push command as a subcommand of the data command
    DataCmd().AddCommand(pushCmd)

    // Define flags for the push command
    pushCmd.Flags().BoolVar(&hostnames, "hostnames", false, "Specify that the data type is hostnames")
    pushCmd.Flags().StringVarP(&inputFile, "input-file", "i", "", "Path to the input file containing data")
}
