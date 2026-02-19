package containers

import (
	"encoding/json"
	"fmt"

	e "github.com/untcha/go-adaptivecards/adaptivecards/core/element"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
	els "github.com/untcha/go-adaptivecards/adaptivecards/elements"
)

// TableRow
// Represents a row of cells within a Table element.
type TableRow struct {
	Type                           m.TypeString          `json:"type"`                                     // Version 1.5
	Cells                          []TableCell           `json:"cells,omitempty"`                          // Version 1.5
	Style                          m.ContainerStyle      `json:"style,omitempty"`                          // Version 1.5
	HorizontalCellContentAlignment m.HorizontalAlignment `json:"horizontalCellContentAlignment,omitempty"` // Version 1.5
	VerticalCellContentAlignment   m.VerticalAlignment   `json:"verticalCellContentAlignment,omitempty"`   // Version 1.5
}

// NewTableRow creates a new TableRow with the specified cells.
func NewTableRow(cells ...TableCell) TableRow {
	return TableRow{
		Type:  m.TypeTableRow,
		Cells: cells,
	}
}

// NewTableRowEmpty creates a new empty TableRow.
func NewTableRowEmpty() TableRow {
	return TableRow{
		Type:  m.TypeTableRow,
		Cells: make([]TableCell, 0),
	}
}

func (tr TableRow) GetType() m.TypeString { return m.TypeTableRow }

// Builder methods for TableRow.

// AddCell adds a cell to the row.
func (tr TableRow) AddCell(cell TableCell) TableRow {
	if cell.Type == "" {
		cell.Type = m.TypeTableCell
	}
	tr.Cells = append(tr.Cells, cell)
	return tr
}

// AddTextCell adds a cell containing a single TextBlock.
func (tr TableRow) AddTextCell(text string) TableRow {
	tb := els.NewTextBlock(text)
	cell := NewTableCell(tb)
	tr.Cells = append(tr.Cells, cell)
	return tr
}

// AddImageCell adds a cell containing a single Image.
func (tr TableRow) AddImageCell(imageURL string) TableRow {
	img := els.NewImage(imageURL)
	cell := NewTableCell(img)
	tr.Cells = append(tr.Cells, cell)
	return tr
}

// AddEmptyCell adds an empty cell to the row.
func (tr TableRow) AddEmptyCell() TableRow {
	cell := NewTableCellEmpty()
	tr.Cells = append(tr.Cells, cell)
	return tr
}

// WithStyle sets the Style field.
func (tr TableRow) WithStyle(style m.ContainerStyle) TableRow {
	tr.Style = style
	return tr
}

// WithHorizontalCellContentAlignment sets the HorizontalCellContentAlignment field.
func (tr TableRow) WithHorizontalCellContentAlignment(align m.HorizontalAlignment) TableRow {
	tr.HorizontalCellContentAlignment = align
	return tr
}

// WithVerticalCellContentAlignment sets the VerticalCellContentAlignment field.
func (tr TableRow) WithVerticalCellContentAlignment(align m.VerticalAlignment) TableRow {
	tr.VerticalCellContentAlignment = align
	return tr
}

// Convenience methods for style.
func (tr TableRow) StyleDefault() TableRow   { tr.Style = m.ContainerStyleDefault; return tr }
func (tr TableRow) StyleEmphasis() TableRow  { tr.Style = m.ContainerStyleEmphasis; return tr }
func (tr TableRow) StyleGood() TableRow      { tr.Style = m.ContainerStyleGood; return tr }
func (tr TableRow) StyleAttention() TableRow { tr.Style = m.ContainerStyleAttention; return tr }
func (tr TableRow) StyleWarning() TableRow   { tr.Style = m.ContainerStyleWarning; return tr }
func (tr TableRow) StyleAccent() TableRow    { tr.Style = m.ContainerStyleAccent; return tr }

// Convenience methods for horizontal alignment.
func (tr TableRow) AlignLeft() TableRow { tr.HorizontalCellContentAlignment = m.HAlignLeft; return tr }
func (tr TableRow) AlignCenter() TableRow {
	tr.HorizontalCellContentAlignment = m.HAlignCenter
	return tr
}
func (tr TableRow) AlignRight() TableRow {
	tr.HorizontalCellContentAlignment = m.HAlignRight
	return tr
}

// Convenience methods for vertical alignment.
func (tr TableRow) AlignTop() TableRow { tr.VerticalCellContentAlignment = m.VAlignTop; return tr }
func (tr TableRow) AlignMiddle() TableRow {
	tr.VerticalCellContentAlignment = m.VAlignCenter
	return tr
}
func (tr TableRow) AlignBottom() TableRow {
	tr.VerticalCellContentAlignment = m.VAlignBottom
	return tr
}

// Validate checks that the TableRow has valid fields.
func (tr TableRow) Validate() error {
	// Validate all cells
	for i, cell := range tr.Cells {
		if err := cell.Validate(); err != nil {
			return fmt.Errorf("TableRow.cells[%d]: %w", i, err)
		}
	}

	// Validate enum values.
	if tr.Style != "" {
		if !tr.Style.IsValid() {
			return m.NewEnumError(
				"TableRow.style",
				string(tr.Style),
				m.AllowedContainerStyleStrings(),
			)
		}
	}

	if tr.HorizontalCellContentAlignment != "" && !tr.HorizontalCellContentAlignment.IsValid() {
		return m.NewEnumError(
			"TableRow.horizontalCellContentAlignment",
			string(tr.HorizontalCellContentAlignment),
			m.AllowedHorizontalAlignmentStrings(),
		)
	}

	if tr.VerticalCellContentAlignment != "" && !tr.VerticalCellContentAlignment.IsValid() {
		return m.NewEnumError(
			"TableRow.verticalCellContentAlignment",
			string(tr.VerticalCellContentAlignment),
			m.AllowedVerticalAlignmentStrings(),
		)
	}

	return nil
}

// MarshalJSON ensures Type is always set.
func (tr TableRow) MarshalJSON() ([]byte, error) {
	trr := tr
	if trr.Type == "" {
		trr.Type = m.TypeTableRow
	}
	type alias TableRow
	return json.Marshal(alias(trr))
}

// UnmarshalJSON decodes TableRow from JSON.
func (tr *TableRow) UnmarshalJSON(b []byte) error {
	type alias TableRow
	var tmp alias
	if err := json.Unmarshal(b, &tmp); err != nil {
		return fmt.Errorf("TableRow: decode: %w", err)
	}

	// Validate type
	if tmp.Type != "" && tmp.Type != m.TypeTableRow {
		return fmt.Errorf("TableRow.type must be %q (got %q)", m.TypeTableRow, tmp.Type)
	}

	// Default the type if omitted
	if tmp.Type == "" {
		tmp.Type = m.TypeTableRow
	}

	// Validate enum values during unmarshaling.
	if tmp.Style != "" {
		if !tmp.Style.IsValid() {
			return m.NewEnumError(
				"TableRow.style",
				string(tmp.Style),
				m.AllowedContainerStyleStrings(),
			)
		}
	}

	if tmp.HorizontalCellContentAlignment != "" && !tmp.HorizontalCellContentAlignment.IsValid() {
		return m.NewEnumError(
			"TableRow.horizontalCellContentAlignment",
			string(tmp.HorizontalCellContentAlignment),
			m.AllowedHorizontalAlignmentStrings(),
		)
	}

	if tmp.VerticalCellContentAlignment != "" && !tmp.VerticalCellContentAlignment.IsValid() {
		return m.NewEnumError(
			"TableRow.verticalCellContentAlignment",
			string(tmp.VerticalCellContentAlignment),
			m.AllowedVerticalAlignmentStrings(),
		)
	}

	*tr = TableRow(tmp)
	return nil
}

// Register TableRow in the element registry.
func init() {
	e.RegisterElement(m.TypeTableRow, func() e.Element { return &TableRow{Type: m.TypeTableRow} })
}
