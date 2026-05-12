package main

import (
	"errors"
	"fmt"
	"math"
)

func safeSqrt(n int) (float64, error) {
	if n < 0 {
		return 0, errors.New("Ошибка: нельзя извлечь корень из отрицательного числа")
	}
	return math.Sqrt(float64(n)), nil
}

func main() {
	var n int
	fmt.Scan(&n)
	sqrtN, err := safeSqrt(n)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sqrtN)
}
