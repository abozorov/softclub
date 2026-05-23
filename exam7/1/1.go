package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func problem1() {
	ch := make(chan string)
	go func() {
		ch <- "hello from goroutine"
	}()
	fmt.Println(<-ch)
	close(ch)
}

func problem2() {
	nums := []int{1, 2, 3, 4, 5}
	ch := make(chan int)
	go func() {
		sum := 0
		for _, v := range nums {
			sum += v
		}
		ch <- sum
	}()

	fmt.Println(<-ch)
}

func goroutine1(w []string, ch chan<- string) {
	defer close(ch)
	for _, v := range w {
		ch <- v
	}
}

func goroutine2(ch <-chan string, chOut chan<- string) {
	defer close(chOut)
	for v := range ch {
		chOut <- strings.ToUpper(v)
	}
}

func problem3() {
	words := []string{"go", "lang", "channel", "gogogogo", "help"}
	ch := make(chan string, 2)
	go goroutine1(words, ch)
	chOut := make(chan string, 2)
	go goroutine2(ch, chOut)

	for v := range chOut {
		fmt.Println(v)
	}
}

func problem4() {
	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int, 3)

	for _, v := range tasks {
		ch <- v
		go func() {
			fmt.Println("processing", v)
			time.Sleep(time.Second)
			<-ch
		}()
	}
}

type SafeMap struct {
	data map[string]int
	mu   *sync.Mutex
}

func (sm *SafeMap) Set(key string, value int) {
	for {
		if sm.mu.TryLock() {
			sm.data[key] = value
			sm.mu.Unlock()
			fmt.Printf("Set key %s, val %d\n", key, value)
			return
		}
	}
}

func (sm *SafeMap) Get(key string) (int, bool) {
	for {
		if sm.mu.TryLock() {
			defer sm.mu.Unlock()

			if v, ok := sm.data[key]; ok {
				return v, ok
			} else {
				return 0, ok
			}
		}
	}
}

func problem5() {
	words := []string{"go", "lang", "channel", "gogochamp", "hello", "help"}

	mp := &SafeMap{
		data: make(map[string]int),
		mu:   new(sync.Mutex),
	}
	wg := new(sync.WaitGroup)
	for _, v := range words {
		wg.Add(2)
		go func() {
			defer wg.Done()
			mp.Set(v, rand.Int())
		}()
		go func() {
			defer wg.Done()
			val, ok := mp.Get(v)
			fmt.Printf("Get key %s, val %d, ok? %t\n", v, val, ok)
		}()
	}
	wg.Wait()
}

func main() {
	// problem1()
	// problem2()
	// problem3()
	// problem4()
	problem5()
}
