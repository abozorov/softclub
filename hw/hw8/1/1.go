package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

type User struct {
	ID       int
	Username string
}

func AddUser(users []User, username string, id int) ([]User, error) {
	username = strings.ToLower(strings.TrimSpace(username))

	if utf8.RuneCountInString(username) == 0 {
		return users, errors.New("username is empty")
	}

	if utf8.RuneCountInString(username) < 3 {
		return users, errors.New("username is too short")
	}

	for _, v := range users {
		if strings.Compare(v.Username, username) == 0 {
			return users, errors.New("username already exists")
		}
	}
	return append(users, User{ID: id, Username: username}), nil
}

func main() {
	users := []User{
		{1, "ali"},
		{2, "umed"},
	}
	username := "  rTT  e  "

	users, err := AddUser(users, username, 3)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(users)
	}
}
