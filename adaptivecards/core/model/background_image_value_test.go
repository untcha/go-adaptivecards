package model

import "testing"

func TestBackgroundImageValueHelperAccessorsAndUnchecked(t *testing.T) {
	var nilValue *BackgroundImageValue
	if nilValue.Object() != nil {
		t.Fatalf("expected nil object for nil receiver")
	}
	if nilValue.URL() != "" {
		t.Fatalf("expected empty url for nil receiver")
	}

	u := BackgroundImageURLUnchecked("not-a-url")
	if got := u.URL(); got != "not-a-url" {
		t.Fatalf("expected unchecked url to be preserved, got %q", got)
	}
	if u.Object() != nil {
		t.Fatalf("expected unchecked url value to have nil object")
	}

	obj := BackgroundImage{URL: "https://example.com/bg.png"}
	o := BackgroundImageObjectUnchecked(obj)
	if got := o.Object(); got == nil || got.URL != obj.URL {
		t.Fatalf("expected unchecked object to be preserved")
	}
	if o.URL() != "" {
		t.Fatalf("expected object-form value to return empty URL()")
	}
}
