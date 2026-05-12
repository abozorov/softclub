package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"
)

var (
	orders   = make(map[int]*Order, 0)
	fileName = "exam5/1/orders.txt"
)

type Order struct {
	ID        int
	Customer  string
	Products  []string
	Total     float64
	CreatedAt time.Time
}

func NewOrder() *Order {
	return &Order{
		Products: make([]string, 0),
	}
}

func (o *Order) PrintOrder() {
	fmt.Printf("Order ID #%d\nCustomer: %s\nPoducts: %v\nTotal sum: %.2f\nCreated: %s\n\n", o.ID, o.Customer, o.Products, o.Total, o.CreatedAt.Format(time.UnixDate))
}

func AddOrder(c string, p []string, t float64) {
	ord := NewOrder()
	ord.ID = len(orders)
	ord.Customer = c
	ord.Products = p
	ord.Total = t
	ord.CreatedAt = time.Now()
	orders[ord.ID] = ord
}

func PrintAll() error {
	if len(orders) == 0 {
		return errors.New("no orders")
	}

	for _, v := range orders {
		fmt.Printf("Order ID #%d\nCustomer: %s\nPoducts: %v\nTotal sum: %.2f\nCreated: %s\n\n", v.ID, v.Customer, v.Products, v.Total, v.CreatedAt.Format(time.UnixDate))
	}
	return nil
}

func GetExpenciveOrder() (o *Order, e error) {
	o = NewOrder()
	for _, v := range orders {
		if o.Total < v.Total {
			o = v
		}
	}
	return
}

func GetOrdersSum() float64 {
	sum := 0.0
	for _, v := range orders {
		sum += v.Total
	}
	return sum
}

func ClientWithManyOrders() (o *Order, e error) {
	o = NewOrder()
	for _, v := range orders {
		if len(o.Products) < len(v.Products) {
			o = v
		}
	}
	return
}

func WriteFile() {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	wrtr := bufio.NewWriter(file)

	for _, v := range orders {
		wrtr.WriteString(fmt.Sprintf("Order ID #%d\nCustomer: %s\nPoducts: %v\nTotal sum: %.2f\nCreated: %s\n\n", v.ID, v.Customer, v.Products, v.Total, v.CreatedAt.Format(time.UnixDate)))
		wrtr.Flush()
	}
}

func main() {

	AddOrder("Bob", []string{"Pencil", "Book", "Notebook", "Table"}, 31514.01)
	AddOrder("Alice", []string{"Poco x7"}, 445.71)
	AddOrder("Georg", []string{"Watch", "Book", "pencil"}, 88771.01)
	PrintAll()

	o, err := GetExpenciveOrder()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("----\nсамый дорогой заказ:")
	o.PrintOrder()

	fmt.Println("Orders sum: ", GetOrdersSum())

	o, err = ClientWithManyOrders()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("----\nКлиента с самым большим количеством заказов:")
	o.PrintOrder()

	WriteFile()
}
