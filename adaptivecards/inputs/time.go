package inputs

import (
	"encoding/json"
	"fmt"

	e "github.com/untcha/go-adaptivecards/adaptivecards/core/element"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// InputTime represents an Input.Time element that lets a user select a time.
// See: https://adaptivecards.io/explorer/Input.Time.html
type InputTime struct {
	InputBase
	Type        m.TypeString `json:"type"`                  // Version 1.0
	Max         string       `json:"max,omitempty"`         // Version 1.0
	Min         string       `json:"min,omitempty"`         // Version 1.0
	Placeholder string       `json:"placeholder,omitempty"` // Version 1.0
	Value       string       `json:"value,omitempty"`       // Version 1.0
}

// NewInputTime creates an Input.Time with the specified id.
func NewInputTime(id string) InputTime {
	return InputTime{
		InputBase: InputBase{ElementBase: e.ElementBase{ID: id}},
		Type:      m.TypeInputTime,
	}
}

// GetType returns the Adaptive Card type discriminator for Input.Time.
func (i InputTime) GetType() m.TypeString { return m.TypeInputTime }

// Validate checks that the InputTime has valid fields.
func (i InputTime) Validate() error {
	if err := i.validateInputBase("input.time"); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements json.Marshaler, ensuring the type field is set.
func (i InputTime) MarshalJSON() ([]byte, error) {
	ii := i
	if ii.Type == "" {
		ii.Type = m.TypeInputTime
	}
	type alias InputTime
	return json.Marshal(alias(ii))
}

// UnmarshalJSON implements json.Unmarshaler and validates the decoded value.
func (i *InputTime) UnmarshalJSON(b []byte) error {
	type alias InputTime
	var tmp alias
	if err := json.Unmarshal(b, &tmp); err != nil {
		return fmt.Errorf("input.time: decode: %w", err)
	}
	if tmp.Type != "" && tmp.Type != m.TypeInputTime {
		return fmt.Errorf("input.time.type must be %q (got %q)", m.TypeInputTime, tmp.Type)
	}
	if tmp.Type == "" {
		tmp.Type = m.TypeInputTime
	}
	val := InputTime(tmp)
	if err := val.Validate(); err != nil {
		return err
	}
	*i = val
	return nil
}

func init() {
	e.RegisterElement(
		m.TypeInputTime,
		func() e.Element { return &InputTime{Type: m.TypeInputTime} },
	)
}
