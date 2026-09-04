package storage

import "orders-processor/internal/models"

func GetOrderList(capacity int) models.OrderList {
	if capacity < 1 {
		orderList := models.OrderList{OrderItems: make([]models.Order, 0, 1), NextId: 1}
		return orderList
	}
	orderList := models.OrderList{OrderItems: make([]models.Order, 0, capacity), NextId: 1}
	return orderList
}

func FillOrderListWithTestData(orderList *models.OrderList) {
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
}
