package main

import (
	"errors"
	"fmt"
	"strings"
)

type Profile struct {
	Name  string
	Email string
	Age   int
}

func checkDog(email string) bool {
	rEmail := []rune(email)

	for _, v := range rEmail {
		if v == rune('@') {
			return true
		}
	}
	return false
}

func UpdateProfile(p *Profile, name string, email string, age int) error {
	if p == nil {
		return errors.New("profile is nil")
	}
	name = strings.TrimSpace(name)

	if strings.Compare(name, "") == 0 {
		return errors.New("name is empty")
	}
	email = strings.ToLower(strings.TrimSpace(email))

	if !checkDog(email) {
		return errors.New("invalid email")
	}

	if age < 0 || age > 120 {
		return errors.New("invalid age")
	}

	p.Name = name
	p.Email = email
	p.Age = age
	
	return nil
}

func main() {
	profile := Profile{Name: "Ali", Email: "ali@mail.com", Age: 20}

	err := UpdateProfile(&profile, "  UmeD  ", "  UMED@MAIL.COM  ", 23)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(profile)
}
