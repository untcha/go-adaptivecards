package model

// ChoiceInputStyle is a style hint for Input.ChoiceSet.
type ChoiceInputStyle string

const (
	ChoiceInputStyleCompact  ChoiceInputStyle = "compact"
	ChoiceInputStyleExpanded ChoiceInputStyle = "expanded"
)

var choiceInputStyleAllowed = []ChoiceInputStyle{
	ChoiceInputStyleCompact,
	ChoiceInputStyleExpanded,
}

// IsValid reports whether s is a valid ChoiceInputStyle value.
func (s ChoiceInputStyle) IsValid() bool {
	return enumIsValid(s, choiceInputStyleAllowed)
}

// ParseChoiceInputStyle parses v into a ChoiceInputStyle or returns an EnumError.
func ParseChoiceInputStyle(v string) (ChoiceInputStyle, error) {
	return enumParse(v, choiceInputStyleAllowed, "ChoiceInputStyle")
}

// UnmarshalJSON implements json.Unmarshaler, enforcing allowed values.
func (s *ChoiceInputStyle) UnmarshalJSON(b []byte) error {
	return enumUnmarshalJSON(s, b, choiceInputStyleAllowed, "ChoiceInputStyle")
}

// MarshalJSON implements json.Marshaler, enforcing allowed values.
func (s ChoiceInputStyle) MarshalJSON() ([]byte, error) {
	return enumMarshalJSON(s, choiceInputStyleAllowed, "ChoiceInputStyle")
}

// TextInputStyle is a style hint for Input.Text.
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

// IsValid reports whether s is a valid TextInputStyle value.
func (s TextInputStyle) IsValid() bool {
	return enumIsValid(s, textInputStyleAllowed)
}

// ParseTextInputStyle parses v into a TextInputStyle or returns an EnumError.
func ParseTextInputStyle(v string) (TextInputStyle, error) {
	return enumParse(v, textInputStyleAllowed, "TextInputStyle")
}

// UnmarshalJSON implements json.Unmarshaler, enforcing allowed values.
func (s *TextInputStyle) UnmarshalJSON(b []byte) error {
	return enumUnmarshalJSON(s, b, textInputStyleAllowed, "TextInputStyle")
}

// MarshalJSON implements json.Marshaler, enforcing allowed values.
func (s TextInputStyle) MarshalJSON() ([]byte, error) {
	return enumMarshalJSON(s, textInputStyleAllowed, "TextInputStyle")
}

// AllowedTextInputStyleStrings returns the list of valid TextInputStyle values as strings.
func AllowedTextInputStyleStrings() []string {
	return enumAllowedStrings(textInputStyleAllowed)
}

// AllowedChoiceInputStyleStrings returns the list of valid ChoiceInputStyle values as strings.
func AllowedChoiceInputStyleStrings() []string {
	return enumAllowedStrings(choiceInputStyleAllowed)
}
