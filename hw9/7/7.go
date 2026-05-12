package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	c, sum := []rune(s), 0

	for _, v:= range c {
		n, err := strconv.Atoi(string(v))
		if err != nil {
			fmt.Println(err)
			return
		}
		sum += n
	}

	fmt.Println(sum)
}
