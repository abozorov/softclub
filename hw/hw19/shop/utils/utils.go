package utils

import (
	"fmt"
	"hw19/shop/products"
)

func PrintProduct(p products.Product) {
	fmt.Printf("\nНазвание: %s\nЦена: %.2f\nКол-во: %d\n",
		p.Name,
		p.Price,
		p.Quantity)
}
