package services

import (
	"bytes"
	"io"
	"orders-processor/internal/order"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestProcessOrders(t *testing.T) {
	product1, err1 := order.StoreProduct("Рубашка", 200.5, 10)
	product2, err2 := order.StoreProduct("Апельсин", 10, 10)
	product3, err3 := order.StoreProduct("Скрепка", 5, 10)

	if err1 != nil || err2 != nil || err3 != nil {
		t.Errorf("ProcessOrders() = Ошибка при заполнении заказа")
	}

	productList := []order.Product{product1, product2, product3}

	order1, err1 := order.StoreOrder(1, "Smith", productList, "new")
	order2, err2 := order.StoreOrder(2, "Smith", productList, "new")
	order3, err3 := order.StoreOrder(3, "Smith", productList, "new")

	if err1 != nil || err2 != nil || err3 != nil {
		t.Errorf("ProcessOrders() = Ошибка при заполнении заказа")
	}

	orderList := []*order.Order{&order1, &order2, &order3}

	tests := []struct {
		name         string
		jobs         []*order.Order
		wantResults  []ProcessResult
		wantOutput   []string
		wantInOutput []string
		workerID     int
	}{
		{
			name: "Несколько работ",
			jobs: orderList,
			wantResults: []ProcessResult{
				ProcessResult{Completed: true, TotalSum: 2155},
				ProcessResult{Completed: true, TotalSum: 2155},
				ProcessResult{Completed: true, TotalSum: 2155},
			},
			wantOutput: []string{"", ""},
			wantInOutput: []string{"Worker 1 начал обработку заказа #1",
				"Worker 1 закончил обработку заказа #1, его итоговая сумма 2155.00",
				"Worker 1 начал обработку заказа #1",
				"Worker 1 закончил обработку заказа #1, его итоговая сумма 2155.00",
				"Worker 1 начал обработку заказа #1",
				"Worker 1 закончил обработку заказа #1, его итоговая сумма 2155.00",
			},
			workerID: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ordersChan := make(chan *order.Order, len(tt.jobs))
			resultChan := make(chan ProcessResult, len(tt.jobs))
			var wg sync.WaitGroup
			wg.Add(1)

			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w
			defer func() { os.Stdout = oldStdout }()

			for i := range tt.jobs {
				ordersChan <- tt.jobs[i]
			}

			close(ordersChan)

			ProcessOrders(ordersChan, resultChan, tt.workerID, &wg)

			w.Close()
			close(resultChan)

			var gotResults []ProcessResult
			for result := range resultChan {
				gotResults = append(gotResults, result)
			}
			if len(gotResults) != len(tt.wantResults) {
				t.Errorf("ProcessOrders() = got %d results, want %d", len(gotResults), len(tt.wantResults))
			}
			for i := range gotResults {
				if gotResults[i] != tt.wantResults[i] {
					t.Errorf("ProcessOrders() = at index %d: got result %#v, want %#v", i, gotResults[i], tt.wantResults[i])
				}
			}

			var buf bytes.Buffer
			_, _ = io.Copy(&buf, r)
			outputStr := buf.String()

			for _, expectedLog := range tt.wantInOutput {
				if !strings.Contains(outputStr, expectedLog) {
					t.Errorf("ProcessOrders() = expected log output to contain %q, but it didn't.\nFull output:\n%s", expectedLog, outputStr)
				}
			}

		})
	}
}
