// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Metadata struct {
	Aws       *AwsMetadata   `json:"aws,omitempty"`
	Cribl     *CriblMetadata `json:"cribl,omitempty"`
	Env       *EnvMetadata   `json:"env,omitempty"`
	HostOs    *OsMetadata    `json:"hostOs,omitempty"`
	Kube      *KubeMetadata  `json:"kube,omitempty"`
	Os        *OsMetadata    `json:"os,omitempty"`
	Timestamp float64        `json:"timestamp"`
}

func (o *Metadata) GetAws() *AwsMetadata {
	if o == nil {
		return nil
	}
	return o.Aws
}

func (o *Metadata) GetCribl() *CriblMetadata {
	if o == nil {
		return nil
	}
	return o.Cribl
}

func (o *Metadata) GetEnv() *EnvMetadata {
	if o == nil {
		return nil
	}
	return o.Env
}

func (o *Metadata) GetHostOs() *OsMetadata {
	if o == nil {
		return nil
	}
	return o.HostOs
}

func (o *Metadata) GetKube() *KubeMetadata {
	if o == nil {
		return nil
	}
	return o.Kube
}

func (o *Metadata) GetOs() *OsMetadata {
	if o == nil {
		return nil
	}
	return o.Os
}

func (o *Metadata) GetTimestamp() float64 {
	if o == nil {
		return 0.0
	}
	return o.Timestamp
}
