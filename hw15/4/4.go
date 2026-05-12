package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format(time.DateTime))

	var days int
	fmt.Scan(&days)
	now = now.AddDate(0, 0, days)
	fmt.Println(now.Format(time.DateTime))
}
