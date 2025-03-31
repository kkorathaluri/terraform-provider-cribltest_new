// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

type OutputConfluentCloudType string

const (
	OutputConfluentCloudTypeConfluentCloud OutputConfluentCloudType = "confluent_cloud"
)

func (e OutputConfluentCloudType) ToPointer() *OutputConfluentCloudType {
	return &e
}
func (e *OutputConfluentCloudType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "confluent_cloud":
		*e = OutputConfluentCloudType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputConfluentCloudType: %v", v)
	}
}

// OutputConfluentCloudMinimumTLSVersion - Minimum TLS version to use when connecting
type OutputConfluentCloudMinimumTLSVersion string

const (
	OutputConfluentCloudMinimumTLSVersionTlSv1  OutputConfluentCloudMinimumTLSVersion = "TLSv1"
	OutputConfluentCloudMinimumTLSVersionTlSv11 OutputConfluentCloudMinimumTLSVersion = "TLSv1.1"
	OutputConfluentCloudMinimumTLSVersionTlSv12 OutputConfluentCloudMinimumTLSVersion = "TLSv1.2"
	OutputConfluentCloudMinimumTLSVersionTlSv13 OutputConfluentCloudMinimumTLSVersion = "TLSv1.3"
)

func (e OutputConfluentCloudMinimumTLSVersion) ToPointer() *OutputConfluentCloudMinimumTLSVersion {
	return &e
}
func (e *OutputConfluentCloudMinimumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = OutputConfluentCloudMinimumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputConfluentCloudMinimumTLSVersion: %v", v)
	}
}

// OutputConfluentCloudMaximumTLSVersion - Maximum TLS version to use when connecting
type OutputConfluentCloudMaximumTLSVersion string

const (
	OutputConfluentCloudMaximumTLSVersionTlSv1  OutputConfluentCloudMaximumTLSVersion = "TLSv1"
	OutputConfluentCloudMaximumTLSVersionTlSv11 OutputConfluentCloudMaximumTLSVersion = "TLSv1.1"
	OutputConfluentCloudMaximumTLSVersionTlSv12 OutputConfluentCloudMaximumTLSVersion = "TLSv1.2"
	OutputConfluentCloudMaximumTLSVersionTlSv13 OutputConfluentCloudMaximumTLSVersion = "TLSv1.3"
)

func (e OutputConfluentCloudMaximumTLSVersion) ToPointer() *OutputConfluentCloudMaximumTLSVersion {
	return &e
}
func (e *OutputConfluentCloudMaximumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = OutputConfluentCloudMaximumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputConfluentCloudMaximumTLSVersion: %v", v)
	}
}

type OutputConfluentCloudTLSSettingsClientSide struct {
	Disabled *bool `default:"false" json:"disabled"`
	// Reject certs that are not authorized by a CA in the CA certificate path, or by another
	//                     trusted CA (e.g., the system's CA). Defaults to Yes. Overrides the toggle from Advanced Settings, when also present.
	RejectUnauthorized *bool `default:"true" json:"rejectUnauthorized"`
	// Server name for the SNI (Server Name Indication) TLS extension. It must be a host name, and not an IP address.
	Servername *string `json:"servername,omitempty"`
	// The name of the predefined certificate.
	CertificateName *string `json:"certificateName,omitempty"`
	// Path on client in which to find CA certificates to verify the server's cert. PEM format. Can reference $ENV_VARS.
	CaPath *string `json:"caPath,omitempty"`
	// Path on client in which to find the private key to use. PEM format. Can reference $ENV_VARS.
	PrivKeyPath *string `json:"privKeyPath,omitempty"`
	// Path on client in which to find certificates to use. PEM format. Can reference $ENV_VARS.
	CertPath *string `json:"certPath,omitempty"`
	// Passphrase to use to decrypt private key.
	Passphrase *string `json:"passphrase,omitempty"`
	// Minimum TLS version to use when connecting
	MinVersion *OutputConfluentCloudMinimumTLSVersion `json:"minVersion,omitempty"`
	// Maximum TLS version to use when connecting
	MaxVersion *OutputConfluentCloudMaximumTLSVersion `json:"maxVersion,omitempty"`
}

func (o OutputConfluentCloudTLSSettingsClientSide) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputConfluentCloudTLSSettingsClientSide) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputConfluentCloudTLSSettingsClientSide) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *OutputConfluentCloudTLSSettingsClientSide) GetRejectUnauthorized() *bool {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *OutputConfluentCloudTLSSettingsClientSide) GetServername() *string {
	if o == nil {
		return nil
	}
	return o.Servername
}

func (o *OutputConfluentCloudTLSSettingsClientSide) GetCertificateName() *string {
	if o == nil {
		return nil
	}
	return o.CertificateName
}

func (o *OutputConfluentCloudTLSSettingsClientSide) GetCaPath() *string {
	if o == nil {
		return nil
	}
	return o.CaPath
}

func (o *OutputConfluentCloudTLSSettingsClientSide) GetPrivKeyPath() *string {
	if o == nil {
		return nil
	}
	return o.PrivKeyPath
}

func (o *OutputConfluentCloudTLSSettingsClientSide) GetCertPath() *string {
	if o == nil {
		return nil
	}
	return o.CertPath
}

func (o *OutputConfluentCloudTLSSettingsClientSide) GetPassphrase() *string {
	if o == nil {
		return nil
	}
	return o.Passphrase
}

func (o *OutputConfluentCloudTLSSettingsClientSide) GetMinVersion() *OutputConfluentCloudMinimumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MinVersion
}

func (o *OutputConfluentCloudTLSSettingsClientSide) GetMaxVersion() *OutputConfluentCloudMaximumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MaxVersion
}

// OutputConfluentCloudAcknowledgments - Control the number of required acknowledgments.
type OutputConfluentCloudAcknowledgments int64

const (
	OutputConfluentCloudAcknowledgmentsOne    OutputConfluentCloudAcknowledgments = 1
	OutputConfluentCloudAcknowledgmentsZero   OutputConfluentCloudAcknowledgments = 0
	OutputConfluentCloudAcknowledgmentsMinus1 OutputConfluentCloudAcknowledgments = -1
)

func (e OutputConfluentCloudAcknowledgments) ToPointer() *OutputConfluentCloudAcknowledgments {
	return &e
}
func (e *OutputConfluentCloudAcknowledgments) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 0:
		fallthrough
	case -1:
		*e = OutputConfluentCloudAcknowledgments(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputConfluentCloudAcknowledgments: %v", v)
	}
}

// OutputConfluentCloudRecordDataFormat - Format to use to serialize events before writing to Kafka.
type OutputConfluentCloudRecordDataFormat string

const (
	OutputConfluentCloudRecordDataFormatJSON     OutputConfluentCloudRecordDataFormat = "json"
	OutputConfluentCloudRecordDataFormatRaw      OutputConfluentCloudRecordDataFormat = "raw"
	OutputConfluentCloudRecordDataFormatProtobuf OutputConfluentCloudRecordDataFormat = "protobuf"
)

func (e OutputConfluentCloudRecordDataFormat) ToPointer() *OutputConfluentCloudRecordDataFormat {
	return &e
}
func (e *OutputConfluentCloudRecordDataFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "json":
		fallthrough
	case "raw":
		fallthrough
	case "protobuf":
		*e = OutputConfluentCloudRecordDataFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputConfluentCloudRecordDataFormat: %v", v)
	}
}

// OutputConfluentCloudCompression - Codec to use to compress the data before sending to Kafka
type OutputConfluentCloudCompression string

const (
	OutputConfluentCloudCompressionNone   OutputConfluentCloudCompression = "none"
	OutputConfluentCloudCompressionGzip   OutputConfluentCloudCompression = "gzip"
	OutputConfluentCloudCompressionSnappy OutputConfluentCloudCompression = "snappy"
	OutputConfluentCloudCompressionLz4    OutputConfluentCloudCompression = "lz4"
)

func (e OutputConfluentCloudCompression) ToPointer() *OutputConfluentCloudCompression {
	return &e
}
func (e *OutputConfluentCloudCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		fallthrough
	case "snappy":
		fallthrough
	case "lz4":
		*e = OutputConfluentCloudCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputConfluentCloudCompression: %v", v)
	}
}

// OutputConfluentCloudAuth - Credentials to use when authenticating with the schema registry using basic HTTP authentication
type OutputConfluentCloudAuth struct {
	// Enable authentication
	Disabled *bool `default:"true" json:"disabled"`
	// Select or create a secret that references your credentials
	CredentialsSecret *string `json:"credentialsSecret,omitempty"`
}

func (o OutputConfluentCloudAuth) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputConfluentCloudAuth) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputConfluentCloudAuth) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *OutputConfluentCloudAuth) GetCredentialsSecret() *string {
	if o == nil {
		return nil
	}
	return o.CredentialsSecret
}

// OutputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion - Minimum TLS version to use when connecting
type OutputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion string

const (
	OutputConfluentCloudKafkaSchemaRegistryMinimumTLSVersionTlSv1  OutputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion = "TLSv1"
	OutputConfluentCloudKafkaSchemaRegistryMinimumTLSVersionTlSv11 OutputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion = "TLSv1.1"
	OutputConfluentCloudKafkaSchemaRegistryMinimumTLSVersionTlSv12 OutputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion = "TLSv1.2"
	OutputConfluentCloudKafkaSchemaRegistryMinimumTLSVersionTlSv13 OutputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion = "TLSv1.3"
)

func (e OutputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion) ToPointer() *OutputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion {
	return &e
}
func (e *OutputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = OutputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion: %v", v)
	}
}

// OutputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion - Maximum TLS version to use when connecting
type OutputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion string

const (
	OutputConfluentCloudKafkaSchemaRegistryMaximumTLSVersionTlSv1  OutputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion = "TLSv1"
	OutputConfluentCloudKafkaSchemaRegistryMaximumTLSVersionTlSv11 OutputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion = "TLSv1.1"
	OutputConfluentCloudKafkaSchemaRegistryMaximumTLSVersionTlSv12 OutputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion = "TLSv1.2"
	OutputConfluentCloudKafkaSchemaRegistryMaximumTLSVersionTlSv13 OutputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion = "TLSv1.3"
)

func (e OutputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion) ToPointer() *OutputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion {
	return &e
}
func (e *OutputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = OutputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion: %v", v)
	}
}

type OutputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide struct {
	Disabled *bool `default:"true" json:"disabled"`
	// Reject certs that are not authorized by a CA in the CA certificate path, or by another
	//                     trusted CA (e.g., the system's CA). Defaults to Yes. Overrides the toggle from Advanced Settings, when also present.
	RejectUnauthorized *bool `default:"true" json:"rejectUnauthorized"`
	// Server name for the SNI (Server Name Indication) TLS extension. It must be a host name, and not an IP address.
	Servername *string `json:"servername,omitempty"`
	// The name of the predefined certificate.
	CertificateName *string `json:"certificateName,omitempty"`
	// Path on client in which to find CA certificates to verify the server's cert. PEM format. Can reference $ENV_VARS.
	CaPath *string `json:"caPath,omitempty"`
	// Path on client in which to find the private key to use. PEM format. Can reference $ENV_VARS.
	PrivKeyPath *string `json:"privKeyPath,omitempty"`
	// Path on client in which to find certificates to use. PEM format. Can reference $ENV_VARS.
	CertPath *string `json:"certPath,omitempty"`
	// Passphrase to use to decrypt private key.
	Passphrase *string `json:"passphrase,omitempty"`
	// Minimum TLS version to use when connecting
	MinVersion *OutputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion `json:"minVersion,omitempty"`
	// Maximum TLS version to use when connecting
	MaxVersion *OutputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion `json:"maxVersion,omitempty"`
}

func (o OutputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *OutputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide) GetRejectUnauthorized() *bool {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *OutputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide) GetServername() *string {
	if o == nil {
		return nil
	}
	return o.Servername
}

func (o *OutputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide) GetCertificateName() *string {
	if o == nil {
		return nil
	}
	return o.CertificateName
}

func (o *OutputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide) GetCaPath() *string {
	if o == nil {
		return nil
	}
	return o.CaPath
}

func (o *OutputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide) GetPrivKeyPath() *string {
	if o == nil {
		return nil
	}
	return o.PrivKeyPath
}

func (o *OutputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide) GetCertPath() *string {
	if o == nil {
		return nil
	}
	return o.CertPath
}

func (o *OutputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide) GetPassphrase() *string {
	if o == nil {
		return nil
	}
	return o.Passphrase
}

func (o *OutputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide) GetMinVersion() *OutputConfluentCloudKafkaSchemaRegistryMinimumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MinVersion
}

func (o *OutputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide) GetMaxVersion() *OutputConfluentCloudKafkaSchemaRegistryMaximumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MaxVersion
}

type OutputConfluentCloudKafkaSchemaRegistryAuthentication struct {
	// Enable Schema Registry
	Disabled *bool `default:"true" json:"disabled"`
	// URL for accessing the Confluent Schema Registry. Example: http://localhost:8081. To connect over TLS, use https instead of http.
	SchemaRegistryURL *string `default:"http://localhost:8081" json:"schemaRegistryURL"`
	// Maximum time to wait for a Schema Registry connection to complete successfully
	ConnectionTimeout *float64 `default:"30000" json:"connectionTimeout"`
	// Maximum time to wait for the Schema Registry to respond to a request
	RequestTimeout *float64 `default:"30000" json:"requestTimeout"`
	// Maximum number of times to try fetching schemas from the Schema Registry
	MaxRetries *float64 `default:"1" json:"maxRetries"`
	// Credentials to use when authenticating with the schema registry using basic HTTP authentication
	Auth *OutputConfluentCloudAuth                                     `json:"auth,omitempty"`
	TLS  *OutputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide `json:"tls,omitempty"`
	// Used when __keySchemaIdOut is not present, to transform key values, leave blank if key transformation is not required by default.
	DefaultKeySchemaID *float64 `json:"defaultKeySchemaId,omitempty"`
	// Used when __valueSchemaIdOut is not present, to transform _raw, leave blank if value transformation is not required by default.
	DefaultValueSchemaID *float64 `json:"defaultValueSchemaId,omitempty"`
}

func (o OutputConfluentCloudKafkaSchemaRegistryAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputConfluentCloudKafkaSchemaRegistryAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputConfluentCloudKafkaSchemaRegistryAuthentication) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *OutputConfluentCloudKafkaSchemaRegistryAuthentication) GetSchemaRegistryURL() *string {
	if o == nil {
		return nil
	}
	return o.SchemaRegistryURL
}

func (o *OutputConfluentCloudKafkaSchemaRegistryAuthentication) GetConnectionTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.ConnectionTimeout
}

func (o *OutputConfluentCloudKafkaSchemaRegistryAuthentication) GetRequestTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.RequestTimeout
}

func (o *OutputConfluentCloudKafkaSchemaRegistryAuthentication) GetMaxRetries() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRetries
}

func (o *OutputConfluentCloudKafkaSchemaRegistryAuthentication) GetAuth() *OutputConfluentCloudAuth {
	if o == nil {
		return nil
	}
	return o.Auth
}

func (o *OutputConfluentCloudKafkaSchemaRegistryAuthentication) GetTLS() *OutputConfluentCloudKafkaSchemaRegistryTLSSettingsClientSide {
	if o == nil {
		return nil
	}
	return o.TLS
}

func (o *OutputConfluentCloudKafkaSchemaRegistryAuthentication) GetDefaultKeySchemaID() *float64 {
	if o == nil {
		return nil
	}
	return o.DefaultKeySchemaID
}

func (o *OutputConfluentCloudKafkaSchemaRegistryAuthentication) GetDefaultValueSchemaID() *float64 {
	if o == nil {
		return nil
	}
	return o.DefaultValueSchemaID
}

// OutputConfluentCloudSASLMechanism - SASL authentication mechanism to use.
type OutputConfluentCloudSASLMechanism string

const (
	OutputConfluentCloudSASLMechanismPlain       OutputConfluentCloudSASLMechanism = "plain"
	OutputConfluentCloudSASLMechanismScramSha256 OutputConfluentCloudSASLMechanism = "scram-sha-256"
	OutputConfluentCloudSASLMechanismScramSha512 OutputConfluentCloudSASLMechanism = "scram-sha-512"
	OutputConfluentCloudSASLMechanismKerberos    OutputConfluentCloudSASLMechanism = "kerberos"
)

func (e OutputConfluentCloudSASLMechanism) ToPointer() *OutputConfluentCloudSASLMechanism {
	return &e
}
func (e *OutputConfluentCloudSASLMechanism) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "plain":
		fallthrough
	case "scram-sha-256":
		fallthrough
	case "scram-sha-512":
		fallthrough
	case "kerberos":
		*e = OutputConfluentCloudSASLMechanism(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputConfluentCloudSASLMechanism: %v", v)
	}
}

// OutputConfluentCloudAuthentication - Authentication parameters to use when connecting to brokers. Using TLS is highly recommended.
type OutputConfluentCloudAuthentication struct {
	// Enable Authentication
	Disabled *bool `default:"true" json:"disabled"`
	// SASL authentication mechanism to use.
	Mechanism *OutputConfluentCloudSASLMechanism `default:"plain" json:"mechanism"`
}

func (o OutputConfluentCloudAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputConfluentCloudAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputConfluentCloudAuthentication) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *OutputConfluentCloudAuthentication) GetMechanism() *OutputConfluentCloudSASLMechanism {
	if o == nil {
		return nil
	}
	return o.Mechanism
}

// OutputConfluentCloudBackpressureBehavior - Whether to block, drop, or queue events when all receivers are exerting backpressure.
type OutputConfluentCloudBackpressureBehavior string

const (
	OutputConfluentCloudBackpressureBehaviorBlock OutputConfluentCloudBackpressureBehavior = "block"
	OutputConfluentCloudBackpressureBehaviorDrop  OutputConfluentCloudBackpressureBehavior = "drop"
	OutputConfluentCloudBackpressureBehaviorQueue OutputConfluentCloudBackpressureBehavior = "queue"
)

func (e OutputConfluentCloudBackpressureBehavior) ToPointer() *OutputConfluentCloudBackpressureBehavior {
	return &e
}
func (e *OutputConfluentCloudBackpressureBehavior) UnmarshalJSON(data []byte) error {
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
		*e = OutputConfluentCloudBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputConfluentCloudBackpressureBehavior: %v", v)
	}
}

// OutputConfluentCloudPqCompressCompression - Codec to use to compress the persisted data.
type OutputConfluentCloudPqCompressCompression string

const (
	OutputConfluentCloudPqCompressCompressionNone OutputConfluentCloudPqCompressCompression = "none"
	OutputConfluentCloudPqCompressCompressionGzip OutputConfluentCloudPqCompressCompression = "gzip"
)

func (e OutputConfluentCloudPqCompressCompression) ToPointer() *OutputConfluentCloudPqCompressCompression {
	return &e
}
func (e *OutputConfluentCloudPqCompressCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputConfluentCloudPqCompressCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputConfluentCloudPqCompressCompression: %v", v)
	}
}

// OutputConfluentCloudQueueFullBehavior - Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputConfluentCloudQueueFullBehavior string

const (
	OutputConfluentCloudQueueFullBehaviorBlock OutputConfluentCloudQueueFullBehavior = "block"
	OutputConfluentCloudQueueFullBehaviorDrop  OutputConfluentCloudQueueFullBehavior = "drop"
)

func (e OutputConfluentCloudQueueFullBehavior) ToPointer() *OutputConfluentCloudQueueFullBehavior {
	return &e
}
func (e *OutputConfluentCloudQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputConfluentCloudQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputConfluentCloudQueueFullBehavior: %v", v)
	}
}

// OutputConfluentCloudMode - In Error mode, PQ writes events to the filesystem only when it detects a non-retryable Destination error. In Backpressure mode, PQ writes events to the filesystem when it detects backpressure from the Destination or when there are non-retryable Destination errors. In Always On mode, PQ always writes events to the filesystem.
type OutputConfluentCloudMode string

const (
	OutputConfluentCloudModeError        OutputConfluentCloudMode = "error"
	OutputConfluentCloudModeBackpressure OutputConfluentCloudMode = "backpressure"
	OutputConfluentCloudModeAlways       OutputConfluentCloudMode = "always"
)

func (e OutputConfluentCloudMode) ToPointer() *OutputConfluentCloudMode {
	return &e
}
func (e *OutputConfluentCloudMode) UnmarshalJSON(data []byte) error {
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
		*e = OutputConfluentCloudMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputConfluentCloudMode: %v", v)
	}
}

type OutputConfluentCloudPqControls struct {
}

type OutputConfluentCloud struct {
	// Unique ID for this output
	ID   *string                   `json:"id,omitempty"`
	Type *OutputConfluentCloudType `json:"type,omitempty"`
	// Pipeline to process data before sending out to this output
	Pipeline *string `json:"pipeline,omitempty"`
	// Fields to automatically add to events, such as cribl_pipe. Supports wildcards.
	SystemFields []string `json:"systemFields,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Tags for filtering and grouping in @{product}
	Streamtags []string `json:"streamtags,omitempty"`
	// List of Confluent Cloud bootstrap servers to use, such as yourAccount.confluent.cloud:9092.
	Brokers []string                                   `json:"brokers"`
	TLS     *OutputConfluentCloudTLSSettingsClientSide `json:"tls,omitempty"`
	// The topic to publish events to. Can be overridden using the __topicOut field.
	Topic string `json:"topic"`
	// Control the number of required acknowledgments.
	Ack *OutputConfluentCloudAcknowledgments `default:"1" json:"ack"`
	// Format to use to serialize events before writing to Kafka.
	Format *OutputConfluentCloudRecordDataFormat `default:"json" json:"format"`
	// Codec to use to compress the data before sending to Kafka
	Compression *OutputConfluentCloudCompression `default:"gzip" json:"compression"`
	// Maximum size of each record batch before compression. The value must not exceed the Kafka brokers' message.max.bytes setting.
	MaxRecordSizeKB *float64 `default:"768" json:"maxRecordSizeKB"`
	// The maximum number of events you want the Destination to allow in a batch before forcing a flush
	FlushEventCount *float64 `default:"1000" json:"flushEventCount"`
	// The maximum amount of time you want the Destination to wait before forcing a flush. Shorter intervals tend to result in smaller batches being sent.
	FlushPeriodSec      *float64                                               `default:"1" json:"flushPeriodSec"`
	KafkaSchemaRegistry *OutputConfluentCloudKafkaSchemaRegistryAuthentication `json:"kafkaSchemaRegistry,omitempty"`
	// Maximum time to wait for a connection to complete successfully
	ConnectionTimeout *float64 `default:"10000" json:"connectionTimeout"`
	// Maximum time to wait for Kafka to respond to a request
	RequestTimeout *float64 `default:"60000" json:"requestTimeout"`
	// If messages are failing, you can set the maximum number of retries as high as 100 to prevent loss of data.
	MaxRetries *float64 `default:"5" json:"maxRetries"`
	// The maximum wait time for a retry, in milliseconds. Default (and minimum) is 30,000 ms (30 seconds); maximum is 180,000 ms (180 seconds).
	MaxBackOff *float64 `default:"30000" json:"maxBackOff"`
	// Initial value used to calculate the retry, in milliseconds. Maximum is 600,000 ms (10 minutes).
	InitialBackoff *float64 `default:"300" json:"initialBackoff"`
	// Set the backoff multiplier (2-20) to control the retry frequency for failed messages. For faster retries, use a lower multiplier. For slower retries with more delay between attempts, use a higher multiplier. The multiplier is used in an exponential backoff formula; see the Kafka [documentation](https://kafka.js.org/docs/retry-detailed) for details.
	BackoffRate *float64 `default:"2" json:"backoffRate"`
	// Maximum time to wait for Kafka to respond to an authentication request
	AuthenticationTimeout *float64 `default:"10000" json:"authenticationTimeout"`
	// Specifies a time window during which @{product} can reauthenticate if needed. Creates the window measuring backwards from the moment when credentials are set to expire.
	ReauthenticationThreshold *float64 `default:"10000" json:"reauthenticationThreshold"`
	// Authentication parameters to use when connecting to brokers. Using TLS is highly recommended.
	Sasl *OutputConfluentCloudAuthentication `json:"sasl,omitempty"`
	// Whether to block, drop, or queue events when all receivers are exerting backpressure.
	OnBackpressure *OutputConfluentCloudBackpressureBehavior `default:"block" json:"onBackpressure"`
	Description    *string                                   `json:"description,omitempty"`
	// Select a set of Protobuf definitions for the events you want to send
	ProtobufLibraryID *string `json:"protobufLibraryId,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	PqMaxFileSize *string `default:"1 MB" json:"pqMaxFileSize"`
	// The maximum disk space that the queue can consume (as an average per Worker Process) before queueing stops. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `default:"5GB" json:"pqMaxSize"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/<output-id>.
	PqPath *string `default:"\\$CRIBL_HOME/state/queues" json:"pqPath"`
	// Codec to use to compress the persisted data.
	PqCompress *OutputConfluentCloudPqCompressCompression `default:"none" json:"pqCompress"`
	// Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputConfluentCloudQueueFullBehavior `default:"block" json:"pqOnBackpressure"`
	// In Error mode, PQ writes events to the filesystem only when it detects a non-retryable Destination error. In Backpressure mode, PQ writes events to the filesystem when it detects backpressure from the Destination or when there are non-retryable Destination errors. In Always On mode, PQ always writes events to the filesystem.
	PqMode     *OutputConfluentCloudMode       `default:"error" json:"pqMode"`
	PqControls *OutputConfluentCloudPqControls `json:"pqControls,omitempty"`
}

func (o OutputConfluentCloud) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputConfluentCloud) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *OutputConfluentCloud) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OutputConfluentCloud) GetType() *OutputConfluentCloudType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *OutputConfluentCloud) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *OutputConfluentCloud) GetSystemFields() []string {
	if o == nil {
		return nil
	}
	return o.SystemFields
}

func (o *OutputConfluentCloud) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *OutputConfluentCloud) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *OutputConfluentCloud) GetBrokers() []string {
	if o == nil {
		return []string{}
	}
	return o.Brokers
}

func (o *OutputConfluentCloud) GetTLS() *OutputConfluentCloudTLSSettingsClientSide {
	if o == nil {
		return nil
	}
	return o.TLS
}

func (o *OutputConfluentCloud) GetTopic() string {
	if o == nil {
		return ""
	}
	return o.Topic
}

func (o *OutputConfluentCloud) GetAck() *OutputConfluentCloudAcknowledgments {
	if o == nil {
		return nil
	}
	return o.Ack
}

func (o *OutputConfluentCloud) GetFormat() *OutputConfluentCloudRecordDataFormat {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *OutputConfluentCloud) GetCompression() *OutputConfluentCloudCompression {
	if o == nil {
		return nil
	}
	return o.Compression
}

func (o *OutputConfluentCloud) GetMaxRecordSizeKB() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRecordSizeKB
}

func (o *OutputConfluentCloud) GetFlushEventCount() *float64 {
	if o == nil {
		return nil
	}
	return o.FlushEventCount
}

func (o *OutputConfluentCloud) GetFlushPeriodSec() *float64 {
	if o == nil {
		return nil
	}
	return o.FlushPeriodSec
}

func (o *OutputConfluentCloud) GetKafkaSchemaRegistry() *OutputConfluentCloudKafkaSchemaRegistryAuthentication {
	if o == nil {
		return nil
	}
	return o.KafkaSchemaRegistry
}

func (o *OutputConfluentCloud) GetConnectionTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.ConnectionTimeout
}

func (o *OutputConfluentCloud) GetRequestTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.RequestTimeout
}

func (o *OutputConfluentCloud) GetMaxRetries() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRetries
}

func (o *OutputConfluentCloud) GetMaxBackOff() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBackOff
}

func (o *OutputConfluentCloud) GetInitialBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialBackoff
}

func (o *OutputConfluentCloud) GetBackoffRate() *float64 {
	if o == nil {
		return nil
	}
	return o.BackoffRate
}

func (o *OutputConfluentCloud) GetAuthenticationTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.AuthenticationTimeout
}

func (o *OutputConfluentCloud) GetReauthenticationThreshold() *float64 {
	if o == nil {
		return nil
	}
	return o.ReauthenticationThreshold
}

func (o *OutputConfluentCloud) GetSasl() *OutputConfluentCloudAuthentication {
	if o == nil {
		return nil
	}
	return o.Sasl
}

func (o *OutputConfluentCloud) GetOnBackpressure() *OutputConfluentCloudBackpressureBehavior {
	if o == nil {
		return nil
	}
	return o.OnBackpressure
}

func (o *OutputConfluentCloud) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *OutputConfluentCloud) GetProtobufLibraryID() *string {
	if o == nil {
		return nil
	}
	return o.ProtobufLibraryID
}

func (o *OutputConfluentCloud) GetPqMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxFileSize
}

func (o *OutputConfluentCloud) GetPqMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxSize
}

func (o *OutputConfluentCloud) GetPqPath() *string {
	if o == nil {
		return nil
	}
	return o.PqPath
}

func (o *OutputConfluentCloud) GetPqCompress() *OutputConfluentCloudPqCompressCompression {
	if o == nil {
		return nil
	}
	return o.PqCompress
}

func (o *OutputConfluentCloud) GetPqOnBackpressure() *OutputConfluentCloudQueueFullBehavior {
	if o == nil {
		return nil
	}
	return o.PqOnBackpressure
}

func (o *OutputConfluentCloud) GetPqMode() *OutputConfluentCloudMode {
	if o == nil {
		return nil
	}
	return o.PqMode
}

func (o *OutputConfluentCloud) GetPqControls() *OutputConfluentCloudPqControls {
	if o == nil {
		return nil
	}
	return o.PqControls
}
