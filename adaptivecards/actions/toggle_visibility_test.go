package actions

import (
	"encoding/json"
	"testing"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

func TestActionToggleVisibilityValidateAndFactory(t *testing.T) {
	visible := true
	act := NewActionToggleVisibility(
		"Toggle",
		TargetElement{ElementID: "details"},
		TargetElement{Type: "TargetElement", ElementID: "extra", IsVisible: &visible},
	)
	if err := act.Validate(); err != nil {
		t.Fatalf("unexpected validation error: %v", err)
	}

	decoded, err := UnmarshalAction([]byte(`{"type":"Action.ToggleVisibility","title":"Toggle","targetElements":["details",{"type":"TargetElement","elementId":"extra","isVisible":true}]}`))
	if err != nil {
		t.Fatalf("unexpected action unmarshal error: %v", err)
	}
	if decoded.GetType() != m.TypeActionToggleVisibility {
		t.Fatalf("unexpected action type: %s", decoded.GetType())
	}
}

func TestTargetElementUnmarshalErrors(t *testing.T) {
	var target TargetElement
	if err := json.Unmarshal([]byte(`""`), &target); err == nil {
		t.Fatalf("expected empty target element error")
	}
	if err := json.Unmarshal([]byte(`{"type":"Wrong","elementId":"x"}`), &target); err == nil {
		t.Fatalf("expected target element type error")
	}
	if err := json.Unmarshal([]byte(`{"type":"TargetElement","elementId":"  "}`), &target); err == nil {
		t.Fatalf("expected required elementId error")
	}
}

func TestActionToggleVisibilityUnmarshalErrors(t *testing.T) {
	var act ActionToggleVisibility
	if err := json.Unmarshal([]byte(`{"type":"Action.OpenUrl","targetElements":["x"]}`), &act); err == nil {
		t.Fatalf("expected type mismatch error")
	}
	if err := json.Unmarshal([]byte(`{"type":"Action.ToggleVisibility","targetElements":[]}`), &act); err == nil {
		t.Fatalf("expected targetElements required error")
	}
}

func TestActionToggleVisibilityValidateRejectsInvalidTargetType(t *testing.T) {
	act := NewActionToggleVisibility(
		"Toggle",
		TargetElement{Type: "WrongType", ElementID: "x"},
	)
	if err := act.Validate(); err == nil {
		t.Fatalf("expected invalid target type error")
	}
}
