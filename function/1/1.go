package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Status     string    `json:"status"`
	Priority   int       `json:"priority"`
	AssignedTo string    `json:"assigned_to"`
	CreatedAt  time.Time `json:"created_at"`
}

func (t Task) Print() {
	fmt.Printf("id: %d\ntitle: %s\nstatus: %s\npriority: %d\nassigned to: %s\ncreater at: %s\n",
		t.ID,
		t.Title,
		t.Status,
		t.Priority,
		t.AssignedTo,
		t.CreatedAt.Format(time.DateTime),
	)
}

var (
	dataFile = "function/1/file1"
)

func Download(fileName string) (map[int]*Task, error) {
	// открытие файла
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return map[int]*Task{}, err
	}
	defer file.Close()

	// выгрузка данных
	data := make(map[int]*Task)
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return map[int]*Task{}, err

	}

	return data, nil
}

func Save(data map[int]*Task, fileName string) error {
	// открываем файл
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// задаем параметры форматирования
	encoder := json.NewEncoder(file)
	encoder.SetIndent(" ", "	")

	// записываем
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}

func AddTask(t Task, data map[int]*Task) {
	id := len(data)
	t.ID = id
	t.CreatedAt = time.Now()
	data[id] = &t
}

func DataPrint(data map[int]*Task) {
	for _, v := range data {
		fmt.Println("------------------")
		v.Print()
	}
}

func Stats(data map[int]*Task)  {
    done := make([]Task, 0, len(data))
    for 
}

func main() {
	data, err := Download(dataFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	tasks := []Task{
		{
			Title:      "Настроить CI/CD",
			Status:     "in_progress",
			Priority:   1,
			AssignedTo: "dev_user",
		},
		{
			Title:      "Исправить баг в API",
			Status:     "new",
			Priority:   2,
			AssignedTo: "alice_w",
		},
		{
			Title:      "Обновить документацию",
			Status:     "done",
			Priority:   3,
			AssignedTo: "admin",
		},
		{
			Title:      "Провести код-ревью",
			Status:     "new",
			Priority:   1,
			AssignedTo: "admin",
		},
	}
	for _, v := range tasks {
		AddTask(v, data)
	}
	DataPrint(data)

    fmt.Printf("all done: %v\nthe highest priority: %s\nemployee with the largest number of tasks: %s\n", Stats(data))

	Save(data, dataFile)
}
