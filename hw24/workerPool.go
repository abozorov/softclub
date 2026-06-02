package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func worker(client *http.Client, task <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for url := range task {
		req, _ := http.NewRequest("GET", url, nil)
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Ошибка при отправке запроса:", url, err)
			continue
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Ошибка при чтении тела ответа:", url, err)
			continue
		}
		fmt.Printf("Url %s\nStatuscode: %d\nBody: %v\n\n",
			url,
			resp.StatusCode,
			// resp.Header,
			string(body[:50]),
		)
	}
}

func main() {
	client := &http.Client{
		Timeout: time.Second * 2,
	}

	urls := []string{
		"https://google.com",
		"https://api.github.com",
		"https://golang.org",
		"https://jsonplaceholder.typicode.com/users",
		"https://reqres.in/api/users",
		"https://httpbin.org/get",
		"https://pokeapi.co/api/v2/pokemon/pikachu",
	}
	tasks := make(chan string, 3)
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(client, tasks, &wg)
	}

	for _, url := range urls {
		tasks <- url
	}
	close(tasks)
	wg.Wait()
}
