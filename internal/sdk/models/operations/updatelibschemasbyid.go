// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type UpdateLibSchemasByIDRequest struct {
	// Unique ID to PATCH
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Schema object to be updated
	SchemaLibEntry shared.SchemaLibEntry `request:"mediaType=application/json"`
}

func (o *UpdateLibSchemasByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateLibSchemasByIDRequest) GetSchemaLibEntry() shared.SchemaLibEntry {
	if o == nil {
		return shared.SchemaLibEntry{}
	}
	return o.SchemaLibEntry
}

// UpdateLibSchemasByIDResponseBody - a list of Schema objects
type UpdateLibSchemasByIDResponseBody struct {
	// number of items present in the items array
	Count *int64                  `json:"count,omitempty"`
	Items []shared.SchemaLibEntry `json:"items,omitempty"`
}

func (o *UpdateLibSchemasByIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *UpdateLibSchemasByIDResponseBody) GetItems() []shared.SchemaLibEntry {
	if o == nil {
		return nil
	}
	return o.Items
}

type UpdateLibSchemasByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of Schema objects
	Object *UpdateLibSchemasByIDResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *UpdateLibSchemasByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateLibSchemasByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateLibSchemasByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateLibSchemasByIDResponse) GetObject() *UpdateLibSchemasByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *UpdateLibSchemasByIDResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
