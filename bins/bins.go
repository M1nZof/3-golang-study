package bins

import "time"

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

func newBin(id, name string, private bool, createdAt time.Time) *Bin {
	res := Bin{
		Id:        id,
		Private:   private,
		CreatedAt: createdAt,
		Name:      name,
	}
	return &res
}

type BinList struct {
	bins []Bin
}

func newBinList(bins []Bin) *BinList {
	res := BinList{
		bins: bins,
	}
	return &res
}
