package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("hw14/1/file.txt")
	defer file.Close()
	scn := bufio.NewScanner(file)

	n := 0
	for ;scn.Scan();n++ {
		scn.Text()
	} 
	fmt.Println(n)
}
