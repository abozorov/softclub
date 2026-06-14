package main

import (
	"fmt"
	"unicode/utf8"
)

type User struct {
    Name string
    Age  int
}

func updateUser(u *User, name string, age int) {
	
	if utf8.RuneCountInString(name) > 0 {
		u.Name = name
	}
	if age >= 0 {
		u.Age = age
	}
}

func isAdult(u User) bool {
	return u.Age >= 18
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

	fmt.Println(isAdult(users[2]), isAdult(users[1]))

	updateUser(&users[0], "Firdavs", 18)

	fmt.Println(users)
}