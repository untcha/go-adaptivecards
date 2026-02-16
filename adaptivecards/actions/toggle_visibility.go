package actions

import (
	"encoding/json"
	"fmt"
	"strings"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// Action.ToggleVisibility
// An action that toggles the visibility of associated card elements.
// See: https://adaptivecards.io/explorer/Action.ToggleVisibility.html
type ActionToggleVisibility struct {
	ActionBase
	Type           m.TypeString    `json:"type"`                     // Version 1.2
	TargetElements []TargetElement `json:"targetElements,omitempty"` // Version 1.2
}

func (a ActionToggleVisibility) GetType() m.TypeString { return m.TypeActionToggleVisibility }

func NewActionToggleVisibility(title string, targets ...TargetElement) ActionToggleVisibility {
	return ActionToggleVisibility{
		ActionBase:     ActionBase{Title: title},
		Type:           m.TypeActionToggleVisibility,
		TargetElements: targets,
	}
}

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

func (a ActionToggleVisibility) MarshalJSON() ([]byte, error) {
	aa := a
	if aa.Type == "" {
		aa.Type = m.TypeActionToggleVisibility
	}
	type alias ActionToggleVisibility
	return json.Marshal(alias(aa))
}

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
