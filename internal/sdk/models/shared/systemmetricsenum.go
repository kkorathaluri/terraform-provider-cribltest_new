// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"errors"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

type DimKeyFilterType string

const (
	DimKeyFilterTypeStr        DimKeyFilterType = "str"
	DimKeyFilterTypeArrayOfStr DimKeyFilterType = "arrayOfStr"
)

type DimKeyFilter struct {
	Str        *string  `queryParam:"inline"`
	ArrayOfStr []string `queryParam:"inline"`

	Type DimKeyFilterType
}

func CreateDimKeyFilterStr(str string) DimKeyFilter {
	typ := DimKeyFilterTypeStr

	return DimKeyFilter{
		Str:  &str,
		Type: typ,
	}
}

func CreateDimKeyFilterArrayOfStr(arrayOfStr []string) DimKeyFilter {
	typ := DimKeyFilterTypeArrayOfStr

	return DimKeyFilter{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *DimKeyFilter) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = DimKeyFilterTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = DimKeyFilterTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DimKeyFilter", string(data))
}

func (u DimKeyFilter) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type DimKeyFilter: all fields are null")
}

type DimValueFilterType string

const (
	DimValueFilterTypeStr        DimValueFilterType = "str"
	DimValueFilterTypeArrayOfStr DimValueFilterType = "arrayOfStr"
)

type DimValueFilter struct {
	Str        *string  `queryParam:"inline"`
	ArrayOfStr []string `queryParam:"inline"`

	Type DimValueFilterType
}

func CreateDimValueFilterStr(str string) DimValueFilter {
	typ := DimValueFilterTypeStr

	return DimValueFilter{
		Str:  &str,
		Type: typ,
	}
}

func CreateDimValueFilterArrayOfStr(arrayOfStr []string) DimValueFilter {
	typ := DimValueFilterTypeArrayOfStr

	return DimValueFilter{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *DimValueFilter) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = DimValueFilterTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = DimValueFilterTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DimValueFilter", string(data))
}

func (u DimValueFilter) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type DimValueFilter: all fields are null")
}

type MetricNameFilterType string

const (
	MetricNameFilterTypeStr        MetricNameFilterType = "str"
	MetricNameFilterTypeArrayOfStr MetricNameFilterType = "arrayOfStr"
)

type MetricNameFilter struct {
	Str        *string  `queryParam:"inline"`
	ArrayOfStr []string `queryParam:"inline"`

	Type MetricNameFilterType
}

func CreateMetricNameFilterStr(str string) MetricNameFilter {
	typ := MetricNameFilterTypeStr

	return MetricNameFilter{
		Str:  &str,
		Type: typ,
	}
}

func CreateMetricNameFilterArrayOfStr(arrayOfStr []string) MetricNameFilter {
	typ := MetricNameFilterTypeArrayOfStr

	return MetricNameFilter{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *MetricNameFilter) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = MetricNameFilterTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = MetricNameFilterTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for MetricNameFilter", string(data))
}

func (u MetricNameFilter) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type MetricNameFilter: all fields are null")
}

type SystemMetricsEnum struct {
	DimKeyFilter     *DimKeyFilter     `json:"dimKeyFilter,omitempty"`
	DimValueFilter   *DimValueFilter   `json:"dimValueFilter,omitempty"`
	Earliest         *float64          `json:"earliest,omitempty"`
	FilterExpr       *string           `json:"filterExpr,omitempty"`
	MaxValues        *float64          `json:"maxValues,omitempty"`
	MetricNameFilter *MetricNameFilter `json:"metricNameFilter,omitempty"`
}

func (o *SystemMetricsEnum) GetDimKeyFilter() *DimKeyFilter {
	if o == nil {
		return nil
	}
	return o.DimKeyFilter
}

func (o *SystemMetricsEnum) GetDimValueFilter() *DimValueFilter {
	if o == nil {
		return nil
	}
	return o.DimValueFilter
}

func (o *SystemMetricsEnum) GetEarliest() *float64 {
	if o == nil {
		return nil
	}
	return o.Earliest
}

func (o *SystemMetricsEnum) GetFilterExpr() *string {
	if o == nil {
		return nil
	}
	return o.FilterExpr
}

func (o *SystemMetricsEnum) GetMaxValues() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxValues
}

func (o *SystemMetricsEnum) GetMetricNameFilter() *MetricNameFilter {
	if o == nil {
		return nil
	}
	return o.MetricNameFilter
}
