package main

import "fmt"

type Employee struct {
	Name   string
	Salary int
}

func maxSalaryEmp(e []Employee) []Employee {
	mxSalary := 0
	top := make([]Employee, 0, len(e))
	for _, v := range e {
		if v.Salary > mxSalary {
			mxSalary = v.Salary
			top = top[:0]
		}

		if v.Salary == mxSalary {
			top = append(top, v)
		}
	}
	return top
}

func midSalaryEmp(e []Employee) []Employee {
	midSalary := 0.0
	for _, v := range e {
		midSalary += float64(v.Salary)
	}
	midSalary /= float64(len(e))

	top := make([]Employee, 0, len(e))
	for _, v := range e {
		if float64(v.Salary) > midSalary {
			top = append(top, v)
		}
	}
	return top
}

func main() {
	employees := []Employee{
		{Name: "Иван", Salary: 50000},
		{Name: "Мария", Salary: 65000},
		{Name: "Алексей", Salary: 45000},
		{Name: "Ольга", Salary: 70000},
		{Name: "Дмитрий", Salary: 55000},
		{Name: "Елена", Salary: 80000},
		{Name: "Павел", Salary: 48000},
		{Name: "Наталья", Salary: 62000},
		{Name: "Артем", Salary: 53000},
		{Name: "Юлия", Salary: 75000},
	}

	fmt.Printf("max salary %v\nover mid salary %v\n", maxSalaryEmp(employees), midSalaryEmp(employees))
}
