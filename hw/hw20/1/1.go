package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("Problem1-------------")
	foods := []string{
		"Pizza",
		"Burger",
		"Sushi",
		"Pasta",
	}

	wg := sync.WaitGroup{}
	for _, v := range foods {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("Cooking %s\n", v)
			time.Sleep(1 * time.Second)
			fmt.Printf("%s ready\n", v)
		}()
	}
	wg.Wait()
	fmt.Println("Problem2-------------")

	for i := 0; i < 5; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			fmt.Println("Fast worker")
			time.Sleep(200 * time.Millisecond)
		}()

		go func() {
			defer wg.Done()
			fmt.Println("Slow worker")
			time.Sleep(time.Second)
		}()
	}
	wg.Wait()
	fmt.Println("Problem3-------------")

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Started", i)
			time.Sleep(time.Duration((rand.Int() % 5)) * time.Second)
			fmt.Println("Finished", i)
		}()
	}
	wg.Wait()
	fmt.Println("Problem4-------------")
	grades := []int{
		78, 90, 45, 88,
		67, 100, 56, 72,
	}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(s []int) {
			defer wg.Done()
			sum := 0
			for _, v := range s {
				sum += v
			}
			fmt.Println("Part", i, sum)
		}(grades[i*len(grades)/2 : (i+1)*len(grades)/2])
	}
	wg.Wait()
	fmt.Println("Problem5-------------")
	// Задачи 5-8 похожи на задачу 3
}
