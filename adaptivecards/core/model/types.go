package model

// TypeString is the JSON "type" discriminator string used across Adaptive Card objects.
type TypeString string

// URI is a string holding a URI/URL value as serialized in card JSON.
type URI string

// Type marker constants for the supported Adaptive Card object types.
const (
	// Cards
	TypeAdaptiveCard TypeString = "AdaptiveCard"

	// Card Elements
	TypeTextBlock      TypeString = "TextBlock"
	TypeImage          TypeString = "Image"
	TypeContainer      TypeString = "Container"
	TypeColumnSet      TypeString = "ColumnSet"
	TypeColumn         TypeString = "Column"
	TypeActionSet      TypeString = "ActionSet"
	TypeFactSet        TypeString = "FactSet"
	TypeImageSet       TypeString = "ImageSet"
	TypeRichTextBlock  TypeString = "RichTextBlock"
	TypeTextRun        TypeString = "TextRun"
	TypeInputText      TypeString = "Input.Text"
	TypeInputChoice    TypeString = "Input.Choice"
	TypeInputChoiceSet TypeString = "Input.ChoiceSet"
	TypeInputToggle    TypeString = "Input.Toggle"
	TypeInputDate      TypeString = "Input.Date"
	TypeInputTime      TypeString = "Input.Time"
	TypeInputNumber    TypeString = "Input.Number"

	// Containers
	TypeTable                 TypeString = "Table"
	TypeTableRow              TypeString = "TableRow"
	TypeTableCell             TypeString = "TableCell"
	TypeTableColumnDefinition TypeString = "TableColumnDefinition"

	// Actions
	TypeActionOpenURL          TypeString = "Action.OpenUrl"
	TypeActionSubmit           TypeString = "Action.Submit"
	TypeActionToggleVisibility TypeString = "Action.ToggleVisibility"
	TypeActionShowCard         TypeString = "Action.ShowCard"
)
