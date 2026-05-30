package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 1. Очередь заказов
func getOrders(ch chan<- int) {
	defer close(ch)

	for i := 0; i < 5; i++ {
		ch <- rand.Int() % 20
	}
}

func problem1() {
	ch := make(chan int, 5)

	go getOrders(ch)

	for val := range ch {
		fmt.Printf("Обрабатывается: Заказ #%d\n", val)
	}
}

// 2. Фильтрация данных
func moreThan10(numbers []int, ch chan<- int) {
	defer close(ch)
	for _, v := range numbers {
		if v > 10 {
			ch <- v

		}
	}
}

func problem2() {
	numbers := []int{3, 8, 15, 20, 27, 40}
	ch := make(chan int, len(numbers))

	go moreThan10(numbers, ch)

	for val := range ch {
		fmt.Printf("#%d\n", val)
	}
}

// 3. Онлайн-магазин
func problem3() {
	prod := []string{
		"Phone",
		"Laptop",
		"Tablet",
	}
	prodCh := make(chan string)
	go func() {
		defer close(prodCh)
		for _, v := range prod {
			prodCh <- v
		}
	}()

	prc := []int{
		1000,
		2500,
		800,
	}
	prcCh := make(chan int)
	go func() {
		defer close(prcCh)
		for _, v := range prc {
			prcCh <- v
		}
	}()

	for i := 0; i < 3; i++ {
		product := <-prodCh
		price := <-prcCh

		fmt.Printf("product %s, price %d\n", product, price)
	}
}

// 4. Кто закончит раньше (select)
func problem4() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		defer close(ch1)
		time.Sleep(time.Second * 2)
		ch1 <- fmt.Sprintln("First goroutine complete downloading file")
	}()

	go func() {
		defer close(ch2)
		time.Sleep(time.Second * 1)
		ch2 <- fmt.Sprintln("Second goroutine complete downloading file")
	}()

	select {
	case msg := <-ch1:
		fmt.Print(msg)
	case msg := <-ch2:
		fmt.Print(msg)
	}
}

// 5. Мониторинг сервера (select)
func problem5() {
	errors := make(chan string)
	success := make(chan string)

	go func() {
		defer close(errors)
		time.Sleep(time.Microsecond * time.Duration(rand.Int()%1000))
		errors <- fmt.Sprintln("Server Error")
	}()

	go func() {
		defer close(success)
		time.Sleep(time.Microsecond * time.Duration(rand.Int()%1000))
		success <- fmt.Sprintln("Request Success")
	}()

	for {
		ok := false
		select {
		case msg := <-errors:
			fmt.Print(msg)
			ok = true
		case msg := <-success:
			fmt.Print(msg)
			ok = true
		}

		if ok {
			break
		}
	}
}
func main() {
	// problem1()
	// problem2()
	// problem3()
	// problem4()
	problem5()
}
