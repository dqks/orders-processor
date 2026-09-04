package main

import (
	"orders-processor/internal/models"
	"orders-processor/internal/services"
	"orders-processor/internal/storage"
)

func main() {
	orderList := storage.GetOrderList(10)
	storage.FillOrderListWithTestData(&orderList)
	orders := orderList.GetOrdersByIds([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	//services.ProcessOrdersByWorkers(3, orders)
	services.ProcessByWorkers[models.Order, error](3, orders, services.ProcessOrders)
}
