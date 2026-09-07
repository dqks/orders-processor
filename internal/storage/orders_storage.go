package storage

import (
	"errors"
	"fmt"
	"orders-processor/internal/order"
)

func FillOrderListWithTestData(orderList *order.OrderList) error {
	product1, err1 := order.StoreProduct("Рубашка", 200.5, 10)
	product2, err2 := order.StoreProduct("Апельсин", 10, 10)
	product3, err3 := order.StoreProduct("Скрепка", 5, 10)

	if err1 != nil || err2 != nil || err3 != nil {
		return errors.New("Произошла ошибка при заполнении списка заказов")
	}

	productList := []order.Product{product1, product2, product3}

	clientNames := [5]string{"Иван Иванов", "Сергей Петров", "Алеша Попович", "Сергей Иванович", "Сергей Иванович"}

	for _, cl := range clientNames {
		err := orderList.Store(cl, productList, "new")
		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}
