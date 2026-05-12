package main

import (
	"fmt"
	"time"
)

func main() {
	birthDay := time.Date(2001, 8, 25, 0, 0, 0, 0, time.UTC)
	diff := time.Since(birthDay)

	fmt.Println(diff)
}
