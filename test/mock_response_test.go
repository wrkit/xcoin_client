package test

import (
	"context"
	"net/http"
	"os"
	"testing"
)

func TestMockResponses(t *testing.T) {
	ctx := context.Background()

	// Читаем содержимое файлов с реальными ответами
	settingsData, err := os.ReadFile("mock/get_pair_current_setting.json")
	if err != nil {
		t.Fatalf("Failed to read mock settings file: %v", err)
	}

	ordersData, err := os.ReadFile("mock/get_current_order.json")
	if err != nil {
		t.Fatalf("Failed to read mock orders file: %v", err)
	}

	pairsData, err := os.ReadFile("mock/get_list_pair2.json")
	if err != nil {
		t.Fatalf("Failed to read mock pairs file: %v", err)
	}

	tests := []struct {
		name           string
		handlerFunc    http.HandlerFunc
		keyWork        int
		expectedError  bool
		expectedStatus string
		expectedValues map[string]string
	}{
		{
			name:    "get pair settings with stop wait status",
			keyWork: 16,
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				// Используем реальный JSON-ответ из файла
				w.Write(settingsData)
			},
			expectedStatus: "StopWait",
			expectedValues: map[string]string{
				"Val1":      "ATOM",
				"Val2":      "FDUSD",
				"AllProfit": "17,0572",
			},
		},
		{
			name:    "get pair settings with running status",
			keyWork: 6,
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				// Используем тот же JSON-ответ для второго теста
				w.Write(settingsData)
			},
			expectedStatus: "StopWait",
			expectedValues: map[string]string{
				"Val1":      "ATOM",
				"Val2":      "FDUSD",
				"AllProfit": "17,0572",
			},
		},
	}

	// Тесты для GetPairCurOrder
	orderTests := []struct {
		name           string
		handlerFunc    http.HandlerFunc
		keyWork        int
		expectedError  bool
		expectedValues map[string]interface{}
	}{
		{
			name:    "get current orders",
			keyWork: 1,
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.Write(ordersData)
			},
			expectedValues: map[string]interface{}{
				"Symbol":   "ATOMFDUSD",
				"OrderId":  20770493,
				"Price":    11.5550000000,
				"Quantity": 1.7300000000,
				"Status":   2,
				"myManual": 0,
			},
		},
	}

	// Тесты для GetListPair2
	pairTests := []struct {
		name           string
		handlerFunc    http.HandlerFunc
		expectedError  bool
		expectedValues map[string]interface{}
	}{
		{
			name: "get pairs list",
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.Write(pairsData)
			},
			expectedValues: map[string]interface{}{
				"Birga":       "Binance",
				"Description": "ATOMFDUSD up FDUSD",
				"KeyWork":     1,
				"Moneta":      59.5300000000,
				"Price":       13.7180000000,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, server := setupTestClient(tt.handlerFunc)
			defer server.Close()

			result, err := client.GetPairCurrentSettings(ctx, tt.keyWork)

			if tt.expectedError {
				if err == nil {
					t.Error("expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if result == nil {
				t.Error("expected result but got nil")
				return
			}

			// Check status
			for _, kv := range result.Val {
				if kv.Key == "Status" && kv.Value != tt.expectedStatus {
					t.Errorf("expected status %s, got %s", tt.expectedStatus, kv.Value)
				}
			}

			// Check expected values
			for key, expectedValue := range tt.expectedValues {
				found := false
				for _, kv := range result.Val {
					if kv.Key == key {
						found = true
						if kv.Value != expectedValue {
							t.Errorf("expected %s=%s, got %s", key, expectedValue, kv.Value)
						}
						break
					}
				}
				if !found {
					t.Errorf("expected key %s not found in response", key)
				}
			}
		})
	}

	// Выполняем тесты для GetPairCurOrder
	for _, tt := range orderTests {
		t.Run(tt.name, func(t *testing.T) {
			client, server := setupTestClient(tt.handlerFunc)
			defer server.Close()

			result, err := client.GetPairOrders(ctx, tt.keyWork)

			if tt.expectedError {
				if err == nil {
					t.Error("expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if len(result) == 0 {
				t.Error("expected non-empty result but got empty")
				return
			}

			// Проверяем первый ордер
			order := result[0]
			for key, expectedValue := range tt.expectedValues {
				switch key {
				case "Symbol":
					if order.Symbol != expectedValue.(string) {
						t.Errorf("expected Symbol %v, got %v", expectedValue, order.Symbol)
					}
				case "OrderId":
					if order.OrderId != expectedValue.(int) {
						t.Errorf("expected OrderId %v, got %v", expectedValue, order.OrderId)
					}
				case "Price":
					if order.Price != expectedValue.(float64) {
						t.Errorf("expected Price %v, got %v", expectedValue, order.Price)
					}
				case "Quantity":
					if order.Quantity != expectedValue.(float64) {
						t.Errorf("expected Quantity %v, got %v", expectedValue, order.Quantity)
					}
				case "Status":
					if order.Status != expectedValue.(int) {
						t.Errorf("expected Status %v, got %v", expectedValue, order.Status)
					}
				case "myManual":
					if order.MyManual != expectedValue.(int) {
						t.Errorf("expected myManual %v, got %v", expectedValue, order.MyManual)
					}
				}
			}
		})
	}

	// Выполняем тесты для GetListPair2
	for _, tt := range pairTests {
		t.Run(tt.name, func(t *testing.T) {
			client, server := setupTestClient(tt.handlerFunc)
			defer server.Close()

			result, err := client.GetPairList2(ctx)

			if tt.expectedError {
				if err == nil {
					t.Error("expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if len(result) == 0 {
				t.Error("expected non-empty result but got empty")
				return
			}

			// Проверяем первую пару
			pair := result[0]
			for key, expectedValue := range tt.expectedValues {
				switch key {
				case "Birga":
					if pair.Birga != expectedValue.(string) {
						t.Errorf("expected Birga %v, got %v", expectedValue, pair.Setting.Birga)
					}
				case "Description":
					if pair.Description != expectedValue.(string) {
						t.Errorf("expected Description %v, got %v", expectedValue, pair.Description)
					}
				case "KeyWork":
					if pair.KeyWork != expectedValue.(int) {
						t.Errorf("expected KeyWork %v, got %v", expectedValue, pair.KeyWork)
					}
				case "Moneta":
					if pair.Moneta != expectedValue.(float64) {
						t.Errorf("expected Moneta %v, got %v", expectedValue, pair.Moneta)
					}
				case "Price":
					if pair.Price != expectedValue.(float64) {
						t.Errorf("expected Price %v, got %v", expectedValue, pair.Price)
					}
				}
			}
		})
	}
}
