package containers

import (
	"encoding/json"
	"fmt"
	"strings"

	e "github.com/untcha/go-adaptivecards/adaptivecards/core/element"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// Table
// Provides a way to display data in a tabular form.
// See: https://adaptivecards.io/explorer/Table.html
type Table struct {
	e.ElementBase                                          // Embedding e.ElementBase to include common element fields
	Type                           m.TypeString            `json:"type"`                                     // Version 1.5
	Columns                        []TableColumnDefinition `json:"columns,omitempty"`                        // Version 1.5
	Rows                           []TableRow              `json:"rows,omitempty"`                           // Version 1.5
	FirstRowAsHeader               *bool                   `json:"firstRowAsHeader,omitempty"`               // Version 1.5
	ShowGridLines                  *bool                   `json:"showGridLines,omitempty"`                  // Version 1.5
	GridStyle                      m.ContainerStyle        `json:"gridStyle,omitempty"`                      // Version 1.5
	HorizontalCellContentAlignment m.HorizontalAlignment   `json:"horizontalCellContentAlignment,omitempty"` // Version 1.5
	VerticalCellContentAlignment   m.VerticalAlignment     `json:"verticalCellContentAlignment,omitempty"`   // Version 1.5
}

// NewTable creates a new empty Table
func NewTable() Table {
	firstRowAsHeader := true // Default value per spec
	showGridLines := true    // Default value per spec
	return Table{
		ElementBase:      e.ElementBase{},
		Type:             m.TypeTable,
		Columns:          make([]TableColumnDefinition, 0),
		Rows:             make([]TableRow, 0),
		FirstRowAsHeader: &firstRowAsHeader,
		ShowGridLines:    &showGridLines,
		GridStyle:        m.ContainerStyleDefault,
	}
}

// NewTableWithColumnsAndRows creates a new Table with specified columns and rows
func NewTableWithColumnsAndRows(columns []TableColumnDefinition, rows []TableRow) Table {
	firstRowAsHeader := true // Default value per spec
	showGridLines := true    // Default value per spec
	return Table{
		ElementBase:      e.ElementBase{},
		Type:             m.TypeTable,
		Columns:          columns,
		Rows:             rows,
		FirstRowAsHeader: &firstRowAsHeader,
		ShowGridLines:    &showGridLines,
		GridStyle:        m.ContainerStyleDefault,
	}
}

func (t Table) GetType() m.TypeString { return m.TypeTable }

// Builder methods for Table structure

// AddColumn adds a column definition to the table
func (t Table) AddColumn(column TableColumnDefinition) Table {
	if column.Type == "" {
		column.Type = m.TypeTableColumnDefinition
	}
	t.Columns = append(t.Columns, column)
	return t
}

// AddColumnWithWidth adds a column with the specified width
func (t Table) AddColumnWithWidth(width any) Table {
	column := NewTableColumnDefinitionWithWidth(width)
	t.Columns = append(t.Columns, column)
	return t
}

// AddRow adds a row to the table
func (t Table) AddRow(row TableRow) Table {
	if row.Type == "" {
		row.Type = m.TypeTableRow
	}
	t.Rows = append(t.Rows, row)
	return t
}

// AddTextRow adds a row with text cells
func (t Table) AddTextRow(texts ...string) Table {
	row := NewTableRowEmpty()
	for _, text := range texts {
		row = row.AddTextCell(text)
	}
	t.Rows = append(t.Rows, row)
	return t
}

// Configuration methods

// WithFirstRowAsHeader sets whether the first row should be treated as a header
func (t Table) WithFirstRowAsHeader(firstRowAsHeader bool) Table {
	t.FirstRowAsHeader = &firstRowAsHeader
	return t
}

// WithShowGridLines sets whether grid lines should be displayed
func (t Table) WithShowGridLines(showGridLines bool) Table {
	t.ShowGridLines = &showGridLines
	return t
}

// WithGridStyle sets the grid style
func (t Table) WithGridStyle(style m.ContainerStyle) Table {
	t.GridStyle = style
	return t
}

// WithHorizontalCellContentAlignment sets the default horizontal alignment for all cells
func (t Table) WithHorizontalCellContentAlignment(align m.HorizontalAlignment) Table {
	t.HorizontalCellContentAlignment = align
	return t
}

// WithVerticalCellContentAlignment sets the default vertical alignment for all cells
func (t Table) WithVerticalCellContentAlignment(align m.VerticalAlignment) Table {
	t.VerticalCellContentAlignment = align
	return t
}

// Convenience methods for grid style
func (t Table) GridStyleDefault() Table   { t.GridStyle = m.ContainerStyleDefault; return t }
func (t Table) GridStyleEmphasis() Table  { t.GridStyle = m.ContainerStyleEmphasis; return t }
func (t Table) GridStyleGood() Table      { t.GridStyle = m.ContainerStyleGood; return t }
func (t Table) GridStyleAttention() Table { t.GridStyle = m.ContainerStyleAttention; return t }
func (t Table) GridStyleWarning() Table   { t.GridStyle = m.ContainerStyleWarning; return t }
func (t Table) GridStyleAccent() Table    { t.GridStyle = m.ContainerStyleAccent; return t }

// Convenience methods for horizontal alignment
func (t Table) CellsAlignLeft() Table { t.HorizontalCellContentAlignment = m.HAlignLeft; return t }
func (t Table) CellsAlignCenter() Table {
	t.HorizontalCellContentAlignment = m.HAlignCenter
	return t
}

func (t Table) CellsAlignRight() Table { t.HorizontalCellContentAlignment = m.HAlignRight; return t }

// Convenience methods for vertical alignment
func (t Table) CellsAlignTop() Table { t.VerticalCellContentAlignment = m.VAlignTop; return t }

func (t Table) CellsAlignMiddle() Table { t.VerticalCellContentAlignment = m.VAlignCenter; return t }

func (t Table) CellsAlignBottom() Table { t.VerticalCellContentAlignment = m.VAlignBottom; return t }

// Fluent setters for embedded e.ElementBase fields

// WithFallback sets the Fallback field
func (t Table) WithFallback(fallback *e.ElementFallback) Table {
	t.Fallback = fallback
	return t
}

// WithHeight sets the Height field
func (t Table) WithHeight(height m.BlockElementHeight) Table {
	t.Height = height
	return t
}

// WithSeparator sets the Separator field
func (t Table) WithSeparator(separator bool) Table {
	t.Separator = separator
	return t
}

// WithSpacing sets the m.Spacing field
func (t Table) WithSpacing(spacing m.Spacing) Table { t.Spacing = spacing; return t }

// WithID sets the ID field
func (t Table) WithID(id string) Table { t.ID = id; return t }

// WithVisible sets the IsVisible field
func (t Table) WithVisible(visible bool) Table { t.IsVisible = &visible; return t }

// WithRequires sets the Requires field
func (t Table) WithRequires(requires map[string]string) Table {
	t.Requires = requires
	return t
}

// Convenience methods for common spacing values
func (t Table) SpacingDefault() Table    { t.Spacing = m.SpacingDefault; return t }
func (t Table) SpacingNone() Table       { t.Spacing = m.SpacingNone; return t }
func (t Table) SpacingSmall() Table      { t.Spacing = m.SpacingSmall; return t }
func (t Table) SpacingMedium() Table     { t.Spacing = m.SpacingMedium; return t }
func (t Table) SpacingLarge() Table      { t.Spacing = m.SpacingLarge; return t }
func (t Table) SpacingExtraLarge() Table { t.Spacing = m.SpacingExtraLarge; return t }
func (t Table) SpacingPadding() Table    { t.Spacing = m.SpacingPadding; return t }

// Convenience methods for visibility
func (t Table) Hide() Table { visible := false; t.IsVisible = &visible; return t }
func (t Table) Show() Table { visible := true; t.IsVisible = &visible; return t }

// Validate checks that the Table has valid fields
func (t Table) Validate() error {
	// Validate embedded e.ElementBase fields
	if err := t.validateElementBase(); err != nil {
		return err
	}

	// Validate all columns
	for i, column := range t.Columns {
		if err := column.Validate(); err != nil {
			return fmt.Errorf("table.columns[%d]: %w", i, err)
		}
	}

	// Validate all rows
	for i, row := range t.Rows {
		if err := row.Validate(); err != nil {
			return fmt.Errorf("table.rows[%d]: %w", i, err)
		}

		// Validate that row doesn't have more cells than columns
		if len(t.Columns) > 0 && len(row.Cells) > len(t.Columns) {
			return fmt.Errorf(
				"table.rows[%d]: row has %d cells but table has only %d columns",
				i,
				len(row.Cells),
				len(t.Columns),
			)
		}
	}

	// Validate enum values
	if t.GridStyle != "" && !t.GridStyle.IsValid() {
		return m.NewEnumError(
			"Table.gridStyle",
			string(t.GridStyle),
			m.AllowedContainerStyleStrings(),
		)
	}

	if t.HorizontalCellContentAlignment != "" && !t.HorizontalCellContentAlignment.IsValid() {
		return m.NewEnumError(
			"Table.horizontalCellContentAlignment",
			string(t.HorizontalCellContentAlignment),
			m.AllowedHorizontalAlignmentStrings(),
		)
	}

	if t.VerticalCellContentAlignment != "" && !t.VerticalCellContentAlignment.IsValid() {
		return m.NewEnumError(
			"Table.verticalCellContentAlignment",
			string(t.VerticalCellContentAlignment),
			m.AllowedVerticalAlignmentStrings(),
		)
	}

	return nil
}

// Custom JSON marshalling to ensure Type is always set and defaults are applied
func (t Table) MarshalJSON() ([]byte, error) {
	tt := t
	if tt.Type == "" {
		tt.Type = m.TypeTable
	}
	// Set defaults if not specified
	if tt.FirstRowAsHeader == nil {
		firstRowAsHeader := true
		tt.FirstRowAsHeader = &firstRowAsHeader
	}
	if tt.ShowGridLines == nil {
		showGridLines := true
		tt.ShowGridLines = &showGridLines
	}
	if tt.GridStyle == "" {
		tt.GridStyle = m.ContainerStyleDefault
	}

	type alias Table
	return json.Marshal(alias(tt))
}

// Custom JSON unmarshalling
func (t *Table) UnmarshalJSON(b []byte) error {
	type alias Table
	var tmp alias
	if err := json.Unmarshal(b, &tmp); err != nil {
		return fmt.Errorf("Table: decode: %w", err)
	}

	// Validate type
	if tmp.Type != "" && tmp.Type != m.TypeTable {
		return fmt.Errorf("table.type must be %q (got %q)", m.TypeTable, tmp.Type)
	}

	// Default the type if omitted
	if tmp.Type == "" {
		tmp.Type = m.TypeTable
	}

	// Set defaults if not provided
	if tmp.FirstRowAsHeader == nil {
		firstRowAsHeader := true
		tmp.FirstRowAsHeader = &firstRowAsHeader
	}
	if tmp.ShowGridLines == nil {
		showGridLines := true
		tmp.ShowGridLines = &showGridLines
	}
	if tmp.GridStyle == "" {
		tmp.GridStyle = m.ContainerStyleDefault
	}

	// Validate enum values during unmarshaling
	if tmp.GridStyle != "" && !tmp.GridStyle.IsValid() {
		return m.NewEnumError(
			"Table.gridStyle",
			string(tmp.GridStyle),
			m.AllowedContainerStyleStrings(),
		)
	}

	if tmp.HorizontalCellContentAlignment != "" && !tmp.HorizontalCellContentAlignment.IsValid() {
		return m.NewEnumError(
			"Table.horizontalCellContentAlignment",
			string(tmp.HorizontalCellContentAlignment),
			m.AllowedHorizontalAlignmentStrings(),
		)
	}

	if tmp.VerticalCellContentAlignment != "" && !tmp.VerticalCellContentAlignment.IsValid() {
		return m.NewEnumError(
			"Table.verticalCellContentAlignment",
			string(tmp.VerticalCellContentAlignment),
			m.AllowedVerticalAlignmentStrings(),
		)
	}

	*t = Table(tmp)
	return nil
}

// validateElementBase validates the inherited e.ElementBase fields
func (t Table) validateElementBase() error {
	// Validate Height field
	if t.Height != "" && !t.Height.IsValid() {
		return m.NewEnumError(
			"Table.height",
			string(t.Height),
			m.AllowedBlockElementHeightStrings(),
		)
	}

	// Validate m.Spacing enum
	if t.Spacing != "" && !t.Spacing.IsValid() {
		return m.NewEnumError("Table.spacing", string(t.Spacing), m.AllowedSpacingStrings())
	}

	// Validate ID field (basic sanity checks)
	if t.ID != "" {
		id := strings.TrimSpace(t.ID)
		if id == "" {
			return fmt.Errorf("table.id cannot be empty or whitespace-only")
		}
		// ID should not contain certain characters that could cause issues
		if strings.ContainsAny(id, "\n\r\t") {
			return fmt.Errorf("table.id cannot contain newlines or tabs")
		}
	}

	return nil
}

// Register Table in the element registry
func init() {
	e.RegisterElement(m.TypeTable, func() e.Element { return &Table{Type: m.TypeTable} })
}
