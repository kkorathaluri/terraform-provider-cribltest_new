// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type GetLookupFileByIDRequest struct {
	// Unique ID to GET
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetLookupFileByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// GetLookupFileByIDResponseBody - a list of LookupFile objects
type GetLookupFileByIDResponseBody struct {
	// number of items present in the items array
	Count *int64              `json:"count,omitempty"`
	Items []shared.LookupFile `json:"items,omitempty"`
}

func (o *GetLookupFileByIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *GetLookupFileByIDResponseBody) GetItems() []shared.LookupFile {
	if o == nil {
		return nil
	}
	return o.Items
}

type GetLookupFileByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of LookupFile objects
	Object *GetLookupFileByIDResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *GetLookupFileByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLookupFileByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLookupFileByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetLookupFileByIDResponse) GetObject() *GetLookupFileByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *GetLookupFileByIDResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
