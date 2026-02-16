package element

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// ErrUnknownElementType is returned when no factory is registered for a given "type".
var ErrUnknownElementType = errors.New("unknown element type")

// elementFactory constructs a *concrete* Element ready for json.Unmarshal into it.
type elementFactory func() Element

var (
	elemMu       sync.RWMutex
	elemRegistry = make(map[m.TypeString]elementFactory)
)

// RegisterElement registers a constructor for a given element type string (e.g., "TextBlock").
// Panics on duplicate registrations or nil factory.
func RegisterElement(t m.TypeString, fac elementFactory) {
	if fac == nil {
		panic("RegisterElement: nil factory")
	}
	elemMu.Lock()
	defer elemMu.Unlock()
	if _, exists := elemRegistry[t]; exists {
		panic("RegisterElement: duplicate registration for type " + string(t))
	}
	elemRegistry[t] = fac
}

// UnmarshalElement inspects the "type" property of b and unmarshals it into the matching concrete Element.
func UnmarshalElement(b []byte) (Element, error) {
	// 1) Peek the type
	var probe struct {
		Type m.TypeString `json:"type"`
	}
	if err := json.Unmarshal(b, &probe); err != nil {
		return nil, fmt.Errorf("element: probe type: %w", err)
	}
	if probe.Type == "" {
		return nil, fmt.Errorf("element: missing required \"type\" property")
	}

	// 2) Look up the factory
	elemMu.RLock()
	fac, ok := elemRegistry[probe.Type]
	elemMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("%w: %q", ErrUnknownElementType, probe.Type)
	}

	// 3) Construct and decode
	el := fac() // concrete pointer that implements Element
	if err := json.Unmarshal(b, el); err != nil {
		return nil, fmt.Errorf("element %q: %w", probe.Type, err)
	}

	return el, nil
}

// UnmarshalElementsSlice decodes a JSON array into []Element using the registry.
// Useful for Card.Body, Container.Items, Column.Items, etc.
func UnmarshalElementsSlice(rawArr []json.RawMessage) ([]Element, error) {
	out := make([]Element, 0, len(rawArr))
	for i, rm := range rawArr {
		el, err := UnmarshalElement(rm)
		if err != nil {
			return nil, fmt.Errorf("elements[%d]: %w", i, err)
		}
		out = append(out, el)
	}
	return out, nil
}
