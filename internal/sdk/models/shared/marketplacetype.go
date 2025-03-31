// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type MarketplaceType string

const (
	MarketplaceTypeAws MarketplaceType = "aws"
)

func (e MarketplaceType) ToPointer() *MarketplaceType {
	return &e
}
func (e *MarketplaceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "aws":
		*e = MarketplaceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MarketplaceType: %v", v)
	}
}
