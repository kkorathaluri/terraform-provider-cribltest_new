// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type InputCriblTCPTLSSettingsServerSide struct {
	CaPath             types.String `tfsdk:"ca_path"`
	CertificateName    types.String `tfsdk:"certificate_name"`
	CertPath           types.String `tfsdk:"cert_path"`
	CommonNameRegex    types.String `tfsdk:"common_name_regex"`
	Disabled           types.Bool   `tfsdk:"disabled"`
	MaxVersion         types.String `tfsdk:"max_version"`
	MinVersion         types.String `tfsdk:"min_version"`
	Passphrase         types.String `tfsdk:"passphrase"`
	PrivKeyPath        types.String `tfsdk:"priv_key_path"`
	RejectUnauthorized types.String `tfsdk:"reject_unauthorized"`
	RequestCert        types.Bool   `tfsdk:"request_cert"`
}
