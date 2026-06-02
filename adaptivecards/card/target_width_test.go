package card

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"

	con "github.com/untcha/go-adaptivecards/adaptivecards/containers"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
	els "github.com/untcha/go-adaptivecards/adaptivecards/elements"
)

func TestTargetWidthMarshalsWhenSet(t *testing.T) {
	tb := els.NewTextBlock("hi").WithTargetWidth(m.TargetWidthAtLeast(m.TargetWidthWide))
	b, err := json.Marshal(tb)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	if !strings.Contains(string(b), `"targetWidth":"atLeast:wide"`) {
		t.Fatalf("expected targetWidth in JSON, got: %s", b)
	}
}

func TestTargetWidthAbsentByDefault(t *testing.T) {
	b, err := json.Marshal(els.NewTextBlock("hi"))
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	if strings.Contains(string(b), "targetWidth") {
		t.Fatalf("expected no targetWidth key, got: %s", b)
	}
}

func TestTargetWidthRoundTrip(t *testing.T) {
	in := `{"type":"TextBlock","text":"hi","targetWidth":"atMost:narrow"}`
	var tb els.TextBlock
	if err := json.Unmarshal([]byte(in), &tb); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if tb.TargetWidth != "atMost:narrow" {
		t.Fatalf("expected atMost:narrow, got %q", tb.TargetWidth)
	}
}

// TestValidateAcceptsNestedTargetWidth exercises the recursive schema strip:
// a targetWidth on an element nested inside a Container must not trip the
// embedded schema's additionalProperties:false.
func TestValidateAcceptsNestedTargetWidth(t *testing.T) {
	nested := els.NewTextBlock("deep").WithTargetWidth(m.TargetWidthNarrow)
	container := con.NewContainer(nested)

	c, err := NewCard().AddElement(container).Build()
	if err != nil {
		t.Fatalf("build failed: %v", err)
	}
	if err := c.Validate(); err != nil {
		t.Fatalf("expected card with nested targetWidth to validate, got: %v", err)
	}
}

func TestValidateRejectsInvalidTargetWidth(t *testing.T) {
	tb := els.NewTextBlock("hi")
	tb.TargetWidth = "huge"

	c, err := NewCard().AddElement(tb).Build()
	if err != nil {
		t.Fatalf("build failed: %v", err)
	}
	err = c.Validate()
	if err == nil {
		t.Fatalf("expected validation error for invalid targetWidth")
	}
	if !errors.Is(err, m.ErrInvalidEnum) {
		t.Fatalf("expected ErrInvalidEnum, got: %v", err)
	}
	if !strings.Contains(err.Error(), "targetWidth") {
		t.Fatalf("expected targetWidth in error, got: %v", err)
	}
}
