package businnect

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/Businnect/businnect_go/resources"
)

func TestClient_Do_Errors(t *testing.T) {
	tests := []struct {
		name           string
		statusResponse int
		bodyResponse   string
		expectedError  string
	}{
		{
			name:           "Invalid API Token 401",
			statusResponse: http.StatusUnauthorized,
			bodyResponse:   `{"invalid_api_user_token": "USER-API-TOKEN header is not valid"}`,
			expectedError:  "USER-API-TOKEN header is not valid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient("fake-key", "https://api.businnect.com")

			ctx := context.Background()

			title := "Test"
			_, err := client.Micropost.Create(ctx, resources.CreateMicropostRequest{
				Title: &title,
				Body:  "Hello",
			})

			if err == nil {
				t.Fatal("expected an error but got nil")
			}

			if !strings.Contains(err.Error(), tt.expectedError) {
				t.Errorf("expected error to contain '%s', got '%v'", tt.expectedError, err)
			}
		})
	}
}
