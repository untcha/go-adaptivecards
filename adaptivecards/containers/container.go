package containers

import (
	"encoding/json"
	"fmt"
	"strings"

	a "github.com/untcha/go-adaptivecards/adaptivecards/actions"
	e "github.com/untcha/go-adaptivecards/adaptivecards/core/element"
	m "github.com/untcha/go-adaptivecards/adaptivecards/core/model"
)

// Container groups items together.
// See: https://adaptivecards.io/explorer/Container.html
type Container struct {
	e.ElementBase                                       // Embedding e.ElementBase to include common element fields
	Type                     m.TypeString               `json:"type"`                               // Version 1.0
	Items                    []e.Element                `json:"items,omitempty"`                    // Version 1.0
	SelectAction             a.Action                   `json:"selectAction,omitempty"`             // Version 1.1
	Style                    m.ContainerStyle           `json:"style,omitempty"`                    // Version 1.0
	VerticalContentAlignment m.VerticalContentAlignment `json:"verticalContentAlignment,omitempty"` // Version 1.1
	Bleed                    bool                       `json:"bleed,omitempty"`                    // Version 1.2
	BackgroundImage          *m.BackgroundImageValue    `json:"backgroundImage,omitempty"`          // Version 1.2
	MinHeight                string                     `json:"minHeight,omitempty"`                // Version 1.2
	Rtl                      *bool                      `json:"rtl,omitempty"`                      // Version 1.5
}

// NewContainer returns a Container holding the given items.
func NewContainer(items ...e.Element) Container {
	return Container{
		Type:  m.TypeContainer,
		Items: items,
	}
}

// GetType returns the Adaptive Card type discriminator for Container.
func (c Container) GetType() m.TypeString { return m.TypeContainer }

// Validate reports whether the Container and its children are well-formed.
func (c Container) Validate() error {
	if err := c.validateElementBase(); err != nil {
		return err
	}
	if err := validateElements(c.Items); err != nil {
		return fmt.Errorf("container.items: %w", err)
	}
	if c.SelectAction != nil {
		if err := validateSelectAction(c.SelectAction); err != nil {
			return fmt.Errorf("container.selectAction: %w", err)
		}
	}
	if c.Style != "" && !c.Style.IsValid() {
		return m.NewEnumError("Container.style", string(c.Style), m.AllowedContainerStyleStrings())
	}
	if c.VerticalContentAlignment != "" && !c.VerticalContentAlignment.IsValid() {
		return m.NewEnumError(
			"Container.verticalContentAlignment",
			string(c.VerticalContentAlignment),
			m.AllowedVerticalContentAlignmentStrings(),
		)
	}
	if c.BackgroundImage != nil {
		if c.BackgroundImage.Object() != nil {
			if err := c.BackgroundImage.Object().Validate(); err != nil {
				return fmt.Errorf("container.backgroundImage: %w", err)
			}
		} else if c.BackgroundImage.URL() != "" {
			if _, err := m.BackgroundImageURL(c.BackgroundImage.URL()); err != nil {
				return fmt.Errorf("container.backgroundImage: %w", err)
			}
		}
	}
	if c.MinHeight != "" && !isValidPixelWidth(c.MinHeight) {
		return fmt.Errorf(
			"container.minHeight must be in format \"<number>px\" (got %q)",
			c.MinHeight,
		)
	}
	return nil
}

// MarshalJSON implements json.Marshaler, ensuring the type field is set.
func (c Container) MarshalJSON() ([]byte, error) {
	cc := c
	if cc.Type == "" {
		cc.Type = m.TypeContainer
	}
	type alias Container
	return json.Marshal(alias(cc))
}

// UnmarshalJSON implements json.Unmarshaler, decoding interface-typed fields via the element/action factories.
func (c *Container) UnmarshalJSON(b []byte) error {
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(b, &obj); err != nil {
		return fmt.Errorf("container: decode: %w", err)
	}
	if v, ok := obj["type"]; ok {
		var t string
		if err := json.Unmarshal(v, &t); err != nil {
			return fmt.Errorf("container.type: %w", err)
		}
		if t != "" && t != string(m.TypeContainer) {
			return fmt.Errorf("container.type must be %q (got %q)", m.TypeContainer, t)
		}
	}
	itemsRaw := obj["items"]
	selectActionRaw := obj["selectAction"]
	delete(obj, "items")
	delete(obj, "selectAction")

	type alias Container
	var base alias
	baseBytes, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("container: re-encode base: %w", err)
	}
	if err := json.Unmarshal(baseBytes, &base); err != nil {
		return fmt.Errorf("container: decode base: %w", err)
	}
	if base.Type == "" {
		base.Type = m.TypeContainer
	}

	if len(itemsRaw) != 0 {
		var rawArr []json.RawMessage
		if err := json.Unmarshal(itemsRaw, &rawArr); err != nil {
			return fmt.Errorf("container.items: %w", err)
		}
		items, err := e.UnmarshalElementsSlice(rawArr)
		if err != nil {
			return fmt.Errorf("container.items: %w", err)
		}
		base.Items = items
	}

	if len(selectActionRaw) != 0 {
		action, err := a.UnmarshalAction(selectActionRaw)
		if err != nil {
			return fmt.Errorf("container.selectAction: %w", err)
		}
		base.SelectAction = action
	}

	*c = Container(base)
	return nil
}

func (c Container) validateElementBase() error {
	if c.Height != "" && !c.Height.IsValid() {
		return m.NewEnumError(
			"Container.height",
			string(c.Height),
			m.AllowedBlockElementHeightStrings(),
		)
	}
	if c.Spacing != "" && !c.Spacing.IsValid() {
		return m.NewEnumError("Container.spacing", string(c.Spacing), m.AllowedSpacingStrings())
	}
	if c.ID != "" {
		id := strings.TrimSpace(c.ID)
		if id == "" {
			return fmt.Errorf("container.id cannot be empty or whitespace-only")
		}
		if strings.ContainsAny(id, "\n\r\t") {
			return fmt.Errorf("container.id cannot contain newlines or tabs")
		}
	}
	return nil
}

func init() {
	e.RegisterElement(
		m.TypeContainer,
		func() e.Element { return &Container{Type: m.TypeContainer} },
	)
}
