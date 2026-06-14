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

func printUser(u User) {
	fmt.Printf("Name: %s, Age: %d, Adress: %s\n", u.Name, u.Age, u.City)
}

func increaseAge(u *User) {
	u.Age++
}

func (u *User) rename(newName string) {
	u.Name = newName
}

func changeCity(u *User, newCity string) {
	u.City = newCity
}

func growAll(users []User) {
	for i := range users {
		users[i].Age++
	}
}

func findOldest(users []User) User {
	oldest := User{}
	for _, v := range users {
		if oldest.Age < v.Age {
			oldest = v
		}
	}

	return oldest
}

func main() {
	// u := User{
	// 	Name: "Barashek",
	// 	Age: 5,
	// 	Adress: Address{
	// 		City: "Zarafshon 2",
	// 	},
	// }

	// printUser(u)
	// increaseAge(&u)
	// printUser(u)
	// u.rename("Bob")
	// printUser(u)
	// changeCity(&u, "101 mkr")

	users := []User{
		{Name: "Ali", Age: 20},
		{Name: "Umed", Age: 23},
		{Name: "Zafar", Age: 19},
	}
	growAll(users)
	for _, v := range users {
		printUser(v)
	}

	printUser(findOldest(users))
}
