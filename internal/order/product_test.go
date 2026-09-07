package order

import "testing"

func TestStoreProduct(t *testing.T) {

}

func TestValidateProductFields(t *testing.T) {
	tests := []struct {
		name    string
		price   float64
		amount  int
		wantErr bool
	}{
		{
			name:    "Рубашка",
			price:   100,
			amount:  30,
			wantErr: false,
		},
		{
			name:    "",
			price:   100,
			amount:  30,
			wantErr: true,
		},
		{
			name:    "Рубашка",
			price:   0,
			amount:  30,
			wantErr: true,
		},
		{
			name:    "Рубашка",
			price:   100,
			amount:  0,
			wantErr: true,
		},
		{
			name:    "Рубашка",
			price:   0,
			amount:  -1,
			wantErr: true,
		},
		{
			name:    "Рубашка",
			price:   100,
			amount:  -1,
			wantErr: true,
		},
		{
			name:    "",
			price:   0,
			amount:  0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateProductFields(tt.name, tt.price, tt.amount)

			if err != nil && tt.wantErr {

			} else if err == nil && !tt.wantErr {
			} else {
				t.Errorf("ValidateProductFields(%s, %f, %d) = %s", tt.name, tt.price, tt.amount, err)
			}

		})
	}

}

func TestGetTotalPrice(t *testing.T) {
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
