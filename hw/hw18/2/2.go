package main

import (
	"bufio"
	"fmt"
	"os"
)

type Product struct {
	Name     string
	Category string
	Price    float64
	Quantity int
}

func (p Product) Print() string {
	return fmt.Sprintf("\nТовар %s\nCategory %s\nЦена %.2f\nКол-во %d\n",
		p.Name,
		p.Category,
		p.Price,
		p.Quantity)
}

func TotalPrice(products []Product) float64 {
	tP := 0.0
	for _, v := range products {
		tP += v.Price * float64(v.Quantity)
	}
	return tP
}

func MostExpensive(products []Product) Product {
	mxPP := new(Product)
	for _, v := range products {
		if mxPP.Price < v.Price {
			mxPP = &v
		}
	}
	return *mxPP
}

func CategoryStats(products []Product) map[string]int {
	stats := make(map[string]int)
	for _, v := range products {
		stats[v.Category] += v.Quantity
	}
	return stats
}

func DataSaver(stats map[string]int) error {
	file, err := os.OpenFile("hw18/2/stats.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	wrtr := bufio.NewWriter(file)
	for category, count := range stats {
		wrtr.WriteString(fmt.Sprintf("Category: %s | Count: %d\n", category, count))
		wrtr.Flush()
	}
	return nil
}

func main() {
	products := []Product{
		// Electronics
		{Name: "Smartphone X", Category: "Electronics", Price: 799.99, Quantity: 15},
		{Name: "Wireless Headphones", Category: "Electronics", Price: 149.50, Quantity: 40},
		{Name: "Mechanical Keyboard", Category: "Electronics", Price: 89.99, Quantity: 25},
		{Name: "4K Monitor 27\"", Category: "Electronics", Price: 329.00, Quantity: 10},
		{Name: "Gaming Mouse", Category: "Electronics", Price: 59.99, Quantity: 50},
		{Name: "USB-C Hub", Category: "Electronics", Price: 34.95, Quantity: 100},

		// Apparel
		{Name: "Cotton T-Shirt", Category: "Apparel", Price: 19.99, Quantity: 200},
		{Name: "Denim Jeans", Category: "Apparel", Price: 49.99, Quantity: 85},
		{Name: "Hoodie", Category: "Apparel", Price: 39.95, Quantity: 60},
		{Name: "Running Shoes", Category: "Apparel", Price: 89.00, Quantity: 35},
		{Name: "Socks (5-pack)", Category: "Apparel", Price: 12.50, Quantity: 150},
		{Name: "Winter Jacket", Category: "Apparel", Price: 120.00, Quantity: 20},

		// Groceries
		{Name: "Organic Milk", Category: "Groceries", Price: 3.49, Quantity: 80},
		{Name: "Whole Wheat Bread", Category: "Groceries", Price: 2.99, Quantity: 45},
		{Name: "Coffee Beans 1kg", Category: "Groceries", Price: 24.99, Quantity: 30},
		{Name: "Green Tea Box", Category: "Groceries", Price: 5.50, Quantity: 90},
		{Name: "Olive Oil 500ml", Category: "Groceries", Price: 14.20, Quantity: 40},
		{Name: "Dark Chocolate", Category: "Groceries", Price: 3.99, Quantity: 120},

		// Home & Kitchen
		{Name: "Blender 800W", Category: "Home", Price: 65.00, Quantity: 18},
		{Name: "Ceramic Frying Pan", Category: "Home", Price: 29.99, Quantity: 30},
		{Name: "Chef's Knife", Category: "Home", Price: 45.50, Quantity: 22},
		{Name: "Desk Lamp LED", Category: "Home", Price: 24.99, Quantity: 55},
		{Name: "Water Filter Pitcher", Category: "Home", Price: 19.95, Quantity: 70},
		{Name: "Memory Foam Pillow", Category: "Home", Price: 34.99, Quantity: 40},

		// Sports & Outdoors
		{Name: "Yoga Mat", Category: "Sports", Price: 25.00, Quantity: 65},
		{Name: "Dumbbells Set (10kg)", Category: "Sports", Price: 45.00, Quantity: 15},
		{Name: "Water Bottle 1L", Category: "Sports", Price: 15.99, Quantity: 110},
		{Name: "Camping Tent 2-Person", Category: "Sports", Price: 95.00, Quantity: 12},
		{Name: "Backpack 30L", Category: "Sports", Price: 49.95, Quantity: 28},
		{Name: "Bicycle Helmet", Category: "Sports", Price: 35.00, Quantity: 20},
	}
	fmt.Printf("\nAll products price : %.2f\n", TotalPrice(products))
	fmt.Println("\n--------------------\nПодробную информацию о самом дорогом товаре", MostExpensive(products).Print())
	err := DataSaver(CategoryStats(products))
	if err != nil {
		fmt.Println(err)
		return
	}
}
