package test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/wrkit/xcoin_client"
)

func setupTestClient(handler http.HandlerFunc) (xcoin_client.IClient, *httptest.Server) {
	server := httptest.NewServer(handler)

	cfg := xcoin_client.NewConfig(server.URL)

	client := xcoin_client.NewClient(cfg)
	return client, server
}

func TestGetVersion(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name           string
		handlerFunc    http.HandlerFunc
		expectedError  bool
		expectedResult *xcoin_client.VersionResponse
	}{
		{
			name: "successful request",
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet {
					t.Errorf("expected GET request, got %s", r.Method)
				}

				response := &xcoin_client.VersionResponse{
					BotID: "101",
					Pairs: []xcoin_client.PairInfo{
						{
							KeyWork: 1,
							Pair:    "BTC/USD",
						},
					},
				}

				if err := json.NewEncoder(w).Encode(response); err != nil {
					t.Errorf("failed to encode response: %v", err)
				}
			},
			expectedResult: &xcoin_client.VersionResponse{
				BotID: "101",
				Pairs: []xcoin_client.PairInfo{
					{
						KeyWork: 1,
						Pair:    "BTC/USD",
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, server := setupTestClient(tt.handlerFunc)
			defer server.Close()

			result, err := client.GetVersion(ctx)

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

				if result.BotID != tt.expectedResult.BotID {
					t.Errorf("expected version %s, got %s", tt.expectedResult.BotID, result.BotID)
				}

				if len(result.Pairs) != len(tt.expectedResult.Pairs) {
					for i, pair := range tt.expectedResult.Pairs {
						if result.Pairs[i].KeyWork != pair.KeyWork {
							t.Errorf("expected keywork %d, got %d", pair.KeyWork, result.Pairs[i].KeyWork)
						}
						if result.Pairs[i].Pair != pair.Pair {
							t.Errorf("expected pair %s, got %s", pair.Pair, result.Pairs[i].Pair)
						}
					}

				}
			}
		})
	}
}
