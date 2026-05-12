package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	s = strings.ToLower(s)

	fmt.Println(strings.Count(s, "go"))
}