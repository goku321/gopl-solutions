// Write a program wordfreq to report the frequency of each word in an input text
// file. Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into
// words instead of lines.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func wordFreq(fileName string) (map[string]int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	fileReader := bufio.NewScanner(f)
	fileReader.Split(bufio.ScanWords)
	for fileReader.Scan() {
		fmt.Println(fileReader.Text())
	}
	return nil, nil
}

func main() {
	wordFreq("words.txt")
}
