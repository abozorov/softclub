package main

import "fmt"

func Reverse[T any](items []T) []T {
	a := make([]T, len(items))
	len := len(items) - 1

	for i := 0; i <= len; i++ {
		a[i] = items[len-i]
	}
	return a
}

func main() {
	fmt.Println(Reverse([]int{1, 2, 3, 4, 5, 6, 7}))
	fmt.Println(Reverse([]string{"1", "2", "3", "4"}))
}
