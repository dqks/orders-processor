package models

import (
	"errors"
	"fmt"
	"orders-processor/internal/appErrors"
)

type OrderList struct {
	orderItems []Order
	nextId     int
}

func (orderList *OrderList) GetOrderByID(id int) (*Order, error) {
	if id < 0 {
		return nil, errors.New("Невалидный id заказа")
	}
	for i := range orderList.orderItems {
		if orderList.orderItems[i].id == id {
			return &orderList.orderItems[i], nil
		}
	}
	return nil, fmt.Errorf("Заказ по заданному id не найден: %w", appErrors.ErrOrderNotFound)
}

func (orderList *OrderList) Store(clientName string, productList []Product, status string) error {
	order, err := StoreOrder(orderList.nextId, clientName, productList, status)
	if err != nil {
		// Не печатаем ошибку - это не ответственность Store
		return err
	}
	orderList.nextId++
	orderList.orderItems = append(orderList.orderItems, order)
	return nil
}

func (orderList OrderList) Print() {
	fmt.Println("Список всех заказов:")
	for _, o := range orderList.orderItems {
		o.Print()
	}
}

func (orderList *OrderList) GetOrderPriceByID(id int) (float64, error) {
	foundOrder, err := orderList.GetOrderByID(id)

	if errors.Is(err, appErrors.ErrOrderNotFound) || err != nil {
		return 0, err
	}

	var sum float64
	for i := range foundOrder.productList {
		sum += foundOrder.productList[i].Price * float64(foundOrder.productList[i].Amount)
	}

	return sum, nil
}

func (orderList OrderList) GetOrderAmount() int {
	return len(orderList.orderItems)
}

func (orderList OrderList) GetOrdersByIDs(ids []int) []Order {
	orders := make([]Order, 0, len(ids))

	for _, id := range ids {
		order, err := orderList.GetOrderByID(id)
		if !errors.Is(err, appErrors.ErrOrderNotFound) && err == nil {
			orders = append(orders, *order)
		}
	}
	return orders
}

func (o Order) GetTotalPrice() float64 {
	var sum float64
	for _, p := range o.productList {
		sum += p.GetTotalPrice()
	}
	return sum
}

func CreateOrderList(capacity int) OrderList {
	if capacity < 1 {
		orderList := OrderList{orderItems: make([]Order, 0, 1), nextId: 1}
		return orderList
	}
	orderList := OrderList{orderItems: make([]Order, 0, capacity), nextId: 1}
	return orderList
}
