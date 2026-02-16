package actions

import (
	"encoding/json"
	"strings"
	"testing"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

type factoryTestAction struct{}

func (factoryTestAction) GetType() m.TypeString { return m.TypeActionSubmit }

func TestRegisterActionPanicsOnNilFactory(t *testing.T) {
	tpe := m.TypeString("Action.Test.NilFactory")

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic for nil factory")
		}
	}()

	RegisterAction(tpe, nil)
}

func TestRegisterActionPanicsOnDuplicate(t *testing.T) {
	tpe := m.TypeString("Action.Test.Duplicate")
	RegisterAction(tpe, func() Action { return &ActionSubmit{Type: m.TypeActionSubmit} })

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic for duplicate registration")
		}
	}()

	RegisterAction(tpe, func() Action { return &ActionSubmit{Type: m.TypeActionSubmit} })
}

func TestUnmarshalActionRejectsFactoryTypeMismatch(t *testing.T) {
	tpe := m.TypeString("Action.Test.Mismatch")
	RegisterAction(tpe, func() Action { return &factoryTestAction{} })

	_, err := UnmarshalAction([]byte(`{"type":"Action.Test.Mismatch","foo":"bar"}`))
	if err == nil {
		t.Fatalf("expected type mismatch error")
	}
	if !strings.Contains(err.Error(), "type mismatch") {
		t.Fatalf("expected type mismatch in error, got: %v", err)
	}
}

func TestUnmarshalActionsSliceIncludesIndexOnError(t *testing.T) {
	var raw []json.RawMessage
	if err := json.Unmarshal(
		[]byte(`[{"type":"Action.OpenUrl","url":"https://example.com"},{"type":"Action.Unknown"}]`),
		&raw,
	); err != nil {
		t.Fatalf("setup unmarshal failed: %v", err)
	}

	_, err := UnmarshalActionsSlice(raw)
	if err == nil {
		t.Fatalf("expected indexed error")
	}
	if !strings.Contains(err.Error(), "actions[1]") {
		t.Fatalf("expected error to include failing index, got: %v", err)
	}
}

func TestActionFactoryDeserializeRegisteredActionTypes(t *testing.T) {
	cases := []struct {
		name     string
		raw      string
		wantType m.TypeString
	}{
		{
			name:     "openurl",
			raw:      `{"type":"Action.OpenUrl","url":"https://example.com"}`,
			wantType: m.TypeActionOpenURL,
		},
		{
			name:     "submit",
			raw:      `{"type":"Action.Submit","title":"Send","data":{"x":1}}`,
			wantType: m.TypeActionSubmit,
		},
		{
			name:     "toggle_visibility",
			raw:      `{"type":"Action.ToggleVisibility","targetElements":["details"]}`,
			wantType: m.TypeActionToggleVisibility,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			act, err := UnmarshalAction([]byte(tc.raw))
			if err != nil {
				t.Fatalf("unexpected unmarshal action error for %s: %v", tc.raw, err)
			}
			if act.GetType() != tc.wantType {
				t.Fatalf("expected type %q, got %q", tc.wantType, act.GetType())
			}
		})
	}
}
