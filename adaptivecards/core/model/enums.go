package model

// ImageFillMode describes how the image should fill the area.
type ImageFillMode string

const (
	// The background image covers the entire width of the container. Its aspect ratio is preserved.
	// Content may be clipped if the aspect ratio of the image doesn't match the aspect ratio of the container.
	// verticalAlignment is respected (horizontalAlignment is meaningless since it's stretched width).
	// This is the default mode and is the equivalent to the current model.
	ImageFillCover ImageFillMode = "cover"

	// The background image isn't stretched. It is repeated in the x axis as many times as necessary to cover the container's width.
	// verticalAlignment is honored (default is top), horizontalAlignment is ignored.
	ImageFillRepeatHorizontally ImageFillMode = "repeatHorizontally"

	// The background image isn't stretched. It is repeated in the y axis as many times as necessary to cover the container's height.
	// verticalAlignment is ignored, horizontalAlignment is honored (default is left).
	ImageFillRepeatVertically ImageFillMode = "repeatVertically"

	// The background image isn't stretched. It is repeated first in the x axis then in the y axis as many times as necessary to cover the entire container.
	// Both horizontalAlignment and verticalAlignment are honored (defaults are left and top).
	ImageFillRepeat ImageFillMode = "repeat"
)

// All allowed values for ImageFillMode
var imageFillModeAllowed = []ImageFillMode{
	ImageFillCover,
	ImageFillRepeatHorizontally,
	ImageFillRepeatVertically,
	ImageFillRepeat,
}

// IsValid reports whether m is a valid ImageFillMode value.
func (m ImageFillMode) IsValid() bool {
	return enumIsValid(m, imageFillModeAllowed)
}

// allowedImageFillModeStrings returns the list of allowed ImageFillMode values as strings.
func allowedImageFillModeStrings() []string {
	return enumAllowedStrings(imageFillModeAllowed)
}

// ParseImageFillMode parses s into an ImageFillMode or returns an EnumError.
func ParseImageFillMode(s string) (ImageFillMode, error) {
	return enumParse(s, imageFillModeAllowed, "ImageFillMode")
}

// UnmarshalJSON implements json.Unmarshaler, enforcing allowed values.
func (m *ImageFillMode) UnmarshalJSON(b []byte) error {
	return enumUnmarshalJSON(m, b, imageFillModeAllowed, "ImageFillMode")
}

// MarshalJSON implements json.Marshaler, enforcing allowed values.
func (m ImageFillMode) MarshalJSON() ([]byte, error) {
	return enumMarshalJSON(m, imageFillModeAllowed, "ImageFillMode")
}

// HorizontalAlignment describes how the image should be aligned if it must be cropped or if using repeat fill mode.
type HorizontalAlignment string

const (
	HAlignLeft   HorizontalAlignment = "left"
	HAlignCenter HorizontalAlignment = "center"
	HAlignRight  HorizontalAlignment = "right"
)

// All allowed values for HorizontalAlignment
var horizontalAlignmentAllowed = []HorizontalAlignment{
	HAlignLeft,
	HAlignCenter,
	HAlignRight,
}

// IsValid reports whether h is a valid horizontalAlignmentAllowed value.
func (h HorizontalAlignment) IsValid() bool {
	return enumIsValid(h, horizontalAlignmentAllowed)
}

// allowedHorizontalAlignmentStrings returns the list of allowed HorizontalAlignment values as strings.
func allowedHorizontalAlignmentStrings() []string {
	return enumAllowedStrings(horizontalAlignmentAllowed)
}

// ParseHorizontalAlignment parses s into a HorizontalAlignment or returns an EnumError.
func ParseHorizontalAlignment(s string) (HorizontalAlignment, error) {
	return enumParse(s, horizontalAlignmentAllowed, "HorizontalAlignment")
}

// UnmarshalJSON implements json.Unmarshaler, enforcing allowed values.
func (h *HorizontalAlignment) UnmarshalJSON(b []byte) error {
	return enumUnmarshalJSON(h, b, horizontalAlignmentAllowed, "HorizontalAlignment")
}

// MarshalJSON implements json.Marshaler, enforcing allowed values.
func (h HorizontalAlignment) MarshalJSON() ([]byte, error) {
	return enumMarshalJSON(h, horizontalAlignmentAllowed, "HorizontalAlignment")
}

// VerticalAlignment describes how the image should be aligned if it must be cropped or if using repeat fill mode.
type VerticalAlignment string

const (
	VAlignTop    VerticalAlignment = "top"
	VAlignCenter VerticalAlignment = "center"
	VAlignBottom VerticalAlignment = "bottom"
)

// All allowed values for VerticalAlignment
var verticalAlignmentAllowed = []VerticalAlignment{
	VAlignTop,
	VAlignCenter,
	VAlignBottom,
}

// IsValid reports whether m is a valid verticalAlignmentAllowed value.
func (v VerticalAlignment) IsValid() bool {
	return enumIsValid(v, verticalAlignmentAllowed)
}

// allowedVerticalAlignmentStrings returns the list of allowed VerticalAlignment values as strings.
func allowedVerticalAlignmentStrings() []string {
	return enumAllowedStrings(verticalAlignmentAllowed)
}

// ParseVerticalAlignment parses s into a VerticalAlignment or returns an EnumError.
func ParseVerticalAlignment(s string) (VerticalAlignment, error) {
	return enumParse(s, verticalAlignmentAllowed, "VerticalAlignment")
}

// UnmarshalJSON implements json.Unmarshaler, enforcing allowed values.
func (v *VerticalAlignment) UnmarshalJSON(b []byte) error {
	return enumUnmarshalJSON(v, b, verticalAlignmentAllowed, "VerticalAlignment")
}

// MarshalJSON implements json.Marshaler, enforcing allowed values.
func (v VerticalAlignment) MarshalJSON() ([]byte, error) {
	return enumMarshalJSON(v, verticalAlignmentAllowed, "VerticalAlignment")
}

// VerticalContentAlignment defines how the content should be aligned vertically within the container.
// Only relevant for fixed-height cards, or cards with a minHeight specified.
type VerticalContentAlignment string

const (
	VContentAlignTop    VerticalContentAlignment = "top"
	VContentAlignCenter VerticalContentAlignment = "center"
	VContentAlignBottom VerticalContentAlignment = "bottom"
)

// All allowed values for VerticalContentAlignment
var verticalContentAlignmentAllowed = []VerticalContentAlignment{
	VContentAlignTop,
	VContentAlignCenter,
	VContentAlignBottom,
}

// IsValid reports whether v is a valid VerticalContentAlignment value.
func (v VerticalContentAlignment) IsValid() bool {
	return enumIsValid(v, verticalContentAlignmentAllowed)
}

// allowedVerticalContentAlignmentStrings returns the list of allowed VerticalContentAlignment values as strings.
func allowedVerticalContentAlignmentStrings() []string {
	return enumAllowedStrings(verticalContentAlignmentAllowed)
}

// ParseVerticalContentAlignment parses s into a VerticalContentAlignment or returns an EnumError.
func ParseVerticalContentAlignment(s string) (VerticalContentAlignment, error) {
	return enumParse(s, verticalContentAlignmentAllowed, "VerticalContentAlignment")
}

// UnmarshalJSON implements json.Unmarshaler, enforcing allowed values.
func (v *VerticalContentAlignment) UnmarshalJSON(b []byte) error {
	return enumUnmarshalJSON(v, b, verticalContentAlignmentAllowed, "VerticalContentAlignment")
}

// MarshalJSON implements json.Marshaler, enforcing allowed values.
func (v VerticalContentAlignment) MarshalJSON() ([]byte, error) {
	return enumMarshalJSON(v, verticalContentAlignmentAllowed, "VerticalContentAlignment")
}

// TextColor controls the color of TextBlock elements.
type TextColor string

const (
	ColorDefault   TextColor = "default"
	ColorDark      TextColor = "dark"
	ColorLight     TextColor = "light"
	ColorAccent    TextColor = "accent"
	ColorGood      TextColor = "good"
	ColorWarning   TextColor = "warning"
	ColorAttention TextColor = "attention"
)

// All allowed values for TextColor
var textColorAllowed = []TextColor{
	ColorDefault,
	ColorDark,
	ColorLight,
	ColorAccent,
	ColorGood,
	ColorWarning,
	ColorAttention,
}

// IsValid reports whether c is a valid TextColor value.
func (c TextColor) IsValid() bool {
	return enumIsValid(c, textColorAllowed)
}

// allowedTextColorStrings returns the list of allowed TextColor values as strings.
func allowedTextColorStrings() []string {
	return enumAllowedStrings(textColorAllowed)
}

// ParseTextColor parses s into a TextColor or returns an EnumError.
func ParseTextColor(s string) (TextColor, error) {
	return enumParse[TextColor](s, textColorAllowed, "TextColor")
}

// UnmarshalJSON implements json.Unmarshaler, enforcing allowed values.
func (c *TextColor) UnmarshalJSON(b []byte) error {
	return enumUnmarshalJSON(c, b, textColorAllowed, "TextColor")
}

// MarshalJSON implements json.Marshaler, enforcing allowed values.
func (c TextColor) MarshalJSON() ([]byte, error) {
	return enumMarshalJSON(c, textColorAllowed, "TextColor")
}

// FontType is the type of font to use for rendering.
type FontType string

const (
	FontTypeDefault   FontType = "default"
	FontTypeMonospace FontType = "monospace"
)

// All allowed values for FontType
var fontTypeAllowed = []FontType{
	FontTypeDefault,
	FontTypeMonospace,
}

// IsValid reports whether f is a valid FontType value.
func (f FontType) IsValid() bool {
	return enumIsValid(f, fontTypeAllowed)
}

// allowedFontTypeStrings returns the list of allowed FontType values as strings.
func allowedFontTypeStrings() []string {
	return enumAllowedStrings(fontTypeAllowed)
}

// ParseFontType parses s into a FontType or returns an EnumError.
func ParseFontType(s string) (FontType, error) {
	return enumParse(s, fontTypeAllowed, "FontType")
}

// UnmarshalJSON implements json.Unmarshaler, enforcing allowed values.
func (f *FontType) UnmarshalJSON(b []byte) error {
	return enumUnmarshalJSON(f, b, fontTypeAllowed, "FontType")
}

// MarshalJSON implements json.Marshaler, enforcing allowed values.
func (f FontType) MarshalJSON() ([]byte, error) {
	return enumMarshalJSON(f, fontTypeAllowed, "FontType")
}

// FontSize controls the size of text.
type FontSize string

const (
	SizeDefault    FontSize = "default"
	SizeSmall      FontSize = "small"
	SizeMedium     FontSize = "medium"
	SizeLarge      FontSize = "large"
	SizeExtraLarge FontSize = "extraLarge"
)

// All allowed values for FontSize
var fontSizeAllowed = []FontSize{
	SizeDefault,
	SizeSmall,
	SizeMedium,
	SizeLarge,
	SizeExtraLarge,
}

// IsValid reports whether s is a valid FontSize value.
func (s FontSize) IsValid() bool {
	return enumIsValid(s, fontSizeAllowed)
}

// allowedFontSizeStrings returns the list of allowed FontSize values as strings.
func allowedFontSizeStrings() []string {
	return enumAllowedStrings(fontSizeAllowed)
}

// ParseFontSize parses s into a FontSize or returns an EnumError.
func ParseFontSize(s string) (FontSize, error) {
	return enumParse(s, fontSizeAllowed, "FontSize")
}

// UnmarshalJSON implements json.Unmarshaler, enforcing allowed values.
func (s *FontSize) UnmarshalJSON(b []byte) error {
	return enumUnmarshalJSON(s, b, fontSizeAllowed, "FontSize")
}

// MarshalJSON implements json.Marshaler, enforcing allowed values.
func (s FontSize) MarshalJSON() ([]byte, error) {
	return enumMarshalJSON(s, fontSizeAllowed, "FontSize")
}

// FontWeight controls the weight of TextBlock elements.
type FontWeight string

const (
	WeightDefault FontWeight = "default"
	WeightLighter FontWeight = "lighter"
	WeightBolder  FontWeight = "bolder"
)

// All allowed values for FontWeight
var fontWeightAllowed = []FontWeight{
	WeightDefault,
	WeightLighter,
	WeightBolder,
}

// IsValid reports whether w is a valid FontWeight value.
func (w FontWeight) IsValid() bool {
	return enumIsValid(w, fontWeightAllowed)
}

// allowedFontWeightStrings returns the list of allowed FontWeight values as strings.
func allowedFontWeightStrings() []string {
	return enumAllowedStrings(fontWeightAllowed)
}

// ParseFontWeight parses s into a FontWeight or returns an EnumError.
func ParseFontWeight(s string) (FontWeight, error) {
	return enumParse(s, fontWeightAllowed, "FontWeight")
}

// UnmarshalJSON implements json.Unmarshaler, enforcing allowed values.
func (w *FontWeight) UnmarshalJSON(b []byte) error {
	return enumUnmarshalJSON(w, b, fontWeightAllowed, "FontWeight")
}

// MarshalJSON implements json.Marshaler, enforcing allowed values.
func (w FontWeight) MarshalJSON() ([]byte, error) {
	return enumMarshalJSON(w, fontWeightAllowed, "FontWeight")
}

// TextBlockStyle is the style of a TextBlock for accessibility purposes.
type TextBlockStyle string

const (
	TextBlockStyleDefault TextBlockStyle = "default"
	TextBlockStyleHeading TextBlockStyle = "heading"
)

// All allowed values for TextBlockStyle
var textBlockStyleAllowed = []TextBlockStyle{
	TextBlockStyleDefault,
	TextBlockStyleHeading,
}

// IsValid reports whether s is a valid TextBlockStyle value.
func (s TextBlockStyle) IsValid() bool {
	return enumIsValid(s, textBlockStyleAllowed)
}

// allowedTextBlockStyleStrings returns the list of allowed TextBlockStyle values as strings.
func allowedTextBlockStyleStrings() []string {
	return enumAllowedStrings(textBlockStyleAllowed)
}

// ParseTextBlockStyle parses s into a TextBlockStyle or returns an EnumError.
func ParseTextBlockStyle(s string) (TextBlockStyle, error) {
	return enumParse(s, textBlockStyleAllowed, "TextBlockStyle")
}

// UnmarshalJSON implements json.Unmarshaler, enforcing allowed values.
func (s *TextBlockStyle) UnmarshalJSON(b []byte) error {
	return enumUnmarshalJSON(s, b, textBlockStyleAllowed, "TextBlockStyle")
}

// MarshalJSON implements json.Marshaler, enforcing allowed values.
func (s TextBlockStyle) MarshalJSON() ([]byte, error) {
	return enumMarshalJSON(s, textBlockStyleAllowed, "TextBlockStyle")
}

// BlockElementHeight specifies the height of the element.
type BlockElementHeight string

const (
	BlockElementHeightAuto    BlockElementHeight = "auto"
	BlockElementHeightStretch BlockElementHeight = "stretch"
)

// All allowed values for BlockElementHeight
var blockElementHeightAllowed = []BlockElementHeight{
	BlockElementHeightAuto,
	BlockElementHeightStretch,
}

// IsValid reports whether h is a valid BlockElementHeight value.
func (h BlockElementHeight) IsValid() bool {
	return enumIsValid(h, blockElementHeightAllowed)
}

// allowedBlockElementHeightStrings returns the list of allowed BlockElementHeight values as strings.
func allowedBlockElementHeightStrings() []string {
	return enumAllowedStrings(blockElementHeightAllowed)
}

// ParseBlockElementHeight parses s into a BlockElementHeight or returns an EnumError.
func ParseBlockElementHeight(s string) (BlockElementHeight, error) {
	return enumParse(s, blockElementHeightAllowed, "BlockElementHeight")
}

// UnmarshalJSON implements json.Unmarshaler, enforcing allowed values.
func (h *BlockElementHeight) UnmarshalJSON(b []byte) error {
	return enumUnmarshalJSON(h, b, blockElementHeightAllowed, "BlockElementHeight")
}

// MarshalJSON implements json.Marshaler, enforcing allowed values.
func (h BlockElementHeight) MarshalJSON() ([]byte, error) {
	return enumMarshalJSON(h, blockElementHeightAllowed, "BlockElementHeight")
}

// Spacing controls the amount of spacing between this element and the preceding element.
type Spacing string

const (
	SpacingDefault    Spacing = "default"
	SpacingNone       Spacing = "none"
	SpacingSmall      Spacing = "small"
	SpacingMedium     Spacing = "medium"
	SpacingLarge      Spacing = "large"
	SpacingExtraLarge Spacing = "extraLarge"
	SpacingPadding    Spacing = "padding"
)

// All allowed values for Spacing
var spacingAllowed = []Spacing{
	SpacingDefault,
	SpacingNone,
	SpacingSmall,
	SpacingMedium,
	SpacingLarge,
	SpacingExtraLarge,
	SpacingPadding,
}

// IsValid reports whether s is a valid Spacing value.
func (s Spacing) IsValid() bool {
	return enumIsValid(s, spacingAllowed)
}

// allowedSpacingStrings returns the list of allowed Spacing values as strings.
func allowedSpacingStrings() []string {
	return enumAllowedStrings(spacingAllowed)
}

// ParseSpacing parses s into a Spacing or returns an EnumError.
func ParseSpacing(s string) (Spacing, error) {
	return enumParse(s, spacingAllowed, "Spacing")
}

// UnmarshalJSON implements json.Unmarshaler, enforcing allowed values.
func (s *Spacing) UnmarshalJSON(b []byte) error {
	return enumUnmarshalJSON(s, b, spacingAllowed, "Spacing")
}

// MarshalJSON implements json.Marshaler, enforcing allowed values.
func (s Spacing) MarshalJSON() ([]byte, error) {
	return enumMarshalJSON(s, spacingAllowed, "Spacing")
}

// ImageSize controls the approximate size of the image. The physical dimensions will vary per host.
type ImageSize string

const (
	ImageSizeAuto    ImageSize = "auto"
	ImageSizeStretch ImageSize = "stretch"
	ImageSizeSmall   ImageSize = "small"
	ImageSizeMedium  ImageSize = "medium"
	ImageSizeLarge   ImageSize = "large"
)

// All allowed values for ImageSize
var imageSizeAllowed = []ImageSize{
	ImageSizeAuto,
	ImageSizeStretch,
	ImageSizeSmall,
	ImageSizeMedium,
	ImageSizeLarge,
}

// IsValid reports whether s is a valid ImageSize value.
func (s ImageSize) IsValid() bool {
	return enumIsValid(s, imageSizeAllowed)
}

// allowedImageSizeStrings returns the list of allowed ImageSize values as strings.
func allowedImageSizeStrings() []string {
	return enumAllowedStrings(imageSizeAllowed)
}

// ParseImageSize parses s into a ImageSize or returns an EnumError.
func ParseImageSize(s string) (ImageSize, error) {
	return enumParse(s, imageSizeAllowed, "ImageSize")
}

// UnmarshalJSON implements json.Unmarshaler, enforcing allowed values.
func (s *ImageSize) UnmarshalJSON(b []byte) error {
	return enumUnmarshalJSON(s, b, imageSizeAllowed, "ImageSize")
}

// MarshalJSON implements json.Marshaler, enforcing allowed values.
func (s ImageSize) MarshalJSON() ([]byte, error) {
	return enumMarshalJSON(s, imageSizeAllowed, "ImageSize")
}

// ImageStyle controls how an Image is displayed.
type ImageStyle string

const (
	ImageStyleDefault ImageStyle = "default"
	ImageStylePerson  ImageStyle = "person"
)

// All allowed values for ImageStyle
var imageStyleAllowed = []ImageStyle{
	ImageStyleDefault,
	ImageStylePerson,
}

// IsValid reports whether s is a valid ImageStyle value.
func (s ImageStyle) IsValid() bool {
	return enumIsValid(s, imageStyleAllowed)
}

// allowedImageStyleStrings returns the list of allowed ImageStyle values as strings.
func allowedImageStyleStrings() []string {
	return enumAllowedStrings(imageStyleAllowed)
}

// ParseImageStyle parses s into a ImageStyle or returns an EnumError.
func ParseImageStyle(s string) (ImageStyle, error) {
	return enumParse(s, imageStyleAllowed, "ImageStyle")
}

// UnmarshalJSON implements json.Unmarshaler, enforcing allowed values.
func (s *ImageStyle) UnmarshalJSON(b []byte) error {
	return enumUnmarshalJSON(s, b, imageStyleAllowed, "ImageStyle")
}

// MarshalJSON implements json.Marshaler, enforcing allowed values.
func (s ImageStyle) MarshalJSON() ([]byte, error) {
	return enumMarshalJSON(s, imageStyleAllowed, "ImageStyle")
}

// ActionStyle controls the style of an Action, which influences how the action is displayed, spoken, etc.
type ActionStyle string

const (
	ActionStyleDefault     ActionStyle = "default"
	ActionStylePositive    ActionStyle = "positive"
	ActionStyleDestructive ActionStyle = "destructive"
)

// All allowed values for ActionStyle
var actionStyleAllowed = []ActionStyle{
	ActionStyleDefault,
	ActionStylePositive,
	ActionStyleDestructive,
}

// IsValid reports whether s is a valid ActionStyle value.
func (s ActionStyle) IsValid() bool {
	return enumIsValid(s, actionStyleAllowed)
}

// ParseActionStyle parses s into a ActionStyle or returns an EnumError.
func ParseActionStyle(s string) (ActionStyle, error) {
	return enumParse(s, actionStyleAllowed, "ActionStyle")
}

// UnmarshalJSON implements json.Unmarshaler, enforcing allowed values.
func (s *ActionStyle) UnmarshalJSON(b []byte) error {
	return enumUnmarshalJSON(s, b, actionStyleAllowed, "ActionStyle")
}

// MarshalJSON implements json.Marshaler, enforcing allowed values.
func (s ActionStyle) MarshalJSON() ([]byte, error) {
	return enumMarshalJSON(s, actionStyleAllowed, "ActionStyle")
}

// ActionMode determines whether the action should be displayed as a button or in the overflow menu.
type ActionMode string

const (
	ActionModePrimary   ActionMode = "primary"
	ActionModeSecondary ActionMode = "secondary"
)

// All allowed values for ActionMode
var actionModeAllowed = []ActionMode{
	ActionModePrimary,
	ActionModeSecondary,
}

// IsValid reports whether s is a valid ActionMode value.
func (s ActionMode) IsValid() bool {
	return enumIsValid(s, actionModeAllowed)
}

// ParseActionMode parses s into a ActionMode or returns an EnumError.
func ParseActionMode(s string) (ActionMode, error) {
	return enumParse[ActionMode](s, actionModeAllowed, "ActionMode")
}

// UnmarshalJSON implements json.Unmarshaler, enforcing allowed values.
func (s *ActionMode) UnmarshalJSON(b []byte) error {
	return enumUnmarshalJSON(s, b, actionModeAllowed, "ActionMode")
}

// MarshalJSON implements json.Marshaler, enforcing allowed values.
func (s ActionMode) MarshalJSON() ([]byte, error) {
	return enumMarshalJSON(s, actionModeAllowed, "ActionMode")
}

// ContainerStyle defines the style of the grid. This property currently only controls the grid's color.
type ContainerStyle string

const (
	ContainerStyleDefault   ContainerStyle = "default"
	ContainerStyleEmphasis  ContainerStyle = "emphasis"
	ContainerStyleGood      ContainerStyle = "good"      // Added in (Schema) Version 1.2
	ContainerStyleAttention ContainerStyle = "attention" // Added in (Schema) Version 1.2
	ContainerStyleWarning   ContainerStyle = "warning"   // Added in (Schema) Version 1.2
	ContainerStyleAccent    ContainerStyle = "accent"    // Added in (Schema) Version 1.2
)

// All allowed values for ContainerStyle
var containerStyleAllowed = []ContainerStyle{
	ContainerStyleDefault,
	ContainerStyleEmphasis,
	ContainerStyleGood,
	ContainerStyleAttention,
	ContainerStyleWarning,
	ContainerStyleAccent,
}

// IsValid reports whether s is a valid ContainerStyle value.
func (s ContainerStyle) IsValid() bool {
	return enumIsValid(s, containerStyleAllowed)
}

// allowedContainerStyleStrings returns the list of allowed ContainerStyle values as strings.
func allowedContainerStyleStrings() []string {
	return enumAllowedStrings(containerStyleAllowed)
}

// ParseContainerStyle parses s into a ContainerStyle or returns an EnumError.
func ParseContainerStyle(s string) (ContainerStyle, error) {
	return enumParse(s, containerStyleAllowed, "ContainerStyle")
}

// UnmarshalJSON implements json.Unmarshaler, enforcing allowed values.
func (s *ContainerStyle) UnmarshalJSON(b []byte) error {
	return enumUnmarshalJSON(s, b, containerStyleAllowed, "ContainerStyle")
}

// MarshalJSON implements json.Marshaler, enforcing allowed values.
func (s ContainerStyle) MarshalJSON() ([]byte, error) {
	return enumMarshalJSON(s, containerStyleAllowed, "ContainerStyle")
}

// AllowedImageFillModeStrings returns the list of valid ImageFillMode values as strings.
func AllowedImageFillModeStrings() []string { return allowedImageFillModeStrings() }

// AllowedHorizontalAlignmentStrings returns the list of valid HorizontalAlignment values as strings.
func AllowedHorizontalAlignmentStrings() []string { return allowedHorizontalAlignmentStrings() }

// AllowedVerticalAlignmentStrings returns the list of valid VerticalAlignment values as strings.
func AllowedVerticalAlignmentStrings() []string { return allowedVerticalAlignmentStrings() }

// AllowedVerticalContentAlignmentStrings returns the list of valid VerticalContentAlignment values as strings.
func AllowedVerticalContentAlignmentStrings() []string {
	return allowedVerticalContentAlignmentStrings()
}

// AllowedTextColorStrings returns the list of valid TextColor values as strings.
func AllowedTextColorStrings() []string { return allowedTextColorStrings() }

// AllowedFontTypeStrings returns the list of valid FontType values as strings.
func AllowedFontTypeStrings() []string { return allowedFontTypeStrings() }

// AllowedFontSizeStrings returns the list of valid FontSize values as strings.
func AllowedFontSizeStrings() []string { return allowedFontSizeStrings() }

// AllowedFontWeightStrings returns the list of valid FontWeight values as strings.
func AllowedFontWeightStrings() []string { return allowedFontWeightStrings() }

// AllowedSpacingStrings returns the list of valid Spacing values as strings.
func AllowedSpacingStrings() []string { return allowedSpacingStrings() }

// AllowedBlockElementHeightStrings returns the list of valid BlockElementHeight values as strings.
func AllowedBlockElementHeightStrings() []string { return allowedBlockElementHeightStrings() }

// AllowedImageSizeStrings returns the list of valid ImageSize values as strings.
func AllowedImageSizeStrings() []string { return allowedImageSizeStrings() }

// AllowedImageStyleStrings returns the list of valid ImageStyle values as strings.
func AllowedImageStyleStrings() []string { return allowedImageStyleStrings() }

// AllowedContainerStyleStrings returns the list of valid ContainerStyle values as strings.
func AllowedContainerStyleStrings() []string { return allowedContainerStyleStrings() }

// AllowedTextBlockStyleStrings returns the list of valid TextBlockStyle values as strings.
func AllowedTextBlockStyleStrings() []string { return allowedTextBlockStyleStrings() }
