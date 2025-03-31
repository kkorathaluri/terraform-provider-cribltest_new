// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

// ListSavedQueryResponseBody - a list of SavedQuery objects
type ListSavedQueryResponseBody struct {
	// number of items present in the items array
	Count *int64              `json:"count,omitempty"`
	Items []shared.SavedQuery `json:"items,omitempty"`
}

func (o *ListSavedQueryResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *ListSavedQueryResponseBody) GetItems() []shared.SavedQuery {
	if o == nil {
		return nil
	}
	return o.Items
}

type ListSavedQueryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of SavedQuery objects
	Object *ListSavedQueryResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *ListSavedQueryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListSavedQueryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListSavedQueryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListSavedQueryResponse) GetObject() *ListSavedQueryResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *ListSavedQueryResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
