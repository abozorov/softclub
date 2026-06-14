package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	usersFile    = "classwork9/1/users.json"
	studentsFile = "classwork9/1/students.json"
	productFile  = "classwork9/1/products.json"
	ordersFile   = "classwork9/1/orders.json"
)

//------------------------------------------------------------------------------------------------------
//------------------------------------------------------------------------------------------------------
//------------------------------------------------------------------------------------------------------

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func hw1() {
	users := []User{
		{Name: "Иван", Age: 25},
		{Name: "Мария", Age: 30},
		{Name: "Алексей", Age: 19},
		{Name: "Ольга", Age: 42},
		{Name: "Дмитрий", Age: 28},
		{Name: "Елена", Age: 35},
	}
	file, err := os.OpenFile(usersFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	flowFile := json.NewEncoder(file)
	flowFile.SetIndent("", "	")
	err = flowFile.Encode(users)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func hw2() {
	users := []User{}
	file, err := os.OpenFile(usersFile, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	flowFile := json.NewDecoder(file)
	err = flowFile.Decode(&users)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(users)
}
func hw12() {
	users := []User{}
	file, err := os.OpenFile(usersFile, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	flowFile := json.NewDecoder(file)
	err = flowFile.Decode(&users)
	if err != nil {
		fmt.Println(err)
		return
	}

	oldest := new(User)
	youngest := new(User)
	youngest.Age = 999
	ageSum := 0.0

	for _, v := range users {
		if oldest.Age < v.Age {
			oldest = &v
		}
		if youngest.Age > v.Age {
			youngest = &v
		}
		ageSum += float64(v.Age)
	}
	fmt.Printf("Oldest User name: %s, Age: %d\n", oldest.Name, oldest.Age)
	fmt.Printf("Youngest User name: %s, Age: %d\n", youngest.Name, youngest.Age)
	fmt.Printf("Mid Age of all users: %.2f\n", ageSum/float64(len(users)))

}

//------------------------------------------------------------------------------------------------------
//------------------------------------------------------------------------------------------------------
//------------------------------------------------------------------------------------------------------

type Student struct {
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func hw3() {
	students := []Student{
		{Name: "Иван", Grade: 41},
		{Name: "Мария", Grade: 71},
		{Name: "Алексей", Grade: 100},
	}
	file, err := os.OpenFile(studentsFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	flowFile := json.NewEncoder(file)
	flowFile.SetIndent("", "	")
	err = flowFile.Encode(students)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func hw4() {
	students := []Student{}
	file, err := os.OpenFile(studentsFile, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	flowFile := json.NewDecoder(file)
	err = flowFile.Decode(&students)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(students) == 0 {
		fmt.Println("mid grade", 0.0)
		return
	}
	gradeSum := 0.0
	for _, v := range students {
		gradeSum += float64(v.Grade)
	}
	fmt.Printf("mid grade %.2f\n", gradeSum/float64(len(students)))
}

func hw5() {
	var name string
	fmt.Scan(&name)
	name = strings.TrimSpace(name)

	students := []Student{}
	file, err := os.OpenFile(studentsFile, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	flowFile := json.NewDecoder(file)
	err = flowFile.Decode(&students)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range students {
		if v.Name == name {
			fmt.Printf("Found: %s - %d\n", v.Name, v.Grade)
			return
		}
	}
	fmt.Println("Student not found")
}

func saveStudents(students []Student) error {
	file, err := os.OpenFile(studentsFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	flowFile := json.NewEncoder(file)
	flowFile.SetIndent("", "	")
	err = flowFile.Encode(students)
	if err != nil {
		return err
	}
	return nil
}

func loadStudents() ([]Student, error) {
	students := []Student{}
	file, err := os.OpenFile(studentsFile, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return []Student{}, err
	}
	defer file.Close()
	flowFile := json.NewDecoder(file)
	err = flowFile.Decode(&students)
	if err != nil {
		return []Student{}, err
	}
	return students, nil
}

func hw6() {
	students, err := loadStudents()
	if err != nil {
		fmt.Println(err)
		return
	}
	newStudent := new(Student)
	_, err = fmt.Scan(&newStudent.Name, &newStudent.Grade)
	if err != nil {
		fmt.Println(err)
		return
	}
	newStudent.Name = strings.TrimSpace(newStudent.Name)
	students = append(students, *newStudent)
	err = saveStudents(students)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func hw7() {
	students, err := loadStudents()
	if err != nil {
		fmt.Println(err)
		return
	}

	var name string
	_, err = fmt.Scan(&name)
	if err != nil {
		fmt.Println(err)
		return
	}
	name = strings.TrimSpace(name)

	ok := true
	for i, v := range students {
		if v.Name == name {
			students = append(students[:i], students[i+1:]...)
			ok = false
			break
		}
	}
	if ok {
		fmt.Println("there is no such student")
		return
	}

	err = saveStudents(students)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Deleted")
}

func hw8() {
	students, err := loadStudents()
	if err != nil {
		fmt.Println(err)
		return
	}
	var sortType string
	fmt.Scan(&sortType)
	switch sortType {
	case "name":
		sort.Slice(students, func(i, j int) bool {
			return students[i].Name < students[j].Name
		})
	case "grade":
		sort.Slice(students, func(i, j int) bool {
			return students[i].Grade < students[j].Grade
		})
	default:
		fmt.Println("there is no such option")
		return
	}
	err = saveStudents(students)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//------------------------------------------------------------------------------------------------------
//------------------------------------------------------------------------------------------------------
//------------------------------------------------------------------------------------------------------

type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func saveProducts(products []Product) error {
	file, err := os.OpenFile(productFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	flowFile := json.NewEncoder(file)
	flowFile.SetIndent("", "	")
	err = flowFile.Encode(products)
	if err != nil {
		return err
	}
	return nil
}

func loadProducts() ([]Product, error) {
	products := []Product{}
	file, err := os.OpenFile(productFile, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return []Product{}, err
	}
	defer file.Close()
	flowFile := json.NewDecoder(file)
	err = flowFile.Decode(&products)
	if err != nil {
		return []Product{}, err
	}
	return products, nil
}

func TotalProductsPrice(products []Product) float64 {
	tP := 0.0
	for _, v := range products {
		tP += v.Price * float64(v.Quantity)
	}
	return tP
}

func hw9() {
	products := []Product{
		{Name: "Smartphone", Price: 699.99, Quantity: 15},
		{Name: "Laptop Pro", Price: 1249.50, Quantity: 7},
		{Name: "Wireless Earbuds", Price: 89.99, Quantity: 42},
		{Name: "Mechanical Keyboard", Price: 110.00, Quantity: 20},
		{Name: "Gaming Mouse", Price: 45.50, Quantity: 55},
		{Name: "4K Monitor 27\"", Price: 299.99, Quantity: 10},
		{Name: "USB-C Hub", Price: 34.95, Quantity: 80},
		{Name: "External SSD 1TB", Price: 95.00, Quantity: 24},
		{Name: "Desk Lamp LED", Price: 25.00, Quantity: 35},
		{Name: "HD Webcam", Price: 59.90, Quantity: 18},
	}
	err := saveProducts(products)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\nTotalProductsPrice : %.2f\n", TotalProductsPrice(products))
}

func MostExpensiveProduct(products []Product) Product {
	mxPP := new(Product)
	for _, v := range products {
		if mxPP.Price < v.Price {
			mxPP = &v
		}
	}
	return *mxPP
}

func hw10() {
	products, err := loadProducts()
	if err != nil {
		fmt.Println(err)
		return
	}
	mostExp := MostExpensiveProduct(products)
	fmt.Println(mostExp)
}

//------------------------------------------------------------------------------------------------------
//------------------------------------------------------------------------------------------------------
//------------------------------------------------------------------------------------------------------

type Order struct {
	ID       int     `json:"id"`
	Customer string  `json:"customer"`
	Total    float64 `json:"total"`
}

func hw11() {
	orders := []Order{
		{ID: 101, Customer: "Alice Smith", Total: 250.50},
		{ID: 102, Customer: "Bob Johnson", Total: 89.99},
		{ID: 103, Customer: "Charlie Brown", Total: 1200.00},
		{ID: 104, Customer: "Diana Prince", Total: 45.00},
		{ID: 105, Customer: "Evan Wright", Total: 310.25},
		{ID: 106, Customer: "Fiona Gallagher", Total: 15.75},
		{ID: 107, Customer: "George Brooks", Total: 620.00},
		{ID: 108, Customer: "Hannah Abbott", Total: 95.40},
		{ID: 109, Customer: "Ian Malcolm", Total: 150.00},
		{ID: 110, Customer: "Julia Roberts", Total: 830.15},
	}
	file, err := os.OpenFile(ordersFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "	")
	err = encoder.Encode(orders)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Вывести заказы
	fmt.Println("All orders: ", orders)

	// найти заказ с максимальной суммой
	maxSumOrder := new(Order)
	for _, v := range orders {
		if maxSumOrder.Total < v.Total {
			maxSumOrder = &v
		}
	}
	fmt.Println("max Sum Order: ", *maxSumOrder)
}

func main() {
	// hw1()
	// hw2()
	// hw3()
	// hw4()
	// hw5()
	// hw6()
	// hw7()
	// hw8()
	// hw9()
	// hw10()
	hw11()
	// hw12()

}
