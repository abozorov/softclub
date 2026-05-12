package main

import (
	"bytes"
	"fmt"
)

func main() {
	var a, rs string
	fmt.Scan(&a)
	s := []byte(a)

	for _, v := range s {
		rs = string(v) + rs
		fmt.Println(string(v))
		// fmt.Printf("%x \n", v) 

	}

	fmt.Println(s, rs)

	if bytes.Equal(s, []byte(rs)) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}