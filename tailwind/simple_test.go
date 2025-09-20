package tailwind

import (
	"testing"
)

func TestSimpleTypedConfig(t *testing.T) {
	// Test that we can create a default config
	config := DefaultConfig()
	if config.Theme.Colors.Red50 == nil {
		t.Error("Expected Red50 to be defined in default config")
	}

	// Test that we can access color values
	red50 := config.Theme.Colors.Red50
	if red50.String() != "red-50" {
		t.Errorf("Expected Red50 name to be 'red-50', got %s", red50.String())
	}

	// Test that we can convert to CSS value
	cssValue := red50.ToCSSValue()
	if cssValue == nil {
		t.Error("Expected CSS value to be non-nil")
	}

	// Test spacing config
	if config.Theme.Spacing.Size1 == nil {
		t.Error("Expected Spacing.Size1 to be defined")
	}

	// Test font size config
	if config.Theme.FontSize.Xs.Size == nil {
		t.Error("Expected FontSize.Xs.Size to be defined")
	}

	t.Logf("✅ Basic typed config structure works correctly")
}

func TestConfigValueTypes(t *testing.T) {
	// Test StaticColor
	color := ColorFromHex("test-color", "#ff0000")
	if color.Name != "test-color" {
		t.Errorf("Expected color name to be 'test-color', got %s", color.Name)
	}

	// Test StaticLength
	length := LengthFromPx("test-length", 16)
	if length.Name != "test-length" {
		t.Errorf("Expected length name to be 'test-length', got %s", length.Name)
	}

	// Test that CSS values are generated properly
	cssColorValue := color.ToCSSValue()
	if cssColorValue.String() != "#ff0000" {
		t.Errorf("Expected CSS color value to be '#ff0000', got %s", cssColorValue.String())
	}

	cssLengthValue := length.ToCSSValue()
	if cssLengthValue.String() != "16px" {
		t.Errorf("Expected CSS length value to be '16px', got %s", cssLengthValue.String())
	}

	t.Logf("✅ Value type conversions work correctly")
}
