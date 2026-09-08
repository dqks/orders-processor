package order

import (
	"bytes"
	"testing"
)

func TestStoreProduct(t *testing.T) {
	tests := []struct {
		testName, name string
		price          float64
		amount         int
		wantErr        bool
	}{
		{
			testName: "Валидный продукт",
			name:     "Рубашка",
			price:    100,
			amount:   30,
			wantErr:  false,
		},
		{
			testName: "Продукт без названия",
			name:     "",
			price:    100,
			amount:   30,
			wantErr:  true,
		},
		{
			testName: "Продукт с ценой 0",
			name:     "Рубашка",
			price:    0,
			amount:   30,
			wantErr:  true,
		},
		{
			testName: "Продукт без количеством 0",
			name:     "Рубашка",
			price:    100,
			amount:   0,
			wantErr:  true,
		},
		{
			testName: "Проудкт с отрицательной ценой",
			name:     "Рубашка",
			price:    -1,
			amount:   0,
			wantErr:  true,
		},
		{
			testName: "Проудкт с отрицательным количеством",
			name:     "Рубашка",
			price:    100,
			amount:   -1,
			wantErr:  true,
		},
		{
			testName: "Продукт со всеми полями по умолчанию",
			name:     "",
			price:    0,
			amount:   0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := StoreProduct(tt.name, tt.price, tt.amount)
			if (err != nil && !tt.wantErr) || (err == nil && tt.wantErr) {
				t.Errorf("StoreProduct(%s, %f, %d) = %s", tt.name, tt.price, tt.amount, err)
			}
		})
	}
}

func TestValidateProductFields(t *testing.T) {
	tests := []struct {
		testName, name string
		price          float64
		amount         int
		wantErr        bool
	}{
		{
			testName: "Валидный продукт",
			name:     "Рубашка",
			price:    100,
			amount:   30,
			wantErr:  false,
		},
		{
			testName: "Продукт без названия",
			name:     "",
			price:    100,
			amount:   30,
			wantErr:  true,
		},
		{
			testName: "Продукт с ценой 0",
			name:     "Рубашка",
			price:    0,
			amount:   30,
			wantErr:  true,
		},
		{
			testName: "Продукт без количеством 0",
			name:     "Рубашка",
			price:    100,
			amount:   0,
			wantErr:  true,
		},
		{
			testName: "Проудкт с отрицательной ценой",
			name:     "Рубашка",
			price:    -1,
			amount:   0,
			wantErr:  true,
		},
		{
			testName: "Проудкт с отрицательным количеством",
			name:     "Рубашка",
			price:    100,
			amount:   -1,
			wantErr:  true,
		},
		{
			testName: "Продукт со всеми полями по умолчанию",
			name:     "",
			price:    0,
			amount:   0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			err := ValidateProductFields(tt.name, tt.price, tt.amount)
			if (err != nil && !tt.wantErr) || (err == nil && tt.wantErr) {
				t.Errorf("StoreProduct(%s, %f, %d) = %s", tt.name, tt.price, tt.amount, err)
			}
		})
	}
}

func TestPrintProduct(t *testing.T) {
	product, err := StoreProduct("Рубашка", 200, 100)
	if err != nil {
		t.Errorf("TestPrintProduct() = %s, при %s, %.2f, %d", err, product.name, product.price, product.amount)
	}

	var buf bytes.Buffer
	PrintProduct(&buf, product)

	if buf.String() != "Название: Рубашка, цена: 200.00, количество: 100\n" {
		t.Errorf("TestPrintProduct() = %s, при %s, %.2f, %d", err, product.name, product.price, product.amount)
	}
}

func TestProductGetTotalPrice(t *testing.T) {
	product := Product{
		name:   "Test",
		price:  10,
		amount: 200,
	}

	expected := float64(2000)

	result := product.GetTotalPrice()

	if expected != result {
		t.Errorf("product.GetTotalPrice(%f, %d) = %f; ожидалось %f", product.price, product.amount, result, expected)
	}
}
