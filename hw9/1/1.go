package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string

	fmt.Scan(&s)
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	if n < 0 {
		fmt.Println("Отрицательное число")
		return
	}
	fmt.Println(n * n)
}