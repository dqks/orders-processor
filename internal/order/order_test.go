package order

import (
	"bytes"
	"slices"
	"testing"
)

var productList []Product = []Product{
	{name: "Рубашка", price: 100, amount: 20},
	{name: "Апельсин", price: 200, amount: 20},
	{name: "Яблоко", price: 300, amount: 50},
}

func TestID(t *testing.T) {
	order, err := StoreOrder(1, "Smith", productList, "new")
	if err != nil {
		t.Errorf("ID() = %s, при %d, %#v, %s", err, order.id, order.productList, order.status)
	}
	expected := 1
	result := order.ID()
	if result != expected {
		t.Errorf("ID() = %d, при %d, %#v, %s", result, order.id, order.productList, order.status)
	}
}

func TestClientName(t *testing.T) {
	order, err := StoreOrder(1, "Smith", productList, "new")
	if err != nil {
		t.Errorf("ClientName() = %s, при %d, %#v, %s", err, order.id, order.productList, order.status)
	}
	expected := "Smith"
	result := order.ClientName()
	if expected != result {
		t.Errorf("ClientName() = %s, при %d, %#v, %s", result, order.id, order.productList, order.status)
	}
}

func TestProductList(t *testing.T) {
	order, err := StoreOrder(1, "Smith", productList, "new")
	if err != nil {
		t.Errorf("ProductList() = %s, при %d, %#v, %s", err, order.id, order.productList, order.status)
	}
	expected := order.productList
	result := order.ProductList()
	if !slices.Equal(expected, result) {
		t.Errorf("ProductList() = %#v, при %d, %#v, %s", result, order.id, order.productList, order.status)
	}
}

func TestStatus(t *testing.T) {
	order, err := StoreOrder(1, "Smith", productList, "new")
	if err != nil {
		t.Errorf("Status() = %s, при %d, %#v, %s", err, order.id, order.productList, order.status)
	}
	expected := order.status
	result := order.Status()
	if expected != result {
		t.Errorf("Status() = %s, при %d, %#v, %s", result, order.id, order.productList, order.status)
	}
}

func TestPrint(t *testing.T) {
	order, err := StoreOrder(1, "Smith", productList, "new")
	if err != nil {
		t.Errorf("Print() = %s, при %d, %#v, %s", err, order.id, order.productList, order.status)
	}
	var buf bytes.Buffer
	expected := "id: 1, клиент: Smith, статус: new\n" +
		"Товары:\n" +
		"Название: Рубашка, цена: 100.00, количество: 20\n" +
		"Название: Апельсин, цена: 200.00, количество: 20\n" +
		"Название: Яблоко, цена: 300.00, количество: 50\n" +
		"\n"
	order.Print(&buf)
	if buf.String() != expected {
		t.Errorf("Print(&buf), при %d, %#v, %s", order.id, order.productList, order.status)
		t.Errorf("При: %s", buf.String())
	}
}

func TestSetClientName(t *testing.T) {
	order, err := StoreOrder(1, "Smith", productList, "new")
	if err != nil {
		t.Errorf("SetClientName() = %s, при %d, %#v, %s", err, order.id, order.productList, order.status)
	}

	tests := []struct {
		testName   string
		clientName string
		wantErr    bool
	}{
		{
			testName:   "Валидный",
			clientName: "John",
			wantErr:    false,
		},
		{
			testName:   "Пустой",
			clientName: "",
			wantErr:    true,
		},
		{
			testName:   "Пустой с пробелами",
			clientName: "      ",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			err := order.SetClientName(tt.clientName)
			if (err == nil && tt.wantErr) || (err != nil && !tt.wantErr) {
				t.Errorf("SetClientName(%s) = %s, при %d, %#v, %s", tt.clientName, err, order.id, order.productList, order.status)
			}
		})
	}
}

func TestSetStatus(t *testing.T) {
	order, err := StoreOrder(1, "Smith", productList, "new")
	if err != nil {
		t.Errorf("SetStatus() = %s, при %d, %#v, %s", err, order.id, order.productList, order.status)
	}

	tests := []struct {
		testName string
		status   string
		wantErr  bool
	}{
		{
			testName: "Валидный new",
			status:   "new",
			wantErr:  false,
		},
		{
			testName: "Валидный сompleted",
			status:   "completed",
			wantErr:  false,
		},
		{
			testName: "Валидный error",
			status:   "error",
			wantErr:  false,
		},
		{
			testName: "Невалидный пустой",
			status:   "",
			wantErr:  true,
		},
		{
			testName: "Невалидный",
			status:   "test",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			err := order.SetStatus(tt.status)
			if (err == nil && tt.wantErr) || (err != nil && !tt.wantErr) {
				t.Errorf("SetStatus(%s) = %s, при %d, %#v, %s", tt.status, err, order.id, order.productList, order.status)
			}
		})
	}
}

func TestStoreOrder(t *testing.T) {
	tests := []struct {
		testName    string
		id          int
		clientName  string
		productList []Product
		status      string
		wantErr     bool
	}{
		{
			testName:    "Валидные данные",
			id:          1,
			clientName:  "Smith",
			productList: productList,
			status:      "new",
			wantErr:     false,
		},
		{
			testName:    "ID < 0",
			id:          -1,
			clientName:  "Smith",
			productList: productList,
			status:      "new",
			wantErr:     true,
		},
		{
			testName:    "Пустое имя клиента",
			id:          1,
			clientName:  "",
			productList: productList,
			status:      "new",
			wantErr:     true,
		},
		{
			testName:    "Пустой статус",
			id:          1,
			clientName:  "Smith",
			productList: productList,
			status:      "",
			wantErr:     true,
		},
		{
			testName:    "Невалидный статус",
			id:          1,
			clientName:  "Smith",
			productList: productList,
			status:      "dadas",
			wantErr:     true,
		},
		{
			testName:    "Невалидный продукт",
			id:          1,
			clientName:  "Smith",
			productList: append(productList, Product{name: "", price: 10, amount: 20}),
			status:      "dadas",
			wantErr:     true,
		},
		{
			testName:    "Невалидный продукт",
			id:          1,
			clientName:  "Smith",
			productList: append(productList, Product{name: "Рубашка", price: 0, amount: 20}),
			status:      "dadas",
			wantErr:     true,
		},
		{
			testName:    "Невалидный продукт",
			id:          1,
			clientName:  "Smith",
			productList: append(productList, Product{name: "Рубашка", price: -1, amount: 20}),
			status:      "dadas",
			wantErr:     true,
		},
		{
			testName:    "Невалидный продукт",
			id:          1,
			clientName:  "Smith",
			productList: append(productList, Product{name: "Рубашка", price: 1, amount: 0}),
			status:      "dadas",
			wantErr:     true,
		},
		{
			testName:    "Невалидный продукт",
			id:          1,
			clientName:  "Smith",
			productList: append(productList, Product{name: "Рубашка", price: 1, amount: -1}),
			status:      "dadas",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			_, err := StoreOrder(tt.id, tt.clientName, tt.productList, tt.status)
			if (err == nil && tt.wantErr) || (err != nil && !tt.wantErr) {
				t.Errorf("StoreOrder(%d, %s, %#v, %s) = %s", tt.id, tt.clientName, tt.productList, tt.status, err)
			}
		})
	}
}

func TestValidateOrderFields(t *testing.T) {
	tests := []struct {
		testName    string
		id          int
		clientName  string
		productList []Product
		status      string
		wantErr     bool
	}{
		{
			testName:    "Валидные данные",
			id:          1,
			clientName:  "Smith",
			productList: productList,
			status:      "new",
			wantErr:     false,
		},
		{
			testName:    "ID < 0",
			id:          -1,
			clientName:  "Smith",
			productList: productList,
			status:      "new",
			wantErr:     true,
		},
		{
			testName:    "Пустое имя клиента",
			id:          1,
			clientName:  "",
			productList: productList,
			status:      "new",
			wantErr:     true,
		},
		{
			testName:    "Пустой статус",
			id:          1,
			clientName:  "Smith",
			productList: productList,
			status:      "",
			wantErr:     true,
		},
		{
			testName:    "Невалидный статус",
			id:          1,
			clientName:  "Smith",
			productList: productList,
			status:      "dadas",
			wantErr:     true,
		},
		{
			testName:    "Невалидный продукт",
			id:          1,
			clientName:  "Smith",
			productList: append(productList, Product{name: "", price: 10, amount: 20}),
			status:      "dadas",
			wantErr:     true,
		},
		{
			testName:    "Невалидный продукт",
			id:          1,
			clientName:  "Smith",
			productList: append(productList, Product{name: "Рубашка", price: 0, amount: 20}),
			status:      "dadas",
			wantErr:     true,
		},
		{
			testName:    "Невалидный продукт",
			id:          1,
			clientName:  "Smith",
			productList: append(productList, Product{name: "Рубашка", price: -1, amount: 20}),
			status:      "dadas",
			wantErr:     true,
		},
		{
			testName:    "Невалидный продукт",
			id:          1,
			clientName:  "Smith",
			productList: append(productList, Product{name: "Рубашка", price: 1, amount: 0}),
			status:      "dadas",
			wantErr:     true,
		},
		{
			testName:    "Невалидный продукт",
			id:          1,
			clientName:  "Smith",
			productList: append(productList, Product{name: "Рубашка", price: 1, amount: -1}),
			status:      "dadas",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			err := ValidateOrderFields(tt.id, tt.clientName, tt.productList, tt.status)
			if (err == nil && tt.wantErr) || (err != nil && !tt.wantErr) {
				t.Errorf("ValidateOrderFields(%d, %s, %#v, %s) = %s", tt.id, tt.clientName, tt.productList, tt.status, err)
			}
		})
	}
}

func TestOrderGetTotalPrice(t *testing.T) {
	order, err := StoreOrder(1, "Smith", productList, "new")
	if err != nil {
		t.Errorf("OrderGetTotalPrice() = %s, при %d, %#v, %s", err, order.id, order.productList, order.status)
	}
	expected := float64(21000)
	result := order.GetTotalPrice()
	if result != expected {
		t.Errorf("OrderGetTotalPrice() = %f, при %d, %#v, %s", result, order.id, order.productList, order.status)
	}
}
