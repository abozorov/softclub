package main

import (
	"fmt"
	"strings"
)

type Student struct {
	ID    int
	Name  string
	Score int
}

type Error struct {
	Error error
}

func valid(s Student) (bool, error) {
	if (strings.TrimSpace(s.Name) == "") {
		return false, fmt.Errorf("student ID %d: empty name", s.ID)
	} else if s.Score < 0 || s.Score > 100 {
		return false, fmt.Errorf("student ID %d: invalid score", s.ID)
	}
	return true, nil
}

func check(s []Student) (chStud []Student, errs []Error) {
	for _, v := range s {
		if ok, err := valid(v); ok && v.Score >= 60 {
			chStud = append(chStud, v)
		} else if !ok {
			errs = append(errs, Error{Error: err})
		}
	}
	return chStud, errs
}

func main() {
	students := []Student{
		{ID: 1, Name: "Ali", Score: 75},
		{ID: 2, Name: "	", Score: 80},
		{ID: 3, Name: "Vali", Score: 40},
		{ID: 4, Name: "Umed", Score: 90},
	}

	fmt.Println(check(students))
}
