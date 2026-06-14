package main

import "fmt"

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	if n * m >= k {
		fmt.Println("YES");
	} else {
		fmt.Println("NO");
	}
}
