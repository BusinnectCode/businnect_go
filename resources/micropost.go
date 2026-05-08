package resources

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/Businnect/businnect_go/interfaces"
	"github.com/Businnect/businnect_go/schemas"
)

type MicropostResource struct {
	client interfaces.ClientInterface
}

type CreateMicropostRequest struct {
	Body                           string
	Title                          *string
	PostType                       schemas.PostType
	CommunityName                  *string
	ParentMicropostPublicID        *string
	ParentMicropostCommentPublicID *string
	AllowComments                  bool
	IsDraft                        bool
	File                           *FilePayload
}

type FilePayload struct {
	Filename string
	Reader   io.Reader
	MimeType string
}

func NewMicropostResource(c interfaces.ClientInterface) *MicropostResource {
	return &MicropostResource{client: c}
}

func (r *MicropostResource) Create(ctx context.Context, params CreateMicropostRequest) (*schemas.CreatedWithPublicIdAndLinkResponse, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	fields := map[string]string{
		"body":           params.Body,
		"allow_comments": strconv.FormatBool(params.AllowComments),
		"is_draft":       strconv.FormatBool(params.IsDraft),
	}

	if params.PostType == "" {
		fields["post_type"] = string(schemas.PostTypeText)
	} else {
		fields["post_type"] = string(params.PostType)
	}

	for key, val := range fields {
		if err := w.WriteField(key, val); err != nil {
			return nil, err
		}
	}

	if params.Title != nil {
		_ = w.WriteField("title", *params.Title)
	}
	if params.CommunityName != nil {
		_ = w.WriteField("community_name", *params.CommunityName)
	}
	if params.ParentMicropostPublicID != nil {
		_ = w.WriteField("parent_micropost_public_id", *params.ParentMicropostPublicID)
	}
	if params.ParentMicropostCommentPublicID != nil {
		_ = w.WriteField("parent_micropost_comment_public_id", *params.ParentMicropostCommentPublicID)
	}

	if params.File != nil {
		fw, err := w.CreateFormFile("file", params.File.Filename)
		if err != nil {
			return nil, err
		}
		if _, err = io.Copy(fw, params.File.Reader); err != nil {
			return nil, err
		}
	}

	if err := w.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", "/api/v1/client/micropost", &b)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", w.FormDataContentType())

	resBody, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resBody.Close()

	var result schemas.CreatedWithPublicIdAndLinkResponse
	if err := json.NewDecoder(resBody).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *MicropostResource) Delete(ctx context.Context, publicID string) error {
	data := map[string]string{
		"public_id": publicID,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodDelete, "/api/v1/client/micropost", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	httpRequest.Header.Set("Content-Type", "application/json")

	respBody, err := r.client.Do(httpRequest)
	if err != nil {
		return err
	}

	defer respBody.Close()
	return nil
}

func (r *MicropostResource) Vote(ctx context.Context, publicID string) (*schemas.UserMicropostVoteResponse, error) {
	data := map[string]string{
		"public_id": publicID,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodPatch, "/api/v1/client/micropost/vote", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	httpRequest.Header.Set("Content-Type", "application/json")

	respBody, err := r.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer respBody.Close()

	var response schemas.UserMicropostVoteResponse
	if err := json.NewDecoder(respBody).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
