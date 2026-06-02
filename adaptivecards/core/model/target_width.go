package model

import (
	"encoding/json"
	"strings"
)

// TargetWidth controls at which card widths an element is rendered (responsive
// visibility). It is a documented Adaptive Cards host feature but is NOT part of
// the published JSON schema (any version) — hosts that don't support it ignore
// it, and the library strips it before strict schema validation.
//
// Values are either a bare width bucket (e.g. "narrow") or a bucket prefixed
// with a comparison operator (e.g. "atLeast:wide", "atMost:narrow").
//
// Behavior & host support (important):
//
//   - It filters on the CARD's rendered width, not the device. A normal Teams
//     message card renders in the "standard" bucket, so "atLeast:wide" elements
//     are effectively never shown unless the card is also made full-width
//     (see Card.SetFullWidth / the msteams.width host extension). targetWidth and
//     full-width are companions: full-width lets the card reach "wide", and
//     targetWidth then tailors what appears at each width.
//   - Support is inconsistent across hosts. Teams Desktop honors targetWidth;
//     some clients (e.g. Teams mobile, as of this writing) ignore it and show the
//     element regardless — the documented graceful-degradation fallback for an
//     out-of-schema feature. Design cards so that an element being shown when the
//     property is ignored is acceptable.
//
// See: https://learn.microsoft.com/en-us/adaptive-cards/authoring-cards/responsive-layout
type TargetWidth string

// Bare width buckets.
const (
	TargetWidthVeryNarrow TargetWidth = "veryNarrow"
	TargetWidthNarrow     TargetWidth = "narrow"
	TargetWidthStandard   TargetWidth = "standard"
	TargetWidthWide       TargetWidth = "wide"
)

// Operator prefixes for range forms.
const (
	targetWidthAtLeastPrefix = "atLeast:"
	targetWidthAtMostPrefix  = "atMost:"
)

// targetWidthBuckets are the allowed bare bucket values.
var targetWidthBuckets = []TargetWidth{
	TargetWidthVeryNarrow,
	TargetWidthNarrow,
	TargetWidthStandard,
	TargetWidthWide,
}

// TargetWidthAtLeast builds the "atLeast:<bucket>" form (element shown at the
// given width and wider).
func TargetWidthAtLeast(bucket TargetWidth) TargetWidth {
	return TargetWidth(targetWidthAtLeastPrefix + string(bucket))
}

// TargetWidthAtMost builds the "atMost:<bucket>" form (element shown at the
// given width and narrower).
func TargetWidthAtMost(bucket TargetWidth) TargetWidth {
	return TargetWidth(targetWidthAtMostPrefix + string(bucket))
}

// isTargetWidthBucket reports whether s is one of the bare bucket values.
func isTargetWidthBucket(s string) bool {
	return enumIsValid(TargetWidth(s), targetWidthBuckets)
}

// IsValid reports whether w is a valid TargetWidth: a bare bucket, or an
// "atLeast:"/"atMost:" prefix followed by a valid bucket. The empty value is
// not valid (callers guard with `w != "" && !w.IsValid()`).
func (w TargetWidth) IsValid() bool {
	s := string(w)
	if after, ok := strings.CutPrefix(s, targetWidthAtLeastPrefix); ok {
		return isTargetWidthBucket(after)
	}
	if after, ok := strings.CutPrefix(s, targetWidthAtMostPrefix); ok {
		return isTargetWidthBucket(after)
	}
	return isTargetWidthBucket(s)
}

// AllowedTargetWidthStrings returns the documented values, including the prefix
// templates, for use in validation error messages.
func AllowedTargetWidthStrings() []string {
	out := enumAllowedStrings(targetWidthBuckets)
	return append(out, targetWidthAtLeastPrefix+"<bucket>", targetWidthAtMostPrefix+"<bucket>")
}

// ParseTargetWidth parses s into a TargetWidth or returns an EnumError.
func ParseTargetWidth(s string) (TargetWidth, error) {
	v := TargetWidth(s)
	if !v.IsValid() {
		return "", newEnumError("targetWidth", s, AllowedTargetWidthStrings())
	}
	return v, nil
}

// UnmarshalJSON enforces allowed values during decode. Because the field lives
// on the embedded ElementBase, this gives every element eager validation of
// targetWidth for free. An empty/absent value is accepted.
func (w *TargetWidth) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if s == "" {
		*w = ""
		return nil
	}
	v := TargetWidth(s)
	if !v.IsValid() {
		return newEnumError("targetWidth", s, AllowedTargetWidthStrings())
	}
	*w = v
	return nil
}

// MarshalJSON guards against emitting an invalid non-empty value.
func (w TargetWidth) MarshalJSON() ([]byte, error) {
	if w != "" && !w.IsValid() {
		return nil, newEnumError("targetWidth", string(w), AllowedTargetWidthStrings())
	}
	return json.Marshal(string(w))
}
