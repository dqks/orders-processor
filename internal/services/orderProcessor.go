package services

import (
	"fmt"
	"orders-processor/internal/models"
)

func ProcessOrders(ordersChan <-chan *models.Order, resultChan chan<- error, n int) {
	for order := range ordersChan {
		var err error = nil
		id := order.GetID()
		productList := order.GetProductList()
		fmt.Printf("Worker %d начал обработку заказа #%d\n", n, id)
		// Валидация полей заказа
		err = models.ValidateOrderFields(id, order.GetClientName(), productList, order.GetStatus())
		if err != nil {
			errStatus := order.SetStatus("error")
			if errStatus != nil {
				resultChan <- errStatus
				fmt.Printf("Worker %d не смог обработать заказ #%d\n", n, id)
			} else {
				resultChan <- err
				fmt.Printf("Worker %d не смог обработать заказ #%d\n", n, id)
			}

		}
		// Если все прошло валидацию
		if err == nil {
			errStatus := order.SetStatus("completed")
			if errStatus != nil {
				resultChan <- errStatus
				fmt.Printf("Worker %d не смог обработать заказ #%d\n", n, id)
			} else {
				resultChan <- nil
				fmt.Printf("Worker %d закончил обработку заказа #%d, его итоговая сумма %.2f\n", n, id, order.GetTotalPrice())
			}

		}
	}
}

func ProcessOrderResult(err error) {
	fmt.Println(err)
}
