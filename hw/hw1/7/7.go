package main

import "fmt"

func main() {
	var a, b int
	var operation string
	fmt.Scan(&a, &b, &operation)

	switch operation {
	case "+":
		fmt.Printf("a + b = %d\n", a+b)
	case "-":
		fmt.Printf("a - b = %d\n", a-b)
	case "*":
		fmt.Printf("a * b = %d\n", a*b)
	case "/":
		if b != 0 {
			fmt.Printf("a / b = %d\n", a/b)
		} else {
			fmt.Println("Ошибка: деление на ноль")
		}
	default:
		fmt.Println("Ошибка: неверный операнд")
	}
}
