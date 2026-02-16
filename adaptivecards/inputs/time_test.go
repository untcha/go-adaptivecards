package inputs

import (
	"encoding/json"
	"testing"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

func TestInputTimeValidateAndUnmarshal(t *testing.T) {
	in := NewInputTime("startTime")
	in.Min = "09:00"
	in.Max = "18:00"
	in.Value = "13:15"
	if err := in.Validate(); err != nil {
		t.Fatalf("unexpected validation error: %v", err)
	}

	var decoded InputTime
	if err := json.Unmarshal([]byte(`{"type":"Input.Time","id":"t","min":"09:00","max":"18:00","value":"10:30"}`), &decoded); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if decoded.Type != m.TypeInputTime {
		t.Fatalf("expected input time type")
	}

	schemaTolerant := NewInputTime("x")
	schemaTolerant.Min = "not-a-time"
	schemaTolerant.Max = "still-not-a-time"
	schemaTolerant.Value = "25:99"
	if err := schemaTolerant.Validate(); err != nil {
		t.Fatalf("expected schema-tolerant validation, got error: %v", err)
	}
}
