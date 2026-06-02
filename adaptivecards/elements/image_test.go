package elements

import (
	"encoding/json"
	"strings"
	"testing"

	a "github.com/untcha/go-adaptivecards/adaptivecards/actions"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

type fakeShowCardAction struct{}

func (fakeShowCardAction) GetType() m.TypeString { return m.TypeActionShowCard }

func TestImageGetType(t *testing.T) {
	if NewImage("https://example.com").GetType() != m.TypeImage {
		t.Fatalf("unexpected image type")
	}
}

func TestImageFluentMethodsAndJSON(t *testing.T) {
	action, err := a.NewActionOpenURL("open", "https://example.com")
	if err != nil {
		t.Fatalf("setup action failed: %v", err)
	}

	img := NewImageEmpty().
		WithURL("https://example.com/img.png").
		WithAltText("alt").
		WithBackgroundColor("#ffffff").
		WithHeight(m.BlockElementHeightAuto).
		WithAlign(m.HAlignCenter).
		WithSelectAction(action).
		WithSize(m.ImageSizeLarge).
		WithStyle(m.ImageStylePerson).
		WithWidth("42px").
		WithSeparator(true).
		WithSpacing(m.SpacingSmall).
		WithID("img-1").
		WithVisible(true).
		WithRequires(map[string]string{"feature": "1.0"}).
		SpacingPadding().
		Show()

	if err := img.Validate(); err != nil {
		t.Fatalf("unexpected image validate error: %v", err)
	}

	b, err := json.Marshal(img)
	if err != nil {
		t.Fatalf("marshal image failed: %v", err)
	}

	var decoded Image
	if err := json.Unmarshal(b, &decoded); err != nil {
		t.Fatalf("unmarshal image failed: %v", err)
	}
	if decoded.Type != m.TypeImage {
		t.Fatalf("expected image type")
	}
}

func TestImageValidateErrorUsesImageFieldNamesForSpacing(t *testing.T) {
	img := NewImage("https://example.com/cat.png")
	img.Spacing = "invalid-spacing"

	err := img.Validate()
	if err == nil {
		t.Fatalf("expected validation error")
	}
	if !strings.Contains(err.Error(), "Image.spacing") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestImageValidateErrorUsesImageFieldNamesForID(t *testing.T) {
	img := NewImage("https://example.com/cat.png")
	img.ID = "\n"

	err := img.Validate()
	if err == nil {
		t.Fatalf("expected validation error")
	}
	if !strings.Contains(err.Error(), "Image.id") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestImageValidateAdditionalErrorBranches(t *testing.T) {
	cases := []Image{
		NewImage(""),
		NewImage("https://example.com/i.png").WithHeightPx("20"),
		func() Image {
			i := NewImage("https://example.com/i.png")
			i.Height = m.BlockElementHeight("bad")
			return i
		}(),
		func() Image { i := NewImage("https://example.com/i.png"); i.Height = 42; return i }(),
		NewImage("https://example.com/i.png").WithAlign(m.HorizontalAlignment("bad")),
		NewImage("https://example.com/i.png").WithSize(m.ImageSize("bad")),
		NewImage("https://example.com/i.png").WithStyle(m.ImageStyle("bad")),
	}
	for idx, img := range cases {
		if err := img.Validate(); err == nil {
			t.Fatalf("expected image validation error for case %d", idx)
		}
	}
}

func TestImageValidateAcceptsHeightTokensAndPixels(t *testing.T) {
	cases := []Image{
		NewImage("https://example.com/i.png").WithHeightPx("20px"),
		NewImage("https://example.com/i.png").WithHeightPx("auto"),
		NewImage("https://example.com/i.png").WithHeightPx("stretch"),
	}
	for idx, img := range cases {
		if err := img.Validate(); err != nil {
			t.Fatalf("expected valid image height for case %d, got error: %v", idx, err)
		}
	}
}

func TestImageUnmarshalSelectActionAndErrors(t *testing.T) {
	var img Image
	if err := json.Unmarshal([]byte(`{"type":"Image","url":"https://example.com/i.png","selectAction":{"type":"Action.OpenUrl","url":"https://example.com"}}`), &img); err != nil {
		t.Fatalf("unexpected image selectAction unmarshal error: %v", err)
	}
	if img.SelectAction == nil {
		t.Fatalf("expected selectAction set")
	}

	if err := json.Unmarshal([]byte(`{"type":"Image","url":"https://example.com/i.png","selectAction":{"type":"Nope"}}`), &img); err == nil {
		t.Fatalf("expected image selectAction error")
	}

	if err := json.Unmarshal([]byte(`{"type":"Image","url":"not-a-url"}`), &img); err != nil {
		t.Fatalf("expected uri-reference url to be accepted, got: %v", err)
	}
}

func TestImageValidateRejectsShowCardSelectAction(t *testing.T) {
	img := NewImage("https://example.com/cat.png").WithSelectAction(fakeShowCardAction{})
	if err := img.Validate(); err == nil {
		t.Fatalf("expected show card selectAction validation error")
	}
}
