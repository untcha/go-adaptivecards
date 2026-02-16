package inputs

import (
	"fmt"

	a "github.com/untcha/go-adaptivecards/adaptivecards/actions"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

func validateAction(act a.Action) error {
	type validator interface{ Validate() error }
	if v, ok := act.(validator); ok {
		return v.Validate()
	}
	return nil
}

func validateSelectAction(act a.Action) error {
	if act.GetType() == m.TypeActionShowCard {
		return fmt.Errorf("Action.ShowCard is not supported in selectAction")
	}
	return validateAction(act)
}
