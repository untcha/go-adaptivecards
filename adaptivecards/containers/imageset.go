package containers

import (
	"encoding/json"
	"fmt"
	"strings"

	e "github.com/untcha/go-adaptivecards/adaptivecards/core/element"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
	els "github.com/untcha/go-adaptivecards/adaptivecards/elements"
)

// ImageSet displays a collection of Images similar to a gallery. Acceptable formats are PNG, JPEG, and GIF.
// See: https://adaptivecards.io/explorer/ImageSet.html
type ImageSet struct {
	e.ElementBase              // Embedding e.ElementBase to include common element fields
	Type          m.TypeString `json:"type"`                // Version 1.0
	Images        []els.Image  `json:"images,omitempty"`    // Version 1.0
	ImageSize     m.ImageSize  `json:"imageSize,omitempty"` // Version 1.0
}

// NewImageSet returns an ImageSet holding the given images.
func NewImageSet(images ...els.Image) ImageSet {
	return ImageSet{
		ElementBase: e.ElementBase{},
		Type:        m.TypeImageSet,
		Images:      images,
	}
}

// GetType returns the Adaptive Card type discriminator for ImageSet.
func (i ImageSet) GetType() m.TypeString { return m.TypeImageSet }

// Validate reports whether the ImageSet and its images are well-formed.
func (i ImageSet) Validate() error {
	if err := i.validateElementBase(); err != nil {
		return err
	}
	if len(i.Images) == 0 {
		return fmt.Errorf("imageSet.images is required")
	}
	for idx, img := range i.Images {
		if err := img.Validate(); err != nil {
			return fmt.Errorf("imageSet.images[%d]: %w", idx, err)
		}
	}
	if i.ImageSize != "" && !i.ImageSize.IsValid() {
		return m.NewEnumError(
			"ImageSet.imageSize",
			string(i.ImageSize),
			m.AllowedImageSizeStrings(),
		)
	}
	return nil
}

// MarshalJSON implements json.Marshaler, ensuring the type field is set and the image size is normalized.
func (i ImageSet) MarshalJSON() ([]byte, error) {
	ii := i
	if ii.Type == "" {
		ii.Type = m.TypeImageSet
	}
	ii.ImageSize = ii.normalizedImageSize()
	type alias ImageSet
	return json.Marshal(alias(ii))
}

// UnmarshalJSON implements json.Unmarshaler, normalizing the image size and validating the decoded value.
func (i *ImageSet) UnmarshalJSON(b []byte) error {
	type alias ImageSet
	var tmp alias
	if err := json.Unmarshal(b, &tmp); err != nil {
		return fmt.Errorf("imageSet: decode: %w", err)
	}
	if tmp.Type != "" && tmp.Type != m.TypeImageSet {
		return fmt.Errorf("imageSet.type must be %q (got %q)", m.TypeImageSet, tmp.Type)
	}
	if tmp.Type == "" {
		tmp.Type = m.TypeImageSet
	}
	val := ImageSet(tmp)
	val.ImageSize = val.normalizedImageSize()
	if err := val.Validate(); err != nil {
		return err
	}
	*i = val
	return nil
}

func (i ImageSet) normalizedImageSize() m.ImageSize {
	if i.ImageSize == m.ImageSizeAuto || i.ImageSize == m.ImageSizeStretch {
		return m.ImageSizeMedium
	}
	return i.ImageSize
}

func (i ImageSet) validateElementBase() error {
	if i.Height != "" && !i.Height.IsValid() {
		return m.NewEnumError(
			"ImageSet.height",
			string(i.Height),
			m.AllowedBlockElementHeightStrings(),
		)
	}
	if i.Spacing != "" && !i.Spacing.IsValid() {
		return m.NewEnumError("ImageSet.spacing", string(i.Spacing), m.AllowedSpacingStrings())
	}
	if i.ID != "" {
		id := strings.TrimSpace(i.ID)
		if id == "" {
			return fmt.Errorf("imageSet.id cannot be empty or whitespace-only")
		}
		if strings.ContainsAny(id, "\n\r\t") {
			return fmt.Errorf("imageSet.id cannot contain newlines or tabs")
		}
	}
	return nil
}

func init() {
	e.RegisterElement(m.TypeImageSet, func() e.Element { return &ImageSet{Type: m.TypeImageSet} })
}
