package test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/wrkit/xcoin_client"
)

func TestGetPairList2(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name           string
		handlerFunc    http.HandlerFunc
		expectedError  bool
		expectedResult []xcoin_client.Pair
	}{
		{
			name: "successful request",
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet {
					t.Errorf("expected GET request, got %s", r.Method)
				}

				response := []xcoin_client.Pair{
					{
						KeyWork:     1,
						Description: "BTC/USD",
						Status:      "Active",
						StatusInt:   1,
						Setting: xcoin_client.Setting{
							Val1:  "BTC",
							Val2:  "USD",
							Birga: 1,
						},
					},
					{
						KeyWork:     2,
						Description: "ETH/USD",
						Status:      "Active",
						StatusInt:   1,
						Setting: xcoin_client.Setting{
							Val1:  "ETH",
							Val2:  "USD",
							Birga: 1,
						},
					},
				}

				json.NewEncoder(w).Encode(response)
			},
			expectedResult: []xcoin_client.Pair{
				{
					KeyWork:     1,
					Description: "BTC/USD",
					Status:      "Active",
					StatusInt:   1,
					Setting: xcoin_client.Setting{
						Val1:  "BTC",
						Val2:  "USD",
						Birga: 1,
					},
				},
				{
					KeyWork:     2,
					Description: "ETH/USD",
					Status:      "Active",
					StatusInt:   1,
					Setting: xcoin_client.Setting{
						Val1:  "ETH",
						Val2:  "USD",
						Birga: 1,
					},
				},
			},
		},
		{
			name: "server error",
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			expectedError: true,
		},
		{
			name: "invalid response format",
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

			result, err := client.GetPairList2(ctx)

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
					t.Errorf("expected %d pairs, got %d", len(tt.expectedResult), len(result))
					return
				}

				for i, pair := range tt.expectedResult {
					if result[i].KeyWork != pair.KeyWork {
						t.Errorf("expected KeyWork %d, got %d", pair.KeyWork, result[i].KeyWork)
					}
					if result[i].Description != pair.Description {
						t.Errorf("expected Description %s, got %s", pair.Description, result[i].Description)
					}
					if result[i].Moneta != pair.Moneta {
						t.Errorf("expected Status %f, got %f", pair.Moneta, result[i].Moneta)
					}
					if result[i].Setting.Val1 != pair.Setting.Val1 {
						t.Errorf("expected Val1 %s, got %s", pair.Setting.Val1, result[i].Setting.Val1)
					}
					if result[i].Setting.Val2 != pair.Setting.Val2 {
						t.Errorf("expected Val2 %s, got %s", pair.Setting.Val2, result[i].Setting.Val2)
					}
				}
			}
		})
	}
}
