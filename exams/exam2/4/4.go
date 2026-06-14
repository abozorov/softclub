package main

import (
	"fmt"
)

type User struct {
    Name string
    Age  int
}

func addUser(users []User, u User) []User {
	return append(users, u)
}

func updateUserByName(users []User, name string, newName string, newAge int) {

	for i := 0; i < len(users); i++ {
		if users[i].Name == name {
			users[i].Name = newName
			users[i].Age = newAge
			return
		}
	}
}

func printUsers(users []User) {
	for _, v := range users {
		fmt.Printf("Name: %s, Age: %d\n", v.Name, v.Age)
	}
	fmt.Println()
}


func main() {
	var users []User = []User{
		{
			Name: "Bob",
			Age: 17,
		},
		{
			Name: "Jake",
			Age: 19,
		},
		{
			Name: "Danie",
			Age: 11,
		},
	}

	users = addUser(users, User{Name: "Garry", Age: 18})
	printUsers(users)

	updateUserByName(users, "Bob", "FIrdavs", 19)
	printUsers(users)

	updateUserByName(users, "qwe", "hfhhf", 19)
	printUsers(users)
}