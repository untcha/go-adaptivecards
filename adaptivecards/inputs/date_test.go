package inputs

import (
	"encoding/json"
	"testing"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

func TestInputDateValidateAndUnmarshal(t *testing.T) {
	in := NewInputDate("startDate")
	in.Min = "2026-01-01"
	in.Max = "2026-12-31"
	in.Value = "2026-06-15"
	if err := in.Validate(); err != nil {
		t.Fatalf("unexpected validation error: %v", err)
	}

	var decoded InputDate
	if err := json.Unmarshal([]byte(`{"type":"Input.Date","id":"d","min":"2026-01-01","max":"2026-12-31","value":"2026-10-01"}`), &decoded); err != nil {
		t.Fatalf("unexpected unmarshal error: %v", err)
	}
	if decoded.Type != m.TypeInputDate {
		t.Fatalf("expected input date type")
	}

	schemaTolerant := NewInputDate("x")
	schemaTolerant.Min = "not-a-date"
	schemaTolerant.Max = "definitely-not-a-date"
	schemaTolerant.Value = "also-not-a-date"
	if err := schemaTolerant.Validate(); err != nil {
		t.Fatalf("expected schema-tolerant validation, got error: %v", err)
	}
}
