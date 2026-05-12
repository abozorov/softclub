package main

import "fmt"

func Merge[T any](a, b []T) []T {
	return append(a, b...)
}

func main() {
	fmt.Println(Merge([]int{6, 2, 7}, []int{1, 2, 3, 4, 1, 5}))
	fmt.Println(Merge([]string{"1", "3", "4"}, []string{"1", "2", "2"}))
}
