// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type NotificationTarget struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func (o *NotificationTarget) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *NotificationTarget) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}
