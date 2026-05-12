package main

import (
	"fmt"
)

func increment(a *int) {
	*a++
}

func main() {
	var a int
	fmt.Scan(&a)
	increment(&a)
	fmt.Println(a)
}
