package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Grade  float64 `json:"grade"`
	Passed bool    `json:"passed"`
}

func Download(s *[]Student) error {
	file, err := os.OpenFile("exam6/1/students.json", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	jDec := json.NewDecoder(file)
	err = jDec.Decode(s)
	if err != nil {
		return err
	}
	return nil
}

func GetMidGrade(s []Student) float64 {
	midGrade := 0.0
	for _, v := range s {
		midGrade += v.Grade
	}
	if len(s) == 0 {
		return 0.0
	}
	midGrade /= float64(len(s))
	return midGrade
}

func GetSuccessList(s []Student) (list []Student) {
	for _, v := range s {
		if v.Passed {
			list = append(list, v)
		}
	}
	return
}

func (s Student) Print() {
	fmt.Printf("name %s\nage %d\ngrade %.2f\n",
		s.Name,
		s.Age,
		s.Grade,
	)
}

func main() {
	students := make([]Student, 0)
	err := Download(&students)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(students)

	fmt.Printf("All students: %d\n", len(students))

	fmt.Printf("Mid grade for all students: %.2f\n", GetMidGrade(students))

	fmt.Println("List of names of students who successfully passed the exam")
	list := GetSuccessList(students)
	for _, v := range list {
		fmt.Printf("\nstudent #%d\n", v.ID)
		v.Print()
	}
}
