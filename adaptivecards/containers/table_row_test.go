package containers

import (
	"encoding/json"
	"testing"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
	els "github.com/untcha/go-adaptivecards/adaptivecards/elements"
)

func TestTableRowGetType(t *testing.T) {
	if NewTableRowEmpty().GetType() != m.TypeTableRow {
		t.Fatalf("unexpected table row type")
	}
}

func TestTableRowBuilderMethods(t *testing.T) {
	row := NewTableRowEmpty().
		AddTextCell("A").
		AddImageCell("https://example.com/a.png").
		AddCell(NewTableCell(els.NewTextBlock("B"))).
		AddEmptyCell().
		WithStyle(m.ContainerStyleAccent).
		WithHorizontalCellContentAlignment(m.HAlignCenter).
		WithVerticalCellContentAlignment(m.VAlignCenter)

	if len(row.Cells) != 4 {
		t.Fatalf("expected 4 cells, got %d", len(row.Cells))
	}
	if row.Style != m.ContainerStyleAccent {
		t.Fatalf("expected style accent, got %q", row.Style)
	}
	if row.HorizontalCellContentAlignment != m.HAlignCenter {
		t.Fatalf("expected horizontal alignment center, got %q", row.HorizontalCellContentAlignment)
	}
	if row.VerticalCellContentAlignment != m.VAlignCenter {
		t.Fatalf("expected vertical alignment center, got %q", row.VerticalCellContentAlignment)
	}
}

func TestTableRowDecodeErrorBranches(t *testing.T) {
	var row TableRow
	if err := json.Unmarshal([]byte(`{"type":"Wrong","cells":[]}`), &row); err == nil {
		t.Fatalf("expected table row type error")
	}
	if err := json.Unmarshal([]byte(`{"cells":[],"style":"bad"}`), &row); err == nil {
		t.Fatalf("expected table row style error")
	}
	if err := json.Unmarshal([]byte(`{"cells":[],"horizontalCellContentAlignment":"bad"}`), &row); err == nil {
		t.Fatalf("expected row horizontal alignment error")
	}
}
