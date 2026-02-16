package inputs

import (
	"encoding/json"
	"testing"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

func TestInputNumberValidateAndUnmarshal(t *testing.T) {
	min := 1.0
	max := 10.0
	value := 5.0

	in := NewInputNumber("count")
	in.Min = &min
	in.Max = &max
	in.Value = &value
	if err := in.Validate(); err != nil {
		t.Fatalf("unexpected validation error: %v", err)
	}

	var decoded InputNumber
	if err := json.Unmarshal([]byte(`{"type":"Input.Number","id":"n","min":0,"max":10,"value":3}`), &decoded); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if decoded.Type != m.TypeInputNumber {
		t.Fatalf("expected input number type")
	}

	outOfRange := NewInputNumber("n")
	outOfRange.Min = &min
	outOfRange.Max = &max
	badValue := 11.0
	outOfRange.Value = &badValue
	if err := outOfRange.Validate(); err != nil {
		t.Fatalf("expected schema-tolerant validation, got error: %v", err)
	}
}
