package inputs

import (
	"encoding/json"
	"fmt"

	e "github.com/untcha/go-adaptivecards/adaptivecards/core/element"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// Input.Date
// Lets a user choose a date.
// See: https://adaptivecards.io/explorer/Input.Date.html
type InputDate struct {
	InputBase
	Type        m.TypeString `json:"type"`                  // Version 1.0
	Max         string       `json:"max,omitempty"`         // Version 1.0
	Min         string       `json:"min,omitempty"`         // Version 1.0
	Placeholder string       `json:"placeholder,omitempty"` // Version 1.0
	Value       string       `json:"value,omitempty"`       // Version 1.0
}

func NewInputDate(id string) InputDate {
	return InputDate{
		InputBase: InputBase{ElementBase: e.ElementBase{ID: id}},
		Type:      m.TypeInputDate,
	}
}

func (i InputDate) GetType() m.TypeString { return m.TypeInputDate }

func (i InputDate) Validate() error {
	if err := i.validateInputBase("input.date"); err != nil {
		return err
	}
	return nil
}

func (i InputDate) MarshalJSON() ([]byte, error) {
	ii := i
	if ii.Type == "" {
		ii.Type = m.TypeInputDate
	}
	type alias InputDate
	return json.Marshal(alias(ii))
}

func (i *InputDate) UnmarshalJSON(b []byte) error {
	type alias InputDate
	var tmp alias
	if err := json.Unmarshal(b, &tmp); err != nil {
		return fmt.Errorf("input.date: decode: %w", err)
	}
	if tmp.Type != "" && tmp.Type != m.TypeInputDate {
		return fmt.Errorf("input.date.type must be %q (got %q)", m.TypeInputDate, tmp.Type)
	}
	if tmp.Type == "" {
		tmp.Type = m.TypeInputDate
	}
	val := InputDate(tmp)
	if err := val.Validate(); err != nil {
		return err
	}
	*i = val
	return nil
}

func init() {
	e.RegisterElement(
		m.TypeInputDate,
		func() e.Element { return &InputDate{Type: m.TypeInputDate} },
	)
}
