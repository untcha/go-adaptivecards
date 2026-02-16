package actions

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// ErrUnknownActionType is returned when no factory is registered for a given "type".
var ErrUnknownActionType = errors.New("unknown action type")

// actionFactory constructs a *concrete* Action ready for json.Unmarshal into it.
type actionFactory func() Action

var (
	actMu       sync.RWMutex
	actRegistry = make(map[m.TypeString]actionFactory)
)

// RegisterAction registers a constructor for a given action type string (e.g., "Action.Submit").
// Panics on duplicate registrations or nil factory.
func RegisterAction(t m.TypeString, fac actionFactory) {
	if fac == nil {
		panic("RegisterAction: nil factory")
	}
	actMu.Lock()
	defer actMu.Unlock()
	if _, ok := actRegistry[t]; ok {
		panic("RegisterAction: duplicate registration for type " + string(t))
	}
	actRegistry[t] = fac
}

// UnmarshalAction inspects the "type" property of b and unmarshals it into the matching concrete Action.
func UnmarshalAction(b []byte) (Action, error) {
	var probe struct {
		Type m.TypeString `json:"type"`
	}
	if err := json.Unmarshal(b, &probe); err != nil {
		return nil, fmt.Errorf("action: probe type: %w", err)
	}
	if probe.Type == "" {
		return nil, fmt.Errorf("action: missing required \"type\" property")
	}

	actMu.RLock()
	fac, ok := actRegistry[probe.Type]
	actMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("%w: %q", ErrUnknownActionType, probe.Type)
	}

	a := fac()
	if err := json.Unmarshal(b, a); err != nil {
		return nil, fmt.Errorf("action %q: %w", probe.Type, err)
	}

	// Ensure the concrete type matches the probed JSON type.
	if gt := a.GetType(); gt != probe.Type {
		return nil, fmt.Errorf("action %q: type mismatch (GetType=%q)", probe.Type, gt)
	}

	return a, nil
}

// UnmarshalActionsSlice decodes a JSON array into []Action using the registry.
func UnmarshalActionsSlice(raw []json.RawMessage) ([]Action, error) {
	out := make([]Action, 0, len(raw))
	for i, r := range raw {
		a, err := UnmarshalAction(r)
		if err != nil {
			return nil, fmt.Errorf("actions[%d]: %w", i, err)
		}
		out = append(out, a)
	}
	return out, nil
}
