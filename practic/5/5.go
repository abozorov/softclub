package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]float64, n)
	sum := 0.0

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		sum += a[i]
	}
	sum /= float64(n)
	k := 0

	for i := 0; i < n; i++ {
		if a[i] > sum {
			k ++
		}
	}

	fmt.Println(k)
}
