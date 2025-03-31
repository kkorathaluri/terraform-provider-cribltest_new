// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type DeleteSavedQueryByIDRequest struct {
	// Unique ID to DELETE
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteSavedQueryByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// DeleteSavedQueryByIDResponseBody - a list of SavedQuery objects
type DeleteSavedQueryByIDResponseBody struct {
	// number of items present in the items array
	Count *int64              `json:"count,omitempty"`
	Items []shared.SavedQuery `json:"items,omitempty"`
}

func (o *DeleteSavedQueryByIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *DeleteSavedQueryByIDResponseBody) GetItems() []shared.SavedQuery {
	if o == nil {
		return nil
	}
	return o.Items
}

type DeleteSavedQueryByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of SavedQuery objects
	Object *DeleteSavedQueryByIDResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *DeleteSavedQueryByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteSavedQueryByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteSavedQueryByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteSavedQueryByIDResponse) GetObject() *DeleteSavedQueryByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *DeleteSavedQueryByIDResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
