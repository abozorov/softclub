package main

import (
	"fmt"
	"hw19/jsonapp/models"
	"hw19/jsonapp/storage"
)

func main() {
	users := []models.User{
		{Name: "Alice Smith", Age: 25},
		{Name: "Bob Johnson", Age: 34},
		{Name: "Charlie Brown", Age: 22},
		{Name: "Diana Prince", Age: 29},
		{Name: "Evan Wright", Age: 41},
		{Name: "Fiona Gallagher", Age: 19},
		{Name: "George Brooks", Age: 53},
		{Name: "Hannah Abbott", Age: 27},
		{Name: "Ian Malcolm", Age: 45},
		{Name: "Julia Roberts", Age: 31},
		{Name: "Kevin Mitnick", Age: 38},
		{Name: "Laura Croft", Age: 26},
		{Name: "Michael Scott", Age: 43},
		{Name: "Natalie Portman", Age: 35},
		{Name: "Oscar Martinez", Age: 48},
		{Name: "Pam Beesly", Age: 28},
		{Name: "Quentin Tarantino", Age: 50},
		{Name: "Rachel Green", Age: 24},
		{Name: "Sam Winchester", Age: 33},
		{Name: "Tina Turner", Age: 60},
		{Name: "Ulysses Grant", Age: 47},
		{Name: "Victor Vance", Age: 36},
		{Name: "Wendy Darling", Age: 21},
		{Name: "Xavier Charles", Age: 55},
		{Name: "Yisra Khan", Age: 23},
		{Name: "Zachary Levi", Age: 39},
		{Name: "Arthur Dent", Age: 30},
		{Name: "Bruce Wayne", Age: 32},
		{Name: "Clark Kent", Age: 27},
		{Name: "Emma Watson", Age: 31},
	}

	err := storage.SaveUsers(users)
	if err != nil {
		fmt.Println(err)
		return
	}

	newUsers, err := storage.LoadUsers()
	for _, v := range newUsers {
		fmt.Printf("-----\nUser name: %s\nUSerAge: %d\n",
			v.Name,
			v.Age,
		)
	}
}
