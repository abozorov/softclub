package main

import "fmt"

func Min[T int | float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Min(3, 5))
	fmt.Println(Min(2.5, 1.1))
}
