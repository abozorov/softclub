package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// 1. Atomic (sync/atomic)
// 🟢 Лёгкая: «Счётчик кликов»
func problem1() {
	sum := new(atomic.Int64{})
	wg := new(sync.WaitGroup)

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			sum.Add(100)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(sum.Load())
}

// 🟡 Средняя: «Глобальный флаг завершения»
func problem2() {
	stop := new(atomic.Int64)
	wg := new(sync.WaitGroup)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				fmt.Println("load goroutine", i)
				if stop.Load() == 1 {
					break
				}
			}
		}()
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		stop.Add(1)
	}()
	wg.Wait()
}

// 🔒 2. Mutex (sync.Mutex)
// 🟢 Лёгкая: «Безопасное увеличение»
func problem3() {
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	sum := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			sum += 1000
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(sum)
}

// 🟡 Средняя: «Банковский счёт»
type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (ba *BankAccount) Deposit(s int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if ba.mu.TryLock() {
			ba.balance += s
			ba.mu.Unlock()
			break
		}
	}
	fmt.Printf("Баланс был пополнен на %d с, остаток %d\n", s, ba.balance)
}

func (ba *BankAccount) Withdraw(s int, wg *sync.WaitGroup) {
	defer wg.Done()
	ok := true
	for {
		if ba.mu.TryLock() {
			if ba.balance >= s {
				ba.balance -= s
			} else {
				ok = false
				fmt.Println("недостаточно средств для снятия на сумму", s)
			}
			ba.mu.Unlock()
			break
		}
	}
	if ok {
		fmt.Printf("Было снято %d с, остаток %d\n", s, ba.balance)
	}
}

func problem4() {

	acc := BankAccount{
		balance: 0,
		mu:      sync.Mutex{},
	}
	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		wg.Add(2)
		go acc.Deposit(rand.Int()%100, wg)
		go acc.Withdraw(rand.Int()%100, wg)
	}
	wg.Wait()
}

//🔁 3. RWMutex (sync.RWMutex)

func main() {
	// problem1()
	// problem2()
	// problem3()
	problem4()

}
