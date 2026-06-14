package main

import (
	"fmt"
)

func Map[T any, R any](items []T, fn func(T) R) []R {
	a := make([]R, 0)
	for _, v := range items {
		a = append(a, fn(v))
	}
	return a
}

func main() {
	fmt.Println(Map([]int{1, 2, 3, 4, 1, 5, 6, 2, 7}, func(i int) int {
		return i * i
	}))
	fmt.Println(Map([]string{"1", "3", "4", "1", "2", "2"}, func(s string) string {
		return s + "Hello" 
	}))
}
