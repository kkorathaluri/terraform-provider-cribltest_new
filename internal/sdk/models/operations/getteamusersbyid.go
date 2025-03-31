// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type GetTeamUsersByIDRequest struct {
	// Team Name
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetTeamUsersByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// GetTeamUsersByIDResponseBody - a list of string objects
type GetTeamUsersByIDResponseBody struct {
	// number of items present in the items array
	Count *int64   `json:"count,omitempty"`
	Items []string `json:"items,omitempty"`
}

func (o *GetTeamUsersByIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *GetTeamUsersByIDResponseBody) GetItems() []string {
	if o == nil {
		return nil
	}
	return o.Items
}

type GetTeamUsersByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of string objects
	Object *GetTeamUsersByIDResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *GetTeamUsersByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTeamUsersByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTeamUsersByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTeamUsersByIDResponse) GetObject() *GetTeamUsersByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *GetTeamUsersByIDResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
