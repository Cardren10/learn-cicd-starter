package auth

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name           string
		headers        http.Header
		expectedKey    string
		expectedError  error
		errorContains  string
	}{
		{
			name:           "Valid API Key",
			headers:        http.Header{"Authorization": []string{"ApiKey test-key-123"}},
			expectedKey:    "test-key-123",
			expectedError:  nil,
		},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				key, err := GetAPIKey(tt.headers)
	
				// Check the returned key
				if key != tt.expectedKey {
					t.Errorf("expected key %q, got %q", tt.expectedKey, key)
				}
	
				// Check the error
				if tt.expectedError != nil {
					if err != tt.expectedError {
						t.Errorf("expected error %v, got %v", tt.expectedError, err)
					}
				} else if tt.errorContains != "" {
					if err == nil {
						t.Errorf("expected error containing %q, got nil", tt.errorContains)
					} else if !strings.Contains(err.Error(), tt.errorContains) {
						t.Errorf("expected error containing %q, got %v", tt.errorContains, err)
					}
				} else if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			})
		}
	}