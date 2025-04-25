package test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/wrkit/xcoin_client"
)

func TestUpdatePairSettings(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name          string
		handlerFunc   http.HandlerFunc
		keyWork       int
		settings      *xcoin_client.KeyValueResponse
		expectedError bool
	}{
		{
			name:    "successful request",
			keyWork: 123,
			settings: &xcoin_client.KeyValueResponse{
				Val: []xcoin_client.KeyValue{
					{
						Key:   "Val1",
						Value: "BTC",
					},
					{
						Key:   "Val2",
						Value: "USD",
					},
				},
			},
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodPost {
					t.Errorf("expected POST request, got %s", r.Method)
				}
				if r.Header.Get("Content-Type") != "application/json" {
					t.Errorf("expected Content-Type application/json, got %s", r.Header.Get("Content-Type"))
				}

				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					t.Fatalf("failed to read request body: %v", err)
				}

				var settings xcoin_client.KeyValueResponse
				if err := json.Unmarshal(body, &settings); err != nil {
					t.Fatalf("failed to unmarshal request body: %v", err)
				}

				w.WriteHeader(http.StatusOK)
			},
		},
		{
			name:    "server error",
			keyWork: 456,
			settings: &xcoin_client.KeyValueResponse{
				Val: []xcoin_client.KeyValue{
					{
						Key:   "Val1",
						Value: "BTC",
					},
				},
			},
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			expectedError: true,
		},
		{
			name:    "invalid settings",
			keyWork: 789,
			settings: &xcoin_client.KeyValueResponse{
				Val: []xcoin_client.KeyValue{},
			},
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

			err := client.UpdatePairSettings(ctx, tt.keyWork, tt.settings)

			if tt.expectedError && err == nil {
				t.Error("expected error but got none")
			}

			if !tt.expectedError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
