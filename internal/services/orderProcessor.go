package services

import (
	"fmt"
	"orders-processor/internal/models"
)

func ProcessOrders(ordersChan <-chan *models.Order, resultChan chan<- error, n int) {
	var err error = nil
	for order := range ordersChan {
		fmt.Printf("Worker %d начал обработку заказа #%d\n", n, order.ID)
		// Валидация полей заказа
		err = order.ValidateOrder()
		if err != nil {
			resultChan <- err
			fmt.Printf("Worker %d не смог обработать заказ #%d\n", n, order.ID)
			order.UpdateStatus("error")
		}
		// Валидация полей продуктов заказа
		if err == nil {
			for _, p := range order.ProductList {
				err = models.ValidateProduct(p)
				if err != nil {
					resultChan <- err
					fmt.Printf("Worker %d не смог обработать заказ #%d\n", n, order.ID)
					order.UpdateStatus("error")
				}
			}
		}
		// Если все прошло валидацию
		if err == nil {
			order.UpdateStatus("completed")
			fmt.Printf("Worker %d закончил обработку заказа #%d, его итоговая сумма %.2f\n", n, order.ID, order.GetTotalPrice())
			resultChan <- nil
		}
	}
}
