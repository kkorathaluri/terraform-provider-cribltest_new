// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type InputWindowsMetricsCPU struct {
	Detail types.Bool   `tfsdk:"detail"`
	Mode   types.String `tfsdk:"mode"`
	PerCPU types.Bool   `tfsdk:"per_cpu"`
	Time   types.Bool   `tfsdk:"time"`
}
