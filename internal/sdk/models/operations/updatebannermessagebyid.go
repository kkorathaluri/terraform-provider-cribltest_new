// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type UpdateBannerMessageByIDRequest struct {
	// Unique ID to PATCH
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// BannerMessage object to be updated
	BannerMessage shared.BannerMessage `request:"mediaType=application/json"`
}

func (o *UpdateBannerMessageByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateBannerMessageByIDRequest) GetBannerMessage() shared.BannerMessage {
	if o == nil {
		return shared.BannerMessage{}
	}
	return o.BannerMessage
}

// UpdateBannerMessageByIDResponseBody - a list of BannerMessage objects
type UpdateBannerMessageByIDResponseBody struct {
	// number of items present in the items array
	Count *int64                 `json:"count,omitempty"`
	Items []shared.BannerMessage `json:"items,omitempty"`
}

func (o *UpdateBannerMessageByIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *UpdateBannerMessageByIDResponseBody) GetItems() []shared.BannerMessage {
	if o == nil {
		return nil
	}
	return o.Items
}

type UpdateBannerMessageByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of BannerMessage objects
	Object *UpdateBannerMessageByIDResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *UpdateBannerMessageByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateBannerMessageByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateBannerMessageByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateBannerMessageByIDResponse) GetObject() *UpdateBannerMessageByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *UpdateBannerMessageByIDResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
