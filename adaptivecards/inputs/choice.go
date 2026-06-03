package inputs

import (
	"fmt"
	"strings"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// InputChoice represents an Input.Choice that describes a choice for use in a ChoiceSet.
// See: https://adaptivecards.io/explorer/Input.Choice.html
type InputChoice struct {
	Type  m.TypeString `json:"type,omitempty"`
	Title string       `json:"title"`
	Value string       `json:"value"`
}

// NewInputChoice creates a choice option.
func NewInputChoice(title, value string) InputChoice {
	return InputChoice{
		Type:  m.TypeInputChoice,
		Title: title,
		Value: value,
	}
}

// Validate checks that the choice contains a title and value.
func (c InputChoice) Validate() error {
	if c.Type != "" && c.Type != m.TypeInputChoice {
		return fmt.Errorf("input.choice.type must be %q (got %q)", m.TypeInputChoice, c.Type)
	}
	if strings.TrimSpace(c.Title) == "" {
		return fmt.Errorf("input.choice.title is required")
	}
	if strings.TrimSpace(c.Value) == "" {
		return fmt.Errorf("input.choice.value is required")
	}
	return nil
}
