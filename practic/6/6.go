package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := n - 1; i >= 0; i-- {
		fmt.Print(a[i], " ")
	}
	
}
