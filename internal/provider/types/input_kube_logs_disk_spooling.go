// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type InputKubeLogsDiskSpooling struct {
	Compress    types.String `tfsdk:"compress"`
	Enable      types.Bool   `tfsdk:"enable"`
	MaxDataSize types.String `tfsdk:"max_data_size"`
	MaxDataTime types.String `tfsdk:"max_data_time"`
	TimeWindow  types.String `tfsdk:"time_window"`
}
