package element

import "encoding/json"

// FallbackOption is the string form of an element fallback, such as "drop".
type FallbackOption string

// FallbackDrop is the "drop" fallback option, instructing the host to drop the
// element when it cannot be rendered.
const FallbackDrop FallbackOption = "drop"

// ElementFallback models the element-level "fallback" property, which is either
// a string option ("drop") or a fallback Element. Option and Content are
// mutually exclusive.
type ElementFallback struct {
	Option  *FallbackOption
	Content Element
}

// ElementFallbackDropOption returns an ElementFallback representing the "drop" option.
func ElementFallbackDropOption() *ElementFallback {
	o := FallbackDrop
	return &ElementFallback{Option: &o}
}

// ElementFallbackContent returns an ElementFallback wrapping a fallback element.
func ElementFallbackContent(e Element) *ElementFallback { return &ElementFallback{Content: e} }

// MarshalJSON encodes the fallback either as a string option (e.g. "drop") or as an element object.
func (f ElementFallback) MarshalJSON() ([]byte, error) {
	if f.Option != nil {
		return json.Marshal(string(*f.Option)) // "drop"
	}
	if f.Content != nil {
		return json.Marshal(f.Content) // Element JSON
	}
	// treat nil as JSON null (omit via omitempty in parent)
	return []byte("null"), nil
}

// UnmarshalJSON decodes the fallback from either a string option (e.g. "drop") or an element object.
func (f *ElementFallback) UnmarshalJSON(b []byte) error {
	// try string option
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		o := FallbackOption(s)
		*f = ElementFallback{Option: &o}
		return nil
	}
	// else assume element object; delegate to your element factory
	el, err := UnmarshalElement(b) // implement: inspect "type" and decode
	if err != nil {
		return err
	}
	*f = ElementFallback{Content: el}
	return nil
}
