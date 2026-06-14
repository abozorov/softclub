package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	switch {
	case 0 <= n && n <= 5:
		fmt.Println("Ночь")
	case 6 <= n && n <= 11:
		fmt.Println("Утро")
	case 12 <= n && n <= 17:
		fmt.Println("День")
	case 18 <= n && n <= 23:
		fmt.Println("Вечер")
	default:
		fmt.Println("Ошибка")
	}
}
