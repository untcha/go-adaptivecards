package elements

import (
	"encoding/json"
	"fmt"

	a "github.com/untcha/go-adaptivecards/adaptivecards/actions"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// TextRun defines a single run of formatted text.
// A TextRun with no properties set can be represented in the json as string
// containing the text as a shorthand for the json object. These two representations are equivalent.
// See: https://adaptivecards.io/explorer/TextRun.html
type TextRun struct {
	Type          m.TypeString `json:"type,omitempty"`          // Version 1.2
	Text          string       `json:"text"`                    // Version 1.2
	Color         m.TextColor  `json:"color,omitempty"`         // Version 1.2
	FontType      m.FontType   `json:"fontType,omitempty"`      // Version 1.2
	Highlight     bool         `json:"highlight,omitempty"`     // Version 1.2
	IsSubtle      *bool        `json:"isSubtle,omitempty"`      // Version 1.2
	Italic        bool         `json:"italic,omitempty"`        // Version 1.2
	SelectAction  a.Action     `json:"selectAction,omitempty"`  // Version 1.2
	Size          m.FontSize   `json:"size,omitempty"`          // Version 1.2
	Strikethrough bool         `json:"strikethrough,omitempty"` // Version 1.2
	Underline     bool         `json:"underline,omitempty"`     // Version 1.3
	Weight        m.FontWeight `json:"weight,omitempty"`        // Version 1.2
}

// NewTextRun creates a TextRun with the specified text.
func NewTextRun(text string) TextRun {
	return TextRun{
		Type: m.TypeTextRun,
		Text: text,
	}
}

// GetType returns the element type of the TextRun.
func (r TextRun) GetType() m.TypeString { return m.TypeTextRun }

// Validate checks that the TextRun has valid fields.
func (r TextRun) Validate() error {
	if r.Text == "" {
		return fmt.Errorf("textRun.text is required")
	}
	if r.Color != "" && !r.Color.IsValid() {
		return m.NewEnumError("TextRun.color", string(r.Color), m.AllowedTextColorStrings())
	}
	if r.FontType != "" && !r.FontType.IsValid() {
		return m.NewEnumError("TextRun.fontType", string(r.FontType), m.AllowedFontTypeStrings())
	}
	if r.Size != "" && !r.Size.IsValid() {
		return m.NewEnumError("TextRun.size", string(r.Size), m.AllowedFontSizeStrings())
	}
	if r.Weight != "" && !r.Weight.IsValid() {
		return m.NewEnumError("TextRun.weight", string(r.Weight), m.AllowedFontWeightStrings())
	}
	if r.SelectAction != nil {
		if err := validateSelectAction(r.SelectAction); err != nil {
			return fmt.Errorf("textRun.selectAction: %w", err)
		}
	}
	return nil
}

// MarshalJSON encodes the TextRun, emitting a plain string when only Text is set.
func (r TextRun) MarshalJSON() ([]byte, error) {
	rr := r
	if rr.onlyText() {
		return json.Marshal(rr.Text)
	}
	if rr.Type == "" {
		rr.Type = m.TypeTextRun
	}
	type alias TextRun
	return json.Marshal(alias(rr))
}

// UnmarshalJSON decodes a TextRun from either a JSON string shorthand or a JSON object.
func (r *TextRun) UnmarshalJSON(b []byte) error {
	var textOnly string
	if err := json.Unmarshal(b, &textOnly); err == nil {
		tmp := TextRun{
			Type: m.TypeTextRun,
			Text: textOnly,
		}
		if err := tmp.Validate(); err != nil {
			return err
		}
		*r = tmp
		return nil
	}

	var obj map[string]json.RawMessage
	if err := json.Unmarshal(b, &obj); err != nil {
		return fmt.Errorf("textRun: decode: %w", err)
	}
	if v, ok := obj["type"]; ok {
		var t string
		if err := json.Unmarshal(v, &t); err != nil {
			return fmt.Errorf("textRun.type: %w", err)
		}
		if t != "" && t != string(m.TypeTextRun) {
			return fmt.Errorf("textRun.type must be %q (got %q)", m.TypeTextRun, t)
		}
	}
	selectActionRaw := obj["selectAction"]
	delete(obj, "selectAction")

	type alias TextRun
	var base alias
	baseBytes, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("textRun: re-encode base: %w", err)
	}
	if err := json.Unmarshal(baseBytes, &base); err != nil {
		return fmt.Errorf("textRun: decode base: %w", err)
	}
	if base.Type == "" {
		base.Type = m.TypeTextRun
	}
	if len(selectActionRaw) != 0 {
		act, err := a.UnmarshalAction(selectActionRaw)
		if err != nil {
			return fmt.Errorf("textRun.selectAction: %w", err)
		}
		base.SelectAction = act
	}
	tmp := TextRun(base)
	if err := tmp.Validate(); err != nil {
		return err
	}
	*r = tmp
	return nil
}

func (r TextRun) onlyText() bool {
	return r.Type == "" &&
		r.Text != "" &&
		r.Color == "" &&
		r.FontType == "" &&
		!r.Highlight &&
		r.IsSubtle == nil &&
		!r.Italic &&
		r.SelectAction == nil &&
		r.Size == "" &&
		!r.Strikethrough &&
		!r.Underline &&
		r.Weight == ""
}
