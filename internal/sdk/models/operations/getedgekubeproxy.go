// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type GetEdgeKubeProxyRequest struct {
	// string optional
	Path *string `queryParam:"style=form,explode=true,name=path"`
}

func (o *GetEdgeKubeProxyRequest) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

// GetEdgeKubeProxyResponseBody - a list of object objects
type GetEdgeKubeProxyResponseBody struct {
	// number of items present in the items array
	Count *int64   `json:"count,omitempty"`
	Items []string `json:"items,omitempty"`
}

func (o *GetEdgeKubeProxyResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *GetEdgeKubeProxyResponseBody) GetItems() []string {
	if o == nil {
		return nil
	}
	return o.Items
}

type GetEdgeKubeProxyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of object objects
	Object *GetEdgeKubeProxyResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *GetEdgeKubeProxyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEdgeKubeProxyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEdgeKubeProxyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetEdgeKubeProxyResponse) GetObject() *GetEdgeKubeProxyResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *GetEdgeKubeProxyResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
