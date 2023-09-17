/*
Copyright © 2023 Prince Kumar  <Princekrvert@gmail.com>

*/
package cmd

import (
	"crypto/sha1"
	"fmt"
	"log"
	"os"

	"github.com/princekrvert/doval/word"
	"github.com/spf13/cobra"
)

// sha1Cmd represents the sha1 command
var sha1Cmd = &cobra.Command{
	Use:   "sha1",
	Short: "sha1 Hash ",
	Long:  `Paste your sha1 hash here`,
	Run: func(cmd *cobra.Command, args []string) {
		//Now check the sha1 hash and for word file if given
		wordlist, _ := cmd.Flags().GetString("wordlist")
		if wordlist != "" {
			// now pass the path to the word package
			if len(args) != 1 {
				log.Panic("Please Enter a hash to decrypt or provide a hash one by one")
			}
			words, Nooflines := word.Eachword(wordlist)
			for index, word := range words {
				// now convert the word list to sha1 and compare with given hash ..
				hash := sha1.Sum([]byte(word))
				Hashstring := fmt.Sprintf("%x", hash)
				fmt.Printf("\033[31;m Word count %d::%d\n", Nooflines, index+1)
				if Hashstring == args[0] {
					// match found
					fmt.Printf("\033[32;1m Match found: ")
					fmt.Printf("\033[35;1m Hash is :: %s", word)
					os.Exit(0)
				} else {
					fmt.Printf("\033[33;1m Trying %s\n", word)

				}

			}
		} else {
			fmt.Println("use the default wordlist.")
			fmt.Println(wordlist)
		}
	},
}

func init() {
	rootCmd.AddCommand(sha1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	sha1Cmd.PersistentFlags().StringVarP(&Wordlist, "wordlist", "w", "/usr/share/wordlists/rockyou.txt", "Path to wordlist file")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//sha1Cmd.Flags().String(wordlist, "w", "Path of the wordlist file")
	//fmt.Println(sha1Cmd.Args)
}
