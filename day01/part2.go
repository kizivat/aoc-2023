package main

import (
	"unicode"
)

var digitStrings = [...]string{
	"one", "two", "three", "four",
	"five", "six", "seven", "eight", "nine",
}

// trie for the digit strings
type trie struct {
	children [26]*trie
	value    uint8
}

func (t *trie) insert(s string, value uint8) {
	if len(s) == 0 {
		t.value = value
		return
	}
	c := s[0]
	if t.children[c-'a'] == nil {
		t.children[c-'a'] = &trie{}
	}
	t.children[c-'a'].insert(s[1:], value)
}

func (t *trie) find(s string) uint8 {
	if len(s) == 0 {
		return t.value
	}
	c := s[0]
	if c < 'a' || c > 'z' || t.children[c-'a'] == nil {
		return t.value
	}
	return t.children[c-'a'].find(s[1:])
}

func findDigit(trie *trie, text string, out chan uint8) {
	for i, c := range text {
		if unicode.IsDigit(rune(c)) {
			out <- uint8(c - '0')
			return
		} else if trie.find(text[i:]) > 0 {
			out <- trie.find(text[i:])
			return
		}
	}
	out <- 0 // no digit found
}

func findAnyFirstDigit(text string, out chan uint8) {
	// init the trie
	var digitTrie = &trie{}
	for i, s := range digitStrings {
		digitTrie.insert(s, uint8(i+1))
	}
	findDigit(digitTrie, text, out)
}

func findAnyLastDigit(text string, out chan uint8) {
	// reverse the text
	serchText := reverse(text)
	// init the reverse trie
	var digitTrie = &trie{}
	for i, s := range digitStrings {
		digitTrie.insert(reverse(s), uint8(i+1))
	}
	findDigit(digitTrie, serchText, out)
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
