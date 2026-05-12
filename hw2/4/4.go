package main

import "fmt"

func replaceIfGreater(a *int, b *int) {
	
	if *a < *b {
		*a, *b = *b, *a
	}
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	replaceIfGreater(&a, &b)
	fmt.Println(a, b)
}
