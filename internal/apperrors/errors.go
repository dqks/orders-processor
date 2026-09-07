package appErrors

import "errors"

var (
	ErrOrderNotFound = errors.New("Заказ не найден")
)
