package main

import "fmt"

func fill(a **int) {
	b := new(int)
	fmt.Println(*b, b)
	*b = 100
	a = &b

	fmt.Println(**a, *a, a, *b, b)
}

func main() {
	var a **int
	fmt.Println(a)

	fill(a)

}
