// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type TrustPolicy struct {
	ID     string        `json:"id"`
	Policy AMTrustPolicy `json:"policy"`
}

func (o *TrustPolicy) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *TrustPolicy) GetPolicy() AMTrustPolicy {
	if o == nil {
		return AMTrustPolicy{}
	}
	return o.Policy
}
