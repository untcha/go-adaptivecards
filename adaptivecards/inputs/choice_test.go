package inputs

import (
	"testing"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

func TestInputChoiceValidateType(t *testing.T) {
	validEmptyType := InputChoice{Title: "A", Value: "a"}
	if err := validEmptyType.Validate(); err != nil {
		t.Fatalf("expected empty type to be valid: %v", err)
	}

	validCanonical := InputChoice{Type: m.TypeInputChoice, Title: "A", Value: "a"}
	if err := validCanonical.Validate(); err != nil {
		t.Fatalf("expected canonical type to be valid: %v", err)
	}

	invalidType := InputChoice{Type: m.TypeInputText, Title: "A", Value: "a"}
	if err := invalidType.Validate(); err == nil {
		t.Fatalf("expected invalid type error")
	}
}
