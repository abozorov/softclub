package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	switch n {
	case 1:
		fmt.Printf("%d дней\n", 31)
	case 2:
		fmt.Printf("%d дней\n", 28)
	case 3:
		fmt.Printf("%d дней\n", 31)
	case 4:
		fmt.Printf("%d дней\n", 30)
	case 5:
		fmt.Printf("%d дней\n", 31)
	case 6:
		fmt.Printf("%d дней\n", 30)
	case 7:
		fmt.Printf("%d дней\n", 31)
	case 8:
		fmt.Printf("%d дней\n", 31)
	case 9:
		fmt.Printf("%d дней\n", 30)
	case 10:
		fmt.Printf("%d дней\n", 31)
	case 11:
		fmt.Printf("%d дней\n", 30)
	case 12:
		fmt.Printf("%d дней\n", 31)
	default:
		fmt.Println("Ошибка")
	}
}
