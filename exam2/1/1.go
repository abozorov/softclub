package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func normalizeName(name string) string {
	name = strings.TrimSpace(name)
	rName := []rune(name)
	if len(rName) == 0 {
		return "Unknown"
	}
	rName[0] = unicode.ToUpper(rName[0])
	return string(rName)
}

func main() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Println(normalizeName(str))
}