package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

func main() {
	var in io.Reader

	if len(os.Args) > 1 {
		fname := os.Args[1]
		fmt.Println("opening file", fname)
		f, err := os.Open(fname)
		if err != nil {
			log.Fatalf("Error opening file %q - %s\n", fname, err)
		}
		defer f.Close()

		in = bufio.NewReader(f)
	} else {
		in = os.Stdin
	}

	words, err := scanForWords(in)
	if err != nil {
		log.Fatalf("Error scanning words %s\n", err)
	}
	// fmt.Printf("words: %q\n", words)

	wordStats := values(uniqueWords(words))
	sort.Sort(sort.Reverse(ByCount(wordStats)))

	for _, stats := range wordStats {
		fmt.Printf("%d\t%s\n", stats.count, stats.word)
	}
}
