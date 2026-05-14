package resources_test

import (
	"context"
	"os"
	"testing"

	businnect "github.com/BusinnectCode/businnect_go"
	"github.com/BusinnectCode/businnect_go/resources"
)

func TestAdminBlogGet_Integration(t *testing.T) {
	token := os.Getenv("BUSINNECT_API_TOKEN")
	if token == "" {
		t.Skip("Skipping integration test: BUSINNECT_API_TOKEN not set")
	}

	client := businnect.NewClient(token, "https://api.businnect.com")
	listRes, err := client.AdminBlog.List(context.Background())
	if err != nil {
		t.Fatalf("failed to list admin blogs before get: %v", err)
	}

	if listRes == nil || len(listRes.List) == 0 {
		t.Skip("Skipping integration test: no admin blog entries available")
	}

	slug := listRes.List[0].Slug
	if slug == "" {
		t.Skip("Skipping integration test: admin blog entry has empty slug")
	}

	res, err := client.AdminBlog.Get(context.Background(), slug)
	if err != nil {
		t.Fatalf("integration test failed: %v", err)
	}

	if res != nil {
		if _, ok := res["slug"]; !ok {
			t.Log("admin blog response is not null but has no slug field")
		}
	}
}

func TestAdminBlogList_Integration(t *testing.T) {
	token := os.Getenv("BUSINNECT_API_TOKEN")
	if token == "" {
		t.Skip("Skipping integration test: BUSINNECT_API_TOKEN not set")
	}

	client := businnect.NewClient(token, "https://api.businnect.com")

	limit := 5
	skip := 0
	params := resources.ListAdminBlogsRequest{
		Limit: &limit,
		Skip:  &skip,
	}

	res, err := client.AdminBlog.ListWithParams(context.Background(), params)
	if err != nil {
		t.Fatalf("integration test failed: %v", err)
	}

	if res == nil {
		t.Fatal("expected non-nil list response")
	}

	if res.TotalList < 0 || res.TotalPages < 0 {
		t.Fatalf("expected non-negative totals, got total_list=%d total_pages=%d", res.TotalList, res.TotalPages)
	}

	if len(res.List) > limit {
		t.Fatalf("expected at most %d items, got %d", limit, len(res.List))
	}

	if len(res.List) > 0 {
		title := res.List[0].Title
		category := string(res.List[0].Category)

		filteredRes, filterErr := client.AdminBlog.ListWithParams(context.Background(), resources.ListAdminBlogsRequest{
			Title:    &title,
			Category: &category,
			Limit:    &limit,
			Skip:     &skip,
		})
		if filterErr != nil {
			t.Fatalf("integration filtered list test failed: %v", filterErr)
		}

		if filteredRes == nil {
			t.Fatal("expected non-nil filtered list response")
		}
	}
}
