package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func filter(u []User) ([]User, error) {

	if len(u) == 0 {
		return []User{}, errors.New("Пустой слайс")
	}

	fu := make([]User, 0, len(u))

	for _, v := range u {
		if v.Age >= 18 {
			fu = append(fu, v)
		}
	}

	if len(fu) == 0 {
		return []User{}, errors.New("Нет пользователей больше 18 лет")	
	}

	return fu, nil
}

func main() {
	users := []User{
		{ID: 1, Name: "Ali", Age: 17},
		{ID: 2, Name: "Vali", Age: 18},
		{ID: 3, Name: "Umed", Age: 15},
	}
	filtered, err := filter(users)
	
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(filtered)
}
