// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type GetJobsByIDRequest struct {
	// Job instance id
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetJobsByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// GetJobsByIDResponseBody - a list of JobInfo objects
type GetJobsByIDResponseBody struct {
	// number of items present in the items array
	Count *int64           `json:"count,omitempty"`
	Items []shared.JobInfo `json:"items,omitempty"`
}

func (o *GetJobsByIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *GetJobsByIDResponseBody) GetItems() []shared.JobInfo {
	if o == nil {
		return nil
	}
	return o.Items
}

type GetJobsByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of JobInfo objects
	Object *GetJobsByIDResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *GetJobsByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetJobsByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetJobsByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetJobsByIDResponse) GetObject() *GetJobsByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *GetJobsByIDResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
