// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type InputOffice365ServiceRetryRules struct {
	Codes               []types.Number `tfsdk:"codes"`
	EnableHeader        types.Bool     `tfsdk:"enable_header"`
	Interval            types.Number   `tfsdk:"interval"`
	Limit               types.Number   `tfsdk:"limit"`
	Multiplier          types.Number   `tfsdk:"multiplier"`
	RetryConnectReset   types.Bool     `tfsdk:"retry_connect_reset"`
	RetryConnectTimeout types.Bool     `tfsdk:"retry_connect_timeout"`
	Type                types.String   `tfsdk:"type"`
}
