package main

import (
	"bufio"
	"fmt"
)

type WordCounter int

func (wc *WordCounter) Write(p []byte) (int, error) {
	fmt.Println("counting words")
	return len(p), nil
}

func main() {
	
}

