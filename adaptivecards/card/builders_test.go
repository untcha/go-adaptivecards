package card

import (
	"strings"
	"testing"

	a "github.com/untcha/go-adaptivecards/adaptivecards/actions"
	tbl "github.com/untcha/go-adaptivecards/adaptivecards/containers"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
	els "github.com/untcha/go-adaptivecards/adaptivecards/elements"
)

func TestSetSelectActionNilSetsBuildErr(t *testing.T) {
	card := NewCard().SetSelectAction(nil)

	if card == nil {
		t.Fatalf("expected non-nil card")
	}

	_, err := card.Build()
	if err == nil {
		t.Fatalf("expected build error for nil selectAction")
	}
	if !strings.Contains(err.Error(), "selectAction cannot be nil") {
		t.Fatalf("expected nil selectAction error, got: %v", err)
	}
}

func TestSetSelectActionRejectsShowCardType(t *testing.T) {
	card := NewCard().SetSelectAction(newFakeAction(m.TypeActionShowCard))
	_, err := card.Build()
	if err == nil {
		t.Fatalf("expected build error for Action.ShowCard")
	}
	if !strings.Contains(err.Error(), "Action.ShowCard is not supported") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestBuilderBackgroundImageAndCardSettings(t *testing.T) {
	card := NewCard().
		SetBackgroundImageURL("https://example.com/bg.png").
		SetCardMinHeight("240px").
		SetCardRtl(true).
		SetVerticalContentAlignment(m.VContentAlignCenter)

	if card.CardBuildErr() != nil {
		t.Fatalf("unexpected build error: %v", card.CardBuildErr())
	}
	if card.BackgroundImage == nil {
		t.Fatalf("expected background image")
	}
	if card.MinHeight != "240px" {
		t.Fatalf("unexpected min height: %s", card.MinHeight)
	}
	if card.Rtl == nil || !*card.Rtl {
		t.Fatalf("expected rtl=true")
	}
	if card.VerticalContentAlignment != m.VContentAlignCenter {
		t.Fatalf("unexpected vertical content alignment: %s", card.VerticalContentAlignment)
	}
}

func TestBuilderValidationErrorsAreStored(t *testing.T) {
	card := NewCard().SetCardMinHeight("invalid")
	if card.CardBuildErr() == nil {
		t.Fatalf("expected build error")
	}

	card = NewCard().SetVerticalContentAlignment("bad")
	if card.CardBuildErr() == nil {
		t.Fatalf("expected build error for invalid vertical alignment")
	}

	card = NewCard().SetBackgroundImageURL("invalid-url")
	if card.CardBuildErr() == nil {
		t.Fatalf("expected build error for invalid background image url")
	}
}

func TestBackgroundImageBuilderConvenienceMethods(t *testing.T) {
	card := NewCard().
		SetBackgroundImage(m.BackgroundImage{URL: "https://example.com/a.png"}).
		SetBackgroundImageWithFillMode("https://example.com/b.png", m.ImageFillCover).
		SetBackgroundImageWithAlignment("https://example.com/c.png", m.HAlignCenter, m.VAlignCenter).
		SetBackgroundImageComplete("https://example.com/d.png", m.ImageFillRepeat, m.HAlignLeft, m.VAlignTop).
		SetBackgroundImageCover("https://example.com/e.png").
		SetBackgroundImageRepeat("https://example.com/f.png").
		SetBackgroundImageCenterCover("https://example.com/g.png")

	if card.CardBuildErr() != nil {
		t.Fatalf("unexpected build error: %v", card.CardBuildErr())
	}
	if card.BackgroundImage == nil {
		t.Fatalf("expected background image to be set")
	}
}

func TestBuilderErrorPreservationAndNilReceiverBranches(t *testing.T) {
	var nilCard *Card
	if nilCard.SetSelectAction(newFakeAction(m.TypeActionOpenURL)) != nil {
		t.Fatalf("expected nil receiver to stay nil")
	}
	if nilCard.Title("x") != nil {
		t.Fatalf("expected nil receiver for Title")
	}
	if nilCard.AddElement(els.NewTextBlock("x")) != nil {
		t.Fatalf("expected nil receiver for AddElement")
	}
	if nilCard.AddTextBlock(els.NewTextBlock("x")) != nil {
		t.Fatalf("expected nil receiver for AddTextBlock")
	}
	if nilCard.AddAction(newFakeAction(m.TypeActionOpenURL)) != nil {
		t.Fatalf("expected nil receiver for AddAction")
	}
	if nilCard.AddTable(tbl.NewTable()) != nil {
		t.Fatalf("expected nil receiver for AddTable")
	}
	if nilCard.AddImage("https://example.com/image.png") != nil {
		t.Fatalf("expected nil receiver for AddImage")
	}
	if nilCard.AddContainer(els.NewTextBlock("x")) != nil {
		t.Fatalf("expected nil receiver for AddContainer")
	}
	if nilCard.SetBackgroundImageURL("https://example.com/bg.png") != nil {
		t.Fatalf("expected nil receiver for SetBackgroundImageURL")
	}
	if nilCard.SetBackgroundImage(m.BackgroundImage{URL: "https://example.com/bg.png"}) != nil {
		t.Fatalf("expected nil receiver for SetBackgroundImage")
	}
	if nilCard.SetCardMinHeight("10px") != nil {
		t.Fatalf("expected nil receiver for SetCardMinHeight")
	}
	if nilCard.SetCardRtl(true) != nil {
		t.Fatalf("expected nil receiver for SetCardRtl")
	}
	if nilCard.SetVerticalContentAlignment(m.VContentAlignTop) != nil {
		t.Fatalf("expected nil receiver for SetVerticalContentAlignment")
	}
	if nilCard.SetLang("de-DE") != nil {
		t.Fatalf("expected nil receiver for SetLang")
	}
	if nilCard.SetSpeak("hello") != nil {
		t.Fatalf("expected nil receiver for SetSpeak")
	}
	if nilCard.SetFallbackText("fallback") != nil {
		t.Fatalf("expected nil receiver for SetFallbackText")
	}

	c := NewCard().SetCardMinHeight("bad")
	before := c.CardBuildErr()
	c.SetSelectAction(newFakeAction(m.TypeActionOpenURL))
	if c.CardBuildErr() == nil || c.CardBuildErr() != before {
		t.Fatalf("expected existing buildErr to be preserved")
	}
	c.SetLang("de-DE").SetSpeak("ignored").SetFallbackText("ignored").AddImage("https://example.com/image.png").AddContainer(els.NewTextBlock("ignored"))
	if c.Lang != "" {
		t.Fatalf("expected lang unchanged when buildErr is set")
	}
	if c.Speak != "" {
		t.Fatalf("expected speak unchanged when buildErr is set")
	}
	if c.FallbackText != "" {
		t.Fatalf("expected fallbackText unchanged when buildErr is set")
	}
	if len(c.Body) != 0 {
		t.Fatalf("expected body unchanged when buildErr is set")
	}

	c = NewCard().SetBackgroundImage(m.BackgroundImage{URL: ""})
	if c.CardBuildErr() == nil {
		t.Fatalf("expected invalid background image object error")
	}

	if NewCard().SetBackgroundImageWithFillMode("bad-url", m.ImageFillCover).CardBuildErr() == nil {
		t.Fatalf("expected invalid url error for SetBackgroundImageWithFillMode")
	}
	if NewCard().SetBackgroundImageWithAlignment("bad-url", m.HAlignCenter, m.VAlignCenter).CardBuildErr() == nil {
		t.Fatalf("expected invalid url error for SetBackgroundImageWithAlignment")
	}
	if NewCard().SetBackgroundImageComplete("bad-url", m.ImageFillCover, m.HAlignCenter, m.VAlignCenter).CardBuildErr() == nil {
		t.Fatalf("expected invalid url error for SetBackgroundImageComplete")
	}

	card := NewCard().SetCardRtl(true)
	if card.Rtl == nil || !*card.Rtl {
		t.Fatalf("expected rtl set")
	}
}

func TestBuildAndCardBuildErrNilReceiver(t *testing.T) {
	var nilCard *Card

	if err := nilCard.CardBuildErr(); err == nil {
		t.Fatalf("expected error for nil CardBuildErr receiver")
	}

	gotCard, err := nilCard.Build()
	if err == nil {
		t.Fatalf("expected error for nil Build receiver")
	}
	if gotCard != nil {
		t.Fatalf("expected nil card from nil receiver Build")
	}
}

func TestBuilderAddersSetTypes(t *testing.T) {
	tb := els.TextBlock{Text: "hello"}
	table := tbl.Table{}
	act, err := a.NewActionOpenURL("open", "https://example.com")
	if err != nil {
		t.Fatalf("setup action failed: %v", err)
	}
	card := NewCard().Title("title").AddTextBlock(tb).AddTable(table).AddElement(els.NewImage("https://example.com/i.png")).AddAction(act)

	if len(card.Body) != 4 {
		t.Fatalf("expected 4 body elements, got %d", len(card.Body))
	}
	if len(card.Actions) != 1 {
		t.Fatalf("expected 1 action")
	}
}

func TestAdditionalConvenienceBuilders(t *testing.T) {
	card := NewCard().
		SetLang("de-DE").
		SetSpeak("Hallo Welt").
		SetFallbackText("Fallback plain text").
		AddImage("https://example.com/image.png").
		AddContainer(els.NewTextBlock("inside"))

	if card.CardBuildErr() != nil {
		t.Fatalf("unexpected build error: %v", card.CardBuildErr())
	}
	if card.Lang != "de-DE" {
		t.Fatalf("expected lang to be set")
	}
	if card.Speak != "Hallo Welt" {
		t.Fatalf("expected speak to be set")
	}
	if card.FallbackText != "Fallback plain text" {
		t.Fatalf("expected fallbackText to be set")
	}
	if len(card.Body) != 2 {
		t.Fatalf("expected 2 body elements, got %d", len(card.Body))
	}
	if card.Body[0].GetType() != m.TypeImage {
		t.Fatalf("expected first body element to be Image")
	}
	if card.Body[1].GetType() != m.TypeContainer {
		t.Fatalf("expected second body element to be Container")
	}
}

func TestActionConvenienceBuilders(t *testing.T) {
	card := NewCard().
		AddActionOpenURL("Open", "https://example.com").
		AddActionSubmit("Send", map[string]any{"x": 1}).
		AddActionToggleVisibility("Toggle", a.TargetElement{ElementID: "details"})

	if card.CardBuildErr() != nil {
		t.Fatalf("unexpected build error: %v", card.CardBuildErr())
	}
	if len(card.Actions) != 3 {
		t.Fatalf("expected 3 actions, got %d", len(card.Actions))
	}
	if card.Actions[0].GetType() != m.TypeActionOpenURL {
		t.Fatalf("expected first action type %q", m.TypeActionOpenURL)
	}
	if card.Actions[1].GetType() != m.TypeActionSubmit {
		t.Fatalf("expected second action type %q", m.TypeActionSubmit)
	}
	if card.Actions[2].GetType() != m.TypeActionToggleVisibility {
		t.Fatalf("expected third action type %q", m.TypeActionToggleVisibility)
	}
}

func TestActionConvenienceBuildersErrorAndNilBranches(t *testing.T) {
	card := NewCard().AddActionOpenURL("Bad", "not-a-url")
	if card.CardBuildErr() == nil {
		t.Fatalf("expected build error for invalid open url")
	}

	before := card.CardBuildErr()
	card.AddActionSubmit("Ignored", nil)
	if card.CardBuildErr() != before {
		t.Fatalf("expected existing buildErr to be preserved")
	}

	var nilCard *Card
	if nilCard.AddActionOpenURL("Open", "https://example.com") != nil {
		t.Fatalf("expected nil receiver for AddActionOpenURL")
	}
	if nilCard.AddActionSubmit("Send", nil) != nil {
		t.Fatalf("expected nil receiver for AddActionSubmit")
	}
	if nilCard.AddActionToggleVisibility("Toggle", a.TargetElement{ElementID: "x"}) != nil {
		t.Fatalf("expected nil receiver for AddActionToggleVisibility")
	}
}
