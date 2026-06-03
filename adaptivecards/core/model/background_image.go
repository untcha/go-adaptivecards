package model

import (
	"encoding/json"
	"fmt"
)

// BackgroundImage specifies a background image. Acceptable formats are PNG, JPEG, and GIF.
// See: https://adaptivecards.io/explorer/BackgroundImage.html
type BackgroundImage struct {
	URL                 string              `json:"url"`                           // Version 1.2
	FillMode            ImageFillMode       `json:"fillMode,omitempty"`            // Version 1.2
	HorizontalAlignment HorizontalAlignment `json:"horizontalAlignment,omitempty"` // Version 1.2
	VerticalAlignment   VerticalAlignment   `json:"verticalAlignment,omitempty"`   // Version 1.2
}

// BackgroundImageValue allows backgroundImage to be encoded as either a URL string
// or a full object.
type BackgroundImageValue struct {
	obj     *BackgroundImage
	urlOnly string
}

// validateBackgroundImageURL validates and normalizes a background image URL.
// Returns the normalized URL string or an error if invalid.
func validateBackgroundImageURL(u string) (string, error) {
	return validateImageURL(u, ErrInvalidBackgroundImageURL)
}

// NewBackgroundImage creates a new BackgroundImage with the specified URL, fill mode, and alignments.
// URL is validated immediately (critical field), other parameters are validated later in Validate().
func NewBackgroundImage(
	url string,
	fillMode ImageFillMode,
	horizontalAlignment HorizontalAlignment,
	verticalAlignment VerticalAlignment,
) (BackgroundImage, error) {
	validURL, err := validateBackgroundImageURL(url)
	if err != nil {
		return BackgroundImage{}, err
	}

	return BackgroundImage{
		URL:                 validURL,
		FillMode:            fillMode,
		HorizontalAlignment: horizontalAlignment,
		VerticalAlignment:   verticalAlignment,
	}, nil
}

// MustNewBackgroundImage is like NewBackgroundImage but panics on error.
// Useful for tests and demos where you know the values are valid.
func MustNewBackgroundImage(
	url string,
	fillMode ImageFillMode,
	horizontalAlignment HorizontalAlignment,
	verticalAlignment VerticalAlignment,
) BackgroundImage {
	bg, err := NewBackgroundImage(url, fillMode, horizontalAlignment, verticalAlignment)
	if err != nil {
		panic(err)
	}
	return bg
}

// Validate performs comprehensive validation of the BackgroundImage.
// This should be called before using the BackgroundImage in a card.
func (bg BackgroundImage) Validate() error {
	// Validate URL (required field)
	if bg.URL == "" {
		return fmt.Errorf("BackgroundImage.url is required and cannot be empty")
	}

	// Re-validate URL format (defense in depth)
	if _, err := validateBackgroundImageURL(bg.URL); err != nil {
		return fmt.Errorf("BackgroundImage.url: %w", err)
	}

	// Validate enums if provided (optional fields)
	if bg.FillMode != "" && !bg.FillMode.IsValid() {
		return newEnumError(
			"BackgroundImage.fillMode",
			string(bg.FillMode),
			allowedImageFillModeStrings(),
		)
	}

	if bg.HorizontalAlignment != "" && !bg.HorizontalAlignment.IsValid() {
		return newEnumError(
			"BackgroundImage.horizontalAlignment",
			string(bg.HorizontalAlignment),
			allowedHorizontalAlignmentStrings(),
		)
	}

	if bg.VerticalAlignment != "" && !bg.VerticalAlignment.IsValid() {
		return newEnumError(
			"BackgroundImage.verticalAlignment",
			string(bg.VerticalAlignment),
			allowedVerticalAlignmentStrings(),
		)
	}

	return nil
}

// Builder methods for BackgroundImage are lenient; validation happens in Validate.

// WithFillMode sets the fill mode for the background image.
func (bg BackgroundImage) WithFillMode(mode ImageFillMode) BackgroundImage {
	bg.FillMode = mode
	return bg
}

// WithHorizontalAlignment sets the horizontal alignment for the background image.
func (bg BackgroundImage) WithHorizontalAlignment(align HorizontalAlignment) BackgroundImage {
	bg.HorizontalAlignment = align
	return bg
}

// WithVerticalAlignment sets the vertical alignment for the background image.
func (bg BackgroundImage) WithVerticalAlignment(align VerticalAlignment) BackgroundImage {
	bg.VerticalAlignment = align
	return bg
}

// Convenience methods for common fill modes.

// FillCover sets the fill mode to "cover".
func (bg BackgroundImage) FillCover() BackgroundImage {
	bg.FillMode = ImageFillCover
	return bg
}

// FillRepeatHorizontally sets the fill mode to "repeatHorizontally".
func (bg BackgroundImage) FillRepeatHorizontally() BackgroundImage {
	bg.FillMode = ImageFillRepeatHorizontally
	return bg
}

// FillRepeatVertically sets the fill mode to "repeatVertically".
func (bg BackgroundImage) FillRepeatVertically() BackgroundImage {
	bg.FillMode = ImageFillRepeatVertically
	return bg
}

// FillRepeat sets the fill mode to "repeat".
func (bg BackgroundImage) FillRepeat() BackgroundImage {
	bg.FillMode = ImageFillRepeat
	return bg
}

// Convenience methods for alignments.

// AlignLeft sets horizontal alignment to "left".
func (bg BackgroundImage) AlignLeft() BackgroundImage {
	bg.HorizontalAlignment = HAlignLeft
	return bg
}

// AlignCenter sets horizontal alignment to "center".
func (bg BackgroundImage) AlignCenter() BackgroundImage {
	bg.HorizontalAlignment = HAlignCenter
	return bg
}

// AlignRight sets horizontal alignment to "right".
func (bg BackgroundImage) AlignRight() BackgroundImage {
	bg.HorizontalAlignment = HAlignRight
	return bg
}

// AlignTop sets vertical alignment to "top".
func (bg BackgroundImage) AlignTop() BackgroundImage {
	bg.VerticalAlignment = VAlignTop
	return bg
}

// AlignMiddle sets vertical alignment to "center".
func (bg BackgroundImage) AlignMiddle() BackgroundImage {
	bg.VerticalAlignment = VAlignCenter
	return bg
}

// AlignBottom sets vertical alignment to "bottom".
func (bg BackgroundImage) AlignBottom() BackgroundImage {
	bg.VerticalAlignment = VAlignBottom
	return bg
}

// BackgroundImageValue helper functions.

// BackgroundImageURL creates a BackgroundImageValue from a URL string.
func BackgroundImageURL(u string) (*BackgroundImageValue, error) {
	validURL, err := validateBackgroundImageURL(u)
	if err != nil {
		return nil, err
	}
	return &BackgroundImageValue{urlOnly: validURL}, nil
}

// MustBackgroundImageURL is a convenience for tests/demos — panics on error.
func MustBackgroundImageURL(u string) *BackgroundImageValue {
	v, err := BackgroundImageURL(u)
	if err != nil {
		panic(err)
	}
	return v
}

// BackgroundImageObject creates a BackgroundImageValue from a BackgroundImage object.
func BackgroundImageObject(obj BackgroundImage) (*BackgroundImageValue, error) {
	// Validate the object using the standard Validate method
	if err := obj.Validate(); err != nil {
		return nil, err
	}
	return &BackgroundImageValue{obj: &obj}, nil
}

// MustBackgroundImageObject is a convenience for tests/demos - it panics on error.
func MustBackgroundImageObject(obj BackgroundImage) *BackgroundImageValue {
	v, err := BackgroundImageObject(obj)
	if err != nil {
		panic(err)
	}
	return v
}

// BackgroundImageURLUnchecked stores a raw URL without validation.
func BackgroundImageURLUnchecked(u string) *BackgroundImageValue {
	return &BackgroundImageValue{urlOnly: u}
}

// BackgroundImageObjectUnchecked stores a raw object without validation.
func BackgroundImageObjectUnchecked(obj BackgroundImage) *BackgroundImageValue {
	return &BackgroundImageValue{obj: &obj}
}

// Object returns the object-form value when set.
func (v *BackgroundImageValue) Object() *BackgroundImage {
	if v == nil {
		return nil
	}
	return v.obj
}

// URL returns the URL-form value when set.
func (v *BackgroundImageValue) URL() string {
	if v == nil {
		return ""
	}
	return v.urlOnly
}

// MarshalJSON implements custom JSON marshalling for BackgroundImageValue.
func (v BackgroundImageValue) MarshalJSON() ([]byte, error) {
	if v.obj == nil {
		return json.Marshal(v.urlOnly) // string form
	}
	return json.Marshal(v.obj) // object form
}

// UnmarshalJSON implements custom JSON unmarshalling for BackgroundImageValue.
func (v *BackgroundImageValue) UnmarshalJSON(b []byte) error {
	// Try string first
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		// Validate the URL when unmarshaling
		validURL, err := validateBackgroundImageURL(s)
		if err != nil {
			return err
		}
		*v = BackgroundImageValue{urlOnly: validURL}
		return nil
	}

	// Try object form
	var o BackgroundImage
	if err := json.Unmarshal(b, &o); err != nil {
		return err
	}

	// Validate the unmarshaled object
	if err := o.Validate(); err != nil {
		return err
	}

	*v = BackgroundImageValue{obj: &o}
	return nil
}
