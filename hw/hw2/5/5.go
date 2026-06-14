package main

import "fmt"

func addFive(a *int) {
	*a += 5
}

func main() {
	a := [5]int{1, 4, 5, 1, 6}
	
	for i := 0; i < len(a); i ++ {
		if a[i] % 2 == 0 {
			addFive(&a[i])
		}
	}

	fmt.Println(a)
}
