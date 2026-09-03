package models

import (
	"errors"
	"fmt"
	"strings"
)

type OrderList struct {
	OrderItems []Order
	NextId     int
}

type Order struct {
	Id          int
	ClientName  string
	ProductList []Product
	Status      string
}

func StoreOrder(id int, clientName string, productList []Product, status string) (Order, error) {
	if len((strings.TrimSpace(clientName))) == 0 || len(productList) == 0 || len(strings.TrimSpace(status)) == 0 {
		return Order{}, errors.New("Невалидные данные заказа")
	}
	return Order{Id: id, ClientName: clientName, ProductList: productList, Status: status}, nil
}

func (orderList *OrderList) Store(clientName string, productList []Product, status string) error {
	order, err := StoreOrder(orderList.NextId, clientName, productList, status)
	orderList.NextId++
	if err != nil {
		return errors.New("Ошибка при создании заказа")
	}

	orderList.OrderItems = append(orderList.OrderItems, order)

	return nil
}

func (orderList OrderList) Show() {
	fmt.Println("Список всех заказов:")
	for _, o := range orderList.OrderItems {
		o.Show()
	}
}

func (o Order) Show() {
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

func (orderList *OrderList) GetOrderPriceById(id int) (float64, error) {
	if id < 0 {
		return 0, errors.New("Заказ по указанному ID не найден")
	}

	var foundOrder *Order

	for i := range orderList.OrderItems {
		if orderList.OrderItems[i].Id == id {
			foundOrder = &orderList.OrderItems[i]
		}
	}

	if foundOrder == nil {
		return 0, errors.New("Заказ по указанному ID не найден")
	}

	var sum float64
	for i := range foundOrder.ProductList {
		sum += foundOrder.ProductList[i].Price * float64(foundOrder.ProductList[i].Amount)
	}

	return sum, nil
}
