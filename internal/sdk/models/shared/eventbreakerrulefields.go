// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type EventBreakerRuleFields struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (o *EventBreakerRuleFields) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *EventBreakerRuleFields) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}
