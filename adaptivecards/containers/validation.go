package containers

import (
	"fmt"

	a "github.com/untcha/go-adaptivecards/adaptivecards/actions"
	e "github.com/untcha/go-adaptivecards/adaptivecards/core/element"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

func validateElements(items []e.Element) error {
	for i, item := range items {
		if err := validateElement(item); err != nil {
			return fmt.Errorf("element[%d]: %w", i, err)
		}
	}
	return nil
}

func validateElement(item e.Element) error {
	type validator interface{ Validate() error }
	if v, ok := item.(validator); ok {
		return v.Validate()
	}
	return nil
}

func validateAction(act a.Action) error {
	type validator interface{ Validate() error }
	if v, ok := act.(validator); ok {
		return v.Validate()
	}
	return nil
}

func validateActions(actions []a.Action) error {
	for i, act := range actions {
		if err := validateAction(act); err != nil {
			return fmt.Errorf("action[%d]: %w", i, err)
		}
	}
	return nil
}

func validateSelectAction(act a.Action) error {
	if act.GetType() == m.TypeActionShowCard {
		return fmt.Errorf("Action.ShowCard is not supported in selectAction")
	}
	return validateAction(act)
}
