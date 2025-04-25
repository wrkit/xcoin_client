package test

import (
	"context"
	"net/http"
	"testing"

	"github.com/wrkit/xcoin_client"
)

func TestExecuteCommand(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name          string
		handlerFunc   http.HandlerFunc
		keyWork       int
		command       xcoin_client.CommandType
		expectedError bool
	}{
		{
			name:    "successful request",
			keyWork: 123,
			command: xcoin_client.CommandRunned,
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
			name:    "server error",
			keyWork: 456,
			command: xcoin_client.CommandStop,
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			expectedError: true,
		},
		{
			name:    "invalid command",
			keyWork: 789,
			command: "invalid",
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

			err := client.ExecuteCommand(ctx, tt.keyWork, tt.command)

			if tt.expectedError && err == nil {
				t.Error("expected error but got none")
			}

			if !tt.expectedError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
