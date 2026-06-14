package main

import (
	"fmt"
	"math/rand"
	"strconv"
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
	ba.mu.Lock()
	ba.balance += s
	ba.mu.Unlock()

	fmt.Printf("Баланс был пополнен на %d с, остаток %d\n", s, ba.balance)
}

func (ba *BankAccount) Withdraw(s int, wg *sync.WaitGroup) {
	defer wg.Done()
	ok := true
	ba.mu.Lock()
	if ba.balance >= s {
		ba.balance -= s
	} else {
		ok = false
		fmt.Println("недостаточно средств для снятия на сумму", s)
	}
	ba.mu.Unlock()

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

// 🔁 3. RWMutex (sync.RWMutex)
// 🟢 Лёгкая: «Чтение и запись»
func problem5() {
	mu := new(sync.RWMutex)
	wg := new(sync.WaitGroup)
	number := 0
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.RLock()
			fmt.Println(i, "number:", number)
			mu.RUnlock()
		}()
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			mu.Lock()
			number = rand.Int()
			mu.Unlock()
		}
	}()
	wg.Wait()
	fmt.Println("number:", number)
}

// 🟡 Средняя: «Конфигурация сервера»
type Config struct {
	data map[string]string
	mu   sync.RWMutex
}

func (c *Config) Get(key string) (string, bool) {
	c.mu.RLock()
	inf, ok := c.data[key]
	c.mu.RUnlock()
	return inf, ok
}

func (c *Config) Set(key, value string) {
	c.mu.Lock()
	c.data[key] = value
	c.mu.Unlock()
}

func problem6() {
	conf := Config{
		data: make(map[string]string),
		mu:   sync.RWMutex{},
	}
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			key := strconv.Itoa(rand.Int() % 5)
			v, ok := conf.Get(key)
			fmt.Printf("data: key %s, val %s, %v\n", key, v, ok)
		}()
	}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				key := strconv.Itoa(rand.Int() % 5)
				conf.Set(key, strconv.Itoa(rand.Int()%78))
			}
		}()
	}
	wg.Wait()
}

// 🗺 4. sync.Map
// 🟢 Лёгкая: «Кеш пользователей»
func problem7() {
	users := new(sync.Map)
	users.Store(1, "Alice")
	users.Store(2, "Bob")
	users.Store(3, "Charlie")
	users.Store(4, "David")
	users.Store(5, "Eve")

	users.Range(func(key, value any) bool {
		fmt.Printf("ID: %v, User: %v\n", key, value)
		return true
	})
}

// 🟡 Средняя: «Конкурентный логин-кеш»
func problem8() {
	users := new(sync.Map)
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		// горутина выполняющее какое либо рандомное действие из 3
		go func() {
			defer wg.Done()
			action := rand.Int() % 3
			switch action {
			case 0:
				users.Store(i, "Alice")
				fmt.Printf("action: %d, Zapisano id %d, val %s\n", action, i, "Alice")
			case 1:
				id := rand.Int() % 10
				if val, ok := users.LoadOrStore(id, "Bob"); ok {
					fmt.Printf("action: %d,Zapisano id %d, val %s\n", action, id, val)
				} else {
					fmt.Printf("action: %d, Zagruzheno %d, val %s\n", action, id, val)
				}
			default:
				id := rand.Int() % 10
				if val, ok := users.Load(id); ok {
					fmt.Printf("action: %d, Zagruzheno id %d, val %s\n", action, id, val)
				} else {
					fmt.Printf("action: %d, Net val vith id %d\n", action, id)
				}
			}
		}()
	}
	wg.Wait()
}

// ⏱ 5. sync.Once
// 🟢 Лёгкая: «Инициализация конфигурации»
func initConfig() {
	fmt.Println("config loaded")
}

func problem9() {
	once := sync.Once{}
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(initConfig)
		}()
	}
	wg.Wait()
}

// 🟡 Средняя: «Ленивая инициализация БД»
type DBConnect struct {
	Connect bool
}

var (
	Connection = DBConnect{
		Connect: false,
	}
)

func initDB() {
	for i := 0; i < 2; i++ {
		time.Sleep(time.Second)
		fmt.Println("Инициализация", 2-i, "сек")
	}
	Connection.Connect = true
}

func problem10() {
	once := sync.Once{}
	wg := new(sync.WaitGroup)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(initDB)

			if Connection.Connect {
				fmt.Printf("Connetion to DB is aviable! %d\n", i)
			}
		}()
	}
	wg.Wait()
}

func main() {
	// problem1()
	// problem2()
	// problem3()
	// problem4()
	// problem5()
	// problem6()
	// problem7()
	// problem8()
	// problem9()
	// problem10()
}
