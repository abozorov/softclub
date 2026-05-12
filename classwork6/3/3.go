package main

import "fmt"

func SumSlice[T int | float64](nums []T) (sum T) {
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(SumSlice([]int{1, 2, 3}))
	fmt.Println(SumSlice([]float64{1.5, 2.5, 3.3}))
}