package main

import (
	"fmt"
)

func DifferenceUnique(a, b []int) []int {
	usB := make(map[int]struct{}, len(a)+len(b))

	for _, v := range b {
		usB[v] = struct{}{}
	}
	usA := make(map[int]struct{}, len(a)+len(b))
	ans := make([]int, 0, len(a))

	for _, v := range a {
		if _, ok := usB[v]; !ok {
			if _, ok := usA[v]; !ok {
				usA[v] = struct{}{}
				ans = append(ans, v)
			}
		}
	}
	usA = nil
	usB = nil
	return ans
}

func main() {
	a := []int{0, 1, 2, 3, 2, 3, 4, 5, 5, 6}
	b := []int{2, 4, 7, 3}
	fmt.Println(DifferenceUnique(a, b))
}
