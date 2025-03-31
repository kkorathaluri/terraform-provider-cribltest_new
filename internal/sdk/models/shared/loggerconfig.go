// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type LoggerConfig struct {
	Channels            []LoggerEntry `json:"channels"`
	DefaultRedactFields []string      `json:"defaultRedactFields,omitempty"`
	ID                  string        `json:"id"`
	LimitRate           float64       `json:"limitRate"`
	MaxSizeBytes        float64       `json:"maxSizeBytes"`
	RedactFields        []string      `json:"redactFields"`
	RedactLabel         string        `json:"redactLabel"`
}

func (o *LoggerConfig) GetChannels() []LoggerEntry {
	if o == nil {
		return []LoggerEntry{}
	}
	return o.Channels
}

func (o *LoggerConfig) GetDefaultRedactFields() []string {
	if o == nil {
		return nil
	}
	return o.DefaultRedactFields
}

func (o *LoggerConfig) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *LoggerConfig) GetLimitRate() float64 {
	if o == nil {
		return 0.0
	}
	return o.LimitRate
}

func (o *LoggerConfig) GetMaxSizeBytes() float64 {
	if o == nil {
		return 0.0
	}
	return o.MaxSizeBytes
}

func (o *LoggerConfig) GetRedactFields() []string {
	if o == nil {
		return []string{}
	}
	return o.RedactFields
}

func (o *LoggerConfig) GetRedactLabel() string {
	if o == nil {
		return ""
	}
	return o.RedactLabel
}
