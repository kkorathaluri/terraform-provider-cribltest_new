// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type DeleteTeamByIDRequest struct {
	// Unique ID to DELETE
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteTeamByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// DeleteTeamByIDResponseBody - a list of Team objects
type DeleteTeamByIDResponseBody struct {
	// number of items present in the items array
	Count *int64        `json:"count,omitempty"`
	Items []shared.Team `json:"items,omitempty"`
}

func (o *DeleteTeamByIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *DeleteTeamByIDResponseBody) GetItems() []shared.Team {
	if o == nil {
		return nil
	}
	return o.Items
}

type DeleteTeamByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of Team objects
	Object *DeleteTeamByIDResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *DeleteTeamByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteTeamByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteTeamByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteTeamByIDResponse) GetObject() *DeleteTeamByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *DeleteTeamByIDResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
