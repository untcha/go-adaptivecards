![go-adaptivecards README header](assets/readme-header.png)

# go-adaptivecards

Experimental Go library for building and validating Microsoft [Adaptive Cards](https://adaptivecards.io/).

## Important Status

**This project is experimental and not production-ready.**

- API and behavior can change without notice.
- Breaking changes are expected.
- The implementation is **not** at 100% Microsoft Adaptive Cards feature parity.
- Use with caution, especially in production systems.

## Scope

This repository provides:

- Strongly typed Go models for selected Adaptive Card features
- Builder-style APIs for composing cards
- Logical validation and JSON Schema validation support against [schema versions 1.5.0](https://adaptivecards.io/schemas/1.5.0/adaptive-card.json)
- JSON factory-based decoding for interface fields (`Element`, `Action`)
- Optional webhook posting helper for Teams/workflow endpoints

## Install

```bash
go get github.com/untcha/go-adaptivecards
```

## Quick Example

Quick "Hello, World!" example (from `examples/simple/main.go`):

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/untcha/go-adaptivecards/adaptivecards/card"
	"github.com/untcha/go-adaptivecards/adaptivecards/elements"
	"github.com/untcha/go-adaptivecards/adaptivecards/webhook"
)

func main() {
	// Webhook URL
	url := os.Getenv("TEAMS_TEST_WORKFLOW_URL")

	textblock := elements.NewTextBlock("Hello, World!")

	card, err := card.NewCard().
		AddElement(textblock).
		Build()

	if err != nil {
		panic(err)
	}

	json, _ := card.MarshalJSON()
	fmt.Println(string(json))

	// Post to Teams webhook if URL is provided
	if url != "" {
		if err := webhook.PostToWorkflowRaw(context.Background(), url, card); err != nil {
			panic(err)
		}
	}
}
```

## Feature Matrix (Current State)

Legend:

- `[x]` Type/feature exists in this repo
- `[ ]` Not implemented yet
- Checked items can still be partial vs full official schema behavior

### AdaptiveCard Root

- [x] `AdaptiveCard`
- [x] `body`
- [x] `actions`
- [x] `selectAction`
- [x] `backgroundImage`
- [x] `minHeight`
- [x] `rtl`
- [x] `lang`
- [x] `speak`
- [x] `fallbackText`
- [x] `verticalContentAlignment`
- [ ] `refresh` (placeholder only)
- [ ] `authentication` (placeholder only)
- [ ] `metadata` (placeholder only)

### Actions

- [x] `Action.OpenUrl`
- [x] `Action.Submit`
- [x] `Action.ToggleVisibility`
- [ ] `Action.ShowCard`
- [ ] `Action.Execute`

### Elements

- [x] `TextBlock`
- [x] `RichTextBlock`
- [x] `TextRun`
- [x] `Image`
- [ ] `Media`

### Containers

- [x] `Container`
- [x] `Column`
- [x] `ColumnSet`
- [x] `ActionSet`
- [x] `FactSet`
- [x] `ImageSet`
- [x] `Table` (including `TableRow`, `TableCell`, `TableColumnDefinition`)

### Inputs

- [x] `Input.Text`
- [x] `Input.Number`
- [x] `Input.Date`
- [x] `Input.Time`
- [x] `Input.ChoiceSet`
- [x] `Input.Toggle`

### Tooling / Validation

- [x] Logical validation methods on card/types
- [x] JSON factory decode for action/element interfaces
- [x] Embedded Adaptive Card schema validation support
- [x] Webhook helper (`webhook.PostToWorkflowRaw`)

## Package Layout

- `adaptivecards/card`: AdaptiveCard root model and card-level builders/validation
- `adaptivecards/actions`: action types and action factory
- `adaptivecards/elements`: leaf elements and element validation
- `adaptivecards/containers`: container-style elements, including Table
- `adaptivecards/inputs`: input elements and validation
- `adaptivecards/core/model`: shared enums, value objects, low-level validation
- `adaptivecards/core/element`: element interfaces, fallback model, element factory
- `adaptivecards/schema`: schema validation integration
- `adaptivecards/webhook`: helpers to send card JSON to webhook endpoints

## Development

Common commands:

```bash
task test
task lint
task check
task project:update:schema
```

`Taskfile.yml` holds the generic Go-library tasks (shared across repos); repo-specific
tasks live in `Taskfile.project.yml` and are included under the `project:` namespace.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
