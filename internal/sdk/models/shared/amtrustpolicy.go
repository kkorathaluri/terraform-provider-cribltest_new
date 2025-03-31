// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"errors"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
)

type ActionType string

const (
	ActionTypeStr        ActionType = "str"
	ActionTypeArrayOfStr ActionType = "arrayOfStr"
)

type Action struct {
	Str        *string  `queryParam:"inline"`
	ArrayOfStr []string `queryParam:"inline"`

	Type ActionType
}

func CreateActionStr(str string) Action {
	typ := ActionTypeStr

	return Action{
		Str:  &str,
		Type: typ,
	}
}

func CreateActionArrayOfStr(arrayOfStr []string) Action {
	typ := ActionTypeArrayOfStr

	return Action{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *Action) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = ActionTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = ActionTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Action", string(data))
}

func (u Action) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type Action: all fields are null")
}

type StringEquals struct {
}

type AMTrustPolicyCondition struct {
	StringEquals *StringEquals `json:"StringEquals,omitempty"`
}

func (o *AMTrustPolicyCondition) GetStringEquals() *StringEquals {
	if o == nil {
		return nil
	}
	return o.StringEquals
}

type Principal struct {
	Aws string `json:"AWS"`
}

func (o *Principal) GetAws() string {
	if o == nil {
		return ""
	}
	return o.Aws
}

type Statement struct {
	Action    Action                  `json:"Action"`
	Condition *AMTrustPolicyCondition `json:"Condition,omitempty"`
	Effect    string                  `json:"Effect"`
	Principal Principal               `json:"Principal"`
}

func (o *Statement) GetAction() Action {
	if o == nil {
		return Action{}
	}
	return o.Action
}

func (o *Statement) GetCondition() *AMTrustPolicyCondition {
	if o == nil {
		return nil
	}
	return o.Condition
}

func (o *Statement) GetEffect() string {
	if o == nil {
		return ""
	}
	return o.Effect
}

func (o *Statement) GetPrincipal() Principal {
	if o == nil {
		return Principal{}
	}
	return o.Principal
}

type AMTrustPolicy struct {
	Statement []Statement `json:"Statement"`
	Version   string      `json:"Version"`
}

func (o *AMTrustPolicy) GetStatement() []Statement {
	if o == nil {
		return []Statement{}
	}
	return o.Statement
}

func (o *AMTrustPolicy) GetVersion() string {
	if o == nil {
		return ""
	}
	return o.Version
}
