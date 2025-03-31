// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type AppScopeProcessStatus string

const (
	AppScopeProcessStatusGreen  AppScopeProcessStatus = "green"
	AppScopeProcessStatusYellow AppScopeProcessStatus = "yellow"
	AppScopeProcessStatusRed    AppScopeProcessStatus = "red"
)

func (e AppScopeProcessStatus) ToPointer() *AppScopeProcessStatus {
	return &e
}
func (e *AppScopeProcessStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "green":
		fallthrough
	case "yellow":
		fallthrough
	case "red":
		*e = AppScopeProcessStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AppScopeProcessStatus: %v", v)
	}
}
