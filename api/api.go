package api

import "golang-study/bins"

type Storage interface {
	Write(bins.BinList) error
	Read() (*bins.BinList, error)
}

type BinService struct {
	storage Storage
}

func NewBinService(storage Storage) *BinService {
	return &BinService{storage: storage}
}

func (s *BinService) SaveBins(list bins.BinList) error {
	return s.storage.Write(list)
}

func (s *BinService) LoadBins() (*bins.BinList, error) {
	return s.storage.Read()
}
