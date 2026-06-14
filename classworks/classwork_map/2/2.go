package main

import "fmt"

func main() {
	var prod map[string]int = map[string]int{
		"banana": 10,
		"milk": 20,
		"apple": 10,
	}
	
	if v, ok := prod["apple"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("нет яблока")
	}
}