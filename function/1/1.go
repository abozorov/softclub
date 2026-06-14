package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type req struct {
	url         string
	requestType string
}

func worker(client *http.Client, wg *sync.WaitGroup, url, reqType string, clientID int) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// create request
			userID := rand.Int() % 100
			reqBody := bytes.NewBuffer([]byte(fmt.Sprintf("{\"id\":%d,\"name\":\"Bozoorov\"}", userID)))
			newUrl := url
			if clientID == 2 {
				newUrl = url + strconv.Itoa(userID)
			}
			req, _ := http.NewRequest(reqType, newUrl, reqBody)
			req.Header.Add("client_id", strconv.Itoa(clientID))
			req.Header.Add("Authorization", "secret")

			// Do request
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
			fmt.Printf("Url %s\nRequest type: %s\nStatuscode: %d\nBody: %v\n\n",
				newUrl,
				reqType,
				resp.StatusCode,
				// resp.Header,
				string(body),
			)
		}()
	}
}

func main() {
	client := &http.Client{
		Timeout: time.Second * 20,
	}
	wg := sync.WaitGroup{}

	urls := []string{
		"http://localhost:8080/users",
		"http://localhost:8080/user/",
		"http://localhost:8080/user",
	}

	wg.Add(3)
	go worker(client, &wg, urls[0], "GET", 1)
	go worker(client, &wg, urls[1], "GET", 2)
	go worker(client, &wg, urls[2], "POST", 3)

	wg.Wait()
}
