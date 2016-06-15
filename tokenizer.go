package main

import (
	"bufio"
	"io"
	"strings"
)

func scanForWords(in io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(in)

	// use predefined ScanWords splitter func
	// consider something fancier here to remove punctuation
	scanner.Split(bufio.ScanWords)

	words := []string{}
	for scanner.Scan() {
		word := strings.ToLower(strings.TrimSpace(scanner.Text()))
		words = append(words, word)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}
