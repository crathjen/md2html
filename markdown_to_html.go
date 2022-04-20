package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var isHeader *regexp.Regexp = regexp.MustCompile("^(#+) ")

var isLink *regexp.Regexp = regexp.MustCompile(`\[(.+?)\]\((.+?)\)`)


func ConvertMarkdown2HTML(reader io.Reader, writer io.Writer) error{

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		io.WriteString(writer, ConvertLineMarkdown2HTML(line) + "\n")
	}
	return nil
}

func ConvertLineMarkdown2HTML(line string) string {


	line = isLink.ReplaceAllStringFunc(line, func(s string) string {
		subgroups := isLink.FindStringSubmatch(s)
		return fmt.Sprintf(`<a href="%s">%s</a>`, subgroups[2], subgroups[1])
	})

	matchIndex := isHeader.FindStringIndex(line)
	if matchIndex == nil {
		return WrapLineWithTag(line, "p")
	}else {
		return WrapLineWithTag(line[matchIndex[1]:], "h"+ strconv.Itoa(matchIndex[1]-1))
	}
}


func WrapLineWithTag(line string, tag string) string {
	return fmt.Sprintf(`<%s>%s</%s>`, tag, line, tag)
}

func main() {
	out := bufio.NewWriter(os.Stdout)
	err := ConvertMarkdown2HTML(os.Stdin, out)
	if err != nil {
		fmt.Printf("conversion error: %v", err.Error())
	}
	err = out.Flush()
	if err != nil {
		fmt.Printf("flush error: %v", err.Error())
	}
}

//[Link text](https://www.example.com)