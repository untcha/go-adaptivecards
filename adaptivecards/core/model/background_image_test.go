package model

import (
	"encoding/json"
	"strings"
	"testing"
)

// TestNewBackgroundImage and related tests provide comprehensive coverage for background_image.go.
//
// This test suite covers:
//
// 1. Constructors and factory functions:
//    - NewBackgroundImage() - validates URL immediately, allows invalid enums for flexible building
//    - MustNewBackgroundImage() - panic version for tests/demos
//    - BackgroundImageURL() - creates URL-only BackgroundImageValue
//    - MustBackgroundImageURL() - panic version
//    - BackgroundImageObject() - creates object BackgroundImageValue with validation
//    - MustBackgroundImageObject() - panic version
//
// 2. Validation:
//    - Validate() method testing all field combinations
//    - URL validation (required field, format checking)
//    - Enum validation for all optional fields (FillMode, HorizontalAlignment, VerticalAlignment)
//    - Empty enum handling (valid when empty)
//    - Defense-in-depth URL validation
//
// 3. Builder pattern methods:
//    - WithFillMode(), WithHorizontalAlignment(), WithVerticalAlignment()
//    - Immutability testing (original objects unchanged)
//    - Method chaining capabilities
//
// 4. Convenience methods:
//    - Fill mode: FillCover(), FillRepeatHorizontally(), FillRepeatVertically(), FillRepeat()
//    - Horizontal alignment: AlignLeft(), AlignCenter(), AlignRight()
//    - Vertical alignment: AlignTop(), AlignMiddle(), AlignBottom()
//
// 5. JSON serialization and deserialization:
//    - BackgroundImageValue dual-form marshaling (string vs object)
//    - UnmarshalJSON with validation during deserialization
//    - Edge cases and error handling
//    - Round-trip testing
//
// 6. Integration and workflow testing:
//    - Complete workflow from creation to validation to modification
//    - Real-world usage patterns
//
// 7. Edge cases and error conditions:
//    - URL validation edge cases (malformed URLs, invalid schemes)
//    - Enum validation edge cases
//    - JSON parsing edge cases
//    - State consistency validation
//    - Builder pattern immutability verification
//
// Coverage achieved: ~100% for most functions, 92.9%+ for complex JSON handling

func TestNewBackgroundImage(t *testing.T) {
	tests := []struct {
		name                string
		url                 string
		fillMode            ImageFillMode
		horizontalAlignment HorizontalAlignment
		verticalAlignment   VerticalAlignment
		expectError         bool
		errorContains       string
	}{
		{
			name:                "valid with all parameters",
			url:                 "https://example.com/image.jpg",
			fillMode:            ImageFillCover,
			horizontalAlignment: HAlignCenter,
			verticalAlignment:   VAlignTop,
			expectError:         false,
		},
		{
			name:                "valid with minimal parameters",
			url:                 "https://example.com/image.png",
			fillMode:            "",
			horizontalAlignment: "",
			verticalAlignment:   "",
			expectError:         false,
		},
		{
			name:                "invalid URL - empty",
			url:                 "",
			fillMode:            ImageFillCover,
			horizontalAlignment: HAlignCenter,
			verticalAlignment:   VAlignTop,
			expectError:         true,
			errorContains:       "invalid BackgroundImage URL",
		},
		{
			name:                "invalid URL - malformed",
			url:                 "not-a-url",
			fillMode:            ImageFillCover,
			horizontalAlignment: HAlignCenter,
			verticalAlignment:   VAlignTop,
			expectError:         true,
			errorContains:       "invalid BackgroundImage URL",
		},
		{
			name:                "valid with invalid enum (should not error in constructor)",
			url:                 "https://example.com/image.gif",
			fillMode:            "invalid",
			horizontalAlignment: "invalid",
			verticalAlignment:   "invalid",
			expectError:         false, // Constructor is lenient, validation happens in Validate()
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bg, err := NewBackgroundImage(tt.url, tt.fillMode, tt.horizontalAlignment, tt.verticalAlignment)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
					return
				}
				if !strings.Contains(err.Error(), tt.errorContains) {
					t.Errorf("expected error to contain %q, got %q", tt.errorContains, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if bg.URL != tt.url {
				t.Errorf("expected URL %q, got %q", tt.url, bg.URL)
			}
			if bg.FillMode != tt.fillMode {
				t.Errorf("expected FillMode %q, got %q", tt.fillMode, bg.FillMode)
			}
			if bg.HorizontalAlignment != tt.horizontalAlignment {
				t.Errorf("expected HorizontalAlignment %q, got %q", tt.horizontalAlignment, bg.HorizontalAlignment)
			}
			if bg.VerticalAlignment != tt.verticalAlignment {
				t.Errorf("expected VerticalAlignment %q, got %q", tt.verticalAlignment, bg.VerticalAlignment)
			}
		})
	}
}

func TestMustNewBackgroundImage(t *testing.T) {
	t.Run("valid parameters should not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("unexpected panic: %v", r)
			}
		}()

		bg := MustNewBackgroundImage("https://example.com/image.jpg", ImageFillCover, HAlignCenter, VAlignTop)
		if bg.URL != "https://example.com/image.jpg" {
			t.Errorf("expected URL https://example.com/image.jpg, got %q", bg.URL)
		}
	})

	t.Run("invalid parameters should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic but none occurred")
			}
		}()

		MustNewBackgroundImage("", ImageFillCover, HAlignCenter, VAlignTop)
	})
}

func TestBackgroundImageValidate(t *testing.T) {
	tests := []struct {
		name          string
		bg            BackgroundImage
		expectError   bool
		errorContains string
	}{
		{
			name: "valid background image",
			bg: BackgroundImage{
				URL:                 "https://example.com/image.jpg",
				FillMode:            ImageFillCover,
				HorizontalAlignment: HAlignCenter,
				VerticalAlignment:   VAlignTop,
			},
			expectError: false,
		},
		{
			name: "valid with minimal fields",
			bg: BackgroundImage{
				URL: "https://example.com/image.png",
			},
			expectError: false,
		},
		{
			name: "empty URL",
			bg: BackgroundImage{
				URL:      "",
				FillMode: ImageFillCover,
			},
			expectError:   true,
			errorContains: "url is required and cannot be empty",
		},
		{
			name: "invalid fill mode",
			bg: BackgroundImage{
				URL:      "https://example.com/image.jpg",
				FillMode: "invalid",
			},
			expectError:   true,
			errorContains: "BackgroundImage.fillMode",
		},
		{
			name: "invalid horizontal alignment",
			bg: BackgroundImage{
				URL:                 "https://example.com/image.jpg",
				HorizontalAlignment: "invalid",
			},
			expectError:   true,
			errorContains: "BackgroundImage.horizontalAlignment",
		},
		{
			name: "invalid vertical alignment",
			bg: BackgroundImage{
				URL:               "https://example.com/image.jpg",
				VerticalAlignment: "invalid",
			},
			expectError:   true,
			errorContains: "BackgroundImage.verticalAlignment",
		},
		{
			name: "all enums empty (should be valid)",
			bg: BackgroundImage{
				URL:                 "https://example.com/image.jpg",
				FillMode:            "",
				HorizontalAlignment: "",
				VerticalAlignment:   "",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.bg.Validate()

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
					return
				}
				if !strings.Contains(err.Error(), tt.errorContains) {
					t.Errorf("expected error to contain %q, got %q", tt.errorContains, err.Error())
				}
			} else if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestBackgroundImageBuilderMethods(t *testing.T) {
	bg := BackgroundImage{URL: "https://example.com/image.jpg"}

	t.Run("WithFillMode", func(t *testing.T) {
		result := bg.WithFillMode(ImageFillRepeat)
		if result.FillMode != ImageFillRepeat {
			t.Errorf("expected FillMode %q, got %q", ImageFillRepeat, result.FillMode)
		}
		if result.URL != bg.URL {
			t.Errorf("URL should be preserved")
		}
	})

	t.Run("WithHorizontalAlignment", func(t *testing.T) {
		result := bg.WithHorizontalAlignment(HAlignRight)
		if result.HorizontalAlignment != HAlignRight {
			t.Errorf("expected HorizontalAlignment %q, got %q", HAlignRight, result.HorizontalAlignment)
		}
	})

	t.Run("WithVerticalAlignment", func(t *testing.T) {
		result := bg.WithVerticalAlignment(VAlignBottom)
		if result.VerticalAlignment != VAlignBottom {
			t.Errorf("expected VerticalAlignment %q, got %q", VAlignBottom, result.VerticalAlignment)
		}
	})

	t.Run("chaining methods", func(t *testing.T) {
		result := bg.WithFillMode(ImageFillCover).
			WithHorizontalAlignment(HAlignCenter).
			WithVerticalAlignment(VAlignTop)

		if result.FillMode != ImageFillCover {
			t.Errorf("expected FillMode %q, got %q", ImageFillCover, result.FillMode)
		}
		if result.HorizontalAlignment != HAlignCenter {
			t.Errorf("expected HorizontalAlignment %q, got %q", HAlignCenter, result.HorizontalAlignment)
		}
		if result.VerticalAlignment != VAlignTop {
			t.Errorf("expected VerticalAlignment %q, got %q", VAlignTop, result.VerticalAlignment)
		}
	})
}

func TestBackgroundImageConvenienceMethods(t *testing.T) {
	bg := BackgroundImage{URL: "https://example.com/image.jpg"}

	t.Run("fill mode convenience methods", func(t *testing.T) {
		tests := []struct {
			name     string
			method   func(BackgroundImage) BackgroundImage
			expected ImageFillMode
		}{
			{"FillCover", BackgroundImage.FillCover, ImageFillCover},
			{"FillRepeatHorizontally", BackgroundImage.FillRepeatHorizontally, ImageFillRepeatHorizontally},
			{"FillRepeatVertically", BackgroundImage.FillRepeatVertically, ImageFillRepeatVertically},
			{"FillRepeat", BackgroundImage.FillRepeat, ImageFillRepeat},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := tt.method(bg)
				if result.FillMode != tt.expected {
					t.Errorf("expected FillMode %q, got %q", tt.expected, result.FillMode)
				}
			})
		}
	})

	t.Run("horizontal alignment convenience methods", func(t *testing.T) {
		tests := []struct {
			name     string
			method   func(BackgroundImage) BackgroundImage
			expected HorizontalAlignment
		}{
			{"AlignLeft", BackgroundImage.AlignLeft, HAlignLeft},
			{"AlignCenter", BackgroundImage.AlignCenter, HAlignCenter},
			{"AlignRight", BackgroundImage.AlignRight, HAlignRight},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := tt.method(bg)
				if result.HorizontalAlignment != tt.expected {
					t.Errorf("expected HorizontalAlignment %q, got %q", tt.expected, result.HorizontalAlignment)
				}
			})
		}
	})

	t.Run("vertical alignment convenience methods", func(t *testing.T) {
		tests := []struct {
			name     string
			method   func(BackgroundImage) BackgroundImage
			expected VerticalAlignment
		}{
			{"AlignTop", BackgroundImage.AlignTop, VAlignTop},
			{"AlignMiddle", BackgroundImage.AlignMiddle, VAlignCenter},
			{"AlignBottom", BackgroundImage.AlignBottom, VAlignBottom},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := tt.method(bg)
				if result.VerticalAlignment != tt.expected {
					t.Errorf("expected VerticalAlignment %q, got %q", tt.expected, result.VerticalAlignment)
				}
			})
		}
	})
}

func TestBackgroundImageURL(t *testing.T) {
	tests := []struct {
		name          string
		url           string
		expectError   bool
		errorContains string
	}{
		{
			name:        "valid URL",
			url:         "https://example.com/image.jpg",
			expectError: false,
		},
		{
			name:          "empty URL",
			url:           "",
			expectError:   true,
			errorContains: "invalid BackgroundImage URL",
		},
		{
			name:          "invalid URL",
			url:           "not-a-url",
			expectError:   true,
			errorContains: "invalid BackgroundImage URL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := BackgroundImageURL(tt.url)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
					return
				}
				if !strings.Contains(err.Error(), tt.errorContains) {
					t.Errorf("expected error to contain %q, got %q", tt.errorContains, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if result == nil {
				t.Errorf("expected non-nil result")
				return
			}

			if result.urlOnly != tt.url {
				t.Errorf("expected urlOnly %q, got %q", tt.url, result.urlOnly)
			}
			if result.obj != nil {
				t.Errorf("expected obj to be nil for URL-only value")
			}
		})
	}
}

func TestMustBackgroundImageURL(t *testing.T) {
	t.Run("valid URL should not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("unexpected panic: %v", r)
			}
		}()

		result := MustBackgroundImageURL("https://example.com/image.jpg")
		if result.urlOnly != "https://example.com/image.jpg" {
			t.Errorf("expected urlOnly https://example.com/image.jpg, got %q", result.urlOnly)
		}
	})

	t.Run("invalid URL should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic but none occurred")
			}
		}()

		MustBackgroundImageURL("")
	})
}

func TestBackgroundImageObject(t *testing.T) {
	tests := []struct {
		name          string
		obj           BackgroundImage
		expectError   bool
		errorContains string
	}{
		{
			name: "valid object",
			obj: BackgroundImage{
				URL:                 "https://example.com/image.jpg",
				FillMode:            ImageFillCover,
				HorizontalAlignment: HAlignCenter,
				VerticalAlignment:   VAlignTop,
			},
			expectError: false,
		},
		{
			name: "invalid object - empty URL",
			obj: BackgroundImage{
				URL:      "",
				FillMode: ImageFillCover,
			},
			expectError:   true,
			errorContains: "url is required and cannot be empty",
		},
		{
			name: "invalid object - invalid enum",
			obj: BackgroundImage{
				URL:      "https://example.com/image.jpg",
				FillMode: "invalid",
			},
			expectError:   true,
			errorContains: "BackgroundImage.fillMode",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := BackgroundImageObject(tt.obj)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
					return
				}
				if !strings.Contains(err.Error(), tt.errorContains) {
					t.Errorf("expected error to contain %q, got %q", tt.errorContains, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if result == nil {
				t.Errorf("expected non-nil result")
				return
			}

			if result.obj == nil {
				t.Errorf("expected obj to be non-nil for object value")
				return
			}
			if result.urlOnly != "" {
				t.Errorf("expected urlOnly to be empty for object value")
			}

			if result.obj.URL != tt.obj.URL {
				t.Errorf("expected URL %q, got %q", tt.obj.URL, result.obj.URL)
			}
		})
	}
}

func TestMustBackgroundImageObject(t *testing.T) {
	t.Run("valid object should not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("unexpected panic: %v", r)
			}
		}()

		obj := BackgroundImage{
			URL:      "https://example.com/image.jpg",
			FillMode: ImageFillCover,
		}
		result := MustBackgroundImageObject(obj)
		if result.obj.URL != obj.URL {
			t.Errorf("expected URL %q, got %q", obj.URL, result.obj.URL)
		}
	})

	t.Run("invalid object should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic but none occurred")
			}
		}()

		obj := BackgroundImage{URL: ""} // Invalid - empty URL
		MustBackgroundImageObject(obj)
	})
}

func TestBackgroundImageValueMarshalJSON(t *testing.T) {
	t.Run("URL-only value", func(t *testing.T) {
		value := &BackgroundImageValue{urlOnly: "https://example.com/image.jpg"}

		data, err := json.Marshal(value)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}

		expected := `"https://example.com/image.jpg"`
		if string(data) != expected {
			t.Errorf("expected JSON %s, got %s", expected, string(data))
		}
	})

	t.Run("object value", func(t *testing.T) {
		obj := &BackgroundImage{
			URL:                 "https://example.com/image.jpg",
			FillMode:            ImageFillCover,
			HorizontalAlignment: HAlignCenter,
			VerticalAlignment:   VAlignTop,
		}
		value := &BackgroundImageValue{obj: obj}

		data, err := json.Marshal(value)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}

		// Unmarshal to verify structure
		var result map[string]any
		if err := json.Unmarshal(data, &result); err != nil {
			t.Errorf("failed to unmarshal result: %v", err)
			return
		}

		if result["url"] != "https://example.com/image.jpg" {
			t.Errorf("expected URL https://example.com/image.jpg, got %v", result["url"])
		}
		if result["fillMode"] != "cover" {
			t.Errorf("expected fillMode cover, got %v", result["fillMode"])
		}
		if result["horizontalAlignment"] != "center" {
			t.Errorf("expected horizontalAlignment center, got %v", result["horizontalAlignment"])
		}
		if result["verticalAlignment"] != "top" {
			t.Errorf("expected verticalAlignment top, got %v", result["verticalAlignment"])
		}
	})
}

func TestBackgroundImageValueUnmarshalJSON(t *testing.T) {
	t.Run("string form", func(t *testing.T) {
		jsonData := `"https://example.com/image.jpg"`
		var value BackgroundImageValue

		err := json.Unmarshal([]byte(jsonData), &value)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}

		if value.urlOnly != "https://example.com/image.jpg" {
			t.Errorf("expected urlOnly https://example.com/image.jpg, got %q", value.urlOnly)
		}
		if value.obj != nil {
			t.Errorf("expected obj to be nil for string form")
		}
	})

	t.Run("object form", func(t *testing.T) {
		jsonData := `{
			"url": "https://example.com/image.jpg",
			"fillMode": "cover",
			"horizontalAlignment": "center",
			"verticalAlignment": "top"
		}`
		var value BackgroundImageValue

		err := json.Unmarshal([]byte(jsonData), &value)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}

		if value.urlOnly != "" {
			t.Errorf("expected urlOnly to be empty for object form")
		}
		if value.obj == nil {
			t.Errorf("expected obj to be non-nil for object form")
			return
		}

		if value.obj.URL != "https://example.com/image.jpg" {
			t.Errorf("expected URL https://example.com/image.jpg, got %q", value.obj.URL)
		}
		if value.obj.FillMode != ImageFillCover {
			t.Errorf("expected FillMode %q, got %q", ImageFillCover, value.obj.FillMode)
		}
		if value.obj.HorizontalAlignment != HAlignCenter {
			t.Errorf("expected HorizontalAlignment %q, got %q", HAlignCenter, value.obj.HorizontalAlignment)
		}
		if value.obj.VerticalAlignment != VAlignTop {
			t.Errorf("expected VerticalAlignment %q, got %q", VAlignTop, value.obj.VerticalAlignment)
		}
	})

	t.Run("invalid string form", func(t *testing.T) {
		jsonData := `""`
		var value BackgroundImageValue

		err := json.Unmarshal([]byte(jsonData), &value)
		if err == nil {
			t.Errorf("expected error for empty URL but got none")
			return
		}
		if !strings.Contains(err.Error(), "invalid BackgroundImage URL") {
			t.Errorf("expected error to contain 'invalid BackgroundImage URL', got %q", err.Error())
		}
	})

	t.Run("invalid object form", func(t *testing.T) {
		jsonData := `{
			"url": "https://example.com/image.jpg",
			"fillMode": "invalid"
		}`
		var value BackgroundImageValue

		err := json.Unmarshal([]byte(jsonData), &value)
		if err == nil {
			t.Errorf("expected error for invalid fillMode but got none")
			return
		}
		if !strings.Contains(err.Error(), "invalid ImageFillMode") {
			t.Errorf("expected error to contain 'invalid ImageFillMode', got %q", err.Error())
		}
	})

	t.Run("invalid JSON", func(t *testing.T) {
		jsonData := `{invalid json}`
		var value BackgroundImageValue

		err := json.Unmarshal([]byte(jsonData), &value)
		if err == nil {
			t.Errorf("expected error for invalid JSON but got none")
		}
	})
}

func TestBackgroundImageIntegration(t *testing.T) {
	t.Run("complete workflow", func(t *testing.T) {
		// Create a background image using constructor
		bg, err := NewBackgroundImage(
			"https://example.com/image.jpg",
			ImageFillCover,
			HAlignCenter,
			VAlignTop,
		)
		if err != nil {
			t.Errorf("unexpected error creating background image: %v", err)
			return
		}

		// Validate it
		if err := bg.Validate(); err != nil {
			t.Errorf("unexpected validation error: %v", err)
			return
		}

		// Use builder methods to modify
		modified := bg.FillRepeat().AlignRight().AlignBottom()

		// Validate modified version
		if err := modified.Validate(); err != nil {
			t.Errorf("unexpected validation error for modified: %v", err)
			return
		}

		// Check modifications
		if modified.FillMode != ImageFillRepeat {
			t.Errorf("expected FillMode %q, got %q", ImageFillRepeat, modified.FillMode)
		}
		if modified.HorizontalAlignment != HAlignRight {
			t.Errorf("expected HorizontalAlignment %q, got %q", HAlignRight, modified.HorizontalAlignment)
		}
		if modified.VerticalAlignment != VAlignBottom {
			t.Errorf("expected VerticalAlignment %q, got %q", VAlignBottom, modified.VerticalAlignment)
		}

		// Original should be unchanged
		if bg.FillMode != ImageFillCover {
			t.Errorf("original should be unchanged, expected FillMode %q, got %q", ImageFillCover, bg.FillMode)
		}
	})

	t.Run("JSON roundtrip with BackgroundImageValue", func(t *testing.T) {
		// Test string form
		stringValue, err := BackgroundImageURL("https://example.com/image.jpg")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}

		data, err := json.Marshal(stringValue)
		if err != nil {
			t.Errorf("unexpected marshal error: %v", err)
			return
		}

		var unmarshaled BackgroundImageValue
		if err := json.Unmarshal(data, &unmarshaled); err != nil {
			t.Errorf("unexpected unmarshal error: %v", err)
			return
		}

		if unmarshaled.urlOnly != "https://example.com/image.jpg" {
			t.Errorf("expected urlOnly https://example.com/image.jpg, got %q", unmarshaled.urlOnly)
		}

		// Test object form
		bg := MustNewBackgroundImage("https://example.com/image.jpg", ImageFillCover, HAlignCenter, VAlignTop)
		objectValue, err := BackgroundImageObject(bg)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}

		data, err = json.Marshal(objectValue)
		if err != nil {
			t.Errorf("unexpected marshal error: %v", err)
			return
		}

		var unmarshaledObject BackgroundImageValue
		if err := json.Unmarshal(data, &unmarshaledObject); err != nil {
			t.Errorf("unexpected unmarshal error: %v", err)
			return
		}

		if unmarshaledObject.obj == nil {
			t.Errorf("expected obj to be non-nil")
			return
		}
		if unmarshaledObject.obj.URL != bg.URL {
			t.Errorf("expected URL %q, got %q", bg.URL, unmarshaledObject.obj.URL)
		}
	})
}

func TestBackgroundImageEdgeCases(t *testing.T) {
	t.Run("URL validation with malformed URLs", func(t *testing.T) {
		invalidURLs := []string{
			"ht tp://invalid space.com",
			"ftp://not-http.com",
			"javascript:alert('xss')",
			"data:text/html,<script>alert(1)</script>",
		}

		for _, url := range invalidURLs {
			t.Run("invalid_url_"+url, func(t *testing.T) {
				_, err := NewBackgroundImage(url, ImageFillCover, HAlignCenter, VAlignTop)
				if err == nil {
					t.Errorf("expected error for URL %q but got none", url)
				}
			})
		}
	})

	t.Run("validate with malformed URL in existing object", func(t *testing.T) {
		// Create an object with manually set invalid URL to test defense-in-depth validation
		bg := BackgroundImage{
			URL:      "ht tp://invalid space.com", // This bypasses constructor validation
			FillMode: ImageFillCover,
		}

		err := bg.Validate()
		if err == nil {
			t.Errorf("expected validation error for malformed URL but got none")
		}
		if !strings.Contains(err.Error(), "BackgroundImage.url") {
			t.Errorf("expected error to mention BackgroundImage.url, got %q", err.Error())
		}
	})

	t.Run("enum edge cases", func(t *testing.T) {
		bg := BackgroundImage{URL: "https://example.com/image.jpg"}

		// Test with borderline values that should be valid
		validBg := bg.WithFillMode(ImageFillCover).
			WithHorizontalAlignment(HAlignLeft).
			WithVerticalAlignment(VAlignTop)

		if err := validBg.Validate(); err != nil {
			t.Errorf("unexpected error for valid enums: %v", err)
		}

		// Test with empty enum values (should be valid)
		emptyBg := bg.WithFillMode("").
			WithHorizontalAlignment("").
			WithVerticalAlignment("")

		if err := emptyBg.Validate(); err != nil {
			t.Errorf("unexpected error for empty enums: %v", err)
		}
	})

	t.Run("JSON edge cases", func(t *testing.T) {
		// Test with minimal JSON object
		minimalJSON := `{"url": "https://example.com/image.jpg"}`
		var value BackgroundImageValue
		if err := json.Unmarshal([]byte(minimalJSON), &value); err != nil {
			t.Errorf("unexpected error for minimal JSON: %v", err)
		}

		if value.obj == nil {
			t.Errorf("expected obj to be non-nil for minimal JSON object")
		}

		// Test with object that has empty enum values (should cause validation error during unmarshaling)
		emptyEnumJSON := `{
			"url": "https://example.com/image.jpg",
			"fillMode": "",
			"horizontalAlignment": "",
			"verticalAlignment": ""
		}`
		var emptyEnumValue BackgroundImageValue
		if err := json.Unmarshal([]byte(emptyEnumJSON), &emptyEnumValue); err == nil {
			t.Errorf("expected error for empty enum in JSON but got none")
		}

		// Test string with encoded spaces (should be valid as URL encoding handles this)
		spacedStringJSON := `"https://example.com/image%20with%20spaces.jpg"`
		var spacedValue BackgroundImageValue
		if err := json.Unmarshal([]byte(spacedStringJSON), &spacedValue); err != nil {
			t.Errorf("unexpected error for URL-encoded string: %v", err)
		}
	})

	t.Run("builder pattern immutability", func(t *testing.T) {
		original := BackgroundImage{
			URL:                 "https://example.com/image.jpg",
			FillMode:            ImageFillCover,
			HorizontalAlignment: HAlignCenter,
			VerticalAlignment:   VAlignTop,
		}

		// Test that builder methods don't modify original
		modified := original.WithFillMode(ImageFillRepeat)
		if original.FillMode != ImageFillCover {
			t.Errorf("original object was modified when it should be immutable")
		}
		if modified.FillMode != ImageFillRepeat {
			t.Errorf("modified object doesn't have expected value")
		}

		// Test chaining doesn't affect original
		chained := original.FillRepeat().AlignRight().AlignBottom()
		if original.FillMode != ImageFillCover || original.HorizontalAlignment != HAlignCenter || original.VerticalAlignment != VAlignTop {
			t.Errorf("chained operations modified original object")
		}
		if chained.FillMode != ImageFillRepeat || chained.HorizontalAlignment != HAlignRight || chained.VerticalAlignment != VAlignBottom {
			t.Errorf("chained operations didn't produce expected result")
		}
	})

	t.Run("BackgroundImageValue state consistency", func(t *testing.T) {
		// Test that BackgroundImageValue maintains proper state
		urlValue := MustBackgroundImageURL("https://example.com/image.jpg")
		if urlValue.obj != nil {
			t.Errorf("URL-only value should have nil obj")
		}
		if urlValue.urlOnly == "" {
			t.Errorf("URL-only value should have non-empty urlOnly")
		}

		bg := MustNewBackgroundImage("https://example.com/image.jpg", ImageFillCover, HAlignCenter, VAlignTop)
		objValue := MustBackgroundImageObject(bg)
		if objValue.obj == nil {
			t.Errorf("Object value should have non-nil obj")
		}
		if objValue.urlOnly != "" {
			t.Errorf("Object value should have empty urlOnly")
		}
	})
}
