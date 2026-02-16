package model

import (
	"errors"
	"testing"
)

func TestPublicEnumHelpers(t *testing.T) {
	allowed := EnumAllowedStrings([]TextColor{ColorDefault, ColorAccent})
	if len(allowed) != 2 || allowed[0] != "default" || allowed[1] != "accent" {
		t.Fatalf("unexpected allowed values: %#v", allowed)
	}

	err := NewEnumError("TextColor", "nope", allowed)
	var enumErr *EnumError
	if !errors.As(err, &enumErr) {
		t.Fatalf("expected EnumError, got %T", err)
	}
	if !errors.Is(err, ErrInvalidEnum) {
		t.Fatalf("expected ErrInvalidEnum marker")
	}
}
