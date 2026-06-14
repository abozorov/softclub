package main

import (
	"fmt"
)

func unicWordCount(words []string) int {
	used := make(map[string]struct{})
	for _, v := range words {
		used[v] = struct{}{}
	}
	unicWordCount := len(used)
	used = nil
	return unicWordCount
}

func main() {
	var n int
	fmt.Scan(&n)
	words := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&words[i])
	}
	fmt.Println(unicWordCount(words))
}
