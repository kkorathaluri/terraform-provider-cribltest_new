// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

type OutputPrometheusType string

const (
	OutputPrometheusTypePrometheus OutputPrometheusType = "prometheus"
)

func (e OutputPrometheusType) ToPointer() *OutputPrometheusType {
	return &e
}
func (e *OutputPrometheusType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "prometheus":
		*e = OutputPrometheusType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputPrometheusType: %v", v)
	}
}

type OutputPrometheusExtraHTTPHeaders struct {
	// Field name
	Name *string `json:"name,omitempty"`
	// Field value
	Value string `json:"value"`
}

func (o *OutputPrometheusExtraHTTPHeaders) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *OutputPrometheusExtraHTTPHeaders) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// OutputPrometheusFailedRequestLoggingMode - Data to log when a request fails. All headers are redacted by default, unless listed as safe headers below.
type OutputPrometheusFailedRequestLoggingMode string

const (
	OutputPrometheusFailedRequestLoggingModePayload           OutputPrometheusFailedRequestLoggingMode = "payload"
	OutputPrometheusFailedRequestLoggingModePayloadAndHeaders OutputPrometheusFailedRequestLoggingMode = "payloadAndHeaders"
	OutputPrometheusFailedRequestLoggingModeNone              OutputPrometheusFailedRequestLoggingMode = "none"
)

func (e OutputPrometheusFailedRequestLoggingMode) ToPointer() *OutputPrometheusFailedRequestLoggingMode {
	return &e
}
func (e *OutputPrometheusFailedRequestLoggingMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "payload":
		fallthrough
	case "payloadAndHeaders":
		fallthrough
	case "none":
		*e = OutputPrometheusFailedRequestLoggingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputPrometheusFailedRequestLoggingMode: %v", v)
	}
}

type OutputPrometheusResponseRetrySettings struct {
	// The HTTP response status code that will trigger retries
	HTTPStatus float64 `json:"httpStatus"`
	// How long, in milliseconds, Cribl Stream should wait before initiating backoff. Maximum interval is 600,000 ms (10 minutes).
	InitialBackoff *float64 `default:"1000" json:"initialBackoff"`
	// Base for exponential backoff. A value of 2 (default) means Cribl Stream will retry after 2 seconds, then 4 seconds, then 8 seconds, etc.
	BackoffRate *float64 `default:"2" json:"backoffRate"`
	// The maximum backoff interval, in milliseconds, Cribl Stream should apply. Default (and minimum) is 10,000 ms (10 seconds); maximum is 180,000 ms (180 seconds).
	MaxBackoff *float64 `default:"10000" json:"maxBackoff"`
}

func (o OutputPrometheusResponseRetrySettings) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputPrometheusResponseRetrySettings) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputPrometheusResponseRetrySettings) GetHTTPStatus() float64 {
	if o == nil {
		return 0.0
	}
	return o.HTTPStatus
}

func (o *OutputPrometheusResponseRetrySettings) GetInitialBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialBackoff
}

func (o *OutputPrometheusResponseRetrySettings) GetBackoffRate() *float64 {
	if o == nil {
		return nil
	}
	return o.BackoffRate
}

func (o *OutputPrometheusResponseRetrySettings) GetMaxBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBackoff
}

type OutputPrometheusTimeoutRetrySettings struct {
	// Enable to retry on request timeout
	TimeoutRetry *bool `default:"false" json:"timeoutRetry"`
	// How long, in milliseconds, Cribl Stream should wait before initiating backoff. Maximum interval is 600,000 ms (10 minutes).
	InitialBackoff *float64 `default:"1000" json:"initialBackoff"`
	// Base for exponential backoff. A value of 2 (default) means Cribl Stream will retry after 2 seconds, then 4 seconds, then 8 seconds, etc.
	BackoffRate *float64 `default:"2" json:"backoffRate"`
	// The maximum backoff interval, in milliseconds, Cribl Stream should apply. Default (and minimum) is 10,000 ms (10 seconds); maximum is 180,000 ms (180 seconds).
	MaxBackoff *float64 `default:"10000" json:"maxBackoff"`
}

func (o OutputPrometheusTimeoutRetrySettings) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputPrometheusTimeoutRetrySettings) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputPrometheusTimeoutRetrySettings) GetTimeoutRetry() *bool {
	if o == nil {
		return nil
	}
	return o.TimeoutRetry
}

func (o *OutputPrometheusTimeoutRetrySettings) GetInitialBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialBackoff
}

func (o *OutputPrometheusTimeoutRetrySettings) GetBackoffRate() *float64 {
	if o == nil {
		return nil
	}
	return o.BackoffRate
}

func (o *OutputPrometheusTimeoutRetrySettings) GetMaxBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBackoff
}

// OutputPrometheusBackpressureBehavior - Whether to block, drop, or queue events when all receivers are exerting backpressure.
type OutputPrometheusBackpressureBehavior string

const (
	OutputPrometheusBackpressureBehaviorBlock OutputPrometheusBackpressureBehavior = "block"
	OutputPrometheusBackpressureBehaviorDrop  OutputPrometheusBackpressureBehavior = "drop"
	OutputPrometheusBackpressureBehaviorQueue OutputPrometheusBackpressureBehavior = "queue"
)

func (e OutputPrometheusBackpressureBehavior) ToPointer() *OutputPrometheusBackpressureBehavior {
	return &e
}
func (e *OutputPrometheusBackpressureBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		fallthrough
	case "queue":
		*e = OutputPrometheusBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputPrometheusBackpressureBehavior: %v", v)
	}
}

// OutputPrometheusAuthenticationType - Remote Write authentication type
type OutputPrometheusAuthenticationType string

const (
	OutputPrometheusAuthenticationTypeNone              OutputPrometheusAuthenticationType = "none"
	OutputPrometheusAuthenticationTypeBasic             OutputPrometheusAuthenticationType = "basic"
	OutputPrometheusAuthenticationTypeCredentialsSecret OutputPrometheusAuthenticationType = "credentialsSecret"
	OutputPrometheusAuthenticationTypeToken             OutputPrometheusAuthenticationType = "token"
	OutputPrometheusAuthenticationTypeTextSecret        OutputPrometheusAuthenticationType = "textSecret"
	OutputPrometheusAuthenticationTypeOauth             OutputPrometheusAuthenticationType = "oauth"
)

func (e OutputPrometheusAuthenticationType) ToPointer() *OutputPrometheusAuthenticationType {
	return &e
}
func (e *OutputPrometheusAuthenticationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "basic":
		fallthrough
	case "credentialsSecret":
		fallthrough
	case "token":
		fallthrough
	case "textSecret":
		fallthrough
	case "oauth":
		*e = OutputPrometheusAuthenticationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputPrometheusAuthenticationType: %v", v)
	}
}

// OutputPrometheusCompression - Codec to use to compress the persisted data.
type OutputPrometheusCompression string

const (
	OutputPrometheusCompressionNone OutputPrometheusCompression = "none"
	OutputPrometheusCompressionGzip OutputPrometheusCompression = "gzip"
)

func (e OutputPrometheusCompression) ToPointer() *OutputPrometheusCompression {
	return &e
}
func (e *OutputPrometheusCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputPrometheusCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputPrometheusCompression: %v", v)
	}
}

// OutputPrometheusQueueFullBehavior - Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputPrometheusQueueFullBehavior string

const (
	OutputPrometheusQueueFullBehaviorBlock OutputPrometheusQueueFullBehavior = "block"
	OutputPrometheusQueueFullBehaviorDrop  OutputPrometheusQueueFullBehavior = "drop"
)

func (e OutputPrometheusQueueFullBehavior) ToPointer() *OutputPrometheusQueueFullBehavior {
	return &e
}
func (e *OutputPrometheusQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputPrometheusQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputPrometheusQueueFullBehavior: %v", v)
	}
}

// OutputPrometheusMode - In Error mode, PQ writes events to the filesystem only when it detects a non-retryable Destination error. In Backpressure mode, PQ writes events to the filesystem when it detects backpressure from the Destination or when there are non-retryable Destination errors. In Always On mode, PQ always writes events to the filesystem.
type OutputPrometheusMode string

const (
	OutputPrometheusModeError        OutputPrometheusMode = "error"
	OutputPrometheusModeBackpressure OutputPrometheusMode = "backpressure"
	OutputPrometheusModeAlways       OutputPrometheusMode = "always"
)

func (e OutputPrometheusMode) ToPointer() *OutputPrometheusMode {
	return &e
}
func (e *OutputPrometheusMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "error":
		fallthrough
	case "backpressure":
		fallthrough
	case "always":
		*e = OutputPrometheusMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputPrometheusMode: %v", v)
	}
}

type OutputPrometheusPqControls struct {
}

type OutputPrometheusOauthParams struct {
	// OAuth parameter name
	Name string `json:"name"`
	// OAuth parameter value
	Value string `json:"value"`
}

func (o *OutputPrometheusOauthParams) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *OutputPrometheusOauthParams) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type OutputPrometheusOauthHeaders struct {
	// OAuth header name
	Name string `json:"name"`
	// OAuth header value
	Value string `json:"value"`
}

func (o *OutputPrometheusOauthHeaders) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *OutputPrometheusOauthHeaders) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type OutputPrometheus struct {
	// Unique ID for this output
	ID   *string              `json:"id,omitempty"`
	Type OutputPrometheusType `json:"type"`
	// Pipeline to process data before sending out to this output
	Pipeline *string `json:"pipeline,omitempty"`
	// Fields to automatically add to events, such as cribl_pipe. Supports wildcards. These fields are added as dimensions to generated metrics.
	SystemFields []string `json:"systemFields,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Tags for filtering and grouping in @{product}
	Streamtags []string `json:"streamtags,omitempty"`
	// The endpoint to send metrics to.
	URL string `json:"url"`
	// A JS expression that can be used to rename metrics. E.g.: name.replace(/\./g, '_') will replace all '.' characters in a metric's name with the supported '_' character. Use the 'name' global variable to access the metric's name.  You can access event fields' values via __e.<fieldName>.
	MetricRenameExpr *string `default:"name.replace(/[^a-zA-Z0-9_]/g, '_')" json:"metricRenameExpr"`
	// Whether to generate and send metadata (`type` and `metricFamilyName`) requests.
	SendMetadata *bool `default:"true" json:"sendMetadata"`
	// Maximum number of ongoing requests before blocking
	Concurrency *float64 `default:"5" json:"concurrency"`
	// Maximum size, in KB, of the request body
	MaxPayloadSizeKB *float64 `default:"4096" json:"maxPayloadSizeKB"`
	// Maximum number of events to include in the request body. Default is 0 (unlimited).
	MaxPayloadEvents *float64 `default:"0" json:"maxPayloadEvents"`
	// Reject certificates not authorized by a CA in the CA certificate path or by another trusted CA (such as the system's).
	//         Defaults to Yes. When this setting is also present in TLS Settings (Client Side),
	//         that value will take precedence.
	RejectUnauthorized *bool `default:"true" json:"rejectUnauthorized"`
	// Amount of time, in seconds, to wait for a request to complete before canceling it
	TimeoutSec *float64 `default:"30" json:"timeoutSec"`
	// Maximum time between requests. Small values could cause the payload size to be smaller than the configured Max body size.
	FlushPeriodSec *float64 `default:"1" json:"flushPeriodSec"`
	// Headers to add to all events.
	ExtraHTTPHeaders []OutputPrometheusExtraHTTPHeaders `json:"extraHttpHeaders,omitempty"`
	// Enables round-robin DNS lookup. When a DNS server returns multiple addresses, @{product} will cycle through them in the order returned. For optimal performance, consider enabling this setting for non-load balanced destinations.
	UseRoundRobinDNS *bool `default:"false" json:"useRoundRobinDns"`
	// Data to log when a request fails. All headers are redacted by default, unless listed as safe headers below.
	FailedRequestLoggingMode *OutputPrometheusFailedRequestLoggingMode `default:"none" json:"failedRequestLoggingMode"`
	// List of headers that are safe to log in plain text
	SafeHeaders []string `json:"safeHeaders,omitempty"`
	// Automatically retry after unsuccessful response status codes, such as 429 (Too Many Requests) or 503 (Service Unavailable).
	ResponseRetrySettings []OutputPrometheusResponseRetrySettings `json:"responseRetrySettings,omitempty"`
	TimeoutRetrySettings  *OutputPrometheusTimeoutRetrySettings   `json:"timeoutRetrySettings,omitempty"`
	// Honor any Retry-After header that specifies a delay (in seconds) no longer than 180 seconds after the retry request. @{product} limits the delay to 180 seconds, even if the Retry-After header specifies a longer delay. When enabled, takes precedence over user-configured retry options. When disabled, all Retry-After headers are ignored.
	ResponseHonorRetryAfterHeader *bool `default:"false" json:"responseHonorRetryAfterHeader"`
	// Whether to block, drop, or queue events when all receivers are exerting backpressure.
	OnBackpressure *OutputPrometheusBackpressureBehavior `default:"block" json:"onBackpressure"`
	// Remote Write authentication type
	AuthType    *OutputPrometheusAuthenticationType `default:"none" json:"authType"`
	Description *string                             `json:"description,omitempty"`
	// How frequently metrics metadata is sent out. Value cannot be smaller than the base Flush period (sec) set above.
	MetricsFlushPeriodSec *float64 `default:"60" json:"metricsFlushPeriodSec"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	PqMaxFileSize *string `default:"1 MB" json:"pqMaxFileSize"`
	// The maximum disk space that the queue can consume (as an average per Worker Process) before queueing stops. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `default:"5GB" json:"pqMaxSize"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/<output-id>.
	PqPath *string `default:"\\$CRIBL_HOME/state/queues" json:"pqPath"`
	// Codec to use to compress the persisted data.
	PqCompress *OutputPrometheusCompression `default:"none" json:"pqCompress"`
	// Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputPrometheusQueueFullBehavior `default:"block" json:"pqOnBackpressure"`
	// In Error mode, PQ writes events to the filesystem only when it detects a non-retryable Destination error. In Backpressure mode, PQ writes events to the filesystem when it detects backpressure from the Destination or when there are non-retryable Destination errors. In Always On mode, PQ always writes events to the filesystem.
	PqMode     *OutputPrometheusMode       `default:"error" json:"pqMode"`
	PqControls *OutputPrometheusPqControls `json:"pqControls,omitempty"`
	// Username for Basic authentication
	Username *string `json:"username,omitempty"`
	// Password for Basic authentication
	Password *string `json:"password,omitempty"`
	// Bearer token to include in the authorization header
	Token *string `json:"token,omitempty"`
	// Select or create a secret that references your credentials
	CredentialsSecret *string `json:"credentialsSecret,omitempty"`
	// Select or create a stored text secret
	TextSecret *string `json:"textSecret,omitempty"`
	// URL for OAuth
	LoginURL *string `json:"loginUrl,omitempty"`
	// Secret parameter name to pass in request body
	SecretParamName *string `json:"secretParamName,omitempty"`
	// Secret parameter value to pass in request body
	Secret *string `json:"secret,omitempty"`
	// Name of the auth token attribute in the OAuth response. Can be top-level (e.g., 'token'); or nested, using a period (e.g., 'data.token').
	TokenAttributeName *string `json:"tokenAttributeName,omitempty"`
	// JavaScript expression to compute the Authorization header value to pass in requests. The value `${token}` is used to reference the token obtained from authentication, e.g.: `Bearer ${token}`.
	AuthHeaderExpr *string `default:"Bearer \\${token}" json:"authHeaderExpr"`
	// How often the OAuth token should be refreshed.
	TokenTimeoutSecs *float64 `default:"3600" json:"tokenTimeoutSecs"`
	// Additional parameters to send in the OAuth login request. @{product} will combine the secret with these parameters, and will send the URL-encoded result in a POST request to the endpoint specified in the 'Login URL'. We'll automatically add the content-type header 'application/x-www-form-urlencoded' when sending this request.
	OauthParams []OutputPrometheusOauthParams `json:"oauthParams,omitempty"`
	// Additional headers to send in the OAuth login request. @{product} will automatically add the content-type header 'application/x-www-form-urlencoded' when sending this request.
	OauthHeaders []OutputPrometheusOauthHeaders `json:"oauthHeaders,omitempty"`
}

func (o OutputPrometheus) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputPrometheus) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *OutputPrometheus) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OutputPrometheus) GetType() OutputPrometheusType {
	if o == nil {
		return OutputPrometheusType("")
	}
	return o.Type
}

func (o *OutputPrometheus) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *OutputPrometheus) GetSystemFields() []string {
	if o == nil {
		return nil
	}
	return o.SystemFields
}

func (o *OutputPrometheus) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *OutputPrometheus) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *OutputPrometheus) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *OutputPrometheus) GetMetricRenameExpr() *string {
	if o == nil {
		return nil
	}
	return o.MetricRenameExpr
}

func (o *OutputPrometheus) GetSendMetadata() *bool {
	if o == nil {
		return nil
	}
	return o.SendMetadata
}

func (o *OutputPrometheus) GetConcurrency() *float64 {
	if o == nil {
		return nil
	}
	return o.Concurrency
}

func (o *OutputPrometheus) GetMaxPayloadSizeKB() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxPayloadSizeKB
}

func (o *OutputPrometheus) GetMaxPayloadEvents() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxPayloadEvents
}

func (o *OutputPrometheus) GetRejectUnauthorized() *bool {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *OutputPrometheus) GetTimeoutSec() *float64 {
	if o == nil {
		return nil
	}
	return o.TimeoutSec
}

func (o *OutputPrometheus) GetFlushPeriodSec() *float64 {
	if o == nil {
		return nil
	}
	return o.FlushPeriodSec
}

func (o *OutputPrometheus) GetExtraHTTPHeaders() []OutputPrometheusExtraHTTPHeaders {
	if o == nil {
		return nil
	}
	return o.ExtraHTTPHeaders
}

func (o *OutputPrometheus) GetUseRoundRobinDNS() *bool {
	if o == nil {
		return nil
	}
	return o.UseRoundRobinDNS
}

func (o *OutputPrometheus) GetFailedRequestLoggingMode() *OutputPrometheusFailedRequestLoggingMode {
	if o == nil {
		return nil
	}
	return o.FailedRequestLoggingMode
}

func (o *OutputPrometheus) GetSafeHeaders() []string {
	if o == nil {
		return nil
	}
	return o.SafeHeaders
}

func (o *OutputPrometheus) GetResponseRetrySettings() []OutputPrometheusResponseRetrySettings {
	if o == nil {
		return nil
	}
	return o.ResponseRetrySettings
}

func (o *OutputPrometheus) GetTimeoutRetrySettings() *OutputPrometheusTimeoutRetrySettings {
	if o == nil {
		return nil
	}
	return o.TimeoutRetrySettings
}

func (o *OutputPrometheus) GetResponseHonorRetryAfterHeader() *bool {
	if o == nil {
		return nil
	}
	return o.ResponseHonorRetryAfterHeader
}

func (o *OutputPrometheus) GetOnBackpressure() *OutputPrometheusBackpressureBehavior {
	if o == nil {
		return nil
	}
	return o.OnBackpressure
}

func (o *OutputPrometheus) GetAuthType() *OutputPrometheusAuthenticationType {
	if o == nil {
		return nil
	}
	return o.AuthType
}

func (o *OutputPrometheus) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *OutputPrometheus) GetMetricsFlushPeriodSec() *float64 {
	if o == nil {
		return nil
	}
	return o.MetricsFlushPeriodSec
}

func (o *OutputPrometheus) GetPqMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxFileSize
}

func (o *OutputPrometheus) GetPqMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxSize
}

func (o *OutputPrometheus) GetPqPath() *string {
	if o == nil {
		return nil
	}
	return o.PqPath
}

func (o *OutputPrometheus) GetPqCompress() *OutputPrometheusCompression {
	if o == nil {
		return nil
	}
	return o.PqCompress
}

func (o *OutputPrometheus) GetPqOnBackpressure() *OutputPrometheusQueueFullBehavior {
	if o == nil {
		return nil
	}
	return o.PqOnBackpressure
}

func (o *OutputPrometheus) GetPqMode() *OutputPrometheusMode {
	if o == nil {
		return nil
	}
	return o.PqMode
}

func (o *OutputPrometheus) GetPqControls() *OutputPrometheusPqControls {
	if o == nil {
		return nil
	}
	return o.PqControls
}

func (o *OutputPrometheus) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *OutputPrometheus) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *OutputPrometheus) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

func (o *OutputPrometheus) GetCredentialsSecret() *string {
	if o == nil {
		return nil
	}
	return o.CredentialsSecret
}

func (o *OutputPrometheus) GetTextSecret() *string {
	if o == nil {
		return nil
	}
	return o.TextSecret
}

func (o *OutputPrometheus) GetLoginURL() *string {
	if o == nil {
		return nil
	}
	return o.LoginURL
}

func (o *OutputPrometheus) GetSecretParamName() *string {
	if o == nil {
		return nil
	}
	return o.SecretParamName
}

func (o *OutputPrometheus) GetSecret() *string {
	if o == nil {
		return nil
	}
	return o.Secret
}

func (o *OutputPrometheus) GetTokenAttributeName() *string {
	if o == nil {
		return nil
	}
	return o.TokenAttributeName
}

func (o *OutputPrometheus) GetAuthHeaderExpr() *string {
	if o == nil {
		return nil
	}
	return o.AuthHeaderExpr
}

func (o *OutputPrometheus) GetTokenTimeoutSecs() *float64 {
	if o == nil {
		return nil
	}
	return o.TokenTimeoutSecs
}

func (o *OutputPrometheus) GetOauthParams() []OutputPrometheusOauthParams {
	if o == nil {
		return nil
	}
	return o.OauthParams
}

func (o *OutputPrometheus) GetOauthHeaders() []OutputPrometheusOauthHeaders {
	if o == nil {
		return nil
	}
	return o.OauthHeaders
}
