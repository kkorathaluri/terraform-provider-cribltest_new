// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

type OutputClickHouseType string

const (
	OutputClickHouseTypeClickHouse OutputClickHouseType = "click_house"
)

func (e OutputClickHouseType) ToPointer() *OutputClickHouseType {
	return &e
}
func (e *OutputClickHouseType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "click_house":
		*e = OutputClickHouseType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputClickHouseType: %v", v)
	}
}

type OutputClickHouseAuthenticationType string

const (
	OutputClickHouseAuthenticationTypeNone               OutputClickHouseAuthenticationType = "none"
	OutputClickHouseAuthenticationTypeBasic              OutputClickHouseAuthenticationType = "basic"
	OutputClickHouseAuthenticationTypeCredentialsSecret  OutputClickHouseAuthenticationType = "credentialsSecret"
	OutputClickHouseAuthenticationTypeSslUserCertificate OutputClickHouseAuthenticationType = "sslUserCertificate"
	OutputClickHouseAuthenticationTypeToken              OutputClickHouseAuthenticationType = "token"
	OutputClickHouseAuthenticationTypeTextSecret         OutputClickHouseAuthenticationType = "textSecret"
	OutputClickHouseAuthenticationTypeOauth              OutputClickHouseAuthenticationType = "oauth"
)

func (e OutputClickHouseAuthenticationType) ToPointer() *OutputClickHouseAuthenticationType {
	return &e
}
func (e *OutputClickHouseAuthenticationType) UnmarshalJSON(data []byte) error {
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
	case "sslUserCertificate":
		fallthrough
	case "token":
		fallthrough
	case "textSecret":
		fallthrough
	case "oauth":
		*e = OutputClickHouseAuthenticationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputClickHouseAuthenticationType: %v", v)
	}
}

// OutputClickHouseFormat - Data format to use when sending data to ClickHouse. Defaults to JSON Compact.
type OutputClickHouseFormat string

const (
	OutputClickHouseFormatJSONCompactEachRowWithNames OutputClickHouseFormat = "json-compact-each-row-with-names"
	OutputClickHouseFormatJSONEachRow                 OutputClickHouseFormat = "json-each-row"
)

func (e OutputClickHouseFormat) ToPointer() *OutputClickHouseFormat {
	return &e
}
func (e *OutputClickHouseFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "json-compact-each-row-with-names":
		fallthrough
	case "json-each-row":
		*e = OutputClickHouseFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputClickHouseFormat: %v", v)
	}
}

// OutputClickHouseMappingType - How event fields are mapped to ClickHouse columns.
type OutputClickHouseMappingType string

const (
	OutputClickHouseMappingTypeAutomatic OutputClickHouseMappingType = "automatic"
	OutputClickHouseMappingTypeCustom    OutputClickHouseMappingType = "custom"
)

func (e OutputClickHouseMappingType) ToPointer() *OutputClickHouseMappingType {
	return &e
}
func (e *OutputClickHouseMappingType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "automatic":
		fallthrough
	case "custom":
		*e = OutputClickHouseMappingType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputClickHouseMappingType: %v", v)
	}
}

// OutputClickHouseMinimumTLSVersion - Minimum TLS version to use when connecting
type OutputClickHouseMinimumTLSVersion string

const (
	OutputClickHouseMinimumTLSVersionTlSv1  OutputClickHouseMinimumTLSVersion = "TLSv1"
	OutputClickHouseMinimumTLSVersionTlSv11 OutputClickHouseMinimumTLSVersion = "TLSv1.1"
	OutputClickHouseMinimumTLSVersionTlSv12 OutputClickHouseMinimumTLSVersion = "TLSv1.2"
	OutputClickHouseMinimumTLSVersionTlSv13 OutputClickHouseMinimumTLSVersion = "TLSv1.3"
)

func (e OutputClickHouseMinimumTLSVersion) ToPointer() *OutputClickHouseMinimumTLSVersion {
	return &e
}
func (e *OutputClickHouseMinimumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = OutputClickHouseMinimumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputClickHouseMinimumTLSVersion: %v", v)
	}
}

// OutputClickHouseMaximumTLSVersion - Maximum TLS version to use when connecting
type OutputClickHouseMaximumTLSVersion string

const (
	OutputClickHouseMaximumTLSVersionTlSv1  OutputClickHouseMaximumTLSVersion = "TLSv1"
	OutputClickHouseMaximumTLSVersionTlSv11 OutputClickHouseMaximumTLSVersion = "TLSv1.1"
	OutputClickHouseMaximumTLSVersionTlSv12 OutputClickHouseMaximumTLSVersion = "TLSv1.2"
	OutputClickHouseMaximumTLSVersionTlSv13 OutputClickHouseMaximumTLSVersion = "TLSv1.3"
)

func (e OutputClickHouseMaximumTLSVersion) ToPointer() *OutputClickHouseMaximumTLSVersion {
	return &e
}
func (e *OutputClickHouseMaximumTLSVersion) UnmarshalJSON(data []byte) error {
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
		*e = OutputClickHouseMaximumTLSVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputClickHouseMaximumTLSVersion: %v", v)
	}
}

type OutputClickHouseTLSSettingsClientSide struct {
	Disabled *bool `default:"true" json:"disabled"`
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
	MinVersion *OutputClickHouseMinimumTLSVersion `json:"minVersion,omitempty"`
	// Maximum TLS version to use when connecting
	MaxVersion *OutputClickHouseMaximumTLSVersion `json:"maxVersion,omitempty"`
}

func (o OutputClickHouseTLSSettingsClientSide) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputClickHouseTLSSettingsClientSide) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputClickHouseTLSSettingsClientSide) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *OutputClickHouseTLSSettingsClientSide) GetServername() *string {
	if o == nil {
		return nil
	}
	return o.Servername
}

func (o *OutputClickHouseTLSSettingsClientSide) GetCertificateName() *string {
	if o == nil {
		return nil
	}
	return o.CertificateName
}

func (o *OutputClickHouseTLSSettingsClientSide) GetCaPath() *string {
	if o == nil {
		return nil
	}
	return o.CaPath
}

func (o *OutputClickHouseTLSSettingsClientSide) GetPrivKeyPath() *string {
	if o == nil {
		return nil
	}
	return o.PrivKeyPath
}

func (o *OutputClickHouseTLSSettingsClientSide) GetCertPath() *string {
	if o == nil {
		return nil
	}
	return o.CertPath
}

func (o *OutputClickHouseTLSSettingsClientSide) GetPassphrase() *string {
	if o == nil {
		return nil
	}
	return o.Passphrase
}

func (o *OutputClickHouseTLSSettingsClientSide) GetMinVersion() *OutputClickHouseMinimumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MinVersion
}

func (o *OutputClickHouseTLSSettingsClientSide) GetMaxVersion() *OutputClickHouseMaximumTLSVersion {
	if o == nil {
		return nil
	}
	return o.MaxVersion
}

type OutputClickHouseExtraHTTPHeaders struct {
	// Field name
	Name *string `json:"name,omitempty"`
	// Field value
	Value string `json:"value"`
}

func (o *OutputClickHouseExtraHTTPHeaders) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *OutputClickHouseExtraHTTPHeaders) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// OutputClickHouseFailedRequestLoggingMode - Data to log when a request fails. All headers are redacted by default, unless listed as safe headers below.
type OutputClickHouseFailedRequestLoggingMode string

const (
	OutputClickHouseFailedRequestLoggingModePayload           OutputClickHouseFailedRequestLoggingMode = "payload"
	OutputClickHouseFailedRequestLoggingModePayloadAndHeaders OutputClickHouseFailedRequestLoggingMode = "payloadAndHeaders"
	OutputClickHouseFailedRequestLoggingModeNone              OutputClickHouseFailedRequestLoggingMode = "none"
)

func (e OutputClickHouseFailedRequestLoggingMode) ToPointer() *OutputClickHouseFailedRequestLoggingMode {
	return &e
}
func (e *OutputClickHouseFailedRequestLoggingMode) UnmarshalJSON(data []byte) error {
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
		*e = OutputClickHouseFailedRequestLoggingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputClickHouseFailedRequestLoggingMode: %v", v)
	}
}

type OutputClickHouseResponseRetrySettings struct {
	// The HTTP response status code that will trigger retries
	HTTPStatus float64 `json:"httpStatus"`
	// How long, in milliseconds, Cribl Stream should wait before initiating backoff. Maximum interval is 600,000 ms (10 minutes).
	InitialBackoff *float64 `default:"1000" json:"initialBackoff"`
	// Base for exponential backoff. A value of 2 (default) means Cribl Stream will retry after 2 seconds, then 4 seconds, then 8 seconds, etc.
	BackoffRate *float64 `default:"2" json:"backoffRate"`
	// The maximum backoff interval, in milliseconds, Cribl Stream should apply. Default (and minimum) is 10,000 ms (10 seconds); maximum is 180,000 ms (180 seconds).
	MaxBackoff *float64 `default:"10000" json:"maxBackoff"`
}

func (o OutputClickHouseResponseRetrySettings) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputClickHouseResponseRetrySettings) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputClickHouseResponseRetrySettings) GetHTTPStatus() float64 {
	if o == nil {
		return 0.0
	}
	return o.HTTPStatus
}

func (o *OutputClickHouseResponseRetrySettings) GetInitialBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialBackoff
}

func (o *OutputClickHouseResponseRetrySettings) GetBackoffRate() *float64 {
	if o == nil {
		return nil
	}
	return o.BackoffRate
}

func (o *OutputClickHouseResponseRetrySettings) GetMaxBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBackoff
}

type OutputClickHouseTimeoutRetrySettings struct {
	// Enable to retry on request timeout
	TimeoutRetry *bool `default:"false" json:"timeoutRetry"`
	// How long, in milliseconds, Cribl Stream should wait before initiating backoff. Maximum interval is 600,000 ms (10 minutes).
	InitialBackoff *float64 `default:"1000" json:"initialBackoff"`
	// Base for exponential backoff. A value of 2 (default) means Cribl Stream will retry after 2 seconds, then 4 seconds, then 8 seconds, etc.
	BackoffRate *float64 `default:"2" json:"backoffRate"`
	// The maximum backoff interval, in milliseconds, Cribl Stream should apply. Default (and minimum) is 10,000 ms (10 seconds); maximum is 180,000 ms (180 seconds).
	MaxBackoff *float64 `default:"10000" json:"maxBackoff"`
}

func (o OutputClickHouseTimeoutRetrySettings) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputClickHouseTimeoutRetrySettings) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OutputClickHouseTimeoutRetrySettings) GetTimeoutRetry() *bool {
	if o == nil {
		return nil
	}
	return o.TimeoutRetry
}

func (o *OutputClickHouseTimeoutRetrySettings) GetInitialBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialBackoff
}

func (o *OutputClickHouseTimeoutRetrySettings) GetBackoffRate() *float64 {
	if o == nil {
		return nil
	}
	return o.BackoffRate
}

func (o *OutputClickHouseTimeoutRetrySettings) GetMaxBackoff() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBackoff
}

// OutputClickHouseBackpressureBehavior - Whether to block, drop, or queue events when all receivers are exerting backpressure.
type OutputClickHouseBackpressureBehavior string

const (
	OutputClickHouseBackpressureBehaviorBlock OutputClickHouseBackpressureBehavior = "block"
	OutputClickHouseBackpressureBehaviorDrop  OutputClickHouseBackpressureBehavior = "drop"
	OutputClickHouseBackpressureBehaviorQueue OutputClickHouseBackpressureBehavior = "queue"
)

func (e OutputClickHouseBackpressureBehavior) ToPointer() *OutputClickHouseBackpressureBehavior {
	return &e
}
func (e *OutputClickHouseBackpressureBehavior) UnmarshalJSON(data []byte) error {
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
		*e = OutputClickHouseBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputClickHouseBackpressureBehavior: %v", v)
	}
}

type OutputClickHouseOauthParams struct {
	// OAuth parameter name
	Name string `json:"name"`
	// OAuth parameter value
	Value string `json:"value"`
}

func (o *OutputClickHouseOauthParams) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *OutputClickHouseOauthParams) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type OutputClickHouseOauthHeaders struct {
	// OAuth header name
	Name string `json:"name"`
	// OAuth header value
	Value string `json:"value"`
}

func (o *OutputClickHouseOauthHeaders) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *OutputClickHouseOauthHeaders) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type ColumnMappings struct {
	// Name of the column in ClickHouse that will store field value
	ColumnName string `json:"columnName"`
	// Type of the column in the ClickHouse database
	ColumnType *string `json:"columnType,omitempty"`
	// JavaScript expression to compute value to be inserted into ClickHouse table
	ColumnValueExpression string `json:"columnValueExpression"`
}

func (o *ColumnMappings) GetColumnName() string {
	if o == nil {
		return ""
	}
	return o.ColumnName
}

func (o *ColumnMappings) GetColumnType() *string {
	if o == nil {
		return nil
	}
	return o.ColumnType
}

func (o *ColumnMappings) GetColumnValueExpression() string {
	if o == nil {
		return ""
	}
	return o.ColumnValueExpression
}

// OutputClickHouseCompression - Codec to use to compress the persisted data.
type OutputClickHouseCompression string

const (
	OutputClickHouseCompressionNone OutputClickHouseCompression = "none"
	OutputClickHouseCompressionGzip OutputClickHouseCompression = "gzip"
)

func (e OutputClickHouseCompression) ToPointer() *OutputClickHouseCompression {
	return &e
}
func (e *OutputClickHouseCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputClickHouseCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputClickHouseCompression: %v", v)
	}
}

// OutputClickHouseQueueFullBehavior - Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputClickHouseQueueFullBehavior string

const (
	OutputClickHouseQueueFullBehaviorBlock OutputClickHouseQueueFullBehavior = "block"
	OutputClickHouseQueueFullBehaviorDrop  OutputClickHouseQueueFullBehavior = "drop"
)

func (e OutputClickHouseQueueFullBehavior) ToPointer() *OutputClickHouseQueueFullBehavior {
	return &e
}
func (e *OutputClickHouseQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputClickHouseQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputClickHouseQueueFullBehavior: %v", v)
	}
}

// OutputClickHouseMode - In Error mode, PQ writes events to the filesystem only when it detects a non-retryable Destination error. In Backpressure mode, PQ writes events to the filesystem when it detects backpressure from the Destination or when there are non-retryable Destination errors. In Always On mode, PQ always writes events to the filesystem.
type OutputClickHouseMode string

const (
	OutputClickHouseModeError        OutputClickHouseMode = "error"
	OutputClickHouseModeBackpressure OutputClickHouseMode = "backpressure"
	OutputClickHouseModeAlways       OutputClickHouseMode = "always"
)

func (e OutputClickHouseMode) ToPointer() *OutputClickHouseMode {
	return &e
}
func (e *OutputClickHouseMode) UnmarshalJSON(data []byte) error {
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
		*e = OutputClickHouseMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputClickHouseMode: %v", v)
	}
}

type OutputClickHousePqControls struct {
}

type OutputClickHouse struct {
	// Unique ID for this output
	ID   *string               `json:"id,omitempty"`
	Type *OutputClickHouseType `json:"type,omitempty"`
	// Pipeline to process data before sending out to this output
	Pipeline *string `json:"pipeline,omitempty"`
	// Fields to automatically add to events, such as cribl_pipe. Supports wildcards.
	SystemFields []string `json:"systemFields,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Tags for filtering and grouping in @{product}
	Streamtags []string `json:"streamtags,omitempty"`
	// URL of the ClickHouse instance. Example: http://localhost:8123/
	URL      string                              `json:"url"`
	AuthType *OutputClickHouseAuthenticationType `default:"none" json:"authType"`
	Database string                              `json:"database"`
	// Name of the ClickHouse table where data will be inserted. Name can contain letters (A-Z, a-z), numbers (0-9), and the character "_", and must start with either a letter or the character "_".
	TableName string `json:"tableName"`
	// Data format to use when sending data to ClickHouse. Defaults to JSON Compact.
	Format *OutputClickHouseFormat `default:"json-compact-each-row-with-names" json:"format"`
	// How event fields are mapped to ClickHouse columns.
	MappingType *OutputClickHouseMappingType `default:"automatic" json:"mappingType"`
	// Collect data into batches for later processing. Disable to write to a ClickHouse table immediately.
	AsyncInserts *bool                                  `default:"false" json:"asyncInserts"`
	TLS          *OutputClickHouseTLSSettingsClientSide `json:"tls,omitempty"`
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
	ExtraHTTPHeaders []OutputClickHouseExtraHTTPHeaders `json:"extraHttpHeaders,omitempty"`
	// Enables round-robin DNS lookup. When a DNS server returns multiple addresses, @{product} will cycle through them in the order returned. For optimal performance, consider enabling this setting for non-load balanced destinations.
	UseRoundRobinDNS *bool `default:"false" json:"useRoundRobinDns"`
	// Data to log when a request fails. All headers are redacted by default, unless listed as safe headers below.
	FailedRequestLoggingMode *OutputClickHouseFailedRequestLoggingMode `default:"none" json:"failedRequestLoggingMode"`
	// List of headers that are safe to log in plain text
	SafeHeaders []string `json:"safeHeaders,omitempty"`
	// Automatically retry after unsuccessful response status codes, such as 429 (Too Many Requests) or 503 (Service Unavailable).
	ResponseRetrySettings []OutputClickHouseResponseRetrySettings `json:"responseRetrySettings,omitempty"`
	TimeoutRetrySettings  *OutputClickHouseTimeoutRetrySettings   `json:"timeoutRetrySettings,omitempty"`
	// Honor any Retry-After header that specifies a delay (in seconds) no longer than 180 seconds after the retry request. @{product} limits the delay to 180 seconds, even if the Retry-After header specifies a longer delay. When enabled, takes precedence over user-configured retry options. When disabled, all Retry-After headers are ignored.
	ResponseHonorRetryAfterHeader *bool `default:"false" json:"responseHonorRetryAfterHeader"`
	// Log the most recent event that fails to match the table schema
	DumpFormatErrorsToDisk *bool `default:"false" json:"dumpFormatErrorsToDisk"`
	// Whether to block, drop, or queue events when all receivers are exerting backpressure.
	OnBackpressure *OutputClickHouseBackpressureBehavior `default:"block" json:"onBackpressure"`
	Description    *string                               `json:"description,omitempty"`
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
	OauthParams []OutputClickHouseOauthParams `json:"oauthParams,omitempty"`
	// Additional headers to send in the OAuth login request. @{product} will automatically add the content-type header 'application/x-www-form-urlencoded' when sending this request.
	OauthHeaders []OutputClickHouseOauthHeaders `json:"oauthHeaders,omitempty"`
	// Username for certificate authentication
	SQLUsername *string `json:"sqlUsername,omitempty"`
	// Cribl will wait for confirmation that data has been fully inserted into the ClickHouse database before proceeding. Disabling this option can increase throughput, but Cribl won’t be able to verify data has been completely inserted.
	WaitForAsyncInserts *bool `default:"true" json:"waitForAsyncInserts"`
	// Fields to exclude from sending to ClickHouse
	ExcludeMappingFields []string `json:"excludeMappingFields,omitempty"`
	// Retrieves the table schema from ClickHouse and populates the Column Mapping table
	DescribeTable  *string          `json:"describeTable,omitempty"`
	ColumnMappings []ColumnMappings `json:"columnMappings,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	PqMaxFileSize *string `default:"1 MB" json:"pqMaxFileSize"`
	// The maximum disk space that the queue can consume (as an average per Worker Process) before queueing stops. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `default:"5GB" json:"pqMaxSize"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/<output-id>.
	PqPath *string `default:"\\$CRIBL_HOME/state/queues" json:"pqPath"`
	// Codec to use to compress the persisted data.
	PqCompress *OutputClickHouseCompression `default:"none" json:"pqCompress"`
	// Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputClickHouseQueueFullBehavior `default:"block" json:"pqOnBackpressure"`
	// In Error mode, PQ writes events to the filesystem only when it detects a non-retryable Destination error. In Backpressure mode, PQ writes events to the filesystem when it detects backpressure from the Destination or when there are non-retryable Destination errors. In Always On mode, PQ always writes events to the filesystem.
	PqMode     *OutputClickHouseMode       `default:"error" json:"pqMode"`
	PqControls *OutputClickHousePqControls `json:"pqControls,omitempty"`
}

func (o OutputClickHouse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputClickHouse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *OutputClickHouse) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OutputClickHouse) GetType() *OutputClickHouseType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *OutputClickHouse) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *OutputClickHouse) GetSystemFields() []string {
	if o == nil {
		return nil
	}
	return o.SystemFields
}

func (o *OutputClickHouse) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *OutputClickHouse) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *OutputClickHouse) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *OutputClickHouse) GetAuthType() *OutputClickHouseAuthenticationType {
	if o == nil {
		return nil
	}
	return o.AuthType
}

func (o *OutputClickHouse) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *OutputClickHouse) GetTableName() string {
	if o == nil {
		return ""
	}
	return o.TableName
}

func (o *OutputClickHouse) GetFormat() *OutputClickHouseFormat {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *OutputClickHouse) GetMappingType() *OutputClickHouseMappingType {
	if o == nil {
		return nil
	}
	return o.MappingType
}

func (o *OutputClickHouse) GetAsyncInserts() *bool {
	if o == nil {
		return nil
	}
	return o.AsyncInserts
}

func (o *OutputClickHouse) GetTLS() *OutputClickHouseTLSSettingsClientSide {
	if o == nil {
		return nil
	}
	return o.TLS
}

func (o *OutputClickHouse) GetConcurrency() *float64 {
	if o == nil {
		return nil
	}
	return o.Concurrency
}

func (o *OutputClickHouse) GetMaxPayloadSizeKB() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxPayloadSizeKB
}

func (o *OutputClickHouse) GetMaxPayloadEvents() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxPayloadEvents
}

func (o *OutputClickHouse) GetCompress() *bool {
	if o == nil {
		return nil
	}
	return o.Compress
}

func (o *OutputClickHouse) GetRejectUnauthorized() *bool {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *OutputClickHouse) GetTimeoutSec() *float64 {
	if o == nil {
		return nil
	}
	return o.TimeoutSec
}

func (o *OutputClickHouse) GetFlushPeriodSec() *float64 {
	if o == nil {
		return nil
	}
	return o.FlushPeriodSec
}

func (o *OutputClickHouse) GetExtraHTTPHeaders() []OutputClickHouseExtraHTTPHeaders {
	if o == nil {
		return nil
	}
	return o.ExtraHTTPHeaders
}

func (o *OutputClickHouse) GetUseRoundRobinDNS() *bool {
	if o == nil {
		return nil
	}
	return o.UseRoundRobinDNS
}

func (o *OutputClickHouse) GetFailedRequestLoggingMode() *OutputClickHouseFailedRequestLoggingMode {
	if o == nil {
		return nil
	}
	return o.FailedRequestLoggingMode
}

func (o *OutputClickHouse) GetSafeHeaders() []string {
	if o == nil {
		return nil
	}
	return o.SafeHeaders
}

func (o *OutputClickHouse) GetResponseRetrySettings() []OutputClickHouseResponseRetrySettings {
	if o == nil {
		return nil
	}
	return o.ResponseRetrySettings
}

func (o *OutputClickHouse) GetTimeoutRetrySettings() *OutputClickHouseTimeoutRetrySettings {
	if o == nil {
		return nil
	}
	return o.TimeoutRetrySettings
}

func (o *OutputClickHouse) GetResponseHonorRetryAfterHeader() *bool {
	if o == nil {
		return nil
	}
	return o.ResponseHonorRetryAfterHeader
}

func (o *OutputClickHouse) GetDumpFormatErrorsToDisk() *bool {
	if o == nil {
		return nil
	}
	return o.DumpFormatErrorsToDisk
}

func (o *OutputClickHouse) GetOnBackpressure() *OutputClickHouseBackpressureBehavior {
	if o == nil {
		return nil
	}
	return o.OnBackpressure
}

func (o *OutputClickHouse) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *OutputClickHouse) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *OutputClickHouse) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *OutputClickHouse) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

func (o *OutputClickHouse) GetCredentialsSecret() *string {
	if o == nil {
		return nil
	}
	return o.CredentialsSecret
}

func (o *OutputClickHouse) GetTextSecret() *string {
	if o == nil {
		return nil
	}
	return o.TextSecret
}

func (o *OutputClickHouse) GetLoginURL() *string {
	if o == nil {
		return nil
	}
	return o.LoginURL
}

func (o *OutputClickHouse) GetSecretParamName() *string {
	if o == nil {
		return nil
	}
	return o.SecretParamName
}

func (o *OutputClickHouse) GetSecret() *string {
	if o == nil {
		return nil
	}
	return o.Secret
}

func (o *OutputClickHouse) GetTokenAttributeName() *string {
	if o == nil {
		return nil
	}
	return o.TokenAttributeName
}

func (o *OutputClickHouse) GetAuthHeaderExpr() *string {
	if o == nil {
		return nil
	}
	return o.AuthHeaderExpr
}

func (o *OutputClickHouse) GetTokenTimeoutSecs() *float64 {
	if o == nil {
		return nil
	}
	return o.TokenTimeoutSecs
}

func (o *OutputClickHouse) GetOauthParams() []OutputClickHouseOauthParams {
	if o == nil {
		return nil
	}
	return o.OauthParams
}

func (o *OutputClickHouse) GetOauthHeaders() []OutputClickHouseOauthHeaders {
	if o == nil {
		return nil
	}
	return o.OauthHeaders
}

func (o *OutputClickHouse) GetSQLUsername() *string {
	if o == nil {
		return nil
	}
	return o.SQLUsername
}

func (o *OutputClickHouse) GetWaitForAsyncInserts() *bool {
	if o == nil {
		return nil
	}
	return o.WaitForAsyncInserts
}

func (o *OutputClickHouse) GetExcludeMappingFields() []string {
	if o == nil {
		return nil
	}
	return o.ExcludeMappingFields
}

func (o *OutputClickHouse) GetDescribeTable() *string {
	if o == nil {
		return nil
	}
	return o.DescribeTable
}

func (o *OutputClickHouse) GetColumnMappings() []ColumnMappings {
	if o == nil {
		return nil
	}
	return o.ColumnMappings
}

func (o *OutputClickHouse) GetPqMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxFileSize
}

func (o *OutputClickHouse) GetPqMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxSize
}

func (o *OutputClickHouse) GetPqPath() *string {
	if o == nil {
		return nil
	}
	return o.PqPath
}

func (o *OutputClickHouse) GetPqCompress() *OutputClickHouseCompression {
	if o == nil {
		return nil
	}
	return o.PqCompress
}

func (o *OutputClickHouse) GetPqOnBackpressure() *OutputClickHouseQueueFullBehavior {
	if o == nil {
		return nil
	}
	return o.PqOnBackpressure
}

func (o *OutputClickHouse) GetPqMode() *OutputClickHouseMode {
	if o == nil {
		return nil
	}
	return o.PqMode
}

func (o *OutputClickHouse) GetPqControls() *OutputClickHousePqControls {
	if o == nil {
		return nil
	}
	return o.PqControls
}
