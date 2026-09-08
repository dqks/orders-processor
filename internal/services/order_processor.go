package services

import (
	"fmt"
	"orders-processor/internal/order"
	"sync"
)

type ProcessResult struct {
	Completed bool
	TotalSum  float64
}

func ProcessOrders(ordersChan <-chan *order.Order, resultChan chan<- ProcessResult, n int, wg *sync.WaitGroup) {
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
		err = order.ValidateOrderFields(
			id,
			o.ClientName(),
			productList,
			o.Status(),
		)
		if err != nil {
			if errStatus := o.SetStatus("error"); err != nil {
				resultChan <- ProcessResult{Completed: false, TotalSum: 0}
				fmt.Printf("Worker %d не смог обработать заказ #%d - %s\n",
					n,
					id,
					errStatus,
				)
				continue
			}
			resultChan <- ProcessResult{Completed: false, TotalSum: 0}
			fmt.Printf("Worker %d не смог обработать заказ #%d - %s\n",
				n,
				id,
				err,
			)
			continue
		}
		// Если все прошло валидацию
		if err == nil {
			if errStatus := o.SetStatus("completed"); errStatus != nil {
				resultChan <- ProcessResult{Completed: false, TotalSum: 0}
				fmt.Printf("Worker %d не смог обработать заказ #%d - %s\n",
					n,
					id,
					errStatus,
				)
				continue
			}
			resultChan <- ProcessResult{Completed: true, TotalSum: o.GetTotalPrice()}
			fmt.Printf("Worker %d закончил обработку заказа #%d, его итоговая сумма %.2f\n",
				n,
				id,
				o.GetTotalPrice(),
			)
			continue
		}
	}
}

func ProcessOrderResult() func(res ProcessResult, opened bool) {
	stats := make(map[string]float64)
	return func(res ProcessResult, opened bool) {
		if res.Completed {
			stats["completed"] = stats["completed"] + 1
			stats["totalSum"] = stats["totalSum"] + res.TotalSum

		} else if !res.Completed && opened {
			stats["failed"] = stats["failed"] + 1
		}
		if !opened {
			fmt.Println(stats)
		}
	}
}
