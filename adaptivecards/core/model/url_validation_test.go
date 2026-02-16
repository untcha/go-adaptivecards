package model

import "testing"

func TestURLValidationPublicWrappers(t *testing.T) {
	if _, err := ValidateActionURL("https://example.com", ErrInvalidActionOpenURLURL); err != nil {
		t.Fatalf("expected valid action url: %v", err)
	}
	if _, err := ValidateActionURL("relative/path", ErrInvalidActionOpenURLURL); err == nil {
		t.Fatalf("expected relative action url to be rejected")
	}

	if _, err := ValidateImageURL("https://example.com/image.png", ErrInvalidImageURL); err != nil {
		t.Fatalf("expected valid strict image url: %v", err)
	}
	if _, err := ValidateImageURL("not-a-url", ErrInvalidImageURL); err == nil {
		t.Fatalf("expected strict image url to reject non-absolute value")
	}

	if _, err := ValidateImageURIReference("not-a-url", ErrInvalidImageURL); err != nil {
		t.Fatalf("expected uri-reference image url to accept relative value: %v", err)
	}
	if _, err := ValidateImageURIReference("data:image/png;base64,AAAA", ErrInvalidImageURL); err != nil {
		t.Fatalf("expected uri-reference image url to accept data URI: %v", err)
	}
	if _, err := ValidateImageURIReference("   ", ErrInvalidImageURL); err == nil {
		t.Fatalf("expected empty uri-reference to fail")
	}
}
