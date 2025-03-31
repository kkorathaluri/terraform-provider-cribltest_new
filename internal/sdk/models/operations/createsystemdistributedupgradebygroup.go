// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type CreateSystemDistributedUpgradeByGroupRequest struct {
	// Group to upgrade
	Group string `pathParam:"style=simple,explode=false,name=group"`
	// distributedUpgrade object
	DistributedUpgradeRequest shared.DistributedUpgradeRequest `request:"mediaType=application/json"`
}

func (o *CreateSystemDistributedUpgradeByGroupRequest) GetGroup() string {
	if o == nil {
		return ""
	}
	return o.Group
}

func (o *CreateSystemDistributedUpgradeByGroupRequest) GetDistributedUpgradeRequest() shared.DistributedUpgradeRequest {
	if o == nil {
		return shared.DistributedUpgradeRequest{}
	}
	return o.DistributedUpgradeRequest
}

// CreateSystemDistributedUpgradeByGroupResponseBody - a list of string objects
type CreateSystemDistributedUpgradeByGroupResponseBody struct {
	// number of items present in the items array
	Count *int64   `json:"count,omitempty"`
	Items []string `json:"items,omitempty"`
}

func (o *CreateSystemDistributedUpgradeByGroupResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *CreateSystemDistributedUpgradeByGroupResponseBody) GetItems() []string {
	if o == nil {
		return nil
	}
	return o.Items
}

type CreateSystemDistributedUpgradeByGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of string objects
	Object *CreateSystemDistributedUpgradeByGroupResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *CreateSystemDistributedUpgradeByGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateSystemDistributedUpgradeByGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateSystemDistributedUpgradeByGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateSystemDistributedUpgradeByGroupResponse) GetObject() *CreateSystemDistributedUpgradeByGroupResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *CreateSystemDistributedUpgradeByGroupResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
