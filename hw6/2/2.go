package main

import (
	"errors"
	"fmt"
)

type Product struct {
	ID    int
	Name  string
	Price int
}

func filter(p []Product, minPrice int) (int, error) {
	if len(p) == 0 {
		return 0, errors.New("Пустой слайс")
	}
	if minPrice < 0 {
		return 0, errors.New("Минимальная сумма отрицательна")
	}

	count := 0
	for _, v := range p {
		if v.Price < 0 {
			return 0, errors.New("Есть отрицательная цена")
		}
		if v.Price > minPrice {
			count++
		}
	}

	return count, nil
}

func main() {
	products := []Product{
		{ID: 1, Name: "Phone", Price: 1000},
		{ID: 2, Name: "Mouse", Price: 150},
		{ID: 3, Name: "Keyboard", Price: 300},
	}
	minPrice := 200

	filteredCount, err := filter(products, minPrice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(filteredCount)
}
