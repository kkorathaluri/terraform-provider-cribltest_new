// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type InputOffice365MsgTrace struct {
	AuthType               types.String                        `tfsdk:"auth_type"`
	CertOptions            *CertOptions                        `tfsdk:"cert_options"`
	ClientID               types.String                        `tfsdk:"client_id"`
	ClientSecret           types.String                        `tfsdk:"client_secret"`
	Connections            []InputOffice365MsgTraceConnections `tfsdk:"connections"`
	CredentialsSecret      types.String                        `tfsdk:"credentials_secret"`
	Description            types.String                        `tfsdk:"description"`
	Disabled               types.Bool                          `tfsdk:"disabled"`
	DisableTimeFilter      types.Bool                          `tfsdk:"disable_time_filter"`
	EndDate                types.String                        `tfsdk:"end_date"`
	Environment            types.String                        `tfsdk:"environment"`
	ID                     types.String                        `tfsdk:"id"`
	Interval               types.Number                        `tfsdk:"interval"`
	JobTimeout             types.String                        `tfsdk:"job_timeout"`
	KeepAliveTime          types.Number                        `tfsdk:"keep_alive_time"`
	LogLevel               types.String                        `tfsdk:"log_level"`
	MaxMissedKeepAlives    types.Number                        `tfsdk:"max_missed_keep_alives"`
	MaxTaskReschedule      types.Number                        `tfsdk:"max_task_reschedule"`
	Metadata               []InputOffice365MsgTraceMetadata    `tfsdk:"metadata"`
	Password               types.String                        `tfsdk:"password"`
	Pipeline               types.String                        `tfsdk:"pipeline"`
	PlanType               types.String                        `tfsdk:"plan_type"`
	Pq                     *InputOffice365MsgTracePq           `tfsdk:"pq"`
	PqEnabled              types.Bool                          `tfsdk:"pq_enabled"`
	RescheduleDroppedTasks types.Bool                          `tfsdk:"reschedule_dropped_tasks"`
	Resource               types.String                        `tfsdk:"resource"`
	RetryRules             *InputOffice365MsgTraceRetryRules   `tfsdk:"retry_rules"`
	SendToRoutes           types.Bool                          `tfsdk:"send_to_routes"`
	StartDate              types.String                        `tfsdk:"start_date"`
	Streamtags             []types.String                      `tfsdk:"streamtags"`
	TenantID               types.String                        `tfsdk:"tenant_id"`
	TextSecret             types.String                        `tfsdk:"text_secret"`
	Timeout                types.Number                        `tfsdk:"timeout"`
	TTL                    types.String                        `tfsdk:"ttl"`
	Type                   types.String                        `tfsdk:"type"`
	URL                    types.String                        `tfsdk:"url"`
	Username               types.String                        `tfsdk:"username"`
}
