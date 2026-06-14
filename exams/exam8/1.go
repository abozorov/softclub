package main

import (
	"errors"
	"fmt"
)

type PaymentProcessor interface {
	Process(amount float64) error
	GetBalance() float64
}

// card payment
type CardPayment struct {
	Balance float64
}

func (card *CardPayment) Process(amount float64) error {
	if amount <= 0 {
		return errors.New("Amount <= 0")
	}
	card.Balance += amount
	return nil
}

func (cash *CardPayment) GetBalance() float64 {
	return cash.Balance
}

// cash payment
type CashPayment struct {
	Balance float64
}

func (cash *CashPayment) Process(amount float64) error {
	if amount <= 0 {
		return errors.New("Amount <= 0")
	}
	cash.Balance += amount
	return nil
}

func (cash *CashPayment) GetBalance() float64 {
	return cash.Balance
}

// run payment
func Pay(p PaymentProcessor, amount float64) error {
	if err := p.Process(amount); err != nil {
		return fmt.Errorf("Func Pay: %w", err)
	}
	return nil
}

func main() {
	payments := []PaymentProcessor{
		&CardPayment{
			Balance: 99,
		},
		&CashPayment{
			Balance: 99,
		},
	}
	amounts := []float64{100, 50, -10}

	for _, v := range amounts {
		for _, vP := range payments {
			err := Pay(vP, v)

			if !errors.Is(nil, err) {
				fmt.Printf("Error %s\nPayment %T\nAmount %.2f\nBalance %.2f\n\n",
					err.Error(),
					vP,
					v,
					vP.GetBalance(),
				)
			} else {
				fmt.Printf("SUCCESS\nPayment %T\nAmount %.2f\nBalance %.2f\n\n",
					vP,
					v,
					vP.GetBalance(),
				)
			}
		}
	}
}
