package order

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type Order struct {
	id          int
	clientName  string
	productList []Product
	status      string
}

func (o *Order) ID() int {
	return o.id
}

func (o *Order) ClientName() string {
	return o.clientName
}

func (o *Order) ProductList() []Product {
	return o.productList
}

func (o Order) Print(w io.Writer) {
	fmt.Fprintf(w, "id: %d, клиент: %s, статус: %s\n", o.id, o.clientName, o.status)
	fmt.Fprintln(w, "Товары:")
	for _, p := range o.productList {
		PrintProduct(w, p)
	}
	fmt.Fprintln(w, "")
}

func (o *Order) SetStatus(newStatus string) error {
	if newStatus != "new" && newStatus != "completed" && newStatus != "error" {
		return errors.New("Присваивается невалидный статус заказа")
	}
	o.status = newStatus
	return nil
}

func (o *Order) Status() string {
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
		err := ValidateProductFields(p.name, p.price, p.amount)
		if err != nil {
			return err
		}
	}

	return nil
}

func (o Order) GetTotalPrice() float64 {
	var sum float64
	for _, p := range o.productList {
		sum += p.GetTotalPrice()
	}
	return sum
}
