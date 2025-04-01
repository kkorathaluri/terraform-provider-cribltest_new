// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type InputGrafana1 struct {
	ActivityLogSampleRate types.Number                       `tfsdk:"activity_log_sample_rate"`
	CaptureHeaders        types.Bool                         `tfsdk:"capture_headers"`
	Connections           []InputGrafanaConnections          `tfsdk:"connections"`
	Description           types.String                       `tfsdk:"description"`
	Disabled              types.Bool                         `tfsdk:"disabled"`
	EnableHealthCheck     types.Bool                         `tfsdk:"enable_health_check"`
	EnableProxyHeader     types.Bool                         `tfsdk:"enable_proxy_header"`
	Environment           types.String                       `tfsdk:"environment"`
	Host                  types.String                       `tfsdk:"host"`
	ID                    types.String                       `tfsdk:"id"`
	IPAllowlistRegex      types.String                       `tfsdk:"ip_allowlist_regex"`
	IPDenylistRegex       types.String                       `tfsdk:"ip_denylist_regex"`
	KeepAliveTimeout      types.Number                       `tfsdk:"keep_alive_timeout"`
	LokiAPI               types.String                       `tfsdk:"loki_api"`
	LokiAuth              *LokiAuth                          `tfsdk:"loki_auth"`
	MaxActiveReq          types.Number                       `tfsdk:"max_active_req"`
	MaxRequestsPerSocket  types.Int64                        `tfsdk:"max_requests_per_socket"`
	Metadata              []InputGrafanaMetadata             `tfsdk:"metadata"`
	Pipeline              types.String                       `tfsdk:"pipeline"`
	Port                  types.Number                       `tfsdk:"port"`
	Pq                    *InputGrafanaPq                    `tfsdk:"pq"`
	PqEnabled             types.Bool                         `tfsdk:"pq_enabled"`
	PrometheusAPI         types.String                       `tfsdk:"prometheus_api"`
	PrometheusAuth        *PrometheusAuth                    `tfsdk:"prometheus_auth"`
	RequestTimeout        types.Number                       `tfsdk:"request_timeout"`
	SendToRoutes          types.Bool                         `tfsdk:"send_to_routes"`
	SocketTimeout         types.Number                       `tfsdk:"socket_timeout"`
	Streamtags            []types.String                     `tfsdk:"streamtags"`
	TLS                   *InputGrafanaTLSSettingsServerSide `tfsdk:"tls"`
	Type                  types.String                       `tfsdk:"type"`
}
