package services

import (
	"fmt"
	"orders-processor/internal/models"
)

func ProcessOrdersByWorkers(workerNum int, orders []models.Order) {
	orderNum := len(orders)
	orderChan := make(chan *models.Order, orderNum)

	resultChan := make(chan error, orderNum)

	for i := 1; i <= workerNum; i++ {
		go ProcessOrders(orderChan, resultChan, i)
	}

	for i := 0; i < orderNum; i++ {
		orderChan <- &orders[i]
	}

	close(orderChan)

	for i := 1; i <= orderNum; i++ {
		fmt.Println(<-resultChan, " Результат")
	}
}

func ProcessOrders(ordersChan <-chan *models.Order, resultChan chan<- error, n int) {
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
