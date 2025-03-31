// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type UpdateDatasetByIDRequest struct {
	// Unique ID to PATCH
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Dataset object to be updated
	RequestBody any `request:"mediaType=application/json"`
}

func (o *UpdateDatasetByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateDatasetByIDRequest) GetRequestBody() any {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

// UpdateDatasetByIDResponseBody - a list of Dataset objects
type UpdateDatasetByIDResponseBody struct {
	// number of items present in the items array
	Count *int64 `json:"count,omitempty"`
	Items []any  `json:"items,omitempty"`
}

func (o *UpdateDatasetByIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *UpdateDatasetByIDResponseBody) GetItems() []any {
	if o == nil {
		return nil
	}
	return o.Items
}

type UpdateDatasetByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of Dataset objects
	Object *UpdateDatasetByIDResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *UpdateDatasetByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateDatasetByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateDatasetByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateDatasetByIDResponse) GetObject() *UpdateDatasetByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *UpdateDatasetByIDResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
