// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type KMSHealth struct {
	Auth       KMSHealthTest `json:"auth"`
	Connection KMSHealthTest `json:"connection"`
	System     KMSHealthTest `json:"system"`
}

func (o *KMSHealth) GetAuth() KMSHealthTest {
	if o == nil {
		return KMSHealthTest{}
	}
	return o.Auth
}

func (o *KMSHealth) GetConnection() KMSHealthTest {
	if o == nil {
		return KMSHealthTest{}
	}
	return o.Connection
}

func (o *KMSHealth) GetSystem() KMSHealthTest {
	if o == nil {
		return KMSHealthTest{}
	}
	return o.System
}
