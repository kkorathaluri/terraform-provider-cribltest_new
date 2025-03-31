// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type DeleteGlobalVariableLibVarsByPackAndIDRequest struct {
	// Unique ID to DELETE for pack
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// pack ID to DELETE
	Pack string `pathParam:"style=simple,explode=false,name=pack"`
}

func (o *DeleteGlobalVariableLibVarsByPackAndIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeleteGlobalVariableLibVarsByPackAndIDRequest) GetPack() string {
	if o == nil {
		return ""
	}
	return o.Pack
}

// DeleteGlobalVariableLibVarsByPackAndIDResponseBody - a list of Global Variable objects
type DeleteGlobalVariableLibVarsByPackAndIDResponseBody struct {
	// number of items present in the items array
	Count *int64             `json:"count,omitempty"`
	Items []shared.GlobalVar `json:"items,omitempty"`
}

func (o *DeleteGlobalVariableLibVarsByPackAndIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *DeleteGlobalVariableLibVarsByPackAndIDResponseBody) GetItems() []shared.GlobalVar {
	if o == nil {
		return nil
	}
	return o.Items
}

type DeleteGlobalVariableLibVarsByPackAndIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of Global Variable objects
	Object *DeleteGlobalVariableLibVarsByPackAndIDResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *DeleteGlobalVariableLibVarsByPackAndIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteGlobalVariableLibVarsByPackAndIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteGlobalVariableLibVarsByPackAndIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteGlobalVariableLibVarsByPackAndIDResponse) GetObject() *DeleteGlobalVariableLibVarsByPackAndIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *DeleteGlobalVariableLibVarsByPackAndIDResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
