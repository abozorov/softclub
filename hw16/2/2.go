package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Reader interface {
	Read(path string) ([]string, error)
}

type TxtReader struct{}

func (t *TxtReader) Read(path string) ([]string, error) {
	file, err := os.OpenFile(path+".txt", os.O_RDONLY, 0644)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	text := make([]string, 0)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	return text, nil
}

type CSVReader struct{}

func (c *CSVReader) Read(path string) ([]string, error) {
	file, err := os.OpenFile(path+".csv", os.O_RDONLY, 0644)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	text := make([]string, 0)
	for scanner.Scan() {
		text = append(text, strings.Split(scanner.Text(), ",")...)
	}
	return text, nil
}

func ReadAll(readers []Reader, path string) (allText [][]string) {
	for _, v := range readers {

		text, err := v.Read(path)
		if err != nil {
			fmt.Printf("Reader: %v\nError: %s\n", v, err)
			continue
		}
		allText = append(allText, text)
	}
	return
}

func main() {
	readers := []Reader{
		&TxtReader{},
		&CSVReader{},
	}
	fmt.Println(ReadAll(readers, "hw16/2/file"))
}
