// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

type SchemaLibEntry struct {
	ID          string  `json:"id"`
	Description *string `json:"description,omitempty"`
	// JSON schema matching standards of draft version 2019-09
	Schema               string `json:"schema"`
	AdditionalProperties any    `additionalProperties:"true" json:"-"`
}

func (s SchemaLibEntry) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SchemaLibEntry) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SchemaLibEntry) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *SchemaLibEntry) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *SchemaLibEntry) GetSchema() string {
	if o == nil {
		return ""
	}
	return o.Schema
}

func (o *SchemaLibEntry) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}
