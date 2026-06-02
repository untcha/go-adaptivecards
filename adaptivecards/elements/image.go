package elements

import (
	"encoding/json"
	"fmt"
	"strings"

	a "github.com/untcha/go-adaptivecards/adaptivecards/actions"
	e "github.com/untcha/go-adaptivecards/adaptivecards/core/element"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// Image
// Displays an image. Acceptable formats are PNG, JPEG, and GIF
// See: https://adaptivecards.io/explorer/Image.html
type Image struct {
	e.ElementBase                             // Embedding e.ElementBase to include common element fields
	Type                m.TypeString          `json:"type"`                          // Version 1.0
	URL                 m.URI                 `json:"url"`                           // Version 1.0
	AltText             string                `json:"altText,omitempty"`             // Version 1.0
	BackgroundColor     string                `json:"backgroundColor,omitempty"`     // Version 1.1
	Height              any                   `json:"height,omitempty"`              // Version 1.1 - string or m.BlockElementHeight
	HorizontalAlignment m.HorizontalAlignment `json:"horizontalAlignment,omitempty"` // Version 1.0
	SelectAction        a.Action              `json:"selectAction,omitempty"`        // Version 1.1
	Size                m.ImageSize           `json:"size,omitempty"`                // Version 1.0
	Style               m.ImageStyle          `json:"style,omitempty"`               // Version 1.0
	Width               string                `json:"width,omitempty"`               // Version 1.1
}

// NewImageEmpty creates an Image with no URL set.
func NewImageEmpty() Image {
	return Image{
		ElementBase: e.ElementBase{},
		Type:        m.TypeImage,
	}
}

// NewImage creates an Image with the specified URL.
func NewImage(u string) Image {
	return Image{
		ElementBase: e.ElementBase{},
		Type:        m.TypeImage,
		URL:         m.URI(u),
	}
}

func (i Image) GetType() m.TypeString { return m.TypeImage }

// Fluent setters return a copy of Image with the specified field set.

// WithURL sets the URL field.
func (i Image) WithURL(u string) Image { i.URL = m.URI(u); return i }

// WithAltText sets the AltText field.
func (i Image) WithAltText(alt string) Image { i.AltText = alt; return i }

// WithBackgroundColor sets the BackgroundColor field.
func (i Image) WithBackgroundColor(color string) Image { i.BackgroundColor = color; return i }

// WithHeight sets the Height field (BlockElementHeight enum).
func (i Image) WithHeight(h m.BlockElementHeight) Image { i.Height = h; return i }

// WithHeightPx sets the Height field (string like "50px").
func (i Image) WithHeightPx(h string) Image { i.Height = h; return i }

// WithAlign sets the HorizontalAlignment field.
func (i Image) WithAlign(a m.HorizontalAlignment) Image { i.HorizontalAlignment = a; return i }

// AlignLeft sets HorizontalAlignment to left.
func (i Image) AlignLeft() Image { i.HorizontalAlignment = m.HAlignLeft; return i }

// AlignCenter sets HorizontalAlignment to center.
func (i Image) AlignCenter() Image { i.HorizontalAlignment = m.HAlignCenter; return i }

// AlignRight sets HorizontalAlignment to right.
func (i Image) AlignRight() Image { i.HorizontalAlignment = m.HAlignRight; return i }

// WithSelectAction sets the SelectAction field.
func (i Image) WithSelectAction(a a.Action) Image { i.SelectAction = a; return i }

// WithSize sets the Size field.
func (i Image) WithSize(sz m.ImageSize) Image { i.Size = sz; return i }

// SizeAuto sets Size to auto.
func (i Image) SizeAuto() Image { i.Size = m.ImageSizeAuto; return i }

// SizeStretch sets Size to stretch.
func (i Image) SizeStretch() Image { i.Size = m.ImageSizeStretch; return i }

// SizeSmall sets Size to small.
func (i Image) SizeSmall() Image { i.Size = m.ImageSizeSmall; return i }

// SizeMedium sets Size to medium.
func (i Image) SizeMedium() Image { i.Size = m.ImageSizeMedium; return i }

// SizeLarge sets Size to large.
func (i Image) SizeLarge() Image { i.Size = m.ImageSizeLarge; return i }

// WithStyle sets the Style field.
func (i Image) WithStyle(s m.ImageStyle) Image { i.Style = s; return i }

// StyleDefault sets Style to default.
func (i Image) StyleDefault() Image { i.Style = m.ImageStyleDefault; return i }

// StylePerson sets Style to person.
func (i Image) StylePerson() Image { i.Style = m.ImageStylePerson; return i }

// WithWidth sets the Width field (string like "50px").
func (i Image) WithWidth(w string) Image { i.Width = w; return i }

// Fluent setters for embedded e.ElementBase fields.

// WithFallback sets the Fallback field.
func (i Image) WithFallback(fallback *e.ElementFallback) Image {
	i.Fallback = fallback
	return i
}

// WithSeparator sets the Separator field.
func (i Image) WithSeparator(separator bool) Image {
	i.Separator = separator
	return i
}

// WithSpacing sets the Spacing field.
func (i Image) WithSpacing(spacing m.Spacing) Image { i.Spacing = spacing; return i }

// WithTargetWidth sets the TargetWidth field (responsive visibility).
func (i Image) WithTargetWidth(tw m.TargetWidth) Image { i.TargetWidth = tw; return i }

// WithID sets the ID field.
func (i Image) WithID(id string) Image { i.ID = id; return i }

// WithVisible sets the IsVisible field.
func (i Image) WithVisible(visible bool) Image { i.IsVisible = &visible; return i }

// WithRequires sets the Requires field.
func (i Image) WithRequires(requires map[string]string) Image {
	i.Requires = requires
	return i
}

// Convenience methods for common spacing values.
func (i Image) SpacingDefault() Image    { i.Spacing = m.SpacingDefault; return i }
func (i Image) SpacingNone() Image       { i.Spacing = m.SpacingNone; return i }
func (i Image) SpacingSmall() Image      { i.Spacing = m.SpacingSmall; return i }
func (i Image) SpacingMedium() Image     { i.Spacing = m.SpacingMedium; return i }
func (i Image) SpacingLarge() Image      { i.Spacing = m.SpacingLarge; return i }
func (i Image) SpacingExtraLarge() Image { i.Spacing = m.SpacingExtraLarge; return i }
func (i Image) SpacingPadding() Image    { i.Spacing = m.SpacingPadding; return i }

// Convenience methods for visibility.
func (i Image) Hide() Image { visible := false; i.IsVisible = &visible; return i }
func (i Image) Show() Image { visible := true; i.IsVisible = &visible; return i }

// Validate checks that the Image has valid fields.
func (i Image) Validate() error {
	// Validate embedded e.ElementBase fields
	if err := i.validateElementBase(); err != nil {
		return err
	}

	if i.URL == "" {
		return fmt.Errorf("image.url is required and cannot be empty")
	}

	// Validate URL format
	if _, err := m.ValidateImageURIReference(string(i.URL), m.ErrInvalidImageURL); err != nil {
		return fmt.Errorf("image.url: %w", err)
	}

	// Validate height if provided (string like "50px" or BlockElementHeight enum).
	if i.Height != nil {
		switch h := i.Height.(type) {
		case string:
			if h != "" {
				if strings.HasSuffix(h, "px") {
					break
				}
				if !m.BlockElementHeight(h).IsValid() {
					return fmt.Errorf(
						"image.height as string must be '<n>px' or a valid height token (got %q)",
						h,
					)
				}
			}
		case m.BlockElementHeight:
			if !h.IsValid() {
				return m.NewEnumError("Image.height", string(h), m.AllowedBlockElementHeightStrings())
			}
		default:
			return fmt.Errorf("image.height must be string or BlockElementHeight")
		}
	}

	// Validate enums
	if i.HorizontalAlignment != "" && !i.HorizontalAlignment.IsValid() {
		return m.NewEnumError(
			"Image.horizontalAlignment",
			string(i.HorizontalAlignment),
			m.AllowedHorizontalAlignmentStrings(),
		)
	}
	if i.Size != "" && !i.Size.IsValid() {
		return m.NewEnumError("Image.size", string(i.Size), m.AllowedImageSizeStrings())
	}
	if i.Style != "" && !i.Style.IsValid() {
		return m.NewEnumError("Image.style", string(i.Style), m.AllowedImageStyleStrings())
	}
	if i.SelectAction != nil {
		if err := validateSelectAction(i.SelectAction); err != nil {
			return fmt.Errorf("image.selectAction: %w", err)
		}
	}

	return nil
}

// MarshalJSON ensures Type is always set.
func (i Image) MarshalJSON() ([]byte, error) {
	ii := i
	if ii.Type == "" {
		ii.Type = m.TypeImage
	}
	type alias Image
	return json.Marshal(alias(ii))
}

// UnmarshalJSON decodes Image and its selectAction field.
func (i *Image) UnmarshalJSON(b []byte) error {
	// Decode into a generic map to handle interface fields explicitly.
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(b, &obj); err != nil {
		return fmt.Errorf("Image: decode: %w", err)
	}

	// Validate/normalize type
	if v, ok := obj["type"]; ok {
		var t string
		if err := json.Unmarshal(v, &t); err != nil {
			return fmt.Errorf("image.type: %w", err)
		}
		if t != "" && t != string(m.TypeImage) {
			return fmt.Errorf("image.type must be %q (got %q)", m.TypeImage, t)
		}
	}

	// Extract raw selectAction and remove it before base decode
	selectActionRaw := obj["selectAction"]
	delete(obj, "selectAction")

	// Decode the remaining fields into base
	type alias Image
	var base alias
	baseBytes, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("Image: re-encode base: %w", err)
	}
	if err := json.Unmarshal(baseBytes, &base); err != nil {
		return fmt.Errorf("Image: decode base: %w", err)
	}

	// Default the type if omitted
	if base.Type == "" {
		base.Type = m.TypeImage
	}

	// Decode selectAction
	if len(selectActionRaw) != 0 {
		act, err := a.UnmarshalAction(selectActionRaw)
		if err != nil {
			return fmt.Errorf("Image.selectAction: %w", err)
		}
		base.SelectAction = act
	}

	val := Image(base)
	if err := val.Validate(); err != nil {
		return err
	}
	*i = val
	return nil
}

func (i Image) validateElementBase() error {
	return i.ElementBase.Validate("Image")
}

// Register Image in the element registry.
func init() {
	e.RegisterElement(m.TypeImage, func() e.Element { return &Image{Type: m.TypeImage} })
}
