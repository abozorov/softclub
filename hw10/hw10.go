package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func problem1() {
	var str string
	fmt.Scan(&str)

	fmt.Println(func(s string) bool {
		if utf8.RuneCountInString(s) < 6 {
			return false
		}
		d, l := 0, 0
		for _, v := range s {
			if unicode.IsDigit(v) {
				d++
			}
			if unicode.Is(unicode.Latin, v) {
				l++
			}
		}
		return d > 0 && l > 0
	}(str))
}

func problem2() {
	words := []string{"go", "backend", "api", "database", "sql"}
	n := 3
	fmt.Println(func(w []string, n int) int {
		count := 0
		for _, v := range w {
			if utf8.RuneCountInString(v) > n {
				count++
			}
		}
		return count
	}(words, n))
}

func problem3() {
	nums := []int{4, -2, 7, -9, 0, 3}
	fmt.Println(func(a []int) []int {
		b := make([]int, 0, len(a))
		for _, v := range a {
			if v >= 0 {
				b = append(b, v)
			}
		}
		return b
	}(nums))
}

func problem4() {
	names := []string{"ali", "UMED", " rustam ", "SaId"}
	fmt.Println(func(a []string) []string {
		for i := range a {
			a[i] = strings.ToLower(strings.TrimSpace(a[i]))
			b := []rune(a[i])
			b[0] = unicode.ToUpper(b[0])
			a[i] = string(b)
		}

		return a
	}(names))
}

func problem5() {
	s := "a1b2277c3 5"
	fmt.Println(func(s string) int {
		sum := 0
		for _, v := range s {
			if unicode.IsDigit(rune(v)) {
				sum += int(int(rune(v)) - 48)
			}
		}
		return sum
	}(s))
}

func problem6() {
	s := "aa1b2277c3 5"
	fmt.Println(func(s string) string {
		us := make(map[rune]int)
		for _, v := range s {
			us[v]++
		}
		for _, v := range s {
			if us[v] == 1 {
				return string(v)
			}
		}
		return "not found"
	}(s))
}

func problem7() {
	add := func(a, b int) int {
		return a + b
	}

	sub := func(a, b int) int {
		return a - b
	}

	mul := func(a, b int) int {
		return a * b
	}

	div := func(a, b int) (int, error) {
		if b == 0 {
			return 0, errors.New("Divide by zero")
		}
		return a / b, nil
	}

	var (
		a, b int
		o    string
	)
	fmt.Scan(&a, &b, &o)

	switch o {
	case "+":
		fmt.Println(add(a, b))
	case "-":
		fmt.Println(sub(a, b))
	case "*":
		fmt.Println(mul(a, b))
	case "/":
		ans, err := div(a, b)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(ans)
	default:
		fmt.Println("error")
	}
}

type Order struct {
	ID     int
	Amount int
	Status string
}

func problem8() {
	orders := []Order{{1, 500, "new"}, {2, 1200, "paid"}, {3, 300, "paid"}, {4, 900, "failed"}}

	fmt.Println(func(o []Order) int {
		paid := 0
		for _, v := range o {
			if v.Status == "paid" {
				paid += v.Amount
			}
		}
		return paid
	}(orders))
}

func problem9() {
	id := func() func() int {
		id := 0
		return func() int {
			id++
			return id
		}
	}()
	fmt.Println(id())
	fmt.Println(id())
	fmt.Println(id())
}

type User struct {
	Name string
	Age  int
}

func problem10() {
	users := []User{{"Ali", 17}, {"Umed", 20}, {"Rustam", 15}, {"Said", 22}}
	fmt.Printf("%v", func(u []User) []User {
		f := make([]User, 0, len(u))
		for _, v := range u {
			if v.Age >= 18 {
				f = append(f, v)
			}
		}
		return f
	}(users))
}

func main() {
	// problem1()
	// problem2()
	// problem3()
	// problem4()
	// problem5()
	// problem6()
	// problem7()
	// problem8()
	// problem9()
	// problem10()
}
