package model

import (
	"fmt"
	"net/url"
	"slices"
	"strings"
)

// URLValidationOptions provides configuration for URL validation.
// This allows for future extensibility without breaking existing code.
type URLValidationOptions struct {
	// AllowDataURIs enables support for data: scheme URIs (e.g., for inline images)
	AllowDataURIs bool
	// AllowedSchemes specifies custom allowed schemes. If empty, defaults to http/https
	AllowedSchemes []string
	// RequireHost specifies whether a host is required (default: true)
	RequireHost bool
}

// DefaultURLValidationOptions returns the standard validation options used
// throughout the library: http/https only, host required, no data URIs.
func DefaultURLValidationOptions() URLValidationOptions {
	return URLValidationOptions{
		AllowDataURIs:  false,
		AllowedSchemes: []string{"http", "https"},
		RequireHost:    true,
	}
}

// validateURL validates and normalizes a given URL string.
// It wraps validation errors with the provided baseErr for context.
// Returns the normalized URL string or an error if invalid.
// The URL must be an absolute http/https URL with a valid host.
func validateURL(u string, baseErr error) (string, error) {
	return validateURLWithOptions(u, baseErr, DefaultURLValidationOptions())
}

// validateURLWithOptions validates and normalizes a URL string with custom options.
// This provides flexibility for future requirements while maintaining backward compatibility.
func validateURLWithOptions(u string, baseErr error, opts URLValidationOptions) (string, error) {
	s := strings.TrimSpace(u)
	if s == "" {
		return "", fmt.Errorf("%w: empty URL", baseErr)
	}

	parsed, err := url.Parse(s)
	if err != nil {
		return "", fmt.Errorf("%w: parse URL %v", baseErr, err)
	}

	// Check if URL is absolute
	if !parsed.IsAbs() {
		return "", fmt.Errorf("%w: must be absolute URL", baseErr)
	}

	// Determine allowed schemes
	allowedSchemes := opts.AllowedSchemes
	if len(allowedSchemes) == 0 {
		allowedSchemes = []string{"http", "https"}
	}

	// Add data scheme if allowed
	if opts.AllowDataURIs {
		allowedSchemes = append(allowedSchemes, "data")
	}

	// Validate scheme
	if !slices.Contains(allowedSchemes, parsed.Scheme) {
		if len(allowedSchemes) == 1 {
			return "", fmt.Errorf("%w: unsupported scheme %q (want %s)", baseErr, parsed.Scheme, allowedSchemes[0])
		}
		return "", fmt.Errorf("%w: unsupported scheme %q (want %s)", baseErr, parsed.Scheme, strings.Join(allowedSchemes, "/"))
	}

	// Validate host (if required and not a data URI)
	if opts.RequireHost && parsed.Scheme != "data" && parsed.Host == "" {
		return "", fmt.Errorf("%w: missing host", baseErr)
	}

	return parsed.String(), nil
}

// Helper functions for common validation scenarios

// validateImageURL validates URLs specifically for images, which might support data URIs in the future
func validateImageURL(u string, baseErr error) (string, error) {
	// Keep strict behavior for existing call sites (e.g. background images).
	return validateURL(u, baseErr)
}

// validateImageURIReference validates image URLs as uri-reference (schema-compatible):
// absolute URLs, relative references, and data URIs are all allowed.
func validateImageURIReference(u string, baseErr error) (string, error) {
	s := strings.TrimSpace(u)
	if s == "" {
		return "", fmt.Errorf("%w: empty URL", baseErr)
	}
	parsed, err := url.Parse(s)
	if err != nil {
		return "", fmt.Errorf("%w: parse URL %v", baseErr, err)
	}
	return parsed.String(), nil
}

// validateActionURL validates URLs for actions, which should always be external links
func validateActionURL(u string, baseErr error) (string, error) {
	// Actions should always be external http/https URLs
	return validateURL(u, baseErr)
}

// ValidateImageURL validates image URLs with strict rules used by legacy call sites
// (absolute http/https URLs with host).
func ValidateImageURL(u string, baseErr error) (string, error) {
	return validateImageURL(u, baseErr)
}

// ValidateImageURIReference validates image URLs as schema-compatible uri-reference.
// Relative references and data URIs are accepted.
func ValidateImageURIReference(u string, baseErr error) (string, error) {
	return validateImageURIReference(u, baseErr)
}

// ValidateActionURL validates action URLs as absolute external links.
func ValidateActionURL(u string, baseErr error) (string, error) {
	return validateActionURL(u, baseErr)
}
