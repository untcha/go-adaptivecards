package model

import (
	"encoding/json"
	"slices"
	"strings"
)

// stringEnum is any defined type whose underlying type is string.
type stringEnum interface{ ~string }

// NewEnumError creates a typed enum validation error for a public field.
// It wraps the shared ErrInvalidEnum marker and includes allowed values.
func NewEnumError(field, got string, allowed []string) error {
	return newEnumError(field, got, allowed)
}

// EnumAllowedStrings converts a typed enum value list into a []string copy.
// The returned slice is detached from the input.
func EnumAllowedStrings[T ~string](values []T) []string {
	return enumAllowedStrings(values)
}

// enumIsValid reports whether v is in allowed.
func enumIsValid[T stringEnum](v T, allowed []T) bool {
	return slices.Contains(allowed, v)
	// for _, a := range allowed {
	// 	if v == a {
	// 		return true
	// 	}
	// }
	// return false
}

// enumAllowedStrings returns a []string copy for messages.
func enumAllowedStrings[T stringEnum](allowed []T) []string {
	out := make([]string, len(allowed))
	for i, v := range allowed {
		out[i] = string(v)
	}
	return out
}

// enumParse checks s against allowed and returns a typed value or an EnumError.
func enumParse[T stringEnum](s string, allowed []T, field string) (T, error) {
	if v, ok := enumParseCanonical(s, allowed); ok {
		return v, nil
	}
	var zero T
	return zero, newEnumError(field, s, enumAllowedStrings(allowed))
}

// enumUnmarshalJSON decodes a JSON string into dst, enforcing allowed values.
func enumUnmarshalJSON[T stringEnum](dst *T, b []byte, allowed []T, field string) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	v, ok := enumParseCanonical(s, allowed)
	if !ok {
		return newEnumError(field, s, enumAllowedStrings(allowed))
	}
	*dst = v
	return nil
}

func enumParseCanonical[T stringEnum](s string, allowed []T) (T, bool) {
	v := T(s)
	if enumIsValid(v, allowed) {
		return v, true
	}
	for _, a := range allowed {
		if strings.EqualFold(s, string(a)) {
			return a, true
		}
	}
	var zero T
	return zero, false
}

// enumMarshalJSON encodes v, optionally guarding invalid non-empty values.
func enumMarshalJSON[T stringEnum](v T, allowed []T, field string) ([]byte, error) {
	var zero T
	if v != zero && !enumIsValid(v, allowed) {
		return nil, newEnumError(field, string(v), enumAllowedStrings(allowed))
	}
	return json.Marshal(string(v))
}
