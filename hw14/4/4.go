package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	
	var str string
	fmt.Scan(&str)
	txt, _ := os.ReadFile("hw14/4/text.txt")
	words := strings.Fields(string(txt))
	for _, v := range words {
		if v == str {
			fmt.Println("Найдено")
			return
		}
	}
	fmt.Println("Не найдено")
}
