package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 int
	fmt.Scan(&x1, &y1, &x2, &y2)

	if x1 == x2 && y1 == y2 {
		fmt.Println("Нет (фигура на месте)")
	} else if x1 == x2 {
		fmt.Println("Да (вертикаль)")
	} else if y1 == y2 {
		fmt.Println("Да (горизонталь)")
	} else if math.Abs(float64(x1)-float64(x2)) == math.Abs(float64(y1)-float64(y2)) {
		fmt.Println("Да (диагональ)")
	} else {
		fmt.Println("Нет")
	}
}
