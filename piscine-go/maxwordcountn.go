package main

import (
	"fmt"
	"wordcount"
)

func main() {
	text := "the quick brown fox jumps over the lazy dog"
	wordCounts := wordcount.MaxWordCountN(text, 3)
	for word, count := range wordCounts {
		fmt.Printf("%s: %d\n", word, count)
	}
}
