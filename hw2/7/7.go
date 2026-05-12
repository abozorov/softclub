package main

import "fmt"

func sum(a, b, result *int) {
	
	*result = *a + *b
}

func main() {
	var a, b, result int
	fmt.Scan(&a, &b)
	sum(&a, &b, &result)
	fmt.Println(result)
}
