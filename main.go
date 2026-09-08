package main

import (
	"orders-processor/internal/order"
	"orders-processor/internal/services"
	"orders-processor/internal/storage"
	"orders-processor/internal/worker"
)

func main() {
	orderList := order.CreateOrderList(10)
	if err := storage.FillOrderListWithTestData(&orderList); err != nil {
		panic(err)
	}
	orders := orderList.GetOrdersByIDs([]int{1, 2, 3, 4, 5})
	worker.ProcessByWorkers(
		3,
		orders,
		services.ProcessOrders,
		services.ProcessOrderResult())
}
