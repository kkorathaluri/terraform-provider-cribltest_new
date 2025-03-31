// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

type InputFirehoseType string

const (
	InputFirehoseTypeFirehose InputFirehoseType = "firehose"
)

func (e InputFirehoseType) ToPointer() *InputFirehoseType {
	return &e
}
func (e *InputFirehoseType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "firehose":
		*e = InputFirehoseType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputFirehoseType: %v", v)
	}
}

type InputFirehoseConnections struct {
	Pipeline *string `json:"pipeline,omitempty"`
	Output   string  `json:"output"`
}

func (o *InputFirehoseConnections) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputFirehoseConnections) GetOutput() string {
	if o == nil {
		return ""
	}
	return o.Output
}

// InputFirehoseMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputFirehoseMode string

const (
	InputFirehoseModeSmart  InputFirehoseMode = "smart"
	InputFirehoseModeAlways InputFirehoseMode = "always"
)

func (e InputFirehoseMode) ToPointer() *InputFirehoseMode {
	return &e
}
func (e *InputFirehoseMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputFirehoseMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputFirehoseMode: %v", v)
	}
}

// InputFirehoseCompression - Codec to use to compress the persisted data
type InputFirehoseCompression string

const (
	InputFirehoseCompressionNone InputFirehoseCompression = "none"
	InputFirehoseCompressionGzip InputFirehoseCompression = "gzip"
)

func (e InputFirehoseCompression) ToPointer() *InputFirehoseCompression {
	return &e
}
func (e *InputFirehoseCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputFirehoseCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputFirehoseCompression: %v", v)
	}
}

type InputFirehosePq struct {
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputFirehoseMode `default:"always" json:"mode"`
	// The maximum number of events to hold in memory before writing the events to disk
	MaxBufferSize *float64 `default:"1000" json:"maxBufferSize"`
	// The number of events to send downstream before committing that Stream has read them
	CommitFrequency *float64 `default:"42" json:"commitFrequency"`
	// The maximum size to store in each queue file before closing and optionally compressing. Enter a numeral with units of KB, MB, etc.
	MaxFileSize *string `default:"1 MB" json:"maxFileSize"`
	// The maximum disk space that the queue can consume (as an average per Worker Process) before queueing stops. Enter a numeral with units of KB, MB, etc.
	MaxSize *string `default:"5GB" json:"maxSize"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/inputs/<input-id>
	Path *string `default:"\\$CRIBL_HOME/state/queues" json:"path"`
	// Codec to use to compress the persisted data
	Compress *InputFirehoseCompression `default:"none" json:"compress"`
}

func (i InputFirehosePq) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputFirehosePq) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputFirehosePq) GetMode() *InputFirehoseMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputFirehosePq) GetMaxBufferSize() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBufferSize
}

func (o *InputFirehosePq) GetCommitFrequency() *float64 {
	if o == nil {
		return nil
	}
	return o.CommitFrequency
}

func (o *InputFirehosePq) GetMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxFileSize
}

func (o *InputFirehosePq) GetMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxSize
}

func (o *InputFirehosePq) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *InputFirehosePq) GetCompress() *InputFirehoseCompression {
	if o == nil {
		return nil
	}
	return o.Compress
}

// InputFirehoseMinimumTLSVersion - Minimum TLS version to accept from connections
type InputFirehoseMinimumTLSVersion string

const (
	InputFirehoseMinimumTLSVersionTlSv1  InputFirehoseMinimumTLSVersion = "TLSv1"
	InputFirehoseMinimumTLSVersionTlSv11 InputFirehoseMinimumTLSVersion = "TLSv1.1"
	InputFirehoseMinimumTLSVersionTlSv12 InputFirehoseMinimumTLSVersion = "TLSv1.2"
	InputFirehoseMinimumTLSVersionTlSv13 InputFirehoseMinimumTLSVersion = "TLSv1.3"
)

func (e InputFirehoseMinimumTLSVersion) ToPointer() *InputFirehoseMinimumTLSVersion {
	return &e
}
func (e *InputFirehoseMinimumTLSVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TLSv1":
		fallthrough
	case "TLSv1.1":
		fallthrough
	case "TLSv1.2":
		fallthrough
	case "TLSv1.3":
		*e = InputFirehoseMinimumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputFirehoseMinimumTLSVersion: %v", v)
	}
}

// InputFirehoseMaximumTLSVersion - Maximum TLS version to accept from connections
type InputFirehoseMaximumTLSVersion string

const (
	InputFirehoseMaximumTLSVersionTlSv1  InputFirehoseMaximumTLSVersion = "TLSv1"
	InputFirehoseMaximumTLSVersionTlSv11 InputFirehoseMaximumTLSVersion = "TLSv1.1"
	InputFirehoseMaximumTLSVersionTlSv12 InputFirehoseMaximumTLSVersion = "TLSv1.2"
	InputFirehoseMaximumTLSVersionTlSv13 InputFirehoseMaximumTLSVersion = "TLSv1.3"
)

func (e InputFirehoseMaximumTLSVersion) ToPointer() *InputFirehoseMaximumTLSVersion {
	return &e
}
func (e *InputFirehoseMaximumTLSVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TLSv1":
		fallthrough
	case "TLSv1.1":
		fallthrough
	case "TLSv1.2":
		fallthrough
	case "TLSv1.3":
		*e = InputFirehoseMaximumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputFirehoseMaximumTLSVersion: %v", v)
	}
}

type InputFirehoseTLSSettingsServerSide struct {
	Disabled *bool `default:"true" json:"disabled"`
	// The name of the predefined certificate
	CertificateName *string `json:"certificateName,omitempty"`
	// Path on server containing the private key to use. PEM format. Can reference $ENV_VARS.
	PrivKeyPath *string `json:"privKeyPath,omitempty"`
	// Passphrase to use to decrypt private key
	Passphrase *string `json:"passphrase,omitempty"`
	// Path on server containing certificates to use. PEM format. Can reference $ENV_VARS.
	CertPath *string `json:"certPath,omitempty"`
	// Path on server containing CA certificates to use. PEM format. Can reference $ENV_VARS.
	CaPath *string `json:"caPath,omitempty"`
	// Require clients to present their certificates. Used to perform client authentication using SSL certs.
	RequestCert        *bool `default:"false" json:"requestCert"`
	RejectUnauthorized any   `json:"rejectUnauthorized,omitempty"`
	CommonNameRegex    any   `json:"commonNameRegex,omitempty"`
	// Minimum TLS version to accept from connections
	MinVersion *InputFirehoseMinimumTLSVersion `json:"minVersion,omitempty"`
	// Maximum TLS version to accept from connections
	MaxVersion *InputFirehoseMaximumTLSVersion `json:"maxVersion,omitempty"`
}

func (i InputFirehoseTLSSettingsServerSide) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputFirehoseTLSSettingsServerSide) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputFirehoseTLSSettingsServerSide) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputFirehoseTLSSettingsServerSide) GetCertificateName() *string {
	if o == nil {
		return nil
	}
	return o.CertificateName
}

func (o *InputFirehoseTLSSettingsServerSide) GetPrivKeyPath() *string {
	if o == nil {
		return nil
	}
	return o.PrivKeyPath
}

func (o *InputFirehoseTLSSettingsServerSide) GetPassphrase() *string {
	if o == nil {
		return nil
	}
	return o.Passphrase
}

func (o *InputFirehoseTLSSettingsServerSide) GetCertPath() *string {
	if o == nil {
		return nil
	}
	return o.CertPath
}

func (o *InputFirehoseTLSSettingsServerSide) GetCaPath() *string {
	if o == nil {
		return nil
	}
	return o.CaPath
}

func (o *InputFirehoseTLSSettingsServerSide) GetRequestCert() *bool {
	if o == nil {
		return nil
	}
	return o.RequestCert
}

func (o *InputFirehoseTLSSettingsServerSide) GetRejectUnauthorized() any {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *InputFirehoseTLSSettingsServerSide) GetCommonNameRegex() any {
	if o == nil {
		return nil
	}
	return o.CommonNameRegex
}

func (o *InputFirehoseTLSSettingsServerSide) GetMinVersion() *InputFirehoseMinimumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MinVersion
}

func (o *InputFirehoseTLSSettingsServerSide) GetMaxVersion() *InputFirehoseMaximumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MaxVersion
}

type InputFirehoseMetadata struct {
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

func (o *InputFirehoseMetadata) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InputFirehoseMetadata) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type InputFirehose struct {
	// Unique ID for this input
	ID       *string            `json:"id,omitempty"`
	Type     *InputFirehoseType `json:"type,omitempty"`
	Disabled *bool              `default:"false" json:"disabled"`
	// Pipeline to process data from this Source before sending it through the Routes
	Pipeline *string `json:"pipeline,omitempty"`
	// Select whether to send data to Routes, or directly to Destinations.
	SendToRoutes *bool `default:"true" json:"sendToRoutes"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Use a disk queue to minimize data loss when connected services block. See [Cribl Docs](https://docs.cribl.io/stream/persistent-queues) for PQ defaults (Cribl-managed Cloud Workers) and configuration options (on-prem and hybrid Workers).
	PqEnabled *bool `default:"false" json:"pqEnabled"`
	// Tags for filtering and grouping in @{product}
	Streamtags []string `json:"streamtags,omitempty"`
	// Direct connections to Destinations, and optionally via a Pipeline or a Pack
	Connections []InputFirehoseConnections `json:"connections,omitempty"`
	Pq          *InputFirehosePq           `json:"pq,omitempty"`
	// Address to bind on. Defaults to 0.0.0.0 (all addresses).
	Host *string `default:"0.0.0.0" json:"host"`
	// Port to listen on
	Port float64 `json:"port"`
	// Shared secrets to be provided by any client (Authorization: <token>). If empty, unauthorized access is permitted.
	AuthTokens []string                            `json:"authTokens,omitempty"`
	TLS        *InputFirehoseTLSSettingsServerSide `json:"tls,omitempty"`
	// Maximum number of active requests per Worker Process. Use 0 for unlimited.
	MaxActiveReq *float64 `default:"256" json:"maxActiveReq"`
	// Maximum number of requests per socket before @{product} instructs the client to close the connection. Default is 0 (unlimited).
	MaxRequestsPerSocket *int64 `default:"0" json:"maxRequestsPerSocket"`
	// Enable when clients are connecting through a proxy that supports the x-forwarded-for header to keep the client's original IP address on the event instead of the proxy's IP address
	EnableProxyHeader *bool `default:"false" json:"enableProxyHeader"`
	// Toggle this to Yes to add request headers to events, in the __headers field.
	CaptureHeaders *bool `default:"false" json:"captureHeaders"`
	// How often request activity is logged at the `info` level. A value of 1 would log every request, 10 every 10th request, etc.
	ActivityLogSampleRate *float64 `default:"100" json:"activityLogSampleRate"`
	// How long to wait for an incoming request to complete before aborting it. Use 0 to disable.
	RequestTimeout *float64 `default:"0" json:"requestTimeout"`
	// How long @{product} should wait before assuming that an inactive socket has timed out. To wait forever, set to 0.
	SocketTimeout *float64 `default:"0" json:"socketTimeout"`
	// After the last response is sent, @{product} will wait this long for additional data before closing the socket connection. Minimum 1 sec.; maximum 600 sec. (10 min.).
	KeepAliveTimeout *float64 `default:"5" json:"keepAliveTimeout"`
	// Enable to expose the /cribl_health endpoint, which returns 200 OK when this Source is healthy
	EnableHealthCheck *bool `default:"false" json:"enableHealthCheck"`
	// Messages from matched IP addresses will be processed, unless also matched by the denylist.
	IPAllowlistRegex *string `default:"/.*/" json:"ipAllowlistRegex"`
	// Messages from matched IP addresses will be ignored. This takes precedence over the allowlist.
	IPDenylistRegex *string `default:"/^\\$/" json:"ipDenylistRegex"`
	// Fields to add to events from this input
	Metadata    []InputFirehoseMetadata `json:"metadata,omitempty"`
	Description *string                 `json:"description,omitempty"`
}

func (i InputFirehose) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputFirehose) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *InputFirehose) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *InputFirehose) GetType() *InputFirehoseType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *InputFirehose) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputFirehose) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputFirehose) GetSendToRoutes() *bool {
	if o == nil {
		return nil
	}
	return o.SendToRoutes
}

func (o *InputFirehose) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *InputFirehose) GetPqEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.PqEnabled
}

func (o *InputFirehose) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *InputFirehose) GetConnections() []InputFirehoseConnections {
	if o == nil {
		return nil
	}
	return o.Connections
}

func (o *InputFirehose) GetPq() *InputFirehosePq {
	if o == nil {
		return nil
	}
	return o.Pq
}

func (o *InputFirehose) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *InputFirehose) GetPort() float64 {
	if o == nil {
		return 0.0
	}
	return o.Port
}

func (o *InputFirehose) GetAuthTokens() []string {
	if o == nil {
		return nil
	}
	return o.AuthTokens
}

func (o *InputFirehose) GetTLS() *InputFirehoseTLSSettingsServerSide {
	if o == nil {
		return nil
	}
	return o.TLS
}

func (o *InputFirehose) GetMaxActiveReq() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxActiveReq
}

func (o *InputFirehose) GetMaxRequestsPerSocket() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxRequestsPerSocket
}

func (o *InputFirehose) GetEnableProxyHeader() *bool {
	if o == nil {
		return nil
	}
	return o.EnableProxyHeader
}

func (o *InputFirehose) GetCaptureHeaders() *bool {
	if o == nil {
		return nil
	}
	return o.CaptureHeaders
}

func (o *InputFirehose) GetActivityLogSampleRate() *float64 {
	if o == nil {
		return nil
	}
	return o.ActivityLogSampleRate
}

func (o *InputFirehose) GetRequestTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.RequestTimeout
}

func (o *InputFirehose) GetSocketTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.SocketTimeout
}

func (o *InputFirehose) GetKeepAliveTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.KeepAliveTimeout
}

func (o *InputFirehose) GetEnableHealthCheck() *bool {
	if o == nil {
		return nil
	}
	return o.EnableHealthCheck
}

func (o *InputFirehose) GetIPAllowlistRegex() *string {
	if o == nil {
		return nil
	}
	return o.IPAllowlistRegex
}

func (o *InputFirehose) GetIPDenylistRegex() *string {
	if o == nil {
		return nil
	}
	return o.IPDenylistRegex
}

func (o *InputFirehose) GetMetadata() []InputFirehoseMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *InputFirehose) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}
