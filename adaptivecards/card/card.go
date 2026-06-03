package card

import (
	"encoding/json"
	"errors"
	"fmt"

	a "github.com/untcha/go-adaptivecards/adaptivecards/actions"
	e "github.com/untcha/go-adaptivecards/adaptivecards/core/element"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// Version is a card schema version string (e.g. "1.5").
type Version string

// Uri is a URI string as serialized in card JSON.
type Uri string

// TODO: Forward declarations for complex types not yet implemented

// Refresh is a placeholder for the AdaptiveCard refresh property (not yet implemented).
type Refresh any

// Authentication is a placeholder for the AdaptiveCard authentication property (not yet implemented).
type Authentication any

// Metadata is a placeholder for the AdaptiveCard metadata property (not yet implemented).
type Metadata any

// Schema constants describing the Adaptive Cards version this library targets.
const (
	// Version15 is the Adaptive Cards schema version targeted by this library.
	Version15 Version = "1.5"
	// Schema15 is the $schema URI emitted on generated cards.
	Schema15 Uri = "http://adaptivecards.io/schemas/adaptive-card.json"
)

// Card represents the root AdaptiveCard object.
// See: https://adaptivecards.io/explorer/AdaptiveCard.html
type Card struct {
	Type                     m.TypeString               `json:"type"`                               // Version 1.0
	Version                  Version                    `json:"version"`                            // Version 1.0
	Refresh                  any                        `json:"refresh,omitempty"`                  // Version 1.4 - not yet implemented
	Authentication           any                        `json:"authentication,omitempty"`           // Version 1.4 - not yet implemented
	Body                     []e.Element                `json:"body,omitempty"`                     // Version 1.0
	Actions                  []a.Action                 `json:"actions,omitempty"`                  // Version 1.0
	SelectAction             a.Action                   `json:"selectAction,omitempty"`             // Version 1.1
	FallbackText             string                     `json:"fallbackText,omitempty"`             // Version 1.0
	BackgroundImage          *m.BackgroundImageValue    `json:"backgroundImage,omitempty"`          // Version 1.2, 1.0
	MetaData                 any                        `json:"metadata,omitempty"`                 // Version 1.6 - not yet implemented
	MinHeight                string                     `json:"minHeight,omitempty"`                // Version 1.2
	Rtl                      *bool                      `json:"rtl,omitempty"`                      // Version 1.5 - supports null
	MSTeams                  *MSTeams                   `json:"msteams,omitempty"`                  // Teams host extension; not in AC schema
	Speak                    string                     `json:"speak,omitempty"`                    // Version 1.0
	Lang                     string                     `json:"lang,omitempty"`                     // Version 1.0
	VerticalContentAlignment m.VerticalContentAlignment `json:"verticalContentAlignment,omitempty"` // Version 1.1
	Schema                   Uri                        `json:"$schema,omitempty"`                  // Version 1.0
	buildErr                 error                      `json:"-"`                                  // internal only
}

// NewCard constructs a Card with default values.
// - Type is "AdaptiveCard"
// - Version is "1.5"
// - Body and Actions are initialized as empty slices
// - Schema is set to the 1.5 schema URI
func NewCard() *Card {
	return &Card{
		Type:    m.TypeAdaptiveCard,
		Version: Version15,
		Body:    make([]e.Element, 0),
		Actions: make([]a.Action, 0),
		Schema:  Schema15,
	}
}

// CardBuildErr returns any error encountered during the building process.
func (c *Card) CardBuildErr() error {
	if c == nil {
		return errors.New("card is nil")
	}
	return c.buildErr
}

// Build finalizes the Card building process and returns the Card and any error encountered.
// If there was an error during building, the returned Card may be incomplete or invalid.
func (c *Card) Build() (*Card, error) {
	if c == nil {
		return nil, errors.New("card is nil")
	}
	return c, c.buildErr
}

// MarshalJSON serializes the Card, applying default Type, Version, and Schema
// values and using an alias type to avoid infinite recursion.
func (c *Card) MarshalJSON() ([]byte, error) {
	if c == nil {
		return []byte("null"), nil
	}

	cc := *c // Copy (shallow) to avoid modifying the original
	// Set default values if not already set
	if cc.Type == "" {
		cc.Type = m.TypeAdaptiveCard
	}
	if cc.Version == "" {
		cc.Version = Version15
	}
	if cc.Schema == "" {
		cc.Schema = Schema15
	}

	// Ensure Body elements are properly marshalled
	// Create an alias to avoid infinite recursion in MarshalJSON
	// This is a common pattern in Go for custom JSON marshalling.
	type alias Card
	return json.Marshal(struct {
		*alias
	}{
		alias: (*alias)(&cc), // Embed the alias to inherit all fields; before (*alias)(c)
	})
}

// UnmarshalJSON decodes a Card from JSON, resolving the interface-typed fields
// (body, actions, selectAction) through their respective factories before
// decoding the remaining fields.
func (c *Card) UnmarshalJSON(b []byte) error {
	// Decode into a generic map to handle interface fields explicitly.
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(b, &obj); err != nil {
		return fmt.Errorf("Card: decode: %w", err)
	}

	// Validate/normalize type
	if v, ok := obj["type"]; ok {
		var t string
		if err := json.Unmarshal(v, &t); err != nil {
			return fmt.Errorf("card.type: %w", err)
		}
		if t != "" && t != string(m.TypeAdaptiveCard) {
			return fmt.Errorf("AdaptiveCard.type must be %q (got %q)", m.TypeAdaptiveCard, t)
		}
	}

	// Extract raw arrays/objects for interface fields and remove them before base decode
	bodyRaw := obj["body"]
	actionsRaw := obj["actions"]
	selectActionRaw := obj["selectAction"]
	delete(obj, "body")
	delete(obj, "actions")
	delete(obj, "selectAction")

	// Decode the remaining fields into base
	type alias Card
	var base alias
	baseBytes, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("Card: re-encode base: %w", err)
	}
	if err := json.Unmarshal(baseBytes, &base); err != nil {
		return fmt.Errorf("Card: decode base: %w", err)
	}

	// Defaults
	if base.Type == "" {
		base.Type = m.TypeAdaptiveCard
	}
	if base.Version == "" {
		base.Version = Version15
	}
	if base.Schema == "" {
		base.Schema = Schema15
	}

	// Decode body
	if len(bodyRaw) != 0 {
		var rawArr []json.RawMessage
		if err := json.Unmarshal(bodyRaw, &rawArr); err != nil {
			return fmt.Errorf("card.body: %w", err)
		}
		els, err := e.UnmarshalElementsSlice(rawArr)
		if err != nil {
			return fmt.Errorf("card.body: %w", err)
		}
		base.Body = els
	}

	// Decode actions
	if len(actionsRaw) != 0 {
		var rawArr []json.RawMessage
		if err := json.Unmarshal(actionsRaw, &rawArr); err != nil {
			return fmt.Errorf("card.actions: %w", err)
		}
		acts, err := a.UnmarshalActionsSlice(rawArr)
		if err != nil {
			return fmt.Errorf("card.actions: %w", err)
		}
		base.Actions = acts
	}

	// Decode selectAction
	if len(selectActionRaw) != 0 {
		act, err := a.UnmarshalAction(selectActionRaw)
		if err != nil {
			return fmt.Errorf("Card.selectAction: %w", err)
		}
		base.SelectAction = act
	}

	*c = Card(base)
	return nil
}
