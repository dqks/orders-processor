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

//func StoreProduct(name string, price float64, amount int) (Product, error) {
//	err := ValidateProductFields(name, price, amount)
//	if err != nil {
//		return Product{}, err
//	}
//	return Product{Name: name, Price: price, Amount: amount}, nil
//}

func ValidateProductFields(name string, price float64, amount int) error {
	if price <= 0 || amount <= 0 || len(strings.TrimSpace(name)) == 0 {
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
