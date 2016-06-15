package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

var verbose bool

func init() {
	flag.BoolVar(&verbose, "v", false, "verbose/debug output")
	flag.Parse()
}

func main() {
	var in io.Reader

	if len(flag.Args()) > 0 {
		fname := flag.Args()[0]
		debug("opening file %q\n", fname)
		f, err := os.Open(fname)
		if err != nil {
			log.Fatalf("Error opening file %q - %s\n", fname, err)
		}
		defer f.Close()

		in = bufio.NewReader(f)
	} else {
		in = os.Stdin
	}

	words, err := tokenizeWords(in)
	if err != nil {
		log.Fatalf("Error tokenizing words %s\n", err)
	}
	debug("words: %s\n", words)

	wordStats := findUnique(words).statValues()
	sort.Sort(sort.Reverse(ByCount(wordStats)))

	for _, stats := range wordStats {
		fmt.Printf("%d\t%s\n", stats.count, stats.word)
	}
}

func debug(format string, args ...interface{}) {
	if !verbose {
		return
	}

	fmt.Fprintf(os.Stderr, format, args)
}
