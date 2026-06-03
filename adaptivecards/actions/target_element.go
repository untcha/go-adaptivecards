package actions

import (
	"encoding/json"
	"fmt"
	"strings"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

const targetElementType = "TargetElement"

// TargetElement represents an entry for Action.ToggleVisibility's targetElements property.
// See: https://adaptivecards.io/explorer/TargetElement.html
type TargetElement struct {
	Type      m.TypeString `json:"type,omitempty"`
	ElementID string       `json:"elementId"`
	IsVisible *bool        `json:"isVisible,omitempty"`
}

// UnmarshalJSON decodes a TargetElement from either a bare elementId string or a full object.
func (t *TargetElement) UnmarshalJSON(b []byte) error {
	var id string
	if err := json.Unmarshal(b, &id); err == nil {
		id = strings.TrimSpace(id)
		if id == "" {
			return fmt.Errorf("targetElement.elementId is required")
		}
		*t = TargetElement{ElementID: id}
		return nil
	}

	type alias TargetElement
	var tmp alias
	if err := json.Unmarshal(b, &tmp); err != nil {
		return fmt.Errorf("targetElement: decode: %w", err)
	}
	if tmp.Type != "" && tmp.Type != targetElementType {
		return fmt.Errorf("targetElement.type must be %q (got %q)", targetElementType, tmp.Type)
	}
	tmp.ElementID = strings.TrimSpace(tmp.ElementID)
	if tmp.ElementID == "" {
		return fmt.Errorf("targetElement.elementId is required")
	}
	*t = TargetElement(tmp)
	return nil
}
