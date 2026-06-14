package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	c := []rune(s)

	for i := 0; i < len(c)/2; i++ {
		c[i], c[len(c)-i-1] = c[len(c)-i-1], c[i]
	}

	n, err := strconv.Atoi(string(c))
	if err != nil {
		fmt.Println(err)
		return

	}
	fmt.Println(n)
}
