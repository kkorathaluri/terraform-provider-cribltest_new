// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type OutputMskAuth struct {
	CredentialsSecret types.String `tfsdk:"credentials_secret"`
	Disabled          types.Bool   `tfsdk:"disabled"`
}
