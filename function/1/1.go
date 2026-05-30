package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "Internal Server Error",
					http.StatusInternalServerError)
				fmt.Println("Panic recovered:", err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var userDto User
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, userDto)
}
	
func users(w http.ResponseWriter, r *http.Request) {
	userSlices := []User{
		{Name: "Ali", Age: 10, Gender: "Male"},
		{Name: "Alisa", Age: 9, Gender: "female"},
	}

	userSlicesData, err := json.Marshal(userSlices)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("ошибка, json.Marshal: ", err)
		return
	}

	_, err = w.Write(userSlicesData)
	if err != nil {
		fmt.Println("ошибка, w.Write: ", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/create/user", createUser)
	mux.HandleFunc("/users", users)
	handler := recoverMiddleware(loggingMiddleware(mux))
	fmt.Println("Сервер запущен на :8080")
	http.ListenAndServe(":8080", handler)
}
