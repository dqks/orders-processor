package order

import (
	"bytes"
	"errors"
	"fmt"
	"slices"
	"testing"
)

func fillOrderListWithTestData(orderList *OrderList) error {
	product1, err1 := StoreProduct("Рубашка", 200.5, 10)
	product2, err2 := StoreProduct("Апельсин", 10, 10)
	product3, err3 := StoreProduct("Скрепка", 5, 10)

	if err1 != nil || err2 != nil || err3 != nil {
		return errors.New("Произошла ошибка при заполнении списка заказов")
	}

	productList := []Product{product1, product2, product3}

	clientNames := [5]string{"Иван Иванов", "Сергей Петров", "Алеша Попович", "Сергей Иванович", "Сергей Иванович"}

	for _, cl := range clientNames {
		err := orderList.Store(cl, productList, "new")
		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}

func TestGetOrderByID(t *testing.T) {
	orderList := CreateOrderList(5)
	err := fillOrderListWithTestData(&orderList)
	if err != nil {
		t.Errorf("GetOrderByID() = %s", err)
	}

	tests := []struct {
		testName string
		id       int
		wantErr  bool
	}{
		{
			testName: "Id найден",
			id:       1,
			wantErr:  false,
		},
		{
			testName: "Не найден id",
			id:       999,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			_, err := orderList.GetOrderByID(tt.id)

			if err != nil && !tt.wantErr {
				t.Errorf("GetOrderByID() = %s", err)
			}
		})
	}
}

func TestStore(t *testing.T) {
	orderList := CreateOrderList(5)
	err := fillOrderListWithTestData(&orderList)
	if err != nil {
		t.Errorf("Store() = %s", err)
	}
	var productList []Product = []Product{
		{name: "Рубашка", price: 100, amount: 20},
		{name: "Апельсин", price: 200, amount: 20},
		{name: "Яблоко", price: 300, amount: 50},
	}

	tests := []struct {
		testName    string
		clientName  string
		productList []Product
		status      string
		wantErr     bool
	}{
		{
			testName:    "Валидные данные",
			clientName:  "Smith",
			productList: productList,
			status:      "new",
			wantErr:     false,
		},
		{
			testName:    "Пустое имя клиента",
			clientName:  "",
			productList: productList,
			status:      "new",
			wantErr:     true,
		},
		{
			testName:    "Пустой статус",
			clientName:  "Smith",
			productList: productList,
			status:      "",
			wantErr:     true,
		},
		{
			testName:    "Невалидный статус",
			clientName:  "Smith",
			productList: productList,
			status:      "dadas",
			wantErr:     true,
		},
		{
			testName:    "Невалидный продукт",
			clientName:  "Smith",
			productList: append(productList, Product{name: "", price: 10, amount: 20}),
			status:      "dadas",
			wantErr:     true,
		},
		{
			testName:    "Невалидный продукт",
			clientName:  "Smith",
			productList: append(productList, Product{name: "Рубашка", price: 0, amount: 20}),
			status:      "dadas",
			wantErr:     true,
		},
		{
			testName:    "Невалидный продукт",
			clientName:  "Smith",
			productList: append(productList, Product{name: "Рубашка", price: -1, amount: 20}),
			status:      "dadas",
			wantErr:     true,
		},
		{
			testName:    "Невалидный продукт",
			clientName:  "Smith",
			productList: append(productList, Product{name: "Рубашка", price: 1, amount: 0}),
			status:      "dadas",
			wantErr:     true,
		},
		{
			testName:    "Невалидный продукт",
			clientName:  "Smith",
			productList: append(productList, Product{name: "Рубашка", price: 1, amount: -1}),
			status:      "dadas",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			err := orderList.Store(tt.clientName, tt.productList, tt.status)
			if (err == nil && tt.wantErr) || (err != nil && !tt.wantErr) {
				t.Errorf("Store(%s, %#v, %s) = %s", tt.clientName, tt.productList, tt.status, err)
			}
		})
	}

}

func TestOrderListPrint(t *testing.T) {
	orderList := CreateOrderList(1)
	product, err := StoreProduct("Рубашка", 100, 20)

	if err != nil {
		t.Errorf("Print() при %#v", product)
	}

	productList := []Product{product}
	orderList.Store("Smith", productList, "new")

	var buf bytes.Buffer
	expected := "Список всех заказов:\nid: 1, клиент: Smith, статус: new\n" +
		"Товары:\n" +
		"Название: Рубашка, цена: 100.00, количество: 20\n" +
		"\n"
	orderList.Print(&buf)
	if buf.String() != expected {
		t.Errorf("Print(&buf), при %d, %#v, %s", orderList.Orders()[0].ID(), orderList.Orders()[0].ProductList(), orderList.Orders()[0].Status())
		t.Errorf("При: %s", buf.String())
	}
}

func TestGetOrderPriceByID(t *testing.T) {
	orderList := CreateOrderList(5)
	err := fillOrderListWithTestData(&orderList)
	if err != nil {
		t.Errorf("GetOrderPriceByID() = %s", err)
	}

	tests := []struct {
		testName    string
		id          int
		expectedSum float64
		wantErr     bool
	}{
		{
			testName:    "Валидный кейс",
			id:          1,
			expectedSum: 2155,
			wantErr:     false,
		},
		{
			testName:    "Не найден id",
			id:          999,
			expectedSum: 0,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			sum, err := orderList.GetOrderPriceByID(tt.id)

			if err != nil && !tt.wantErr {
				t.Errorf("GetOrderPriceByID(%d) = %s", tt.id, err)
			}

			if sum != float64(tt.expectedSum) {
				t.Errorf("GetOrderPriceByID(%d) = %f", tt.id, sum)
			}
		})
	}
}

func TestGetOrderAmount(t *testing.T) {
	orderList := CreateOrderList(5)
	err := fillOrderListWithTestData(&orderList)
	if err != nil {
		t.Errorf("GetOrderAmount() = %s", err)
	}

	ordersAmount1 := orderList.GetOrderAmount()

	product, err := StoreProduct("Рубашка", 100, 20)

	if err != nil {
		t.Errorf("GetOrderAmount() = %s", err)
	}

	orderList.Store("Smith", []Product{product}, "new")
	ordersAmount2 := orderList.GetOrderAmount()

	if ordersAmount1+1 != ordersAmount2 {
		t.Errorf("GetOrderAmount() = %d, при %d", ordersAmount2, ordersAmount1)
	}
}

func TestGetOrdersByIDs(t *testing.T) {
	orderList := CreateOrderList(5)
	err := fillOrderListWithTestData(&orderList)
	if err != nil {
		t.Errorf("GetOrdersByIDs() = %s", err)
	}

	tests := []struct {
		testName    string
		ids         []int
		expectedSum float64
		wantErr     bool
	}{
		{
			testName:    "Валидный кейс",
			ids:         []int{1, 2, 3},
			expectedSum: 2155,
			wantErr:     false,
		},
		{
			testName:    "Не найден id",
			ids:         []int{1, 2, 3},
			expectedSum: 0,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			foundOrders := orderList.GetOrdersByIDs(tt.ids)

			for i := range foundOrders {
				if !slices.Contains(tt.ids, foundOrders[i].ID()) {
					t.Errorf("GetOrdersByIDs(%v) = %v", tt.ids, foundOrders)
				}
			}
		})
	}
}

func TestCreateOrderList(t *testing.T) {
	tests := []struct {
		testName                      string
		capacityArg, expectedCapacity int
	}{
		{
			testName:         "Емкость 10",
			capacityArg:      10,
			expectedCapacity: 10,
		},
		{
			testName:         "Емкость меньше 1",
			capacityArg:      0,
			expectedCapacity: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			orderList := CreateOrderList(tt.capacityArg)
			if tt.expectedCapacity != cap(orderList.Orders()) {
				t.Errorf("CreateOrderList(%d), cap(orderList) = %d", tt.capacityArg, cap(orderList.Orders()))
			}
		})
	}
}
