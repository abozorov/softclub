package main

import (
	"fmt"
	"sort"
	"strconv"
)

// 🔍 Задача 1. Поиск пропущенного числа
func problem1() {
	var n, sum int
	fmt.Scan(&n)

	for i := 1; i < n; i++ {
		var a int
		fmt.Scan(&a)
		sum += a
	}

	fmt.Println(((1+n)*n)/2 - sum)
}

// 🔠 Задача 2. Первый уникальный символ
func problem2() {
	var word string
	fmt.Scan(&word)
	s := []rune(word)
	us := make(map[rune]int, len(s))

	for _, v := range s {
		us[v]++
	}
	for _, v := range s {
		if us[v] == 1 {
			fmt.Println(string(v))
			return
		}
	}
	fmt.Println("No exist")
}

// 🔄 Задача 3. Проверка анаграммы
func problem3() {
	var w1, w2 string
	fmt.Scan(&w1, &w2)
	s1 := []rune(w1)
	s2 := []rune(w2)

	if len(s1) != len(s2) {
		fmt.Println(false)
		return
	}

	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})

	sort.Slice(s2, func(i, j int) bool {
		return s2[i] < s2[j]
	})

	if string(s1) != string(s2) {
		fmt.Println(false)
		return
	}

	fmt.Println(true)
}

// 🤐 Задача 4. Сжатие строки
func problem4() {
	var s string
	fmt.Scan(&s)

	r := []rune(s)
	ans := make([]rune, 0, 2*len(r))
	k := 1
	r = append(r, '\n')

	for i := 1; i < len(r); i++ {
		if r[i] == r[i-1] {
			k++
		} else {
			ans = append(ans, r[i-1])
			ans = append(ans, []rune(strconv.Itoa(k))...)
			k = 1
		}
	}
	fmt.Println(string(ans))
}

// 👥 Задача 5. Поиск повторяющихся пользователей
func problem5() {
	var n int
	fmt.Scan(&n)
	us := make(map[int]int, n)

	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		us[a]++
	}

	for k, v := range us {
		if v > 1 {
			fmt.Print(k, " ")
		}
	}
	fmt.Println()
}

// 💡 Задача 6. Лампочки
func problem6() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j += i {
			a[j]++
		}
	}

	for i := 1; i <= n; i++ {
		if a[i]%2 > 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

// 🛗 Задача 7. Лифт
func problem7() {
	var n int
	fmt.Scan(&n)
	var mn, mx, cur int
	for ; n > 0; n-- {
		var a int
		fmt.Scan(&a)
		cur += a
		mn = min(mn, cur)
		mx = max(mx, cur)
	}

	fmt.Printf("current floor %d\nmax floor %d\nmin floor %d\n",
		cur,
		mx,
		mn,
	)
}

func main() {
	// problem1()
	// problem2()
	// problem3()
	// problem4()
	// problem5()
	// problem6()
	problem7()

}

// a1, a1+d, a1+2d
// n * a1
