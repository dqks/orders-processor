package models

import (
	"errors"
	"fmt"
	"strings"
)

type Product struct {
	Name   string
	Price  float64
	Amount int
}

func ValidateProduct(p Product) error {
	if p.Price <= 0 || p.Amount <= 0 || len(strings.TrimSpace(p.Name)) == 0 {
		return errors.New("Невалидная информация о товаре")
	}
	return nil
}

func PrintProduct(p Product) {
	fmt.Printf("Название: %s, цена: %.2f, количество: %d\n", p.Name, p.Price, p.Amount)
}

func (p Product) GetTotalPrice() float64 {
	return p.Price * float64(p.Amount)
}
