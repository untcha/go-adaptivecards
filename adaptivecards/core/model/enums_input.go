package model

// Style hint for Input.ChoiceSet.
type ChoiceInputStyle string

const (
	ChoiceInputStyleCompact  ChoiceInputStyle = "compact"
	ChoiceInputStyleExpanded ChoiceInputStyle = "expanded"
)

var choiceInputStyleAllowed = []ChoiceInputStyle{
	ChoiceInputStyleCompact,
	ChoiceInputStyleExpanded,
}

func (s ChoiceInputStyle) IsValid() bool {
	return enumIsValid(s, choiceInputStyleAllowed)
}

func ParseChoiceInputStyle(v string) (ChoiceInputStyle, error) {
	return enumParse(v, choiceInputStyleAllowed, "ChoiceInputStyle")
}

func (s *ChoiceInputStyle) UnmarshalJSON(b []byte) error {
	return enumUnmarshalJSON(s, b, choiceInputStyleAllowed, "ChoiceInputStyle")
}

func (s ChoiceInputStyle) MarshalJSON() ([]byte, error) {
	return enumMarshalJSON(s, choiceInputStyleAllowed, "ChoiceInputStyle")
}

// Style hint for Input.Text.
type TextInputStyle string

const (
	TextInputStyleText     TextInputStyle = "text"
	TextInputStyleTel      TextInputStyle = "tel"
	TextInputStyleURL      TextInputStyle = "url"
	TextInputStyleEmail    TextInputStyle = "email"
	TextInputStylePassword TextInputStyle = "password"
)

var textInputStyleAllowed = []TextInputStyle{
	TextInputStyleText,
	TextInputStyleTel,
	TextInputStyleURL,
	TextInputStyleEmail,
	TextInputStylePassword,
}

func (s TextInputStyle) IsValid() bool {
	return enumIsValid(s, textInputStyleAllowed)
}

func ParseTextInputStyle(v string) (TextInputStyle, error) {
	return enumParse(v, textInputStyleAllowed, "TextInputStyle")
}

func (s *TextInputStyle) UnmarshalJSON(b []byte) error {
	return enumUnmarshalJSON(s, b, textInputStyleAllowed, "TextInputStyle")
}

func (s TextInputStyle) MarshalJSON() ([]byte, error) {
	return enumMarshalJSON(s, textInputStyleAllowed, "TextInputStyle")
}

func AllowedTextInputStyleStrings() []string {
	return enumAllowedStrings(textInputStyleAllowed)
}

func AllowedChoiceInputStyleStrings() []string {
	return enumAllowedStrings(choiceInputStyleAllowed)
}
