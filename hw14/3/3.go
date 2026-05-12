package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	name := ""
	file, _ := os.OpenFile("hw14/3/users.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	wrtr := bufio.NewWriter(file)

	for fmt.Scan(&name); name != "exit"; {
		wrtr.WriteString(name + ", ")
		fmt.Scan(&name)
		wrtr.Flush()
	}
	wrtr.WriteString(".")
	wrtr.Flush()
}
