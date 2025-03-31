// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

// ListNotificationTargetResponseBody - a list of NotificationTarget objects
type ListNotificationTargetResponseBody struct {
	// number of items present in the items array
	Count *int64                      `json:"count,omitempty"`
	Items []shared.NotificationTarget `json:"items,omitempty"`
}

func (o *ListNotificationTargetResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *ListNotificationTargetResponseBody) GetItems() []shared.NotificationTarget {
	if o == nil {
		return nil
	}
	return o.Items
}

type ListNotificationTargetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of NotificationTarget objects
	Object *ListNotificationTargetResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *ListNotificationTargetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListNotificationTargetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListNotificationTargetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListNotificationTargetResponse) GetObject() *ListNotificationTargetResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *ListNotificationTargetResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
