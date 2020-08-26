package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var compiledLinkRegex = regexp.MustCompile(`\[[^][]+]\((https?://[^()]+)\)`)
var compiledItalicsRegex = regexp.MustCompile(`\**(?:^|[^*])(\*(\w+(\s\w+)*)\*)`)
var compiledBoldRegex = regexp.MustCompile(`\**(?:^|[^*])(\*\*(\w+(\s\w+)*)\*\*)`)

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
	usage := `

	To execute this program successfully, be sure to specify the full path of an MD file as an argument.
	There are some located inside the input directory of this project.
	Otherwise, provide a different MD as long as you specify the full path to it as well.
	
	E.g.:
	
		1.) go run markdown2html.go input/sample1.md
		
		2.) go build markdown2html.go
		   ./markdown2html /path/to/your/input/file.md	
		
	`
	println(usage)
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

func modifyLink(line, md, url string) string {
	modification := fmt.Sprintf("<a href=\"%s\">%s</a>", url, md[1:strings.IndexByte(md, ']')])
	return strings.Replace(line, md, modification, 1)
}

func searchForLinks(line string) string {
	links := compiledLinkRegex.FindAllStringSubmatch(line, -1)
	if len(links) > 0 {
		for link := range links {
			_, err := url.ParseRequestURI(links[link][1])
			if err != nil {
				log.Fatal(err)
			}
			line = modifyLink(line, links[link][0], links[link][1])
		}
	}
	return line
}

func modifyBoldOrItalics(line, md, tag string) string {
	modification := fmt.Sprintf("<%s>%s</%s>", tag, strings.Trim(md, "*"), tag)
	return strings.Replace(line, md, modification, 1)
}

func searchForBoldOrItalics(line string) string {
	bolds := compiledBoldRegex.FindAllStringSubmatch(line, -1)
	if len(bolds) > 0 {
		for bold := range bolds {
			line = modifyBoldOrItalics(line, bolds[bold][1], "strong")
		}
	}
	italics := compiledItalicsRegex.FindAllStringSubmatch(line, -1)
	if len(italics) > 0 {
		for italic := range italics {
			line = modifyBoldOrItalics(line, italics[italic][1], "em")
		}
	}
	return line
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
	if strings.HasPrefix(tag, "h") {
		line = strings.Trim(line, "#")
	}
	line = searchForLinks(line)
	line = searchForBoldOrItalics(line)
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
