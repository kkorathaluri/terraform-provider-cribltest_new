// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type OutputCrowdstrikeNextGenSiemResponseRetrySettings struct {
	BackoffRate    types.Number `tfsdk:"backoff_rate"`
	HTTPStatus     types.Number `tfsdk:"http_status"`
	InitialBackoff types.Number `tfsdk:"initial_backoff"`
	MaxBackoff     types.Number `tfsdk:"max_backoff"`
}
