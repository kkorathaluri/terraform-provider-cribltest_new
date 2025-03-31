// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type GetSystemProjectsVersionCountByProjectIDRequest struct {
	// Project ID
	ProjectID string `pathParam:"style=simple,explode=false,name=projectId"`
	// Commit ID
	ID *string `queryParam:"style=form,explode=true,name=ID"`
}

func (o *GetSystemProjectsVersionCountByProjectIDRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *GetSystemProjectsVersionCountByProjectIDRequest) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// GetSystemProjectsVersionCountByProjectIDResponseBody - a list of any objects
type GetSystemProjectsVersionCountByProjectIDResponseBody struct {
	// number of items present in the items array
	Count *int64           `json:"count,omitempty"`
	Items []map[string]any `json:"items,omitempty"`
}

func (o *GetSystemProjectsVersionCountByProjectIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *GetSystemProjectsVersionCountByProjectIDResponseBody) GetItems() []map[string]any {
	if o == nil {
		return nil
	}
	return o.Items
}

type GetSystemProjectsVersionCountByProjectIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of any objects
	Object *GetSystemProjectsVersionCountByProjectIDResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *GetSystemProjectsVersionCountByProjectIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSystemProjectsVersionCountByProjectIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSystemProjectsVersionCountByProjectIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSystemProjectsVersionCountByProjectIDResponse) GetObject() *GetSystemProjectsVersionCountByProjectIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *GetSystemProjectsVersionCountByProjectIDResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
