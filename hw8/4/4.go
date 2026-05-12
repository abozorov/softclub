package main

import "fmt"

type Order struct {
	ID     int
	UserID int
	Amount int
	Status string
}

func BuildUserStats(orders []Order) map[int]map[string]int {
	stats := make(map[int]map[string]int)

	for _, v := range orders {
		if _, ok := stats[v.UserID]; !ok {
			stats[v.UserID] = make(map[string]int)
		}
		stats[v.UserID][v.Status] += v.Amount
	}
	return stats
}

func main() {
	orders := []Order{
		{1, 10, 100, "success"},
		{2, 10, 50, "failed"},
		{3, 20, 300, "success"},
		{4, 30, 200, "success"},
	}

	fmt.Printf("%v\n", BuildUserStats(orders))
}
