package main

import (
	"fmt"
	"time"
)

func main() {
	date1 := time.Now()
	date2 := date1.AddDate(1, 0, 0)

	fmt.Println("After? ", date1.After(date2))
	fmt.Println("Before? ", date1.Before(date2))
	fmt.Println("Equal? ", date1.Equal(date2))
}
