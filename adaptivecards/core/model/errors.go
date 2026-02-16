package model

import (
	"errors"
	"strconv"
	"strings"
)

// ErrInvalidImageURL is returned when an image URL is invalid.
var ErrInvalidImageURL = errors.New("invalid Image URL")

// ErrInvalidBackgroundImageURL is returned when a background image URL is invalid.
var ErrInvalidBackgroundImageURL = errors.New("invalid BackgroundImage URL")

// ErrInvalidActionOpenURLURL is returned when an Action.OpenUrl URL is invalid.
var ErrInvalidActionOpenURLURL = errors.New("invalid Action.OpenUrl URL")

// ErrInvalidBackgroundImageFillMode is returned when a background image fill mode is invalid.
var ErrInvalidBackgroundImageFillMode = errors.New("invalid BackgroundImage fillMode")

// ErrInvalidMinHeight is returned when minHeight does not match the expected format.
var ErrInvalidMinHeight = errors.New("invalid minHeight")

// ErrInvalidEnum marks enum validation failures.
var ErrInvalidEnum = errors.New("invalid enum value")

// EnumError describes an invalid enum value.
type EnumError struct {
	Field   string   // e.g. "ImageFillMode" or "BackgroundImage.fillMode"
	Got     string   // the invalid raw value
	Allowed []string // printable list of allowed values
}

func (e *EnumError) Error() string {
	return "invalid " + e.Field + " " + strconv.Quote(e.Got) +
		" (allowed: " + strings.Join(e.Allowed, ", ") + ")"
}

func (e *EnumError) Unwrap() error { return ErrInvalidEnum }

// helper constructor to avoid repeating format logic
func newEnumError(field, got string, allowed []string) error {
	return &EnumError{Field: field, Got: got, Allowed: allowed}
}
