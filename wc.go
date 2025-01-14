package main

import (
	"flag"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func getLines(data []byte) int {
	lines := 0
	for _, char := range data {
		if char == '\n' {
			lines++
		}
	}
	return lines
}

func getWords(data []byte) int {
	words := 0
	inWord := false

	for _, char := range data {
		if unicode.IsSpace(rune(char)) {
			inWord = false
		} else if !inWord {
			inWord = true
			words++
		}
	}
	return words
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage <command> <command> <..> <filename>")
		return
	}

	l := flag.Bool("l", false, "count lines")
	w := flag.Bool("w", false, "count words")
	m := flag.Bool("m", false, "count characters")
	c := flag.Bool("c", false, "count bytes")

	flag.Parse()

	if !*l && !*w && !*m && !*c {
		*l = true
		*w = true
		*m = true
		*c = true
	}

	filename := os.Args[len(os.Args)-1]

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if *l {
		fmt.Println(getLines(data))
	}
	if *c {
		fmt.Println(len(data))
	}
	if *w {
		fmt.Println(getWords(data))
	}
	if *m {
		fmt.Println(utf8.RuneCount(data))
	}
}
