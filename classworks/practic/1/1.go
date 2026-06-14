package main

import "fmt" 

func main() {
	var n, sum int
	fmt.Scan(&n)

	for i := 2; i <= n; i += 2 {
		sum += i
	}
	fmt.Println(sum)
}