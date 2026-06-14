package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	fl = "hw14/7/data.txt"
)

func createToDo(file *os.File) {
	fmt.Println("ToDo text:")
	rdr := bufio.NewScanner(os.Stdin)

	if rdr.Scan() {
		todo := rdr.Text()
		wrt := bufio.NewWriter(file)
		wrt.WriteString(fmt.Sprintf("ToDO: %s\n", todo))
		wrt.Flush()
	}
}

func loadAllToDo() {
	todos, err := os.ReadFile(fl)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(todos))
}

func main() {
	commands := `
Выерите команду:
	0. Выход
	1. Добавить задачу
	2. Вывести все задачи`

	file, err := os.OpenFile(fl, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for {
		fmt.Println(commands)
		var c string
		fmt.Scanln(&c)
		switch c {
		case "0":
			fmt.Println("До скорых встреч :D")
			return
		case "1":
			createToDo(file)
		case "2":
			loadAllToDo()
		default:
			fmt.Println("Такой команды нет")
		}
	}
}
