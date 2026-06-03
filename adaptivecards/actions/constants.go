package actions

import m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"

// ActionStyleDefault and the other constants in this block are convenience
// aliases re-exported from the core model package. ActionStyle* values mirror
// the corresponding model.ActionStyle values, and ActionMode* values mirror the
// corresponding model.ActionMode values, so callers can reference them without
// importing the model package directly.
const (
	ActionStyleDefault     = m.ActionStyleDefault
	ActionStylePositive    = m.ActionStylePositive
	ActionStyleDestructive = m.ActionStyleDestructive
	ActionModePrimary      = m.ActionModePrimary
	ActionModeSecondary    = m.ActionModeSecondary
)
