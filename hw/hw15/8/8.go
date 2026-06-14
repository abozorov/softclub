package main

import (
	"fmt"
	"time"
)

func main() {
	var timer int
	fmt.Scan(&timer)

	fmt.Println("Таймер запущен...")
	for ;timer > 0;timer-- {
		fmt.Printf("Осталдось %d секунд\n", timer)
		time.Sleep(time.Duration(int(1e9)))
	}
	fmt.Println("Время вышло")
}
