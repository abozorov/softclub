package main

import (
	"fmt"
	"unicode/utf8"
)

func groupingOfWords(words []string) map[int][]string {
	mp := make(map[int][]string)
	for _, v := range words {
		mp[utf8.RuneCountInString(v)] = append(mp[utf8.RuneCountInString(v)], v)
	}
	
	return mp
}

func main() {
	var n int
	fmt.Scan(&n)
	words := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&words[i])
	}
	fmt.Println(groupingOfWords(words))
}
