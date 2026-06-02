# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

While the project is `0.x`, the API is experimental and may change between minor
versions (see the status note in the README).

## [Unreleased]

_Targeted for `v0.3.0`._

### Added

- **`targetWidth` responsive visibility** on all elements (`ElementBase`). Lets an
  element declare at which card widths it renders — bare buckets (`veryNarrow`,
  `narrow`, `standard`, `wide`) or range forms (`atLeast:<bucket>`,
  `atMost:<bucket>`). Like `msteams`, this is a documented Adaptive Cards host
  feature that is **not** part of the published JSON schema; it is validated
  logically and stripped before strict schema validation.
  - `m.TargetWidth` enum with `TargetWidthAtLeast`/`TargetWidthAtMost` builders,
    `IsValid()`, `ParseTargetWidth`, and `AllowedTargetWidthStrings()`.
  - `WithTargetWidth(...)` fluent setters on `TextBlock`, `Image`, and `Table`.
  - `examples/responsive/main.go`.

### Changed

- Centralized the per-element `validateElementBase()` logic into a single shared
  `ElementBase.Validate(typeName)` (previously duplicated across 10 element/
  container files). Base-field validation error messages now consistently use the
  capitalized element type name (e.g. `Image.id`, previously `image.id`).
- `schema.Validate` recursively strips non-schema host keys (currently
  `targetWidth`) at any depth before validating, so nested responsive elements
  pass strict schema validation.

## [0.2.0] - 2026-06-02

### Added

- **Microsoft Teams `msteams` host extension** on the card root. This is a
  Teams-specific extension, not part of the Adaptive Cards schema; it is
  ignored by other renderers.
  - `MSTeams` type with a typed `MSTeamsWidth` enum and the `MSTeamsWidthFull`
    constant (`adaptivecards/card/msteams.go`).
  - `Card.MSTeams *MSTeams` field — optional and pointer-typed, so cards that
    don't use it serialize exactly as before.
  - Builders `Card.SetFullWidth(bool)` (emits `"msteams":{"width":"Full"}`) and
    `Card.SetMSTeams(MSTeams)`.
  - Full-width cards round-trip through `MarshalJSON`/`UnmarshalJSON` and pass
    the validated `webhook.PostToWorkflowRaw` path (no post-marshal JSON
    injection required).
  - `examples/full_width/main.go` demonstrating the feature.

### Changed

- `Card.Validate()` now validates the `MSTeams` extension logically and strips
  it from a copy before JSON-schema validation, since the embedded schema sets
  `additionalProperties:false` on `AdaptiveCard`.

## [0.1.0] - 2026-03-16

### Added

- Initial experimental release: strongly typed models and builder APIs for
  selected Adaptive Cards features, logical + embedded JSON-schema validation
  (schema 1.5.0), factory-based decoding for `Element`/`Action` interfaces, and
  an optional Teams/workflow webhook helper. See the README feature matrix for
  the implemented surface.

[Unreleased]: https://github.com/untcha/go-adaptivecards/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/untcha/go-adaptivecards/compare/v0.1.0...v0.2.0
