package element

import (
	"encoding/json"
	"strings"
	"testing"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

const (
	testFactoryTypeA = m.TypeString("Test.FactoryElementA")
	testFactoryTypeB = m.TypeString("Test.FactoryElementB")
)

type testElement struct {
	Type  m.TypeString `json:"type"`
	Value string       `json:"value,omitempty"`
}

func (t testElement) GetType() m.TypeString { return t.Type }

func init() {
	RegisterElement(testFactoryTypeA, func() Element {
		return &testElement{Type: testFactoryTypeA}
	})
	RegisterElement(testFactoryTypeB, func() Element {
		return &testElement{Type: testFactoryTypeB}
	})
}

func TestUnmarshalElementErrors(t *testing.T) {
	if _, err := UnmarshalElement([]byte(`{"text":"missing type"}`)); err == nil {
		t.Fatalf("expected missing type error")
	}

	if _, err := UnmarshalElement([]byte(`{"type":"Not.Registered"}`)); err == nil {
		t.Fatalf("expected unknown type error")
	}

	if _, err := UnmarshalElement([]byte(`{`)); err == nil {
		t.Fatalf("expected invalid json error")
	}
}

func TestUnmarshalElementsSliceIncludesIndexOnError(t *testing.T) {
	var raw []json.RawMessage
	if err := json.Unmarshal(
		[]byte(`[{"type":"Test.FactoryElementA","value":"ok"},{"type":"Not.Registered"}]`),
		&raw,
	); err != nil {
		t.Fatalf("setup unmarshal failed: %v", err)
	}

	_, err := UnmarshalElementsSlice(raw)
	if err == nil {
		t.Fatalf("expected indexed error")
	}
	if !strings.Contains(err.Error(), "elements[1]") {
		t.Fatalf("expected error to include failing index, got: %v", err)
	}
}

func TestRegisterElementPanicsOnNilFactory(t *testing.T) {
	typ := m.TypeString("Test.NilFactoryElement")
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic for nil factory")
		}
	}()
	RegisterElement(typ, nil)
}

func TestRegisterElementPanicsOnDuplicate(t *testing.T) {
	typ := m.TypeString("Test.DuplicateElement")
	RegisterElement(typ, func() Element { return &testElement{} })

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic for duplicate registration")
		}
	}()
	RegisterElement(typ, func() Element { return &testElement{} })
}

func TestElementFactoryDeserializeRegisteredTestElementTypes(t *testing.T) {
	cases := []string{
		`{"type":"Test.FactoryElementA","value":"a"}`,
		`{"type":"Test.FactoryElementB","value":"b"}`,
	}

	for _, raw := range cases {
		el, err := UnmarshalElement([]byte(raw))
		if err != nil {
			t.Fatalf("unexpected unmarshal element error for %s: %v", raw, err)
		}
		if el.GetType() == "" {
			t.Fatalf("expected non-empty type")
		}
	}
}
