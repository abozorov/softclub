package main

import (
	"fmt"
)

type Address struct {
	City string
}

type User struct {
	Address
	Name string
	Age  int
}

func findAdults(users []User) []User {
	older18 := make([]User, 0, len(users))

	for _, v := range users {
		if v.Age >= 18 {
			older18 = append(older18, v)
		}
	}
	return older18
}

func increaseMinors(users []User) {

	for i := range users {
		if users[i].Age < 18 {
			users[i].Age++
		}
	}
}

func countFromCity(users []User, city string) int {
	var count int

	for i := range users {
		if users[i].City == city {
			count++
		}
	}
	return count
}

func (u User) isAdult() bool {
	return u.Age >= 18
}

func renameAll(users []*User, newName string) {
	for _, v := range users {
		v.Name = newName // Автоматически выходит из поинтера
	}
}

func getAverageAge(users []User) float64 {
	var sr float64
	for _, v := range users {
		sr += float64(v.Age)
	}
	return sr / float64(len(users))
}

func main() {

	users := []User{
		{
			Name: "Ali",
			Age:  17,
			Address: Address{
				City: "Vahdat",
			},
		}, {
			Name: "Umed",
			Age:  23,
			Address: Address{
				City: "Kulob",
			},
		}, {
			Name: "Zafar",
			Age:  15,
			Address: Address{
				City: "Dushanbe",
			},
		}, {
			Name: "Bob",
			Age:  22,
			Address: Address{
				City: "Kulob",
			},
		},
	}
	fmt.Println(findAdults(users))

	increaseMinors(users)
	fmt.Println(users)

	fmt.Println(countFromCity(users, "Kulob"))
	fmt.Println(users[0].isAdult())

	p_users := []*User{
		{Name: "Ali", Age: 20}, // автоматически берет ссылку
		{Name: "Umed", Age: 23},// автоматически берет ссылку
	}
	renameAll(p_users, "B.I.G")
	
	for _, v := range p_users {
		fmt.Println(v)
	}

	fmt.Println(getAverageAge(users))
}
