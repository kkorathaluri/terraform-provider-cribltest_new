// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

type InputType string

const (
	InputTypeSplunk               InputType = "splunk"
	InputTypeSplunkHec            InputType = "splunk_hec"
	InputTypeSyslog               InputType = "syslog"
	InputTypeTcpjson              InputType = "tcpjson"
	InputTypeGrafana              InputType = "grafana"
	InputTypeLoki                 InputType = "loki"
	InputTypeHTTP                 InputType = "http"
	InputTypeHTTPRaw              InputType = "http_raw"
	InputTypeFirehose             InputType = "firehose"
	InputTypeElastic              InputType = "elastic"
	InputTypeKafka                InputType = "kafka"
	InputTypeConfluentCloud       InputType = "confluent_cloud"
	InputTypeMsk                  InputType = "msk"
	InputTypeKinesis              InputType = "kinesis"
	InputTypeEventhub             InputType = "eventhub"
	InputTypeAzureBlob            InputType = "azure_blob"
	InputTypeMetrics              InputType = "metrics"
	InputTypeSqs                  InputType = "sqs"
	InputTypeS3                   InputType = "s3"
	InputTypeS3Inventory          InputType = "s3_inventory"
	InputTypeSnmp                 InputType = "snmp"
	InputTypeCrowdstrike          InputType = "crowdstrike"
	InputTypeTCP                  InputType = "tcp"
	InputTypeRawUDP               InputType = "raw_udp"
	InputTypeNetflow              InputType = "netflow"
	InputTypeOffice365Service     InputType = "office365_service"
	InputTypeOffice365Mgmt        InputType = "office365_mgmt"
	InputTypeOffice365MsgTrace    InputType = "office365_msg_trace"
	InputTypePrometheus           InputType = "prometheus"
	InputTypeEdgePrometheus       InputType = "edge_prometheus"
	InputTypePrometheusRw         InputType = "prometheus_rw"
	InputTypeAppscope             InputType = "appscope"
	InputTypeGooglePubsub         InputType = "google_pubsub"
	InputTypeOpenTelemetry        InputType = "open_telemetry"
	InputTypeModelDrivenTelemetry InputType = "model_driven_telemetry"
	InputTypeDatadogAgent         InputType = "datadog_agent"
	InputTypeWef                  InputType = "wef"
	InputTypeWiz                  InputType = "wiz"
	InputTypeZscalerHec           InputType = "zscaler_hec"
	InputTypeDatagen              InputType = "datagen"
	InputTypeCribl                InputType = "cribl"
	InputTypeCriblmetrics         InputType = "criblmetrics"
	InputTypeCriblHTTP            InputType = "cribl_http"
	InputTypeCriblTCP             InputType = "cribl_tcp"
	InputTypeWinEventLogs         InputType = "win_event_logs"
	InputTypeSystemMetrics        InputType = "system_metrics"
	InputTypeWindowsMetrics       InputType = "windows_metrics"
	InputTypeSystemState          InputType = "system_state"
	InputTypeKubeMetrics          InputType = "kube_metrics"
	InputTypeKubeLogs             InputType = "kube_logs"
	InputTypeKubeEvents           InputType = "kube_events"
	InputTypeExec                 InputType = "exec"
	InputTypeSplunkSearch         InputType = "splunk_search"
	InputTypeFile                 InputType = "file"
	InputTypeJournalFiles         InputType = "journal_files"
	InputTypeSecurityLake         InputType = "security_lake"
)

func (e InputType) ToPointer() *InputType {
	return &e
}
func (e *InputType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "splunk":
		fallthrough
	case "splunk_hec":
		fallthrough
	case "syslog":
		fallthrough
	case "tcpjson":
		fallthrough
	case "grafana":
		fallthrough
	case "loki":
		fallthrough
	case "http":
		fallthrough
	case "http_raw":
		fallthrough
	case "firehose":
		fallthrough
	case "elastic":
		fallthrough
	case "kafka":
		fallthrough
	case "confluent_cloud":
		fallthrough
	case "msk":
		fallthrough
	case "kinesis":
		fallthrough
	case "eventhub":
		fallthrough
	case "azure_blob":
		fallthrough
	case "metrics":
		fallthrough
	case "sqs":
		fallthrough
	case "s3":
		fallthrough
	case "s3_inventory":
		fallthrough
	case "snmp":
		fallthrough
	case "crowdstrike":
		fallthrough
	case "tcp":
		fallthrough
	case "raw_udp":
		fallthrough
	case "netflow":
		fallthrough
	case "office365_service":
		fallthrough
	case "office365_mgmt":
		fallthrough
	case "office365_msg_trace":
		fallthrough
	case "prometheus":
		fallthrough
	case "edge_prometheus":
		fallthrough
	case "prometheus_rw":
		fallthrough
	case "appscope":
		fallthrough
	case "google_pubsub":
		fallthrough
	case "open_telemetry":
		fallthrough
	case "model_driven_telemetry":
		fallthrough
	case "datadog_agent":
		fallthrough
	case "wef":
		fallthrough
	case "wiz":
		fallthrough
	case "zscaler_hec":
		fallthrough
	case "datagen":
		fallthrough
	case "cribl":
		fallthrough
	case "criblmetrics":
		fallthrough
	case "cribl_http":
		fallthrough
	case "cribl_tcp":
		fallthrough
	case "win_event_logs":
		fallthrough
	case "system_metrics":
		fallthrough
	case "windows_metrics":
		fallthrough
	case "system_state":
		fallthrough
	case "kube_metrics":
		fallthrough
	case "kube_logs":
		fallthrough
	case "kube_events":
		fallthrough
	case "exec":
		fallthrough
	case "splunk_search":
		fallthrough
	case "file":
		fallthrough
	case "journal_files":
		fallthrough
	case "security_lake":
		*e = InputType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputType: %v", v)
	}
}

type InputDatagenConnections struct {
	Pipeline *string `json:"pipeline,omitempty"`
	Output   string  `json:"output"`
}

func (o *InputDatagenConnections) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputDatagenConnections) GetOutput() string {
	if o == nil {
		return ""
	}
	return o.Output
}

// InputDatagenMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputDatagenMode string

const (
	InputDatagenModeSmart  InputDatagenMode = "smart"
	InputDatagenModeAlways InputDatagenMode = "always"
)

func (e InputDatagenMode) ToPointer() *InputDatagenMode {
	return &e
}
func (e *InputDatagenMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputDatagenMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputDatagenMode: %v", v)
	}
}

// InputDatagenCompression - Codec to use to compress the persisted data
type InputDatagenCompression string

const (
	InputDatagenCompressionNone InputDatagenCompression = "none"
	InputDatagenCompressionGzip InputDatagenCompression = "gzip"
)

func (e InputDatagenCompression) ToPointer() *InputDatagenCompression {
	return &e
}
func (e *InputDatagenCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputDatagenCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputDatagenCompression: %v", v)
	}
}

type InputDatagenPq struct {
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputDatagenMode `default:"always" json:"mode"`
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
	Compress *InputDatagenCompression `default:"none" json:"compress"`
}

func (i InputDatagenPq) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputDatagenPq) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputDatagenPq) GetMode() *InputDatagenMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputDatagenPq) GetMaxBufferSize() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBufferSize
}

func (o *InputDatagenPq) GetCommitFrequency() *float64 {
	if o == nil {
		return nil
	}
	return o.CommitFrequency
}

func (o *InputDatagenPq) GetMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxFileSize
}

func (o *InputDatagenPq) GetMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxSize
}

func (o *InputDatagenPq) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *InputDatagenPq) GetCompress() *InputDatagenCompression {
	if o == nil {
		return nil
	}
	return o.Compress
}

type Samples struct {
	// Name of the datagen file
	Sample string `json:"sample"`
	// Maximum no. of events to generate per second per worker node. Defaults to 10.
	EventsPerSec *float64 `default:"10" json:"eventsPerSec"`
}

func (s Samples) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Samples) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Samples) GetSample() string {
	if o == nil {
		return ""
	}
	return o.Sample
}

func (o *Samples) GetEventsPerSec() *float64 {
	if o == nil {
		return nil
	}
	return o.EventsPerSec
}

type InputDatagenMetadata struct {
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

func (o *InputDatagenMetadata) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InputDatagenMetadata) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type InputDatagen struct {
	// Unique ID for this input
	ID       *string   `json:"id,omitempty"`
	Type     InputType `json:"type"`
	Disabled *bool     `default:"false" json:"disabled"`
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
	Connections []InputDatagenConnections `json:"connections,omitempty"`
	Pq          *InputDatagenPq           `json:"pq,omitempty"`
	// List of datagens
	Samples []Samples `json:"samples"`
	// Fields to add to events from this input
	Metadata    []InputDatagenMetadata `json:"metadata,omitempty"`
	Description *string                `json:"description,omitempty"`
}

func (i InputDatagen) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputDatagen) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *InputDatagen) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *InputDatagen) GetType() InputType {
	if o == nil {
		return InputType("")
	}
	return o.Type
}

func (o *InputDatagen) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputDatagen) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputDatagen) GetSendToRoutes() *bool {
	if o == nil {
		return nil
	}
	return o.SendToRoutes
}

func (o *InputDatagen) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *InputDatagen) GetPqEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.PqEnabled
}

func (o *InputDatagen) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *InputDatagen) GetConnections() []InputDatagenConnections {
	if o == nil {
		return nil
	}
	return o.Connections
}

func (o *InputDatagen) GetPq() *InputDatagenPq {
	if o == nil {
		return nil
	}
	return o.Pq
}

func (o *InputDatagen) GetSamples() []Samples {
	if o == nil {
		return []Samples{}
	}
	return o.Samples
}

func (o *InputDatagen) GetMetadata() []InputDatagenMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *InputDatagen) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}
