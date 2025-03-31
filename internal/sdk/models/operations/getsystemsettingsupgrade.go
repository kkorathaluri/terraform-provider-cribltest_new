// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

// GetSystemSettingsUpgradeResponseBody - a list of UpgradeResult objects
type GetSystemSettingsUpgradeResponseBody struct {
	// number of items present in the items array
	Count *int64                 `json:"count,omitempty"`
	Items []shared.UpgradeResult `json:"items,omitempty"`
}

func (o *GetSystemSettingsUpgradeResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *GetSystemSettingsUpgradeResponseBody) GetItems() []shared.UpgradeResult {
	if o == nil {
		return nil
	}
	return o.Items
}

type GetSystemSettingsUpgradeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of UpgradeResult objects
	Object *GetSystemSettingsUpgradeResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *GetSystemSettingsUpgradeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSystemSettingsUpgradeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSystemSettingsUpgradeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSystemSettingsUpgradeResponse) GetObject() *GetSystemSettingsUpgradeResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *GetSystemSettingsUpgradeResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
