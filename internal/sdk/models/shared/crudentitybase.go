// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CrudEntityBase struct {
	ID string `json:"id"`
}

func (o *CrudEntityBase) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}
