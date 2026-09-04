package main

import (
	"fmt"
	"orders-processor/internal/models"
)

func processOrders(n int, ordersChan <-chan *models.Order, resultChan chan<- error) {
	var err error = nil
	for order := range ordersChan {
		fmt.Printf("Worker %d начал обработку заказа #%d\n", n, order.ID)
		for _, p := range order.ProductList {
			err = models.ValidateProduct(p)
			if err != nil {
				resultChan <- err
				fmt.Printf("Worker %d не смог обработать заказ #%d\n", n, order.ID)
				order.UpdateStatus("error")
			}
		}
		if err == nil {
			order.UpdateStatus("completed")
			fmt.Printf("Worker %d закончил обработку заказа #%d\n", n, order.ID)
			resultChan <- nil
		}
	}
}

func main() {
	orderList := models.OrderList{OrderItems: make([]models.Order, 0, 10), NextId: 1}

	productList := []models.Product{
		{Name: "Рубашка", Price: 200.5, Amount: 10},
		{Name: "Апельсин", Price: 50, Amount: 1},
		{Name: "Скрепка", Price: 15, Amount: 3},
	}

	orderList.Store("Илья Иванов", append(productList, models.Product{Name: "Рубашка", Price: 200.5, Amount: -10}), "new")
	orderList.Store("Иван Иванов", productList, "new")
	orderList.Store("Сергей Петров", productList, "new")
	orderList.Store("Илья Иванов", append(productList, models.Product{Name: "Рубашка", Price: 200.5, Amount: -10}), "new")
	orderList.Store("Илья Иванов", append(productList, models.Product{Name: "Рубашка", Price: 200.5, Amount: -10}), "new")
	orderList.Store("Илья Иванов", append(productList, models.Product{Name: "Рубашка", Price: 200.5, Amount: -10}), "new")
	orderList.Store("Алеша Попович", productList, "new")
	orderList.Store("Сергей Иванович", productList, "new")
	orderList.Store("Сергей Иванович", productList, "new")

	ordersAmount := orderList.GetOrderAmount()

	orders := make([]models.Order, 0, ordersAmount)

	for i := 1; i <= ordersAmount; i++ {
		order, err := orderList.GetOrderById(i)
		if err == nil {
			orders = append(orders, *order)
		}
	}
	actOrdersAmount := len(orders)

	workers := 3

	orderChan := make(chan *models.Order, actOrdersAmount)

	resultChan := make(chan error, actOrdersAmount)

	for i := 1; i <= workers; i++ {
		go processOrders(i, orderChan, resultChan)
	}

	for i := 0; i < actOrdersAmount; i++ {
		orderChan <- &orders[i]
	}

	close(orderChan)

	for i := 1; i <= actOrdersAmount; i++ {
		fmt.Println(<-resultChan, " Результат")
	}

}
