package actions

import m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"

// Action is the common interface implemented by all Adaptive Card actions.
type Action interface {
	GetType() m.TypeString
}

// ActionBase holds common fields shared by all actions and is embedded in concrete action types.
type ActionBase struct {
	Title     string            `json:"title,omitempty"`     // Version 1.0
	IconURL   m.URI             `json:"iconUrl,omitempty"`   // Version 1.1
	ID        string            `json:"id,omitempty"`        // Version 1.0
	Style     m.ActionStyle     `json:"style,omitempty"`     // Version 1.2
	Fallback  *ActionFallback   `json:"fallback,omitempty"`  // Version 1.2
	Tooltip   string            `json:"tooltip,omitempty"`   // Version 1.5
	IsEnabled *bool             `json:"isEnabled,omitempty"` // Version 1.5
	Mode      m.ActionMode      `json:"mode,omitempty"`      // Version 1.5
	Requires  map[string]string `json:"requires,omitempty"`  // Version 1.2
}
