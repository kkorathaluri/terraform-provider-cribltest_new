// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CriblLib string

const (
	CriblLibCribl       CriblLib = "cribl"
	CriblLibCriblCustom CriblLib = "cribl-custom"
	CriblLibCustom      CriblLib = "custom"
)

func (e CriblLib) ToPointer() *CriblLib {
	return &e
}
func (e *CriblLib) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cribl":
		fallthrough
	case "cribl-custom":
		fallthrough
	case "custom":
		*e = CriblLib(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CriblLib: %v", v)
	}
}
