package inputs

import (
	"encoding/json"
	"fmt"
	"strings"

	e "github.com/untcha/go-adaptivecards/adaptivecards/core/element"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// Input.Toggle
// Lets a user choose between two options.
// See: https://adaptivecards.io/explorer/Input.Toggle.html
type InputToggle struct {
	InputBase
	Type     m.TypeString `json:"type"`               // Version 1.0
	Title    string       `json:"title"`              // Version 1.0
	Value    string       `json:"value,omitempty"`    // Version 1.0
	ValueOff string       `json:"valueOff,omitempty"` // Version 1.0
	ValueOn  string       `json:"valueOn,omitempty"`  // Version 1.0
	Wrap     bool         `json:"wrap,omitempty"`     // Version 1.2
}

func (i InputToggle) GetType() m.TypeString { return m.TypeInputToggle }

func NewInputToggle(id, title string) InputToggle {
	return InputToggle{
		InputBase: InputBase{ElementBase: e.ElementBase{ID: id}},
		Type:      m.TypeInputToggle,
		Title:     title,
		ValueOff:  "false",
		ValueOn:   "true",
	}
}

func (i InputToggle) Validate() error {
	if err := i.validateInputBase("input.toggle"); err != nil {
		return err
	}
	if strings.TrimSpace(i.Title) == "" {
		return fmt.Errorf("input.toggle.title is required")
	}
	return nil
}

func (i InputToggle) MarshalJSON() ([]byte, error) {
	ii := i
	if ii.Type == "" {
		ii.Type = m.TypeInputToggle
	}
	if ii.ValueOff == "" {
		ii.ValueOff = "false"
	}
	if ii.ValueOn == "" {
		ii.ValueOn = "true"
	}
	type alias InputToggle
	return json.Marshal(alias(ii))
}

func (i *InputToggle) UnmarshalJSON(b []byte) error {
	type alias InputToggle
	var tmp alias
	if err := json.Unmarshal(b, &tmp); err != nil {
		return fmt.Errorf("input.toggle: decode: %w", err)
	}
	if tmp.Type != "" && tmp.Type != m.TypeInputToggle {
		return fmt.Errorf("input.toggle.type must be %q (got %q)", m.TypeInputToggle, tmp.Type)
	}
	if tmp.Type == "" {
		tmp.Type = m.TypeInputToggle
	}
	if tmp.ValueOff == "" {
		tmp.ValueOff = "false"
	}
	if tmp.ValueOn == "" {
		tmp.ValueOn = "true"
	}
	val := InputToggle(tmp)
	if err := val.Validate(); err != nil {
		return err
	}
	*i = val
	return nil
}

func init() {
	e.RegisterElement(
		m.TypeInputToggle,
		func() e.Element { return &InputToggle{Type: m.TypeInputToggle} },
	)
}
