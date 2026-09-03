package main

import (
	"fmt"
	"orders-processor/internal/models"
)

func main() {
	orderList := models.OrderList{OrderItems: make([]models.Order, 0, 10), NextId: 1}
	productList := []models.Product{
		{Name: "Рубашка", Price: 200.5, Amount: 10},
		{Name: "Апельсин", Price: 50, Amount: 1},
		{Name: "Скрепка", Price: 15, Amount: 3},
	}
	orderList.Store("Иван Иванов", productList, "Оплачено")

	fmt.Println(orderList.GetOrderPriceById(1))
	orderList.OrderItems[0].Show()

}
