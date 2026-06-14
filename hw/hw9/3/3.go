package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	c := []rune(s)

	fmt.Println(c[0])
	c[0]++
	fmt.Println(string(c[0]))

}