package resources

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/BusinnectCode/businnect_go/interfaces"
	"github.com/BusinnectCode/businnect_go/schemas"
)

type AdminBlogResource struct {
	client interfaces.ClientInterface
}

type ListAdminBlogsRequest struct {
	Title    *string
	Category *string
	Skip     *int
	Limit    *int
}

func NewAdminBlogResource(c interfaces.ClientInterface) *AdminBlogResource {
	return &AdminBlogResource{client: c}
}

func (r *AdminBlogResource) Get(ctx context.Context, slug string) (schemas.AdminBlogArticle, error) {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, "/api/v1/client/admin/blog", nil)
	if err != nil {
		return nil, err
	}

	query := url.Values{}
	query.Set("slug", slug)
	httpReq.URL.RawQuery = query.Encode()

	respBody, err := r.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer respBody.Close()

	var result schemas.AdminBlogArticle
	if err := json.NewDecoder(respBody).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *AdminBlogResource) List(ctx context.Context) (*schemas.AdminBlogListResponse, error) {
	return r.ListWithParams(ctx, ListAdminBlogsRequest{})
}

func (r *AdminBlogResource) ListWithParams(ctx context.Context, params ListAdminBlogsRequest) (*schemas.AdminBlogListResponse, error) {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, "/api/v1/client/admin/blog/list", nil)
	if err != nil {
		return nil, err
	}

	query := url.Values{}
	if params.Title != nil {
		query.Set("title", *params.Title)
	}
	if params.Category != nil {
		query.Set("category", *params.Category)
	}
	if params.Skip != nil {
		query.Set("skip", strconv.Itoa(*params.Skip))
	}
	if params.Limit != nil {
		query.Set("limit", strconv.Itoa(*params.Limit))
	}
	httpReq.URL.RawQuery = query.Encode()

	respBody, err := r.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer respBody.Close()

	var result schemas.AdminBlogListResponse
	if err := json.NewDecoder(respBody).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
