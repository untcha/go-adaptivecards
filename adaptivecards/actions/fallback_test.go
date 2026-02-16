package actions

import (
	"encoding/json"
	"testing"
)

func TestActionFallbackUnmarshalObjectBranches(t *testing.T) {
	var af ActionFallback
	if err := json.Unmarshal([]byte(`{"type":"Action.OpenUrl","url":"https://example.com"}`), &af); err != nil {
		t.Fatalf("unexpected action fallback object unmarshal error: %v", err)
	}
	if af.Content == nil {
		t.Fatalf("expected action fallback content")
	}
	if err := json.Unmarshal([]byte(`{"type":"Nope"}`), &af); err == nil {
		t.Fatalf("expected action fallback object error")
	}
}
