package tailwind

import (
	"testing"
)

func TestTypedConfig(t *testing.T) {
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
}

func TestUtilityGenerator(t *testing.T) {
	// Test that we can create a utility generator
	config := DefaultConfig()
	generator := NewUtilityGenerator(config)
	
	// Test generating utilities
	stylesheet := generator.GenerateUtilities()
	
	// Check that some CSS was generated
	cssOutput := stylesheet.String()
	if len(cssOutput) == 0 {
		t.Error("Expected CSS output to be non-empty")
	}
}

func TestBasicConfigStructure(t *testing.T) {
	config := DefaultConfig()
	
	// Test that major config sections are present
	if config.Theme.Colors.Red50 == nil {
		t.Error("Expected Colors.Red50 to be defined")
	}
	
	if config.Theme.Spacing.Size1 == nil {
		t.Error("Expected Spacing.Size1 to be defined")
	}
	
	if config.Theme.FontSize.Xs.Size == nil {
		t.Error("Expected FontSize.Xs.Size to be defined")
	}
}

// Test that individual value types work correctly
func TestValueTypes(t *testing.T) {
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
}