package main

import "fmt"

func setZero(a *int) {
	*a = 0
}

func main() {
	var a int
	fmt.Scan(&a)
	setZero(&a)
	fmt.Println(a)
}
