package models

import (
	"errors"
	"fmt"
	"strings"
)

type Order struct {
	Id          int
	ClientName  string
	ProductList []Product
	Status      string
}

func (o Order) Print() {
	fmt.Printf("ID: %d, клиент: %s, статус: %s\n", o.Id, o.ClientName, o.Status)
	fmt.Println("Товары:")
	for _, p := range o.ProductList {
		PrintProduct(p)
	}
	fmt.Println("")
}

func (o *Order) UpdateStatus(newStatus string) error {
	if len(strings.TrimSpace(newStatus)) == 0 {
		return errors.New("Получен невалидный статус")
	}
	o.Status = newStatus
	return nil
}
