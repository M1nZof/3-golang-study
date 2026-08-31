package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang-study/bins"
	"golang-study/file"
	"os"
)

type Storage interface {
	Write(bins.BinList) error
	Read() (string, error)
}

type JsonStorage struct {
	FilePath string
}

func (js *JsonStorage) Write(content bins.BinList) error {
	data, err := json.Marshal(content)
	if err != nil {
		fmt.Println("Ошибка сериализации структуры:", err)
		return err
	}
	file, err := os.Create(js.FilePath)
	if err != nil {
		fmt.Println("Ошибка создания файла", err)
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Ошибка записи данных в файл", err)
		return err
	}
	return nil
}

func (js *JsonStorage) Read() (string, error) {
	if !file.IsJsonExtension(js.FilePath) {
		return "", errors.New("Файл не является JSON")
	}
	data, err := os.ReadFile(js.FilePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func NewJsonStorage() Storage {
	return &JsonStorage{
		FilePath: "my_awesome_bin.json",
	}
}
