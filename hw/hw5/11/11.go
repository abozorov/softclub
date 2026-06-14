package main

import (
	"fmt"
)

func someFunc(n int) {
	if n == 0 {
		panic("Поймана ошибка")
	}
}

func recoverPanic() {
	if pnc := recover(); pnc != nil {
		fmt.Printf("Recovered: %v\n", pnc)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	defer recoverPanic()
	someFunc(n)
}
