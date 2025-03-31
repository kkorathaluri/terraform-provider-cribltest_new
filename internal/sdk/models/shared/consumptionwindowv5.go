// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConsumptionWindowV5 string

const (
	ConsumptionWindowV5Monthly ConsumptionWindowV5 = "monthly"
)

func (e ConsumptionWindowV5) ToPointer() *ConsumptionWindowV5 {
	return &e
}
func (e *ConsumptionWindowV5) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "monthly":
		*e = ConsumptionWindowV5(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConsumptionWindowV5: %v", v)
	}
}
