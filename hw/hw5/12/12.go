package main

import (
	"fmt"
	"strconv"
)

func main() {

	num, err := strconv.Atoi("sd")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)
}
