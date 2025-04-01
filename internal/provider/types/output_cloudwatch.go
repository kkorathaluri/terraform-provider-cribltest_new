// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type OutputCloudwatch struct {
	AssumeRoleArn           types.String                `tfsdk:"assume_role_arn"`
	AssumeRoleExternalID    types.String                `tfsdk:"assume_role_external_id"`
	AwsAPIKey               types.String                `tfsdk:"aws_api_key"`
	AwsAuthenticationMethod types.String                `tfsdk:"aws_authentication_method"`
	AwsSecret               types.String                `tfsdk:"aws_secret"`
	AwsSecretKey            types.String                `tfsdk:"aws_secret_key"`
	Description             types.String                `tfsdk:"description"`
	DurationSeconds         types.Number                `tfsdk:"duration_seconds"`
	EnableAssumeRole        types.Bool                  `tfsdk:"enable_assume_role"`
	Endpoint                types.String                `tfsdk:"endpoint"`
	Environment             types.String                `tfsdk:"environment"`
	FlushPeriodSec          types.Number                `tfsdk:"flush_period_sec"`
	ID                      types.String                `tfsdk:"id"`
	LogGroupName            types.String                `tfsdk:"log_group_name"`
	LogStreamName           types.String                `tfsdk:"log_stream_name"`
	MaxQueueSize            types.Number                `tfsdk:"max_queue_size"`
	MaxRecordSizeKB         types.Number                `tfsdk:"max_record_size_kb"`
	OnBackpressure          types.String                `tfsdk:"on_backpressure"`
	Pipeline                types.String                `tfsdk:"pipeline"`
	PqCompress              types.String                `tfsdk:"pq_compress"`
	PqControls              *OutputCloudwatchPqControls `tfsdk:"pq_controls"`
	PqMaxFileSize           types.String                `tfsdk:"pq_max_file_size"`
	PqMaxSize               types.String                `tfsdk:"pq_max_size"`
	PqMode                  types.String                `tfsdk:"pq_mode"`
	PqOnBackpressure        types.String                `tfsdk:"pq_on_backpressure"`
	PqPath                  types.String                `tfsdk:"pq_path"`
	Region                  types.String                `tfsdk:"region"`
	RejectUnauthorized      types.Bool                  `tfsdk:"reject_unauthorized"`
	ReuseConnections        types.Bool                  `tfsdk:"reuse_connections"`
	Streamtags              []types.String              `tfsdk:"streamtags"`
	SystemFields            []types.String              `tfsdk:"system_fields"`
	Type                    types.String                `tfsdk:"type"`
}
