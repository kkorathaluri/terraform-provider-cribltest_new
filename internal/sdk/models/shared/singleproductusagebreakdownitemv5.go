// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
	"time"
)

type SingleProductUsageBreakdownItemV5 struct {
	Date    time.Time                                  `json:"date"`
	Credits float64                                    `json:"credits"`
	PerUnit []SingleProductUsageBreakdownPerUnitItemV5 `json:"perUnit"`
}

func (s SingleProductUsageBreakdownItemV5) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SingleProductUsageBreakdownItemV5) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SingleProductUsageBreakdownItemV5) GetDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Date
}

func (o *SingleProductUsageBreakdownItemV5) GetCredits() float64 {
	if o == nil {
		return 0.0
	}
	return o.Credits
}

func (o *SingleProductUsageBreakdownItemV5) GetPerUnit() []SingleProductUsageBreakdownPerUnitItemV5 {
	if o == nil {
		return []SingleProductUsageBreakdownPerUnitItemV5{}
	}
	return o.PerUnit
}
