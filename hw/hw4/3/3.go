package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Println(strings.Join(strings.Split(s, " "), ""))
}