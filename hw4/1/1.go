package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println(string(s[0]), string(s[len([]rune(s)) - 1]))
}