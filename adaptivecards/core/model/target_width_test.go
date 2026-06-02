package model

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestTargetWidthIsValid(t *testing.T) {
	cases := []struct {
		in   TargetWidth
		want bool
	}{
		// bare buckets
		{TargetWidthVeryNarrow, true},
		{TargetWidthNarrow, true},
		{TargetWidthStandard, true},
		{TargetWidthWide, true},
		// prefixed range forms
		{"atLeast:narrow", true},
		{"atMost:wide", true},
		{"atLeast:veryNarrow", true},
		{"atMost:standard", true},
		// invalid
		{"", false},
		{"huge", false},
		{"atLeast:", false},
		{"atMost:", false},
		{"atLeast:huge", false},
		{"atMost:tiny", false},
		{"atLeast:atLeast:wide", false},
		{"exactly:wide", false},
		{"Narrow", false}, // case-sensitive
	}
	for _, c := range cases {
		if got := c.in.IsValid(); got != c.want {
			t.Errorf("TargetWidth(%q).IsValid() = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestTargetWidthBuilders(t *testing.T) {
	if got := TargetWidthAtLeast(TargetWidthWide); got != "atLeast:wide" {
		t.Errorf("TargetWidthAtLeast = %q", got)
	}
	if got := TargetWidthAtMost(TargetWidthNarrow); got != "atMost:narrow" {
		t.Errorf("TargetWidthAtMost = %q", got)
	}
}

func TestParseTargetWidth(t *testing.T) {
	v, err := ParseTargetWidth("atLeast:wide")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v != "atLeast:wide" {
		t.Fatalf("got %q", v)
	}

	_, err = ParseTargetWidth("nope")
	if err == nil {
		t.Fatalf("expected error for invalid value")
	}
	if !errors.Is(err, ErrInvalidEnum) {
		t.Fatalf("expected ErrInvalidEnum, got %v", err)
	}
}

func TestTargetWidthUnmarshalJSON(t *testing.T) {
	var w TargetWidth
	if err := json.Unmarshal([]byte(`"atMost:narrow"`), &w); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if w != "atMost:narrow" {
		t.Fatalf("got %q", w)
	}

	// empty string is accepted (absent/cleared)
	if err := json.Unmarshal([]byte(`""`), &w); err != nil {
		t.Fatalf("unexpected error for empty: %v", err)
	}

	if err := json.Unmarshal([]byte(`"bogus"`), &w); err == nil {
		t.Fatalf("expected error for invalid value")
	}
}

func TestTargetWidthMarshalJSON(t *testing.T) {
	b, err := json.Marshal(TargetWidthWide)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(b) != `"wide"` {
		t.Fatalf("got %s", b)
	}

	if _, err := json.Marshal(TargetWidth("bogus")); err == nil {
		t.Fatalf("expected error marshaling invalid value")
	}
}
