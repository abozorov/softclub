package main

import (
	"fmt"
)

func main() {
	var f float64
	fmt.Scan(&f)
	i := int(f)


	fmt.Printf("full %d\nRaznica %.1f\n",i,  f - float64(i))
}