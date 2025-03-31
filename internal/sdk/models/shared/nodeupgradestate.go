// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type NodeUpgradeState int64

const (
	NodeUpgradeStateZero  NodeUpgradeState = 0
	NodeUpgradeStateOne   NodeUpgradeState = 1
	NodeUpgradeStateTwo   NodeUpgradeState = 2
	NodeUpgradeStateThree NodeUpgradeState = 3
)

func (e NodeUpgradeState) ToPointer() *NodeUpgradeState {
	return &e
}
func (e *NodeUpgradeState) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		*e = NodeUpgradeState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NodeUpgradeState: %v", v)
	}
}
