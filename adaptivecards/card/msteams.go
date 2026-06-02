package card

import m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"

// MSTeamsWidth is the Teams-specific card width host extension value.
type MSTeamsWidth string

// MSTeamsWidthFull renders the card across the full message column in Teams.
// It is currently the only documented value.
const MSTeamsWidthFull MSTeamsWidth = "Full"

// IsValid reports whether the width is a documented value.
func (w MSTeamsWidth) IsValid() bool { return w == MSTeamsWidthFull }

// AllowedMSTeamsWidthStrings returns the documented msteams.width values.
func AllowedMSTeamsWidthStrings() []string { return []string{string(MSTeamsWidthFull)} }

// MSTeams holds Microsoft Teams host-specific extensions to an Adaptive Card.
//
// This object is NOT part of the Adaptive Cards schema; it is a Teams host
// extension honored only by Teams clients and ignored by other renderers.
// Microsoft recommends testing narrow form factors (mobile, meeting side
// panels) for truncation when using full-width cards.
//
// See: https://learn.microsoft.com/en-us/microsoftteams/platform/task-modules-and-cards/cards/cards-format#full-width-adaptive-card
type MSTeams struct {
	Width MSTeamsWidth `json:"width,omitempty"` // Teams host extension
	// Future: Entities []Mention `json:"entities,omitempty"`, AllowExpand bool, etc.
}

// Validate checks that the MSTeams extension carries documented values.
func (t MSTeams) Validate() error {
	if t.Width != "" && !t.Width.IsValid() {
		return m.NewEnumError("msteams.width", string(t.Width), AllowedMSTeamsWidthStrings())
	}
	return nil
}
