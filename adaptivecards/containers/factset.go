package containers

import (
	"encoding/json"
	"fmt"

	e "github.com/untcha/go-adaptivecards/adaptivecards/core/element"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// FactSet
// The FactSet element displays a series of facts (i.e. name/value pairs) in a tabular form.
// See: https://adaptivecards.io/explorer/FactSet.html
type FactSet struct {
	e.ElementBase              // Embedding e.ElementBase to include common element fields
	Type          m.TypeString `json:"type"`            // Version 1.0
	Facts         []Fact       `json:"facts,omitempty"` // Version 1.0
}

func NewFactSet(facts ...Fact) FactSet {
	return FactSet{
		Type:  m.TypeFactSet,
		Facts: facts,
	}
}

func (f FactSet) GetType() m.TypeString { return m.TypeFactSet }

func (f FactSet) Validate() error {
	if err := f.validateElementBase(); err != nil {
		return err
	}
	if len(f.Facts) == 0 {
		return fmt.Errorf("factSet.facts is required")
	}
	for i, fact := range f.Facts {
		if err := fact.Validate(); err != nil {
			return fmt.Errorf("factSet.facts[%d]: %w", i, err)
		}
	}
	return nil
}

func (f FactSet) MarshalJSON() ([]byte, error) {
	ff := f
	if ff.Type == "" {
		ff.Type = m.TypeFactSet
	}
	type alias FactSet
	return json.Marshal(alias(ff))
}

func (f *FactSet) UnmarshalJSON(b []byte) error {
	type alias FactSet
	var tmp alias
	if err := json.Unmarshal(b, &tmp); err != nil {
		return fmt.Errorf("factSet: decode: %w", err)
	}
	if tmp.Type != "" && tmp.Type != m.TypeFactSet {
		return fmt.Errorf("factSet.type must be %q (got %q)", m.TypeFactSet, tmp.Type)
	}
	if tmp.Type == "" {
		tmp.Type = m.TypeFactSet
	}
	val := FactSet(tmp)
	if err := val.Validate(); err != nil {
		return err
	}
	*f = val
	return nil
}

func (f FactSet) validateElementBase() error {
	return f.ElementBase.Validate("FactSet")
}

func init() {
	e.RegisterElement(m.TypeFactSet, func() e.Element { return &FactSet{Type: m.TypeFactSet} })
}
