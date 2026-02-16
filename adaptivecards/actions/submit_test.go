package actions

import (
	"testing"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

func TestActionSubmitFactoryAndValidation(t *testing.T) {
	a := NewActionSubmit("Send", map[string]any{"k": "v"})
	a.AssociatedInputs = "auto"
	if err := a.Validate(); err != nil {
		t.Fatalf("unexpected validate error: %v", err)
	}

	a.AssociatedInputs = "bad"
	if err := a.Validate(); err == nil {
		t.Fatalf("expected associatedInputs error")
	}

	act, err := UnmarshalAction([]byte(`{"type":"Action.Submit","title":"Send","data":{"x":1},"associatedInputs":"none"}`))
	if err != nil {
		t.Fatalf("unexpected unmarshal action error: %v", err)
	}
	if act.GetType() != m.TypeActionSubmit {
		t.Fatalf("unexpected action type: %s", act.GetType())
	}
}
