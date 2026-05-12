package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	logs     = make(map[string]int)
	fileName = "exam5/2/log.log"
)

func ReadFile() error {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := bufio.NewScanner(file)

	for reader.Scan() {
		line := strings.Fields(reader.Text())
		// fmt.Println(line)
		logs[line[0]]++
	}
	return nil
}

func main() {
	err := ReadFile()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("INFO: %d\nERROR: %d\nWARNING: %d\n", logs["INFO"], logs["ERROR"], logs["WARNING"])

	if logs["INFO"] >= logs["ERROR"] && logs["INFO"] >= logs["WARNING"] {
		fmt.Println("more logs: INFO")
	} else if logs["ERROR"] >= logs["INFO"] && logs["ERROR"] >= logs["WARNING"] {
		fmt.Println("more logs: ERROR")
	} else {
		fmt.Println("more logs: WARNING")

	}
}
