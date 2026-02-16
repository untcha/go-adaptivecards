package elements

import (
	"encoding/json"
	"testing"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

func TestRichTextBlockValidateAndUnmarshal(t *testing.T) {
	block := NewRichTextBlock(
		NewTextRun("Part A"),
		TextRun{Text: "Part B", Weight: m.WeightDefault},
	)
	block.HorizontalAlignment = m.HAlignCenter
	if err := block.Validate(); err != nil {
		t.Fatalf("unexpected richtext validation error: %v", err)
	}

	var decoded RichTextBlock
	if err := json.Unmarshal([]byte(`{"type":"RichTextBlock","inlines":["A",{"type":"TextRun","text":"B"}]}`), &decoded); err != nil {
		t.Fatalf("unexpected richtext unmarshal error: %v", err)
	}
	if decoded.Type != m.TypeRichTextBlock {
		t.Fatalf("expected richtextblock type")
	}
	if len(decoded.Inlines) != 2 {
		t.Fatalf("expected 2 inlines")
	}
}

func TestRichTextBlockValidationErrors(t *testing.T) {
	block := NewRichTextBlock()
	if err := block.Validate(); err == nil {
		t.Fatalf("expected inlines required error")
	}

	var decoded RichTextBlock
	if err := json.Unmarshal([]byte(`{"type":"Wrong","inlines":["x"]}`), &decoded); err == nil {
		t.Fatalf("expected richtext type mismatch")
	}
}
