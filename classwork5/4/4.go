package main

import (
	"fmt"
	"unicode"
)

type Client struct {
	Name  string
	Money int
}

func main() {
	clients := []Client{
		{Name: "Олег", Money: 1500},
		{Name: "Anna", Money: 4200},
		{Name: "Виктор", Money: 800},
		{Name: "Марина", Money: 12500},
		{Name: "Arina", Money: 300},
		{Name: "Екатерина", Money: 7800},
		{Name: "Денис", Money: 2100},
		{Name: "Anastasia", Money: 5600},
		{Name: "Михаил", Money: 90},
		{Name: "Татьяна", Money: 3400},
	}

	fmt.Printf("over 1000: %v\nfirst bukva A: %v\n", func(c []Client) []Client {
		a := make([]Client, 0, len(c))
		for _, v := range c {
			if v.Money > 1000 {
				a = append(a, v)
			}
		}
		return a
	}(clients), func(c []Client) []Client {
		a := make([]Client, 0, len(c))
		for _, v := range c {
			if unicode.ToLower([]rune(v.Name)[0]) == 'a' {
				a = append(a, v)
			}
		}
		return a
	}(clients))
}
