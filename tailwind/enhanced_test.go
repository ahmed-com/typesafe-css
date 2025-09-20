package tailwind

import (
	"testing"
	
	"github.com/ahmed-com/typesafe-css/css"
)

func TestEnhancedTheme(t *testing.T) {
	theme := DefaultTheme()

	// Test comprehensive color palette
	expectedColors := []string{
		"inherit", "current", "transparent", "black", "white",
		"slate-500", "gray-500", "zinc-500", "neutral-500", "stone-500",
		"red-500", "orange-500", "amber-500", "yellow-500", "lime-500",
		"green-500", "emerald-500", "teal-500", "cyan-500", "sky-500",
		"blue-500", "indigo-500", "violet-500", "purple-500", "fuchsia-500",
		"pink-500", "rose-500",
	}
	
	for _, color := range expectedColors {
		if _, exists := theme.Colors[color]; !exists {
			t.Errorf("Expected color %s to exist in theme", color)
		}
	}

	// Test new theme categories
	if len(theme.BorderRadius) == 0 {
		t.Error("Expected BorderRadius to be populated")
	}
	
	if len(theme.BorderWidth) == 0 {
		t.Error("Expected BorderWidth to be populated")
	}
	
	if len(theme.BoxShadow) == 0 {
		t.Error("Expected BoxShadow to be populated")
	}
	
	if len(theme.Opacity) == 0 {
		t.Error("Expected Opacity to be populated")
	}
	
	if len(theme.FontWeight) == 0 {
		t.Error("Expected FontWeight to be populated")
	}
	
	if len(theme.LineHeight) == 0 {
		t.Error("Expected LineHeight to be populated")
	}
	
	if len(theme.ZIndex) == 0 {
		t.Error("Expected ZIndex to be populated")
	}
	
	// Test specific values
	if theme.BorderRadius["lg"] != css.Rem(0.5) {
		t.Errorf("Expected BorderRadius['lg'] to be 0.5rem, got %v", theme.BorderRadius["lg"])
	}
	
	if theme.Opacity["50"] != "0.5" {
		t.Errorf("Expected Opacity['50'] to be '0.5', got %v", theme.Opacity["50"])
	}
	
	if theme.FontWeight["bold"] != "700" {
		t.Errorf("Expected FontWeight['bold'] to be '700', got %v", theme.FontWeight["bold"])
	}
}

func TestNewUtilities(t *testing.T) {
	// Test border radius utilities
	rule := RoundedLg()
	if rule.Selector != ".rounded-lg" {
		t.Errorf("Expected selector '.rounded-lg', got %s", rule.Selector)
	}
	
	// Test shadow utilities
	rule = ShadowLg()
	if rule.Selector != ".shadow-lg" {
		t.Errorf("Expected selector '.shadow-lg', got %s", rule.Selector)
	}
	
	// Test opacity utilities
	rule = Opacity50()
	if rule.Selector != ".opacity-50" {
		t.Errorf("Expected selector '.opacity-50', got %s", rule.Selector)
	}
	
	// Test z-index utilities
	rule = Z10()
	if rule.Selector != ".z-10" {
		t.Errorf("Expected selector '.z-10', got %s", rule.Selector)
	}
	
	// Test filter utilities
	rule = BlurClass("sm")
	if rule.Selector != ".blur-sm" {
		t.Errorf("Expected selector '.blur-sm', got %s", rule.Selector)
	}
}

func TestCustomThemeWithNewCategories(t *testing.T) {
	customTheme := CustomTheme(&Theme{
		Colors: map[string]css.Color{
			"brand": css.Hex("#6366f1"),
		},
		BorderRadius: map[string]css.Length{
			"huge": css.Rem(5),
		},
		Opacity: map[string]string{
			"15": "0.15",
		},
		FontWeight: map[string]string{
			"super": "950",
		},
	})
	
	// Should have the custom values
	if customTheme.Colors["brand"] != css.Hex("#6366f1") {
		t.Error("Custom color not merged correctly")
	}
	
	if customTheme.BorderRadius["huge"] != css.Rem(5) {
		t.Error("Custom border radius not merged correctly")
	}
	
	if customTheme.Opacity["15"] != "0.15" {
		t.Error("Custom opacity not merged correctly")
	}
	
	if customTheme.FontWeight["super"] != "950" {
		t.Error("Custom font weight not merged correctly")
	}
	
	// Should still have the default values
	if customTheme.Colors["blue-500"] != css.Hex("#3b82f6") {
		t.Error("Default color not preserved in custom theme")
	}
	
	if customTheme.BorderRadius["lg"] != css.Rem(0.5) {
		t.Error("Default border radius not preserved in custom theme")
	}
}

func TestThemeBasedUtilities(t *testing.T) {
	// Create a custom theme
	customTheme := &Theme{
		Colors: map[string]css.Color{
			"custom": css.Hex("#123456"),
		},
		BorderRadius: map[string]css.Length{
			"custom": css.Rem(2),
		},
		Opacity: map[string]string{
			"custom": "0.33",
		},
	}
	
	manager := NewUtilityManager(customTheme)
	
	// Test that utilities use the custom theme
	bgRule := BackgroundColor(manager, "custom")
	if bgRule.Selector != ".bg-custom" {
		t.Errorf("Expected selector '.bg-custom', got %s", bgRule.Selector)
	}
	
	radiusRule := BorderRadius(manager, "custom")
	if radiusRule.Selector != ".rounded-custom" {
		t.Errorf("Expected selector '.rounded-custom', got %s", radiusRule.Selector)
	}
	
	opacityRule := Opacity(manager, "custom")
	if opacityRule.Selector != ".opacity-custom" {
		t.Errorf("Expected selector '.opacity-custom', got %s", opacityRule.Selector)
	}
}