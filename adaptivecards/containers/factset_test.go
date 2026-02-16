package containers

import (
	"encoding/json"
	"testing"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

func TestFactSetValidateAndUnmarshal(t *testing.T) {
	fs := NewFactSet(
		NewFact("Env", "Prod"),
		NewFact("Version", "1.0.0"),
	)
	if err := fs.Validate(); err != nil {
		t.Fatalf("unexpected validation error: %v", err)
	}

	var decoded FactSet
	if err := json.Unmarshal([]byte(`{"type":"FactSet","facts":[{"title":"A","value":"a"}]}`), &decoded); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if decoded.Type != m.TypeFactSet {
		t.Fatalf("expected factset type")
	}

	bad := NewFactSet(Fact{Title: "", Value: "x"})
	if err := bad.Validate(); err == nil {
		t.Fatalf("expected invalid fact error")
	}
}
