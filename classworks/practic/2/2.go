package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n <= 1 {
		fmt.Println("not prime")
		return
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			fmt.Println("not prime")
			return
		}
	}
	fmt.Println("prime")
}
