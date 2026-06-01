package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func sendRequest(client *http.Client, url string, workers chan<- struct{}) {
	defer func() {
		workers <- struct{}{}
	}()
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Ошибка при отправке запроса:", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Ошибка при чтении тела ответа:", url, err)
		return
	}
	fmt.Printf("Url %s\nStatuscode: %d\nBody: %v\n\n",
		url,
		resp.StatusCode,
		// resp.Header,
		string(body[:50]),
	)
}

func main() {
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	time.Sleep(time.Second*5)
	urls := []string{
		"https://google.com",
		"https://api.github.com",
		"https://golang.org",
		"https://jsonplaceholder.typicode.com/users",
		"https://reqres.in/api/users",
		"https://httpbin.org/get",
		"https://pokeapi.co/api/v2/pokemon/pikachu",
	}
	workers := make(chan struct{}, 3)

	for _, v := range urls {
		go sendRequest(client, v, workers)
	}

	for ok := true; ok; {
		_, ok = <-workers
	}
}
