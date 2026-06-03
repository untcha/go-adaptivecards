package actions

import (
	"encoding/json"
	"fmt"
)

// FallbackOption is the string form of an action fallback. The only valid
// value is FallbackDrop.
type FallbackOption string

// FallbackDrop is the only valid FallbackOption value and instructs hosts to
// drop the action when it is unsupported.
const FallbackDrop FallbackOption = "drop"

// ActionFallback describes the fallback behavior for an action when the host
// does not support it. Option, when set, must be FallbackDrop. Option and
// Content are mutually exclusive: set Option to drop the action, or set Content
// to an alternate Action to use instead.
type ActionFallback struct {
	Option  *FallbackOption
	Content Action
}

// ActionFallbackDropOption returns an ActionFallback whose Option is set to
// FallbackDrop.
func ActionFallbackDropOption() *ActionFallback {
	o := FallbackDrop
	return &ActionFallback{Option: &o}
}

// ActionFallbackContent returns an ActionFallback whose Content is set to the
// given alternate Action.
func ActionFallbackContent(a Action) *ActionFallback { return &ActionFallback{Content: a} }

// MarshalJSON encodes the ActionFallback as either the "drop" string or the
// embedded action's JSON.
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

// UnmarshalJSON decodes the ActionFallback from either the "drop" string or an
// action object resolved through the action factory.
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
