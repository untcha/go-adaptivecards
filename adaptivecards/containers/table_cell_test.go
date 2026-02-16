package containers

import (
	"encoding/json"
	"strings"
	"testing"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
	els "github.com/untcha/go-adaptivecards/adaptivecards/elements"
)

type fakeOpenURLAction struct{}

func (fakeOpenURLAction) GetType() m.TypeString { return m.TypeActionOpenURL }

func TestTableCellGetType(t *testing.T) {
	if NewTableCell(els.NewTextBlock("x")).GetType() != m.TypeTableCell {
		t.Fatalf("unexpected table cell type")
	}
}

func TestTableCellConvenienceMethodsAndUnmarshalErrors(t *testing.T) {
	cell := NewTableCellEmpty().
		AddElement(els.NewTextBlock("x")).
		AddTextBlock(els.TextBlock{Text: "y"}).
		AddImage(els.Image{URL: "https://example.com/img.png"}).
		WithSelectAction(fakeOpenURLAction{}).
		WithStyle(m.ContainerStyleDefault).
		WithVerticalContentAlignment(m.VContentAlignTop).
		WithBleed(true).
		WithMinHeight("20px").
		WithRtl(true).
		StyleDefault().
		StyleEmphasis().
		StyleGood().
		StyleAttention().
		StyleWarning().
		StyleAccent().
		AlignTop().
		AlignCenter().
		AlignBottom()

	if err := cell.Validate(); err != nil {
		t.Fatalf("unexpected tablecell validate error: %v", err)
	}

	var tc TableCell
	if err := json.Unmarshal([]byte(`{"type":"Wrong","items":[]}`), &tc); err == nil {
		t.Fatalf("expected type error")
	}
	if err := json.Unmarshal([]byte(`{"type":"TableCell","selectAction":{"type":"Nope"},"items":[{"type":"TextBlock","text":"x"}]}`), &tc); err == nil {
		t.Fatalf("expected selectAction error")
	}
}

func TestTableCellWithBackgroundImageURLKeepsInvalidInputForValidation(t *testing.T) {
	tc := NewTableCell(els.NewTextBlock("cell")).WithBackgroundImageURL("not-a-url")

	if tc.BackgroundImage == nil {
		t.Fatalf("expected background image to be set")
	}

	err := tc.Validate()
	if err == nil {
		t.Fatalf("expected validation error for invalid background image url")
	}
	if !strings.Contains(err.Error(), "TableCell.backgroundImage") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestTableCellWithBackgroundImageKeepsInvalidObjectForValidation(t *testing.T) {
	invalid := m.BackgroundImage{
		URL:      "",
		FillMode: m.ImageFillCover,
	}

	tc := NewTableCell(els.NewTextBlock("cell")).WithBackgroundImage(invalid)
	if tc.BackgroundImage == nil {
		t.Fatalf("expected background image to be set")
	}

	err := tc.Validate()
	if err == nil {
		t.Fatalf("expected validation error for invalid background image object")
	}
	if !strings.Contains(err.Error(), "TableCell.backgroundImage") {
		t.Fatalf("unexpected error: %v", err)
	}
}
