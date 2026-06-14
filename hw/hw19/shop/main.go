package main

import (
	"fmt"
	"hw19/shop/products"
	"hw19/shop/utils"
)

func main() {
	list := []products.Product{
		{Name: "Laptop Pro 14", Price: 1299.99, Quantity: 8},
		{Name: "Wireless Mouse", Price: 29.50, Quantity: 45},
		{Name: "Mechanical Keyboard", Price: 89.99, Quantity: 20},
		{Name: "4K Monitor 27\"", Price: 349.00, Quantity: 12},
		{Name: "USB-C Cable 2m", Price: 15.95, Quantity: 100},
		{Name: "External SSD 1TB", Price: 110.00, Quantity: 25},
		{Name: "Bluetooth Speaker", Price: 59.99, Quantity: 30},
		{Name: "HD Webcam", Price: 65.50, Quantity: 15},
		{Name: "Desk Lamp LED", Price: 24.99, Quantity: 40},
		{Name: "Gaming Headset", Price: 79.90, Quantity: 18},
		{Name: "Power Bank 20k", Price: 42.00, Quantity: 60},
		{Name: "Laptop Stand", Price: 35.00, Quantity: 22},
		{Name: "Ergonomic Chair", Price: 249.99, Quantity: 5},
	}
	for _, v := range list {
		utils.PrintProduct(v)
	}

	fmt.Printf("\nall price : %.2f", products.TotalPrice(list))
	fmt.Println("\n--------------------\nПодробную информацию о самом дорогом товаре")
	utils.PrintProduct(products.MostExpensive(list))
}
