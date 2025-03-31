// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

type OutputWavefrontType string

const (
	OutputWavefrontTypeWavefront OutputWavefrontType = "wavefront"
)

func (e OutputWavefrontType) ToPointer() *OutputWavefrontType {
	return &e
}
func (e *OutputWavefrontType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "wavefront":
		*e = OutputWavefrontType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputWavefrontType: %v", v)
	}
}

// OutputWavefrontAuthenticationMethod - Select Manual to enter an auth token directly, or select Secret to use a text secret to authenticate
type OutputWavefrontAuthenticationMethod string

const (
	OutputWavefrontAuthenticationMethodManual OutputWavefrontAuthenticationMethod = "manual"
	OutputWavefrontAuthenticationMethodSecret OutputWavefrontAuthenticationMethod = "secret"
)

func (e OutputWavefrontAuthenticationMethod) ToPointer() *OutputWavefrontAuthenticationMethod {
	return &e
}
func (e *OutputWavefrontAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "manual":
		fallthrough
	case "secret":
		*e = OutputWavefrontAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputWavefrontAuthenticationMethod: %v", v)
	}
}

type OutputWavefrontExtraHTTPHeaders struct {
	// Field name
	Name *string `json:"name,omitempty"`
	// Field value
	Value string `json:"value"`
}

func (o *OutputWavefrontExtraHTTPHeaders) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *OutputWavefrontExtraHTTPHeaders) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// OutputWavefrontFailedRequestLoggingMode - Data to log when a request fails. All headers are redacted by default, unless listed as safe headers below.
type OutputWavefrontFailedRequestLoggingMode string

const (
	OutputWavefrontFailedRequestLoggingModePayload           OutputWavefrontFailedRequestLoggingMode = "payload"
	OutputWavefrontFailedRequestLoggingModePayloadAndHeaders OutputWavefrontFailedRequestLoggingMode = "payloadAndHeaders"
	OutputWavefrontFailedRequestLoggingModeNone              OutputWavefrontFailedRequestLoggingMode = "none"
)

func (e OutputWavefrontFailedRequestLoggingMode) ToPointer() *OutputWavefrontFailedRequestLoggingMode {
	return &e
}
func (e *OutputWavefrontFailedRequestLoggingMode) UnmarshalJSON(data []byte) error {
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
		*e = OutputWavefrontFailedRequestLoggingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputWavefrontFailedRequestLoggingMode: %v", v)
	}
}

type OutputWavefrontResponseRetrySettings struct {
	// The HTTP response status code that will trigger retries
	HTTPStatus float64 `json:"httpStatus"`
	// How long, in milliseconds, Cribl Stream should wait before initiating backoff. Maximum interval is 600,000 ms (10 minutes).
	InitialBackoff *float64 `default:"1000" json:"initialBackoff"`
	// Base for exponential backoff. A value of 2 (default) means Cribl Stream will retry after 2 seconds, then 4 seconds, then 8 seconds, etc.
	BackoffRate *float64 `default:"2" json:"backoffRate"`
	// The maximum backoff interval, in milliseconds, Cribl Stream should apply. Default (and minimum) is 10,000 ms (10 seconds); maximum is 180,000 ms (180 seconds).
	MaxBackoff *float64 `default:"10000" json:"maxBackoff"`
}

func (o OutputWavefrontResponseRetrySettings) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputWavefrontResponseRetrySettings) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputWavefrontResponseRetrySettings) GetHTTPStatus() float64 {
	if o == nil {
		return 0.0
	}
	return o.HTTPStatus
}

func (o *OutputWavefrontResponseRetrySettings) GetInitialBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialBackoff
}

func (o *OutputWavefrontResponseRetrySettings) GetBackoffRate() *float64 {
	if o == nil {
		return nil
	}
	return o.BackoffRate
}

func (o *OutputWavefrontResponseRetrySettings) GetMaxBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBackoff
}

type OutputWavefrontTimeoutRetrySettings struct {
	// Enable to retry on request timeout
	TimeoutRetry *bool `default:"false" json:"timeoutRetry"`
	// How long, in milliseconds, Cribl Stream should wait before initiating backoff. Maximum interval is 600,000 ms (10 minutes).
	InitialBackoff *float64 `default:"1000" json:"initialBackoff"`
	// Base for exponential backoff. A value of 2 (default) means Cribl Stream will retry after 2 seconds, then 4 seconds, then 8 seconds, etc.
	BackoffRate *float64 `default:"2" json:"backoffRate"`
	// The maximum backoff interval, in milliseconds, Cribl Stream should apply. Default (and minimum) is 10,000 ms (10 seconds); maximum is 180,000 ms (180 seconds).
	MaxBackoff *float64 `default:"10000" json:"maxBackoff"`
}

func (o OutputWavefrontTimeoutRetrySettings) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputWavefrontTimeoutRetrySettings) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputWavefrontTimeoutRetrySettings) GetTimeoutRetry() *bool {
	if o == nil {
		return nil
	}
	return o.TimeoutRetry
}

func (o *OutputWavefrontTimeoutRetrySettings) GetInitialBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialBackoff
}

func (o *OutputWavefrontTimeoutRetrySettings) GetBackoffRate() *float64 {
	if o == nil {
		return nil
	}
	return o.BackoffRate
}

func (o *OutputWavefrontTimeoutRetrySettings) GetMaxBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBackoff
}

// OutputWavefrontBackpressureBehavior - Whether to block, drop, or queue events when all receivers are exerting backpressure.
type OutputWavefrontBackpressureBehavior string

const (
	OutputWavefrontBackpressureBehaviorBlock OutputWavefrontBackpressureBehavior = "block"
	OutputWavefrontBackpressureBehaviorDrop  OutputWavefrontBackpressureBehavior = "drop"
	OutputWavefrontBackpressureBehaviorQueue OutputWavefrontBackpressureBehavior = "queue"
)

func (e OutputWavefrontBackpressureBehavior) ToPointer() *OutputWavefrontBackpressureBehavior {
	return &e
}
func (e *OutputWavefrontBackpressureBehavior) UnmarshalJSON(data []byte) error {
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
		*e = OutputWavefrontBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputWavefrontBackpressureBehavior: %v", v)
	}
}

// OutputWavefrontCompression - Codec to use to compress the persisted data.
type OutputWavefrontCompression string

const (
	OutputWavefrontCompressionNone OutputWavefrontCompression = "none"
	OutputWavefrontCompressionGzip OutputWavefrontCompression = "gzip"
)

func (e OutputWavefrontCompression) ToPointer() *OutputWavefrontCompression {
	return &e
}
func (e *OutputWavefrontCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputWavefrontCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputWavefrontCompression: %v", v)
	}
}

// OutputWavefrontQueueFullBehavior - Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputWavefrontQueueFullBehavior string

const (
	OutputWavefrontQueueFullBehaviorBlock OutputWavefrontQueueFullBehavior = "block"
	OutputWavefrontQueueFullBehaviorDrop  OutputWavefrontQueueFullBehavior = "drop"
)

func (e OutputWavefrontQueueFullBehavior) ToPointer() *OutputWavefrontQueueFullBehavior {
	return &e
}
func (e *OutputWavefrontQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputWavefrontQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputWavefrontQueueFullBehavior: %v", v)
	}
}

// OutputWavefrontMode - In Error mode, PQ writes events to the filesystem only when it detects a non-retryable Destination error. In Backpressure mode, PQ writes events to the filesystem when it detects backpressure from the Destination or when there are non-retryable Destination errors. In Always On mode, PQ always writes events to the filesystem.
type OutputWavefrontMode string

const (
	OutputWavefrontModeError        OutputWavefrontMode = "error"
	OutputWavefrontModeBackpressure OutputWavefrontMode = "backpressure"
	OutputWavefrontModeAlways       OutputWavefrontMode = "always"
)

func (e OutputWavefrontMode) ToPointer() *OutputWavefrontMode {
	return &e
}
func (e *OutputWavefrontMode) UnmarshalJSON(data []byte) error {
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
		*e = OutputWavefrontMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputWavefrontMode: %v", v)
	}
}

type OutputWavefrontPqControls struct {
}

type OutputWavefront struct {
	// Unique ID for this output
	ID   *string             `json:"id,omitempty"`
	Type OutputWavefrontType `json:"type"`
	// Pipeline to process data before sending out to this output
	Pipeline *string `json:"pipeline,omitempty"`
	// Fields to automatically add to events, such as cribl_pipe. Supports wildcards.
	SystemFields []string `json:"systemFields,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Tags for filtering and grouping in @{product}
	Streamtags []string `json:"streamtags,omitempty"`
	// Select Manual to enter an auth token directly, or select Secret to use a text secret to authenticate
	AuthType *OutputWavefrontAuthenticationMethod `default:"manual" json:"authType"`
	// WaveFront domain name, e.g. "longboard"
	Domain *string `default:"longboard" json:"domain"`
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
	ExtraHTTPHeaders []OutputWavefrontExtraHTTPHeaders `json:"extraHttpHeaders,omitempty"`
	// Enables round-robin DNS lookup. When a DNS server returns multiple addresses, @{product} will cycle through them in the order returned. For optimal performance, consider enabling this setting for non-load balanced destinations.
	UseRoundRobinDNS *bool `default:"false" json:"useRoundRobinDns"`
	// Data to log when a request fails. All headers are redacted by default, unless listed as safe headers below.
	FailedRequestLoggingMode *OutputWavefrontFailedRequestLoggingMode `default:"none" json:"failedRequestLoggingMode"`
	// List of headers that are safe to log in plain text
	SafeHeaders []string `json:"safeHeaders,omitempty"`
	// Automatically retry after unsuccessful response status codes, such as 429 (Too Many Requests) or 503 (Service Unavailable).
	ResponseRetrySettings []OutputWavefrontResponseRetrySettings `json:"responseRetrySettings,omitempty"`
	TimeoutRetrySettings  *OutputWavefrontTimeoutRetrySettings   `json:"timeoutRetrySettings,omitempty"`
	// Honor any Retry-After header that specifies a delay (in seconds) no longer than 180 seconds after the retry request. @{product} limits the delay to 180 seconds, even if the Retry-After header specifies a longer delay. When enabled, takes precedence over user-configured retry options. When disabled, all Retry-After headers are ignored.
	ResponseHonorRetryAfterHeader *bool `default:"false" json:"responseHonorRetryAfterHeader"`
	// Whether to block, drop, or queue events when all receivers are exerting backpressure.
	OnBackpressure *OutputWavefrontBackpressureBehavior `default:"block" json:"onBackpressure"`
	Description    *string                              `json:"description,omitempty"`
	// WaveFront API authentication token (see [here](https://docs.wavefront.com/wavefront_api.html#generating-an-api-token))
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
	PqCompress *OutputWavefrontCompression `default:"none" json:"pqCompress"`
	// Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputWavefrontQueueFullBehavior `default:"block" json:"pqOnBackpressure"`
	// In Error mode, PQ writes events to the filesystem only when it detects a non-retryable Destination error. In Backpressure mode, PQ writes events to the filesystem when it detects backpressure from the Destination or when there are non-retryable Destination errors. In Always On mode, PQ always writes events to the filesystem.
	PqMode     *OutputWavefrontMode       `default:"error" json:"pqMode"`
	PqControls *OutputWavefrontPqControls `json:"pqControls,omitempty"`
}

func (o OutputWavefront) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputWavefront) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *OutputWavefront) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OutputWavefront) GetType() OutputWavefrontType {
	if o == nil {
		return OutputWavefrontType("")
	}
	return o.Type
}

func (o *OutputWavefront) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *OutputWavefront) GetSystemFields() []string {
	if o == nil {
		return nil
	}
	return o.SystemFields
}

func (o *OutputWavefront) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *OutputWavefront) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *OutputWavefront) GetAuthType() *OutputWavefrontAuthenticationMethod {
	if o == nil {
		return nil
	}
	return o.AuthType
}

func (o *OutputWavefront) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *OutputWavefront) GetConcurrency() *float64 {
	if o == nil {
		return nil
	}
	return o.Concurrency
}

func (o *OutputWavefront) GetMaxPayloadSizeKB() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxPayloadSizeKB
}

func (o *OutputWavefront) GetMaxPayloadEvents() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxPayloadEvents
}

func (o *OutputWavefront) GetCompress() *bool {
	if o == nil {
		return nil
	}
	return o.Compress
}

func (o *OutputWavefront) GetRejectUnauthorized() *bool {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *OutputWavefront) GetTimeoutSec() *float64 {
	if o == nil {
		return nil
	}
	return o.TimeoutSec
}

func (o *OutputWavefront) GetFlushPeriodSec() *float64 {
	if o == nil {
		return nil
	}
	return o.FlushPeriodSec
}

func (o *OutputWavefront) GetExtraHTTPHeaders() []OutputWavefrontExtraHTTPHeaders {
	if o == nil {
		return nil
	}
	return o.ExtraHTTPHeaders
}

func (o *OutputWavefront) GetUseRoundRobinDNS() *bool {
	if o == nil {
		return nil
	}
	return o.UseRoundRobinDNS
}

func (o *OutputWavefront) GetFailedRequestLoggingMode() *OutputWavefrontFailedRequestLoggingMode {
	if o == nil {
		return nil
	}
	return o.FailedRequestLoggingMode
}

func (o *OutputWavefront) GetSafeHeaders() []string {
	if o == nil {
		return nil
	}
	return o.SafeHeaders
}

func (o *OutputWavefront) GetResponseRetrySettings() []OutputWavefrontResponseRetrySettings {
	if o == nil {
		return nil
	}
	return o.ResponseRetrySettings
}

func (o *OutputWavefront) GetTimeoutRetrySettings() *OutputWavefrontTimeoutRetrySettings {
	if o == nil {
		return nil
	}
	return o.TimeoutRetrySettings
}

func (o *OutputWavefront) GetResponseHonorRetryAfterHeader() *bool {
	if o == nil {
		return nil
	}
	return o.ResponseHonorRetryAfterHeader
}

func (o *OutputWavefront) GetOnBackpressure() *OutputWavefrontBackpressureBehavior {
	if o == nil {
		return nil
	}
	return o.OnBackpressure
}

func (o *OutputWavefront) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *OutputWavefront) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

func (o *OutputWavefront) GetTextSecret() *string {
	if o == nil {
		return nil
	}
	return o.TextSecret
}

func (o *OutputWavefront) GetPqMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxFileSize
}

func (o *OutputWavefront) GetPqMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxSize
}

func (o *OutputWavefront) GetPqPath() *string {
	if o == nil {
		return nil
	}
	return o.PqPath
}

func (o *OutputWavefront) GetPqCompress() *OutputWavefrontCompression {
	if o == nil {
		return nil
	}
	return o.PqCompress
}

func (o *OutputWavefront) GetPqOnBackpressure() *OutputWavefrontQueueFullBehavior {
	if o == nil {
		return nil
	}
	return o.PqOnBackpressure
}

func (o *OutputWavefront) GetPqMode() *OutputWavefrontMode {
	if o == nil {
		return nil
	}
	return o.PqMode
}

func (o *OutputWavefront) GetPqControls() *OutputWavefrontPqControls {
	if o == nil {
		return nil
	}
	return o.PqControls
}
