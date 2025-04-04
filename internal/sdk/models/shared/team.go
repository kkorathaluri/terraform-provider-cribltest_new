// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Team struct {
	Description string   `json:"description"`
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Roles       []string `json:"roles"`
	SsoGroupIds []string `json:"ssoGroupIds,omitempty"`
}

func (o *Team) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *Team) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Team) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Team) GetRoles() []string {
	if o == nil {
		return []string{}
	}
	return o.Roles
}

func (o *Team) GetSsoGroupIds() []string {
	if o == nil {
		return nil
	}
	return o.SsoGroupIds
}
