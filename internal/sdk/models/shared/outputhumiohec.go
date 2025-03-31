// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

type OutputHumioHecType string

const (
	OutputHumioHecTypeHumioHec OutputHumioHecType = "humio_hec"
)

func (e OutputHumioHecType) ToPointer() *OutputHumioHecType {
	return &e
}
func (e *OutputHumioHecType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "humio_hec":
		*e = OutputHumioHecType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputHumioHecType: %v", v)
	}
}

type OutputHumioHecExtraHTTPHeaders struct {
	// Field name
	Name *string `json:"name,omitempty"`
	// Field value
	Value string `json:"value"`
}

func (o *OutputHumioHecExtraHTTPHeaders) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *OutputHumioHecExtraHTTPHeaders) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// OutputHumioHecFailedRequestLoggingMode - Data to log when a request fails. All headers are redacted by default, unless listed as safe headers below.
type OutputHumioHecFailedRequestLoggingMode string

const (
	OutputHumioHecFailedRequestLoggingModePayload           OutputHumioHecFailedRequestLoggingMode = "payload"
	OutputHumioHecFailedRequestLoggingModePayloadAndHeaders OutputHumioHecFailedRequestLoggingMode = "payloadAndHeaders"
	OutputHumioHecFailedRequestLoggingModeNone              OutputHumioHecFailedRequestLoggingMode = "none"
)

func (e OutputHumioHecFailedRequestLoggingMode) ToPointer() *OutputHumioHecFailedRequestLoggingMode {
	return &e
}
func (e *OutputHumioHecFailedRequestLoggingMode) UnmarshalJSON(data []byte) error {
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
		*e = OutputHumioHecFailedRequestLoggingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputHumioHecFailedRequestLoggingMode: %v", v)
	}
}

// RequestFormat - When set to JSON, the event is automatically formatted with required fields before sending. When set to Raw, only the event's `_raw` value is sent.
type RequestFormat string

const (
	RequestFormatJSON RequestFormat = "JSON"
	RequestFormatRaw  RequestFormat = "raw"
)

func (e RequestFormat) ToPointer() *RequestFormat {
	return &e
}
func (e *RequestFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "JSON":
		fallthrough
	case "raw":
		*e = RequestFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestFormat: %v", v)
	}
}

// OutputHumioHecAuthenticationMethod - Select Manual to enter an auth token directly, or select Secret to use a text secret to authenticate
type OutputHumioHecAuthenticationMethod string

const (
	OutputHumioHecAuthenticationMethodManual OutputHumioHecAuthenticationMethod = "manual"
	OutputHumioHecAuthenticationMethodSecret OutputHumioHecAuthenticationMethod = "secret"
)

func (e OutputHumioHecAuthenticationMethod) ToPointer() *OutputHumioHecAuthenticationMethod {
	return &e
}
func (e *OutputHumioHecAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "manual":
		fallthrough
	case "secret":
		*e = OutputHumioHecAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputHumioHecAuthenticationMethod: %v", v)
	}
}

type OutputHumioHecResponseRetrySettings struct {
	// The HTTP response status code that will trigger retries
	HTTPStatus float64 `json:"httpStatus"`
	// How long, in milliseconds, Cribl Stream should wait before initiating backoff. Maximum interval is 600,000 ms (10 minutes).
	InitialBackoff *float64 `default:"1000" json:"initialBackoff"`
	// Base for exponential backoff. A value of 2 (default) means Cribl Stream will retry after 2 seconds, then 4 seconds, then 8 seconds, etc.
	BackoffRate *float64 `default:"2" json:"backoffRate"`
	// The maximum backoff interval, in milliseconds, Cribl Stream should apply. Default (and minimum) is 10,000 ms (10 seconds); maximum is 180,000 ms (180 seconds).
	MaxBackoff *float64 `default:"10000" json:"maxBackoff"`
}

func (o OutputHumioHecResponseRetrySettings) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputHumioHecResponseRetrySettings) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputHumioHecResponseRetrySettings) GetHTTPStatus() float64 {
	if o == nil {
		return 0.0
	}
	return o.HTTPStatus
}

func (o *OutputHumioHecResponseRetrySettings) GetInitialBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialBackoff
}

func (o *OutputHumioHecResponseRetrySettings) GetBackoffRate() *float64 {
	if o == nil {
		return nil
	}
	return o.BackoffRate
}

func (o *OutputHumioHecResponseRetrySettings) GetMaxBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBackoff
}

type OutputHumioHecTimeoutRetrySettings struct {
	// Enable to retry on request timeout
	TimeoutRetry *bool `default:"false" json:"timeoutRetry"`
	// How long, in milliseconds, Cribl Stream should wait before initiating backoff. Maximum interval is 600,000 ms (10 minutes).
	InitialBackoff *float64 `default:"1000" json:"initialBackoff"`
	// Base for exponential backoff. A value of 2 (default) means Cribl Stream will retry after 2 seconds, then 4 seconds, then 8 seconds, etc.
	BackoffRate *float64 `default:"2" json:"backoffRate"`
	// The maximum backoff interval, in milliseconds, Cribl Stream should apply. Default (and minimum) is 10,000 ms (10 seconds); maximum is 180,000 ms (180 seconds).
	MaxBackoff *float64 `default:"10000" json:"maxBackoff"`
}

func (o OutputHumioHecTimeoutRetrySettings) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputHumioHecTimeoutRetrySettings) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputHumioHecTimeoutRetrySettings) GetTimeoutRetry() *bool {
	if o == nil {
		return nil
	}
	return o.TimeoutRetry
}

func (o *OutputHumioHecTimeoutRetrySettings) GetInitialBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialBackoff
}

func (o *OutputHumioHecTimeoutRetrySettings) GetBackoffRate() *float64 {
	if o == nil {
		return nil
	}
	return o.BackoffRate
}

func (o *OutputHumioHecTimeoutRetrySettings) GetMaxBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBackoff
}

// OutputHumioHecBackpressureBehavior - Whether to block, drop, or queue events when all receivers are exerting backpressure.
type OutputHumioHecBackpressureBehavior string

const (
	OutputHumioHecBackpressureBehaviorBlock OutputHumioHecBackpressureBehavior = "block"
	OutputHumioHecBackpressureBehaviorDrop  OutputHumioHecBackpressureBehavior = "drop"
	OutputHumioHecBackpressureBehaviorQueue OutputHumioHecBackpressureBehavior = "queue"
)

func (e OutputHumioHecBackpressureBehavior) ToPointer() *OutputHumioHecBackpressureBehavior {
	return &e
}
func (e *OutputHumioHecBackpressureBehavior) UnmarshalJSON(data []byte) error {
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
		*e = OutputHumioHecBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputHumioHecBackpressureBehavior: %v", v)
	}
}

// OutputHumioHecCompression - Codec to use to compress the persisted data.
type OutputHumioHecCompression string

const (
	OutputHumioHecCompressionNone OutputHumioHecCompression = "none"
	OutputHumioHecCompressionGzip OutputHumioHecCompression = "gzip"
)

func (e OutputHumioHecCompression) ToPointer() *OutputHumioHecCompression {
	return &e
}
func (e *OutputHumioHecCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputHumioHecCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputHumioHecCompression: %v", v)
	}
}

// OutputHumioHecQueueFullBehavior - Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputHumioHecQueueFullBehavior string

const (
	OutputHumioHecQueueFullBehaviorBlock OutputHumioHecQueueFullBehavior = "block"
	OutputHumioHecQueueFullBehaviorDrop  OutputHumioHecQueueFullBehavior = "drop"
)

func (e OutputHumioHecQueueFullBehavior) ToPointer() *OutputHumioHecQueueFullBehavior {
	return &e
}
func (e *OutputHumioHecQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputHumioHecQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputHumioHecQueueFullBehavior: %v", v)
	}
}

// OutputHumioHecMode - In Error mode, PQ writes events to the filesystem only when it detects a non-retryable Destination error. In Backpressure mode, PQ writes events to the filesystem when it detects backpressure from the Destination or when there are non-retryable Destination errors. In Always On mode, PQ always writes events to the filesystem.
type OutputHumioHecMode string

const (
	OutputHumioHecModeError        OutputHumioHecMode = "error"
	OutputHumioHecModeBackpressure OutputHumioHecMode = "backpressure"
	OutputHumioHecModeAlways       OutputHumioHecMode = "always"
)

func (e OutputHumioHecMode) ToPointer() *OutputHumioHecMode {
	return &e
}
func (e *OutputHumioHecMode) UnmarshalJSON(data []byte) error {
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
		*e = OutputHumioHecMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputHumioHecMode: %v", v)
	}
}

type OutputHumioHecPqControls struct {
}

type OutputHumioHec struct {
	// Unique ID for this output
	ID   *string             `json:"id,omitempty"`
	Type *OutputHumioHecType `json:"type,omitempty"`
	// Pipeline to process data before sending out to this output
	Pipeline *string `json:"pipeline,omitempty"`
	// Fields to automatically add to events, such as cribl_pipe. Supports wildcards.
	SystemFields []string `json:"systemFields,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Tags for filtering and grouping in @{product}
	Streamtags []string `json:"streamtags,omitempty"`
	// URL to a CrowdStrike Falcon LogScale endpoint to send events to, e.g., https://cloud.us.humio.com/api/v1/ingest/hec for JSON and https://cloud.us.humio.com/api/v1/ingest/hec/raw for raw
	URL *string `default:"https://cloud.us.humio.com/api/v1/ingest/hec" json:"url"`
	// Maximum number of ongoing requests before blocking
	Concurrency *float64 `default:"5" json:"concurrency"`
	// Maximum size, in KB, of the request body
	MaxPayloadSizeKB *float64 `default:"4096" json:"maxPayloadSizeKB"`
	// Maximum number of events to include in the request body. Default is 0 (unlimited).
	MaxPayloadEvents *float64 `default:"0" json:"maxPayloadEvents"`
	// Compress the payload body before sending
	Compress *bool `default:"true" json:"compress"`
	// Reject certificates not authorized by a CA in the CA certificate path or by another trusted CA (such as the system's).
	//         Defaults to Yes. When this setting is also present in TLS Settings (Client Side),
	//         that value will take precedence.
	RejectUnauthorized *bool `default:"true" json:"rejectUnauthorized"`
	// Amount of time, in seconds, to wait for a request to complete before canceling it
	TimeoutSec *float64 `default:"30" json:"timeoutSec"`
	// Maximum time between requests. Small values could cause the payload size to be smaller than the configured Max body size.
	FlushPeriodSec *float64 `default:"1" json:"flushPeriodSec"`
	// Headers to add to all events.
	ExtraHTTPHeaders []OutputHumioHecExtraHTTPHeaders `json:"extraHttpHeaders,omitempty"`
	// Enables round-robin DNS lookup. When a DNS server returns multiple addresses, @{product} will cycle through them in the order returned. For optimal performance, consider enabling this setting for non-load balanced destinations.
	UseRoundRobinDNS *bool `default:"true" json:"useRoundRobinDns"`
	// Data to log when a request fails. All headers are redacted by default, unless listed as safe headers below.
	FailedRequestLoggingMode *OutputHumioHecFailedRequestLoggingMode `default:"none" json:"failedRequestLoggingMode"`
	// List of headers that are safe to log in plain text
	SafeHeaders []string `json:"safeHeaders,omitempty"`
	// When set to JSON, the event is automatically formatted with required fields before sending. When set to Raw, only the event's `_raw` value is sent.
	Format *RequestFormat `default:"JSON" json:"format"`
	// Select Manual to enter an auth token directly, or select Secret to use a text secret to authenticate
	AuthType *OutputHumioHecAuthenticationMethod `default:"manual" json:"authType"`
	// Automatically retry after unsuccessful response status codes, such as 429 (Too Many Requests) or 503 (Service Unavailable).
	ResponseRetrySettings []OutputHumioHecResponseRetrySettings `json:"responseRetrySettings,omitempty"`
	TimeoutRetrySettings  *OutputHumioHecTimeoutRetrySettings   `json:"timeoutRetrySettings,omitempty"`
	// Honor any Retry-After header that specifies a delay (in seconds) no longer than 180 seconds after the retry request. @{product} limits the delay to 180 seconds, even if the Retry-After header specifies a longer delay. When enabled, takes precedence over user-configured retry options. When disabled, all Retry-After headers are ignored.
	ResponseHonorRetryAfterHeader *bool `default:"false" json:"responseHonorRetryAfterHeader"`
	// Whether to block, drop, or queue events when all receivers are exerting backpressure.
	OnBackpressure *OutputHumioHecBackpressureBehavior `default:"block" json:"onBackpressure"`
	Description    *string                             `json:"description,omitempty"`
	// CrowdStrike Falcon LogScale authentication token
	Token *string `json:"token,omitempty"`
	// Select or create a stored text secret
	TextSecret *string `json:"textSecret,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	PqMaxFileSize *string `default:"1 MB" json:"pqMaxFileSize"`
	// The maximum disk space that the queue can consume (as an average per Worker Process) before queueing stops. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `default:"5GB" json:"pqMaxSize"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/<output-id>.
	PqPath *string `default:"\\$CRIBL_HOME/state/queues" json:"pqPath"`
	// Codec to use to compress the persisted data.
	PqCompress *OutputHumioHecCompression `default:"none" json:"pqCompress"`
	// Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputHumioHecQueueFullBehavior `default:"block" json:"pqOnBackpressure"`
	// In Error mode, PQ writes events to the filesystem only when it detects a non-retryable Destination error. In Backpressure mode, PQ writes events to the filesystem when it detects backpressure from the Destination or when there are non-retryable Destination errors. In Always On mode, PQ always writes events to the filesystem.
	PqMode     *OutputHumioHecMode       `default:"error" json:"pqMode"`
	PqControls *OutputHumioHecPqControls `json:"pqControls,omitempty"`
}

func (o OutputHumioHec) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputHumioHec) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *OutputHumioHec) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OutputHumioHec) GetType() *OutputHumioHecType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *OutputHumioHec) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *OutputHumioHec) GetSystemFields() []string {
	if o == nil {
		return nil
	}
	return o.SystemFields
}

func (o *OutputHumioHec) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *OutputHumioHec) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *OutputHumioHec) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *OutputHumioHec) GetConcurrency() *float64 {
	if o == nil {
		return nil
	}
	return o.Concurrency
}

func (o *OutputHumioHec) GetMaxPayloadSizeKB() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxPayloadSizeKB
}

func (o *OutputHumioHec) GetMaxPayloadEvents() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxPayloadEvents
}

func (o *OutputHumioHec) GetCompress() *bool {
	if o == nil {
		return nil
	}
	return o.Compress
}

func (o *OutputHumioHec) GetRejectUnauthorized() *bool {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *OutputHumioHec) GetTimeoutSec() *float64 {
	if o == nil {
		return nil
	}
	return o.TimeoutSec
}

func (o *OutputHumioHec) GetFlushPeriodSec() *float64 {
	if o == nil {
		return nil
	}
	return o.FlushPeriodSec
}

func (o *OutputHumioHec) GetExtraHTTPHeaders() []OutputHumioHecExtraHTTPHeaders {
	if o == nil {
		return nil
	}
	return o.ExtraHTTPHeaders
}

func (o *OutputHumioHec) GetUseRoundRobinDNS() *bool {
	if o == nil {
		return nil
	}
	return o.UseRoundRobinDNS
}

func (o *OutputHumioHec) GetFailedRequestLoggingMode() *OutputHumioHecFailedRequestLoggingMode {
	if o == nil {
		return nil
	}
	return o.FailedRequestLoggingMode
}

func (o *OutputHumioHec) GetSafeHeaders() []string {
	if o == nil {
		return nil
	}
	return o.SafeHeaders
}

func (o *OutputHumioHec) GetFormat() *RequestFormat {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *OutputHumioHec) GetAuthType() *OutputHumioHecAuthenticationMethod {
	if o == nil {
		return nil
	}
	return o.AuthType
}

func (o *OutputHumioHec) GetResponseRetrySettings() []OutputHumioHecResponseRetrySettings {
	if o == nil {
		return nil
	}
	return o.ResponseRetrySettings
}

func (o *OutputHumioHec) GetTimeoutRetrySettings() *OutputHumioHecTimeoutRetrySettings {
	if o == nil {
		return nil
	}
	return o.TimeoutRetrySettings
}

func (o *OutputHumioHec) GetResponseHonorRetryAfterHeader() *bool {
	if o == nil {
		return nil
	}
	return o.ResponseHonorRetryAfterHeader
}

func (o *OutputHumioHec) GetOnBackpressure() *OutputHumioHecBackpressureBehavior {
	if o == nil {
		return nil
	}
	return o.OnBackpressure
}

func (o *OutputHumioHec) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *OutputHumioHec) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

func (o *OutputHumioHec) GetTextSecret() *string {
	if o == nil {
		return nil
	}
	return o.TextSecret
}

func (o *OutputHumioHec) GetPqMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxFileSize
}

func (o *OutputHumioHec) GetPqMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxSize
}

func (o *OutputHumioHec) GetPqPath() *string {
	if o == nil {
		return nil
	}
	return o.PqPath
}

func (o *OutputHumioHec) GetPqCompress() *OutputHumioHecCompression {
	if o == nil {
		return nil
	}
	return o.PqCompress
}

func (o *OutputHumioHec) GetPqOnBackpressure() *OutputHumioHecQueueFullBehavior {
	if o == nil {
		return nil
	}
	return o.PqOnBackpressure
}

func (o *OutputHumioHec) GetPqMode() *OutputHumioHecMode {
	if o == nil {
		return nil
	}
	return o.PqMode
}

func (o *OutputHumioHec) GetPqControls() *OutputHumioHecPqControls {
	if o == nil {
		return nil
	}
	return o.PqControls
}
