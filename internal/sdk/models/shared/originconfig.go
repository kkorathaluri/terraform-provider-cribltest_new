// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type OriginConfig struct {
	FilterExpression *string       `json:"filterExpression,omitempty"`
	Origin           DatasetOrigin `json:"origin"`
}

func (o *OriginConfig) GetFilterExpression() *string {
	if o == nil {
		return nil
	}
	return o.FilterExpression
}

func (o *OriginConfig) GetOrigin() DatasetOrigin {
	if o == nil {
		return DatasetOrigin("")
	}
	return o.Origin
}
