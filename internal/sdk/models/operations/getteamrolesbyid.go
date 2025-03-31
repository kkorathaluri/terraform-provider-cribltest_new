// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type GetTeamRolesByIDRequest struct {
	// user id
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetTeamRolesByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// GetTeamRolesByIDResponseBody - a list of string objects
type GetTeamRolesByIDResponseBody struct {
	// number of items present in the items array
	Count *int64   `json:"count,omitempty"`
	Items []string `json:"items,omitempty"`
}

func (o *GetTeamRolesByIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *GetTeamRolesByIDResponseBody) GetItems() []string {
	if o == nil {
		return nil
	}
	return o.Items
}

type GetTeamRolesByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of string objects
	Object *GetTeamRolesByIDResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *GetTeamRolesByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTeamRolesByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTeamRolesByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTeamRolesByIDResponse) GetObject() *GetTeamRolesByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *GetTeamRolesByIDResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
