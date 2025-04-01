// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type InputCriblTCP struct {
	Connections         []InputCriblTCPConnections          `tfsdk:"connections"`
	Description         types.String                        `tfsdk:"description"`
	Disabled            types.Bool                          `tfsdk:"disabled"`
	EnableLoadBalancing types.Bool                          `tfsdk:"enable_load_balancing"`
	EnableProxyHeader   types.Bool                          `tfsdk:"enable_proxy_header"`
	Environment         types.String                        `tfsdk:"environment"`
	Host                types.String                        `tfsdk:"host"`
	ID                  types.String                        `tfsdk:"id"`
	MaxActiveCxn        types.Number                        `tfsdk:"max_active_cxn"`
	Metadata            []InputCriblTCPMetadata             `tfsdk:"metadata"`
	Pipeline            types.String                        `tfsdk:"pipeline"`
	Port                types.Number                        `tfsdk:"port"`
	Pq                  *InputCriblTCPPq                    `tfsdk:"pq"`
	PqEnabled           types.Bool                          `tfsdk:"pq_enabled"`
	SendToRoutes        types.Bool                          `tfsdk:"send_to_routes"`
	SocketEndingMaxWait types.Number                        `tfsdk:"socket_ending_max_wait"`
	SocketIdleTimeout   types.Number                        `tfsdk:"socket_idle_timeout"`
	SocketMaxLifespan   types.Number                        `tfsdk:"socket_max_lifespan"`
	Streamtags          []types.String                      `tfsdk:"streamtags"`
	TLS                 *InputCriblTCPTLSSettingsServerSide `tfsdk:"tls"`
	Type                types.String                        `tfsdk:"type"`
}
