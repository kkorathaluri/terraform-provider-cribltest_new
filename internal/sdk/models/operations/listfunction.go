// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

// ListFunctionResponseBody - a list of Function objects
type ListFunctionResponseBody struct {
	// number of items present in the items array
	Count *int64            `json:"count,omitempty"`
	Items []shared.Function `json:"items,omitempty"`
}

func (o *ListFunctionResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *ListFunctionResponseBody) GetItems() []shared.Function {
	if o == nil {
		return nil
	}
	return o.Items
}

type ListFunctionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of Function objects
	Object *ListFunctionResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *ListFunctionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListFunctionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListFunctionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListFunctionResponse) GetObject() *ListFunctionResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *ListFunctionResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
