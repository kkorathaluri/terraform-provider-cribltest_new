// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type OutputCriblTCPHosts struct {
	Host       types.String `tfsdk:"host"`
	Port       types.Number `tfsdk:"port"`
	Servername types.String `tfsdk:"servername"`
	TLS        types.String `tfsdk:"tls"`
	Weight     types.Number `tfsdk:"weight"`
}
