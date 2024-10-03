package cmd

import (
	"log"
	"dope-recon/services"
	"dope-recon/utils"

	"github.com/spf13/cobra"
)

var domainList string
var wordlist string

// Define the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the domain fuzzing process",
	Run: func(cmd *cobra.Command, args []string) {
		if domainList == "" {
			log.Fatal("Error: Domain list not provided.")
		}
		if !utils.FileExists(domainList) {
			log.Fatalf("Error: Domain list file '%s' not found.\n", domainList)
		}
		if wordlist == "" {
			wordlist = "/home/superuser/wordlists/megadns.txt"
		}
		if !utils.FileExists(wordlist) {
			log.Fatalf("Error: Wordlist file '%s' not found.\n", wordlist)
		}

		// Call the service to handle domain fuzzing
		err := services.FuzzDomains(domainList, wordlist)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&domainList, "domain", "d", "", "Domain list file")
	runCmd.Flags().StringVarP(&wordlist, "wordlist", "w", "", "Wordlist file")
}
