package elements

import (
	"encoding/json"
	"fmt"
	"strings"

	e "github.com/untcha/go-adaptivecards/adaptivecards/core/element"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// RichTextBlock
// Defines an array of inlines, allowing for inline text formatting.
// See: https://adaptivecards.io/explorer/RichTextBlock.html
type RichTextBlock struct {
	e.ElementBase
	Type                m.TypeString          `json:"type"`                          // Version 1.2
	Inlines             []TextRun             `json:"inlines,omitempty"`             // Version 1.2
	HorizontalAlignment m.HorizontalAlignment `json:"horizontalAlignment,omitempty"` // Version 1.2
}

func (r RichTextBlock) GetType() m.TypeString { return m.TypeRichTextBlock }

func NewRichTextBlock(inlines ...TextRun) RichTextBlock {
	return RichTextBlock{
		ElementBase: e.ElementBase{},
		Type:        m.TypeRichTextBlock,
		Inlines:     inlines,
	}
}

func (r RichTextBlock) Validate() error {
	if err := r.validateElementBase(); err != nil {
		return err
	}
	if len(r.Inlines) == 0 {
		return fmt.Errorf("richTextBlock.inlines is required")
	}
	for i, inline := range r.Inlines {
		if err := inline.Validate(); err != nil {
			return fmt.Errorf("richTextBlock.inlines[%d]: %w", i, err)
		}
	}
	if r.HorizontalAlignment != "" && !r.HorizontalAlignment.IsValid() {
		return m.NewEnumError(
			"RichTextBlock.horizontalAlignment",
			string(r.HorizontalAlignment),
			m.AllowedHorizontalAlignmentStrings(),
		)
	}
	return nil
}

func (r RichTextBlock) validateElementBase() error {
	if r.Height != "" && !r.Height.IsValid() {
		return m.NewEnumError(
			"RichTextBlock.height",
			string(r.Height),
			m.AllowedBlockElementHeightStrings(),
		)
	}
	if r.Spacing != "" && !r.Spacing.IsValid() {
		return m.NewEnumError("RichTextBlock.spacing", string(r.Spacing), m.AllowedSpacingStrings())
	}
	if r.ID != "" {
		id := strings.TrimSpace(r.ID)
		if id == "" {
			return fmt.Errorf("richTextBlock.id cannot be empty or whitespace-only")
		}
		if strings.ContainsAny(id, "\n\r\t") {
			return fmt.Errorf("richTextBlock.id cannot contain newlines or tabs")
		}
	}
	return nil
}

func (r RichTextBlock) MarshalJSON() ([]byte, error) {
	rr := r
	if rr.Type == "" {
		rr.Type = m.TypeRichTextBlock
	}
	type alias RichTextBlock
	return json.Marshal(alias(rr))
}

func (r *RichTextBlock) UnmarshalJSON(b []byte) error {
	type alias RichTextBlock
	var tmp alias
	if err := json.Unmarshal(b, &tmp); err != nil {
		return fmt.Errorf("richTextBlock: decode: %w", err)
	}
	if tmp.Type != "" && tmp.Type != m.TypeRichTextBlock {
		return fmt.Errorf("richTextBlock.type must be %q (got %q)", m.TypeRichTextBlock, tmp.Type)
	}
	if tmp.Type == "" {
		tmp.Type = m.TypeRichTextBlock
	}
	val := RichTextBlock(tmp)
	if err := val.Validate(); err != nil {
		return err
	}
	*r = val
	return nil
}

func init() {
	e.RegisterElement(
		m.TypeRichTextBlock,
		func() e.Element { return &RichTextBlock{Type: m.TypeRichTextBlock} },
	)
}
