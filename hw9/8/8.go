package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)
	c:= []rune(s)
	l, d := 0, 0

	for _, v:= range c {
		if unicode.IsDigit(v) {
			d++
		} else if unicode.IsLetter(v) {
			l++
		}
	}

	fmt.Printf("Числа %d\nбуквы %d\n", d, l)
}
