// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type OutputDynatraceOtlp struct {
	AuthTokenName                 types.String                               `tfsdk:"auth_token_name"`
	Compress                      types.String                               `tfsdk:"compress"`
	Concurrency                   types.Number                               `tfsdk:"concurrency"`
	ConnectionTimeout             types.Number                               `tfsdk:"connection_timeout"`
	Description                   types.String                               `tfsdk:"description"`
	Endpoint                      types.String                               `tfsdk:"endpoint"`
	EndpointType                  types.String                               `tfsdk:"endpoint_type"`
	Environment                   types.String                               `tfsdk:"environment"`
	ExtraHTTPHeaders              []OutputDynatraceOtlpExtraHTTPHeaders      `tfsdk:"extra_http_headers"`
	FailedRequestLoggingMode      types.String                               `tfsdk:"failed_request_logging_mode"`
	FlushPeriodSec                types.Number                               `tfsdk:"flush_period_sec"`
	HTTPCompress                  types.String                               `tfsdk:"http_compress"`
	HTTPLogsEndpointOverride      types.String                               `tfsdk:"http_logs_endpoint_override"`
	HTTPMetricsEndpointOverride   types.String                               `tfsdk:"http_metrics_endpoint_override"`
	HTTPTracesEndpointOverride    types.String                               `tfsdk:"http_traces_endpoint_override"`
	ID                            types.String                               `tfsdk:"id"`
	KeepAlive                     types.Bool                                 `tfsdk:"keep_alive"`
	KeepAliveTime                 types.Number                               `tfsdk:"keep_alive_time"`
	MaxPayloadSizeKB              types.Number                               `tfsdk:"max_payload_size_kb"`
	Metadata                      []OutputDynatraceOtlpMetadata              `tfsdk:"metadata"`
	OnBackpressure                types.String                               `tfsdk:"on_backpressure"`
	OtlpVersion                   types.String                               `tfsdk:"otlp_version"`
	Pipeline                      types.String                               `tfsdk:"pipeline"`
	PqCompress                    types.String                               `tfsdk:"pq_compress"`
	PqControls                    *OutputDynatraceOtlpPqControls             `tfsdk:"pq_controls"`
	PqMaxFileSize                 types.String                               `tfsdk:"pq_max_file_size"`
	PqMaxSize                     types.String                               `tfsdk:"pq_max_size"`
	PqMode                        types.String                               `tfsdk:"pq_mode"`
	PqOnBackpressure              types.String                               `tfsdk:"pq_on_backpressure"`
	PqPath                        types.String                               `tfsdk:"pq_path"`
	Protocol                      types.String                               `tfsdk:"protocol"`
	RejectUnauthorized            types.Bool                                 `tfsdk:"reject_unauthorized"`
	ResponseHonorRetryAfterHeader types.Bool                                 `tfsdk:"response_honor_retry_after_header"`
	ResponseRetrySettings         []OutputDynatraceOtlpResponseRetrySettings `tfsdk:"response_retry_settings"`
	SafeHeaders                   []types.String                             `tfsdk:"safe_headers"`
	Streamtags                    []types.String                             `tfsdk:"streamtags"`
	SystemFields                  []types.String                             `tfsdk:"system_fields"`
	TimeoutRetrySettings          *OutputDynatraceOtlpTimeoutRetrySettings   `tfsdk:"timeout_retry_settings"`
	TimeoutSec                    types.Number                               `tfsdk:"timeout_sec"`
	TokenSecret                   types.String                               `tfsdk:"token_secret"`
	Type                          types.String                               `tfsdk:"type"`
	UseRoundRobinDNS              types.Bool                                 `tfsdk:"use_round_robin_dns"`
}
