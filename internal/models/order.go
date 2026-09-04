package models

import (
	"errors"
	"fmt"
	"strings"
)

type Order struct {
	ID          int
	ClientName  string
	ProductList []Product
	status      string
}

func (o Order) Print() {
	fmt.Printf("ID: %d, клиент: %s, статус: %s\n", o.ID, o.ClientName, o.status)
	fmt.Println("Товары:")
	for _, p := range o.ProductList {
		PrintProduct(p)
	}
	fmt.Println("")
}

func (o *Order) UpdateStatus(newStatus string) error {
	if newStatus != "new" && newStatus != "completed" && newStatus != "error" {
		return errors.New("Получен невалидный статус")
	}
	o.status = newStatus
	return nil
}

func (o *Order) GetStatus() string {
	return o.status
}

func StoreOrder(id int, clientName string, productList []Product, status string) (Order, error) {
	if id < 0 {
		return Order{}, errors.New("Невалидный ID для создания заказа")
	}

	if len(strings.TrimSpace(clientName)) == 0 || len(productList) == 0 || len(strings.TrimSpace(status)) == 0 {
		return Order{}, errors.New("Невалидные данные для создания заказа")
	}

	if status != "new" && status != "completed" && status != "error" {
		return Order{}, errors.New("Невалидный статус для создания заказа")
	}

	for _, p := range productList {
		err := ValidateProduct(p)
		if err != nil {
			return Order{}, err
		}
	}

	return Order{ID: id, ClientName: clientName, ProductList: productList, status: status}, nil
}
