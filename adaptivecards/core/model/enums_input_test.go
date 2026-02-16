package model

import (
	"encoding/json"
	"testing"
)

func TestInputStyleEnumsParseAndJSON(t *testing.T) {
	if got, err := ParseTextInputStyle("EMAIL"); err != nil || got != TextInputStyleEmail {
		t.Fatalf("expected ParseTextInputStyle to normalize EMAIL, got %q (err=%v)", got, err)
	}
	if got, err := ParseChoiceInputStyle("COMPACT"); err != nil || got != ChoiceInputStyleCompact {
		t.Fatalf("expected ParseChoiceInputStyle to normalize COMPACT, got %q (err=%v)", got, err)
	}

	var tis TextInputStyle
	if err := json.Unmarshal([]byte(`"PASSWORD"`), &tis); err != nil {
		t.Fatalf("expected TextInputStyle unmarshal to accept case-insensitive value: %v", err)
	}
	if tis != TextInputStylePassword {
		t.Fatalf("expected canonical password style, got %q", tis)
	}
	b, err := json.Marshal(tis)
	if err != nil {
		t.Fatalf("marshal text input style failed: %v", err)
	}
	if string(b) != `"password"` {
		t.Fatalf("expected canonical lowercase marshal, got %s", string(b))
	}

	var cis ChoiceInputStyle
	if err := json.Unmarshal([]byte(`"EXPANDED"`), &cis); err != nil {
		t.Fatalf("expected ChoiceInputStyle unmarshal to accept case-insensitive value: %v", err)
	}
	if cis != ChoiceInputStyleExpanded {
		t.Fatalf("expected canonical expanded style, got %q", cis)
	}
}
