// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type DeleteNotificationTargetByIDRequest struct {
	// Unique ID to DELETE
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteNotificationTargetByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// DeleteNotificationTargetByIDResponseBody - a list of NotificationTarget objects
type DeleteNotificationTargetByIDResponseBody struct {
	// number of items present in the items array
	Count *int64                      `json:"count,omitempty"`
	Items []shared.NotificationTarget `json:"items,omitempty"`
}

func (o *DeleteNotificationTargetByIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *DeleteNotificationTargetByIDResponseBody) GetItems() []shared.NotificationTarget {
	if o == nil {
		return nil
	}
	return o.Items
}

type DeleteNotificationTargetByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of NotificationTarget objects
	Object *DeleteNotificationTargetByIDResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *DeleteNotificationTargetByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteNotificationTargetByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteNotificationTargetByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteNotificationTargetByIDResponse) GetObject() *DeleteNotificationTargetByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *DeleteNotificationTargetByIDResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
