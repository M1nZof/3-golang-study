package main

import (
	"fmt"
	"time"

	"golang-study/api"
	"golang-study/bins"
	"golang-study/storage"
)

func main() {
	st := storage.NewJsonStorage()

	service := api.NewBinService(st)

	bin := bins.NewBin("1", "my-bin", false, time.Now())
	list := bins.NewBinList([]bins.Bin{*bin})

	if err := service.SaveBins(*list); err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}

	loaded, err := service.LoadBins()
	if err != nil {
		fmt.Println("Ошибка загрузки:", err)
		return
	}

	fmt.Printf("Загружено бинов: %d\n", len(loaded.Bins))
}
