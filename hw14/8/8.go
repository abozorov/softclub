package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum := 0
	mx := -10000000000
	mn := 10000000000
	file, err := os.OpenFile("hw14/8/numbers.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	rdr := bufio.NewScanner(file)

	for rdr.Scan() {
		s := rdr.Text()
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
			return
		}
		sum = sum + num
		if mx < num {
			mx = num
		}
		if mn > num {
			mn = num
		}
	}
	fmt.Printf("Сумма: %d. Максимум: %d. Минимум: %d.\n", sum, mx, mn)
}
