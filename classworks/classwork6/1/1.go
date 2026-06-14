package main

import "fmt"

func Print[T any](value T) {
	fmt.Printf("%v\n", value)
}

func main() {
	Print(10)
	Print("hello")
	Print(true)
	Print([]int{1, 2, 3, 4, 5, 6})

}
