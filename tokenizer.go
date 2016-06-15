package main

import (
	"bufio"
	"io"
	"strings"
)

func tokenizeWords(in io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(in)

	// use predefined ScanWords splitter func
	// consider something fancier here to remove punctuation
	scanner.Split(bufio.ScanWords)

	words := []string{}
	for scanner.Scan() {
		w := strings.ToLower(strings.TrimSpace(scanner.Text()))
		words = append(words, w)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}
