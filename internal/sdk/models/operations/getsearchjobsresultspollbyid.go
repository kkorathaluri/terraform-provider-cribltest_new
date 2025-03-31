// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type GetSearchJobsResultsPollByIDRequest struct {
	// search job id
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// maximum number of events returned
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// starting offset of the events
	Offset *float64 `queryParam:"style=form,explode=true,name=offset"`
	// lower bound of the time range, inclusive
	LowerBound *float64 `queryParam:"style=form,explode=true,name=lowerBound"`
	// upper bound of the time range, exclusive
	UpperBound *float64 `queryParam:"style=form,explode=true,name=upperBound"`
	// last known status of the Search Job. Used to return immediatelyupon status change if the status was queued.
	LastJobStatus *string `queryParam:"style=form,explode=true,name=lastJobStatus"`
}

func (o *GetSearchJobsResultsPollByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetSearchJobsResultsPollByIDRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetSearchJobsResultsPollByIDRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetSearchJobsResultsPollByIDRequest) GetLowerBound() *float64 {
	if o == nil {
		return nil
	}
	return o.LowerBound
}

func (o *GetSearchJobsResultsPollByIDRequest) GetUpperBound() *float64 {
	if o == nil {
		return nil
	}
	return o.UpperBound
}

func (o *GetSearchJobsResultsPollByIDRequest) GetLastJobStatus() *string {
	if o == nil {
		return nil
	}
	return o.LastJobStatus
}

type GetSearchJobsResultsPollByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// SearchResultsResults for the Search /results and /results-poll endpoints. object
	SearchJobResults *shared.SearchJobResults
	// Unexpected error
	Error *shared.Error
}

func (o *GetSearchJobsResultsPollByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSearchJobsResultsPollByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSearchJobsResultsPollByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSearchJobsResultsPollByIDResponse) GetSearchJobResults() *shared.SearchJobResults {
	if o == nil {
		return nil
	}
	return o.SearchJobResults
}

func (o *GetSearchJobsResultsPollByIDResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
