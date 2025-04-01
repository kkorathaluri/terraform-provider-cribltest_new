// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type OutputNewrelicTimeoutRetrySettings struct {
	BackoffRate    types.Number `tfsdk:"backoff_rate"`
	InitialBackoff types.Number `tfsdk:"initial_backoff"`
	MaxBackoff     types.Number `tfsdk:"max_backoff"`
	TimeoutRetry   types.Bool   `tfsdk:"timeout_retry"`
}
