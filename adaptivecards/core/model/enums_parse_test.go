package model

import "testing"

func TestEnumParseFunctions(t *testing.T) {
	if _, err := ParseImageFillMode("cover"); err != nil {
		t.Fatalf("ParseImageFillMode valid failed: %v", err)
	}
	if _, err := ParseHorizontalAlignment("left"); err != nil {
		t.Fatalf("ParseHorizontalAlignment valid failed: %v", err)
	}
	if _, err := ParseVerticalAlignment("top"); err != nil {
		t.Fatalf("ParseVerticalAlignment valid failed: %v", err)
	}
	if _, err := ParseVerticalContentAlignment("center"); err != nil {
		t.Fatalf("ParseVerticalContentAlignment valid failed: %v", err)
	}
	if _, err := ParseTextColor("accent"); err != nil {
		t.Fatalf("ParseTextColor valid failed: %v", err)
	}
	if _, err := ParseFontType("default"); err != nil {
		t.Fatalf("ParseFontType valid failed: %v", err)
	}
	if _, err := ParseFontSize("medium"); err != nil {
		t.Fatalf("ParseFontSize valid failed: %v", err)
	}
	if _, err := ParseFontWeight("bolder"); err != nil {
		t.Fatalf("ParseFontWeight valid failed: %v", err)
	}
	if _, err := ParseTextBlockStyle("heading"); err != nil {
		t.Fatalf("ParseTextBlockStyle valid failed: %v", err)
	}
	if _, err := ParseBlockElementHeight("auto"); err != nil {
		t.Fatalf("ParseBlockElementHeight valid failed: %v", err)
	}
	if _, err := ParseSpacing("small"); err != nil {
		t.Fatalf("ParseSpacing valid failed: %v", err)
	}
	if _, err := ParseImageSize("small"); err != nil {
		t.Fatalf("ParseImageSize valid failed: %v", err)
	}
	if _, err := ParseImageStyle("default"); err != nil {
		t.Fatalf("ParseImageStyle valid failed: %v", err)
	}
	if _, err := ParseActionStyle("positive"); err != nil {
		t.Fatalf("ParseActionStyle valid failed: %v", err)
	}
	if _, err := ParseActionMode("primary"); err != nil {
		t.Fatalf("ParseActionMode valid failed: %v", err)
	}
	if _, err := ParseContainerStyle("default"); err != nil {
		t.Fatalf("ParseContainerStyle valid failed: %v", err)
	}
}

func TestEnumParseFunctionsRejectInvalidValues(t *testing.T) {
	if _, err := ParseTextColor("not-valid"); err == nil {
		t.Fatalf("expected ParseTextColor invalid error")
	}
	if _, err := ParseActionStyle("not-valid"); err == nil {
		t.Fatalf("expected ParseActionStyle invalid error")
	}
	if _, err := ParseContainerStyle("not-valid"); err == nil {
		t.Fatalf("expected ParseContainerStyle invalid error")
	}
}
