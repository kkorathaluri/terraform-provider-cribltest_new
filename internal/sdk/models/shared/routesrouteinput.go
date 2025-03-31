// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

type RoutesRouteInput struct {
	Name string `json:"name"`
	// Disable this routing rule
	Disabled *bool `json:"disabled,omitempty"`
	// JavaScript expression to select data to route
	Filter *string `default:"true" json:"filter"`
	// Pipeline to send the matching data to
	Pipeline string `json:"pipeline"`
	// Enable to use a JavaScript expression that evaluates to the name of the Output below
	EnableOutputExpression *bool   `default:"false" json:"enableOutputExpression"`
	Output                 any     `json:"output,omitempty"`
	OutputExpression       any     `json:"outputExpression,omitempty"`
	Description            *string `json:"description,omitempty"`
	// Flag to control whether the event gets consumed by this Route (Final), or cloned into it
	Final                *bool `default:"true" json:"final"`
	AdditionalProperties any   `additionalProperties:"true" json:"-"`
}

func (r RoutesRouteInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RoutesRouteInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RoutesRouteInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *RoutesRouteInput) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *RoutesRouteInput) GetFilter() *string {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *RoutesRouteInput) GetPipeline() string {
	if o == nil {
		return ""
	}
	return o.Pipeline
}

func (o *RoutesRouteInput) GetEnableOutputExpression() *bool {
	if o == nil {
		return nil
	}
	return o.EnableOutputExpression
}

func (o *RoutesRouteInput) GetOutput() any {
	if o == nil {
		return nil
	}
	return o.Output
}

func (o *RoutesRouteInput) GetOutputExpression() any {
	if o == nil {
		return nil
	}
	return o.OutputExpression
}

func (o *RoutesRouteInput) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *RoutesRouteInput) GetFinal() *bool {
	if o == nil {
		return nil
	}
	return o.Final
}

func (o *RoutesRouteInput) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}
