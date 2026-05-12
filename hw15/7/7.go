package main

import (
	"fmt"
	"time"
)

func main() {
	var str string
	fmt.Scan(&str)
	date, err := time.Parse("2006-01-02", str)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Year: %d\nMonth: %d\nDay: %d\n", date.Year(), date.Month(), date.Day())
}
