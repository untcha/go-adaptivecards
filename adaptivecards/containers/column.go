package containers

import (
	"encoding/json"
	"fmt"
	"strings"

	a "github.com/untcha/go-adaptivecards/adaptivecards/actions"
	e "github.com/untcha/go-adaptivecards/adaptivecards/core/element"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// Column
// Defines a container that is part of a ColumnSet.
// See: https://adaptivecards.io/explorer/Column.html
type Column struct {
	e.ElementBase                                       // Embedding e.ElementBase to include common element fields
	Type                     m.TypeString               `json:"type"`                               // Version 1.0
	Items                    []e.Element                `json:"items,omitempty"`                    // Version 1.0
	BackgroundImage          *m.BackgroundImageValue    `json:"backgroundImage,omitempty"`          // Version 1.2
	Bleed                    bool                       `json:"bleed,omitempty"`                    // Version 1.2
	MinHeight                string                     `json:"minHeight,omitempty"`                // Version 1.2
	Rtl                      *bool                      `json:"rtl,omitempty"`                      // Version 1.5
	SelectAction             a.Action                   `json:"selectAction,omitempty"`             // Version 1.1
	Style                    m.ContainerStyle           `json:"style,omitempty"`                    // Version 1.2
	VerticalContentAlignment m.VerticalContentAlignment `json:"verticalContentAlignment,omitempty"` // Version 1.1
	Width                    any                        `json:"width,omitempty"`                    // Version 1.0
}

func NewColumn(items ...e.Element) Column {
	return Column{
		Type:  m.TypeColumn,
		Items: items,
	}
}

func (c Column) GetType() m.TypeString { return m.TypeColumn }

func (c Column) Validate() error {
	if err := c.validateElementBase(); err != nil {
		return err
	}
	if err := validateElements(c.Items); err != nil {
		return fmt.Errorf("column.items: %w", err)
	}
	if c.SelectAction != nil {
		if err := validateSelectAction(c.SelectAction); err != nil {
			return fmt.Errorf("column.selectAction: %w", err)
		}
	}
	if c.Style != "" && !c.Style.IsValid() {
		return m.NewEnumError("Column.style", string(c.Style), m.AllowedContainerStyleStrings())
	}
	if c.VerticalContentAlignment != "" && !c.VerticalContentAlignment.IsValid() {
		return m.NewEnumError(
			"Column.verticalContentAlignment",
			string(c.VerticalContentAlignment),
			m.AllowedVerticalContentAlignmentStrings(),
		)
	}
	if c.BackgroundImage != nil {
		if c.BackgroundImage.Object() != nil {
			if err := c.BackgroundImage.Object().Validate(); err != nil {
				return fmt.Errorf("column.backgroundImage: %w", err)
			}
		} else if c.BackgroundImage.URL() != "" {
			if _, err := m.BackgroundImageURL(c.BackgroundImage.URL()); err != nil {
				return fmt.Errorf("column.backgroundImage: %w", err)
			}
		}
	}
	if c.MinHeight != "" && !isValidPixelWidth(c.MinHeight) {
		return fmt.Errorf("column.minHeight must be in format \"<number>px\" (got %q)", c.MinHeight)
	}
	if c.Width != nil {
		switch w := c.Width.(type) {
		case string:
			if w != "" && !isValidWidthString(w) {
				return fmt.Errorf("column.width as string must be \"auto\", \"stretch\", or in format \"<number>px\" (got %q)", w)
			}
		case int:
			if w < 1 {
				return fmt.Errorf("column.width as number must be >= 1 (got %d)", w)
			}
		case float64:
			if w < 1 {
				return fmt.Errorf("column.width as number must be >= 1 (got %f)", w)
			}
		default:
			return fmt.Errorf("column.width must be string or number")
		}
	}
	return nil
}

func (c Column) MarshalJSON() ([]byte, error) {
	cc := c
	if cc.Type == "" {
		cc.Type = m.TypeColumn
	}
	type alias Column
	return json.Marshal(alias(cc))
}

func (c *Column) UnmarshalJSON(b []byte) error {
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(b, &obj); err != nil {
		return fmt.Errorf("column: decode: %w", err)
	}
	if v, ok := obj["type"]; ok {
		var t string
		if err := json.Unmarshal(v, &t); err != nil {
			return fmt.Errorf("column.type: %w", err)
		}
		if t != "" && t != string(m.TypeColumn) {
			return fmt.Errorf("column.type must be %q (got %q)", m.TypeColumn, t)
		}
	}
	itemsRaw := obj["items"]
	selectActionRaw := obj["selectAction"]
	delete(obj, "items")
	delete(obj, "selectAction")

	type alias Column
	var base alias
	baseBytes, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("column: re-encode base: %w", err)
	}
	if err := json.Unmarshal(baseBytes, &base); err != nil {
		return fmt.Errorf("column: decode base: %w", err)
	}
	if base.Type == "" {
		base.Type = m.TypeColumn
	}

	if len(itemsRaw) != 0 {
		var rawArr []json.RawMessage
		if err := json.Unmarshal(itemsRaw, &rawArr); err != nil {
			return fmt.Errorf("column.items: %w", err)
		}
		items, err := e.UnmarshalElementsSlice(rawArr)
		if err != nil {
			return fmt.Errorf("column.items: %w", err)
		}
		base.Items = items
	}
	if len(selectActionRaw) != 0 {
		action, err := a.UnmarshalAction(selectActionRaw)
		if err != nil {
			return fmt.Errorf("column.selectAction: %w", err)
		}
		base.SelectAction = action
	}

	*c = Column(base)
	return nil
}

func (c Column) validateElementBase() error {
	if c.Spacing != "" && !c.Spacing.IsValid() {
		return m.NewEnumError("Column.spacing", string(c.Spacing), m.AllowedSpacingStrings())
	}
	if c.ID != "" {
		id := strings.TrimSpace(c.ID)
		if id == "" {
			return fmt.Errorf("column.id cannot be empty or whitespace-only")
		}
		if strings.ContainsAny(id, "\n\r\t") {
			return fmt.Errorf("column.id cannot contain newlines or tabs")
		}
	}
	return nil
}

func init() {
	e.RegisterElement(m.TypeColumn, func() e.Element { return &Column{Type: m.TypeColumn} })
}
