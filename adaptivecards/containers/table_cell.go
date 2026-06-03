package containers

import (
	"encoding/json"
	"fmt"

	a "github.com/untcha/go-adaptivecards/adaptivecards/actions"
	e "github.com/untcha/go-adaptivecards/adaptivecards/core/element"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
	els "github.com/untcha/go-adaptivecards/adaptivecards/elements"
)

// TableCell represents a cell within a row of a Table element.
type TableCell struct {
	Type                     m.TypeString               `json:"type"`                               // Version 1.5
	Items                    []e.Element                `json:"items"`                              // Version 1.5
	SelectAction             a.Action                   `json:"selectAction,omitempty"`             // Version 1.1
	Style                    m.ContainerStyle           `json:"style,omitempty"`                    // Version 1.5
	VerticalContentAlignment m.VerticalContentAlignment `json:"verticalContentAlignment,omitempty"` // Version 1.1
	Bleed                    bool                       `json:"bleed,omitempty"`                    // Version 1.2
	BackgroundImage          *m.BackgroundImageValue    `json:"backgroundImage,omitempty"`          // Version 1.2
	MinHeight                string                     `json:"minHeight,omitempty"`                // Version 1.2
	Rtl                      *bool                      `json:"rtl,omitempty"`                      // Version 1.5
}

// NewTableCell creates a new TableCell with the specified items.
func NewTableCell(items ...e.Element) TableCell {
	return TableCell{
		Type:  m.TypeTableCell,
		Items: items,
	}
}

// NewTableCellEmpty creates a new empty TableCell.
func NewTableCellEmpty() TableCell {
	return TableCell{
		Type:  m.TypeTableCell,
		Items: make([]e.Element, 0),
	}
}

// GetType returns the Adaptive Card type discriminator for TableCell.
func (tc TableCell) GetType() m.TypeString { return m.TypeTableCell }

// Builder methods for TableCell.

// AddElement adds an element to the cell.
func (tc TableCell) AddElement(el e.Element) TableCell {
	tc.Items = append(tc.Items, el)
	return tc
}

// AddTextBlock adds a TextBlock to the cell.
func (tc TableCell) AddTextBlock(tb els.TextBlock) TableCell {
	if tb.Type == "" {
		tb.Type = m.TypeTextBlock
	}
	tc.Items = append(tc.Items, tb)
	return tc
}

// AddImage adds an Image to the cell.
func (tc TableCell) AddImage(img els.Image) TableCell {
	if img.Type == "" {
		img.Type = m.TypeImage
	}
	tc.Items = append(tc.Items, img)
	return tc
}

// WithSelectAction sets the SelectAction field.
func (tc TableCell) WithSelectAction(action a.Action) TableCell {
	tc.SelectAction = action
	return tc
}

// WithStyle sets the Style field.
func (tc TableCell) WithStyle(style m.ContainerStyle) TableCell {
	tc.Style = style
	return tc
}

// WithVerticalContentAlignment sets the VerticalContentAlignment field.
func (tc TableCell) WithVerticalContentAlignment(align m.VerticalContentAlignment) TableCell {
	tc.VerticalContentAlignment = align
	return tc
}

// WithBleed sets the Bleed field.
func (tc TableCell) WithBleed(bleed bool) TableCell {
	tc.Bleed = bleed
	return tc
}

// WithBackgroundImage sets BackgroundImage from a BackgroundImage object.
func (tc TableCell) WithBackgroundImage(bg m.BackgroundImage) TableCell {
	// Preserve provided state even if invalid; Validate() reports concrete errors later.
	tc.BackgroundImage = m.BackgroundImageObjectUnchecked(bg)
	return tc
}

// WithBackgroundImageURL sets BackgroundImage from a URL string.
func (tc TableCell) WithBackgroundImageURL(url string) TableCell {
	// Preserve provided state even if invalid; Validate() reports concrete errors later.
	tc.BackgroundImage = m.BackgroundImageURLUnchecked(url)
	return tc
}

// WithMinHeight sets the MinHeight field.
func (tc TableCell) WithMinHeight(minHeight string) TableCell {
	tc.MinHeight = minHeight
	return tc
}

// WithRtl sets the Rtl field.
func (tc TableCell) WithRtl(rtl bool) TableCell {
	tc.Rtl = &rtl
	return tc
}

// Convenience methods for style.
func (tc TableCell) StyleDefault() TableCell   { tc.Style = m.ContainerStyleDefault; return tc }
func (tc TableCell) StyleEmphasis() TableCell  { tc.Style = m.ContainerStyleEmphasis; return tc }
func (tc TableCell) StyleGood() TableCell      { tc.Style = m.ContainerStyleGood; return tc }
func (tc TableCell) StyleAttention() TableCell { tc.Style = m.ContainerStyleAttention; return tc }
func (tc TableCell) StyleWarning() TableCell   { tc.Style = m.ContainerStyleWarning; return tc }
func (tc TableCell) StyleAccent() TableCell    { tc.Style = m.ContainerStyleAccent; return tc }

// Convenience methods for vertical alignment.
func (tc TableCell) AlignTop() TableCell { tc.VerticalContentAlignment = m.VContentAlignTop; return tc }
func (tc TableCell) AlignCenter() TableCell {
	tc.VerticalContentAlignment = m.VContentAlignCenter
	return tc
}
func (tc TableCell) AlignBottom() TableCell {
	tc.VerticalContentAlignment = m.VContentAlignBottom
	return tc
}

// Validate checks that the TableCell has valid fields.
func (tc TableCell) Validate() error {
	// Items is required
	if len(tc.Items) == 0 {
		return fmt.Errorf("TableCell.items is required and cannot be empty")
	}

	// Validate all items
	for i, item := range tc.Items {
		if err := validateElement(item); err != nil {
			return fmt.Errorf("TableCell.items[%d]: %w", i, err)
		}
	}

	// Validate SelectAction if present (Action.ShowCard is not supported).
	if tc.SelectAction != nil {
		if tc.SelectAction.GetType() == m.TypeActionShowCard {
			return fmt.Errorf("TableCell.selectAction: Action.ShowCard is not supported")
		}
		if err := validateAction(tc.SelectAction); err != nil {
			return fmt.Errorf("TableCell.selectAction: %w", err)
		}
	}

	// Validate enum values.
	if tc.Style != "" {
		if !tc.Style.IsValid() {
			return m.NewEnumError(
				"TableCell.style",
				string(tc.Style),
				m.AllowedContainerStyleStrings(),
			)
		}
	}

	if tc.VerticalContentAlignment != "" && !tc.VerticalContentAlignment.IsValid() {
		return m.NewEnumError(
			"TableCell.verticalContentAlignment",
			string(tc.VerticalContentAlignment),
			m.AllowedVerticalContentAlignmentStrings(),
		)
	}

	// Validate background image if present.
	if tc.BackgroundImage != nil {
		if tc.BackgroundImage.Object() != nil {
			if err := tc.BackgroundImage.Object().Validate(); err != nil {
				return fmt.Errorf("TableCell.backgroundImage: %w", err)
			}
		} else if tc.BackgroundImage.URL() != "" {
			if _, err := m.BackgroundImageURL(tc.BackgroundImage.URL()); err != nil {
				return fmt.Errorf("TableCell.backgroundImage: %w", err)
			}
		}
	}

	// Validate MinHeight format if provided
	if tc.MinHeight != "" && !isValidPixelWidth(tc.MinHeight) {
		return fmt.Errorf(
			"TableCell.minHeight must be in format \"<number>px\" (got %q)",
			tc.MinHeight,
		)
	}

	return nil
}

// MarshalJSON ensures Type is always set.
func (tc TableCell) MarshalJSON() ([]byte, error) {
	tcc := tc
	if tcc.Type == "" {
		tcc.Type = m.TypeTableCell
	}
	type alias TableCell
	return json.Marshal(alias(tcc))
}

// UnmarshalJSON decodes TableCell and its interface fields.
func (tc *TableCell) UnmarshalJSON(b []byte) error {
	// Decode into a generic map to handle interface fields explicitly
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(b, &obj); err != nil {
		return fmt.Errorf("TableCell: decode: %w", err)
	}

	// Validate/normalize type
	if v, ok := obj["type"]; ok {
		var t string
		if err := json.Unmarshal(v, &t); err != nil {
			return fmt.Errorf("TableCell.type: %w", err)
		}
		if t != "" && t != string(m.TypeTableCell) {
			return fmt.Errorf("TableCell.type must be %q (got %q)", m.TypeTableCell, t)
		}
	}

	// Extract raw arrays/objects for interface fields and remove them before base decode
	itemsRaw := obj["items"]
	selectActionRaw := obj["selectAction"]
	delete(obj, "items")
	delete(obj, "selectAction")

	// Decode the remaining fields into base
	type alias TableCell
	var base alias
	baseBytes, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("TableCell: re-encode base: %w", err)
	}
	if err := json.Unmarshal(baseBytes, &base); err != nil {
		return fmt.Errorf("TableCell: decode base: %w", err)
	}

	// Default the type if omitted
	if base.Type == "" {
		base.Type = m.TypeTableCell
	}

	// Decode items (required)
	if len(itemsRaw) == 0 {
		return fmt.Errorf("TableCell.items is required")
	}
	var rawArr []json.RawMessage
	if err := json.Unmarshal(itemsRaw, &rawArr); err != nil {
		return fmt.Errorf("TableCell.items: %w", err)
	}
	items, err := e.UnmarshalElementsSlice(rawArr)
	if err != nil {
		return fmt.Errorf("TableCell.items: %w", err)
	}
	base.Items = items

	// Decode selectAction if present
	if len(selectActionRaw) != 0 {
		action, err := a.UnmarshalAction(selectActionRaw)
		if err != nil {
			return fmt.Errorf("TableCell.selectAction: %w", err)
		}
		base.SelectAction = action
	}

	*tc = TableCell(base)
	return nil
}

// Register TableCell in the element registry.
func init() {
	e.RegisterElement(
		m.TypeTableCell,
		func() e.Element { return &TableCell{Type: m.TypeTableCell} },
	)
}
