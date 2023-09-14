/*
Copyright © 2023 Prince Kumar  <Princekrvert@gmail.com>

*/
package cmd

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/princekrvert/doval/word"
	"github.com/spf13/cobra"
)

var Wordlist string

// sha256Cmd represents the sha256 command
var sha256Cmd = &cobra.Command{
	Use:   "sha256",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//Now check the sha256 hash and for word file if given
		wordlist, _ := cmd.Flags().GetString("wordlist")
		if wordlist != "" {
			// now pass the path to the word package
			words := word.Eachword(wordlist)
			for _, word := range words {
				// now convert the word list to sha256 and compare with given hash ..
				hash := sha256.Sum256([]byte(word))
				Hashstring, err := fmt.Printf("%x", hash)
				fmt.Print("\033[H\033[2J")
				if err != nil {
					log.Fatal("Somthing went wrong,Please try again")
				} else {
					// check if the word maches the hash
					if args[0] == strconv.Itoa(Hashstring) {
						fmt.Println("\033[32;1m Match found")
						fmt.Printf("%s Hash : String %d", args[0], Hashstring)
						os.Exit(0)
					} else {
						fmt.Printf("Trying %s : ", word)
						fmt.Println(Hashstring)
					}
				}

			}
		} else {
			fmt.Println("use the default wordlist.")
			fmt.Println(wordlist)
		}
	},
}

func init() {
	rootCmd.AddCommand(sha256Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	sha256Cmd.PersistentFlags().StringVarP(&Wordlist, "wordlist", "w", "/usr/share/wordlists/rockyou.txt", "Path to wordlist file")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//sha256Cmd.Flags().String(wordlist, "w", "Path of the wordlist file")
	//fmt.Println(sha256Cmd.Args)
}
