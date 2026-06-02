package schema

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/santhosh-tekuri/jsonschema/v6"
	"github.com/untcha/go-adaptivecards/internal/schema"
)

var (
	compileOnce sync.Once
	compiled    *jsonschema.Schema
	compileErr  error
)

const schemaURL = "mem://adaptive-cards/v1.5/schema.json"

// Lazily compile the schema once.
func ensureCompiled() error {
	compileOnce.Do(func() {
		c := jsonschema.NewCompiler()
		// IMPORTANT: let the file's $schema (draft-06) control the draft.
		// Do NOT set c.DefaultDraft(...).

		// Unmarshal to a generic value and add as a resource.
		var schemaAny any
		if err := json.Unmarshal(schema.AdaptiveCardSchema, &schemaAny); err != nil {
			compileErr = fmt.Errorf("decode embedded schema: %w", err)
			return
		}
		// Some upstream schema versions encode optional fields as keys like "rtl?".
		// Normalize these keys so runtime JSON ("rtl") validates as intended.
		schemaAny = normalizeOptionalKeySuffix(schemaAny)

		// Older jsonschema versions expect AddResource(url string, v any).
		// Passing the decoded JSON value avoids the "*bytes.Reader" issue.
		if err := c.AddResource(schemaURL, schemaAny); err != nil {
			compileErr = fmt.Errorf("add schema resource: %w", err)
			return
		}

		compiled, compileErr = c.Compile(schemaURL)
	})
	return compileErr
}

func Validate(v any) error {
	if v == nil {
		return errors.New("value is nil")
	}
	if err := ensureCompiled(); err != nil {
		return fmt.Errorf("compile schema: %w", err)
	}

	// Validate a normal Go JSON value (map/slice/primitive), not raw bytes/readers.
	var normalized any
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("marshal value: %w", err)
	}
	if err := json.Unmarshal(b, &normalized); err != nil {
		return fmt.Errorf("prepare for validation: %w", err)
	}
	// Remove documented host features / extensions that are not part of the
	// published JSON schema (which uses additionalProperties:false). They are
	// validated logically by the model, not by the schema.
	normalized = stripNonSchemaKeys(normalized)
	if err := compiled.Validate(normalized); err != nil {
		return fmt.Errorf("adaptive card schema: %w", err)
	}
	return nil
}

// nonSchemaKeys lists JSON property names that are valid Adaptive Cards host
// features but are absent from the published JSON schema. They are stripped
// (at any depth) before strict schema validation.
//
// Note: the card-root "msteams" extension is stripped by Card.Validate before
// it ever reaches here; this set covers fields that live on nested elements.
var nonSchemaKeys = map[string]struct{}{
	"targetWidth": {}, // documented responsive-visibility feature; not in any schema version
}

// stripNonSchemaKeys recursively removes nonSchemaKeys entries from a decoded
// JSON value (map/slice tree).
func stripNonSchemaKeys(v any) any {
	switch x := v.(type) {
	case map[string]any:
		out := make(map[string]any, len(x))
		for k, vv := range x {
			if _, drop := nonSchemaKeys[k]; drop {
				continue
			}
			out[k] = stripNonSchemaKeys(vv)
		}
		return out
	case []any:
		for i := range x {
			x[i] = stripNonSchemaKeys(x[i])
		}
		return x
	default:
		return v
	}
}

func normalizeOptionalKeySuffix(v any) any {
	switch x := v.(type) {
	case map[string]any:
		out := make(map[string]any, len(x))
		for k, vv := range x {
			nk := strings.TrimSuffix(k, "?")
			// Keep an existing normalized key if present, so we don't override explicit keys.
			if _, exists := out[nk]; !exists {
				out[nk] = normalizeOptionalKeySuffix(vv)
			}
		}
		return out
	case []any:
		out := make([]any, len(x))
		for i := range x {
			out[i] = normalizeOptionalKeySuffix(x[i])
		}
		return out
	default:
		return v
	}
}
