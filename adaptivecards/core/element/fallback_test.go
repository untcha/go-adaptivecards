package element

import (
	"encoding/json"
	"strings"
	"testing"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

const fallbackTestType = m.TypeString("Test.FallbackElement")

type fallbackTestElement struct {
	Type  m.TypeString `json:"type"`
	Value string       `json:"value,omitempty"`
}

func (e fallbackTestElement) GetType() m.TypeString { return e.Type }

func init() {
	RegisterElement(fallbackTestType, func() Element {
		return &fallbackTestElement{Type: fallbackTestType}
	})
}

func TestElementFallbackJSON(t *testing.T) {
	drop := ElementFallbackDropOption()
	b, err := json.Marshal(drop)
	if err != nil {
		t.Fatalf("marshal drop fallback failed: %v", err)
	}
	if string(b) != `"drop"` {
		t.Fatalf("unexpected drop fallback json: %s", string(b))
	}

	content := ElementFallbackContent(fallbackTestElement{Type: fallbackTestType, Value: "fallback"})
	b, err = json.Marshal(content)
	if err != nil {
		t.Fatalf("marshal content fallback failed: %v", err)
	}
	if !strings.Contains(string(b), `"type":"Test.FallbackElement"`) {
		t.Fatalf("unexpected fallback element json: %s", string(b))
	}

	var decoded ElementFallback
	if err := json.Unmarshal([]byte(`{"type":"Test.FallbackElement","value":"ok"}`), &decoded); err != nil {
		t.Fatalf("unmarshal fallback content failed: %v", err)
	}
	if decoded.Content == nil {
		t.Fatalf("expected fallback content")
	}
}

func TestElementFallbackUnmarshalObjectBranches(t *testing.T) {
	var ef ElementFallback
	if err := json.Unmarshal([]byte(`{"type":"Test.FallbackElement","value":"ok"}`), &ef); err != nil {
		t.Fatalf("unexpected element fallback object unmarshal error: %v", err)
	}
	if ef.Content == nil {
		t.Fatalf("expected element fallback content")
	}
	if err := json.Unmarshal([]byte(`{"type":"Nope"}`), &ef); err == nil {
		t.Fatalf("expected element fallback object error")
	}
}
