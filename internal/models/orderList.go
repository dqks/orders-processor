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

func StoreOrder(id int, clientName string, productList []Product, status string) (Order, error) {
	if len(strings.TrimSpace(clientName)) == 0 || len(productList) == 0 || len(strings.TrimSpace(status)) == 0 {
		return Order{}, errors.New("Невалидные данные заказа")
	}
	return Order{Id: id, ClientName: clientName, ProductList: productList, Status: status}, nil
}

func (orderList *OrderList) GetOrderById(id int) (*Order, error) {
	if id < 0 {
		return nil, errors.New("Невалидный ID заказа")
	}
	for i := range orderList.OrderItems {
		if orderList.OrderItems[i].Id == id {
			return &orderList.OrderItems[i], nil
		}
	}
	return nil, errors.New("Заказ не найден")
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

func (orderList OrderList) Print() {
	fmt.Println("Список всех заказов:")
	for _, o := range orderList.OrderItems {
		o.Print()
	}
}

func (orderList *OrderList) GetOrderPriceById(id int) (float64, error) {
	if id < 0 {
		return 0, errors.New("Заказ по указанному ID не найден")
	}

	foundOrder, err := orderList.GetOrderById(id)

	if err != nil {
		return 0, errors.New("Заказ по указанному ID не найден")
	}

	var sum float64
	for i := range foundOrder.ProductList {
		sum += foundOrder.ProductList[i].Price * float64(foundOrder.ProductList[i].Amount)
	}

	return sum, nil
}
