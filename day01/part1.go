package main

import (
	"unicode"
)

func findFirstDigit(text string, out chan uint8) {
	for _, c := range text {
		if unicode.IsDigit(rune(c)) {
			out <- uint8(c - '0')
			return
		}
	}
	out <- 0 // no digit found
}

func findLastDigit(text string, out chan uint8) {
	for i := len(text) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(text[i])) {
			out <- uint8(text[i] - '0')
			return
		}
	}
	out <- 0 // no digit found
}
