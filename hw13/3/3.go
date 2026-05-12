package main

import "fmt"

func Unique[T comparable](items []T) []T {
	a := make([]T, 0, len(items))
	us := make(map[T]struct{}, 0)
	for _, v := range items {
		if _, ok := us[v]; !ok {
			us[v] = struct{}{}
			a = append(a, v)
		}
	}
	return a
}

func main() {
	fmt.Println(Unique([]int{1, 2, 3, 4, 1, 5, 6, 2, 7}))
	fmt.Println(Unique([]string{"1", "2", "2", "1",  "3", "4"}))
}
