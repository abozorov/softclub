package storage

import (
	"encoding/json"
	"hw19/jsonapp/models"
	"os"
)

var (
	fileName = "storage/users.json"
)

func SaveUsers(users []models.User) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	jDec := json.NewEncoder(file)
	jDec.SetIndent("", "	") // Для красивой записи в JSON файл
	err = jDec.Encode(&users)
	if err != nil {
		return err
	}
	return nil
}

func LoadUsers() ([]models.User, error) {
	users := make([]models.User, 0)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return []models.User{}, err
	}
	defer file.Close()
	jDec := json.NewDecoder(file)
	err = jDec.Decode(&users)
	if err != nil {
		return []models.User{}, err
	}
	return users, nil
}
