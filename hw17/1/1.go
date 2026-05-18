package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Books struct {
	Books []Book `json:"books"`
}

type Book struct {
	ID        int       `json:"id"`
	Title     string    `json:"books"`
	Author    string    `json:"author"`
	Pages     int       `json:"pages"`
	Available bool      `json:"avialable"`
	CreatedAt time.Time `json:"create_at"`
}

type BookManager interface {
	AddBook(book Book) error
	TakeBook(id int) error
	ReturnBook(id int) error
	ShowBooks()
}

// ✅ Добавлять книгу

// 📌 Ошибки:

// * ID не может быть <= 0
// * название пустое
// * количество страниц <= 0

func (book Book) Print() string {
	return fmt.Sprintf(`
  ID: %d,
  Title: %s,
  Author: %s,
  Pages: %d,
  Available: %t,
  Created_at: %s,
  `,
		book.ID,
		book.Title,
		book.Author,
		book.Pages,
		book.Available,
		book.CreatedAt.Format(time.RFC1123),
	)
}

func (b *Books) AddBook(book Book) error {
	if book.ID <= 0 {
		return errors.New("id не может быть равен или меньше нуля")
	}
	if strings.TrimSpace(book.Title) == "" {
		return errors.New("название не можеть быть пустым")
	}
	if book.Pages <= 0 {
		return errors.New("страница не может быть равен или меньше нуля")
	}

	b.Books = append(b.Books, book)
	return nil
}

// ---

// ✅ Брать книгу

// 📌 Условия:

// * если книга уже занята → ошибка
// * если книга не найдена → ошибка
// * если всё успешно:

//   * `Available = false`

func (b *Books) TakeBook(id int) error {
	var isFind bool
	for _, book := range b.Books {
		if book.ID == id {
			if !book.Available { // if book.Available == false {}
				return errors.New("книга не доступна")
			}
			book.Available = false
			isFind = true
		}
	}

	return handleResp(isFind)
}

// ✅ Возвращать книгу

// 📌 Условия:

// * если книги нет → ошибка
// * если книга уже свободна → ошибка

// ---

func (b *Books) ReturnBook(id int) error {
	var isFind bool
	for _, book := range b.Books {
		if book.ID == id {
			if book.Available {
				return errors.New("книга уже свободна")
			}
			book.Available = true
			isFind = true
		}
	}

	return handleResp(isFind)
}

func handleResp(isFind bool) error {
	if !isFind {
		return errors.New("книга не найдена")
	}
	return nil
}

// ✅ Показывать:

// все книги
// только свободные книги
// только занятые книги

func (b *Books) ShowBooks() {
	for _, book := range b.Books {
		fmt.Print(book.Print())
	}
	fmt.Print("\nAviablevkeys:-----")
	for _, book := range b.Books {
		if book.Available {
			fmt.Print(book.Print())
		}
	}
	fmt.Print("\nNon aviablevkeys:-----")
	for _, book := range b.Books {
		if !book.Available {
			fmt.Print(book.Print())
		}
	}
}

// ✅ Сохранить библиотеку в файл:
func (b *Books) Save() error {
	// открываем файл
	file, err := os.OpenFile("hw17/1/save", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// задаем параметры форматирования
	encoder := json.NewEncoder(file)
	encoder.SetIndent(" ", "	")

	// записываем
	err = encoder.Encode(b.Books)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	books := Books{
		[]Book{
			{
				ID:        1,
				Title:     "The Go Programming Language",
				Author:    "Alan Donovan",
				Pages:     380,
				Available: true,
				CreatedAt: time.Now(),
			},
			{
				ID:        2,
				Title:     "Clean Code",
				Author:    "Cormen",
				Pages:     450,
				Available: true,
				CreatedAt: time.Now(),
			},
			{
				ID:        3,
				Title:     "Design Patterns",
				Author:    "GoF",
				Pages:     395,
				Available: false,
				CreatedAt: time.Now(),
			},
			{
				ID:        4,
				Title:     "Refactoring",
				Author:    "Martin Fowler",
				Pages:     520,
				Available: true,
				CreatedAt: time.Now(),
			},
			{
				ID:        5,
				Title:     "Introduction to Algorithms",
				Author:    "Cormen",
				Pages:     1300,
				Available: true,
				CreatedAt: time.Now(),
			},
		},
	}
	books.ShowBooks()

	bigBook := books.Books[0]
	for _, v := range books.Books {
		if v.Pages > bigBook.Pages {
			bigBook = v
		}
	}
	fmt.Println("\nКнига с самым большим количеством страниц:", bigBook.Print())

	us := make(map[string]int)
	author := ""
	for _, v := range books.Books {
		us[v.Author]++
		if us[author] < us[v.Author] {
			author = v.Author
		}
	}
	fmt.Println("Автора с самым большим количеством книг:", author, us[author])
	books.Save()
}
