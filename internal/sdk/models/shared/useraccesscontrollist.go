// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type UserAccessControlList struct {
	Perms []ResourcePolicy `json:"perms"`
	User  string           `json:"user"`
}

func (o *UserAccessControlList) GetPerms() []ResourcePolicy {
	if o == nil {
		return []ResourcePolicy{}
	}
	return o.Perms
}

func (o *UserAccessControlList) GetUser() string {
	if o == nil {
		return ""
	}
	return o.User
}
