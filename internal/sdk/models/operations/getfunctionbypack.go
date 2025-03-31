// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type GetFunctionByPackRequest struct {
	// pack ID to GET
	Pack string `pathParam:"style=simple,explode=false,name=pack"`
}

func (o *GetFunctionByPackRequest) GetPack() string {
	if o == nil {
		return ""
	}
	return o.Pack
}

// GetFunctionByPackResponseBody - a list of Function objects
type GetFunctionByPackResponseBody struct {
	// number of items present in the items array
	Count *int64            `json:"count,omitempty"`
	Items []shared.Function `json:"items,omitempty"`
}

func (o *GetFunctionByPackResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *GetFunctionByPackResponseBody) GetItems() []shared.Function {
	if o == nil {
		return nil
	}
	return o.Items
}

type GetFunctionByPackResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of Function objects
	Object *GetFunctionByPackResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *GetFunctionByPackResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetFunctionByPackResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetFunctionByPackResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetFunctionByPackResponse) GetObject() *GetFunctionByPackResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *GetFunctionByPackResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
