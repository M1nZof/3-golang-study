package golangstudy

import "time"

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func newBin(id, name string, private bool, createdAt time.Time) *Bin {
	res := Bin{
		id:        id,
		private:   private,
		createdAt: createdAt,
		name:      name,
	}
	return &res
}

type BinList struct {
	bins []Bin
}

func newBinList(bins *[]Bin) *BinList {
	res := BinList{
		bins: *bins,
	}
	return &res
}

func main() {

}
