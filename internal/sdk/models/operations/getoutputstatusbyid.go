// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type GetOutputStatusByIDRequest struct {
	// Unique ID to GET
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetOutputStatusByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// GetOutputStatusByIDResponseBody - a list of OutputStatus objects
type GetOutputStatusByIDResponseBody struct {
	// number of items present in the items array
	Count *int64                `json:"count,omitempty"`
	Items []shared.OutputStatus `json:"items,omitempty"`
}

func (o *GetOutputStatusByIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *GetOutputStatusByIDResponseBody) GetItems() []shared.OutputStatus {
	if o == nil {
		return nil
	}
	return o.Items
}

type GetOutputStatusByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of OutputStatus objects
	Object *GetOutputStatusByIDResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *GetOutputStatusByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetOutputStatusByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetOutputStatusByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetOutputStatusByIDResponse) GetObject() *GetOutputStatusByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *GetOutputStatusByIDResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
