// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type UpdateProjectByIDRequest struct {
	// Unique ID to PATCH
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Project object to be updated
	ProjectConfig shared.ProjectConfig `request:"mediaType=application/json"`
}

func (o *UpdateProjectByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateProjectByIDRequest) GetProjectConfig() shared.ProjectConfig {
	if o == nil {
		return shared.ProjectConfig{}
	}
	return o.ProjectConfig
}

// UpdateProjectByIDResponseBody - a list of Project objects
type UpdateProjectByIDResponseBody struct {
	// number of items present in the items array
	Count *int64                 `json:"count,omitempty"`
	Items []shared.ProjectConfig `json:"items,omitempty"`
}

func (o *UpdateProjectByIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *UpdateProjectByIDResponseBody) GetItems() []shared.ProjectConfig {
	if o == nil {
		return nil
	}
	return o.Items
}

type UpdateProjectByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of Project objects
	Object *UpdateProjectByIDResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *UpdateProjectByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateProjectByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateProjectByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateProjectByIDResponse) GetObject() *UpdateProjectByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *UpdateProjectByIDResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
