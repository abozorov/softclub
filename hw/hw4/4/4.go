package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str, news string
	fmt.Scan(&str)
	s := []rune(str)

	for i := 0; i < len(s); {
		c, k := s[i], 0

		for ;i < len(s) && c == s[i]; i++ {
			k++
		}
		news = news + string(c) + strconv.Itoa(k)
	}
	fmt.Println(news)
}