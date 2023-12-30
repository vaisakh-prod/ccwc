package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var fileName string
var flag string
var text string

func init() {
	file := os.Stdin
	fi, err := file.Stat()
	if err != nil {
		fmt.Println("file.Stat()", err)
	}
	size := fi.Size()
	if size > 0 {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		text = string(stdin)
		if len(os.Args) == 2 {
			flag = os.Args[1]
		}
	} else {
		if len(os.Args) == 3 {
			flag = os.Args[1]
			fileName = os.Args[2]
		}
		if len(os.Args) == 2 {
			fileName = os.Args[1]
		}
	}
}

func bytesCount(text string) string {
	return fmt.Sprintf("  %v %s\n", len(text), fileName)
}

func linesCount(text string) string {
	count := strings.Count(text, "\n")
	return fmt.Sprintf("   %v %s \n", count, fileName)
}

func wordsCount(text string) string {
	count := len(strings.Fields(text))
	return fmt.Sprintf("  %v %s \n", count, fileName)
}

func charactersCount(text string) string {
	count := strings.Count(text, "")
	return fmt.Sprintf("  %v %s \n", count-1, fileName)
}

func defaultCount(text string) string {
	bString := bytesCount(text)
	lString := linesCount(text)
	wString := wordsCount(text)
	wString = strings.Trim(wString, "\n")
	wString = strings.Trim(wString, " ")
	wString = strings.Trim(wString, fileName)
	lString = strings.Trim(lString, "\n")
	lString = strings.Trim(lString, " ")
	lString = strings.Trim(lString, fileName)
	finalString := fmt.Sprintf("  %s %s %s", lString, wString, bString)
	return finalString
}

func main() {

	if fileName != "" {
		bText, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		text = string(bText)
	}
	var finalString string
	switch flag {
	case "-c":
		finalString = bytesCount(text)
	case "-l":
		finalString = linesCount(text)
	case "-w":
		finalString = wordsCount(text)
	case "-m":
		finalString = charactersCount(text)
	default:
		finalString = defaultCount(text)
	}
	fmt.Print(finalString)
}
