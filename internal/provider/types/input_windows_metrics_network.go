// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type InputWindowsMetricsNetwork struct {
	Detail       types.Bool     `tfsdk:"detail"`
	Devices      []types.String `tfsdk:"devices"`
	Mode         types.String   `tfsdk:"mode"`
	PerInterface types.Bool     `tfsdk:"per_interface"`
}
