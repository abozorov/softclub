package main

import (
	"fmt"
)

func cymbolCount(word string) map[string]int {
	mp := make(map[string]int)
	rWord := []rune(word)
	for _, v := range rWord {
		mp[string(v)]++
	}
	
	return mp
}

func main() {
	var str string
	fmt.Scan(&str)
	
	fmt.Println(cymbolCount(str))
}
