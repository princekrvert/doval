package word

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Eachword(filepath string) []string {
	words := []string{}
	// Now check if the path of file is valid
	_, err := os.Stat(filepath)
	if err != nil {
		log.Fatal("\033[31;1m Invalid path, Please provide the correct path")
		os.Exit(1)
	}
	// Now open the file to read
	readFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	defer readFile.Close()

	for _, line := range fileLines {
		words = append(words, line)
	}

	return words
}
