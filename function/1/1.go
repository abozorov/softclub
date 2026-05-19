package main

import (
	"fmt"
	"sync"
	"time"
)

func func1() {
	fmt.Println(1, 2, 3)
}

func func2() {
	fmt.Println(4, 5, 6)
}

func func3() {
	fmt.Println(7, 8, "9++")
}

func timer(wg *sync.WaitGroup, t int) {
	defer wg.Done()
	fmt.Println("timer start")
	for i := 0; i < t; i++ {
		time.Sleep(time.Second)
		fmt.Println("second", i+1)
	}
}

func goFunc(i int, wg *sync.WaitGroup) {
	defer wg.Done()
}

func main() {
	go func1()
	time.Sleep(100 * time.Millisecond)
	go func2()
	time.Sleep(100 * time.Millisecond)
	go func3()
	time.Sleep(100 * time.Millisecond)

	wg := sync.WaitGroup{}
	wg.Add(1)
	timer(&wg, 3)
	wg.Wait()

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Goroutine number:", i)
		}()
	}
	wg.Wait()
}
