package inputs

import (
	"encoding/json"
	"testing"
)

func TestInputToggleValidateAndUnmarshal(t *testing.T) {
	in := NewInputToggle("notify", "Notify me")
	if err := in.Validate(); err != nil {
		t.Fatalf("unexpected validation error: %v", err)
	}

	var decoded InputToggle
	if err := json.Unmarshal([]byte(`{"type":"Input.Toggle","id":"x","title":"T"}`), &decoded); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if decoded.ValueOff != "false" || decoded.ValueOn != "true" {
		t.Fatalf("expected default toggle values")
	}

	bad := NewInputToggle("x", "")
	if err := bad.Validate(); err == nil {
		t.Fatalf("expected required title error")
	}
}
