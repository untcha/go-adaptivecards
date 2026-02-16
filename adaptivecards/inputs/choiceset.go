package inputs

import (
	"encoding/json"
	"fmt"

	e "github.com/untcha/go-adaptivecards/adaptivecards/core/element"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// Input.ChoiceSet
// Allows a user to input a Choice.
// See: https://adaptivecards.io/explorer/Input.ChoiceSet.html
type InputChoiceSet struct {
	InputBase
	Type          m.TypeString       `json:"type"`                    // Version 1.0
	Choices       []InputChoice      `json:"choices,omitempty"`       // Version 1.0
	IsMultiSelect bool               `json:"isMultiSelect,omitempty"` // Version 1.0
	Style         m.ChoiceInputStyle `json:"style,omitempty"`         // Version 1.0
	Value         string             `json:"value,omitempty"`         // Version 1.0
	Placeholder   string             `json:"placeholder,omitempty"`   // Version 1.0
	Wrap          bool               `json:"wrap,omitempty"`          // Version 1.2
}

func (i InputChoiceSet) GetType() m.TypeString { return m.TypeInputChoiceSet }

// NewInputChoiceSet creates an Input.ChoiceSet with the specified id and choices.
func NewInputChoiceSet(id string, choices ...InputChoice) InputChoiceSet {
	return InputChoiceSet{
		InputBase: InputBase{ElementBase: e.ElementBase{ID: id}},
		Type:      m.TypeInputChoiceSet,
		Choices:   choices,
	}
}

// Validate checks that the InputChoiceSet has valid fields.
func (i InputChoiceSet) Validate() error {
	if err := i.validateInputBase("input.choiceSet"); err != nil {
		return err
	}
	for idx, c := range i.Choices {
		if err := c.Validate(); err != nil {
			return fmt.Errorf("input.choiceSet.choices[%d]: %w", idx, err)
		}
	}
	if i.Style != "" && !i.Style.IsValid() {
		return m.NewEnumError(
			"InputChoiceSet.style",
			string(i.Style),
			m.AllowedChoiceInputStyleStrings(),
		)
	}
	return nil
}

func (i InputChoiceSet) MarshalJSON() ([]byte, error) {
	ii := i
	if ii.Type == "" {
		ii.Type = m.TypeInputChoiceSet
	}
	type alias InputChoiceSet
	return json.Marshal(alias(ii))
}

func (i *InputChoiceSet) UnmarshalJSON(b []byte) error {
	type alias InputChoiceSet
	var tmp alias
	if err := json.Unmarshal(b, &tmp); err != nil {
		return fmt.Errorf("input.choiceSet: decode: %w", err)
	}
	if tmp.Type != "" && tmp.Type != m.TypeInputChoiceSet {
		return fmt.Errorf(
			"input.choiceSet.type must be %q (got %q)",
			m.TypeInputChoiceSet,
			tmp.Type,
		)
	}
	if tmp.Type == "" {
		tmp.Type = m.TypeInputChoiceSet
	}
	val := InputChoiceSet(tmp)
	if err := val.Validate(); err != nil {
		return err
	}
	*i = val
	return nil
}

func init() {
	e.RegisterElement(
		m.TypeInputChoiceSet,
		func() e.Element { return &InputChoiceSet{Type: m.TypeInputChoiceSet} },
	)
}
