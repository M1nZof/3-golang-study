package bins

import "time"

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
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
	Bins []Bin `json:"bins"`
}

func newBinList(bins []Bin) *BinList {
	res := BinList{
		Bins: bins,
	}
	return &res
}
