// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

type InputSplunkType string

const (
	InputSplunkTypeSplunk InputSplunkType = "splunk"
)

func (e InputSplunkType) ToPointer() *InputSplunkType {
	return &e
}
func (e *InputSplunkType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "splunk":
		*e = InputSplunkType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSplunkType: %v", v)
	}
}

type InputSplunkConnections struct {
	Pipeline *string `json:"pipeline,omitempty"`
	Output   string  `json:"output"`
}

func (o *InputSplunkConnections) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputSplunkConnections) GetOutput() string {
	if o == nil {
		return ""
	}
	return o.Output
}

// InputSplunkMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputSplunkMode string

const (
	InputSplunkModeSmart  InputSplunkMode = "smart"
	InputSplunkModeAlways InputSplunkMode = "always"
)

func (e InputSplunkMode) ToPointer() *InputSplunkMode {
	return &e
}
func (e *InputSplunkMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputSplunkMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSplunkMode: %v", v)
	}
}

// InputSplunkCompression - Codec to use to compress the persisted data
type InputSplunkCompression string

const (
	InputSplunkCompressionNone InputSplunkCompression = "none"
	InputSplunkCompressionGzip InputSplunkCompression = "gzip"
)

func (e InputSplunkCompression) ToPointer() *InputSplunkCompression {
	return &e
}
func (e *InputSplunkCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputSplunkCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSplunkCompression: %v", v)
	}
}

type InputSplunkPq struct {
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputSplunkMode `default:"always" json:"mode"`
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
	Compress *InputSplunkCompression `default:"none" json:"compress"`
}

func (i InputSplunkPq) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputSplunkPq) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputSplunkPq) GetMode() *InputSplunkMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputSplunkPq) GetMaxBufferSize() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBufferSize
}

func (o *InputSplunkPq) GetCommitFrequency() *float64 {
	if o == nil {
		return nil
	}
	return o.CommitFrequency
}

func (o *InputSplunkPq) GetMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxFileSize
}

func (o *InputSplunkPq) GetMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxSize
}

func (o *InputSplunkPq) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *InputSplunkPq) GetCompress() *InputSplunkCompression {
	if o == nil {
		return nil
	}
	return o.Compress
}

// InputSplunkMinimumTLSVersion - Minimum TLS version to accept from connections
type InputSplunkMinimumTLSVersion string

const (
	InputSplunkMinimumTLSVersionTlSv1  InputSplunkMinimumTLSVersion = "TLSv1"
	InputSplunkMinimumTLSVersionTlSv11 InputSplunkMinimumTLSVersion = "TLSv1.1"
	InputSplunkMinimumTLSVersionTlSv12 InputSplunkMinimumTLSVersion = "TLSv1.2"
	InputSplunkMinimumTLSVersionTlSv13 InputSplunkMinimumTLSVersion = "TLSv1.3"
)

func (e InputSplunkMinimumTLSVersion) ToPointer() *InputSplunkMinimumTLSVersion {
	return &e
}
func (e *InputSplunkMinimumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = InputSplunkMinimumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSplunkMinimumTLSVersion: %v", v)
	}
}

// InputSplunkMaximumTLSVersion - Maximum TLS version to accept from connections
type InputSplunkMaximumTLSVersion string

const (
	InputSplunkMaximumTLSVersionTlSv1  InputSplunkMaximumTLSVersion = "TLSv1"
	InputSplunkMaximumTLSVersionTlSv11 InputSplunkMaximumTLSVersion = "TLSv1.1"
	InputSplunkMaximumTLSVersionTlSv12 InputSplunkMaximumTLSVersion = "TLSv1.2"
	InputSplunkMaximumTLSVersionTlSv13 InputSplunkMaximumTLSVersion = "TLSv1.3"
)

func (e InputSplunkMaximumTLSVersion) ToPointer() *InputSplunkMaximumTLSVersion {
	return &e
}
func (e *InputSplunkMaximumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = InputSplunkMaximumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSplunkMaximumTLSVersion: %v", v)
	}
}

type InputSplunkTLSSettingsServerSide struct {
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
	MinVersion *InputSplunkMinimumTLSVersion `json:"minVersion,omitempty"`
	// Maximum TLS version to accept from connections
	MaxVersion *InputSplunkMaximumTLSVersion `json:"maxVersion,omitempty"`
}

func (i InputSplunkTLSSettingsServerSide) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputSplunkTLSSettingsServerSide) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputSplunkTLSSettingsServerSide) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputSplunkTLSSettingsServerSide) GetCertificateName() *string {
	if o == nil {
		return nil
	}
	return o.CertificateName
}

func (o *InputSplunkTLSSettingsServerSide) GetPrivKeyPath() *string {
	if o == nil {
		return nil
	}
	return o.PrivKeyPath
}

func (o *InputSplunkTLSSettingsServerSide) GetPassphrase() *string {
	if o == nil {
		return nil
	}
	return o.Passphrase
}

func (o *InputSplunkTLSSettingsServerSide) GetCertPath() *string {
	if o == nil {
		return nil
	}
	return o.CertPath
}

func (o *InputSplunkTLSSettingsServerSide) GetCaPath() *string {
	if o == nil {
		return nil
	}
	return o.CaPath
}

func (o *InputSplunkTLSSettingsServerSide) GetRequestCert() *bool {
	if o == nil {
		return nil
	}
	return o.RequestCert
}

func (o *InputSplunkTLSSettingsServerSide) GetRejectUnauthorized() any {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *InputSplunkTLSSettingsServerSide) GetCommonNameRegex() any {
	if o == nil {
		return nil
	}
	return o.CommonNameRegex
}

func (o *InputSplunkTLSSettingsServerSide) GetMinVersion() *InputSplunkMinimumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MinVersion
}

func (o *InputSplunkTLSSettingsServerSide) GetMaxVersion() *InputSplunkMaximumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MaxVersion
}

type InputSplunkMetadata struct {
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

func (o *InputSplunkMetadata) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InputSplunkMetadata) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type AuthTokens struct {
	// Shared secrets to be provided by any Splunk forwarder. If empty, unauthorized access is permitted.
	Token       string  `json:"token"`
	Description *string `json:"description,omitempty"`
}

func (o *AuthTokens) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

func (o *AuthTokens) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

// MaxS2SVersion - The highest S2S protocol version to advertise during handshake
type MaxS2SVersion string

const (
	MaxS2SVersionV3 MaxS2SVersion = "v3"
	MaxS2SVersionV4 MaxS2SVersion = "v4"
)

func (e MaxS2SVersion) ToPointer() *MaxS2SVersion {
	return &e
}
func (e *MaxS2SVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "v3":
		fallthrough
	case "v4":
		*e = MaxS2SVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MaxS2SVersion: %v", v)
	}
}

type InputSplunk struct {
	// Unique ID for this input
	ID       *string          `json:"id,omitempty"`
	Type     *InputSplunkType `json:"type,omitempty"`
	Disabled *bool            `default:"false" json:"disabled"`
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
	Connections []InputSplunkConnections `json:"connections,omitempty"`
	Pq          *InputSplunkPq           `json:"pq,omitempty"`
	// Address to bind on. Defaults to 0.0.0.0 (all addresses).
	Host *string `default:"0.0.0.0" json:"host"`
	// Port to listen on
	Port float64                           `json:"port"`
	TLS  *InputSplunkTLSSettingsServerSide `json:"tls,omitempty"`
	// Regex matching IP addresses that are allowed to establish a connection
	IPWhitelistRegex *string `default:"/.*/" json:"ipWhitelistRegex"`
	// Maximum number of active connections allowed per Worker Process. Use 0 for unlimited.
	MaxActiveCxn *float64 `default:"1000" json:"maxActiveCxn"`
	// How long @{product} should wait before assuming that an inactive socket has timed out. After this time, the connection will be closed. Leave at 0 for no inactive socket monitoring.
	SocketIdleTimeout *float64 `default:"0" json:"socketIdleTimeout"`
	// How long the server will wait after initiating a closure for a client to close its end of the connection. If the client doesn't close the connection within this time, the server will forcefully terminate the socket to prevent resource leaks and ensure efficient connection cleanup and system stability. Leave at 0 for no inactive socket monitoring.
	SocketEndingMaxWait *float64 `default:"30" json:"socketEndingMaxWait"`
	// The maximum duration a socket can remain open, even if active. This helps manage resources and mitigate issues caused by TCP pinning. Set to 0 to disable.
	SocketMaxLifespan *float64 `default:"0" json:"socketMaxLifespan"`
	// Enable if the connection is proxied by a device that supports proxy protocol v1 or v2
	EnableProxyHeader *bool `default:"false" json:"enableProxyHeader"`
	// Fields to add to events from this input
	Metadata []InputSplunkMetadata `json:"metadata,omitempty"`
	// A list of event-breaking rulesets that will be applied, in order, to the input data stream
	BreakerRulesets []string `json:"breakerRulesets,omitempty"`
	// How long (in milliseconds) the Event Breaker will wait for new data to be sent to a specific channel before flushing the data stream out, as is, to the Pipelines
	StaleChannelFlushMs *float64 `default:"10000" json:"staleChannelFlushMs"`
	// Shared secrets to be provided by any Splunk forwarder. If empty, unauthorized access is permitted.
	AuthTokens []AuthTokens `json:"authTokens,omitempty"`
	// The highest S2S protocol version to advertise during handshake
	MaxS2Sversion *MaxS2SVersion `default:"v3" json:"maxS2Sversion"`
	Description   *string        `json:"description,omitempty"`
	// Event Breakers will determine events' time zone from UF-provided metadata, when TZ can't be inferred from the raw event
	UseFwdTimezone *bool `default:"true" json:"useFwdTimezone"`
	// Drop Splunk control fields such as `crcSalt` and `_savedPort`. If disabled, control fields are stored in the internal field `__ctrlFields`.
	DropControlFields *bool `default:"true" json:"dropControlFields"`
	// Extract and process Splunk-generated metrics as Cribl metrics
	ExtractMetrics *bool `default:"false" json:"extractMetrics"`
	// Controls whether to support reading compressed data from a forwarder. Select 'Automatic' to match the forwarder's configuration, or 'Disabled' to reject compressed connections.
	Compress *string `default:"disabled" json:"compress"`
}

func (i InputSplunk) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputSplunk) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *InputSplunk) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *InputSplunk) GetType() *InputSplunkType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *InputSplunk) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputSplunk) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputSplunk) GetSendToRoutes() *bool {
	if o == nil {
		return nil
	}
	return o.SendToRoutes
}

func (o *InputSplunk) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *InputSplunk) GetPqEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.PqEnabled
}

func (o *InputSplunk) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *InputSplunk) GetConnections() []InputSplunkConnections {
	if o == nil {
		return nil
	}
	return o.Connections
}

func (o *InputSplunk) GetPq() *InputSplunkPq {
	if o == nil {
		return nil
	}
	return o.Pq
}

func (o *InputSplunk) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *InputSplunk) GetPort() float64 {
	if o == nil {
		return 0.0
	}
	return o.Port
}

func (o *InputSplunk) GetTLS() *InputSplunkTLSSettingsServerSide {
	if o == nil {
		return nil
	}
	return o.TLS
}

func (o *InputSplunk) GetIPWhitelistRegex() *string {
	if o == nil {
		return nil
	}
	return o.IPWhitelistRegex
}

func (o *InputSplunk) GetMaxActiveCxn() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxActiveCxn
}

func (o *InputSplunk) GetSocketIdleTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.SocketIdleTimeout
}

func (o *InputSplunk) GetSocketEndingMaxWait() *float64 {
	if o == nil {
		return nil
	}
	return o.SocketEndingMaxWait
}

func (o *InputSplunk) GetSocketMaxLifespan() *float64 {
	if o == nil {
		return nil
	}
	return o.SocketMaxLifespan
}

func (o *InputSplunk) GetEnableProxyHeader() *bool {
	if o == nil {
		return nil
	}
	return o.EnableProxyHeader
}

func (o *InputSplunk) GetMetadata() []InputSplunkMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *InputSplunk) GetBreakerRulesets() []string {
	if o == nil {
		return nil
	}
	return o.BreakerRulesets
}

func (o *InputSplunk) GetStaleChannelFlushMs() *float64 {
	if o == nil {
		return nil
	}
	return o.StaleChannelFlushMs
}

func (o *InputSplunk) GetAuthTokens() []AuthTokens {
	if o == nil {
		return nil
	}
	return o.AuthTokens
}

func (o *InputSplunk) GetMaxS2Sversion() *MaxS2SVersion {
	if o == nil {
		return nil
	}
	return o.MaxS2Sversion
}

func (o *InputSplunk) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *InputSplunk) GetUseFwdTimezone() *bool {
	if o == nil {
		return nil
	}
	return o.UseFwdTimezone
}

func (o *InputSplunk) GetDropControlFields() *bool {
	if o == nil {
		return nil
	}
	return o.DropControlFields
}

func (o *InputSplunk) GetExtractMetrics() *bool {
	if o == nil {
		return nil
	}
	return o.ExtractMetrics
}

func (o *InputSplunk) GetCompress() *string {
	if o == nil {
		return nil
	}
	return o.Compress
}
