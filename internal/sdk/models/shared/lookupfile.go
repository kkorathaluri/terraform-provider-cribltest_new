// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

// LookupFile2Mode - Operation mode for CSV-based lookups
type LookupFile2Mode string

const (
	LookupFile2ModeMemory LookupFile2Mode = "memory"
	LookupFile2ModeDisk   LookupFile2Mode = "disk"
)

func (e LookupFile2Mode) ToPointer() *LookupFile2Mode {
	return &e
}
func (e *LookupFile2Mode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "memory":
		fallthrough
	case "disk":
		*e = LookupFile2Mode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LookupFile2Mode: %v", v)
	}
}

// LookupFile2Type - Task type
type LookupFile2Type string

const (
	LookupFile2TypeImport LookupFile2Type = "IMPORT"
	LookupFile2TypeIndex  LookupFile2Type = "INDEX"
)

func (e LookupFile2Type) ToPointer() *LookupFile2Type {
	return &e
}
func (e *LookupFile2Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "IMPORT":
		fallthrough
	case "INDEX":
		*e = LookupFile2Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LookupFile2Type: %v", v)
	}
}

type LookupFilePendingTask struct {
	// Task ID (generated).
	ID *string `json:"id,omitempty"`
	// Task type
	Type *LookupFile2Type `json:"type,omitempty"`
	// Error message if task has failed
	Error *string `json:"error,omitempty"`
}

func (o *LookupFilePendingTask) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *LookupFilePendingTask) GetType() *LookupFile2Type {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *LookupFilePendingTask) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

type LookupFile2 struct {
	// File content.
	Content     *string `json:"content,omitempty"`
	ID          string  `json:"id"`
	Description *string `json:"description,omitempty"`
	// One or more tags related to this lookup. Optional.
	Tags *string `json:"tags,omitempty"`
	// File size. Optional.
	Size *float64 `json:"size,omitempty"`
	// Unique string generated for each modification of this lookup
	Version *string `json:"version,omitempty"`
	// Operation mode for CSV-based lookups
	Mode        *LookupFile2Mode       `default:"memory" json:"mode"`
	PendingTask *LookupFilePendingTask `json:"pendingTask,omitempty"`
}

func (l LookupFile2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LookupFile2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *LookupFile2) GetContent() *string {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *LookupFile2) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *LookupFile2) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *LookupFile2) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *LookupFile2) GetSize() *float64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *LookupFile2) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}

func (o *LookupFile2) GetMode() *LookupFile2Mode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *LookupFile2) GetPendingTask() *LookupFilePendingTask {
	if o == nil {
		return nil
	}
	return o.PendingTask
}

type FileInfo struct {
	Filename string `json:"filename"`
}

func (o *FileInfo) GetFilename() string {
	if o == nil {
		return ""
	}
	return o.Filename
}

// LookupFileMode - Operation mode for CSV-based lookups
type LookupFileMode string

const (
	LookupFileModeMemory LookupFileMode = "memory"
	LookupFileModeDisk   LookupFileMode = "disk"
)

func (e LookupFileMode) ToPointer() *LookupFileMode {
	return &e
}
func (e *LookupFileMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "memory":
		fallthrough
	case "disk":
		*e = LookupFileMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LookupFileMode: %v", v)
	}
}

// LookupFileType - Task type
type LookupFileType string

const (
	LookupFileTypeImport LookupFileType = "IMPORT"
	LookupFileTypeIndex  LookupFileType = "INDEX"
)

func (e LookupFileType) ToPointer() *LookupFileType {
	return &e
}
func (e *LookupFileType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "IMPORT":
		fallthrough
	case "INDEX":
		*e = LookupFileType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LookupFileType: %v", v)
	}
}

type PendingTask struct {
	// Task ID (generated).
	ID *string `json:"id,omitempty"`
	// Task type
	Type *LookupFileType `json:"type,omitempty"`
	// Error message if task has failed
	Error *string `json:"error,omitempty"`
}

func (o *PendingTask) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PendingTask) GetType() *LookupFileType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *PendingTask) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

type LookupFile1 struct {
	FileInfo    *FileInfo `json:"fileInfo,omitempty"`
	ID          string    `json:"id"`
	Description *string   `json:"description,omitempty"`
	// One or more tags related to this lookup. Optional.
	Tags *string `json:"tags,omitempty"`
	// File size. Optional.
	Size *float64 `json:"size,omitempty"`
	// Unique string generated for each modification of this lookup
	Version *string `json:"version,omitempty"`
	// Operation mode for CSV-based lookups
	Mode        *LookupFileMode `default:"memory" json:"mode"`
	PendingTask *PendingTask    `json:"pendingTask,omitempty"`
}

func (l LookupFile1) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LookupFile1) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *LookupFile1) GetFileInfo() *FileInfo {
	if o == nil {
		return nil
	}
	return o.FileInfo
}

func (o *LookupFile1) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *LookupFile1) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *LookupFile1) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *LookupFile1) GetSize() *float64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *LookupFile1) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}

func (o *LookupFile1) GetMode() *LookupFileMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *LookupFile1) GetPendingTask() *PendingTask {
	if o == nil {
		return nil
	}
	return o.PendingTask
}

type LookupFileUnionType string

const (
	LookupFileUnionTypeLookupFile1 LookupFileUnionType = "LookupFile_1"
	LookupFileUnionTypeLookupFile2 LookupFileUnionType = "LookupFile_2"
)

type LookupFile struct {
	LookupFile1 *LookupFile1 `queryParam:"inline"`
	LookupFile2 *LookupFile2 `queryParam:"inline"`

	Type LookupFileUnionType
}

func CreateLookupFileLookupFile1(lookupFile1 LookupFile1) LookupFile {
	typ := LookupFileUnionTypeLookupFile1

	return LookupFile{
		LookupFile1: &lookupFile1,
		Type:        typ,
	}
}

func CreateLookupFileLookupFile2(lookupFile2 LookupFile2) LookupFile {
	typ := LookupFileUnionTypeLookupFile2

	return LookupFile{
		LookupFile2: &lookupFile2,
		Type:        typ,
	}
}

func (u *LookupFile) UnmarshalJSON(data []byte) error {

	var lookupFile1 LookupFile1 = LookupFile1{}
	if err := utils.UnmarshalJSON(data, &lookupFile1, "", true, true); err == nil {
		u.LookupFile1 = &lookupFile1
		u.Type = LookupFileUnionTypeLookupFile1
		return nil
	}

	var lookupFile2 LookupFile2 = LookupFile2{}
	if err := utils.UnmarshalJSON(data, &lookupFile2, "", true, true); err == nil {
		u.LookupFile2 = &lookupFile2
		u.Type = LookupFileUnionTypeLookupFile2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for LookupFile", string(data))
}

func (u LookupFile) MarshalJSON() ([]byte, error) {
	if u.LookupFile1 != nil {
		return utils.MarshalJSON(u.LookupFile1, "", true)
	}

	if u.LookupFile2 != nil {
		return utils.MarshalJSON(u.LookupFile2, "", true)
	}

	return nil, errors.New("could not marshal union type LookupFile: all fields are null")
}

type LookupFile2Input struct {
	// File content.
	Content     *string `json:"content,omitempty"`
	ID          string  `json:"id"`
	Description *string `json:"description,omitempty"`
	// One or more tags related to this lookup. Optional.
	Tags *string `json:"tags,omitempty"`
	// File size. Optional.
	Size *float64 `json:"size,omitempty"`
	// Operation mode for CSV-based lookups
	Mode *LookupFile2Mode `default:"memory" json:"mode"`
}

func (l LookupFile2Input) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LookupFile2Input) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *LookupFile2Input) GetContent() *string {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *LookupFile2Input) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *LookupFile2Input) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *LookupFile2Input) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *LookupFile2Input) GetSize() *float64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *LookupFile2Input) GetMode() *LookupFile2Mode {
	if o == nil {
		return nil
	}
	return o.Mode
}

type LookupFile1Input struct {
	FileInfo    *FileInfo `json:"fileInfo,omitempty"`
	ID          string    `json:"id"`
	Description *string   `json:"description,omitempty"`
	// One or more tags related to this lookup. Optional.
	Tags *string `json:"tags,omitempty"`
	// File size. Optional.
	Size *float64 `json:"size,omitempty"`
	// Operation mode for CSV-based lookups
	Mode *LookupFileMode `default:"memory" json:"mode"`
}

func (l LookupFile1Input) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LookupFile1Input) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *LookupFile1Input) GetFileInfo() *FileInfo {
	if o == nil {
		return nil
	}
	return o.FileInfo
}

func (o *LookupFile1Input) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *LookupFile1Input) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *LookupFile1Input) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *LookupFile1Input) GetSize() *float64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *LookupFile1Input) GetMode() *LookupFileMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

type LookupFileInputType string

const (
	LookupFileInputTypeLookupFile1Input LookupFileInputType = "LookupFile_1_input"
	LookupFileInputTypeLookupFile2Input LookupFileInputType = "LookupFile_2_input"
)

type LookupFileInput struct {
	LookupFile1Input *LookupFile1Input `queryParam:"inline"`
	LookupFile2Input *LookupFile2Input `queryParam:"inline"`

	Type LookupFileInputType
}

func CreateLookupFileInputLookupFile1Input(lookupFile1Input LookupFile1Input) LookupFileInput {
	typ := LookupFileInputTypeLookupFile1Input

	return LookupFileInput{
		LookupFile1Input: &lookupFile1Input,
		Type:             typ,
	}
}

func CreateLookupFileInputLookupFile2Input(lookupFile2Input LookupFile2Input) LookupFileInput {
	typ := LookupFileInputTypeLookupFile2Input

	return LookupFileInput{
		LookupFile2Input: &lookupFile2Input,
		Type:             typ,
	}
}

func (u *LookupFileInput) UnmarshalJSON(data []byte) error {

	var lookupFile1Input LookupFile1Input = LookupFile1Input{}
	if err := utils.UnmarshalJSON(data, &lookupFile1Input, "", true, true); err == nil {
		u.LookupFile1Input = &lookupFile1Input
		u.Type = LookupFileInputTypeLookupFile1Input
		return nil
	}

	var lookupFile2Input LookupFile2Input = LookupFile2Input{}
	if err := utils.UnmarshalJSON(data, &lookupFile2Input, "", true, true); err == nil {
		u.LookupFile2Input = &lookupFile2Input
		u.Type = LookupFileInputTypeLookupFile2Input
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for LookupFileInput", string(data))
}

func (u LookupFileInput) MarshalJSON() ([]byte, error) {
	if u.LookupFile1Input != nil {
		return utils.MarshalJSON(u.LookupFile1Input, "", true)
	}

	if u.LookupFile2Input != nil {
		return utils.MarshalJSON(u.LookupFile2Input, "", true)
	}

	return nil, errors.New("could not marshal union type LookupFileInput: all fields are null")
}
