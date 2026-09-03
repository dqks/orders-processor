package main

import (
	"errors"
	"fmt"
	"orders-processor/internal/models"
)

func processOrders(n int, ordersChan <-chan *models.Order, resultChan chan<- error) {
	for order := range ordersChan {
		fmt.Printf("Worker %d начал обработку заказа #%d\n", n, order.Id)
		for _, p := range order.ProductList {
			err := models.ValidateProduct(p)
			if err != nil {
				resultChan <- errors.New("Невалидная информация о товаре")
				fmt.Printf("Worker %d не смог обработать заказ #%d\n", n, order.Id)
				order.Status = "error"
				return
			}
		}
		order.Status = "completed"
		fmt.Printf("Worker %d закончил обработку заказа #%d\n", n, order.Id)
		resultChan <- nil
	}
}

func main() {
	orderList := models.OrderList{OrderItems: make([]models.Order, 0, 10), NextId: 1}
	productList := []models.Product{
		{Name: "Рубашка", Price: 200.5, Amount: 10},
		{Name: "Апельсин", Price: 50, Amount: 1},
		{Name: "Скрепка", Price: 15, Amount: 3},
	}
	orderList.Store("Иван Иванов", productList, "Оплачено")
	orderList.Store("Сергей Петров", productList, "Ожидание оплаты")
	orderList.Store("Алеша Попович", productList, "Отменено")
	orderList.Store("Сергей Иванович", productList, "Оплачено")
	orderList.Store("Илья Иванов", append(productList, models.Product{Name: "Рубашка", Price: 200.5, Amount: -10}), "Отменено")

	workers := 3
	orders := 5

	orderChan := make(chan *models.Order, orders)
	resultChan := make(chan error, orders)

	for i := 1; i <= workers; i++ {
		go processOrders(i, orderChan, resultChan)
	}

	for i := 1; i <= orders; i++ {
		o, e := orderList.GetOrderById(i)
		if e == nil {
			orderChan <- o
		}
	}

	close(orderChan)

	for i := 1; i <= orders; i++ {
		fmt.Println(<-resultChan)
	}

}
