package main

type Order struct {
	Price int64
	Qty   int64
}

type PriceReader interface {
	Read() int64
}

func (o *Order) Read() int64 {
	return o.Price
}

func main() {
	c := Order{Price: 500, Qty: 60}
	readers := []PriceReader{&c}
	_ = readers[0].Read()
	var pr PriceReader = &c
	_ = pr.Read()
}
