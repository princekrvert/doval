/*
Copyright © 2023 Prince Kumar  <princekrvert@gmail.com>

*/
package cmd

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"

	"github.com/princekrvert/doval/word"
	"github.com/spf13/cobra"
)

var wordlist string

// md5Cmd represents the md5 command
var md5Cmd = &cobra.Command{
	Use:   "md5",
	Short: "This cmd will convert your string into md5 hash",
	Long:  `Enter a text to convert a text into md5 hash `,
	Run: func(cmd *cobra.Command, args []string) {
		wordlist, _ := cmd.Flags().GetString("wordlist")
		if wordlist != "" {
			words, Nooflines := word.Eachword(wordlist)
			for index, word := range words {
				hash := md5.Sum([]byte(word))
				Hashstring := fmt.Sprintf("%x", hash)
				fmt.Printf("\033[31;m Word count %d::%d\n", Nooflines, index+1)
				// check if args0 is provided
				if len(args) != 1 {
					log.Panic("Please Enter a hash to decrypt or provide a hash one by one")
				}
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
			fmt.Println("\033[31;1m Use default wordlist")
		}
	},
}

func init() {
	rootCmd.AddCommand(md5Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// md5Cmd.PersistentFlags().String("foo", "", "A help for foo")
	md5Cmd.PersistentFlags().StringVarP(&wordlist, "wordlist", "w", "/usr/share/wordlists/rockyou.txt", "Path to wordlist file")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// md5Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
