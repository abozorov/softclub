package main

import "fmt"

type Account struct {
    Balance int
}

func deposit(acc *Account, amount int) {
	if amount < 0 {
		return
	}
	acc.Balance += amount
}

func withdraw(acc *Account, amount int) bool {
	if acc.Balance < amount || amount <= 0 {
		return false
	}
	acc.Balance -= amount
	return true
}

func main() {
	var acc Account = Account{
		Balance: 100,
	}

	deposit(&acc, -1)
	fmt.Println(acc)

	deposit(&acc, 10)
	fmt.Println(acc)

	fmt.Println(withdraw(&acc, 111), acc)
	fmt.Println(withdraw(&acc, 110), acc)
	fmt.Println(withdraw(&acc, -5), acc)
}