package main

import (
	"flag"
	"fmt"
	"io"
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

func getChars(data []byte) int {
	return utf8.RuneCount(data)
}

func getBytes(data []byte) int {
	return len(data)
}

func main() {

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

	info, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var data []byte

	if (info.Mode() & os.ModeCharDevice) == 0 {
		data, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Error reading stdin:", err)
			return
		}
	} else if len(flag.Args()) > 0 {
		filename := flag.Args()[0]
		data, err = os.ReadFile(filename)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
	} else {
		fmt.Println("Usage: <command> [options] <filename>")
		return
	}

	if *l {
		fmt.Println(getLines(data))
	}
	if *c {
		fmt.Println(getBytes(data))
	}
	if *w {
		fmt.Println(getWords(data))
	}
	if *m {
		fmt.Println(getChars(data))
	}
}
