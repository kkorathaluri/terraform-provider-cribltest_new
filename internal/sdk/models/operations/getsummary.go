// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

// Mode - product filter
type Mode string

const (
	ModeWorker      Mode = "worker"
	ModeManagedEdge Mode = "managed-edge"
)

func (e Mode) ToPointer() *Mode {
	return &e
}
func (e *Mode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "worker":
		fallthrough
	case "managed-edge":
		*e = Mode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Mode: %v", v)
	}
}

type GetSummaryRequest struct {
	// product filter
	Mode *Mode `queryParam:"style=form,explode=true,name=mode"`
}

func (o *GetSummaryRequest) GetMode() *Mode {
	if o == nil {
		return nil
	}
	return o.Mode
}

// GetSummaryResponseBody - a list of DistributedSummary objects
type GetSummaryResponseBody struct {
	// number of items present in the items array
	Count *int64                      `json:"count,omitempty"`
	Items []shared.DistributedSummary `json:"items,omitempty"`
}

func (o *GetSummaryResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *GetSummaryResponseBody) GetItems() []shared.DistributedSummary {
	if o == nil {
		return nil
	}
	return o.Items
}

type GetSummaryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of DistributedSummary objects
	Object *GetSummaryResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *GetSummaryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSummaryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSummaryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSummaryResponse) GetObject() *GetSummaryResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *GetSummaryResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
