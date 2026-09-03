package models

import "fmt"

type Product struct {
	Name   string
	Price  float64
	Amount int
}

func PrintProduct(p Product) {
	fmt.Printf("Название: %s, цена: %.2f, количество: %d\n", p.Name, p.Price, p.Amount)
}
