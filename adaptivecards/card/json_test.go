package card

import (
	"encoding/json"
	"strings"
	"testing"

	a "github.com/untcha/go-adaptivecards/adaptivecards/actions"
	con "github.com/untcha/go-adaptivecards/adaptivecards/containers"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
	els "github.com/untcha/go-adaptivecards/adaptivecards/elements"
	inp "github.com/untcha/go-adaptivecards/adaptivecards/inputs"
)

func TestCardMarshalJSONAppliesDefaults(t *testing.T) {
	b, err := json.Marshal(&Card{})
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	s := string(b)
	if !strings.Contains(s, `"type":"AdaptiveCard"`) {
		t.Fatalf("expected default type, got: %s", s)
	}
	if !strings.Contains(s, `"version":"1.5"`) {
		t.Fatalf("expected default version, got: %s", s)
	}
	if !strings.Contains(s, `"$schema":"http://adaptivecards.io/schemas/adaptive-card.json"`) {
		t.Fatalf("expected default schema, got: %s", s)
	}
}

func TestCardUnmarshalJSONWithBodyActionsAndSelectAction(t *testing.T) {
	input := `{
		"type":"AdaptiveCard",
		"version":"1.5",
		"body":[{"type":"TextBlock","text":"hello"}],
		"actions":[{"type":"Action.OpenUrl","url":"https://example.com"}],
		"selectAction":{"type":"Action.OpenUrl","url":"https://example.com"}
	}`

	var c Card
	if err := json.Unmarshal([]byte(input), &c); err != nil {
		t.Fatalf("unmarshal card failed: %v", err)
	}
	if len(c.Body) != 1 {
		t.Fatalf("expected one body element")
	}
	if len(c.Actions) != 1 {
		t.Fatalf("expected one action")
	}
	if c.SelectAction == nil {
		t.Fatalf("expected selectAction")
	}
}

func TestCardUnmarshalJSONErrors(t *testing.T) {
	var c Card
	if err := json.Unmarshal([]byte(`{"type":"Wrong"}`), &c); err == nil {
		t.Fatalf("expected card type error")
	}

	if err := json.Unmarshal([]byte(`{"type":"AdaptiveCard","body":[{"type":"Nope"}]}`), &c); err == nil {
		t.Fatalf("expected body element type error")
	}

	if err := json.Unmarshal([]byte(`{"type":"AdaptiveCard","actions":[{"type":"Nope"}]}`), &c); err == nil {
		t.Fatalf("expected action type error")
	}

	if err := json.Unmarshal([]byte(`{"type":"AdaptiveCard","selectAction":{"type":"Nope"}}`), &c); err == nil {
		t.Fatalf("expected selectAction error")
	}
}

func TestCardValidateSerializeDeserializeWithLayoutElements(t *testing.T) {
	submit := a.NewActionSubmit("Submit", map[string]any{"meta": "x"})
	openURL, err := a.NewActionOpenURL("Open", "https://example.com")
	if err != nil {
		t.Fatalf("setup action failed: %v", err)
	}

	column := con.NewColumn(
		els.NewTextBlock("Left"),
		els.NewImage("https://example.com/image.png"),
	)
	column.Width = "stretch"
	column.SelectAction = openURL

	columnSet := con.NewColumnSet(column)
	columnSet.HorizontalAlignment = m.HAlignLeft

	container := con.NewContainer(
		els.NewTextBlock("Header"),
		columnSet,
	)
	container.Style = m.ContainerStyleDefault

	actionSet := con.NewActionSet(submit)

	card := NewCard().
		AddElement(container).
		AddElement(actionSet)

	if err := card.Validate(); err != nil {
		t.Fatalf("unexpected card validate error: %v", err)
	}

	b, err := json.Marshal(card)
	if err != nil {
		t.Fatalf("marshal card failed: %v", err)
	}

	var decoded Card
	if err := json.Unmarshal(b, &decoded); err != nil {
		t.Fatalf("unmarshal card failed: %v", err)
	}
	if len(decoded.Body) != 2 {
		t.Fatalf("expected 2 body elements, got %d", len(decoded.Body))
	}
}

func TestCardValidateSerializeDeserializeWithInputAndFactSetElements(t *testing.T) {
	card := NewCard().
		AddElement(inp.NewInputText("subject")).
		AddElement(inp.NewInputChoiceSet("severity", inp.NewInputChoice("High", "high"))).
		AddElement(inp.NewInputToggle("notify", "Notify")).
		AddElement(con.NewFactSet(con.NewFact("Region", "us-east")))

	if err := card.Validate(); err != nil {
		t.Fatalf("unexpected card validate error: %v", err)
	}

	b, err := json.Marshal(card)
	if err != nil {
		t.Fatalf("marshal card failed: %v", err)
	}
	var decoded Card
	if err := json.Unmarshal(b, &decoded); err != nil {
		t.Fatalf("unmarshal card failed: %v", err)
	}
	if len(decoded.Body) != 4 {
		t.Fatalf("expected 4 body elements, got %d", len(decoded.Body))
	}
}

func TestCardValidateSerializeDeserializeWithAdvancedInputAndRichContentElements(t *testing.T) {
	min := 1.0
	max := 10.0
	val := 3.0

	card := NewCard().
		AddElement(con.NewImageSet(els.NewImage("https://example.com/hero.png"))).
		AddElement(els.NewRichTextBlock(
			els.NewTextRun("Status: "),
			els.TextRun{Text: "OK", Color: m.ColorGood},
		)).
		AddElement(func() inp.InputDate {
			in := inp.NewInputDate("dueDate")
			in.Value = "2026-03-01"
			return in
		}()).
		AddElement(func() inp.InputTime {
			in := inp.NewInputTime("dueTime")
			in.Value = "12:15"
			return in
		}()).
		AddElement(func() inp.InputNumber {
			in := inp.NewInputNumber("amount")
			in.Min = &min
			in.Max = &max
			in.Value = &val
			return in
		}()).
		AddAction(a.NewActionToggleVisibility("Toggle Details", a.TargetElement{ElementID: "details"}))

	if err := card.Validate(); err != nil {
		t.Fatalf("unexpected card validate error: %v", err)
	}

	b, err := json.Marshal(card)
	if err != nil {
		t.Fatalf("marshal card failed: %v", err)
	}
	var decoded Card
	if err := json.Unmarshal(b, &decoded); err != nil {
		t.Fatalf("unmarshal card failed: %v", err)
	}
	if len(decoded.Body) != 5 {
		t.Fatalf("expected 5 body elements, got %d", len(decoded.Body))
	}
	if len(decoded.Actions) != 1 {
		t.Fatalf("expected 1 action, got %d", len(decoded.Actions))
	}
	if decoded.Actions[0].GetType() != m.TypeActionToggleVisibility {
		t.Fatalf("expected action type %q, got %q", m.TypeActionToggleVisibility, decoded.Actions[0].GetType())
	}
}
