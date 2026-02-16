package containers

import (
	"encoding/json"
	"strings"
	"testing"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
	els "github.com/untcha/go-adaptivecards/adaptivecards/elements"
)

func TestImageSetValidateAndUnmarshal(t *testing.T) {
	set := NewImageSet(
		els.NewImage("https://example.com/1.png"),
		els.NewImage("https://example.com/2.png"),
	)
	set.ImageSize = m.ImageSizeAuto // normalized to medium by schema behavior
	if err := set.Validate(); err != nil {
		t.Fatalf("unexpected imageset validation error: %v", err)
	}

	b, err := json.Marshal(set)
	if err != nil {
		t.Fatalf("marshal imageset failed: %v", err)
	}
	if !strings.Contains(string(b), `"imageSize":"medium"`) {
		t.Fatalf("expected imageSize normalization to medium, got: %s", string(b))
	}

	var decoded ImageSet
	if err := json.Unmarshal([]byte(`{"type":"ImageSet","images":[{"type":"Image","url":"https://example.com/a.png"}],"imageSize":"stretch"}`), &decoded); err != nil {
		t.Fatalf("unexpected imageset unmarshal error: %v", err)
	}
	if decoded.ImageSize != m.ImageSizeMedium {
		t.Fatalf("expected decoded imageSize medium, got %q", decoded.ImageSize)
	}
}

func TestImageSetValidationErrors(t *testing.T) {
	empty := NewImageSet()
	if err := empty.Validate(); err == nil {
		t.Fatalf("expected images required error")
	}

	invalid := NewImageSet(els.NewImage("https://example.com/1.png"))
	invalid.ImageSize = m.ImageSize("invalid")
	if err := invalid.Validate(); err == nil {
		t.Fatalf("expected invalid imageSize error")
	}
}
