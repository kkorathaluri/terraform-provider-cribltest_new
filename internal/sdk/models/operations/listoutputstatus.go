// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

// ListOutputStatusResponseBody - a list of OutputStatus objects
type ListOutputStatusResponseBody struct {
	// number of items present in the items array
	Count *int64                `json:"count,omitempty"`
	Items []shared.OutputStatus `json:"items,omitempty"`
}

func (o *ListOutputStatusResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *ListOutputStatusResponseBody) GetItems() []shared.OutputStatus {
	if o == nil {
		return nil
	}
	return o.Items
}

type ListOutputStatusResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of OutputStatus objects
	Object *ListOutputStatusResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *ListOutputStatusResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListOutputStatusResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListOutputStatusResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListOutputStatusResponse) GetObject() *ListOutputStatusResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *ListOutputStatusResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
