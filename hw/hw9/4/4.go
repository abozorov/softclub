package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isVowel(r rune) bool {
	r = unicode.ToLower(r)
	return strings.ContainsRune("aeiouy", r)
}

func main() {
	var s string
	fmt.Scan(&s)
	c := []rune(s)

	for i := range c {
		if isVowel(c[i]) {
			c[i] = rune('*')
		}
	}

	fmt.Println(string(c))
}