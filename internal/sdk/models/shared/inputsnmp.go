// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

type InputSnmpType string

const (
	InputSnmpTypeSnmp InputSnmpType = "snmp"
)

func (e InputSnmpType) ToPointer() *InputSnmpType {
	return &e
}
func (e *InputSnmpType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "snmp":
		*e = InputSnmpType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSnmpType: %v", v)
	}
}

type InputSnmpConnections struct {
	Pipeline *string `json:"pipeline,omitempty"`
	Output   string  `json:"output"`
}

func (o *InputSnmpConnections) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputSnmpConnections) GetOutput() string {
	if o == nil {
		return ""
	}
	return o.Output
}

// InputSnmpMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputSnmpMode string

const (
	InputSnmpModeSmart  InputSnmpMode = "smart"
	InputSnmpModeAlways InputSnmpMode = "always"
)

func (e InputSnmpMode) ToPointer() *InputSnmpMode {
	return &e
}
func (e *InputSnmpMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputSnmpMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSnmpMode: %v", v)
	}
}

// InputSnmpCompression - Codec to use to compress the persisted data
type InputSnmpCompression string

const (
	InputSnmpCompressionNone InputSnmpCompression = "none"
	InputSnmpCompressionGzip InputSnmpCompression = "gzip"
)

func (e InputSnmpCompression) ToPointer() *InputSnmpCompression {
	return &e
}
func (e *InputSnmpCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputSnmpCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputSnmpCompression: %v", v)
	}
}

type InputSnmpPq struct {
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputSnmpMode `default:"always" json:"mode"`
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
	Compress *InputSnmpCompression `default:"none" json:"compress"`
}

func (i InputSnmpPq) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputSnmpPq) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputSnmpPq) GetMode() *InputSnmpMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputSnmpPq) GetMaxBufferSize() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBufferSize
}

func (o *InputSnmpPq) GetCommitFrequency() *float64 {
	if o == nil {
		return nil
	}
	return o.CommitFrequency
}

func (o *InputSnmpPq) GetMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxFileSize
}

func (o *InputSnmpPq) GetMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxSize
}

func (o *InputSnmpPq) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *InputSnmpPq) GetCompress() *InputSnmpCompression {
	if o == nil {
		return nil
	}
	return o.Compress
}

type AuthenticationProtocol string

const (
	AuthenticationProtocolNone   AuthenticationProtocol = "none"
	AuthenticationProtocolMd5    AuthenticationProtocol = "md5"
	AuthenticationProtocolSha    AuthenticationProtocol = "sha"
	AuthenticationProtocolSha224 AuthenticationProtocol = "sha224"
	AuthenticationProtocolSha256 AuthenticationProtocol = "sha256"
	AuthenticationProtocolSha384 AuthenticationProtocol = "sha384"
	AuthenticationProtocolSha512 AuthenticationProtocol = "sha512"
)

func (e AuthenticationProtocol) ToPointer() *AuthenticationProtocol {
	return &e
}
func (e *AuthenticationProtocol) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "md5":
		fallthrough
	case "sha":
		fallthrough
	case "sha224":
		fallthrough
	case "sha256":
		fallthrough
	case "sha384":
		fallthrough
	case "sha512":
		*e = AuthenticationProtocol(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AuthenticationProtocol: %v", v)
	}
}

type V3Users struct {
	Name         string                  `json:"name"`
	AuthProtocol *AuthenticationProtocol `default:"none" json:"authProtocol"`
	AuthKey      any                     `json:"authKey,omitempty"`
	PrivProtocol *string                 `default:"none" json:"privProtocol"`
}

func (v V3Users) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V3Users) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *V3Users) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *V3Users) GetAuthProtocol() *AuthenticationProtocol {
	if o == nil {
		return nil
	}
	return o.AuthProtocol
}

func (o *V3Users) GetAuthKey() any {
	if o == nil {
		return nil
	}
	return o.AuthKey
}

func (o *V3Users) GetPrivProtocol() *string {
	if o == nil {
		return nil
	}
	return o.PrivProtocol
}

// SNMPv3Authentication - Authentication parameters for SNMPv3 trap. Set the log level to debug if you are experiencing authentication or decryption issues.
type SNMPv3Authentication struct {
	V3AuthEnabled *bool `default:"false" json:"v3AuthEnabled"`
	// Pass through traps that don't match any of the configured users. @{product} will not attempt to decrypt these traps.
	AllowUnmatchedTrap *bool `default:"false" json:"allowUnmatchedTrap"`
	// User credentials for receiving v3 traps
	V3Users []V3Users `json:"v3Users,omitempty"`
}

func (s SNMPv3Authentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SNMPv3Authentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SNMPv3Authentication) GetV3AuthEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.V3AuthEnabled
}

func (o *SNMPv3Authentication) GetAllowUnmatchedTrap() *bool {
	if o == nil {
		return nil
	}
	return o.AllowUnmatchedTrap
}

func (o *SNMPv3Authentication) GetV3Users() []V3Users {
	if o == nil {
		return nil
	}
	return o.V3Users
}

type InputSnmpMetadata struct {
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

func (o *InputSnmpMetadata) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InputSnmpMetadata) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type InputSnmp struct {
	// Unique ID for this input
	ID       *string        `json:"id,omitempty"`
	Type     *InputSnmpType `json:"type,omitempty"`
	Disabled *bool          `default:"false" json:"disabled"`
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
	Connections []InputSnmpConnections `json:"connections,omitempty"`
	Pq          *InputSnmpPq           `json:"pq,omitempty"`
	// Address to bind on. For IPv4 (all addresses), use the default '0.0.0.0'. For IPv6, enter '::' (all addresses) or specify an IP address.
	Host *string `default:"0.0.0.0" json:"host"`
	// UDP port to receive SNMP traps on. Defaults to 162.
	Port *float64 `default:"162" json:"port"`
	// Authentication parameters for SNMPv3 trap. Set the log level to debug if you are experiencing authentication or decryption issues.
	SnmpV3Auth *SNMPv3Authentication `json:"snmpV3Auth,omitempty"`
	// Maximum number of events to buffer when downstream is blocking.
	MaxBufferSize *float64 `default:"1000" json:"maxBufferSize"`
	// Regex matching IP addresses that are allowed to send data
	IPWhitelistRegex *string `default:"/.*/" json:"ipWhitelistRegex"`
	// Fields to add to events from this input
	Metadata []InputSnmpMetadata `json:"metadata,omitempty"`
	// Optionally, set the SO_RCVBUF socket option for the UDP socket. This value tells the operating system how many bytes can be buffered in the kernel before events are dropped. Leave blank to use the OS default. Caution: Increasing this value will affect OS memory utilization.
	UDPSocketRxBufSize *float64 `json:"udpSocketRxBufSize,omitempty"`
	// If enabled, parses varbinds as an array of objects that include OID, value, and type
	VarbindsWithTypes *bool   `default:"false" json:"varbindsWithTypes"`
	Description       *string `json:"description,omitempty"`
}

func (i InputSnmp) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputSnmp) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *InputSnmp) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *InputSnmp) GetType() *InputSnmpType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *InputSnmp) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputSnmp) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputSnmp) GetSendToRoutes() *bool {
	if o == nil {
		return nil
	}
	return o.SendToRoutes
}

func (o *InputSnmp) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *InputSnmp) GetPqEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.PqEnabled
}

func (o *InputSnmp) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *InputSnmp) GetConnections() []InputSnmpConnections {
	if o == nil {
		return nil
	}
	return o.Connections
}

func (o *InputSnmp) GetPq() *InputSnmpPq {
	if o == nil {
		return nil
	}
	return o.Pq
}

func (o *InputSnmp) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *InputSnmp) GetPort() *float64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *InputSnmp) GetSnmpV3Auth() *SNMPv3Authentication {
	if o == nil {
		return nil
	}
	return o.SnmpV3Auth
}

func (o *InputSnmp) GetMaxBufferSize() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBufferSize
}

func (o *InputSnmp) GetIPWhitelistRegex() *string {
	if o == nil {
		return nil
	}
	return o.IPWhitelistRegex
}

func (o *InputSnmp) GetMetadata() []InputSnmpMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *InputSnmp) GetUDPSocketRxBufSize() *float64 {
	if o == nil {
		return nil
	}
	return o.UDPSocketRxBufSize
}

func (o *InputSnmp) GetVarbindsWithTypes() *bool {
	if o == nil {
		return nil
	}
	return o.VarbindsWithTypes
}

func (o *InputSnmp) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}
