package main

import (
	"fmt"
)

func main() {

	// Problem 1
	fmt.Println("Problem 1")
	helloGo := func() {
		fmt.Println("Hello Go")
	}
	helloGo()
	
	// Problem 2
	fmt.Println("Problem 2")
	sum := func(a, b int) int {
		return a+b
	}
	fmt.Println(sum(1, 2))

	// Problem 3
	fmt.Println("Problem 3")
	isOdd := func(x int) bool {
		return x % 2 == 0
	}
	fmt.Println(isOdd(4))

	// Problem 4
	fmt.Println("Problem 4")
	nums := []int{1,2,3,4,5}
	for _, v := range nums {
		fmt.Print(func(n int) int {
			return n*2
		}(v), " ")
	}
	fmt.Println()

	// Problem 5
	fmt.Println("Problem 5")
	nums = []int{4, 9, 2, 7, 1}
	fmt.Println(func(n []int) int {
		if len(n) == 0 {
			return -1
		}
		mx := n[0]
		for _, v := range n {
			if mx < v {
				mx = v
			}
		}
		return mx
	}(nums))

	// Problem 6
	fmt.Println("Problem 6")
	countF := func() func()  {
		count := 0
		return func()  {
			count++
			fmt.Println(count)
		}
	}()
	countF()
	countF()
	countF()

	// Problem 7
	fmt.Println("Problem 7")
	nums = []int{1,2,3,4,5,6}
	fmt.Println(func(n []int) []int {
		odds := make([]int, 0, len(n))
		for _, v := range n {
			if v % 2 == 0 {
				odds = append(odds, v)
			}
		}
		return odds
	}(nums))
}
