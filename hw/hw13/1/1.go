package main

import (
	"fmt"
)

func Last[T any](items []T) (T, bool) {
	if len(items) > 0 {
		return items[len(items) - 1], true
	}
	return *new(T), false
}

func main() {
	fmt.Println(Last([]int{1, 2, 3, 4}))
	fmt.Println(Last([]rune{}))
}