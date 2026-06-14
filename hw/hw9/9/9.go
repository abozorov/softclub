package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)
	c:= []rune(s)
	a, d, oper := [2]int{0, 0}, 0, rune(0)
	for _, v := range c {
		if unicode.IsDigit(v) {
			a[d] = a[d] * 10 + int(v) - 48
		} else {
			oper = v
			d++
		}
	}

	if oper == rune('-') {
		fmt.Println(a[0] - a[1])
	} else if oper == rune('+') {
		fmt.Println(a[0] + a[1])
	} else if oper == rune('*') {
		fmt.Println(a[0] * a[1])
	} else {
		fmt.Println(a[0] / a[1])

	}
}
