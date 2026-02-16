package element

import "encoding/json"

type FallbackOption string

const FallbackDrop FallbackOption = "drop"

type ElementFallback struct {
	Option  *FallbackOption
	Content Element
}

func ElementFallbackDropOption() *ElementFallback {
	o := FallbackDrop
	return &ElementFallback{Option: &o}
}

func ElementFallbackContent(e Element) *ElementFallback { return &ElementFallback{Content: e} }

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
