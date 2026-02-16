package card

import (
	"errors"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

type fakeAction struct {
	t m.TypeString
}

func (f fakeAction) GetType() m.TypeString    { return f.t }
func (f fakeAction) Validate() error          { return nil }
func newFakeAction(t m.TypeString) fakeAction { return fakeAction{t: t} }

type fakeActionWithError struct {
	t   m.TypeString
	err error
}

func (f fakeActionWithError) GetType() m.TypeString { return f.t }
func (f fakeActionWithError) Validate() error {
	if f.err == nil {
		return errors.New("validation failed")
	}
	return f.err
}

type fakeElementWithError struct {
	t   m.TypeString
	err error
}

func (f fakeElementWithError) GetType() m.TypeString { return f.t }
func (f fakeElementWithError) Validate() error {
	if f.err == nil {
		return errors.New("validation failed")
	}
	return f.err
}
