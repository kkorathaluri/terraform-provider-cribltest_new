// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type InputWiz struct {
	AuthAudienceOverride types.String            `tfsdk:"auth_audience_override"`
	AuthType             types.String            `tfsdk:"auth_type"`
	AuthURL              types.String            `tfsdk:"auth_url"`
	ClientID             types.String            `tfsdk:"client_id"`
	ClientSecret         types.String            `tfsdk:"client_secret"`
	Connections          []InputWizConnections   `tfsdk:"connections"`
	ContentConfig        []InputWizContentConfig `tfsdk:"content_config"`
	Description          types.String            `tfsdk:"description"`
	Disabled             types.Bool              `tfsdk:"disabled"`
	Endpoint             types.String            `tfsdk:"endpoint"`
	Environment          types.String            `tfsdk:"environment"`
	ID                   types.String            `tfsdk:"id"`
	KeepAliveTime        types.Number            `tfsdk:"keep_alive_time"`
	MaxMissedKeepAlives  types.Number            `tfsdk:"max_missed_keep_alives"`
	Metadata             []InputWizMetadata      `tfsdk:"metadata"`
	Pipeline             types.String            `tfsdk:"pipeline"`
	Pq                   *InputWizPq             `tfsdk:"pq"`
	PqEnabled            types.Bool              `tfsdk:"pq_enabled"`
	RequestTimeout       types.Number            `tfsdk:"request_timeout"`
	RetryRules           *InputWizRetryRules     `tfsdk:"retry_rules"`
	SendToRoutes         types.Bool              `tfsdk:"send_to_routes"`
	Streamtags           []types.String          `tfsdk:"streamtags"`
	TextSecret           types.String            `tfsdk:"text_secret"`
	TTL                  types.String            `tfsdk:"ttl"`
	Type                 types.String            `tfsdk:"type"`
}
