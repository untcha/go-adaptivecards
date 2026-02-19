package elements

import (
	"encoding/json"
	"fmt"
	"strings"

	e "github.com/untcha/go-adaptivecards/adaptivecards/core/element"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// TextBlock
// Displays text, allowing control over font sizes, weight, and color.
// See: https://adaptivecards.io/explorer/TextBlock.html
type TextBlock struct {
	e.ElementBase                             // Embedding e.ElementBase to include common element fields
	Type                m.TypeString          `json:"type"`                          // Version 1.0
	Text                string                `json:"text"`                          // Version 1.0
	Color               m.TextColor           `json:"color,omitempty"`               // Version 1.0
	FontType            m.FontType            `json:"fontType,omitempty"`            // Version 1.2
	HorizontalAlignment m.HorizontalAlignment `json:"horizontalAlignment,omitempty"` // Version 1.0
	IsSubtle            *bool                 `json:"isSubtle,omitempty"`            // Version 1.0
	MaxLines            *float64              `json:"maxLines,omitempty"`            // Version 1.0
	Size                m.FontSize            `json:"size,omitempty"`                // Version 1.0
	Weight              m.FontWeight          `json:"weight,omitempty"`              // Version 1.0
	Wrap                bool                  `json:"wrap,omitempty"`                // Version 1.0
	Style               m.TextBlockStyle      `json:"style,omitempty"`               // Version 1.5
}

// NewTextBlockEmpty creates a TextBlock with empty text.
func NewTextBlockEmpty() TextBlock {
	return TextBlock{
		ElementBase: e.ElementBase{},
		Type:        m.TypeTextBlock,
	}
}

// NewTextBlock creates a TextBlock with the specified text.
func NewTextBlock(text string) TextBlock {
	return TextBlock{
		ElementBase: e.ElementBase{},
		Type:        m.TypeTextBlock,
		Text:        text,
	}
}

func (t TextBlock) GetType() m.TypeString { return m.TypeTextBlock }

// Fluent setters return a copy of TextBlock with the specified field set.

// WithText sets the Text field.
func (t TextBlock) WithText(s string) TextBlock { t.Text = s; return t }

// WithColor sets the Color field.
func (t TextBlock) WithColor(c m.TextColor) TextBlock { t.Color = c; return t }

// WithFontType sets the FontType field.
func (t TextBlock) WithFontType(ft m.FontType) TextBlock { t.FontType = ft; return t }

// WithAlign sets the HorizontalAlignment field.
func (t TextBlock) WithAlign(
	a m.HorizontalAlignment,
) TextBlock {
	t.HorizontalAlignment = a
	return t
}

// AlignLeft sets HorizontalAlignment to left.
func (t TextBlock) AlignLeft() TextBlock { t.HorizontalAlignment = m.HAlignLeft; return t }

// AlignCenter sets HorizontalAlignment to center.
func (t TextBlock) AlignCenter() TextBlock { t.HorizontalAlignment = m.HAlignCenter; return t }

// AlignRight sets HorizontalAlignment to right.
func (t TextBlock) AlignRight() TextBlock { t.HorizontalAlignment = m.HAlignRight; return t }

// WithSubtle sets the IsSubtle field.
func (t TextBlock) WithSubtle(b bool) TextBlock { t.IsSubtle = &b; return t }

// SubtleOn sets IsSubtle to true.
func (t TextBlock) SubtleOn() TextBlock { b := true; t.IsSubtle = &b; return t }

// SubtleOff sets IsSubtle to false.
func (t TextBlock) SubtleOff() TextBlock { b := false; t.IsSubtle = &b; return t }

// WithMaxLines sets the MaxLines field.
func (t TextBlock) WithMaxLines(n int) TextBlock {
	v := float64(n)
	t.MaxLines = &v
	return t
}

// WithSize sets the Size field.
func (t TextBlock) WithSize(sz m.FontSize) TextBlock { t.Size = sz; return t }

// WithWeight sets the Weight field.
func (t TextBlock) WithWeight(w m.FontWeight) TextBlock { t.Weight = w; return t }

// WithWrap sets the Wrap field.
func (t TextBlock) WithWrap(w bool) TextBlock { t.Wrap = w; return t }

// WrapOn sets Wrap to true.
func (t TextBlock) WrapOn() TextBlock { t.Wrap = true; return t }

// WrapOff sets Wrap to false.
func (t TextBlock) WrapOff() TextBlock { t.Wrap = false; return t }

// WithStyle sets the Style field.
func (t TextBlock) WithStyle(s m.TextBlockStyle) TextBlock { t.Style = s; return t }

// Fluent setters for embedded e.ElementBase fields.

// WithFallback sets the Fallback field.
func (t TextBlock) WithFallback(fallback *e.ElementFallback) TextBlock {
	t.Fallback = fallback
	return t
}

// WithHeight sets the Height field (BlockElementHeight).
func (t TextBlock) WithHeight(height m.BlockElementHeight) TextBlock {
	t.Height = height
	return t
}

// WithSeparator sets the Separator field.
func (t TextBlock) WithSeparator(separator bool) TextBlock {
	t.Separator = separator
	return t
}

// WithSpacing sets the Spacing field.
func (t TextBlock) WithSpacing(spacing m.Spacing) TextBlock { t.Spacing = spacing; return t }

// WithID sets the ID field.
func (t TextBlock) WithID(id string) TextBlock { t.ID = id; return t }

// WithVisible sets the IsVisible field.
func (t TextBlock) WithVisible(visible bool) TextBlock { t.IsVisible = &visible; return t }

// WithRequires sets the Requires field.
func (t TextBlock) WithRequires(requires map[string]string) TextBlock {
	t.Requires = requires
	return t
}

// Convenience methods for common spacing values.
func (t TextBlock) SpacingNone() TextBlock       { t.Spacing = m.SpacingNone; return t }
func (t TextBlock) SpacingDefault() TextBlock    { t.Spacing = m.SpacingDefault; return t }
func (t TextBlock) SpacingSmall() TextBlock      { t.Spacing = m.SpacingSmall; return t }
func (t TextBlock) SpacingMedium() TextBlock     { t.Spacing = m.SpacingMedium; return t }
func (t TextBlock) SpacingLarge() TextBlock      { t.Spacing = m.SpacingLarge; return t }
func (t TextBlock) SpacingExtraLarge() TextBlock { t.Spacing = m.SpacingExtraLarge; return t }
func (t TextBlock) SpacingPadding() TextBlock    { t.Spacing = m.SpacingPadding; return t }

// Convenience methods for visibility.
func (t TextBlock) Hide() TextBlock { visible := false; t.IsVisible = &visible; return t }
func (t TextBlock) Show() TextBlock { visible := true; t.IsVisible = &visible; return t }

// Validate checks that the TextBlock has valid fields.
func (t TextBlock) Validate() error {
	// Validate base element fields.
	if err := t.validateElementBase(); err != nil {
		return err
	}

	// Validate TextBlock-specific fields
	if t.Text == "" {
		return fmt.Errorf("TextBlock.text is required and cannot be empty")
	}

	// Validate enum values
	if t.Color != "" && !t.Color.IsValid() {
		return m.NewEnumError("TextBlock.color", string(t.Color), m.AllowedTextColorStrings())
	}
	if t.FontType != "" && !t.FontType.IsValid() {
		return m.NewEnumError("TextBlock.fontType", string(t.FontType), m.AllowedFontTypeStrings())
	}
	if t.HorizontalAlignment != "" && !t.HorizontalAlignment.IsValid() {
		return m.NewEnumError(
			"TextBlock.horizontalAlignment",
			string(t.HorizontalAlignment),
			m.AllowedHorizontalAlignmentStrings(),
		)
	}
	if t.Size != "" && !t.Size.IsValid() {
		return m.NewEnumError("TextBlock.fontSize", string(t.Size), m.AllowedFontSizeStrings())
	}
	if t.Weight != "" && !t.Weight.IsValid() {
		return m.NewEnumError("TextBlock.weight", string(t.Weight), m.AllowedFontWeightStrings())
	}
	if t.Style != "" && !t.Style.IsValid() {
		return m.NewEnumError("TextBlock.style", string(t.Style), m.AllowedTextBlockStyleStrings())
	}
	return nil
}

// MarshalJSON ensures Type is always set.
func (t TextBlock) MarshalJSON() ([]byte, error) {
	tt := t
	if tt.Type == "" {
		tt.Type = m.TypeTextBlock
	}
	type alias TextBlock
	return json.Marshal(alias(tt))
}

// UnmarshalJSON decodes TextBlock from JSON.
func (t *TextBlock) UnmarshalJSON(b []byte) error {
	type alias TextBlock // avoid recursion
	var tmp alias
	if err := json.Unmarshal(b, &tmp); err != nil {
		return fmt.Errorf("TextBlock: decode: %w", err)
	}

	// If "type" is present, it must be "TextBlock"
	if tmp.Type != "" && tmp.Type != m.TypeTextBlock {
		return fmt.Errorf("TextBlock.type must be %q (got %q)", m.TypeTextBlock, tmp.Type)
	}

	// Do NOT enforce tmp.Text != "" here — allow empty while building.
	// The Adaptive Cards schema (or your TextBlock.Validate) will check this later.

	// Validate enum values during unmarshaling for better error reporting
	if tmp.Color != "" && !tmp.Color.IsValid() {
		return m.NewEnumError("TextBlock.color", string(tmp.Color), m.AllowedTextColorStrings())
	}
	if tmp.FontType != "" && !tmp.FontType.IsValid() {
		return m.NewEnumError(
			"TextBlock.fontType",
			string(tmp.FontType),
			m.AllowedFontTypeStrings(),
		)
	}
	if tmp.HorizontalAlignment != "" && !tmp.HorizontalAlignment.IsValid() {
		return m.NewEnumError(
			"TextBlock.horizontalAlignment",
			string(tmp.HorizontalAlignment),
			m.AllowedHorizontalAlignmentStrings(),
		)
	}
	if tmp.Size != "" && !tmp.Size.IsValid() {
		return m.NewEnumError("TextBlock.fontSize", string(tmp.Size), m.AllowedFontSizeStrings())
	}
	if tmp.Weight != "" && !tmp.Weight.IsValid() {
		return m.NewEnumError("TextBlock.weight", string(tmp.Weight), m.AllowedFontWeightStrings())
	}
	if tmp.Style != "" && !tmp.Style.IsValid() {
		return m.NewEnumError(
			"TextBlock.style",
			string(tmp.Style),
			m.AllowedTextBlockStyleStrings(),
		)
	}

	// Default the type if omitted
	if tmp.Type == "" {
		tmp.Type = m.TypeTextBlock
	}

	*t = TextBlock(tmp)
	return nil
}

// validateElementBase validates inherited base element fields.
func (t TextBlock) validateElementBase() error {
	// Validate Height field (BlockElementHeight).
	if t.Height != "" && !t.Height.IsValid() {
		return m.NewEnumError(
			"TextBlock.height",
			fmt.Sprintf("%v", t.Height),
			m.AllowedBlockElementHeightStrings(),
		)
	}

	// Validate spacing enum.
	if t.Spacing != "" && !t.Spacing.IsValid() {
		return m.NewEnumError("TextBlock.spacing", string(t.Spacing), m.AllowedSpacingStrings())
	}

	// Validate ID field (basic sanity checks)
	if t.ID != "" {
		id := strings.TrimSpace(t.ID)
		if id == "" {
			return fmt.Errorf("TextBlock.id cannot be empty or whitespace-only")
		}
		// ID should not contain certain characters that could cause issues
		if strings.ContainsAny(id, "\n\r\t") {
			return fmt.Errorf("TextBlock.id cannot contain newlines or tabs")
		}
	}

	return nil
}

// Register TextBlock in the element registry.
func init() {
	e.RegisterElement(
		m.TypeTextBlock,
		func() e.Element { return &TextBlock{Type: m.TypeTextBlock} },
	)
}
