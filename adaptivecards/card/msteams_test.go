package card

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
	els "github.com/untcha/go-adaptivecards/adaptivecards/elements"
)

func TestSetFullWidthMarshalsMSTeamsWidth(t *testing.T) {
	c, err := NewCard().
		AddElement(els.NewTextBlock("hi")).
		SetFullWidth(true).
		Build()
	if err != nil {
		t.Fatalf("build failed: %v", err)
	}

	b, err := json.Marshal(c)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	if !strings.Contains(string(b), `"msteams":{"width":"Full"}`) {
		t.Fatalf("expected msteams width Full, got: %s", b)
	}
}

func TestDefaultCardOmitsMSTeams(t *testing.T) {
	b, err := json.Marshal(NewCard())
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	if strings.Contains(string(b), "msteams") {
		t.Fatalf("expected no msteams key on default card, got: %s", b)
	}
}

func TestMSTeamsRoundTrip(t *testing.T) {
	in := `{"type":"AdaptiveCard","version":"1.5","msteams":{"width":"Full"}}`

	var c Card
	if err := json.Unmarshal([]byte(in), &c); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if c.MSTeams == nil {
		t.Fatalf("expected msteams to be decoded")
	}
	if c.MSTeams.Width != MSTeamsWidthFull {
		t.Fatalf("expected width Full, got %q", c.MSTeams.Width)
	}
}

func TestValidateAcceptsFullWidthCard(t *testing.T) {
	c, err := NewCard().
		AddElement(els.NewTextBlock("hi")).
		SetFullWidth(true).
		Build()
	if err != nil {
		t.Fatalf("build failed: %v", err)
	}
	// The embedded schema sets additionalProperties:false on AdaptiveCard,
	// so this guards that msteams is stripped before schema validation.
	if err := c.Validate(); err != nil {
		t.Fatalf("expected full-width card to validate, got: %v", err)
	}
}

func TestValidateRejectsInvalidMSTeamsWidth(t *testing.T) {
	c := NewCard().AddElement(els.NewTextBlock("hi"))
	c.MSTeams = &MSTeams{Width: "Wide"}

	err := c.Validate()
	if err == nil {
		t.Fatalf("expected validation error for invalid width")
	}
	if !errors.Is(err, m.ErrInvalidEnum) {
		t.Fatalf("expected ErrInvalidEnum, got: %v", err)
	}
	if !strings.Contains(err.Error(), "msteams.width") {
		t.Fatalf("expected msteams.width in error, got: %v", err)
	}
}

func TestSetFullWidthFalseClearsWidth(t *testing.T) {
	c := NewCard().SetFullWidth(true).SetFullWidth(false)
	if c.MSTeams == nil {
		t.Fatalf("expected MSTeams object to remain after clearing width")
	}
	if c.MSTeams.Width != "" {
		t.Fatalf("expected width cleared, got %q", c.MSTeams.Width)
	}

	b, err := json.Marshal(c)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	if strings.Contains(string(b), `"width"`) {
		t.Fatalf("expected no width key, got: %s", b)
	}
}

func TestSetMSTeamsRejectsInvalidWidth(t *testing.T) {
	_, err := NewCard().SetMSTeams(MSTeams{Width: "Wide"}).Build()
	if err == nil {
		t.Fatalf("expected build error for invalid msteams width")
	}
	if !errors.Is(err, m.ErrInvalidEnum) {
		t.Fatalf("expected ErrInvalidEnum, got: %v", err)
	}
}
