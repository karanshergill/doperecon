// run.go

package cmd

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"sync"

	"github.com/karanshergill/dope-recon/pkg"
	"github.com/karanshergill/dope-recon/utils"
	"github.com/spf13/cobra"
)

var domain string
var domainList string
var wordlist string

// Define the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Gather, Generate and BruteForce Subdomains.",
	Run: func(cmd *cobra.Command, args []string) {

		// Ensure either domain or domainList is provided, but not both
		if domain == "" && domainList == "" {
			log.Fatal("Error: You must provide either a domain (-d) or a domain list (-dl).")
		}
		if domain != "" && domainList != "" {
			log.Fatal("Error: You cannot provide both a single domain (-d) and a domain list (-dl). Choose one.")
		}

		// Load wordlist
		if wordlist == "" {
			// Set default wordlist path
			homeDir, err := os.UserHomeDir()
			if err != nil {
				log.Fatalf("Error fetching user home directory: %v", err)
			}
			wordlist = homeDir + "/wordlists/megadns.txt"
		}
		if !utils.FileExists(wordlist) {
			log.Fatalf("Error: Wordlist file '%s' not found.\n", wordlist)
		}

		subdomains, err := generator.LoadWordlist(wordlist)
		if err != nil {
			log.Fatalf("Error loading wordlist: %v", err)
		}

		// Process domains either from a list or single domain
		var wg sync.WaitGroup
		if domain != "" {
			wg.Add(1)
			go generator.ProcessDomain(domain, subdomains, &wg)
		} else if domainList != "" {
			if !utils.FileExists(domainList) {
				log.Fatalf("Error: Domain list file '%s' not found.\n", domainList)
			}

			// Open and read the domain list file
			file, err := os.Open(domainList)
			if err != nil {
				log.Fatalf("Error opening domain list file: %v", err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				wg.Add(1)
				go generator.ProcessDomain(scanner.Text(), subdomains, &wg)
			}
			if err := scanner.Err(); err != nil {
				log.Fatalf("Error reading domain list: %v", err)
			}
		}

		wg.Wait()
		fmt.Println("Subdomain generation completed.")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&domain, "domain", "d", "", "Domain")
	runCmd.Flags().StringVarP(&domainList, "domainlist", "l", "", "Domain list file")
	runCmd.Flags().StringVarP(&wordlist, "wordlist", "w", "", "Wordlist file")
}
