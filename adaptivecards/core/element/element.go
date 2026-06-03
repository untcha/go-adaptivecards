package element

import m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"

// Element is the common interface implemented by all Adaptive Card elements.
type Element interface {
	GetType() m.TypeString
}

// ElementBase holds common fields shared by all elements and is embedded in concrete element types.
type ElementBase struct {
	Fallback  *ElementFallback     `json:"fallback,omitempty"`  // Version 1.2
	Height    m.BlockElementHeight `json:"height,omitempty"`    // Version 1.1
	Separator bool                 `json:"separator,omitempty"` // Version 1.2
	Spacing   m.Spacing            `json:"spacing,omitempty"`   // Version 1.2
	ID        string               `json:"id,omitempty"`        // Version 1.2
	IsVisible *bool                `json:"isVisible,omitempty"` // Version 1.2
	Requires  map[string]string    `json:"requires,omitempty"`  // Version 1.2
}
