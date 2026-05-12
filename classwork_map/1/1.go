package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	numbers := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}
	used := make(map[int]int, len(numbers))

	for i := 0; i < n; i++ {
		used[numbers[i]]++
	}

	fmt.Println(used)
}