# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

Working style, Go standards, and commit rules live in AGENTS.md (commit messages:
propose only, per `docs/COMMIT_GUIDE.md`; never stage or commit):

@AGENTS.md

## Project

Experimental Go library (Go 1.25+) for building and validating Microsoft Adaptive
Cards against schema **1.5.0**. Not at full feature parity; breaking changes are
acceptable. Keep the feature matrix in `README.md` and `CHANGELOG.md` (Keep a
Changelog + semver) in sync with user-visible changes.

## Commands

Automation is Taskfile-based. `Taskfile.yml` flattens shared tasks from
`taskfiles/common.yml` (run unprefixed) and namespaces repo tasks from
`taskfiles/Taskfile.project.yml` under `project:`. Both `Taskfile.yml` and
`common.yml` come from the shared meta repo and should stay identical across repos.

```bash
task check                 # fmt, lint, vet, build (compile-only), test:existing
task test                  # go test ./...
task lint                  # golangci-lint run (config: .golangci.yml)
task test:coverage         # writes coverage/coverage.html
task project:update:schema # re-download schema into internal/schema/adaptive-card.json

go test ./adaptivecards/card -run TestName -v   # single test
go run ./examples/simple                         # set TEAMS_TEST_WORKFLOW_URL to also post
```

## Architecture

Public packages live under `adaptivecards/` (no `pkg/`); the embedded schema lives
in `internal/schema`.

- **Polymorphism via type registries.** `Element` (`core/element`) and `Action`
  (`actions`) are interfaces with only `GetType()`. Each concrete type registers a
  factory in its package's `init()` (`element.RegisterElement` /
  `actions.RegisterAction`, which panic on duplicates). JSON decoding probes the
  `"type"` field and dispatches through the registry
  (`UnmarshalElement[sSlice]`, `UnmarshalAction[sSlice]`). Any type holding
  `[]Element`/`Action` fields (Card, Container, Column, Table cells, ...) needs a
  custom `UnmarshalJSON` that pulls those fields out as `json.RawMessage` and
  decodes them via the factories. `MarshalJSON` implementations use the
  `type alias T` trick to avoid recursion and default the `type` field.
- **Registration depends on imports.** `card` imports `containers` and `elements`
  but not `inputs`. Decoding JSON with `Input.*` elements fails with
  `ErrUnknownElementType` unless the `inputs` package is imported somewhere.
- **Adding a new element/action:** embed `ElementBase`/`ActionBase`, add the
  `TypeString` constant in `core/model`, implement `GetType`, `Validate`, and JSON
  methods if needed, and register the factory in `init()`. `card.validateElement` /
  `validateAction` fall back to any `Validate() error` method, so explicit switch
  cases are only needed for extra checks (e.g. nested `selectAction`).
- **Fallback** (`ElementFallback`, `ActionFallback`) is either the string `"drop"`
  or a nested element/action; its JSON methods handle both shapes.
- **Builders** on `*Card` (`card/builders.go`) are chainable. They record the first
  failure in `buildErr` and turn later calls into no-ops. `Build()` returns that error.
- **`Card.Validate()` has three stages:** builder error, then logical per-element
  validation, then JSON Schema validation (`adaptivecards/schema`,
  santhosh-tekuri/jsonschema v6, compiled lazily once from the embedded draft-06
  schema; `?`-suffixed optional keys in the upstream schema are normalized). The
  `msteams` host extension is stripped from a shallow copy before schema validation
  because the spec sets `additionalProperties: false`. It is validated logically
  instead. Any new non-spec extension needs the same treatment.
- **`core/model`** holds shared enums (with `Allowed*` helpers and JSON parsing),
  value objects (`URI`, `BackgroundImage`), URL validation, and errors.
- **`webhook`** posts the validated card as the raw request body. A strict
  `URLPolicy` (HTTPS only, private/loopback/link-local blocked, optional host
  allowlist) guards against SSRF. Don't loosen the defaults; callers relax the
  policy through `PostToWorkflowRawWithClientAndPolicy`.

## Tests

Tests are table-driven `*_test.go` files next to the code. `task test:integration`
expects `//go:build integration` files whose test names contain `Integration`.
`task test:unit` filters by `-run '^Test[^I]*$'`, so avoid unit test names starting
with `TestI`.
