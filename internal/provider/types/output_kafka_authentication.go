// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type OutputKafkaAuthentication struct {
	Disabled  types.Bool   `tfsdk:"disabled"`
	Mechanism types.String `tfsdk:"mechanism"`
}
