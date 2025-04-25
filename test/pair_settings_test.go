package test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/wrkit/xcoin_client"
)

func TestGetPairSettings(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name           string
		handlerFunc    http.HandlerFunc
		keyWork        int
		expectedError  bool
		expectedResult *xcoin_client.KeyValueResponse
	}{
		{
			name:    "successful request",
			keyWork: 123,
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet {
					t.Errorf("expected GET request, got %s", r.Method)
				}

				response := &xcoin_client.KeyValueResponse{
					Val: []xcoin_client.KeyValue{
						{
							Key:   "Val1",
							Value: "BTC",
						},
						{
							Key:   "Val2",
							Value: "USD",
						},
						{
							Key:   "Birga",
							Value: "binance",
						},
					},
				}

				json.NewEncoder(w).Encode(response)
			},
			expectedResult: &xcoin_client.KeyValueResponse{
				Val: []xcoin_client.KeyValue{
					{
						Key:   "Val1",
						Value: "BTC",
					},
					{
						Key:   "Val2",
						Value: "USD",
					},
					{
						Key:   "Birga",
						Value: "binance",
					},
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

			result, err := client.GetPairCurrentSettings(ctx, tt.keyWork)

			if !tt.expectedError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if tt.expectedResult != nil {
				if result == nil {
					t.Error("expected result but got nil")
					return
				}

				for i, val := range tt.expectedResult.Val {
					if result.Val[i].Key != val.Key || result.Val[i].Value != val.Value {
						t.Errorf("expected key %s, got %s, Value: %s got %s", val.Key, result.Val[i].Key, val.Value, result.Val[0].Value)
					}

				}

			}
		})
	}
}
