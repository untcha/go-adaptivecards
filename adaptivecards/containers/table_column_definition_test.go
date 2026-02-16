package containers

import (
	"encoding/json"
	"testing"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

func TestTableColumnDefinitionFluentMethods(t *testing.T) {
	col := NewTableColumnDefinition().
		WithWidth(2).
		WithPixelWidth("10px").
		WithRelativeWidth(2).
		WithHorizontalCellContentAlignment(m.HAlignLeft).
		WithVerticalCellContentAlignment(m.VAlignTop).
		AlignCenter().
		AlignMiddle()

	if err := col.Validate(); err != nil {
		t.Fatalf("unexpected column validate error: %v", err)
	}
}

func TestTableColumnDefinitionDecodeErrorBranches(t *testing.T) {
	var col TableColumnDefinition
	if err := json.Unmarshal([]byte(`{"type":"Wrong","width":1}`), &col); err == nil {
		t.Fatalf("expected table column type error")
	}
	if err := json.Unmarshal([]byte(`{"width":0}`), &col); err == nil {
		t.Fatalf("expected invalid width error")
	}
	if err := json.Unmarshal([]byte(`{"width":"bad"}`), &col); err == nil {
		t.Fatalf("expected invalid width string error")
	}
	col = NewTableColumnDefinition().AlignLeft().AlignRight().AlignTop().AlignBottom()
	if err := col.Validate(); err != nil {
		t.Fatalf("unexpected column validate error: %v", err)
	}
}
