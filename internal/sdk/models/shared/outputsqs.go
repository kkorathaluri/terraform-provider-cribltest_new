// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

type OutputSqsType string

const (
	OutputSqsTypeSqs OutputSqsType = "sqs"
)

func (e OutputSqsType) ToPointer() *OutputSqsType {
	return &e
}
func (e *OutputSqsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sqs":
		*e = OutputSqsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSqsType: %v", v)
	}
}

// OutputSqsQueueType - The queue type used (or created). Defaults to Standard.
type OutputSqsQueueType string

const (
	OutputSqsQueueTypeStandard OutputSqsQueueType = "standard"
	OutputSqsQueueTypeFifo     OutputSqsQueueType = "fifo"
)

func (e OutputSqsQueueType) ToPointer() *OutputSqsQueueType {
	return &e
}
func (e *OutputSqsQueueType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "standard":
		fallthrough
	case "fifo":
		*e = OutputSqsQueueType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSqsQueueType: %v", v)
	}
}

// OutputSqsAuthenticationMethod - AWS authentication method. Choose Auto to use IAM roles.
type OutputSqsAuthenticationMethod string

const (
	OutputSqsAuthenticationMethodAuto   OutputSqsAuthenticationMethod = "auto"
	OutputSqsAuthenticationMethodManual OutputSqsAuthenticationMethod = "manual"
	OutputSqsAuthenticationMethodSecret OutputSqsAuthenticationMethod = "secret"
)

func (e OutputSqsAuthenticationMethod) ToPointer() *OutputSqsAuthenticationMethod {
	return &e
}
func (e *OutputSqsAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "auto":
		fallthrough
	case "manual":
		fallthrough
	case "secret":
		*e = OutputSqsAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSqsAuthenticationMethod: %v", v)
	}
}

// OutputSqsSignatureVersion - Signature version to use for signing SQS requests
type OutputSqsSignatureVersion string

const (
	OutputSqsSignatureVersionV2 OutputSqsSignatureVersion = "v2"
	OutputSqsSignatureVersionV4 OutputSqsSignatureVersion = "v4"
)

func (e OutputSqsSignatureVersion) ToPointer() *OutputSqsSignatureVersion {
	return &e
}
func (e *OutputSqsSignatureVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "v2":
		fallthrough
	case "v4":
		*e = OutputSqsSignatureVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSqsSignatureVersion: %v", v)
	}
}

// OutputSqsBackpressureBehavior - Whether to block, drop, or queue events when all receivers are exerting backpressure.
type OutputSqsBackpressureBehavior string

const (
	OutputSqsBackpressureBehaviorBlock OutputSqsBackpressureBehavior = "block"
	OutputSqsBackpressureBehaviorDrop  OutputSqsBackpressureBehavior = "drop"
	OutputSqsBackpressureBehaviorQueue OutputSqsBackpressureBehavior = "queue"
)

func (e OutputSqsBackpressureBehavior) ToPointer() *OutputSqsBackpressureBehavior {
	return &e
}
func (e *OutputSqsBackpressureBehavior) UnmarshalJSON(data []byte) error {
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
		*e = OutputSqsBackpressureBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSqsBackpressureBehavior: %v", v)
	}
}

// OutputSqsCompression - Codec to use to compress the persisted data.
type OutputSqsCompression string

const (
	OutputSqsCompressionNone OutputSqsCompression = "none"
	OutputSqsCompressionGzip OutputSqsCompression = "gzip"
)

func (e OutputSqsCompression) ToPointer() *OutputSqsCompression {
	return &e
}
func (e *OutputSqsCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = OutputSqsCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSqsCompression: %v", v)
	}
}

// OutputSqsQueueFullBehavior - Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
type OutputSqsQueueFullBehavior string

const (
	OutputSqsQueueFullBehaviorBlock OutputSqsQueueFullBehavior = "block"
	OutputSqsQueueFullBehaviorDrop  OutputSqsQueueFullBehavior = "drop"
)

func (e OutputSqsQueueFullBehavior) ToPointer() *OutputSqsQueueFullBehavior {
	return &e
}
func (e *OutputSqsQueueFullBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "drop":
		*e = OutputSqsQueueFullBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSqsQueueFullBehavior: %v", v)
	}
}

// OutputSqsMode - In Error mode, PQ writes events to the filesystem only when it detects a non-retryable Destination error. In Backpressure mode, PQ writes events to the filesystem when it detects backpressure from the Destination or when there are non-retryable Destination errors. In Always On mode, PQ always writes events to the filesystem.
type OutputSqsMode string

const (
	OutputSqsModeError        OutputSqsMode = "error"
	OutputSqsModeBackpressure OutputSqsMode = "backpressure"
	OutputSqsModeAlways       OutputSqsMode = "always"
)

func (e OutputSqsMode) ToPointer() *OutputSqsMode {
	return &e
}
func (e *OutputSqsMode) UnmarshalJSON(data []byte) error {
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
		*e = OutputSqsMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutputSqsMode: %v", v)
	}
}

type OutputSqsPqControls struct {
}

type OutputSqs struct {
	// Unique ID for this output
	ID   *string        `json:"id,omitempty"`
	Type *OutputSqsType `json:"type,omitempty"`
	// Pipeline to process data before sending out to this output
	Pipeline *string `json:"pipeline,omitempty"`
	// Fields to automatically add to events, such as cribl_pipe. Supports wildcards.
	SystemFields []string `json:"systemFields,omitempty"`
	// Optionally, enable this config only on a specified Git branch. If empty, will be enabled everywhere.
	Environment *string `json:"environment,omitempty"`
	// Tags for filtering and grouping in @{product}
	Streamtags []string `json:"streamtags,omitempty"`
	// The name, URL, or ARN of the SQS queue to send events to. When a non-AWS URL is specified, format must be: '{url}/myQueueName'. E.g., 'https://host:port/myQueueName'. Must be a JavaScript expression (which can evaluate to a constant value), enclosed in quotes or backticks. Can be evaluated only at init time. E.g., referencing a Global Variable: `https://host:port/myQueue-${C.vars.myVar}`.
	QueueName string `json:"queueName"`
	// The queue type used (or created). Defaults to Standard.
	QueueType *OutputSqsQueueType `default:"standard" json:"queueType"`
	// SQS queue owner's AWS account ID. Leave empty if SQS queue is in same AWS account.
	AwsAccountID *string `json:"awsAccountId,omitempty"`
	// This parameter applies only to FIFO queues. The tag that specifies that a message belongs to a specific message group. Messages that belong to the same message group are processed in a FIFO manner. Use event field __messageGroupId to override this value.
	MessageGroupID *string `default:"cribl" json:"messageGroupId"`
	// Create queue if it does not exist.
	CreateQueue *bool `default:"true" json:"createQueue"`
	// AWS authentication method. Choose Auto to use IAM roles.
	AwsAuthenticationMethod *OutputSqsAuthenticationMethod `default:"auto" json:"awsAuthenticationMethod"`
	// Secret key
	AwsSecretKey *string `json:"awsSecretKey,omitempty"`
	// AWS Region where the SQS queue is located. Required, unless the Queue entry is a URL or ARN that includes a Region.
	Region *string `json:"region,omitempty"`
	// SQS service endpoint. If empty, defaults to AWS' Region-specific endpoint. Otherwise, it must point to SQS-compatible endpoint.
	Endpoint *string `json:"endpoint,omitempty"`
	// Signature version to use for signing SQS requests
	SignatureVersion *OutputSqsSignatureVersion `default:"v4" json:"signatureVersion"`
	// Reuse connections between requests, which can improve performance
	ReuseConnections *bool `default:"true" json:"reuseConnections"`
	// Reject certificates that cannot be verified against a valid CA, such as self-signed certificates
	RejectUnauthorized *bool `default:"true" json:"rejectUnauthorized"`
	// Use Assume Role credentials to access SQS
	EnableAssumeRole *bool `default:"false" json:"enableAssumeRole"`
	// Amazon Resource Name (ARN) of the role to assume
	AssumeRoleArn *string `json:"assumeRoleArn,omitempty"`
	// External ID to use when assuming role
	AssumeRoleExternalID *string `json:"assumeRoleExternalId,omitempty"`
	// Duration of the assumed role's session, in seconds. Minimum is 900 (15 minutes), default is 3600 (1 hour), and maximum is 43200 (12 hours).
	DurationSeconds *float64 `default:"3600" json:"durationSeconds"`
	// Maximum number of queued batches before blocking.
	MaxQueueSize *float64 `default:"100" json:"maxQueueSize"`
	// Maximum size (KB) of batches to send. Per the SQS spec, the max allowed value is 256 KB.
	MaxRecordSizeKB *float64 `default:"256" json:"maxRecordSizeKB"`
	// Maximum time between requests. Small values could cause the payload size to be smaller than the configured Max record size.
	FlushPeriodSec *float64 `default:"1" json:"flushPeriodSec"`
	// The maximum number of in-progress API requests before backpressure is applied.
	MaxInProgress *float64 `default:"10" json:"maxInProgress"`
	// Whether to block, drop, or queue events when all receivers are exerting backpressure.
	OnBackpressure *OutputSqsBackpressureBehavior `default:"block" json:"onBackpressure"`
	Description    *string                        `json:"description,omitempty"`
	// Access key
	AwsAPIKey *string `json:"awsApiKey,omitempty"`
	// Select or create a stored secret that references your access key and secret key.
	AwsSecret *string `json:"awsSecret,omitempty"`
	// The maximum size to store in each queue file before closing and optionally compressing (KB, MB, etc.).
	PqMaxFileSize *string `default:"1 MB" json:"pqMaxFileSize"`
	// The maximum disk space that the queue can consume (as an average per Worker Process) before queueing stops. Enter a numeral with units of KB, MB, etc.
	PqMaxSize *string `default:"5GB" json:"pqMaxSize"`
	// The location for the persistent queue files. To this field's value, the system will append: /<worker-id>/<output-id>.
	PqPath *string `default:"\\$CRIBL_HOME/state/queues" json:"pqPath"`
	// Codec to use to compress the persisted data.
	PqCompress *OutputSqsCompression `default:"none" json:"pqCompress"`
	// Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.
	PqOnBackpressure *OutputSqsQueueFullBehavior `default:"block" json:"pqOnBackpressure"`
	// In Error mode, PQ writes events to the filesystem only when it detects a non-retryable Destination error. In Backpressure mode, PQ writes events to the filesystem when it detects backpressure from the Destination or when there are non-retryable Destination errors. In Always On mode, PQ always writes events to the filesystem.
	PqMode     *OutputSqsMode       `default:"error" json:"pqMode"`
	PqControls *OutputSqsPqControls `json:"pqControls,omitempty"`
}

func (o OutputSqs) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OutputSqs) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *OutputSqs) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OutputSqs) GetType() *OutputSqsType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *OutputSqs) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *OutputSqs) GetSystemFields() []string {
	if o == nil {
		return nil
	}
	return o.SystemFields
}

func (o *OutputSqs) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *OutputSqs) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *OutputSqs) GetQueueName() string {
	if o == nil {
		return ""
	}
	return o.QueueName
}

func (o *OutputSqs) GetQueueType() *OutputSqsQueueType {
	if o == nil {
		return nil
	}
	return o.QueueType
}

func (o *OutputSqs) GetAwsAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AwsAccountID
}

func (o *OutputSqs) GetMessageGroupID() *string {
	if o == nil {
		return nil
	}
	return o.MessageGroupID
}

func (o *OutputSqs) GetCreateQueue() *bool {
	if o == nil {
		return nil
	}
	return o.CreateQueue
}

func (o *OutputSqs) GetAwsAuthenticationMethod() *OutputSqsAuthenticationMethod {
	if o == nil {
		return nil
	}
	return o.AwsAuthenticationMethod
}

func (o *OutputSqs) GetAwsSecretKey() *string {
	if o == nil {
		return nil
	}
	return o.AwsSecretKey
}

func (o *OutputSqs) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *OutputSqs) GetEndpoint() *string {
	if o == nil {
		return nil
	}
	return o.Endpoint
}

func (o *OutputSqs) GetSignatureVersion() *OutputSqsSignatureVersion {
	if o == nil {
		return nil
	}
	return o.SignatureVersion
}

func (o *OutputSqs) GetReuseConnections() *bool {
	if o == nil {
		return nil
	}
	return o.ReuseConnections
}

func (o *OutputSqs) GetRejectUnauthorized() *bool {
	if o == nil {
		return nil
	}
	return o.RejectUnauthorized
}

func (o *OutputSqs) GetEnableAssumeRole() *bool {
	if o == nil {
		return nil
	}
	return o.EnableAssumeRole
}

func (o *OutputSqs) GetAssumeRoleArn() *string {
	if o == nil {
		return nil
	}
	return o.AssumeRoleArn
}

func (o *OutputSqs) GetAssumeRoleExternalID() *string {
	if o == nil {
		return nil
	}
	return o.AssumeRoleExternalID
}

func (o *OutputSqs) GetDurationSeconds() *float64 {
	if o == nil {
		return nil
	}
	return o.DurationSeconds
}

func (o *OutputSqs) GetMaxQueueSize() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxQueueSize
}

func (o *OutputSqs) GetMaxRecordSizeKB() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRecordSizeKB
}

func (o *OutputSqs) GetFlushPeriodSec() *float64 {
	if o == nil {
		return nil
	}
	return o.FlushPeriodSec
}

func (o *OutputSqs) GetMaxInProgress() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxInProgress
}

func (o *OutputSqs) GetOnBackpressure() *OutputSqsBackpressureBehavior {
	if o == nil {
		return nil
	}
	return o.OnBackpressure
}

func (o *OutputSqs) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *OutputSqs) GetAwsAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.AwsAPIKey
}

func (o *OutputSqs) GetAwsSecret() *string {
	if o == nil {
		return nil
	}
	return o.AwsSecret
}

func (o *OutputSqs) GetPqMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxFileSize
}

func (o *OutputSqs) GetPqMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.PqMaxSize
}

func (o *OutputSqs) GetPqPath() *string {
	if o == nil {
		return nil
	}
	return o.PqPath
}

func (o *OutputSqs) GetPqCompress() *OutputSqsCompression {
	if o == nil {
		return nil
	}
	return o.PqCompress
}

func (o *OutputSqs) GetPqOnBackpressure() *OutputSqsQueueFullBehavior {
	if o == nil {
		return nil
	}
	return o.PqOnBackpressure
}

func (o *OutputSqs) GetPqMode() *OutputSqsMode {
	if o == nil {
		return nil
	}
	return o.PqMode
}

func (o *OutputSqs) GetPqControls() *OutputSqsPqControls {
	if o == nil {
		return nil
	}
	return o.PqControls
}
