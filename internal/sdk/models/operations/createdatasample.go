// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

// CreateDataSampleResponseBody - a list of DataSample objects
type CreateDataSampleResponseBody struct {
	// number of items present in the items array
	Count *int64              `json:"count,omitempty"`
	Items []shared.DataSample `json:"items,omitempty"`
}

func (o *CreateDataSampleResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *CreateDataSampleResponseBody) GetItems() []shared.DataSample {
	if o == nil {
		return nil
	}
	return o.Items
}

type CreateDataSampleResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of DataSample objects
	Object *CreateDataSampleResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *CreateDataSampleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateDataSampleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateDataSampleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateDataSampleResponse) GetObject() *CreateDataSampleResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *CreateDataSampleResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
