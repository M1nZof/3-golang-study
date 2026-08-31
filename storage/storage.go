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
	Read() (*bins.BinList, error)
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

	f, err := os.Create(js.FilePath)
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		fmt.Println("Ошибка записи данных в файл:", err)
		return err
	}
	return nil
}

func (js *JsonStorage) Read() (*bins.BinList, error) {
	if !file.IsJsonExtension(js.FilePath) {
		return nil, errors.New("файл не является JSON")
	}

	data, err := os.ReadFile(js.FilePath)
	if err != nil {
		return nil, err
	}

	var list bins.BinList
	if err := json.Unmarshal(data, &list); err != nil {
		fmt.Println("Ошибка десериализации структуры:", err)
		return nil, err
	}
	return &list, nil
}

func NewJsonStorage() Storage {
	return &JsonStorage{
		FilePath: "my_awesome_bin.json",
	}
}
