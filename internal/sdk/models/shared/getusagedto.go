// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type GetUsageDTO struct {
	StartingOn   string               `json:"startingOn"`
	EndingBefore string               `json:"endingBefore"`
	Data         UsageData            `json:"data"`
	GroupedData  map[string]UsageData `json:"groupedData,omitempty"`
}

func (o *GetUsageDTO) GetStartingOn() string {
	if o == nil {
		return ""
	}
	return o.StartingOn
}

func (o *GetUsageDTO) GetEndingBefore() string {
	if o == nil {
		return ""
	}
	return o.EndingBefore
}

func (o *GetUsageDTO) GetData() UsageData {
	if o == nil {
		return UsageData{}
	}
	return o.Data
}

func (o *GetUsageDTO) GetGroupedData() map[string]UsageData {
	if o == nil {
		return nil
	}
	return o.GroupedData
}
