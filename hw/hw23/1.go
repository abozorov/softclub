package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 1. Генератор чисел с остановкой
func generator(ctx context.Context, numbers chan<- int) {
	defer close(numbers)
	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Генератор остановлен")
			return
		case <-time.After(time.Second):
			numbers <- i
		}
	}
}

func problem1() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := make(chan int, 10)
	go generator(ctx, ch)

	i := 0
	for v := range ch {
		fmt.Println(v)
		i++
		if i == 3 {
			cancel()
		}
	}
}

// 2. Ограничение времени обработки заказа
func processOrder(ctx context.Context) {
	orderTime := time.Second * time.Duration(rand.Int()%6)
	fmt.Println(orderTime)
	avb := time.Duration(0)
	for {
		select {
		case <-ctx.Done():
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				fmt.Println("Заказ отменён!")
			}
			return
		case <-time.After(time.Second):
			avb += time.Second
			if avb >= orderTime {
				fmt.Println("Заказ выполнен!")
				return
			}

			fmt.Println("Обрабатывается заказ...")
		}
	}
}

func problem2() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		processOrder(ctx)
	}()
	wg.Wait()
}

// 3. Передача значения через context
type User struct {
	Name string
}

func printUser(ctx context.Context) {
	id := rand.Int() % 10
	fmt.Println(id)
	if val := ctx.Value(id); val != nil {
		fmt.Println("Пользователь:", val.(User).Name)
	} else {
		fmt.Println("Пользователь не найден")
	}
}

func problem3() {
	ctx := context.WithValue(context.Background(), 1, User{Name: "Bozorov A."})
	ctx = context.WithValue(ctx, 2, User{Name: "Bob"})
	ctx = context.WithValue(ctx, 3, User{Name: "Alice"})
	ctx = context.WithValue(ctx, 8, User{Name: "Jake"})
	ctx = context.WithValue(ctx, 0, User{Name: "Grace"})

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		printUser(ctx)
	}()
	wg.Wait()
}

// 4. Конкурирующие задачи (кто успеет первым)
func problem4() {
	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()
	wg := new(sync.WaitGroup)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ok := false
			workTime := time.Millisecond * time.Duration(rand.Int()%2000)
			for {
				select {
				case <-ctx.Done():
					if !ok {
						fmt.Printf("Worker %d отменён\n", i)
					}
					return
				case <-time.After(workTime):
					ok = true
					fmt.Printf("Worker %d завершился\n", i)
					cancle()
				}
			}
		}()
	}
	wg.Wait()
}

// 5. Цепочка вызовов с context
func service(ctx context.Context) {
	repository(ctx)
	err := ctx.Err()
	if errors.Is(err, context.DeadlineExceeded) {
		fmt.Println("Операция не завершена:", err)
	}
}

func repository(ctx context.Context) {
	workTime := time.Second * time.Duration(rand.Int()%6)
	fmt.Println("worktime", workTime)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Репозиторий остановлен")
			return
		case <-time.After(workTime):
			fmt.Println("Successfuly!")
			return
		}
	}
}

func problem5() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		service(ctx)
	}()
	wg.Wait()
}

func main() {
	// problem1()
	// problem2()
	// problem3()
	// problem4()
	problem5()
}
