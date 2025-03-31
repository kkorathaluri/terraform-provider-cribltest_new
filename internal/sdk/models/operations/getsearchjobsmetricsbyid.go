// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type GetSearchJobsMetricsByIDRequest struct {
	// search job id
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetSearchJobsMetricsByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetSearchJobsMetricsByIDResponse struct {
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

func (o *GetSearchJobsMetricsByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSearchJobsMetricsByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSearchJobsMetricsByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSearchJobsMetricsByIDResponse) GetString() *string {
	if o == nil {
		return nil
	}
	return o.String
}

func (o *GetSearchJobsMetricsByIDResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
