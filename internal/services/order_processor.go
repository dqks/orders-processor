package services

import (
	"fmt"
	"orders-processor/internal/order"
	"sync"
)

func ProcessOrders(ordersChan <-chan *order.Order, resultChan chan<- error, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	// orderChan ожидает либо новые данные, либо закрытие канала
	// Канал должен закрываться там, где в него отдают данные
	// В данном случае ProcessByWorkers
	for o := range ordersChan {
		var err error = nil
		id := o.ID()
		productList := o.ProductList()
		fmt.Printf("Worker %d начал обработку заказа #%d\n", n, id)
		// Валидация полей заказа
		err = order.ValidateOrderFields(id, o.ClientName(), productList, o.Status())
		if err != nil {
			errStatus := o.SetStatus("error")
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
			errStatus := o.SetStatus("completed")
			if errStatus != nil {
				resultChan <- errStatus
				fmt.Printf("Worker %d не смог обработать заказ #%d\n", n, id)
			} else {
				resultChan <- nil
				fmt.Printf("Worker %d закончил обработку заказа #%d, его итоговая сумма %.2f\n", n, id, o.GetTotalPrice())
			}
		}
	}
}

func ProcessOrderResult(err error) {
	fmt.Println(err)
}
