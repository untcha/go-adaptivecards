package inputs

import (
	"encoding/json"
	"strings"
	"testing"

	a "github.com/untcha/go-adaptivecards/adaptivecards/actions"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

func TestInputTextValidateAndUnmarshal(t *testing.T) {
	open, err := a.NewActionOpenURL("Open", "https://example.com")
	if err != nil {
		t.Fatalf("setup failed: %v", err)
	}
	in := NewInputText("ticket")
	in.Style = m.TextInputStyleEmail
	maxLength := 120.0
	in.MaxLength = &maxLength
	in.InlineAction = open
	if err := in.Validate(); err != nil {
		t.Fatalf("unexpected validation error: %v", err)
	}

	var decoded InputText
	if err := json.Unmarshal([]byte(`{"type":"Input.Text","id":"x","style":"text","inlineAction":{"type":"Action.Submit","title":"Go"}}`), &decoded); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if decoded.Type != m.TypeInputText {
		t.Fatalf("expected input text type")
	}
	var decodedCase InputText
	if err := json.Unmarshal([]byte(`{"type":"Input.Text","id":"x","style":"EMAIL"}`), &decodedCase); err != nil {
		t.Fatalf("unexpected case-insensitive unmarshal error: %v", err)
	}
	if decodedCase.Style != m.TextInputStyleEmail {
		t.Fatalf("expected canonical style %q, got %q", m.TextInputStyleEmail, decodedCase.Style)
	}
	out, err := json.Marshal(decodedCase)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	if string(out) == "" || !json.Valid(out) {
		t.Fatalf("marshal produced invalid json")
	}
	if want := `"style":"email"`; !strings.Contains(string(out), want) {
		t.Fatalf("expected canonical lowercase style in json, missing %s in %s", want, string(out))
	}

	bad := NewInputText("")
	if err := bad.Validate(); err == nil {
		t.Fatalf("expected required id error")
	}
}

func TestInputTextMaxLengthSchemaNumber(t *testing.T) {
	in := NewInputText("ticket")
	fractional := 1.5
	in.MaxLength = &fractional
	if err := in.Validate(); err != nil {
		t.Fatalf("unexpected validation error for fractional maxLength: %v", err)
	}

	zero := 0.0
	in.MaxLength = &zero
	if err := in.Validate(); err != nil {
		t.Fatalf("unexpected validation error for zero maxLength: %v", err)
	}

	negative := -1.0
	in.MaxLength = &negative
	if err := in.Validate(); err != nil {
		t.Fatalf("expected schema-tolerant validation for negative maxLength, got: %v", err)
	}
}
