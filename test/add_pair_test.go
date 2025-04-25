package test

import (
	"context"
	"net/http"
	"testing"
)

func TestAddPair(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name          string
		handlerFunc   http.HandlerFunc
		val1          string
		val2          string
		birga         string
		mode          string
		expectedError bool
	}{
		{
			name:  "successful request",
			val1:  "BTC",
			val2:  "USD",
			birga: "binance",
			mode:  "spot",
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodPost {
					t.Errorf("expected POST request, got %s", r.Method)
				}
				if r.Header.Get("Content-Type") != "application/json" {
					t.Errorf("expected Content-Type application/json, got %s", r.Header.Get("Content-Type"))
				}
				w.WriteHeader(http.StatusOK)
			},
		},
		{
			name:  "server error",
			val1:  "BTC",
			val2:  "USD",
			birga: "binance",
			mode:  "spot",
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			expectedError: true,
		},
		{
			name:  "invalid parameters",
			val1:  "",
			val2:  "USD",
			birga: "binance",
			mode:  "spot",
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusBadRequest)
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, server := setupTestClient(tt.handlerFunc)
			defer server.Close()

			err := client.AddPair(ctx, tt.val1, tt.val2, tt.birga, tt.mode)

			if tt.expectedError && err == nil {
				t.Error("expected error but got none")
			}

			if !tt.expectedError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
