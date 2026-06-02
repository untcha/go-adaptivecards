package containers

import (
	"encoding/json"
	"fmt"

	a "github.com/untcha/go-adaptivecards/adaptivecards/actions"
	e "github.com/untcha/go-adaptivecards/adaptivecards/core/element"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// ActionSet
// Displays a set of actions.
// See: https://adaptivecards.io/explorer/ActionSet.html
type ActionSet struct {
	e.ElementBase              // Embedding e.ElementBase to include common element fields
	Type          m.TypeString `json:"type"`              // Version 1.2
	Actions       []a.Action   `json:"actions,omitempty"` // Version 1.2
}

func NewActionSet(actions ...a.Action) ActionSet {
	return ActionSet{
		Type:    m.TypeActionSet,
		Actions: actions,
	}
}

func (as ActionSet) GetType() m.TypeString { return m.TypeActionSet }

func (as ActionSet) Validate() error {
	if err := as.validateElementBase(); err != nil {
		return err
	}
	if err := validateActions(as.Actions); err != nil {
		return fmt.Errorf("actionSet.actions: %w", err)
	}
	return nil
}

func (as ActionSet) MarshalJSON() ([]byte, error) {
	aa := as
	if aa.Type == "" {
		aa.Type = m.TypeActionSet
	}
	type alias ActionSet
	return json.Marshal(alias(aa))
}

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
	return as.ElementBase.Validate("ActionSet")
}

func init() {
	e.RegisterElement(
		m.TypeActionSet,
		func() e.Element { return &ActionSet{Type: m.TypeActionSet} },
	)
}
