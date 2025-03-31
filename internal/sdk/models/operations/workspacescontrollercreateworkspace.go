// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type WorkspacesControllerCreateWorkspaceRequest struct {
	OrganizationID     string                    `pathParam:"style=simple,explode=false,name=organizationId"`
	CreateWorkspaceDTO shared.CreateWorkspaceDTO `request:"mediaType=application/json"`
}

func (o *WorkspacesControllerCreateWorkspaceRequest) GetOrganizationID() string {
	if o == nil {
		return ""
	}
	return o.OrganizationID
}

func (o *WorkspacesControllerCreateWorkspaceRequest) GetCreateWorkspaceDTO() shared.CreateWorkspaceDTO {
	if o == nil {
		return shared.CreateWorkspaceDTO{}
	}
	return o.CreateWorkspaceDTO
}

type WorkspacesControllerCreateWorkspaceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse  *http.Response
	WorkspaceDTO *shared.WorkspaceDTO
}

func (o *WorkspacesControllerCreateWorkspaceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *WorkspacesControllerCreateWorkspaceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *WorkspacesControllerCreateWorkspaceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *WorkspacesControllerCreateWorkspaceResponse) GetWorkspaceDTO() *shared.WorkspaceDTO {
	if o == nil {
		return nil
	}
	return o.WorkspaceDTO
}
