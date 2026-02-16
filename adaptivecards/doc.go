// Package adaptivecards is the repository root namespace.
//
// The functional API is organized in subpackages:
//   - card: AdaptiveCard root model and card-level builders/validation
//   - elements: leaf card elements (for example TextBlock, Image, RichTextBlock)
//   - containers: container-style elements (for example Container, ColumnSet, Table)
//   - actions: card actions (for example Action.OpenUrl, Action.Submit)
//   - inputs: input elements (for example Input.Text, Input.ChoiceSet)
//   - webhook: sending cards to workflow/webhook endpoints
//   - core/model: shared enums, value objects, and low-level validation helpers
//   - core/element: element interfaces, fallback model, and element factory registry
//   - schema: JSON schema validation integration
//
// Import subpackages directly in application code.
package adaptivecards
