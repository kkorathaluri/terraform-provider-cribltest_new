// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type HTTPOutExtraParamConfig struct {
	Name  string         `json:"name"`
	Value map[string]any `json:"value"`
}

func (o *HTTPOutExtraParamConfig) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *HTTPOutExtraParamConfig) GetValue() map[string]any {
	if o == nil {
		return map[string]any{}
	}
	return o.Value
}
