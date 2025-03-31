// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type GetSearchJobMetricsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// string object
	String *string
	// Unexpected error
	Error *shared.Error
}

func (o *GetSearchJobMetricsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSearchJobMetricsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSearchJobMetricsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSearchJobMetricsResponse) GetString() *string {
	if o == nil {
		return nil
	}
	return o.String
}

func (o *GetSearchJobMetricsResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
