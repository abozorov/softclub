package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name, pass string
	fmt.Scan(&name, &pass)

	file, _ := os.OpenFile("hw14/5/text.txt", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)
	defer file.Close()
	wrtr := bufio.NewWriter(file)
	wrtr.WriteString(fmt.Sprintf("User registered: %s- Password: %s\n", name, pass))
	wrtr.Flush()
}
