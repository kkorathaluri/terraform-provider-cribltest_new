// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type GetInvoicesRequest struct {
	OrganizationID string `pathParam:"style=simple,explode=false,name=organizationId"`
}

func (o *GetInvoicesRequest) GetOrganizationID() string {
	if o == nil {
		return ""
	}
	return o.OrganizationID
}

type GetInvoicesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse    *http.Response
	GetInvoicesDTO *shared.GetInvoicesDTO
}

func (o *GetInvoicesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetInvoicesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetInvoicesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetInvoicesResponse) GetGetInvoicesDTO() *shared.GetInvoicesDTO {
	if o == nil {
		return nil
	}
	return o.GetInvoicesDTO
}
