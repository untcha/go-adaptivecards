package containers

import (
	"fmt"
	"strings"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// Fact
// Describes a Fact in a FactSet as a key/value pair.
// See: https://adaptivecards.io/explorer/Fact.html
type Fact struct {
	Type  m.TypeString `json:"type,omitempty"`
	Title string       `json:"title"` // Version 1.0
	Value string       `json:"value"` // Version 1.0
}

func NewFact(title, value string) Fact {
	return Fact{
		Type:  "Fact",
		Title: title,
		Value: value,
	}
}

func (f Fact) Validate() error {
	if strings.TrimSpace(f.Title) == "" {
		return fmt.Errorf("fact.title is required")
	}
	if strings.TrimSpace(f.Value) == "" {
		return fmt.Errorf("fact.value is required")
	}
	return nil
}
