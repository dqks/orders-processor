package order

import (
	"errors"
	"fmt"
	"io"
	"orders-processor/internal/apperrors"
)

type OrderList struct {
	orderItems []Order
	nextId     int
}

// Возвращаю не указатели  а просто слайс и можно будет менять из вне
func (ol *OrderList) Orders() []Order {
	return ol.orderItems
}

func (ol *OrderList) GetOrderByID(id int) (*Order, error) {
	if id < 0 {
		return nil, errors.New("Невалидный id заказа")
	}
	for i := range ol.orderItems {
		if ol.orderItems[i].id == id {
			return &ol.orderItems[i], nil
		}
	}
	return nil, fmt.Errorf("Заказ по заданному id не найден: %w", apperrors.ErrOrderNotFound)
}

func (ol *OrderList) Store(clientName string, productList []Product, status string) error {
	order, err := StoreOrder(ol.nextId, clientName, productList, status)
	if err != nil {
		// Не печатаем ошибку - это не ответственность Store
		return err
	}
	ol.nextId++
	ol.orderItems = append(ol.orderItems, order)
	return nil
}

func (ol OrderList) Print(w io.Writer) {
	fmt.Fprintln(w, "Список всех заказов:")
	for _, o := range ol.orderItems {
		o.Print(w)
	}
}

func (ol *OrderList) GetOrderPriceByID(id int) (float64, error) {
	foundOrder, err := ol.GetOrderByID(id)

	if errors.Is(err, apperrors.ErrOrderNotFound) || err != nil {
		return 0, err
	}

	var sum float64
	for i := range foundOrder.productList {
		sum += foundOrder.productList[i].price * float64(foundOrder.productList[i].amount)
	}

	return sum, nil
}

func (ol OrderList) GetOrderAmount() int {
	return len(ol.orderItems)
}

func (ol OrderList) GetOrdersByIDs(ids []int) []*Order {
	orders := make([]*Order, 0, len(ids))

	for id := range ids {
		order, err := ol.GetOrderByID(id)
		if !errors.Is(err, apperrors.ErrOrderNotFound) && err == nil {
			orders = append(orders, order)
		}
	}
	return orders
}

func CreateOrderList(capacity int) OrderList {
	if capacity < 1 {
		orderList := OrderList{orderItems: make([]Order, 0, 1), nextId: 1}
		return orderList
	}
	orderList := OrderList{orderItems: make([]Order, 0, capacity), nextId: 1}
	return orderList
}
