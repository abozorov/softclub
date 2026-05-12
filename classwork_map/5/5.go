package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	mp := make(map[string]int)
	words := []string{"go", "golang", "code"}

	for i := range words {
		mp[words[i]] = utf8.RuneCountInString(words[i])
	}
	fmt.Println(mp)
}
