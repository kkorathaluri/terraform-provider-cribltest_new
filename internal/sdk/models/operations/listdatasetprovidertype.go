// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

// ListDatasetProviderTypeResponseBody - a list of DatasetProviderType objects
type ListDatasetProviderTypeResponseBody struct {
	// number of items present in the items array
	Count *int64                       `json:"count,omitempty"`
	Items []shared.DatasetProviderType `json:"items,omitempty"`
}

func (o *ListDatasetProviderTypeResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *ListDatasetProviderTypeResponseBody) GetItems() []shared.DatasetProviderType {
	if o == nil {
		return nil
	}
	return o.Items
}

type ListDatasetProviderTypeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of DatasetProviderType objects
	Object *ListDatasetProviderTypeResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *ListDatasetProviderTypeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListDatasetProviderTypeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListDatasetProviderTypeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListDatasetProviderTypeResponse) GetObject() *ListDatasetProviderTypeResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *ListDatasetProviderTypeResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
