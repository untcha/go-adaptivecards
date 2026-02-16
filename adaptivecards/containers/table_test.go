package containers

import (
	"encoding/json"
	"strings"
	"testing"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
	els "github.com/untcha/go-adaptivecards/adaptivecards/elements"
)

// TestNewTable tests the Table constructors
func TestNewTable(t *testing.T) {
	t.Run("NewTable creates table with defaults", func(t *testing.T) {
		table := NewTable()

		if table.Type != m.TypeTable {
			t.Errorf("expected Type %q, got %q", m.TypeTable, table.Type)
		}

		if table.FirstRowAsHeader == nil || *table.FirstRowAsHeader != true {
			t.Errorf("expected FirstRowAsHeader to be true by default")
		}

		if table.ShowGridLines == nil || *table.ShowGridLines != true {
			t.Errorf("expected ShowGridLines to be true by default")
		}

		if table.GridStyle != m.ContainerStyleDefault {
			t.Errorf("expected GridStyle to be default, got %q", table.GridStyle)
		}

		if len(table.Columns) != 0 {
			t.Errorf("expected empty columns slice")
		}

		if len(table.Rows) != 0 {
			t.Errorf("expected empty rows slice")
		}
	})

	t.Run("NewTableWithColumnsAndRows creates table with data", func(t *testing.T) {
		columns := []TableColumnDefinition{
			NewTableColumnDefinitionWithWidth(1),
			NewTableColumnDefinitionWithWidth("100px"),
		}
		rows := []TableRow{
			NewTableRow(NewTableCell(els.NewTextBlock("Cell 1")), NewTableCell(els.NewTextBlock("Cell 2"))),
		}

		table := NewTableWithColumnsAndRows(columns, rows)

		if len(table.Columns) != 2 {
			t.Errorf("expected 2 columns, got %d", len(table.Columns))
		}

		if len(table.Rows) != 1 {
			t.Errorf("expected 1 row, got %d", len(table.Rows))
		}
	})
}

// TestTableBuilderMethods tests the fluent builder methods
func TestTableBuilderMethods(t *testing.T) {
	table := NewTable()

	t.Run("AddColumn", func(t *testing.T) {
		column := NewTableColumnDefinitionWithWidth(2)
		result := table.AddColumn(column)

		if len(result.Columns) != 1 {
			t.Errorf("expected 1 column after adding, got %d", len(result.Columns))
		}

		if result.Columns[0].Width != 2 {
			t.Errorf("expected column width 2, got %v", result.Columns[0].Width)
		}
	})

	t.Run("AddColumnWithWidth", func(t *testing.T) {
		result := table.AddColumnWithWidth("150px")

		if len(result.Columns) != 1 {
			t.Errorf("expected 1 column after adding, got %d", len(result.Columns))
		}

		if result.Columns[0].Width != "150px" {
			t.Errorf("expected column width 150px, got %v", result.Columns[0].Width)
		}
	})

	t.Run("AddTextRow", func(t *testing.T) {
		result := table.AddTextRow("Cell 1", "Cell 2", "Cell 3")

		if len(result.Rows) != 1 {
			t.Errorf("expected 1 row after adding, got %d", len(result.Rows))
		}

		if len(result.Rows[0].Cells) != 3 {
			t.Errorf("expected 3 cells in row, got %d", len(result.Rows[0].Cells))
		}
	})

	t.Run("Configuration methods", func(t *testing.T) {
		result := table.
			WithFirstRowAsHeader(false).
			WithShowGridLines(false).
			GridStyleEmphasis().
			CellsAlignCenter()

		if result.FirstRowAsHeader == nil || *result.FirstRowAsHeader != false {
			t.Errorf("expected FirstRowAsHeader to be false")
		}

		if result.ShowGridLines == nil || *result.ShowGridLines != false {
			t.Errorf("expected ShowGridLines to be false")
		}

		if result.GridStyle != m.ContainerStyleEmphasis {
			t.Errorf("expected GridStyle emphasis, got %q", result.GridStyle)
		}

		if result.HorizontalCellContentAlignment != m.HAlignCenter {
			t.Errorf("expected horizontal alignment center, got %q", result.HorizontalCellContentAlignment)
		}
	})
}

// TestTableValidation tests the Table validation
func TestTableValidation(t *testing.T) {
	t.Run("valid empty table", func(t *testing.T) {
		table := NewTable()
		if err := table.Validate(); err != nil {
			t.Errorf("expected no validation error for empty table, got %v", err)
		}
	})

	t.Run("valid table with data", func(t *testing.T) {
		table := NewTable().
			AddColumnWithWidth(1).
			AddColumnWithWidth("100px").
			AddTextRow("Header 1", "Header 2").
			AddTextRow("Cell 1", "Cell 2")

		if err := table.Validate(); err != nil {
			t.Errorf("expected no validation error for valid table, got %v", err)
		}
	})

	t.Run("invalid column", func(t *testing.T) {
		invalidColumn := TableColumnDefinition{
			Type:  m.TypeTableColumnDefinition,
			Width: "invalid-width",
		}
		table := NewTable().AddColumn(invalidColumn)

		if err := table.Validate(); err == nil {
			t.Errorf("expected validation error for invalid column")
		}
	})

	t.Run("row with too many cells", func(t *testing.T) {
		table := NewTable().
			AddColumnWithWidth(1).                   // Only 1 column
			AddTextRow("Cell 1", "Cell 2", "Cell 3") // 3 cells

		if err := table.Validate(); err == nil {
			t.Errorf("expected validation error for row with too many cells")
		} else if !strings.Contains(err.Error(), "row has 3 cells but table has only 1 columns") {
			t.Errorf("expected specific error message about too many cells, got %v", err)
		}
	})

	t.Run("invalid grid style", func(t *testing.T) {
		table := NewTable()
		table.GridStyle = "invalid-style"

		if err := table.Validate(); err == nil {
			t.Errorf("expected validation error for invalid grid style")
		}
	})
}

// TestTableJSON tests JSON marshaling and unmarshaling
func TestTableJSON(t *testing.T) {
	t.Run("marshal with defaults", func(t *testing.T) {
		table := NewTable()
		data, err := json.Marshal(table)
		if err != nil {
			t.Errorf("unexpected marshal error: %v", err)
			return
		}

		// Check that defaults are included
		jsonStr := string(data)
		if !strings.Contains(jsonStr, `"type":"Table"`) {
			t.Errorf("expected type to be included in JSON")
		}
		if !strings.Contains(jsonStr, `"firstRowAsHeader":true`) {
			t.Errorf("expected firstRowAsHeader default to be included")
		}
		if !strings.Contains(jsonStr, `"showGridLines":true`) {
			t.Errorf("expected showGridLines default to be included")
		}
	})

	t.Run("unmarshal basic table", func(t *testing.T) {
		jsonData := `{
			"type": "Table",
			"columns": [
				{"type": "TableColumnDefinition", "width": 1},
				{"type": "TableColumnDefinition", "width": "100px"}
			],
			"rows": [
				{
					"type": "TableRow",
					"cells": [
						{"type": "TableCell", "items": [{"type": "TextBlock", "text": "Header 1"}]},
						{"type": "TableCell", "items": [{"type": "TextBlock", "text": "Header 2"}]}
					]
				}
			]
		}`

		var table Table
		if err := json.Unmarshal([]byte(jsonData), &table); err != nil {
			t.Errorf("unexpected unmarshal error: %v", err)
			return
		}

		if table.Type != m.TypeTable {
			t.Errorf("expected type Table, got %q", table.Type)
		}

		if len(table.Columns) != 2 {
			t.Errorf("expected 2 columns, got %d", len(table.Columns))
		}

		if len(table.Rows) != 1 {
			t.Errorf("expected 1 row, got %d", len(table.Rows))
		}

		// Check defaults are applied
		if table.FirstRowAsHeader == nil || *table.FirstRowAsHeader != true {
			t.Errorf("expected FirstRowAsHeader default to be applied")
		}
	})

	t.Run("roundtrip", func(t *testing.T) {
		original := NewTable().
			AddColumnWithWidth(2).
			AddColumnWithWidth("auto").
			AddTextRow("Name", "Value").
			AddTextRow("Test", "123").
			WithFirstRowAsHeader(true).
			GridStyleAccent()

		// Marshal
		data, err := json.Marshal(original)
		if err != nil {
			t.Errorf("marshal error: %v", err)
			return
		}

		// Unmarshal
		var unmarshaled Table
		if err := json.Unmarshal(data, &unmarshaled); err != nil {
			t.Errorf("unmarshal error: %v", err)
			return
		}

		// Verify
		if len(unmarshaled.Columns) != len(original.Columns) {
			t.Errorf("expected %d columns, got %d", len(original.Columns), len(unmarshaled.Columns))
		}

		if len(unmarshaled.Rows) != len(original.Rows) {
			t.Errorf("expected %d rows, got %d", len(original.Rows), len(unmarshaled.Rows))
		}

		if unmarshaled.GridStyle != original.GridStyle {
			t.Errorf("expected grid style %q, got %q", original.GridStyle, unmarshaled.GridStyle)
		}
	})
}

// TestTableIntegration tests the Table in a complete table scenario
func TestTableIntegration(t *testing.T) {
	// Create a complete table
	table := NewTable().
		AddColumnWithWidth(1).
		AddColumnWithWidth(1).
		AddTextRow("Product", "Price").
		AddTextRow("Widget", "$10.00").
		WithFirstRowAsHeader(true).
		WithShowGridLines(true).
		GridStyleEmphasis().
		WithID("products-table")

	// Validate the complete table
	if err := table.Validate(); err != nil {
		t.Errorf("table validation error: %v", err)
		return
	}

	// Test JSON marshaling
	jsonData, err := json.Marshal(table)
	if err != nil {
		t.Errorf("JSON marshal error: %v", err)
		return
	}

	// Verify the JSON contains table data
	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"type":"Table"`) {
		t.Errorf("expected Table type in JSON")
	}
	if !strings.Contains(jsonStr, `"products-table"`) {
		t.Errorf("expected table ID in JSON")
	}
}
