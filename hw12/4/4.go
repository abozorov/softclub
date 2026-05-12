package main

import (
	"errors"
	"fmt"
)

type OrderProcessor interface {
	Process(order Order) error
}

type Order struct {
	ID     int
	Status string
}

type NewOrderProcessor struct{}
type CancelOrderProcessor struct{}
type CompleteOrderProcessor struct{}

func (nop NewOrderProcessor) Process(order Order) error {
	if order.Status == "new" {
		return nil
	} else {
		return errors.New("NewOrderProcessor not new")
	}
}

func (caop CancelOrderProcessor) Process(order Order) error {
	if order.Status == "new" || order.Status == "processing" {
		return nil
	} else {
		return errors.New("CancelOrderProcessor not new || not processing")
	}
}

func (coop CompleteOrderProcessor) Process(order Order) error {
	if order.Status == "processing" {
		return nil
	} else {
		return errors.New("CompleteOrderProcessor not processing")
	}
}

func HandleOrder(p OrderProcessor, order *Order) {
	err := p.Process(*order)

	if err == nil {
		fmt.Println("OK")
	} else {
		fmt.Println(err)
	}
}

func main() {
	var p OrderProcessor
	orders := []Order{
		{ID: 1, Status: "new"},
		{ID: 2, Status: "processing"},
		{ID: 3, Status: "new"},
		{ID: 4, Status: "processing"},
		{ID: 5, Status: "processing"},
		{ID: 6, Status: "new"},
	}

	p = NewOrderProcessor{}
	// p = CancelOrderProcessor{}
	// p = CompleteOrderProcessor{}
	for _, v := range orders {
		HandleOrder(p, &v)
	}
}
