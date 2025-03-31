// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/models/shared"
	"net/http"
)

type CreateCriblLakeDatasetByLakeIDRequest struct {
	// lake id that contains the Datasets
	LakeID string `pathParam:"style=simple,explode=false,name=lakeId"`
	// CriblLakeDataset object
	CriblLakeDataset shared.CriblLakeDataset `request:"mediaType=application/json"`
}

func (o *CreateCriblLakeDatasetByLakeIDRequest) GetLakeID() string {
	if o == nil {
		return ""
	}
	return o.LakeID
}

func (o *CreateCriblLakeDatasetByLakeIDRequest) GetCriblLakeDataset() shared.CriblLakeDataset {
	if o == nil {
		return shared.CriblLakeDataset{}
	}
	return o.CriblLakeDataset
}

// CreateCriblLakeDatasetByLakeIDResponseBody - a list of CriblLakeDataset objects
type CreateCriblLakeDatasetByLakeIDResponseBody struct {
	// number of items present in the items array
	Count *int64                    `json:"count,omitempty"`
	Items []shared.CriblLakeDataset `json:"items,omitempty"`
}

func (o *CreateCriblLakeDatasetByLakeIDResponseBody) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *CreateCriblLakeDatasetByLakeIDResponseBody) GetItems() []shared.CriblLakeDataset {
	if o == nil {
		return nil
	}
	return o.Items
}

type CreateCriblLakeDatasetByLakeIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// a list of CriblLakeDataset objects
	Object *CreateCriblLakeDatasetByLakeIDResponseBody
	// Unexpected error
	Error *shared.Error
}

func (o *CreateCriblLakeDatasetByLakeIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateCriblLakeDatasetByLakeIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateCriblLakeDatasetByLakeIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateCriblLakeDatasetByLakeIDResponse) GetObject() *CreateCriblLakeDatasetByLakeIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *CreateCriblLakeDatasetByLakeIDResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}
