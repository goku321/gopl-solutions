package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int

func (wc *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))

	// Set the split function for scanning operation
	scanner.Split(bufio.ScanWords)

	*wc = WordCounter(0)
	for scanner.Scan() {
		*wc++
	}

	return len(p), nil
}

type LineCounter int

func (lc *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))

	count := 0
	for scanner.Scan() {
		count++
	}

	*lc = LineCounter(count)

	return len(p), nil
}

func main() {
	var wc WordCounter
	var lc LineCounter
	const inputString = `A Line was made
	using multiple words
	spanned across multiple lines
	.`
	fmt.Fprintf(&wc, inputString)
	fmt.Fprintf(&lc, inputString)

	fmt.Println(lc)
}
