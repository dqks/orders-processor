package models

import (
	"errors"
	"fmt"
	"strings"
)

type Order struct {
	id          int
	clientName  string
	productList []Product
	status      string
}

func (o *Order) GetID() int {
	return o.id
}

func (o *Order) GetClientName() string {
	return o.clientName
}

func (o *Order) GetProductList() []Product {
	return o.productList
}

func (o Order) Print() {
	fmt.Printf("id: %d, клиент: %s, статус: %s\n", o.id, o.clientName, o.status)
	fmt.Println("Товары:")
	for _, p := range o.productList {
		PrintProduct(p)
	}
	fmt.Println("")
}

func (o *Order) SetStatus(newStatus string) error {
	if newStatus != "new" && newStatus != "completed" && newStatus != "error" {
		return errors.New("Присваивается невалидный статус заказа")
	}
	o.status = newStatus
	return nil
}

func (o *Order) GetStatus() string {
	return o.status
}

func StoreOrder(id int, clientName string, productList []Product, status string) (Order, error) {
	err := ValidateOrderFields(id, clientName, productList, status)

	if err != nil {
		return Order{}, err
	}

	return Order{id: id, clientName: clientName, productList: productList, status: status}, nil
}

func ValidateOrderFields(id int, clientName string, productList []Product, status string) error {
	if id < 0 {
		return errors.New("Передан невалидный id заказа")
	}

	if len(strings.TrimSpace(clientName)) == 0 || len(productList) == 0 {
		return errors.New("Передан невалидные данные заказа")
	}

	if status != "new" && status != "completed" && status != "error" {
		return errors.New("Передан невалидный статус заказа")
	}

	for _, p := range productList {
		err := ValidateProduct(p)
		if err != nil {
			return err
		}
	}

	return nil
}
