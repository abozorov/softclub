package main

import (
	"fmt"
	"math"
)

func firstMatch(numbers []int) int {
	used := make(map[int]struct{})
	for _, v := range numbers {
		if _, ok := used[v]; ok {
			return v
		}
		used[v] = struct{}{}
	}
	used = nil
	return math.MaxInt
}

func main() {
	var n int
	fmt.Scan(&n)
	numbers := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}
	firstMatch := firstMatch(numbers)

	if firstMatch == math.MaxInt {
		fmt.Println("Нет повторов")
	} else {
		fmt.Println(firstMatch)
	}
}
