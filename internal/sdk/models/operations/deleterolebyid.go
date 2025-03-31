// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type DeleteRoleByIDRequest struct {
	// Unique ID to DELETE
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteRoleByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// DeleteRoleByIDResponseBody - a list of Role objects
type DeleteRoleByIDResponseBody struct {
	// number of items present in the items array
	Count *int64        `json:"count,omitempty"`
	Items []shared.Role `json:"items,omitempty"`
}

func (o *DeleteRoleByIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *DeleteRoleByIDResponseBody) GetItems() []shared.Role {
	if o == nil {
		return nil
	}
	return o.Items
}

type DeleteRoleByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of Role objects
	Object *DeleteRoleByIDResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *DeleteRoleByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteRoleByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteRoleByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteRoleByIDResponse) GetObject() *DeleteRoleByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *DeleteRoleByIDResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
