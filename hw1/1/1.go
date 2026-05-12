package main

import "fmt"

func main() {
	var w, d, price int
	fmt.Scan(&w, &d)

	switch {
	case w <= 5:
		price = 500
	case w <= 20:
		price = 1000
	default:
		price = 2000
	}

	if d > 100 {
		price += (d - 100) * 5
	}

	fmt.Println(price)
}
