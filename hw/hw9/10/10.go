package main

import (
	"fmt"
	"unicode"
)

func myAtoi(s string) (int, error) {
	r := []rune(s)
	if len(r) == 0 || (len(r) == 1 && ('0' > r[0] || r[0] > '9')) {
		return -1, fmt.Errorf("Строка %s не число", s)
	}
	n, p, raz := 0, 0, 1
	if r[0] == rune('-') {
		raz = -1
		p++
	}

	for i := p; i < len(r); i++ {
		if !unicode.IsDigit(r[i]) {
			return -1, fmt.Errorf("Строка %s не число", s)
		}
		n = n * 10 + int(r[i]) - 48
	}

	return n*raz, nil
}

func main() {
	var s string
	fmt.Scan(&s)

	n, err := myAtoi(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
}
