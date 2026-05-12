package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)
	c := []rune(s)

	for i := range c {
		if !unicode.IsDigit(c[i]) {
			fmt.Println("Не число")
			return
		}
	}

	fmt.Println("Число")
}