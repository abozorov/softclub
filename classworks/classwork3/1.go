package main

import (
	"fmt"
	"unicode"
)

func isCoolPassword(s_pass string) bool {
	pass := []rune(s_pass)
	if len(pass) < 5 {
		return false
	}

	for i := 0; i < len(pass); i++ {
		if !unicode.Is(unicode.Latin, pass[i]) && !unicode.IsDigit(pass[i]) {
			// fmt.Println("ff", string(pass[i]), !unicode.Is(unicode.Latin, pass[i]))
			return false
		}
	}

	return true
}

func main() {
	var password string
	fmt.Scan(&password)
	
	if isCoolPassword(password) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}
