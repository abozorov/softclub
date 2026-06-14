package main

import "fmt"

func swapFirstLast(nums []int) {
	if len(nums) == 0 {
		return
	}

	nums[0], nums[len(nums)-1] = nums[len(nums)-1], nums[0]
}

func main() {
	a := []int{1, 4, 5, 1, 6}
	
	swapFirstLast(a)
	fmt.Println(a)
}
