package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	txt, _ := os.ReadFile("hw14/2/file.txt")
	str := strings.Fields(string(txt))

	fmt.Println(len(str), str)
}
