// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Breakdown struct {
	InfrastructureCredits float64 `json:"infrastructureCredits"`
	ProcessingCredits     float64 `json:"processingCredits"`
	SearchCredits         float64 `json:"searchCredits"`
	LakeCredits           float64 `json:"lakeCredits"`
	LakehouseCredits      float64 `json:"lakehouseCredits"`
}

func (o *Breakdown) GetInfrastructureCredits() float64 {
	if o == nil {
		return 0.0
	}
	return o.InfrastructureCredits
}

func (o *Breakdown) GetProcessingCredits() float64 {
	if o == nil {
		return 0.0
	}
	return o.ProcessingCredits
}

func (o *Breakdown) GetSearchCredits() float64 {
	if o == nil {
		return 0.0
	}
	return o.SearchCredits
}

func (o *Breakdown) GetLakeCredits() float64 {
	if o == nil {
		return 0.0
	}
	return o.LakeCredits
}

func (o *Breakdown) GetLakehouseCredits() float64 {
	if o == nil {
		return 0.0
	}
	return o.LakehouseCredits
}
