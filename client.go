package businnect

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/BusinnectCode/businnect_go/resources"
)

type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client

	Micropost *resources.MicropostResource
	AdminBlog *resources.AdminBlogResource
}

func NewClient(apiKey string, baseURL string) *Client {
	c := &Client{
		baseURL: strings.TrimSuffix(baseURL, "/"),
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}

	c.Micropost = resources.NewMicropostResource(c)
	c.AdminBlog = resources.NewAdminBlogResource(c)

	return c
}

func (c *Client) Do(req *http.Request) (io.ReadCloser, error) {
	if !strings.HasPrefix(req.URL.String(), "http") {
		pathWithQuery := req.URL.Path
		if req.URL.RawQuery != "" {
			pathWithQuery += "?" + req.URL.RawQuery
		}
		req.URL, _ = req.URL.Parse(c.baseURL + pathWithQuery)
	}

	req.Header.Set("USER-API-TOKEN", c.apiKey)
	req.Header.Set("User-Agent", "businnect-Go-SDK/1.0")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		defer resp.Body.Close()

		var errorResp struct {
			Detail string `json:"detail"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&errorResp); err == nil && errorResp.Detail != "" {
			return nil, fmt.Errorf("api error: %s", errorResp.Detail)
		}

		return nil, fmt.Errorf("api error: status %d", resp.StatusCode)
	}

	return resp.Body, nil
}
