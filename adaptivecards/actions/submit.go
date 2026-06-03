package actions

import (
	"encoding/json"
	"fmt"
	"strings"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// ActionSubmit represents an Action.Submit: it gathers input fields, merges with optional data field, and sends an event to the client.
// It is up to the client to determine how this data is processed.
// For example: With BotFramework bots, the client would send an activity through the messaging medium to the bot.
// The inputs that are gathered are those on the current card, and in the case of a show card those on any parent cards.
// See https://docs.microsoft.com/en-us/adaptive-cards/authoring-cards/input-validation for more details.
// See: https://adaptivecards.io/explorer/Action.Submit.html
type ActionSubmit struct {
	ActionBase
	Type             m.TypeString `json:"type"`                       // Version 1.0
	Data             any          `json:"data,omitempty"`             // Version 1.0
	AssociatedInputs string       `json:"associatedInputs,omitempty"` // Version 1.3
}

// NewActionSubmit creates an Action.Submit with an optional data payload.
func NewActionSubmit(title string, data any) ActionSubmit {
	return ActionSubmit{
		ActionBase: ActionBase{Title: title},
		Type:       m.TypeActionSubmit,
		Data:       data,
	}
}

// GetType returns the action type string for Action.Submit.
func (a ActionSubmit) GetType() m.TypeString { return m.TypeActionSubmit }

// Validate checks that AssociatedInputs, if set, is "auto" or "none".
func (a ActionSubmit) Validate() error {
	if a.AssociatedInputs != "" {
		if !strings.EqualFold(a.AssociatedInputs, "auto") &&
			!strings.EqualFold(a.AssociatedInputs, "none") {
			return fmt.Errorf(
				"action.submit.associatedInputs must be auto or none (got %q)",
				a.AssociatedInputs,
			)
		}
	}
	return nil
}

// MarshalJSON encodes the action, ensuring the Type field is always set.
func (a ActionSubmit) MarshalJSON() ([]byte, error) {
	aa := a
	if aa.Type == "" {
		aa.Type = m.TypeActionSubmit
	}
	type alias ActionSubmit
	return json.Marshal(alias(aa))
}

// UnmarshalJSON decodes the action, verifying the type and validating the result.
func (a *ActionSubmit) UnmarshalJSON(b []byte) error {
	type alias ActionSubmit
	var tmp alias
	if err := json.Unmarshal(b, &tmp); err != nil {
		return fmt.Errorf("action.submit: decode: %w", err)
	}
	if tmp.Type != "" && tmp.Type != m.TypeActionSubmit {
		return fmt.Errorf("action.submit.type must be %q (got %q)", m.TypeActionSubmit, tmp.Type)
	}
	tmp.Type = m.TypeActionSubmit

	s := strings.TrimSpace(tmp.AssociatedInputs)
	if s != "" {
		tmp.AssociatedInputs = s
	}
	if err := ActionSubmit(tmp).Validate(); err != nil {
		return err
	}
	*a = ActionSubmit(tmp)
	return nil
}

func init() {
	RegisterAction(
		m.TypeActionSubmit,
		func() Action { return &ActionSubmit{Type: m.TypeActionSubmit} },
	)
}
