package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a+b <= c || a+c <= b || b+c <= a {
		fmt.Println("Треугольник не существует")
	} else if a == b && b == c {
		fmt.Println("Равносторонний")
	} else if a == b || a == c || b == c {
		fmt.Println("Равнобедренный")
	} else {
		fmt.Println("Разносторонний")
	}
}
