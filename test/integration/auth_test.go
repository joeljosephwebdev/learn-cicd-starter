package yourpackage_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/joeljosephwebdev/learn-cicd-starter/internal/auth" // replace with actual import path
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError error
	}{
		{
			name:          "no authorization header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: auth.ErrNoAuthHeaderIncluded,
		},
		{
			name: "malformed header - no ApiKey prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer xyz123"},
			},
			expectedKey:   "",
			expectedError: errors.New("malformed authorization header"),
		},
		{
			name: "malformed header - only prefix",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectedKey:   "",
			expectedError: errors.New("malformed authorization header"),
		},
		{
			name: "valid ApiKey header",
			headers: http.Header{
				"Authorization": []string{"ApiKey abc123"},
			},
			expectedKey:   "abc123",
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := auth.GetAPIKey(tt.headers)

			if key != tt.expectedKey {
				t.Errorf("expected key %q, got %q", tt.expectedKey, key)
			}

			if tt.expectedError != nil {
				if err == nil || err.Error() != tt.expectedError.Error() {
					t.Errorf("expected error %q, got %v", tt.expectedError, err)
				}
			} else if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}
}
