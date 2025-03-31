// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type UpdatePipelineByIDRequest struct {
	// Unique ID to PATCH
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Pipeline object to be updated
	Pipeline shared.Pipeline `request:"mediaType=application/json"`
}

func (o *UpdatePipelineByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdatePipelineByIDRequest) GetPipeline() shared.Pipeline {
	if o == nil {
		return shared.Pipeline{}
	}
	return o.Pipeline
}

// UpdatePipelineByIDResponseBody - a list of Pipeline objects
type UpdatePipelineByIDResponseBody struct {
	// number of items present in the items array
	Count *int64            `json:"count,omitempty"`
	Items []shared.Pipeline `json:"items,omitempty"`
}

func (o *UpdatePipelineByIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *UpdatePipelineByIDResponseBody) GetItems() []shared.Pipeline {
	if o == nil {
		return nil
	}
	return o.Items
}

type UpdatePipelineByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of Pipeline objects
	Object *UpdatePipelineByIDResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *UpdatePipelineByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdatePipelineByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdatePipelineByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdatePipelineByIDResponse) GetObject() *UpdatePipelineByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *UpdatePipelineByIDResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
