// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SearchJobErrorStateConfig struct {
	Coordinated   bool     `json:"coordinated"`
	ErrorMessages []string `json:"errorMessages"`
}

func (o *SearchJobErrorStateConfig) GetCoordinated() bool {
	if o == nil {
		return false
	}
	return o.Coordinated
}

func (o *SearchJobErrorStateConfig) GetErrorMessages() []string {
	if o == nil {
		return []string{}
	}
	return o.ErrorMessages
}
