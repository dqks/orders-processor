package order

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type Product struct {
	name   string
	price  float64
	amount int
}

func StoreProduct(name string, price float64, amount int) (Product, error) {
	err := ValidateProductFields(name, price, amount)
	if err != nil {
		return Product{}, err
	}
	return Product{name: name, price: price, amount: amount}, nil
}

func ValidateProductFields(name string, price float64, amount int) error {
	if price <= 0 || amount <= 0 || len(strings.TrimSpace(name)) == 0 {
		return errors.New("Невалидная информация о товаре")
	}
	return nil
}

func PrintProduct(w io.Writer, p Product) {
	fmt.Fprintf(w, "Название: %s, цена: %.2f, количество: %d\n", p.name, p.price, p.amount)
}

func (p Product) GetTotalPrice() float64 {
	return p.price * float64(p.amount)
}
