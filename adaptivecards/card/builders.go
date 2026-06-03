package card

import (
	"fmt"
	"strings"

	a "github.com/untcha/go-adaptivecards/adaptivecards/actions"
	t "github.com/untcha/go-adaptivecards/adaptivecards/containers"
	e "github.com/untcha/go-adaptivecards/adaptivecards/core/element"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
	els "github.com/untcha/go-adaptivecards/adaptivecards/elements"
)

// Chainable helpers (keep business logic outside of JSON structs)

// Title appends a heading-styled TextBlock to the card body.
func (c *Card) Title(title string) *Card {
	if c == nil || c.buildErr != nil {
		return c
	}
	// Convention: add a bolder medium TextBlock as "title"
	c.Body = append(c.Body, els.NewTextBlock(title).WithColor(m.ColorDark).WithStyle(m.TextBlockStyleHeading))
	return c
}

// AddElement appends an element to the card body.
func (c *Card) AddElement(el e.Element) *Card {
	if c == nil || c.buildErr != nil {
		return c
	}
	c.Body = append(c.Body, el)
	return c
}

// AddImage creates and adds an Image element.
func (c *Card) AddImage(url string) *Card {
	if c == nil || c.buildErr != nil {
		return c
	}
	c.Body = append(c.Body, els.NewImage(url))
	return c
}

// AddContainer creates and adds a Container element.
func (c *Card) AddContainer(items ...e.Element) *Card {
	if c == nil || c.buildErr != nil {
		return c
	}
	c.Body = append(c.Body, t.NewContainer(items...))
	return c
}

// AddTextBlock adds a TextBlock to the Card's Body.
// If tb.Type is empty, it sets it to "TextBlock".
func (c *Card) AddTextBlock(tb els.TextBlock) *Card {
	if c == nil || c.buildErr != nil {
		return c
	}
	if tb.Type == "" {
		tb.Type = m.TypeTextBlock
	}
	c.Body = append(c.Body, tb)
	return c
}

// AddAction appends an action to the card.
func (c *Card) AddAction(act a.Action) *Card {
	if c == nil || c.buildErr != nil {
		return c
	}
	c.Actions = append(c.Actions, act)
	return c
}

// AddActionOpenURL creates and adds an Action.OpenUrl.
// If URL creation fails, buildErr is set and the chain continues.
func (c *Card) AddActionOpenURL(title, url string) *Card {
	if c == nil || c.buildErr != nil {
		return c
	}
	act, err := a.NewActionOpenURL(title, url)
	if err != nil {
		c.buildErr = err
		return c
	}
	c.Actions = append(c.Actions, act)
	return c
}

// AddActionSubmit creates and adds an Action.Submit.
func (c *Card) AddActionSubmit(title string, data any) *Card {
	if c == nil || c.buildErr != nil {
		return c
	}
	c.Actions = append(c.Actions, a.NewActionSubmit(title, data))
	return c
}

// AddActionToggleVisibility creates and adds an Action.ToggleVisibility.
func (c *Card) AddActionToggleVisibility(title string, targets ...a.TargetElement) *Card {
	if c == nil || c.buildErr != nil {
		return c
	}
	c.Actions = append(c.Actions, a.NewActionToggleVisibility(title, targets...))
	return c
}

// AddTable adds a Table to the Card's Body.
// If table.Type is empty, it sets it to "Table".
func (c *Card) AddTable(table t.Table) *Card {
	if c == nil || c.buildErr != nil {
		return c
	}
	if table.Type == "" {
		table.Type = m.TypeTable
	}
	c.Body = append(c.Body, table)
	return c
}

// SetSelectAction sets the card-level selectAction; Action.ShowCard is rejected.
func (c *Card) SetSelectAction(act a.Action) *Card {
	if c == nil {
		return c
	}
	if c.buildErr != nil {
		return c
	}
	if act == nil {
		c.buildErr = fmt.Errorf("selectAction cannot be nil")
		return c
	}
	// Validate that Action.ShowCard is not used (as per schema)
	if act.GetType() == m.TypeActionShowCard {
		c.buildErr = fmt.Errorf("Action.ShowCard is not supported in selectAction")
		return c
	}
	c.SelectAction = act
	return c
}

// ---- Card - BackgroundImage builder methods ----

// SetBackgroundImageURL sets the BackgroundImage from a URL string.
// If the URL is invalid, it sets c.buildErr and returns the Card for chaining.
func (c *Card) SetBackgroundImageURL(u string) *Card {
	if c == nil {
		return c
	}
	if c.buildErr != nil {
		return c
	}
	v, err := m.BackgroundImageURL(u)
	if err != nil {
		c.buildErr = err
		return c
	}
	c.BackgroundImage = v
	return c
}

// SetBackgroundImage sets the BackgroundImage from a BackgroundImage object.
// If the URL is invalid, it sets c.buildErr and returns the Card for chaining.
func (c *Card) SetBackgroundImage(bg m.BackgroundImage) *Card {
	if c == nil {
		return c
	}
	if c.buildErr != nil {
		return c
	}
	v, err := m.BackgroundImageObject(bg)
	if err != nil {
		c.buildErr = err
		return c
	}
	c.BackgroundImage = v
	return c
}

// SetBackgroundImageWithFillMode sets the BackgroundImage with specified URL and fill mode.
func (c *Card) SetBackgroundImageWithFillMode(url string, fillMode m.ImageFillMode) *Card {
	if c == nil {
		return c
	}
	if c.buildErr != nil {
		return c
	}
	bg, err := m.NewBackgroundImage(url, fillMode, "", "")
	if err != nil {
		c.buildErr = err
		return c
	}
	return c.SetBackgroundImage(bg)
}

// SetBackgroundImageWithAlignment sets the BackgroundImage with specified URL and alignments.
func (c *Card) SetBackgroundImageWithAlignment(url string, horizontalAlign m.HorizontalAlignment, verticalAlign m.VerticalAlignment) *Card {
	if c == nil {
		return c
	}
	if c.buildErr != nil {
		return c
	}
	bg, err := m.NewBackgroundImage(url, "", horizontalAlign, verticalAlign)
	if err != nil {
		c.buildErr = err
		return c
	}
	return c.SetBackgroundImage(bg)
}

// SetBackgroundImageComplete sets the BackgroundImage with all parameters.
func (c *Card) SetBackgroundImageComplete(url string, fillMode m.ImageFillMode, horizontalAlign m.HorizontalAlignment, verticalAlign m.VerticalAlignment) *Card {
	if c == nil {
		return c
	}
	if c.buildErr != nil {
		return c
	}
	bg, err := m.NewBackgroundImage(url, fillMode, horizontalAlign, verticalAlign)
	if err != nil {
		c.buildErr = err
		return c
	}
	return c.SetBackgroundImage(bg)
}

// Convenience methods for common background image configurations

// SetBackgroundImageCover sets a background image with cover fill mode.
func (c *Card) SetBackgroundImageCover(url string) *Card {
	return c.SetBackgroundImageWithFillMode(url, m.ImageFillCover)
}

// SetBackgroundImageRepeat sets a background image with repeat fill mode.
func (c *Card) SetBackgroundImageRepeat(url string) *Card {
	return c.SetBackgroundImageWithFillMode(url, m.ImageFillRepeat)
}

// SetBackgroundImageCenterCover sets a background image with cover fill mode and center alignment.
func (c *Card) SetBackgroundImageCenterCover(url string) *Card {
	return c.SetBackgroundImageComplete(url, m.ImageFillCover, m.HAlignCenter, m.VAlignCenter)
}

// SetLang sets the card language (BCP-47, e.g. "en", "de-DE").
func (c *Card) SetLang(lang string) *Card {
	if c == nil {
		return c
	}
	if c.buildErr != nil {
		return c
	}
	c.Lang = lang
	return c
}

// SetSpeak sets the card speak value for speech synthesis hosts.
func (c *Card) SetSpeak(speak string) *Card {
	if c == nil {
		return c
	}
	if c.buildErr != nil {
		return c
	}
	c.Speak = speak
	return c
}

// SetFallbackText sets plain-text fallback for unsupported hosts.
func (c *Card) SetFallbackText(text string) *Card {
	if c == nil {
		return c
	}
	if c.buildErr != nil {
		return c
	}
	c.FallbackText = text
	return c
}

// ---- Card - MinHeight builder method ----

// SetCardMinHeight sets the Card's minHeight property (e.g. "300px").
func (c *Card) SetCardMinHeight(h string) *Card {
	if c == nil {
		return c
	}
	if c.buildErr != nil {
		return c
	}
	s := strings.TrimSpace(h)
	if s == "" {
		c.buildErr = m.ErrInvalidMinHeight
		return c
	}
	// Basic validation: must end with "px"
	if !strings.HasSuffix(s, "px") {
		c.buildErr = m.ErrInvalidMinHeight
		return c
	}
	c.MinHeight = s
	return c
}

// ---- Card - Rtl builder method ----

// SetCardRtl sets the Card's rtl property (right-to-left).
// Default is false (left-to-right).
func (c *Card) SetCardRtl(rtl bool) *Card {
	if c == nil {
		return c
	}
	if c.buildErr != nil {
		return c
	}
	c.Rtl = &rtl
	return c
}

// ---- Card - MSTeams host extension builder methods ----

// SetFullWidth makes the card span the full message column in Teams clients
// by setting msteams.width = "Full". Passing false clears only the width
// sub-field, leaving any other MSTeams extensions intact.
//
// msteams is a Teams host extension, not part of the Adaptive Cards schema;
// other renderers ignore it.
func (c *Card) SetFullWidth(full bool) *Card {
	if c == nil || c.buildErr != nil {
		return c
	}
	if !full {
		if c.MSTeams != nil {
			c.MSTeams.Width = ""
		}
		return c
	}
	if c.MSTeams == nil {
		c.MSTeams = &MSTeams{}
	}
	c.MSTeams.Width = MSTeamsWidthFull
	return c
}

// SetMSTeams sets the full Teams host-extension object on the card.
// It validates the extension immediately and records buildErr on failure.
func (c *Card) SetMSTeams(t MSTeams) *Card {
	if c == nil || c.buildErr != nil {
		return c
	}
	if err := t.Validate(); err != nil {
		c.buildErr = err
		return c
	}
	c.MSTeams = &t
	return c
}

// ---- Card - VerticalContentAlignment builder method ----

// SetVerticalContentAlignment sets the Card's verticalContentAlignment property.
// Valid values are "top", "center", "bottom".
func (c *Card) SetVerticalContentAlignment(v m.VerticalContentAlignment) *Card {
	if c == nil {
		return c
	}
	if c.buildErr != nil {
		return c
	}
	if v != "" && !v.IsValid() {
		c.buildErr = m.NewEnumError("Card.verticalContentAlignment", string(v), m.AllowedVerticalContentAlignmentStrings())
		return c
	}
	c.VerticalContentAlignment = v
	return c
}
