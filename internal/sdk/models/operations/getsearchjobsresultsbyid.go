// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type GetSearchJobsResultsByIDRequest struct {
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
}

func (o *GetSearchJobsResultsByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetSearchJobsResultsByIDRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetSearchJobsResultsByIDRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetSearchJobsResultsByIDRequest) GetLowerBound() *float64 {
	if o == nil {
		return nil
	}
	return o.LowerBound
}

func (o *GetSearchJobsResultsByIDRequest) GetUpperBound() *float64 {
	if o == nil {
		return nil
	}
	return o.UpperBound
}

type GetSearchJobsResultsByIDResponse struct {
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

func (o *GetSearchJobsResultsByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSearchJobsResultsByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSearchJobsResultsByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSearchJobsResultsByIDResponse) GetSearchJobResults() *shared.SearchJobResults {
	if o == nil {
		return nil
	}
	return o.SearchJobResults
}

func (o *GetSearchJobsResultsByIDResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
