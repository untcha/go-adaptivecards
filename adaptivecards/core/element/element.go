package element

import (
	"fmt"
	"strings"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// Common interface for all card elements
type Element interface {
	GetType() m.TypeString
}

// ElementBase holds common fields shared by all elements and is embedded in concrete element types.
type ElementBase struct {
	Fallback    *ElementFallback     `json:"fallback,omitempty"`    // Version 1.2
	Height      m.BlockElementHeight `json:"height,omitempty"`      // Version 1.1
	Separator   bool                 `json:"separator,omitempty"`   // Version 1.2
	Spacing     m.Spacing            `json:"spacing,omitempty"`     // Version 1.2
	ID          string               `json:"id,omitempty"`          // Version 1.2
	IsVisible   *bool                `json:"isVisible,omitempty"`   // Version 1.2
	Requires    map[string]string    `json:"requires,omitempty"`    // Version 1.2
	TargetWidth m.TargetWidth        `json:"targetWidth,omitempty"` // Version 1.6 (host feature; not in JSON schema)
}

// Validate checks the shared base-element fields. typeName is the concrete
// element type (e.g. "TextBlock") and is used to prefix error messages so they
// remain specific to the element being validated.
//
// This centralizes what used to be duplicated per element. New shared-base
// fields should be validated here once, not in each element.
func (b ElementBase) Validate(typeName string) error {
	if b.Height != "" && !b.Height.IsValid() {
		return m.NewEnumError(typeName+".height", string(b.Height), m.AllowedBlockElementHeightStrings())
	}
	if b.Spacing != "" && !b.Spacing.IsValid() {
		return m.NewEnumError(typeName+".spacing", string(b.Spacing), m.AllowedSpacingStrings())
	}
	if b.TargetWidth != "" && !b.TargetWidth.IsValid() {
		return m.NewEnumError(typeName+".targetWidth", string(b.TargetWidth), m.AllowedTargetWidthStrings())
	}
	if b.ID != "" {
		id := strings.TrimSpace(b.ID)
		if id == "" {
			return fmt.Errorf("%s.id cannot be empty or whitespace-only", typeName)
		}
		// ID should not contain characters that could cause issues.
		if strings.ContainsAny(id, "\n\r\t") {
			return fmt.Errorf("%s.id cannot contain newlines or tabs", typeName)
		}
	}
	return nil
}
