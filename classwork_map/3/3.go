package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"Ali": 20,
		"Umed": 23,
		"Vali": 19,
	}
	fmt.Println(mp)
	delete(mp, "Vali")
	fmt.Println(mp)
}