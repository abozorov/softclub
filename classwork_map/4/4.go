package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"Ali": 80,
		"Umed": 95,
		"Vali": 78,
	}
	
	mx, name := 0, ""
	for k, v := range mp {
		if v > mx {
			name = k
			mx = v
		}
	}
	fmt.Println(name)
}