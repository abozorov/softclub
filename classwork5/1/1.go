package main

import "fmt"

type User struct {
	Login    string
	Password string
}

var (
	users map[string]User
)

func Login(l, p string) bool {
	_, ok := users[l]; 
	return ok && users[l].Password == p
}

func main() {
	users = map[string]User{
		"aaa": {
			"aaa",
			"aaapass",
		},
		"bbb": {
			"bbb",
			"98744",
		},
	}

	for {
		var login, password string
		fmt.Scan(&login, &password)

		if Login(login, password) {
			fmt.Printf("user %s logined\n", login)
			break
		} else {
			fmt.Println("wrong login or password")
		}
	}
}
