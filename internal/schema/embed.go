package schema

import _ "embed"

// AdaptiveCardSchema is the embedded Adaptive Cards 1.5 JSON Schema used for validation.
//
//go:embed adaptive-card.json
var AdaptiveCardSchema []byte
