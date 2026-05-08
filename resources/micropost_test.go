package resources_test

import (
	"context"
	"os"
	"testing"

	businnect "github.com/Businnect/businnect_go"
	"github.com/Businnect/businnect_go/resources"
)

func TestClientCreateMicropost_Integration(t *testing.T) {

	token := os.Getenv("BUSINNECT_API_TOKEN")
	if token == "" {
		t.Skip("Skipping integration test: BUSINNECT_API_TOKEN not set")
	}

	client := businnect.NewClient(token, "https://api.businnect.com")

	title := "Test"
	req := resources.CreateMicropostRequest{
		Title: &title,
		Body:  "Hello",
	}

	ctx := context.Background()
	response, err := client.Micropost.Create(ctx, req)

	if err != nil {
		t.Fatalf("Integration test failed: %v", err)
	}

	if response.PublicID == "" {
		t.Fatalf("Expected a non-empty PublicID in the response")
	}

	t.Cleanup(func() {
		_, voteErr := client.Micropost.Vote(ctx, response.PublicID)
		if voteErr != nil {
			t.Fatalf("Warning: Failed to vote micropost %s before cleanup: %v", response.PublicID, voteErr)
		}

		err := client.Micropost.Delete(ctx, response.PublicID)
		if err != nil {
			t.Fatalf("Warning: Failed to cleanup micropost %s: %v", response.PublicID, err)
		} else {
			t.Logf("Successfully cleaned up micropost: %s", response.PublicID)
		}
	})
}
