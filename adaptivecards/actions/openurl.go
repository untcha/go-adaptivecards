package actions

import (
	"encoding/json"
	"fmt"

	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

var ErrInvalidActionOpenURLURL = m.ErrInvalidActionOpenURLURL

// Action.OpenUrl
// When invoked, show the given url either by launching it in an external web browser
// or showing within an embedded web browser.
// See: https://adaptivecards.io/explorer/Action.OpenUrl.html
type ActionOpenURL struct {
	ActionBase
	Type m.TypeString `json:"type"` // Version 1.0
	URL  m.URI        `json:"url"`  // Version 1.0
}

// NewActionOpenURL constructs an Action.OpenUrl with a validated absolute http/https URL.
func NewActionOpenURL(title, u string) (ActionOpenURL, error) {
	validURL, err := m.ValidateActionURL(u, ErrInvalidActionOpenURLURL)
	if err != nil {
		return ActionOpenURL{}, err
	}

	return ActionOpenURL{
		ActionBase: ActionBase{
			Title: title,
		},
		Type: m.TypeActionOpenURL,
		URL:  m.URI(validURL),
	}, nil
}

func (a ActionOpenURL) GetType() m.TypeString { return m.TypeActionOpenURL }

func (a ActionOpenURL) WithTitle(title string) ActionOpenURL { a.Title = title; return a }

func (a ActionOpenURL) WithURL(u string) (ActionOpenURL, error) {
	validURL, err := m.ValidateActionURL(u, ErrInvalidActionOpenURLURL)
	if err != nil {
		return a, err
	}
	a.URL = m.URI(validURL)
	return a, nil
}

func (a ActionOpenURL) WithIconURL(icon string) (ActionOpenURL, error) {
	validURL, err := m.ValidateActionURL(icon, ErrInvalidActionOpenURLURL)
	if err != nil {
		return a, fmt.Errorf("Action.OpenUrl: invalid iconUrl: %w", err)
	}
	a.IconURL = m.URI(validURL)
	return a, nil
}
func (a ActionOpenURL) WithStyle(s m.ActionStyle) ActionOpenURL { a.Style = s; return a }

func (a ActionOpenURL) StyleDefault() ActionOpenURL { a.Style = ActionStyleDefault; return a }

func (a ActionOpenURL) StylePositive() ActionOpenURL { a.Style = ActionStylePositive; return a }

func (a ActionOpenURL) StyleDestructive() ActionOpenURL { a.Style = ActionStyleDestructive; return a }

func (a ActionOpenURL) WithTooltip(t string) ActionOpenURL { a.Tooltip = t; return a }
func (a ActionOpenURL) WithIsEnabled(b bool) ActionOpenURL { a.IsEnabled = &b; return a }

func (a ActionOpenURL) Enable() ActionOpenURL { b := true; a.IsEnabled = &b; return a }

func (a ActionOpenURL) Disable() ActionOpenURL                { b := false; a.IsEnabled = &b; return a }
func (a ActionOpenURL) WithMode(m m.ActionMode) ActionOpenURL { a.Mode = m; return a }

func (a ActionOpenURL) ModePrimary() ActionOpenURL { a.Mode = ActionModePrimary; return a }

func (a ActionOpenURL) ModeSecondary() ActionOpenURL { a.Mode = ActionModeSecondary; return a }
func (a ActionOpenURL) WithRequires(r map[string]string) ActionOpenURL {
	a.Requires = r
	return a
}
func (a ActionOpenURL) AddRequire(key, value string) ActionOpenURL {
	if a.Requires == nil {
		a.Requires = make(map[string]string)
	}
	a.Requires[key] = value
	return a
}

func (a ActionOpenURL) Validate() error {
	if a.URL == "" {
		return fmt.Errorf("ActionOpenURL.url is required and cannot be empty")
	}
	// Validate URL scheme/host
	if _, err := m.ValidateActionURL(string(a.URL), ErrInvalidActionOpenURLURL); err != nil {
		return fmt.Errorf("ActionOpenURL.url: %w", err)
	}
	// Validate iconUrl if provided
	if a.IconURL != "" {
		if _, err := m.ValidateActionURL(string(a.IconURL), ErrInvalidActionOpenURLURL); err != nil {
			return fmt.Errorf("ActionOpenURL.iconUrl: %w", err)
		}
	}
	return nil
}

// Custom JSON marshalling to ensure Type is always set
func (a ActionOpenURL) MarshalJSON() ([]byte, error) {
	aa := a
	if aa.Type == "" {
		aa.Type = m.TypeActionOpenURL
	}
	type alias ActionOpenURL
	return json.Marshal(alias(aa))
}

func (a *ActionOpenURL) UnmarshalJSON(b []byte) error {
	type alias ActionOpenURL
	var tmp alias
	if err := json.Unmarshal(b, &tmp); err != nil {
		return fmt.Errorf("Action.OpenUrl: decode: %w", err)
	}
	if tmp.Type != "" && tmp.Type != m.TypeActionOpenURL {
		return fmt.Errorf("Action.OpenUrl.type must be %q (got %q)", m.TypeActionOpenURL, tmp.Type)
	}
	if tmp.URL == "" {
		return fmt.Errorf("Action.OpenUrl.url is required")
	}
	tmp.Type = m.TypeActionOpenURL
	val := ActionOpenURL(tmp)
	if err := val.Validate(); err != nil {
		return err
	}
	*a = val
	return nil
}

func init() {
	RegisterAction(
		m.TypeActionOpenURL,
		func() Action { return &ActionOpenURL{Type: m.TypeActionOpenURL} },
	)
}
