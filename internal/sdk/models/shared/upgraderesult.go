// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type UpgradeResult struct {
	AvailableVersions []VersionInfo `json:"availableVersions"`
	CanUpgrade        bool          `json:"canUpgrade"`
	InstalledVersion  VersionInfo   `json:"installedVersion"`
	IsSuccess         bool          `json:"isSuccess"`
	Message           string        `json:"message"`
	UpgradedToVersion VersionInfo   `json:"upgradedToVersion"`
}

func (o *UpgradeResult) GetAvailableVersions() []VersionInfo {
	if o == nil {
		return []VersionInfo{}
	}
	return o.AvailableVersions
}

func (o *UpgradeResult) GetCanUpgrade() bool {
	if o == nil {
		return false
	}
	return o.CanUpgrade
}

func (o *UpgradeResult) GetInstalledVersion() VersionInfo {
	if o == nil {
		return VersionInfo{}
	}
	return o.InstalledVersion
}

func (o *UpgradeResult) GetIsSuccess() bool {
	if o == nil {
		return false
	}
	return o.IsSuccess
}

func (o *UpgradeResult) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *UpgradeResult) GetUpgradedToVersion() VersionInfo {
	if o == nil {
		return VersionInfo{}
	}
	return o.UpgradedToVersion
}
