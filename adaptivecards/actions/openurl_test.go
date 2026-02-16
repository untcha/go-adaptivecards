package actions

import (
	"encoding/json"
	"strings"
	"testing"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

func TestActionOpenURLConstructValidateAndMutate(t *testing.T) {
	act, err := NewActionOpenURL("Open", "https://example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if act.GetType() != m.TypeActionOpenURL {
		t.Fatalf("unexpected action type: %s", act.GetType())
	}

	act = act.WithTitle("Open Site").StylePositive().ModePrimary().Enable()
	act, err = act.WithIconURL("https://example.com/icon.png")
	if err != nil {
		t.Fatalf("unexpected icon url error: %v", err)
	}
	act, err = act.WithURL("https://example.com/new")
	if err != nil {
		t.Fatalf("unexpected url error: %v", err)
	}
	if err := act.Validate(); err != nil {
		t.Fatalf("unexpected validation error: %v", err)
	}
}

func TestActionOpenURLValidationErrors(t *testing.T) {
	_, err := NewActionOpenURL("Open", "not-a-url")
	if err == nil {
		t.Fatalf("expected constructor error")
	}

	act := ActionOpenURL{Type: m.TypeActionOpenURL}
	if err := act.Validate(); err == nil {
		t.Fatalf("expected missing url error")
	}

	act = ActionOpenURL{Type: m.TypeActionOpenURL, URL: "https://example.com"}
	_, err = act.WithIconURL("mailto:test@example.com")
	if err == nil {
		t.Fatalf("expected icon url validation error")
	}
}

func TestActionOpenURLMarshalUnmarshal(t *testing.T) {
	act := ActionOpenURL{URL: "https://example.com"}
	b, err := json.Marshal(act)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	if !strings.Contains(string(b), `"type":"Action.OpenUrl"`) {
		t.Fatalf("expected default type in json, got: %s", string(b))
	}

	var decoded ActionOpenURL
	if err := json.Unmarshal([]byte(`{"type":"Action.OpenUrl","url":"https://example.com"}`), &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if decoded.Type != m.TypeActionOpenURL {
		t.Fatalf("unexpected decoded type: %s", decoded.Type)
	}

	if err := json.Unmarshal([]byte(`{"type":"Action.Submit","url":"https://example.com"}`), &decoded); err == nil {
		t.Fatalf("expected type mismatch error")
	}
	if err := json.Unmarshal([]byte(`{"type":"Action.OpenUrl","url":"not-a-url"}`), &decoded); err == nil {
		t.Fatalf("expected invalid url error")
	}
	if err := json.Unmarshal([]byte(`{"type":"Action.OpenUrl","url":"https://example.com","iconUrl":"mailto:test@example.com"}`), &decoded); err == nil {
		t.Fatalf("expected invalid iconUrl error")
	}
}

func TestActionFactoryUnmarshal(t *testing.T) {
	a, err := UnmarshalAction([]byte(`{"type":"Action.OpenUrl","url":"https://example.com"}`))
	if err != nil {
		t.Fatalf("unexpected unmarshal action error: %v", err)
	}
	if a.GetType() != m.TypeActionOpenURL {
		t.Fatalf("unexpected action type: %s", a.GetType())
	}

	if _, err := UnmarshalAction([]byte(`{"url":"https://example.com"}`)); err == nil {
		t.Fatalf("expected missing type error")
	}
	if _, err := UnmarshalAction([]byte(`{"type":"Action.Unknown","url":"https://example.com"}`)); err == nil {
		t.Fatalf("expected unknown type error")
	}
}

func TestUnmarshalActionsSlice(t *testing.T) {
	var raw []json.RawMessage
	if err := json.Unmarshal([]byte(`[{"type":"Action.OpenUrl","url":"https://example.com"}]`), &raw); err != nil {
		t.Fatalf("setup unmarshal failed: %v", err)
	}

	actions, err := UnmarshalActionsSlice(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(actions) != 1 {
		t.Fatalf("expected one action")
	}
}

func TestActionFallbackJSON(t *testing.T) {
	drop := ActionFallbackDropOption()
	b, err := json.Marshal(drop)
	if err != nil {
		t.Fatalf("marshal drop fallback failed: %v", err)
	}
	if string(b) != `"drop"` {
		t.Fatalf("unexpected drop fallback json: %s", string(b))
	}

	f := ActionFallbackContent(ActionOpenURL{Type: m.TypeActionOpenURL, URL: "https://example.com"})
	b, err = json.Marshal(f)
	if err != nil {
		t.Fatalf("marshal content fallback failed: %v", err)
	}
	if !strings.Contains(string(b), `"Action.OpenUrl"`) {
		t.Fatalf("unexpected fallback action json: %s", string(b))
	}

	var decoded ActionFallback
	if err := json.Unmarshal([]byte(`"drop"`), &decoded); err != nil {
		t.Fatalf("unmarshal fallback option failed: %v", err)
	}
	if decoded.Option == nil || *decoded.Option != FallbackDrop {
		t.Fatalf("expected drop option")
	}
	if err := json.Unmarshal([]byte(`"invalid"`), &decoded); err == nil {
		t.Fatalf("expected invalid fallback option error")
	}
}

func TestActionOpenURLAdditionalConvenienceMethods(t *testing.T) {
	a, err := NewActionOpenURL("Open", "https://example.com")
	if err != nil {
		t.Fatalf("setup error: %v", err)
	}
	a = a.WithStyle(ActionStyleDefault).
		StyleDefault().
		StyleDestructive().
		WithTooltip("tip").
		WithIsEnabled(true).
		Disable().
		WithMode(ActionModePrimary).
		ModeSecondary().
		WithRequires(map[string]string{"a": "1"}).
		AddRequire("b", "2")

	if a.Requires["a"] != "1" || a.Requires["b"] != "2" {
		t.Fatalf("expected requires map to be populated")
	}
	if err := a.Validate(); err != nil {
		t.Fatalf("unexpected validation error: %v", err)
	}
}
