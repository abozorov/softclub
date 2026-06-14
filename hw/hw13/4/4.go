package main

import "fmt"

func CountValues[T comparable](items []T) map[T]int {
	us := make(map[T]int, 0)
	for _, v := range items {
		us[v]++
	}
	return us
}

func main() {
	fmt.Println(CountValues([]int{1, 2, 3, 4, 1, 5, 6, 2, 7}))
	fmt.Println(CountValues([]string{"1", "2", "2", "1", "3", "4"}))
}
