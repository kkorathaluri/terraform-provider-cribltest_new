// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-cribl-terraform/internal/sdk/internal/utils"
	"time"
)

type State string

const (
	StateAssigned State = "ASSIGNED"
)

func (e State) ToPointer() *State {
	return &e
}
func (e *State) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASSIGNED":
		*e = State(v)
		return nil
	default:
		return fmt.Errorf("invalid value for State: %v", v)
	}
}

type AssignedSandboxAssignmentDTO struct {
	OrganizationID string    `json:"organizationId"`
	WorkspaceID    string    `json:"workspaceId"`
	State          State     `json:"state"`
	UserID         string    `json:"userId"`
	CourseID       string    `json:"courseId"`
	AssignedAt     time.Time `json:"assignedAt"`
	ExpiresAt      time.Time `json:"expiresAt"`
}

func (a AssignedSandboxAssignmentDTO) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AssignedSandboxAssignmentDTO) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AssignedSandboxAssignmentDTO) GetOrganizationID() string {
	if o == nil {
		return ""
	}
	return o.OrganizationID
}

func (o *AssignedSandboxAssignmentDTO) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

func (o *AssignedSandboxAssignmentDTO) GetState() State {
	if o == nil {
		return State("")
	}
	return o.State
}

func (o *AssignedSandboxAssignmentDTO) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *AssignedSandboxAssignmentDTO) GetCourseID() string {
	if o == nil {
		return ""
	}
	return o.CourseID
}

func (o *AssignedSandboxAssignmentDTO) GetAssignedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.AssignedAt
}

func (o *AssignedSandboxAssignmentDTO) GetExpiresAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.ExpiresAt
}
