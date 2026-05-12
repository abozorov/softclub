package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	weekDay := now.Weekday()
	fmt.Printf("Today: %v\n", weekDay)

	if weekDay > 5 {
		fmt.Println("The Weekend")
	} else {
		fmt.Println("Work Day")
	}
}
