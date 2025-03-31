// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

// CreateKeyMetadataEntityResponseBody - a list of KeyMetadataEntity objects
type CreateKeyMetadataEntityResponseBody struct {
	// number of items present in the items array
	Count *int64                     `json:"count,omitempty"`
	Items []shared.KeyMetadataEntity `json:"items,omitempty"`
}

func (o *CreateKeyMetadataEntityResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *CreateKeyMetadataEntityResponseBody) GetItems() []shared.KeyMetadataEntity {
	if o == nil {
		return nil
	}
	return o.Items
}

type CreateKeyMetadataEntityResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of KeyMetadataEntity objects
	Object *CreateKeyMetadataEntityResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *CreateKeyMetadataEntityResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateKeyMetadataEntityResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateKeyMetadataEntityResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateKeyMetadataEntityResponse) GetObject() *CreateKeyMetadataEntityResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *CreateKeyMetadataEntityResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
