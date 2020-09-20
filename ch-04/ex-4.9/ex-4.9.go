// Write a program wordfreq to report the frequency of each word in an input text
// file. Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into
// words instead of lines.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func wordFreq(fileName string) (map[string]int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	fileReader := bufio.NewScanner(f)
	fileReader.Split(bufio.ScanWords)
	freq := make(map[string]int, 0)
	for fileReader.Scan() {
		if fileReader.Err() != nil {
			return freq, fileReader.Err()
		}
		freq[strings.ToLower(fileReader.Text())]++
	}
	return freq, nil
}

func main() {
	freq, err := wordFreq("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range freq {
		fmt.Printf("%s occurs %d times\n", k, v)
	}
}
