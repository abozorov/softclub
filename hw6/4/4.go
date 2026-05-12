package main

import (
	"fmt"
	"strings"
)

type Order struct {
	ID         int
	ClientName string
	Amount     int
	Status     string
}

type Error struct {
	Error error
}

func valid(o Order) (bool, error) {
	if strings.TrimSpace(o.ClientName) == "" {
		return false, fmt.Errorf("order ID %d: empty ClientName", o.ID)
	}
	if o.Amount <= 0 {
		return false, fmt.Errorf("order ID %d: Amount <0", o.ID)
	}
	if strings.Compare(o.Status, "new") != 0 &&
		strings.Compare(o.Status, "done") != 0 &&
		strings.Compare(o.Status, "cancelled") != 0 {
		return false, fmt.Errorf("order ID %d: не корректный статус", o.ID)
	}
	return true, nil
}

func chheckOrders(o []Order) (doneOrders, cancledOrders, allOrders int, errs []Error) {

	for _, v := range o {
		if ok, err := valid(v); ok {
			switch v.Status {
			case "done":
				doneOrders ++
				allOrders += v.Amount
			case "cancelled":
				cancledOrders ++
			}
			
		} else {
			errs = append(errs, Error{Error: err})
		}
	}
	return
}

func main() {
	orders := []Order{
		{ID: 1, ClientName: "Ali", Amount: 100, Status: "done"},
		{ID: 2, ClientName: "", Amount: 200, Status: "new"},
		{ID: 3, ClientName: "Vali", Amount: 150, Status: "cancelled"},
		{ID: 4, ClientName: "Umed", Amount: 300, Status: "done"},
		{ID: 5, ClientName: "Salim", Amount: -50, Status: "done"},
	}
	fmt.Println(chheckOrders(orders))
}
