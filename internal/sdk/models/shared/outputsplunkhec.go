// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

type OutputSplunkHecType string

const (
	OutputSplunkHecTypeSplunkHec OutputSplunkHecType = "splunk_hec"
)

func (e OutputSplunkHecType) ToPointer() *OutputSplunkHecType {
	return &e
}
func (e *OutputSplunkHecType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "splunk_hec":
		*e = OutputSplunkHecType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSplunkHecType: %v", v)
	}
}

type OutputSplunkHecExtraHTTPHeaders struct {
	// Field name
	Name *string `json:"name,omitempty"`
	// Field value
	Value string `json:"value"`
}

func (o *OutputSplunkHecExtraHTTPHeaders) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *OutputSplunkHecExtraHTTPHeaders) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// OutputSplunkHecFailedRequestLoggingMode - Data to log when a request fails. All headers are redacted by default, unless listed as safe headers below.
type OutputSplunkHecFailedRequestLoggingMode string

const (
	OutputSplunkHecFailedRequestLoggingModePayload           OutputSplunkHecFailedRequestLoggingMode = "payload"
	OutputSplunkHecFailedRequestLoggingModePayloadAndHeaders OutputSplunkHecFailedRequestLoggingMode = "payloadAndHeaders"
	OutputSplunkHecFailedRequestLoggingModeNone              OutputSplunkHecFailedRequestLoggingMode = "none"
)

func (e OutputSplunkHecFailedRequestLoggingMode) ToPointer() *OutputSplunkHecFailedRequestLoggingMode {
	return &e
}
func (e *OutputSplunkHecFailedRequestLoggingMode) UnmarshalJSON(data []byte) error {
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
		*e = OutputSplunkHecFailedRequestLoggingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSplunkHecFailedRequestLoggingMode: %v", v)
	}
}

// OutputSplunkHecAuthenticationMethod - Select Manual to enter an auth token directly, or select Secret to use a text secret to authenticate
type OutputSplunkHecAuthenticationMethod string

const (
	OutputSplunkHecAuthenticationMethodManual OutputSplunkHecAuthenticationMethod = "manual"
	OutputSplunkHecAuthenticationMethodSecret OutputSplunkHecAuthenticationMethod = "secret"
)

func (e OutputSplunkHecAuthenticationMethod) ToPointer() *OutputSplunkHecAuthenticationMethod {
	return &e
}
func (e *OutputSplunkHecAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "manual":
		fallthrough
	case "secret":
		*e = OutputSplunkHecAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSplunkHecAuthenticationMethod: %v", v)
	}
}

type OutputSplunkHecResponseRetrySettings struct {
	// The HTTP response status code that will trigger retries
	HTTPStatus float64 `json:"httpStatus"`
	// How long, in milliseconds, Cribl Stream should wait before initiating backoff. Maximum interval is 600,000 ms (10 minutes).
	InitialBackoff *float64 `default:"1000" json:"initialBackoff"`
	// Base for exponential backoff. A value of 2 (default) means Cribl Stream will retry after 2 seconds, then 4 seconds, then 8 seconds, etc.
	BackoffRate *float64 `default:"2" json:"backoffRate"`
	// The maximum backoff interval, in milliseconds, Cribl Stream should apply. Default (and minimum) is 10,000 ms (10 seconds); maximum is 180,000 ms (180 seconds).
	MaxBackoff *float64 `default:"10000" json:"maxBackoff"`
}

func (o OutputSplunkHecResponseRetrySettings) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputSplunkHecResponseRetrySettings) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputSplunkHecResponseRetrySettings) GetHTTPStatus() float64 {
	if o == nil {
		return 0.0
	}
	return o.HTTPStatus
}

func (o *OutputSplunkHecResponseRetrySettings) GetInitialBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialBackoff
}

func (o *OutputSplunkHecResponseRetrySettings) GetBackoffRate() *float64 {
	if o == nil {
		return nil
	}
	return o.BackoffRate
}

func (o *OutputSplunkHecResponseRetrySettings) GetMaxBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBackoff
}

type OutputSplunkHecTimeoutRetrySettings struct {
	// Enable to retry on request timeout
	TimeoutRetry *bool `default:"false" json:"timeoutRetry"`
	// How long, in milliseconds, Cribl Stream should wait before initiating backoff. Maximum interval is 600,000 ms (10 minutes).
	InitialBackoff *float64 `default:"1000" json:"initialBackoff"`
	// Base for exponential backoff. A value of 2 (default) means Cribl Stream will retry after 2 seconds, then 4 seconds, then 8 seconds, etc.
	BackoffRate *float64 `default:"2" json:"backoffRate"`
	// The maximum backoff interval, in milliseconds, Cribl Stream should apply. Default (and minimum) is 10,000 ms (10 seconds); maximum is 180,000 ms (180 seconds).
	MaxBackoff *float64 `default:"10000" json:"maxBackoff"`
}

func (o OutputSplunkHecTimeoutRetrySettings) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputSplunkHecTimeoutRetrySettings) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputSplunkHecTimeoutRetrySettings) GetTimeoutRetry() *bool {
	if o == nil {
		return nil
	}
	return o.TimeoutRetry
}

func (o *OutputSplunkHecTimeoutRetrySettings) GetInitialBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialBackoff
}

func (o *OutputSplunkHecTimeoutRetrySettings) GetBackoffRate() *float64 {
	if o == nil {
		return nil
	}
	return o.BackoffRate
}

func (o *OutputSplunkHecTimeoutRetrySettings) GetMaxBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBackoff
}

// OutputSplunkHecBackpressureBehavior - Whether to block, drop, or queue events when all receivers are exerting backpressure.
type OutputSplunkHecBackpressureBehavior string

const (
	OutputSplunkHecBackpressureBehaviorBlock OutputSplunkHecBackpressureBehavior = "block"
	OutputSplunkHecBackpressureBehaviorDrop  OutputSplunkHecBackpressureBehavior = "drop"
	OutputSplunkHecBackpressureBehaviorQueue OutputSplunkHecBackpressureBehavior = "queue"
)

func (e OutputSplunkHecBackpressureBehavior) ToPointer() *OutputSplunkHecBackpressureBehavior {
	return &e
}
func (e *OutputSplunkHecBackpressureBehavior) UnmarshalJSON(data []byte) error {
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
		*e = OutputSplunkHecBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSplunkHecBackpressureBehavior: %v", v)
	}
}

type OutputSplunkHecUrls struct {
	// URL to a Splunk HEC endpoint to send events to, e.g., http://localhost:8088/services/collector/event
	URL *string `default:"http://localhost:8088/services/collector/event" json:"url"`
	// Assign a weight (>0) to each endpoint to indicate its traffic-handling capability
	Weight *float64 `default:"1" json:"weight"`
}

func (o OutputSplunkHecUrls) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputSplunkHecUrls) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputSplunkHecUrls) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *OutputSplunkHecUrls) GetWeight() *float64 {
	if o == nil {
		return nil
	}
	return o.Weight
}

// OutputSplunkHecCompression - Codec to use to compress the persisted data.
type OutputSplunkHecCompression string

const (
	OutputSplunkHecCompressionNone OutputSplunkHecCompression = "none"
	OutputSplunkHecCompressionGzip OutputSplunkHecCompression = "gzip"
)

func (e OutputSplunkHecCompression) ToPointer() *OutputSplunkHecCompression {
	return &e
}
func (e *OutputSplunkHecCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputSplunkHecCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSplunkHecCompression: %v", v)
	}
}

// OutputSplunkHecQueueFullBehavior - Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputSplunkHecQueueFullBehavior string

const (
	OutputSplunkHecQueueFullBehaviorBlock OutputSplunkHecQueueFullBehavior = "block"
	OutputSplunkHecQueueFullBehaviorDrop  OutputSplunkHecQueueFullBehavior = "drop"
)

func (e OutputSplunkHecQueueFullBehavior) ToPointer() *OutputSplunkHecQueueFullBehavior {
	return &e
}
func (e *OutputSplunkHecQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputSplunkHecQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSplunkHecQueueFullBehavior: %v", v)
	}
}

// OutputSplunkHecMode - In Error mode, PQ writes events to the filesystem only when it detects a non-retryable Destination error. In Backpressure mode, PQ writes events to the filesystem when it detects backpressure from the Destination or when there are non-retryable Destination errors. In Always On mode, PQ always writes events to the filesystem.
type OutputSplunkHecMode string

const (
	OutputSplunkHecModeError        OutputSplunkHecMode = "error"
	OutputSplunkHecModeBackpressure OutputSplunkHecMode = "backpressure"
	OutputSplunkHecModeAlways       OutputSplunkHecMode = "always"
)

func (e OutputSplunkHecMode) ToPointer() *OutputSplunkHecMode {
	return &e
}
func (e *OutputSplunkHecMode) UnmarshalJSON(data []byte) error {
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
		*e = OutputSplunkHecMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSplunkHecMode: %v", v)
	}
}

type OutputSplunkHecPqControls struct {
}

type OutputSplunkHec struct {
	// Unique ID for this output
	ID   string              `json:"id"`
	Type OutputSplunkHecType `json:"type"`
	// Pipeline to process data before sending out to this output
	Pipeline *string `json:"pipeline,omitempty"`
	// Fields to automatically add to events, such as cribl_pipe. Supports wildcards.
	SystemFields []string `json:"systemFields,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Tags for filtering and grouping in @{product}
	Streamtags []string `json:"streamtags,omitempty"`
	// Enable for optimal performance. Even if you have one hostname, it can expand to multiple IPs. If disabled, consider enabling round-robin DNS.
	LoadBalanced *bool `default:"true" json:"loadBalanced"`
	// In the Splunk app, define which Splunk processing queue to send the events after HEC processing.
	NextQueue *string `default:"indexQueue" json:"nextQueue"`
	// In the Splunk app, set the value of _TCP_ROUTING for events that do not have _ctrl._TCP_ROUTING set.
	TCPRouting *string `default:"nowhere" json:"tcpRouting"`
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
	ExtraHTTPHeaders []OutputSplunkHecExtraHTTPHeaders `json:"extraHttpHeaders,omitempty"`
	// Data to log when a request fails. All headers are redacted by default, unless listed as safe headers below.
	FailedRequestLoggingMode *OutputSplunkHecFailedRequestLoggingMode `default:"none" json:"failedRequestLoggingMode"`
	// List of headers that are safe to log in plain text
	SafeHeaders []string `json:"safeHeaders,omitempty"`
	// Output metrics in multiple-metric format, supported in Splunk 8.0 and above to allow multiple metrics in a single event.
	EnableMultiMetrics *bool `default:"false" json:"enableMultiMetrics"`
	// Select Manual to enter an auth token directly, or select Secret to use a text secret to authenticate
	AuthType *OutputSplunkHecAuthenticationMethod `default:"manual" json:"authType"`
	// Automatically retry after unsuccessful response status codes, such as 429 (Too Many Requests) or 503 (Service Unavailable).
	ResponseRetrySettings []OutputSplunkHecResponseRetrySettings `json:"responseRetrySettings,omitempty"`
	TimeoutRetrySettings  *OutputSplunkHecTimeoutRetrySettings   `json:"timeoutRetrySettings,omitempty"`
	// Honor any Retry-After header that specifies a delay (in seconds) no longer than 180 seconds after the retry request. @{product} limits the delay to 180 seconds, even if the Retry-After header specifies a longer delay. When enabled, takes precedence over user-configured retry options. When disabled, all Retry-After headers are ignored.
	ResponseHonorRetryAfterHeader *bool `default:"false" json:"responseHonorRetryAfterHeader"`
	// Whether to block, drop, or queue events when all receivers are exerting backpressure.
	OnBackpressure *OutputSplunkHecBackpressureBehavior `default:"block" json:"onBackpressure"`
	Description    *string                              `json:"description,omitempty"`
	// URL to a Splunk HEC endpoint to send events to, e.g., http://localhost:8088/services/collector/event
	URL *string `default:"http://localhost:8088/services/collector/event" json:"url"`
	// Enables round-robin DNS lookup. When a DNS server returns multiple addresses, @{product} will cycle through them in the order returned. For optimal performance, consider enabling this setting for non-load balanced destinations.
	UseRoundRobinDNS *bool `default:"false" json:"useRoundRobinDns"`
	// Exclude all IPs of the current host from the list of any resolved hostnames.
	ExcludeSelf *bool                 `default:"false" json:"excludeSelf"`
	Urls        []OutputSplunkHecUrls `json:"urls,omitempty"`
	// Re-resolve any hostnames every this many seconds and pick up destinations from A records.
	DNSResolvePeriodSec *float64 `default:"600" json:"dnsResolvePeriodSec"`
	// How far back in time to keep traffic stats for load balancing purposes.
	LoadBalanceStatsPeriodSec *float64 `default:"300" json:"loadBalanceStatsPeriodSec"`
	// Splunk HEC authentication token
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
	PqCompress *OutputSplunkHecCompression `default:"none" json:"pqCompress"`
	// Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputSplunkHecQueueFullBehavior `default:"block" json:"pqOnBackpressure"`
	// In Error mode, PQ writes events to the filesystem only when it detects a non-retryable Destination error. In Backpressure mode, PQ writes events to the filesystem when it detects backpressure from the Destination or when there are non-retryable Destination errors. In Always On mode, PQ always writes events to the filesystem.
	PqMode     *OutputSplunkHecMode       `default:"error" json:"pqMode"`
	PqControls *OutputSplunkHecPqControls `json:"pqControls,omitempty"`
}

func (o OutputSplunkHec) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputSplunkHec) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *OutputSplunkHec) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *OutputSplunkHec) GetType() OutputSplunkHecType {
	if o == nil {
		return OutputSplunkHecType("")
	}
	return o.Type
}

func (o *OutputSplunkHec) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *OutputSplunkHec) GetSystemFields() []string {
	if o == nil {
		return nil
	}
	return o.SystemFields
}

func (o *OutputSplunkHec) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *OutputSplunkHec) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *OutputSplunkHec) GetLoadBalanced() *bool {
	if o == nil {
		return nil
	}
	return o.LoadBalanced
}

func (o *OutputSplunkHec) GetNextQueue() *string {
	if o == nil {
		return nil
	}
	return o.NextQueue
}

func (o *OutputSplunkHec) GetTCPRouting() *string {
	if o == nil {
		return nil
	}
	return o.TCPRouting
}

func (o *OutputSplunkHec) GetConcurrency() *float64 {
	if o == nil {
		return nil
	}
	return o.Concurrency
}

func (o *OutputSplunkHec) GetMaxPayloadSizeKB() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxPayloadSizeKB
}

func (o *OutputSplunkHec) GetMaxPayloadEvents() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxPayloadEvents
}

func (o *OutputSplunkHec) GetCompress() *bool {
	if o == nil {
		return nil
	}
	return o.Compress
}

func (o *OutputSplunkHec) GetRejectUnauthorized() *bool {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *OutputSplunkHec) GetTimeoutSec() *float64 {
	if o == nil {
		return nil
	}
	return o.TimeoutSec
}

func (o *OutputSplunkHec) GetFlushPeriodSec() *float64 {
	if o == nil {
		return nil
	}
	return o.FlushPeriodSec
}

func (o *OutputSplunkHec) GetExtraHTTPHeaders() []OutputSplunkHecExtraHTTPHeaders {
	if o == nil {
		return nil
	}
	return o.ExtraHTTPHeaders
}

func (o *OutputSplunkHec) GetFailedRequestLoggingMode() *OutputSplunkHecFailedRequestLoggingMode {
	if o == nil {
		return nil
	}
	return o.FailedRequestLoggingMode
}

func (o *OutputSplunkHec) GetSafeHeaders() []string {
	if o == nil {
		return nil
	}
	return o.SafeHeaders
}

func (o *OutputSplunkHec) GetEnableMultiMetrics() *bool {
	if o == nil {
		return nil
	}
	return o.EnableMultiMetrics
}

func (o *OutputSplunkHec) GetAuthType() *OutputSplunkHecAuthenticationMethod {
	if o == nil {
		return nil
	}
	return o.AuthType
}

func (o *OutputSplunkHec) GetResponseRetrySettings() []OutputSplunkHecResponseRetrySettings {
	if o == nil {
		return nil
	}
	return o.ResponseRetrySettings
}

func (o *OutputSplunkHec) GetTimeoutRetrySettings() *OutputSplunkHecTimeoutRetrySettings {
	if o == nil {
		return nil
	}
	return o.TimeoutRetrySettings
}

func (o *OutputSplunkHec) GetResponseHonorRetryAfterHeader() *bool {
	if o == nil {
		return nil
	}
	return o.ResponseHonorRetryAfterHeader
}

func (o *OutputSplunkHec) GetOnBackpressure() *OutputSplunkHecBackpressureBehavior {
	if o == nil {
		return nil
	}
	return o.OnBackpressure
}

func (o *OutputSplunkHec) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *OutputSplunkHec) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *OutputSplunkHec) GetUseRoundRobinDNS() *bool {
	if o == nil {
		return nil
	}
	return o.UseRoundRobinDNS
}

func (o *OutputSplunkHec) GetExcludeSelf() *bool {
	if o == nil {
		return nil
	}
	return o.ExcludeSelf
}

func (o *OutputSplunkHec) GetUrls() []OutputSplunkHecUrls {
	if o == nil {
		return nil
	}
	return o.Urls
}

func (o *OutputSplunkHec) GetDNSResolvePeriodSec() *float64 {
	if o == nil {
		return nil
	}
	return o.DNSResolvePeriodSec
}

func (o *OutputSplunkHec) GetLoadBalanceStatsPeriodSec() *float64 {
	if o == nil {
		return nil
	}
	return o.LoadBalanceStatsPeriodSec
}

func (o *OutputSplunkHec) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

func (o *OutputSplunkHec) GetTextSecret() *string {
	if o == nil {
		return nil
	}
	return o.TextSecret
}

func (o *OutputSplunkHec) GetPqMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxFileSize
}

func (o *OutputSplunkHec) GetPqMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxSize
}

func (o *OutputSplunkHec) GetPqPath() *string {
	if o == nil {
		return nil
	}
	return o.PqPath
}

func (o *OutputSplunkHec) GetPqCompress() *OutputSplunkHecCompression {
	if o == nil {
		return nil
	}
	return o.PqCompress
}

func (o *OutputSplunkHec) GetPqOnBackpressure() *OutputSplunkHecQueueFullBehavior {
	if o == nil {
		return nil
	}
	return o.PqOnBackpressure
}

func (o *OutputSplunkHec) GetPqMode() *OutputSplunkHecMode {
	if o == nil {
		return nil
	}
	return o.PqMode
}

func (o *OutputSplunkHec) GetPqControls() *OutputSplunkHecPqControls {
	if o == nil {
		return nil
	}
	return o.PqControls
}
