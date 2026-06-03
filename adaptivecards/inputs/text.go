package inputs

import (
	"encoding/json"
	"fmt"
	"math"

	a "github.com/untcha/go-adaptivecards/adaptivecards/actions"
	e "github.com/untcha/go-adaptivecards/adaptivecards/core/element"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// InputText represents an Input.Text element that lets a user enter text.
// See: https://adaptivecards.io/explorer/Input.Text.html
type InputText struct {
	InputBase
	Type         m.TypeString     `json:"type"`                   // Version 1.0
	IsMultiline  bool             `json:"isMultiline,omitempty"`  // Version 1.0
	MaxLength    *float64         `json:"maxLength,omitempty"`    // Version 1.0
	Placeholder  string           `json:"placeholder,omitempty"`  // Version 1.0
	Regex        string           `json:"regex,omitempty"`        // Version 1.3
	Style        m.TextInputStyle `json:"style,omitempty"`        // Version 1.0
	InlineAction a.Action         `json:"inlineAction,omitempty"` // Version 1.2
	Value        string           `json:"value,omitempty"`        // Version 1.0
}

// NewInputText creates an Input.Text with the specified id.
func NewInputText(id string) InputText {
	return InputText{
		InputBase: InputBase{
			ElementBase: e.ElementBase{ID: id},
		},
		Type: m.TypeInputText,
	}
}

// GetType returns the Adaptive Card type discriminator for Input.Text.
func (i InputText) GetType() m.TypeString { return m.TypeInputText }

// Validate checks that the InputText has valid fields.
func (i InputText) Validate() error {
	if err := i.validateInputBase("input.text"); err != nil {
		return err
	}
	if i.MaxLength != nil {
		if math.IsNaN(*i.MaxLength) || math.IsInf(*i.MaxLength, 0) {
			return fmt.Errorf("input.text.maxLength must be a finite number")
		}
	}
	if i.Style != "" && !i.Style.IsValid() {
		return m.NewEnumError("InputText.style", string(i.Style), m.AllowedTextInputStyleStrings())
	}
	if i.InlineAction != nil {
		if err := validateSelectAction(i.InlineAction); err != nil {
			return fmt.Errorf("input.text.inlineAction: %w", err)
		}
	}
	return nil
}

// MarshalJSON implements json.Marshaler, ensuring the type field is set.
func (i InputText) MarshalJSON() ([]byte, error) {
	ii := i
	if ii.Type == "" {
		ii.Type = m.TypeInputText
	}
	type alias InputText
	return json.Marshal(alias(ii))
}

// UnmarshalJSON implements json.Unmarshaler and validates the decoded value.
func (i *InputText) UnmarshalJSON(b []byte) error {
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(b, &obj); err != nil {
		return fmt.Errorf("input.text: decode: %w", err)
	}
	if v, ok := obj["type"]; ok {
		var t string
		if err := json.Unmarshal(v, &t); err != nil {
			return fmt.Errorf("input.text.type: %w", err)
		}
		if t != "" && t != string(m.TypeInputText) {
			return fmt.Errorf("input.text.type must be %q (got %q)", m.TypeInputText, t)
		}
	}
	inlineRaw := obj["inlineAction"]
	delete(obj, "inlineAction")

	type alias InputText
	var base alias
	baseBytes, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("input.text: re-encode base: %w", err)
	}
	if err := json.Unmarshal(baseBytes, &base); err != nil {
		return fmt.Errorf("input.text: decode base: %w", err)
	}
	if base.Type == "" {
		base.Type = m.TypeInputText
	}
	if len(inlineRaw) != 0 {
		act, err := a.UnmarshalAction(inlineRaw)
		if err != nil {
			return fmt.Errorf("input.text.inlineAction: %w", err)
		}
		base.InlineAction = act
	}
	tmp := InputText(base)
	if err := tmp.Validate(); err != nil {
		return err
	}
	*i = tmp
	return nil
}

func init() {
	e.RegisterElement(
		m.TypeInputText,
		func() e.Element { return &InputText{Type: m.TypeInputText} },
	)
}
