// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type GetJobsErrorsByIDAndGroupRequest struct {
	// Instance id of the job to get results for
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Group the job belongs to
	Group string `pathParam:"style=simple,explode=false,name=group"`
}

func (o *GetJobsErrorsByIDAndGroupRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetJobsErrorsByIDAndGroupRequest) GetGroup() string {
	if o == nil {
		return ""
	}
	return o.Group
}

// GetJobsErrorsByIDAndGroupResponseBody - a list of any objects
type GetJobsErrorsByIDAndGroupResponseBody struct {
	// number of items present in the items array
	Count *int64           `json:"count,omitempty"`
	Items []map[string]any `json:"items,omitempty"`
}

func (o *GetJobsErrorsByIDAndGroupResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *GetJobsErrorsByIDAndGroupResponseBody) GetItems() []map[string]any {
	if o == nil {
		return nil
	}
	return o.Items
}

type GetJobsErrorsByIDAndGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of any objects
	Object *GetJobsErrorsByIDAndGroupResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *GetJobsErrorsByIDAndGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetJobsErrorsByIDAndGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetJobsErrorsByIDAndGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetJobsErrorsByIDAndGroupResponse) GetObject() *GetJobsErrorsByIDAndGroupResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *GetJobsErrorsByIDAndGroupResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
