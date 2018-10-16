package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// LoadDict takes a txt file path for the dictionary and returns the file
func LoadDict(p string) *os.File {
	f, err := os.Open(p)
	if err != nil {
		log.Fatal(fmt.Errorf("cannot read file, %v", err))
	}
	return f
}

// Find returns true if the prefix exactly matches the dict and all words in the dictionary which start with the given prefix
func Find(prefix string, dict *os.File) (bool, []string) {
	scanner := bufio.NewScanner(dict)
	words := []string{}
	exactMatch := false
	for scanner.Scan() {
		t := scanner.Text()
		if t == prefix {
			exactMatch = true
		} else if strings.HasPrefix(t, prefix) {
			words = append(words, t)
		}
	}
	return exactMatch, words
}
