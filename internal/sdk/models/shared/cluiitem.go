// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Category string

const (
	CategoryLink Category = "link"
)

func (e Category) ToPointer() *Category {
	return &e
}
func (e *Category) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "link":
		*e = Category(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Category: %v", v)
	}
}

type CluiItem struct {
	Category Category `json:"category"`
	GroupID  *string  `json:"groupId,omitempty"`
	ID       *string  `json:"id,omitempty"`
	Name     *string  `json:"name,omitempty"`
	PackID   *string  `json:"packId,omitempty"`
	SubType  *string  `json:"subType,omitempty"`
	Type     CluiType `json:"type"`
}

func (o *CluiItem) GetCategory() Category {
	if o == nil {
		return Category("")
	}
	return o.Category
}

func (o *CluiItem) GetGroupID() *string {
	if o == nil {
		return nil
	}
	return o.GroupID
}

func (o *CluiItem) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CluiItem) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CluiItem) GetPackID() *string {
	if o == nil {
		return nil
	}
	return o.PackID
}

func (o *CluiItem) GetSubType() *string {
	if o == nil {
		return nil
	}
	return o.SubType
}

func (o *CluiItem) GetType() CluiType {
	if o == nil {
		return CluiType("")
	}
	return o.Type
}
