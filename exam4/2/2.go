package main

import "fmt"

type Formatter interface {
	Format(u User) string
}

type User struct {
	Name string
	Age  int
}

type SimpleFormatter struct{}

func (s SimpleFormatter) Format(u User) string {
	return fmt.Sprintf("Name: %s\nAge: %d\n---\n", u.Name, u.Age)
}

func FormatUsers(users []User, f Formatter) []string {
	a := make([]string, 0, len(users))
	for _, v := range users {
		a = append(a, f.Format(v))
	}
	return a
}

func main() {
	data := []User{
		{Name: "Bob", Age: 19},
		{Name: "Alice", Age: 12},
		{Name: "Anastasia", Age: 54},
		{Name: "Dragon", Age: 1473},
	}
	var f Formatter
	f = SimpleFormatter{}

	formatedData := FormatUsers(data, f)
	for _, v := range formatedData {
		fmt.Print(v)
	}
}
