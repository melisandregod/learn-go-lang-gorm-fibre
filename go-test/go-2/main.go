package main

import "fmt"

type Stock struct {
	Symbol string
	Price string
}

func (s Stock) FullName() string {
	return s.Symbol + ":" + s.Price;
}

func main() {
	nvdia := Stock{
		Symbol: "NVDA",
		Price: "100 USD",
	}

	fmt.Println(nvdia)
}