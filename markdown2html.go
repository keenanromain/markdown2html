package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func readFile(arg string) []string {
	// see if there is a better/safer way to open a file, particularly when the files are in a subdirectory
	file, err := os.Open(arg)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	markdown := []string{}
	for scanner.Scan() {
		markdown = append(markdown, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return markdown
}

func printUsage() {
	fmt.Println("Testing...")
	os.Exit(0)
}

func validateArgs(args []string) string {
	if len(args) != 2 {
		printUsage()
	}
	extension := filepath.Ext(args[1])
	if extension != ".md" {
		printUsage()
	}
	return args[1][strings.Index(args[1], "/")+1 : len(args[1])-len(extension)]
}

func main() {
	args := os.Args
	fileName := validateArgs(args)
	markdown := readFile(args[1])
	fmt.Println(fileName, markdown)
}
