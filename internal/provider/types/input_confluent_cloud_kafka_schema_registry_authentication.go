// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type InputConfluentCloudKafkaSchemaRegistryAuthentication struct {
	Auth              *InputConfluentCloudAuth                                     `tfsdk:"auth"`
	ConnectionTimeout types.Number                                                 `tfsdk:"connection_timeout"`
	Disabled          types.Bool                                                   `tfsdk:"disabled"`
	MaxRetries        types.Number                                                 `tfsdk:"max_retries"`
	RequestTimeout    types.Number                                                 `tfsdk:"request_timeout"`
	SchemaRegistryURL types.String                                                 `tfsdk:"schema_registry_url"`
	TLS               *InputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide `tfsdk:"tls"`
}
