package main

import (
	"io"
	"os"
)

func main() {
	source, _ := os.Open("hw14/6/source.txt")
	defer source.Close()
	copy, _ := os.OpenFile("hw14/6/copy.txt", os.O_WRONLY | os.O_CREATE, 0644)
	defer copy.Close()
	io.Copy(copy, source)
}
