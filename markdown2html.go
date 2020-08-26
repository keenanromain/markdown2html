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
	fileName := filepath.Base(args[1])
	return fileName[0 : len(fileName)-len(extension)]
}

func findTagType(line string) string {
	if strings.HasPrefix(line, "#") {
		return fmt.Sprintf("h%d", strings.Count(line, "#"))
	} else if len(line) > 0 {
		return "p"
	} else {
		return "</ br>"
	}
}

func createHTMLcontent(line string) string {
	tag := findTagType(line)
	if tag == "</ br>" {
		return tag
	}
	return fmt.Sprintf("<%s>%s</%s>", tag, line, tag)
}

func createHTMLwrapper(fileName string, markdown []string) string {
	html := `<!DOCTYPE html><html><head><meta charset="utf-8" name="viewport">
		<link href="https://fonts.googleapis.com/css?family=Roboto" rel="stylesheet">
		<title>` + filepath.Base(fileName) + `</title></head><body>`
	for _, line := range markdown {
		html += createHTMLcontent(line)
	}
	return html + "</body></html>"
}

func createFile(fileName string, markdown []string) {
	if _, err := os.Stat("output"); os.IsNotExist(err) {
		os.Mkdir("output", 0755)
	}
	fileName = fmt.Sprintf("%s/%s.%s", "output", fileName, "html")
	if _, err := os.Stat(fileName); err == nil {
		os.Remove(fileName)
	}
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	html := createHTMLwrapper(fileName, markdown)
	f.WriteString(html)
	fmt.Println(fmt.Sprintf("Finished! Your new HTML file can be found in %s", fileName))
}

func main() {
	args := os.Args
	fileName := validateArgs(args)
	markdown := readFile(args[1])
	createFile(fileName, markdown)
}
