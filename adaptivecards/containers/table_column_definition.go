package containers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// TableColumnDefinition defines the characteristics of a column in a Table element.
type TableColumnDefinition struct {
	Type                           m.TypeString          `json:"type"`                                     // Version 1.5
	Width                          any                   `json:"width,omitempty"`                          // Version 1.5 - string or number
	HorizontalCellContentAlignment m.HorizontalAlignment `json:"horizontalCellContentAlignment,omitempty"` // Version 1.5
	VerticalCellContentAlignment   m.VerticalAlignment   `json:"verticalCellContentAlignment,omitempty"`   // Version 1.5
}

// NewTableColumnDefinition creates a new TableColumnDefinition with default settings.
func NewTableColumnDefinition() TableColumnDefinition {
	return TableColumnDefinition{
		Type:  m.TypeTableColumnDefinition,
		Width: 1, // Default width is 1
	}
}

// NewTableColumnDefinitionWithWidth creates a new TableColumnDefinition with the specified width.
func NewTableColumnDefinitionWithWidth(width any) TableColumnDefinition {
	return TableColumnDefinition{
		Type:  m.TypeTableColumnDefinition,
		Width: width,
	}
}

// Builder methods for TableColumnDefinition.

// WithWidth sets the width of the column.
// Can be a number (relative weight) or string (pixel value like "50px").
func (tcd TableColumnDefinition) WithWidth(width any) TableColumnDefinition {
	tcd.Width = width
	return tcd
}

// WithPixelWidth sets the width as a pixel value (e.g., "50px").
func (tcd TableColumnDefinition) WithPixelWidth(pixels string) TableColumnDefinition {
	tcd.Width = pixels
	return tcd
}

// WithRelativeWidth sets the width as a relative weight (number).
func (tcd TableColumnDefinition) WithRelativeWidth(weight int) TableColumnDefinition {
	if weight < 1 {
		weight = 1
	}
	tcd.Width = weight
	return tcd
}

// WithHorizontalCellContentAlignment sets the horizontal alignment for all cells in this column.
func (tcd TableColumnDefinition) WithHorizontalCellContentAlignment(
	align m.HorizontalAlignment,
) TableColumnDefinition {
	tcd.HorizontalCellContentAlignment = align
	return tcd
}

// WithVerticalCellContentAlignment sets the vertical alignment for all cells in this column.
func (tcd TableColumnDefinition) WithVerticalCellContentAlignment(
	align m.VerticalAlignment,
) TableColumnDefinition {
	tcd.VerticalCellContentAlignment = align
	return tcd
}

// Convenience methods for horizontal alignment.
func (tcd TableColumnDefinition) AlignLeft() TableColumnDefinition {
	tcd.HorizontalCellContentAlignment = m.HAlignLeft
	return tcd
}

func (tcd TableColumnDefinition) AlignCenter() TableColumnDefinition {
	tcd.HorizontalCellContentAlignment = m.HAlignCenter
	return tcd
}

func (tcd TableColumnDefinition) AlignRight() TableColumnDefinition {
	tcd.HorizontalCellContentAlignment = m.HAlignRight
	return tcd
}

// Convenience methods for vertical alignment.
func (tcd TableColumnDefinition) AlignTop() TableColumnDefinition {
	tcd.VerticalCellContentAlignment = m.VAlignTop
	return tcd
}

func (tcd TableColumnDefinition) AlignMiddle() TableColumnDefinition {
	tcd.VerticalCellContentAlignment = m.VAlignCenter
	return tcd
}

func (tcd TableColumnDefinition) AlignBottom() TableColumnDefinition {
	tcd.VerticalCellContentAlignment = m.VAlignBottom
	return tcd
}

// Validate checks that the TableColumnDefinition has valid fields.
func (tcd TableColumnDefinition) Validate() error {
	// Validate width if provided
	if tcd.Width != nil {
		switch w := tcd.Width.(type) {
		case string:
			if w != "" && !isValidWidthString(w) {
				return fmt.Errorf("TableColumnDefinition.width as string must be \"auto\", \"stretch\", or in format \"<number>px\" (got %q)", w)
			}
		case int:
			if w < 1 {
				return fmt.Errorf("TableColumnDefinition.width as number must be >= 1 (got %d)", w)
			}
		case float64:
			if w < 1 {
				return fmt.Errorf("TableColumnDefinition.width as number must be >= 1 (got %f)", w)
			}
		default:
			return fmt.Errorf("TableColumnDefinition.width must be string or number")
		}
	}

	// Validate enum values
	if tcd.HorizontalCellContentAlignment != "" && !tcd.HorizontalCellContentAlignment.IsValid() {
		return m.NewEnumError(
			"TableColumnDefinition.horizontalCellContentAlignment",
			string(tcd.HorizontalCellContentAlignment),
			m.AllowedHorizontalAlignmentStrings(),
		)
	}
	if tcd.VerticalCellContentAlignment != "" && !tcd.VerticalCellContentAlignment.IsValid() {
		return m.NewEnumError(
			"TableColumnDefinition.verticalCellContentAlignment",
			string(tcd.VerticalCellContentAlignment),
			m.AllowedVerticalAlignmentStrings(),
		)
	}

	return nil
}

// isValidWidthString validates width format (pixel width, auto, or stretch).
func isValidWidthString(width string) bool {
	if width == "auto" || width == "stretch" {
		return true
	}
	if len(width) < 3 {
		return false
	}
	if !strings.HasSuffix(width, "px") {
		return false
	}
	numberPart := width[:len(width)-2]
	_, err := strconv.Atoi(numberPart)
	return err == nil
}

// isValidPixelWidth validates pixel width format.
func isValidPixelWidth(width string) bool {
	return isValidWidthString(width)
}

// MarshalJSON ensures Type is always set.
func (tcd TableColumnDefinition) MarshalJSON() ([]byte, error) {
	tcdd := tcd
	if tcdd.Type == "" {
		tcdd.Type = m.TypeTableColumnDefinition
	}
	// Set default width if not provided
	if tcdd.Width == nil {
		tcdd.Width = 1
	}
	type alias TableColumnDefinition
	return json.Marshal(alias(tcdd))
}

// UnmarshalJSON decodes TableColumnDefinition from JSON.
func (tcd *TableColumnDefinition) UnmarshalJSON(b []byte) error {
	type alias TableColumnDefinition
	var tmp alias
	if err := json.Unmarshal(b, &tmp); err != nil {
		return fmt.Errorf("TableColumnDefinition: decode: %w", err)
	}

	// Validate type
	if tmp.Type != "" && tmp.Type != m.TypeTableColumnDefinition {
		return fmt.Errorf(
			"TableColumnDefinition.type must be %q (got %q)",
			m.TypeTableColumnDefinition,
			tmp.Type,
		)
	}

	// Default the type if omitted
	if tmp.Type == "" {
		tmp.Type = m.TypeTableColumnDefinition
	}

	// Default width if not provided
	if tmp.Width == nil {
		tmp.Width = 1
	}

	// Validate width format during unmarshaling
	if tmp.Width != nil {
		switch w := tmp.Width.(type) {
		case string:
			if w != "" && !isValidWidthString(w) {
				return fmt.Errorf("TableColumnDefinition.width as string must be \"auto\", \"stretch\", or in format \"<number>px\" (got %q)", w)
			}
		case float64: // JSON numbers are unmarshaled as float64
			if w < 1 {
				return fmt.Errorf("TableColumnDefinition.width as number must be >= 1 (got %f)", w)
			}
		default:
			return fmt.Errorf("TableColumnDefinition.width must be string or number")
		}
	}

	// Validate enum values during unmarshaling
	if tmp.HorizontalCellContentAlignment != "" && !tmp.HorizontalCellContentAlignment.IsValid() {
		return m.NewEnumError(
			"TableColumnDefinition.horizontalCellContentAlignment",
			string(tmp.HorizontalCellContentAlignment),
			m.AllowedHorizontalAlignmentStrings(),
		)
	}
	if tmp.VerticalCellContentAlignment != "" && !tmp.VerticalCellContentAlignment.IsValid() {
		return m.NewEnumError(
			"TableColumnDefinition.verticalCellContentAlignment",
			string(tmp.VerticalCellContentAlignment),
			m.AllowedVerticalAlignmentStrings(),
		)
	}

	*tcd = TableColumnDefinition(tmp)
	return nil
}
