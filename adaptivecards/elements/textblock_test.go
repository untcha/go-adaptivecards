package elements

import (
	"encoding/json"
	"testing"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

func TestTextBlockGetTypeAndValidateAndUnmarshal(t *testing.T) {
	if NewTextBlock("x").GetType() != m.TypeTextBlock {
		t.Fatalf("unexpected textblock type")
	}

	tb := NewTextBlock("hello").
		WithColor(m.ColorAccent).
		WithFontType(m.FontTypeDefault).
		WithAlign(m.HAlignCenter).
		WithSize(m.SizeMedium).
		WithWeight(m.WeightBolder).
		WithWrap(true).
		WithStyle(m.TextBlockStyleDefault)
	if err := tb.Validate(); err != nil {
		t.Fatalf("unexpected textblock validation error: %v", err)
	}

	var decoded TextBlock
	if err := json.Unmarshal([]byte(`{"text":"hi","maxLines":2}`), &decoded); err != nil {
		t.Fatalf("unexpected textblock unmarshal error: %v", err)
	}
	if decoded.Type != m.TypeTextBlock {
		t.Fatalf("expected textblock type default, got %q", decoded.Type)
	}
}

func TestTextBlockFluentMethods(t *testing.T) {
	tb := NewTextBlockEmpty().
		WithText("hello").
		WithColor(m.ColorAccent).
		WithFontType(m.FontTypeMonospace).
		AlignLeft().
		WithSubtle(true).
		WithMaxLines(2).
		WithSize(m.SizeMedium).
		WithWeight(m.WeightBolder).
		WrapOn().
		WithStyle(m.TextBlockStyleHeading).
		WithHeight(m.BlockElementHeightStretch).
		WithSeparator(true).
		WithSpacing(m.SpacingSmall).
		WithID("id-1").
		WithVisible(true).
		WithRequires(map[string]string{"feature": "1.0"})

	if err := tb.Validate(); err != nil {
		t.Fatalf("unexpected textblock validate error: %v", err)
	}
}

func TestTextBlockValidationErrors(t *testing.T) {
	if err := NewTextBlock("").Validate(); err == nil {
		t.Fatalf("expected required text validation error")
	}

	var decoded TextBlock
	if err := json.Unmarshal([]byte(`{"type":"Wrong","text":"x"}`), &decoded); err == nil {
		t.Fatalf("expected textblock type mismatch error")
	}
	if err := json.Unmarshal([]byte(`{"type":"TextBlock","text":"x","maxLines":-1}`), &decoded); err != nil {
		t.Fatalf("expected schema-tolerant maxLines number, got error: %v", err)
	}
}
