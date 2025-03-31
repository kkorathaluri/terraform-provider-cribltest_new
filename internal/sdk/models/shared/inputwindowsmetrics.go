// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

type InputWindowsMetricsType string

const (
	InputWindowsMetricsTypeWindowsMetrics InputWindowsMetricsType = "windows_metrics"
)

func (e InputWindowsMetricsType) ToPointer() *InputWindowsMetricsType {
	return &e
}
func (e *InputWindowsMetricsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "windows_metrics":
		*e = InputWindowsMetricsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputWindowsMetricsType: %v", v)
	}
}

type InputWindowsMetricsConnections struct {
	Pipeline *string `json:"pipeline,omitempty"`
	Output   string  `json:"output"`
}

func (o *InputWindowsMetricsConnections) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputWindowsMetricsConnections) GetOutput() string {
	if o == nil {
		return ""
	}
	return o.Output
}

// InputWindowsMetricsPqMode - With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
type InputWindowsMetricsPqMode string

const (
	InputWindowsMetricsPqModeSmart  InputWindowsMetricsPqMode = "smart"
	InputWindowsMetricsPqModeAlways InputWindowsMetricsPqMode = "always"
)

func (e InputWindowsMetricsPqMode) ToPointer() *InputWindowsMetricsPqMode {
	return &e
}
func (e *InputWindowsMetricsPqMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smart":
		fallthrough
	case "always":
		*e = InputWindowsMetricsPqMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputWindowsMetricsPqMode: %v", v)
	}
}

// InputWindowsMetricsCompression - Codec to use to compress the persisted data
type InputWindowsMetricsCompression string

const (
	InputWindowsMetricsCompressionNone InputWindowsMetricsCompression = "none"
	InputWindowsMetricsCompressionGzip InputWindowsMetricsCompression = "gzip"
)

func (e InputWindowsMetricsCompression) ToPointer() *InputWindowsMetricsCompression {
	return &e
}
func (e *InputWindowsMetricsCompression) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputWindowsMetricsCompression(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputWindowsMetricsCompression: %v", v)
	}
}

type InputWindowsMetricsPq struct {
	// With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.
	Mode *InputWindowsMetricsPqMode `default:"always" json:"mode"`
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
	Compress *InputWindowsMetricsCompression `default:"none" json:"compress"`
}

func (i InputWindowsMetricsPq) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputWindowsMetricsPq) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputWindowsMetricsPq) GetMode() *InputWindowsMetricsPqMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputWindowsMetricsPq) GetMaxBufferSize() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxBufferSize
}

func (o *InputWindowsMetricsPq) GetCommitFrequency() *float64 {
	if o == nil {
		return nil
	}
	return o.CommitFrequency
}

func (o *InputWindowsMetricsPq) GetMaxFileSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxFileSize
}

func (o *InputWindowsMetricsPq) GetMaxSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxSize
}

func (o *InputWindowsMetricsPq) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *InputWindowsMetricsPq) GetCompress() *InputWindowsMetricsCompression {
	if o == nil {
		return nil
	}
	return o.Compress
}

// InputWindowsMetricsMode - Select level of detail for host metrics
type InputWindowsMetricsMode string

const (
	InputWindowsMetricsModeBasic    InputWindowsMetricsMode = "basic"
	InputWindowsMetricsModeAll      InputWindowsMetricsMode = "all"
	InputWindowsMetricsModeCustom   InputWindowsMetricsMode = "custom"
	InputWindowsMetricsModeDisabled InputWindowsMetricsMode = "disabled"
)

func (e InputWindowsMetricsMode) ToPointer() *InputWindowsMetricsMode {
	return &e
}
func (e *InputWindowsMetricsMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "basic":
		fallthrough
	case "all":
		fallthrough
	case "custom":
		fallthrough
	case "disabled":
		*e = InputWindowsMetricsMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputWindowsMetricsMode: %v", v)
	}
}

// InputWindowsMetricsHostMode - Select the level of details for system metrics
type InputWindowsMetricsHostMode string

const (
	InputWindowsMetricsHostModeBasic    InputWindowsMetricsHostMode = "basic"
	InputWindowsMetricsHostModeAll      InputWindowsMetricsHostMode = "all"
	InputWindowsMetricsHostModeCustom   InputWindowsMetricsHostMode = "custom"
	InputWindowsMetricsHostModeDisabled InputWindowsMetricsHostMode = "disabled"
)

func (e InputWindowsMetricsHostMode) ToPointer() *InputWindowsMetricsHostMode {
	return &e
}
func (e *InputWindowsMetricsHostMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "basic":
		fallthrough
	case "all":
		fallthrough
	case "custom":
		fallthrough
	case "disabled":
		*e = InputWindowsMetricsHostMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputWindowsMetricsHostMode: %v", v)
	}
}

type InputWindowsMetricsSystem struct {
	// Select the level of details for system metrics
	Mode *InputWindowsMetricsHostMode `default:"basic" json:"mode"`
	// Generate metrics for all system information
	Detail *bool `default:"false" json:"detail"`
}

func (i InputWindowsMetricsSystem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputWindowsMetricsSystem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputWindowsMetricsSystem) GetMode() *InputWindowsMetricsHostMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputWindowsMetricsSystem) GetDetail() *bool {
	if o == nil {
		return nil
	}
	return o.Detail
}

// InputWindowsMetricsHostCustomMode - Select the level of details for CPU metrics
type InputWindowsMetricsHostCustomMode string

const (
	InputWindowsMetricsHostCustomModeBasic    InputWindowsMetricsHostCustomMode = "basic"
	InputWindowsMetricsHostCustomModeAll      InputWindowsMetricsHostCustomMode = "all"
	InputWindowsMetricsHostCustomModeCustom   InputWindowsMetricsHostCustomMode = "custom"
	InputWindowsMetricsHostCustomModeDisabled InputWindowsMetricsHostCustomMode = "disabled"
)

func (e InputWindowsMetricsHostCustomMode) ToPointer() *InputWindowsMetricsHostCustomMode {
	return &e
}
func (e *InputWindowsMetricsHostCustomMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "basic":
		fallthrough
	case "all":
		fallthrough
	case "custom":
		fallthrough
	case "disabled":
		*e = InputWindowsMetricsHostCustomMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputWindowsMetricsHostCustomMode: %v", v)
	}
}

type InputWindowsMetricsCPU struct {
	// Select the level of details for CPU metrics
	Mode *InputWindowsMetricsHostCustomMode `default:"basic" json:"mode"`
	// Generate metrics for each CPU
	PerCPU *bool `default:"false" json:"perCpu"`
	// Generate metrics for all CPU states
	Detail *bool `default:"false" json:"detail"`
	// Generate raw, monotonic CPU time counters
	Time *bool `default:"false" json:"time"`
}

func (i InputWindowsMetricsCPU) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputWindowsMetricsCPU) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputWindowsMetricsCPU) GetMode() *InputWindowsMetricsHostCustomMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputWindowsMetricsCPU) GetPerCPU() *bool {
	if o == nil {
		return nil
	}
	return o.PerCPU
}

func (o *InputWindowsMetricsCPU) GetDetail() *bool {
	if o == nil {
		return nil
	}
	return o.Detail
}

func (o *InputWindowsMetricsCPU) GetTime() *bool {
	if o == nil {
		return nil
	}
	return o.Time
}

// InputWindowsMetricsHostCustomMemoryMode - Select the level of details for memory metrics
type InputWindowsMetricsHostCustomMemoryMode string

const (
	InputWindowsMetricsHostCustomMemoryModeBasic    InputWindowsMetricsHostCustomMemoryMode = "basic"
	InputWindowsMetricsHostCustomMemoryModeAll      InputWindowsMetricsHostCustomMemoryMode = "all"
	InputWindowsMetricsHostCustomMemoryModeCustom   InputWindowsMetricsHostCustomMemoryMode = "custom"
	InputWindowsMetricsHostCustomMemoryModeDisabled InputWindowsMetricsHostCustomMemoryMode = "disabled"
)

func (e InputWindowsMetricsHostCustomMemoryMode) ToPointer() *InputWindowsMetricsHostCustomMemoryMode {
	return &e
}
func (e *InputWindowsMetricsHostCustomMemoryMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "basic":
		fallthrough
	case "all":
		fallthrough
	case "custom":
		fallthrough
	case "disabled":
		*e = InputWindowsMetricsHostCustomMemoryMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputWindowsMetricsHostCustomMemoryMode: %v", v)
	}
}

type InputWindowsMetricsMemory struct {
	// Select the level of details for memory metrics
	Mode *InputWindowsMetricsHostCustomMemoryMode `default:"basic" json:"mode"`
	// Generate metrics for all memory states
	Detail *bool `default:"false" json:"detail"`
}

func (i InputWindowsMetricsMemory) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputWindowsMetricsMemory) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputWindowsMetricsMemory) GetMode() *InputWindowsMetricsHostCustomMemoryMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputWindowsMetricsMemory) GetDetail() *bool {
	if o == nil {
		return nil
	}
	return o.Detail
}

// InputWindowsMetricsHostCustomNetworkMode - Select the level of details for network metrics
type InputWindowsMetricsHostCustomNetworkMode string

const (
	InputWindowsMetricsHostCustomNetworkModeBasic    InputWindowsMetricsHostCustomNetworkMode = "basic"
	InputWindowsMetricsHostCustomNetworkModeAll      InputWindowsMetricsHostCustomNetworkMode = "all"
	InputWindowsMetricsHostCustomNetworkModeCustom   InputWindowsMetricsHostCustomNetworkMode = "custom"
	InputWindowsMetricsHostCustomNetworkModeDisabled InputWindowsMetricsHostCustomNetworkMode = "disabled"
)

func (e InputWindowsMetricsHostCustomNetworkMode) ToPointer() *InputWindowsMetricsHostCustomNetworkMode {
	return &e
}
func (e *InputWindowsMetricsHostCustomNetworkMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "basic":
		fallthrough
	case "all":
		fallthrough
	case "custom":
		fallthrough
	case "disabled":
		*e = InputWindowsMetricsHostCustomNetworkMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputWindowsMetricsHostCustomNetworkMode: %v", v)
	}
}

type InputWindowsMetricsNetwork struct {
	// Select the level of details for network metrics
	Mode *InputWindowsMetricsHostCustomNetworkMode `default:"basic" json:"mode"`
	// Network interfaces to include/exclude. All interfaces are included if this list is empty.
	Devices []string `json:"devices,omitempty"`
	// Generate separate metrics for each interface
	PerInterface *bool `default:"false" json:"perInterface"`
	// Generate full network metrics
	Detail *bool `default:"false" json:"detail"`
}

func (i InputWindowsMetricsNetwork) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputWindowsMetricsNetwork) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputWindowsMetricsNetwork) GetMode() *InputWindowsMetricsHostCustomNetworkMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputWindowsMetricsNetwork) GetDevices() []string {
	if o == nil {
		return nil
	}
	return o.Devices
}

func (o *InputWindowsMetricsNetwork) GetPerInterface() *bool {
	if o == nil {
		return nil
	}
	return o.PerInterface
}

func (o *InputWindowsMetricsNetwork) GetDetail() *bool {
	if o == nil {
		return nil
	}
	return o.Detail
}

// InputWindowsMetricsHostCustomDiskMode - Select the level of details for disk metrics
type InputWindowsMetricsHostCustomDiskMode string

const (
	InputWindowsMetricsHostCustomDiskModeBasic    InputWindowsMetricsHostCustomDiskMode = "basic"
	InputWindowsMetricsHostCustomDiskModeAll      InputWindowsMetricsHostCustomDiskMode = "all"
	InputWindowsMetricsHostCustomDiskModeCustom   InputWindowsMetricsHostCustomDiskMode = "custom"
	InputWindowsMetricsHostCustomDiskModeDisabled InputWindowsMetricsHostCustomDiskMode = "disabled"
)

func (e InputWindowsMetricsHostCustomDiskMode) ToPointer() *InputWindowsMetricsHostCustomDiskMode {
	return &e
}
func (e *InputWindowsMetricsHostCustomDiskMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "basic":
		fallthrough
	case "all":
		fallthrough
	case "custom":
		fallthrough
	case "disabled":
		*e = InputWindowsMetricsHostCustomDiskMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputWindowsMetricsHostCustomDiskMode: %v", v)
	}
}

type InputWindowsMetricsDisk struct {
	// Select the level of details for disk metrics
	Mode *InputWindowsMetricsHostCustomDiskMode `default:"basic" json:"mode"`
	// Windows volumes to include/exclude. E.g.: C:, !E:, etc. Wildcards and ! (not) operators are supported. All volumes are included if this list is empty.
	Volumes []string `json:"volumes,omitempty"`
	// Generate separate metrics for each volume
	PerVolume *bool `default:"false" json:"perVolume"`
}

func (i InputWindowsMetricsDisk) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputWindowsMetricsDisk) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputWindowsMetricsDisk) GetMode() *InputWindowsMetricsHostCustomDiskMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputWindowsMetricsDisk) GetVolumes() []string {
	if o == nil {
		return nil
	}
	return o.Volumes
}

func (o *InputWindowsMetricsDisk) GetPerVolume() *bool {
	if o == nil {
		return nil
	}
	return o.PerVolume
}

type InputWindowsMetricsCustom struct {
	System  *InputWindowsMetricsSystem  `json:"system,omitempty"`
	CPU     *InputWindowsMetricsCPU     `json:"cpu,omitempty"`
	Memory  *InputWindowsMetricsMemory  `json:"memory,omitempty"`
	Network *InputWindowsMetricsNetwork `json:"network,omitempty"`
	Disk    *InputWindowsMetricsDisk    `json:"disk,omitempty"`
}

func (o *InputWindowsMetricsCustom) GetSystem() *InputWindowsMetricsSystem {
	if o == nil {
		return nil
	}
	return o.System
}

func (o *InputWindowsMetricsCustom) GetCPU() *InputWindowsMetricsCPU {
	if o == nil {
		return nil
	}
	return o.CPU
}

func (o *InputWindowsMetricsCustom) GetMemory() *InputWindowsMetricsMemory {
	if o == nil {
		return nil
	}
	return o.Memory
}

func (o *InputWindowsMetricsCustom) GetNetwork() *InputWindowsMetricsNetwork {
	if o == nil {
		return nil
	}
	return o.Network
}

func (o *InputWindowsMetricsCustom) GetDisk() *InputWindowsMetricsDisk {
	if o == nil {
		return nil
	}
	return o.Disk
}

type InputWindowsMetricsHost struct {
	// Select level of detail for host metrics
	Mode   *InputWindowsMetricsMode   `default:"basic" json:"mode"`
	Custom *InputWindowsMetricsCustom `json:"custom,omitempty"`
}

func (i InputWindowsMetricsHost) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputWindowsMetricsHost) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputWindowsMetricsHost) GetMode() *InputWindowsMetricsMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *InputWindowsMetricsHost) GetCustom() *InputWindowsMetricsCustom {
	if o == nil {
		return nil
	}
	return o.Custom
}

type InputWindowsMetricsSets struct {
	Name            string `json:"name"`
	Filter          string `json:"filter"`
	IncludeChildren *bool  `default:"false" json:"includeChildren"`
}

func (i InputWindowsMetricsSets) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputWindowsMetricsSets) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputWindowsMetricsSets) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InputWindowsMetricsSets) GetFilter() string {
	if o == nil {
		return ""
	}
	return o.Filter
}

func (o *InputWindowsMetricsSets) GetIncludeChildren() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeChildren
}

type InputWindowsMetricsProcess struct {
	// Configure sets to collect process metrics
	Sets []InputWindowsMetricsSets `json:"sets,omitempty"`
}

func (o *InputWindowsMetricsProcess) GetSets() []InputWindowsMetricsSets {
	if o == nil {
		return nil
	}
	return o.Sets
}

type InputWindowsMetricsMetadata struct {
	Name string `json:"name"`
	// JavaScript expression to compute field's value, enclosed in quotes or backticks. (Can evaluate to a constant.)
	Value string `json:"value"`
}

func (o *InputWindowsMetricsMetadata) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *InputWindowsMetricsMetadata) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type InputWindowsMetricsDataCompressionFormat string

const (
	InputWindowsMetricsDataCompressionFormatNone InputWindowsMetricsDataCompressionFormat = "none"
	InputWindowsMetricsDataCompressionFormatGzip InputWindowsMetricsDataCompressionFormat = "gzip"
)

func (e InputWindowsMetricsDataCompressionFormat) ToPointer() *InputWindowsMetricsDataCompressionFormat {
	return &e
}
func (e *InputWindowsMetricsDataCompressionFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "gzip":
		*e = InputWindowsMetricsDataCompressionFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputWindowsMetricsDataCompressionFormat: %v", v)
	}
}

type InputWindowsMetricsPersistence struct {
	// Spool metrics to disk for Cribl Edge and Search
	Enable *bool `default:"false" json:"enable"`
	// Time span for each file bucket
	TimeWindow *string `default:"10m" json:"timeWindow"`
	// Maximum disk space allowed to be consumed (examples: 420MB, 4GB). When limit is reached, older data will be deleted.
	MaxDataSize *string `default:"1GB" json:"maxDataSize"`
	// Maximum amount of time to retain data (examples: 2h, 4d). When limit is reached, older data will be deleted.
	MaxDataTime *string                                   `default:"24h" json:"maxDataTime"`
	Compress    *InputWindowsMetricsDataCompressionFormat `default:"gzip" json:"compress"`
	// Path to use to write metrics. Defaults to $CRIBL_HOME/state/windows_metrics
	DestPath *string `default:"\\$CRIBL_HOME/state/windows_metrics" json:"destPath"`
}

func (i InputWindowsMetricsPersistence) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputWindowsMetricsPersistence) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InputWindowsMetricsPersistence) GetEnable() *bool {
	if o == nil {
		return nil
	}
	return o.Enable
}

func (o *InputWindowsMetricsPersistence) GetTimeWindow() *string {
	if o == nil {
		return nil
	}
	return o.TimeWindow
}

func (o *InputWindowsMetricsPersistence) GetMaxDataSize() *string {
	if o == nil {
		return nil
	}
	return o.MaxDataSize
}

func (o *InputWindowsMetricsPersistence) GetMaxDataTime() *string {
	if o == nil {
		return nil
	}
	return o.MaxDataTime
}

func (o *InputWindowsMetricsPersistence) GetCompress() *InputWindowsMetricsDataCompressionFormat {
	if o == nil {
		return nil
	}
	return o.Compress
}

func (o *InputWindowsMetricsPersistence) GetDestPath() *string {
	if o == nil {
		return nil
	}
	return o.DestPath
}

type InputWindowsMetrics struct {
	// Unique ID for this input
	ID       string                  `json:"id"`
	Type     InputWindowsMetricsType `json:"type"`
	Disabled *bool                   `default:"false" json:"disabled"`
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
	Connections []InputWindowsMetricsConnections `json:"connections,omitempty"`
	Pq          *InputWindowsMetricsPq           `json:"pq,omitempty"`
	// Time, in seconds, between consecutive metric collections. Default is 10 seconds.
	Interval *float64                    `default:"10" json:"interval"`
	Host     *InputWindowsMetricsHost    `json:"host,omitempty"`
	Process  *InputWindowsMetricsProcess `json:"process,omitempty"`
	// Fields to add to events from this input
	Metadata    []InputWindowsMetricsMetadata   `json:"metadata,omitempty"`
	Persistence *InputWindowsMetricsPersistence `json:"persistence,omitempty"`
	// Enable to use built-in tools (PowerShell) to collect metrics instead of native API (default) [Learn more](https://docs.cribl.io/edge/sources-windows-metrics/#advanced-tab)
	DisableNativeModule *bool   `default:"false" json:"disableNativeModule"`
	Description         *string `json:"description,omitempty"`
}

func (i InputWindowsMetrics) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InputWindowsMetrics) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *InputWindowsMetrics) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *InputWindowsMetrics) GetType() InputWindowsMetricsType {
	if o == nil {
		return InputWindowsMetricsType("")
	}
	return o.Type
}

func (o *InputWindowsMetrics) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *InputWindowsMetrics) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *InputWindowsMetrics) GetSendToRoutes() *bool {
	if o == nil {
		return nil
	}
	return o.SendToRoutes
}

func (o *InputWindowsMetrics) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *InputWindowsMetrics) GetPqEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.PqEnabled
}

func (o *InputWindowsMetrics) GetStreamtags() []string {
	if o == nil {
		return nil
	}
	return o.Streamtags
}

func (o *InputWindowsMetrics) GetConnections() []InputWindowsMetricsConnections {
	if o == nil {
		return nil
	}
	return o.Connections
}

func (o *InputWindowsMetrics) GetPq() *InputWindowsMetricsPq {
	if o == nil {
		return nil
	}
	return o.Pq
}

func (o *InputWindowsMetrics) GetInterval() *float64 {
	if o == nil {
		return nil
	}
	return o.Interval
}

func (o *InputWindowsMetrics) GetHost() *InputWindowsMetricsHost {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *InputWindowsMetrics) GetProcess() *InputWindowsMetricsProcess {
	if o == nil {
		return nil
	}
	return o.Process
}

func (o *InputWindowsMetrics) GetMetadata() []InputWindowsMetricsMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *InputWindowsMetrics) GetPersistence() *InputWindowsMetricsPersistence {
	if o == nil {
		return nil
	}
	return o.Persistence
}

func (o *InputWindowsMetrics) GetDisableNativeModule() *bool {
	if o == nil {
		return nil
	}
	return o.DisableNativeModule
}

func (o *InputWindowsMetrics) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}
