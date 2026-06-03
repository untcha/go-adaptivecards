package containers

import (
	"encoding/json"
	"fmt"
	"strings"

	a "github.com/untcha/go-adaptivecards/adaptivecards/actions"
	e "github.com/untcha/go-adaptivecards/adaptivecards/core/element"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// ActionSet displays a set of actions.
// See: https://adaptivecards.io/explorer/ActionSet.html
type ActionSet struct {
	e.ElementBase              // Embedding e.ElementBase to include common element fields
	Type          m.TypeString `json:"type"`              // Version 1.2
	Actions       []a.Action   `json:"actions,omitempty"` // Version 1.2
}

// NewActionSet returns an ActionSet holding the given actions.
func NewActionSet(actions ...a.Action) ActionSet {
	return ActionSet{
		Type:    m.TypeActionSet,
		Actions: actions,
	}
}

// GetType returns the Adaptive Card type discriminator for ActionSet.
func (as ActionSet) GetType() m.TypeString { return m.TypeActionSet }

// Validate reports whether the ActionSet and its actions are well-formed.
func (as ActionSet) Validate() error {
	if err := as.validateElementBase(); err != nil {
		return err
	}
	if err := validateActions(as.Actions); err != nil {
		return fmt.Errorf("actionSet.actions: %w", err)
	}
	return nil
}

// MarshalJSON implements json.Marshaler, ensuring the type field is set.
func (as ActionSet) MarshalJSON() ([]byte, error) {
	aa := as
	if aa.Type == "" {
		aa.Type = m.TypeActionSet
	}
	type alias ActionSet
	return json.Marshal(alias(aa))
}

// UnmarshalJSON implements json.Unmarshaler, decoding interface-typed fields via the element/action factories.
func (as *ActionSet) UnmarshalJSON(b []byte) error {
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(b, &obj); err != nil {
		return fmt.Errorf("actionSet: decode: %w", err)
	}
	if v, ok := obj["type"]; ok {
		var t string
		if err := json.Unmarshal(v, &t); err != nil {
			return fmt.Errorf("actionSet.type: %w", err)
		}
		if t != "" && t != string(m.TypeActionSet) {
			return fmt.Errorf("actionSet.type must be %q (got %q)", m.TypeActionSet, t)
		}
	}
	actionsRaw := obj["actions"]
	delete(obj, "actions")

	type alias ActionSet
	var base alias
	baseBytes, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("actionSet: re-encode base: %w", err)
	}
	if err := json.Unmarshal(baseBytes, &base); err != nil {
		return fmt.Errorf("actionSet: decode base: %w", err)
	}
	if base.Type == "" {
		base.Type = m.TypeActionSet
	}
	if len(actionsRaw) != 0 {
		var raw []json.RawMessage
		if err := json.Unmarshal(actionsRaw, &raw); err != nil {
			return fmt.Errorf("actionSet.actions: %w", err)
		}
		actions, err := a.UnmarshalActionsSlice(raw)
		if err != nil {
			return fmt.Errorf("actionSet.actions: %w", err)
		}
		base.Actions = actions
	}

	*as = ActionSet(base)
	return nil
}

func (as ActionSet) validateElementBase() error {
	if as.Height != "" && !as.Height.IsValid() {
		return m.NewEnumError(
			"ActionSet.height",
			string(as.Height),
			m.AllowedBlockElementHeightStrings(),
		)
	}
	if as.Spacing != "" && !as.Spacing.IsValid() {
		return m.NewEnumError("ActionSet.spacing", string(as.Spacing), m.AllowedSpacingStrings())
	}
	if as.ID != "" {
		id := strings.TrimSpace(as.ID)
		if id == "" {
			return fmt.Errorf("actionSet.id cannot be empty or whitespace-only")
		}
		if strings.ContainsAny(id, "\n\r\t") {
			return fmt.Errorf("actionSet.id cannot contain newlines or tabs")
		}
	}
	return nil
}

func init() {
	e.RegisterElement(
		m.TypeActionSet,
		func() e.Element { return &ActionSet{Type: m.TypeActionSet} },
	)
}
