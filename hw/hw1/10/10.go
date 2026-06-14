package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 int
	fmt.Scan(&x1, &y1, &x2, &y2)

	if math.Abs(float64(x1) - float64(x2)) == 2 && math.Abs(float64(y1) - float64(y2)) == 1 ||
	   math.Abs(float64(y1) - float64(y2)) == 2 && math.Abs(float64(x1) - float64(x2)) == 1 {
		fmt.Println("Да")
	} else {
		fmt.Println("Нет")
	}
}
