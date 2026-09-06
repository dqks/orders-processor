package storage

import (
	"fmt"
	"orders-processor/internal/models"
)

func FillOrderListWithTestData(orderList *models.OrderList) {
	productList := []models.Product{
		{Name: "Рубашка", Price: 200.5, Amount: 10},
		{Name: "Апельсин", Price: 50, Amount: 1},
		{Name: "Скрепка", Price: 15, Amount: 3},
	}

	clientNames := [5]string{"Иван Иванов", "Сергей Петров", "Алеша Попович", "Сергей Иванович", "Сергей Иванович"}

	for _, cl := range clientNames {
		err := orderList.Store(cl, productList, "new")
		if err != nil {
			fmt.Println(err)
		}
	}
}
