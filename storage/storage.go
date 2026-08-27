package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang-study/bins"
	"golang-study/file"
	"os"
)

const FilePath = "my_awesome_bin.json"

func WriteBinFile(content bins.BinList) error {
	data, err := json.Marshal(content)
	if err != nil {
		fmt.Println("Ошибка сериализации структуры:", err)
		return err
	}
	file, err := os.Create(FilePath)
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

func ReadBinFile() (string, error) {
	if !file.IsJsonExtension(FilePath) {
		return "", errors.New("Файл не является JSON")
	}
	data, err := os.ReadFile(FilePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
