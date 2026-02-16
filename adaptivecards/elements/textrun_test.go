package elements

import (
	"encoding/json"
	"testing"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

func TestTextRunStringShorthandAndObjectUnmarshal(t *testing.T) {
	var run TextRun
	if err := json.Unmarshal([]byte(`"Hello"`), &run); err != nil {
		t.Fatalf("unexpected shorthand unmarshal error: %v", err)
	}
	if run.Text != "Hello" || run.GetType() != m.TypeTextRun {
		t.Fatalf("unexpected shorthand text run: %+v", run)
	}

	if err := json.Unmarshal([]byte(`{"type":"TextRun","text":"Link","selectAction":{"type":"Action.OpenUrl","url":"https://example.com"}}`), &run); err != nil {
		t.Fatalf("unexpected object unmarshal error: %v", err)
	}
	if run.SelectAction == nil {
		t.Fatalf("expected selectAction")
	}
}

func TestTextRunValidationErrors(t *testing.T) {
	run := NewTextRun("")
	if err := run.Validate(); err == nil {
		t.Fatalf("expected required text error")
	}

	if err := json.Unmarshal([]byte(`{"type":"Wrong","text":"x"}`), &run); err == nil {
		t.Fatalf("expected type mismatch error")
	}
	if err := json.Unmarshal([]byte(`{"type":"TextRun","text":"x","selectAction":{"type":"Action.ShowCard"}}`), &run); err == nil {
		t.Fatalf("expected selectAction validation error")
	}
}
