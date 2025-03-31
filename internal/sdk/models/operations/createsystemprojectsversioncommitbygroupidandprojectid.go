// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type CreateSystemProjectsVersionCommitByGroupIDAndProjectIDRequest struct {
	// Group Id
	GroupID string `pathParam:"style=simple,explode=false,name=groupId"`
	// Project Id
	ProjectID string `pathParam:"style=simple,explode=false,name=projectId"`
	// ProjectGitCommitParams object
	ProjectGitCommitParams shared.ProjectGitCommitParams `request:"mediaType=application/json"`
}

func (o *CreateSystemProjectsVersionCommitByGroupIDAndProjectIDRequest) GetGroupID() string {
	if o == nil {
		return ""
	}
	return o.GroupID
}

func (o *CreateSystemProjectsVersionCommitByGroupIDAndProjectIDRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *CreateSystemProjectsVersionCommitByGroupIDAndProjectIDRequest) GetProjectGitCommitParams() shared.ProjectGitCommitParams {
	if o == nil {
		return shared.ProjectGitCommitParams{}
	}
	return o.ProjectGitCommitParams
}

// CreateSystemProjectsVersionCommitByGroupIDAndProjectIDResponseBody - A list of GitCommitSummary objects
type CreateSystemProjectsVersionCommitByGroupIDAndProjectIDResponseBody struct {
	// Number of items present in the items array
	Count *int64                    `json:"count,omitempty"`
	Items []shared.GitCommitSummary `json:"items,omitempty"`
}

func (o *CreateSystemProjectsVersionCommitByGroupIDAndProjectIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *CreateSystemProjectsVersionCommitByGroupIDAndProjectIDResponseBody) GetItems() []shared.GitCommitSummary {
	if o == nil {
		return nil
	}
	return o.Items
}

type CreateSystemProjectsVersionCommitByGroupIDAndProjectIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A list of GitCommitSummary objects
	Object *CreateSystemProjectsVersionCommitByGroupIDAndProjectIDResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *CreateSystemProjectsVersionCommitByGroupIDAndProjectIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateSystemProjectsVersionCommitByGroupIDAndProjectIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateSystemProjectsVersionCommitByGroupIDAndProjectIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateSystemProjectsVersionCommitByGroupIDAndProjectIDResponse) GetObject() *CreateSystemProjectsVersionCommitByGroupIDAndProjectIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *CreateSystemProjectsVersionCommitByGroupIDAndProjectIDResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
