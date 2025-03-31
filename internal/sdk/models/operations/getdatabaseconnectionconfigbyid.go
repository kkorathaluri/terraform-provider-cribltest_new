// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type GetDatabaseConnectionConfigByIDRequest struct {
	// Unique ID to GET
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetDatabaseConnectionConfigByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// GetDatabaseConnectionConfigByIDResponseBody - a list of DatabaseConnectionConfig objects
type GetDatabaseConnectionConfigByIDResponseBody struct {
	// number of items present in the items array
	Count *int64                            `json:"count,omitempty"`
	Items []shared.DatabaseConnectionConfig `json:"items,omitempty"`
}

func (o *GetDatabaseConnectionConfigByIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *GetDatabaseConnectionConfigByIDResponseBody) GetItems() []shared.DatabaseConnectionConfig {
	if o == nil {
		return nil
	}
	return o.Items
}

type GetDatabaseConnectionConfigByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of DatabaseConnectionConfig objects
	Object *GetDatabaseConnectionConfigByIDResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *GetDatabaseConnectionConfigByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDatabaseConnectionConfigByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDatabaseConnectionConfigByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetDatabaseConnectionConfigByIDResponse) GetObject() *GetDatabaseConnectionConfigByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *GetDatabaseConnectionConfigByIDResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
