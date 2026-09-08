package storage

import (
	"orders-processor/internal/order"
	"testing"
)

func TestFillOrderListWithTestData(t *testing.T) {
	orderList := order.CreateOrderList(5)
	err := FillOrderListWithTestData(&orderList)
	if err != nil {
		t.Errorf("TestFillOrderListWithTestData() при %#v", orderList)
	}
}
