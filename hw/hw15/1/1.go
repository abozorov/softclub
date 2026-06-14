package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Printf("Year: %d\nMonth: %d\nDay: %d\nHour: %d\nMinut: %d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute())
}