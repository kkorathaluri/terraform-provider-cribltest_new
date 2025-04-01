// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type OutputSignalfx struct {
	AuthType                      types.String                          `tfsdk:"auth_type"`
	Compress                      types.Bool                            `tfsdk:"compress"`
	Concurrency                   types.Number                          `tfsdk:"concurrency"`
	Description                   types.String                          `tfsdk:"description"`
	Environment                   types.String                          `tfsdk:"environment"`
	ExtraHTTPHeaders              []OutputSignalfxExtraHTTPHeaders      `tfsdk:"extra_http_headers"`
	FailedRequestLoggingMode      types.String                          `tfsdk:"failed_request_logging_mode"`
	FlushPeriodSec                types.Number                          `tfsdk:"flush_period_sec"`
	ID                            types.String                          `tfsdk:"id"`
	MaxPayloadEvents              types.Number                          `tfsdk:"max_payload_events"`
	MaxPayloadSizeKB              types.Number                          `tfsdk:"max_payload_size_kb"`
	OnBackpressure                types.String                          `tfsdk:"on_backpressure"`
	Pipeline                      types.String                          `tfsdk:"pipeline"`
	PqCompress                    types.String                          `tfsdk:"pq_compress"`
	PqControls                    *OutputSignalfxPqControls             `tfsdk:"pq_controls"`
	PqMaxFileSize                 types.String                          `tfsdk:"pq_max_file_size"`
	PqMaxSize                     types.String                          `tfsdk:"pq_max_size"`
	PqMode                        types.String                          `tfsdk:"pq_mode"`
	PqOnBackpressure              types.String                          `tfsdk:"pq_on_backpressure"`
	PqPath                        types.String                          `tfsdk:"pq_path"`
	Realm                         types.String                          `tfsdk:"realm"`
	RejectUnauthorized            types.Bool                            `tfsdk:"reject_unauthorized"`
	ResponseHonorRetryAfterHeader types.Bool                            `tfsdk:"response_honor_retry_after_header"`
	ResponseRetrySettings         []OutputSignalfxResponseRetrySettings `tfsdk:"response_retry_settings"`
	SafeHeaders                   []types.String                        `tfsdk:"safe_headers"`
	Streamtags                    []types.String                        `tfsdk:"streamtags"`
	SystemFields                  []types.String                        `tfsdk:"system_fields"`
	TextSecret                    types.String                          `tfsdk:"text_secret"`
	TimeoutRetrySettings          *OutputSignalfxTimeoutRetrySettings   `tfsdk:"timeout_retry_settings"`
	TimeoutSec                    types.Number                          `tfsdk:"timeout_sec"`
	Token                         types.String                          `tfsdk:"token"`
	Type                          types.String                          `tfsdk:"type"`
	UseRoundRobinDNS              types.Bool                            `tfsdk:"use_round_robin_dns"`
}
