package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func stripPunctuation(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) {
			return -1
		}
		return r
	}, s)
}

func countWords(s string) map[string]int {
	words := strings.Fields(stripPunctuation(s))
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}
	return wordCount
}

var (
	word_count map[string]int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a sentence: ")
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	line = strings.TrimSpace(line)

	word_count = countWords(stripPunctuation(line))
	fmt.Println("Word count:", word_count)
}
