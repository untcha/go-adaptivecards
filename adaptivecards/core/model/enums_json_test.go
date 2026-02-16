package model

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestEnumsJSONRoundtripAndValidation(t *testing.T) {
	if !ColorAccent.IsValid() || !FontTypeDefault.IsValid() || !SizeMedium.IsValid() {
		t.Fatalf("expected valid enum values")
	}

	check := func(v any, validJSON string, dst any, invalidJSON string) {
		if _, err := json.Marshal(v); err != nil {
			t.Fatalf("marshal failed: %v", err)
		}
		if err := json.Unmarshal([]byte(validJSON), dst); err != nil {
			t.Fatalf("unmarshal valid failed: %v", err)
		}
		if err := json.Unmarshal([]byte(invalidJSON), dst); err == nil {
			t.Fatalf("expected invalid enum error for %s", invalidJSON)
		}
	}

	var textColor TextColor
	check(ColorAccent, `"accent"`, &textColor, `"invalid"`)
	var fontType FontType
	check(FontTypeDefault, `"default"`, &fontType, `"invalid"`)
	var fontSize FontSize
	check(SizeMedium, `"medium"`, &fontSize, `"invalid"`)
	var fontWeight FontWeight
	check(WeightBolder, `"bolder"`, &fontWeight, `"invalid"`)
	var textStyle TextBlockStyle
	check(TextBlockStyleHeading, `"heading"`, &textStyle, `"invalid"`)
	var blockHeight BlockElementHeight
	check(BlockElementHeightAuto, `"auto"`, &blockHeight, `"invalid"`)
	var spacing Spacing
	check(SpacingSmall, `"small"`, &spacing, `"invalid"`)
	var imageSize ImageSize
	check(ImageSizeSmall, `"small"`, &imageSize, `"invalid"`)
	var imageStyle ImageStyle
	check(ImageStyleDefault, `"default"`, &imageStyle, `"invalid"`)
	var actionStyle ActionStyle
	check(ActionStyleDefault, `"default"`, &actionStyle, `"invalid"`)
	var actionMode ActionMode
	check(ActionModePrimary, `"primary"`, &actionMode, `"invalid"`)
	var containerStyle ContainerStyle
	check(ContainerStyleDefault, `"default"`, &containerStyle, `"invalid"`)
	var vContent VerticalContentAlignment
	check(VContentAlignTop, `"top"`, &vContent, `"invalid"`)
}

func TestEnumMarshalRejectsInvalidNonEmptyValue(t *testing.T) {
	var s Spacing = "invalid"
	_, err := json.Marshal(s)
	if err == nil {
		t.Fatalf("expected marshal error for invalid spacing")
	}

	var enumErr *EnumError
	if !errors.As(err, &enumErr) {
		t.Fatalf("expected EnumError, got %T", err)
	}
	if !errors.Is(err, ErrInvalidEnum) {
		t.Fatalf("expected ErrInvalidEnum")
	}
}
