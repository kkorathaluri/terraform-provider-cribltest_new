// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type OutputCriblTCP struct {
	Compression               types.String                         `tfsdk:"compression"`
	ConnectionTimeout         types.Number                         `tfsdk:"connection_timeout"`
	Description               types.String                         `tfsdk:"description"`
	DNSResolvePeriodSec       types.Number                         `tfsdk:"dns_resolve_period_sec"`
	Environment               types.String                         `tfsdk:"environment"`
	ExcludeFields             []types.String                       `tfsdk:"exclude_fields"`
	ExcludeSelf               types.Bool                           `tfsdk:"exclude_self"`
	Host                      types.String                         `tfsdk:"host"`
	Hosts                     []OutputCriblTCPHosts                `tfsdk:"hosts"`
	ID                        types.String                         `tfsdk:"id"`
	LoadBalanced              types.Bool                           `tfsdk:"load_balanced"`
	LoadBalanceStatsPeriodSec types.Number                         `tfsdk:"load_balance_stats_period_sec"`
	LogFailedRequests         types.Bool                           `tfsdk:"log_failed_requests"`
	MaxConcurrentSenders      types.Number                         `tfsdk:"max_concurrent_senders"`
	OnBackpressure            types.String                         `tfsdk:"on_backpressure"`
	Pipeline                  types.String                         `tfsdk:"pipeline"`
	Port                      types.Number                         `tfsdk:"port"`
	PqCompress                types.String                         `tfsdk:"pq_compress"`
	PqControls                *OutputCriblTCPPqControls            `tfsdk:"pq_controls"`
	PqMaxFileSize             types.String                         `tfsdk:"pq_max_file_size"`
	PqMaxSize                 types.String                         `tfsdk:"pq_max_size"`
	PqMode                    types.String                         `tfsdk:"pq_mode"`
	PqOnBackpressure          types.String                         `tfsdk:"pq_on_backpressure"`
	PqPath                    types.String                         `tfsdk:"pq_path"`
	Streamtags                []types.String                       `tfsdk:"streamtags"`
	SystemFields              []types.String                       `tfsdk:"system_fields"`
	ThrottleRatePerSec        types.String                         `tfsdk:"throttle_rate_per_sec"`
	TLS                       *OutputCriblTCPTLSSettingsClientSide `tfsdk:"tls"`
	TokenTTLMinutes           types.Number                         `tfsdk:"token_ttl_minutes"`
	Type                      types.String                         `tfsdk:"type"`
	WriteTimeout              types.Number                         `tfsdk:"write_timeout"`
}
