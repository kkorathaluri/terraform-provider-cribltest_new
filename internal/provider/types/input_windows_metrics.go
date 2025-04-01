// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type InputWindowsMetrics struct {
	Connections         []InputWindowsMetricsConnections `tfsdk:"connections"`
	Description         types.String                     `tfsdk:"description"`
	Disabled            types.Bool                       `tfsdk:"disabled"`
	DisableNativeModule types.Bool                       `tfsdk:"disable_native_module"`
	Environment         types.String                     `tfsdk:"environment"`
	Host                *InputWindowsMetricsHost         `tfsdk:"host"`
	ID                  types.String                     `tfsdk:"id"`
	Interval            types.Number                     `tfsdk:"interval"`
	Metadata            []InputWindowsMetricsMetadata    `tfsdk:"metadata"`
	Persistence         *InputWindowsMetricsPersistence  `tfsdk:"persistence"`
	Pipeline            types.String                     `tfsdk:"pipeline"`
	Pq                  *InputWindowsMetricsPq           `tfsdk:"pq"`
	PqEnabled           types.Bool                       `tfsdk:"pq_enabled"`
	Process             *InputWindowsMetricsProcess      `tfsdk:"process"`
	SendToRoutes        types.Bool                       `tfsdk:"send_to_routes"`
	Streamtags          []types.String                   `tfsdk:"streamtags"`
	Type                types.String                     `tfsdk:"type"`
}
