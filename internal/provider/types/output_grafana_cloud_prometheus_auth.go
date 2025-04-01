// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type OutputGrafanaCloudPrometheusAuth struct {
	AuthType          types.String `tfsdk:"auth_type"`
	CredentialsSecret types.String `tfsdk:"credentials_secret"`
	Password          types.String `tfsdk:"password"`
	TextSecret        types.String `tfsdk:"text_secret"`
	Token             types.String `tfsdk:"token"`
	Username          types.String `tfsdk:"username"`
}
