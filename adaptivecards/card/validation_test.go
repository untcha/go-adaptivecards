package card

import (
	"errors"
	"strings"
	"testing"

	a "github.com/untcha/go-adaptivecards/adaptivecards/actions"
	con "github.com/untcha/go-adaptivecards/adaptivecards/containers"
	e "github.com/untcha/go-adaptivecards/adaptivecards/core/element"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
	els "github.com/untcha/go-adaptivecards/adaptivecards/elements"
	"github.com/untcha/go-adaptivecards/adaptivecards/schema"
)

func TestCardValidateNilCard(t *testing.T) {
	var card *Card
	err := card.Validate()
	if err == nil {
		t.Fatalf("expected error for nil card")
	}
	if !strings.Contains(err.Error(), "card is nil") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSchemaValidateNilCard(t *testing.T) {
	err := schema.Validate(nil)
	if err == nil {
		t.Fatalf("expected error for nil card")
	}
	if !strings.Contains(err.Error(), "value is nil") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateSelectActionRejectsShowCard(t *testing.T) {
	err := validateSelectAction(newFakeAction(m.TypeActionShowCard))
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestValidateActionUsesValidatorInterface(t *testing.T) {
	want := errors.New("boom")
	err := validateAction(fakeActionWithError{t: "Action.Custom", err: want})
	if !errors.Is(err, want) {
		t.Fatalf("expected wrapped validation error, got %v", err)
	}
}

func TestValidateElementUsesValidatorInterface(t *testing.T) {
	want := errors.New("boom")
	err := validateElement(fakeElementWithError{t: "CustomElement", err: want})
	if !errors.Is(err, want) {
		t.Fatalf("expected wrapped validation error, got %v", err)
	}
}

func TestCardValidateReportsBackgroundImageError(t *testing.T) {
	card := NewCard()
	card.BackgroundImage = m.BackgroundImageURLUnchecked("not-a-url")

	err := card.Validate()
	if err == nil {
		t.Fatalf("expected validation error")
	}
	if !strings.Contains(err.Error(), "backgroundImage") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCardValidationHelperBranches(t *testing.T) {
	if err := validateSelectAction(newFakeAction(m.TypeActionOpenURL)); err != nil {
		t.Fatalf("unexpected valid selectAction error: %v", err)
	}

	if err := validateActions([]a.Action{fakeActionWithError{t: "Action.X"}}); err == nil {
		t.Fatalf("expected validateActions error")
	}

	if err := validateElements([]e.Element{fakeElementWithError{t: "Element.X"}}); err == nil {
		t.Fatalf("expected validateElements error")
	}

	okAct, err := a.NewActionOpenURL("Open", "https://example.com")
	if err != nil {
		t.Fatalf("setup action failed: %v", err)
	}
	if err := validateAction(okAct); err != nil {
		t.Fatalf("unexpected validateAction error: %v", err)
	}
	if err := validateAction(&okAct); err != nil {
		t.Fatalf("unexpected validateAction ptr error: %v", err)
	}

	img := els.NewImage("https://example.com/img.png").WithSelectAction(newFakeAction(m.TypeActionShowCard))
	if err := validateElement(img); err == nil {
		t.Fatalf("expected image selectAction error")
	}
}

func TestCardValidateAllowsContainerRtl(t *testing.T) {
	card := NewCard().AddElement(func() con.Container {
		c := con.NewContainer(els.NewTextBlock("rtl"))
		b := true
		c.Rtl = &b
		return c
	}())

	if err := card.Validate(); err != nil {
		t.Fatalf("unexpected validation error for container.rtl: %v", err)
	}
}

func TestCardValidateAllowsTableCellRtl(t *testing.T) {
	card := NewCard().AddElement(func() con.Table {
		b := true
		cell := con.NewTableCell(els.NewTextBlock("rtl"))
		cell.Rtl = &b
		row := con.NewTableRow(cell)
		return con.NewTable().AddRow(row)
	}())

	if err := card.Validate(); err != nil {
		t.Fatalf("unexpected validation error for tableCell.rtl: %v", err)
	}
}
