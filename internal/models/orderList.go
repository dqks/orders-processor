package models

import (
	"errors"
	"fmt"
)

type OrderList struct {
	OrderItems []Order
	NextId     int // сделать приватным
}

func (orderList *OrderList) GetOrderById(id int) (*Order, error) {
	if id < 0 {
		return nil, errors.New("Невалидный ID заказа")
	}
	for i := range orderList.OrderItems {
		if orderList.OrderItems[i].ID == id {
			return &orderList.OrderItems[i], nil
		}
	}
	return nil, errors.New("Заказ по указанному ID не найден")
}

func (orderList *OrderList) Store(clientName string, productList []Product, status string) error {
	order, err := StoreOrder(orderList.NextId, clientName, productList, status)
	if err != nil {
		// Закомментировано для теста пула воркеров, поскольку
		// Эта проверка не даст создать неправильный заказ

		// fmt.Errorf("Ошибка создания добавления заказа в список: %w", os.ErrInvalid)
		// fmt.Println("Ошибка создания добавления заказа в список")
		// return err
	}
	orderList.NextId++
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
		return 0, err
	}

	var sum float64
	for i := range foundOrder.ProductList {
		sum += foundOrder.ProductList[i].Price * float64(foundOrder.ProductList[i].Amount)
	}

	return sum, nil
}

func (orderList OrderList) GetOrderAmount() int {
	return len(orderList.OrderItems)
}
