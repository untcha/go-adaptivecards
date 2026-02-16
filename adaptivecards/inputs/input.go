package inputs

import (
	"fmt"

	e "github.com/untcha/go-adaptivecards/adaptivecards/core/element"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// InputBase holds common properties shared by input elements.
type InputBase struct {
	e.ElementBase        // Embedding e.ElementBase to include common element fields
	ErrorMessage  string `json:"errorMessage,omitempty"` // Version 1.3
	IsRequired    *bool  `json:"isRequired,omitempty"`   // Version 1.3
	Label         string `json:"label,omitempty"`        // Version 1.3
}

func (b InputBase) validateInputBase(prefix string) error {
	if b.ID == "" {
		return fmt.Errorf("%s.id is required", prefix)
	}
	if b.Height != "" && !b.Height.IsValid() {
		return m.NewEnumError(
			prefix+".height",
			string(b.Height),
			m.AllowedBlockElementHeightStrings(),
		)
	}
	if b.Spacing != "" && !b.Spacing.IsValid() {
		return m.NewEnumError(
			prefix+".spacing",
			string(b.Spacing),
			m.AllowedSpacingStrings(),
		)
	}
	return nil
}
