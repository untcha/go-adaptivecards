package containers

import (
	"encoding/json"
	"fmt"

	a "github.com/untcha/go-adaptivecards/adaptivecards/actions"
	e "github.com/untcha/go-adaptivecards/adaptivecards/core/element"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// ColumnSet
// ColumnSet divides a region into Columns, allowing elements to sit side-by-side.
// See: https://adaptivecards.io/explorer/ColumnSet.html
type ColumnSet struct {
	e.ElementBase                             // Embedding e.ElementBase to include common element fields
	Type                m.TypeString          `json:"type"`                          // Version 1.0
	Columns             []Column              `json:"columns,omitempty"`             // Version 1.0
	SelectAction        a.Action              `json:"selectAction,omitempty"`        // Version 1.1
	Style               m.ContainerStyle      `json:"style,omitempty"`               // Version 1.2
	Bleed               bool                  `json:"bleed,omitempty"`               // Version 1.2
	MinHeight           string                `json:"minHeight,omitempty"`           // Version 1.2
	HorizontalAlignment m.HorizontalAlignment `json:"horizontalAlignment,omitempty"` // Version 1.0
}

func NewColumnSet(columns ...Column) ColumnSet {
	return ColumnSet{
		Type:    m.TypeColumnSet,
		Columns: columns,
	}
}

func (c ColumnSet) GetType() m.TypeString { return m.TypeColumnSet }

func (c ColumnSet) Validate() error {
	if err := c.validateElementBase(); err != nil {
		return err
	}
	for i, col := range c.Columns {
		if err := col.Validate(); err != nil {
			return fmt.Errorf("columnSet.columns[%d]: %w", i, err)
		}
	}
	if c.SelectAction != nil {
		if err := validateSelectAction(c.SelectAction); err != nil {
			return fmt.Errorf("columnSet.selectAction: %w", err)
		}
	}
	if c.Style != "" && !c.Style.IsValid() {
		return m.NewEnumError("ColumnSet.style", string(c.Style), m.AllowedContainerStyleStrings())
	}
	if c.HorizontalAlignment != "" && !c.HorizontalAlignment.IsValid() {
		return m.NewEnumError(
			"ColumnSet.horizontalAlignment",
			string(c.HorizontalAlignment),
			m.AllowedHorizontalAlignmentStrings(),
		)
	}
	if c.MinHeight != "" && !isValidPixelWidth(c.MinHeight) {
		return fmt.Errorf(
			"columnSet.minHeight must be in format \"<number>px\" (got %q)",
			c.MinHeight,
		)
	}
	return nil
}

func (c ColumnSet) MarshalJSON() ([]byte, error) {
	cc := c
	if cc.Type == "" {
		cc.Type = m.TypeColumnSet
	}
	type alias ColumnSet
	return json.Marshal(alias(cc))
}

func (c *ColumnSet) UnmarshalJSON(b []byte) error {
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(b, &obj); err != nil {
		return fmt.Errorf("columnSet: decode: %w", err)
	}
	if v, ok := obj["type"]; ok {
		var t string
		if err := json.Unmarshal(v, &t); err != nil {
			return fmt.Errorf("columnSet.type: %w", err)
		}
		if t != "" && t != string(m.TypeColumnSet) {
			return fmt.Errorf("columnSet.type must be %q (got %q)", m.TypeColumnSet, t)
		}
	}
	selectActionRaw := obj["selectAction"]
	delete(obj, "selectAction")

	type alias ColumnSet
	var base alias
	baseBytes, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("columnSet: re-encode base: %w", err)
	}
	if err := json.Unmarshal(baseBytes, &base); err != nil {
		return fmt.Errorf("columnSet: decode base: %w", err)
	}
	if base.Type == "" {
		base.Type = m.TypeColumnSet
	}

	if len(selectActionRaw) != 0 {
		action, err := a.UnmarshalAction(selectActionRaw)
		if err != nil {
			return fmt.Errorf("columnSet.selectAction: %w", err)
		}
		base.SelectAction = action
	}

	*c = ColumnSet(base)
	return nil
}

func (c ColumnSet) validateElementBase() error {
	return c.ElementBase.Validate("ColumnSet")
}

func init() {
	e.RegisterElement(
		m.TypeColumnSet,
		func() e.Element { return &ColumnSet{Type: m.TypeColumnSet} },
	)
}
