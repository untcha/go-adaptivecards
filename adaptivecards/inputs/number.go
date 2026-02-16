package inputs

import (
	"encoding/json"
	"fmt"

	e "github.com/untcha/go-adaptivecards/adaptivecards/core/element"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// Input.Number
// Allows a user to enter a number.
// See: https://adaptivecards.io/explorer/Input.Number.html
type InputNumber struct {
	InputBase
	Type        m.TypeString `json:"type"`                  // Version 1.0
	Max         *float64     `json:"max,omitempty"`         // Version 1.0
	Min         *float64     `json:"min,omitempty"`         // Version 1.0
	Placeholder string       `json:"placeholder,omitempty"` // Version 1.0
	Value       *float64     `json:"value,omitempty"`       // Version 1.0
}

func (i InputNumber) GetType() m.TypeString { return m.TypeInputNumber }

func NewInputNumber(id string) InputNumber {
	return InputNumber{
		InputBase: InputBase{ElementBase: e.ElementBase{ID: id}},
		Type:      m.TypeInputNumber,
	}
}

func (i InputNumber) Validate() error {
	if err := i.validateInputBase("input.number"); err != nil {
		return err
	}
	return nil
}

func (i InputNumber) MarshalJSON() ([]byte, error) {
	ii := i
	if ii.Type == "" {
		ii.Type = m.TypeInputNumber
	}
	type alias InputNumber
	return json.Marshal(alias(ii))
}

func (i *InputNumber) UnmarshalJSON(b []byte) error {
	type alias InputNumber
	var tmp alias
	if err := json.Unmarshal(b, &tmp); err != nil {
		return fmt.Errorf("input.number: decode: %w", err)
	}
	if tmp.Type != "" && tmp.Type != m.TypeInputNumber {
		return fmt.Errorf("input.number.type must be %q (got %q)", m.TypeInputNumber, tmp.Type)
	}
	if tmp.Type == "" {
		tmp.Type = m.TypeInputNumber
	}
	val := InputNumber(tmp)
	if err := val.Validate(); err != nil {
		return err
	}
	*i = val
	return nil
}

func init() {
	e.RegisterElement(
		m.TypeInputNumber,
		func() e.Element { return &InputNumber{Type: m.TypeInputNumber} },
	)
}
