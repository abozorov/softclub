package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Order struct {
	ID      int
	User    string
	Product string
	Price   float64
	Status  string
}

var (
	Data = make(map[int]*Order)
)

func GetOrder(s string, id int) (*Order, error) {
	fields := strings.Fields(s)
	if len(fields) != 4 {
		return &Order{}, errors.New("Not enougth Data")
	}
	o := Order{
		ID:      id,
		User:    fields[0],
		Product: fields[1],
		Status:  fields[3],
	}
	price, err := strconv.ParseFloat(fields[2], 64)
	if err != nil {
		return &Order{}, err
	}
	o.Price = price
	return &o, nil
}

func Download() error {
	file, err := os.OpenFile("hw17/2/orders.txt", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		order, err := GetOrder(scanner.Text(), i)
		if err != nil {
			return err
		}
		Data[i] = order
	}
	return nil
}

func main() {
	Download()
	// for _, v := range Data {
	// 	fmt.Println(v)
	// }
	var (
		success, fail int
		sumOfSuccess  float64
	)
	usersOrders := make(map[string][]Order)
	sumOfProduct := make(map[string]float64)

	for _, v := range Data {
		switch v.Status {
		case "success":
			success++
			sumOfSuccess += v.Price
		case "failed":
			fail++
		}
		usersOrders[v.User] = append(usersOrders[v.User], *v)
		sumOfProduct[v.Product] += v.Price
	}
	fmt.Printf(`
		сколько успешных заказов: %d
		сколько неуспешных: %d
		общую сумму успешных заказов: %.2f
		`, success, fail, sumOfSuccess)
	for k, v := range usersOrders {
		fmt.Println(k, v)
	}
	for k, v := range sumOfProduct {
		fmt.Println(k, v)
	}
}
