package card

import (
	"errors"
	"fmt"

	a "github.com/untcha/go-adaptivecards/adaptivecards/actions"
	t "github.com/untcha/go-adaptivecards/adaptivecards/containers"
	e "github.com/untcha/go-adaptivecards/adaptivecards/core/element"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
	els "github.com/untcha/go-adaptivecards/adaptivecards/elements"
	"github.com/untcha/go-adaptivecards/adaptivecards/schema"
)

// Validate checks the Card in three stages: first any accumulated builder
// errors, then logical per-element validation, and finally JSON-schema
// validation against the spec (with the Teams-only msteams extension stripped
// from a shallow copy so it does not violate additionalProperties:false).
func (c *Card) Validate() error {
	if c == nil {
		return errors.New("card is nil")
	}
	// 1) Builder errors first
	if c.CardBuildErr() != nil {
		return c.CardBuildErr()
	}

	// 2) Logical per-element validation
	if err := c.validateLogical(); err != nil {
		return err
	}

	// 3) Full JSON Schema validation (strict, spec-accurate).
	// msteams is a Teams host extension, not part of the Adaptive Cards
	// schema (which sets additionalProperties:false on AdaptiveCard).
	// Strip it from a shallow copy so schema validation only covers the
	// spec-defined fields; the extension is validated logically above.
	toValidate := c
	if c.MSTeams != nil {
		cc := *c
		cc.MSTeams = nil
		toValidate = &cc
	}
	if err := schema.Validate(toValidate); err != nil {
		return err
	}
	return nil
}

func (c *Card) validateLogical() error {
	// Validate Body elements
	if err := validateElements(c.Body); err != nil {
		return fmt.Errorf("body: %w", err)
	}
	// Validate Actions
	if err := validateActions(c.Actions); err != nil {
		return fmt.Errorf("actions: %w", err)
	}
	// Validate SelectAction (Action.ShowCard is not supported)
	if c.SelectAction != nil {
		if err := validateSelectAction(c.SelectAction); err != nil {
			return fmt.Errorf("selectAction: %w", err)
		}
	}
	// Validate MSTeams host extension (not part of the AC schema)
	if c.MSTeams != nil {
		if err := c.MSTeams.Validate(); err != nil {
			return fmt.Errorf("msteams: %w", err)
		}
	}
	// Validate BackgroundImage
	if c.BackgroundImage != nil {
		if c.BackgroundImage.Object() != nil {
			if _, err := m.BackgroundImageObject(*c.BackgroundImage.Object()); err != nil {
				return fmt.Errorf("backgroundImage: %w", err)
			}
		} else if c.BackgroundImage.URL() != "" {
			if _, err := m.BackgroundImageURL(c.BackgroundImage.URL()); err != nil {
				return fmt.Errorf("backgroundImage: %w", err)
			}
		}
	}

	return nil
}

func validateElements(elements []e.Element) error {
	for i, el := range elements {
		if err := validateElement(el); err != nil {
			return fmt.Errorf("element[%d]: %w", i, err)
		}
	}
	return nil
}

// validateSelectAction ensures Action.ShowCard is not used in selectAction
func validateSelectAction(act a.Action) error {
	if act.GetType() == m.TypeActionShowCard {
		return fmt.Errorf("Action.ShowCard is not supported in selectAction")
	}
	// Validate the action itself
	return validateAction(act)
}

func validateElement(el e.Element) error {
	switch v := el.(type) {
	case els.TextBlock:
		return v.Validate()
	case *els.TextBlock:
		return v.Validate()
	case els.Image:
		if err := v.Validate(); err != nil {
			return err
		}
		// Validate selectAction if present
		if v.SelectAction != nil {
			if err := validateSelectAction(v.SelectAction); err != nil {
				return fmt.Errorf("selectAction: %w", err)
			}
		}
		return nil
	case *els.Image:
		if err := v.Validate(); err != nil {
			return err
		}
		// Validate selectAction if present
		if v.SelectAction != nil {
			if err := validateSelectAction(v.SelectAction); err != nil {
				return fmt.Errorf("selectAction: %w", err)
			}
		}
		return nil
	case t.Table:
		return v.Validate()
	case *t.Table:
		return v.Validate()
	default:
		// Prepare for more element types: if it implements Validate, call it.
		type validator interface{ Validate() error }
		if vv, ok := el.(validator); ok {
			return vv.Validate()
		}
		return nil
	}
}

func validateActions(actions []a.Action) error {
	for i, act := range actions {
		if err := validateAction(act); err != nil {
			return fmt.Errorf("action[%d]: %w", i, err)
		}
	}
	return nil
}

func validateAction(act a.Action) error {
	switch v := act.(type) {
	case a.ActionOpenURL:
		return v.Validate()
	case *a.ActionOpenURL:
		return v.Validate()
	default:
		// Prepare for more action types: if it implements Validate, call it.
		type validator interface{ Validate() error }
		if vv, ok := act.(validator); ok {
			return vv.Validate()
		}
	}
	return nil
}
