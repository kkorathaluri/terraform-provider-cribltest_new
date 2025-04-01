// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type OutputNetflow struct {
	Description         types.String         `tfsdk:"description"`
	DNSResolvePeriodSec types.Number         `tfsdk:"dns_resolve_period_sec"`
	Environment         types.String         `tfsdk:"environment"`
	Hosts               []OutputNetflowHosts `tfsdk:"hosts"`
	ID                  types.String         `tfsdk:"id"`
	Pipeline            types.String         `tfsdk:"pipeline"`
	Streamtags          []types.String       `tfsdk:"streamtags"`
	SystemFields        []types.String       `tfsdk:"system_fields"`
	Type                types.String         `tfsdk:"type"`
}
