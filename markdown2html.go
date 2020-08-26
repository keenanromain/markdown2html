package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile() []string {
	// see if there is a better/safer way to open a file, particularly when the files are in a subdirectory
	file, err := os.Open("./input/headerTags.md")
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

func main() {
	markdown := readFile()
	fmt.Println(markdown)
}
