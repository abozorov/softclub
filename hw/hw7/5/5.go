package main

import (
	"fmt"
)

func intersection(a, b []int) []int {
	used := make(map[int]struct{})
	for _, v := range a {
		used[v] = struct{}{}
	}
	inter := make([]int, 0,len(a) + len(b))
	us := make(map[int]struct{})

	for _, v := range b {
		 _, ok := used[v]
		if _, ok2 := us[v]; ok && !ok2 {
			us[v] = struct{}{}
			inter = append(inter, v)
		}
	}
	used = nil
	us = nil
	return inter
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]int, n)
	b := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&b[i])
	}
	
	fmt.Println(intersection(a, b))
}
