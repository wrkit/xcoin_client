package test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/wrkit/xcoin_client"
)

func TestGetPairOrders(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name           string
		handlerFunc    http.HandlerFunc
		keyWork        int
		expectedError  bool
		expectedResult []xcoin_client.Order
	}{
		{
			name:    "successful request",
			keyWork: 123,
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet {
					t.Errorf("expected GET request, got %s", r.Method)
				}

				response := []xcoin_client.Order{
					{
						ClientOrderId: "1",
						OrderId:       1,
						Price:         50000.0,
					},
					{
						ClientOrderId: "2",
						OrderId:       2,
						Price:         60000.0,
					},
				}

				json.NewEncoder(w).Encode(response)
			},
			expectedResult: []xcoin_client.Order{
				{
					ClientOrderId: "1",
					OrderId:       1,
					Price:         50000.0,
				},
				{
					ClientOrderId: "2",
					OrderId:       2,
					Price:         60000.0,
				},
			},
		},
		{
			name:    "server error",
			keyWork: 456,
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			expectedError: true,
		},
		{
			name:    "invalid response format",
			keyWork: 789,
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(`{"invalid":"json"}`))
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, server := setupTestClient(tt.handlerFunc)
			defer server.Close()

			result, err := client.GetPairOrders(ctx, tt.keyWork)

			if tt.expectedError && err == nil {
				t.Error("expected error but got none")
			}

			if !tt.expectedError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if tt.expectedResult != nil {
				if result == nil {
					t.Error("expected result but got nil")
					return
				}

				if len(result) != len(tt.expectedResult) {
					t.Errorf("expected %d orders, got %d", len(tt.expectedResult), len(result))
					return
				}

				for i, order := range result {
					expected := tt.expectedResult[i]
					if order.ClientOrderId != expected.ClientOrderId {
						t.Errorf("order %d: expected ID %s, got %s", i, expected.ClientOrderId, order.ClientOrderId)
					}
					if order.OrderId != expected.OrderId {
						t.Errorf("order %d: expected Amount %d, got %d", i, expected.OrderId, order.OrderId)
					}
					if order.Price != expected.Price {
						t.Errorf("order %d: expected Price %f, got %f", i, expected.Price, order.Price)
					}
				}
			}
		})
	}
}
