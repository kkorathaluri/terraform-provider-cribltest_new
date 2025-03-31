// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

// ListProfilerItemResponseBody - a list of ProfilerItem objects
type ListProfilerItemResponseBody struct {
	// number of items present in the items array
	Count *int64                `json:"count,omitempty"`
	Items []shared.ProfilerItem `json:"items,omitempty"`
}

func (o *ListProfilerItemResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *ListProfilerItemResponseBody) GetItems() []shared.ProfilerItem {
	if o == nil {
		return nil
	}
	return o.Items
}

type ListProfilerItemResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of ProfilerItem objects
	Object *ListProfilerItemResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *ListProfilerItemResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListProfilerItemResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListProfilerItemResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListProfilerItemResponse) GetObject() *ListProfilerItemResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *ListProfilerItemResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
