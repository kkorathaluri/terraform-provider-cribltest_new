// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

// GetVersionInfoResponseBody - a list of GitInfo objects
type GetVersionInfoResponseBody struct {
	// number of items present in the items array
	Count *int64           `json:"count,omitempty"`
	Items []shared.GitInfo `json:"items,omitempty"`
}

func (o *GetVersionInfoResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *GetVersionInfoResponseBody) GetItems() []shared.GitInfo {
	if o == nil {
		return nil
	}
	return o.Items
}

type GetVersionInfoResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of GitInfo objects
	Object *GetVersionInfoResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *GetVersionInfoResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetVersionInfoResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetVersionInfoResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetVersionInfoResponse) GetObject() *GetVersionInfoResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *GetVersionInfoResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
