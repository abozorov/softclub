package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Contains[T comparable](items []T, value T) bool {
	for _, v := range items {
		if v == value {
			return true
		}
	}
	return false
}

func Min[T int | int64 | float64 | float32](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func problem3() {
	students := make(map[string]int)
	file, err := os.OpenFile("classwork7/1/grades.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewScanner(file)

	for reader.Scan() {
		line := strings.Fields(reader.Text())
		// fmt.Println(line)
		students[line[0]], _ = strconv.Atoi(line[1])
	}

	for k, v := range students {
		if v > 60 {
			fmt.Println(k, v)
		}
	}
}

type User struct {
	Name   string
	Role   string
	Active bool
}

func CanAccess(user User, requiredRole string) error {
	if !user.Active {
		return errors.New("User Not Active")
	}
	if user.Role != requiredRole {
		return errors.New("user role is not suitable")
	}
	return nil
}

func problem4() {
	users := []User{
		{Name: "SuperAdmin", Role: "admin", Active: true},
		{Name: "Alice Smith", Role: "manager", Active: true},
		{Name: "Bob Johnson", Role: "user", Active: true},
		{Name: "Charlie Brown", Role: "user", Active: false},
		{Name: "Elena Gilbert", Role: "manager", Active: false},
		{Name: "John Wick", Role: "admin", Active: true},
		{Name: "Dmitry", Role: "user", Active: true},
	}

	for _, v := range users {
		fmt.Printf("User: %s\nCanAccess: %v\n\n", v.Name, CanAccess(v, "admin"))
	}
}

func main() {
	fmt.Println(Contains([]int{1, 2, 3}, 2))                // true
	fmt.Println(Contains([]string{"go", "java"}, "python")) // false
	fmt.Println(Min(10, 5))                                 // 5
	fmt.Println(Min(3.2, 7.1))                              // 3.2

	problem3()
	problem4()

}
