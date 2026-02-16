package actions

import (
	"encoding/json"
	"fmt"
)

type FallbackOption string

const FallbackDrop FallbackOption = "drop"

type ActionFallback struct {
	Option  *FallbackOption
	Content Action
}

func ActionFallbackDropOption() *ActionFallback {
	o := FallbackDrop
	return &ActionFallback{Option: &o}
}

func ActionFallbackContent(a Action) *ActionFallback { return &ActionFallback{Content: a} }

func (f ActionFallback) MarshalJSON() ([]byte, error) {
	if f.Option != nil {
		if *f.Option != FallbackDrop {
			return nil, fmt.Errorf(
				"action.fallback option must be %q (got %q)",
				FallbackDrop,
				*f.Option,
			)
		}
		return json.Marshal(string(*f.Option)) // "drop"
	}
	if f.Content != nil {
		return json.Marshal(f.Content) // Action JSON
	}
	// treat nil as JSON null (omit via omitempty in parent)
	return []byte("null"), nil
}

func (f *ActionFallback) UnmarshalJSON(b []byte) error {
	// try string option
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		o := FallbackOption(s)
		if o != FallbackDrop {
			return fmt.Errorf("action.fallback option must be %q (got %q)", FallbackDrop, o)
		}
		*f = ActionFallback{Option: &o}
		return nil
	}
	// else assume action object; delegate to your action factory
	ac, err := UnmarshalAction(b) // implement: inspect "type" and decode
	if err != nil {
		return err
	}
	*f = ActionFallback{Content: ac}
	return nil
}
