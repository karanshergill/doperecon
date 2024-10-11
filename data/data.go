package data

import (
    "github.com/spf13/cobra"
)

// dataCmd represents the data command
var dataCmd = &cobra.Command{
    Use:   "data",
    Short: "Perform operations on data such as inserting or updating information",
    Long:  `The data command allows you to perform operations on data, including inserting hostnames, updating records, and more.`,
}

// DataCmd returns the dataCmd to be used in other parts of the application
func DataCmd() *cobra.Command {
    return dataCmd
}
