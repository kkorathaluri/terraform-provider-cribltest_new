// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

// Context - Search query context
type Context string

const (
	ContextStream Context = "stream"
	ContextEdge   Context = "edge"
)

func (e Context) ToPointer() *Context {
	return &e
}
func (e *Context) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "stream":
		fallthrough
	case "edge":
		*e = Context(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Context: %v", v)
	}
}

type GetCluiRequest struct {
	// Search query
	Query string `queryParam:"style=form,explode=true,name=query"`
	// Search query context
	Context *Context `queryParam:"style=form,explode=true,name=context"`
}

func (o *GetCluiRequest) GetQuery() string {
	if o == nil {
		return ""
	}
	return o.Query
}

func (o *GetCluiRequest) GetContext() *Context {
	if o == nil {
		return nil
	}
	return o.Context
}

// GetCluiResponseBody - a list of CluiItem objects
type GetCluiResponseBody struct {
	// number of items present in the items array
	Count *int64            `json:"count,omitempty"`
	Items []shared.CluiItem `json:"items,omitempty"`
}

func (o *GetCluiResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *GetCluiResponseBody) GetItems() []shared.CluiItem {
	if o == nil {
		return nil
	}
	return o.Items
}

type GetCluiResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of CluiItem objects
	Object *GetCluiResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *GetCluiResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCluiResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCluiResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetCluiResponse) GetObject() *GetCluiResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *GetCluiResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
