package inputs

import (
	"encoding/json"
	"strings"
	"testing"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

func TestInputChoiceSetValidateAndUnmarshal(t *testing.T) {
	choice := NewInputChoice("A", "a")
	in := NewInputChoiceSet("choice", choice)
	in.Style = m.ChoiceInputStyleCompact
	if err := in.Validate(); err != nil {
		t.Fatalf("unexpected validation error: %v", err)
	}

	var decoded InputChoiceSet
	if err := json.Unmarshal([]byte(`{"type":"Input.ChoiceSet","id":"x","choices":[{"title":"A","value":"a"}],"style":"expanded"}`), &decoded); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if decoded.Type != m.TypeInputChoiceSet {
		t.Fatalf("expected input choice set type")
	}
	var decodedCase InputChoiceSet
	if err := json.Unmarshal([]byte(`{"type":"Input.ChoiceSet","id":"x","choices":[{"title":"A","value":"a"}],"style":"COMPACT"}`), &decodedCase); err != nil {
		t.Fatalf("unexpected case-insensitive unmarshal error: %v", err)
	}
	if decodedCase.Style != m.ChoiceInputStyleCompact {
		t.Fatalf("expected canonical style %q, got %q", m.ChoiceInputStyleCompact, decodedCase.Style)
	}
	out, err := json.Marshal(decodedCase)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	if !strings.Contains(string(out), `"style":"compact"`) {
		t.Fatalf("expected canonical lowercase style in json, got %s", string(out))
	}

	bad := NewInputChoiceSet("x", InputChoice{Title: "", Value: "a"})
	if err := bad.Validate(); err == nil {
		t.Fatalf("expected invalid choice error")
	}
}
