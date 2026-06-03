package actions

import (
	"encoding/json"
	"fmt"
	"strings"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// ActionToggleVisibility represents an Action.ToggleVisibility that toggles the visibility of associated card elements.
// See: https://adaptivecards.io/explorer/Action.ToggleVisibility.html
type ActionToggleVisibility struct {
	ActionBase
	Type           m.TypeString    `json:"type"`                     // Version 1.2
	TargetElements []TargetElement `json:"targetElements,omitempty"` // Version 1.2
}

// NewActionToggleVisibility creates an Action.ToggleVisibility with the given title and target elements.
func NewActionToggleVisibility(title string, targets ...TargetElement) ActionToggleVisibility {
	return ActionToggleVisibility{
		ActionBase:     ActionBase{Title: title},
		Type:           m.TypeActionToggleVisibility,
		TargetElements: targets,
	}
}

// GetType returns the action type string for Action.ToggleVisibility.
func (a ActionToggleVisibility) GetType() m.TypeString { return m.TypeActionToggleVisibility }

// Validate checks that at least one target element is present and that each has a valid elementId and type.
func (a ActionToggleVisibility) Validate() error {
	if len(a.TargetElements) == 0 {
		return fmt.Errorf("action.toggleVisibility.targetElements is required")
	}
	for i, t := range a.TargetElements {
		if strings.TrimSpace(t.ElementID) == "" {
			return fmt.Errorf("action.toggleVisibility.targetElements[%d].elementId is required", i)
		}
		if t.Type != "" && t.Type != targetElementType {
			return fmt.Errorf(
				"action.toggleVisibility.targetElements[%d].type must be %q (got %q)",
				i,
				targetElementType,
				t.Type,
			)
		}
	}
	return nil
}

// MarshalJSON encodes the action, ensuring the Type field is always set.
func (a ActionToggleVisibility) MarshalJSON() ([]byte, error) {
	aa := a
	if aa.Type == "" {
		aa.Type = m.TypeActionToggleVisibility
	}
	type alias ActionToggleVisibility
	return json.Marshal(alias(aa))
}

// UnmarshalJSON decodes the action, verifying the type and validating the result.
func (a *ActionToggleVisibility) UnmarshalJSON(b []byte) error {
	type alias ActionToggleVisibility
	var tmp alias
	if err := json.Unmarshal(b, &tmp); err != nil {
		return fmt.Errorf("action.toggleVisibility: decode: %w", err)
	}
	if tmp.Type != "" && tmp.Type != m.TypeActionToggleVisibility {
		return fmt.Errorf(
			"action.toggleVisibility.type must be %q (got %q)",
			m.TypeActionToggleVisibility,
			tmp.Type,
		)
	}
	tmp.Type = m.TypeActionToggleVisibility
	val := ActionToggleVisibility(tmp)
	if err := val.Validate(); err != nil {
		return err
	}
	*a = val
	return nil
}

func init() {
	RegisterAction(m.TypeActionToggleVisibility, func() Action {
		return &ActionToggleVisibility{Type: m.TypeActionToggleVisibility}
	})
}
