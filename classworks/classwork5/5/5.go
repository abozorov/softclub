package main

import (
	"fmt"
)

type Customer struct {
	ID   int
	Name string
	City string
}

var (
	customers = map[int]*Customer{}
)

func findCustomer(i int) Customer {
	return *customers[i]
}

func deleteCustomer(i int) {
	delete(customers, i)
}

func (c *Customer) updateCity(city string) {
	c.City = city
}

func main() {
	data := []Customer{
		{ID: 1, Name: "John Doe", City: "New York"},
		{ID: 2, Name: "Alice Smith", City: "London"},
		{ID: 3, Name: "Bob Johnson", City: "San Francisco"},
		{ID: 4, Name: "Emma Wilson", City: "Berlin"},
		{ID: 5, Name: "Michael Brown", City: "Chicago"},
		{ID: 6, Name: "Sophia Garcia", City: "Madrid"},
		{ID: 7, Name: "James Miller", City: "Toronto"},
		{ID: 8, Name: "Olivia Davis", City: "Sydney"},
		{ID: 9, Name: "William Taylor", City: "Austin"},
		{ID: 10, Name: "Isabella Martinez", City: "Mexico City"},
	}

	for _, c := range data {
		customers[c.ID] = &c
	}
	fmt.Printf("%v", findCustomer(7))
	deleteCustomer(2)
	customers[5].updateCity("Dushanbe")

	for _, c := range customers {
		if c.City == "Dushanbe" {
			fmt.Println(*c)
		}
	}
}
