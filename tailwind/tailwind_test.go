package tailwind

import (
	"strings"
	"testing"

	"github.com/ahmed-com/typesafe-css/css"
)

func TestDefaultTheme(t *testing.T) {
	theme := DefaultTheme()
	
	// Test some basic theme values
	if theme.Colors["blue-500"] != css.Hex("#3b82f6") {
		t.Errorf("Expected blue-500 to be #3b82f6, got %s", theme.Colors["blue-500"])
	}
	
	if theme.Spacing["4"] != css.Rem(1) {
		t.Errorf("Expected spacing 4 to be 1rem, got %s", theme.Spacing["4"])
	}
	
	if theme.FontSizes["lg"] != css.Rem(1.125) {
		t.Errorf("Expected font size lg to be 1.125rem, got %s", theme.FontSizes["lg"])
	}
}

func TestCustomTheme(t *testing.T) {
	custom := CustomTheme(&Theme{
		Colors: map[string]css.Color{
			"brand": css.Hex("#123456"),
		},
	})
	
	// Should have custom color
	if custom.Colors["brand"] != css.Hex("#123456") {
		t.Errorf("Expected brand color to be #123456, got %s", custom.Colors["brand"])
	}
	
	// Should still have default colors
	if custom.Colors["blue-500"] != css.Hex("#3b82f6") {
		t.Errorf("Expected blue-500 to still be #3b82f6, got %s", custom.Colors["blue-500"])
	}
}

func TestUtilityManager(t *testing.T) {
	manager := NewUtilityManager(DefaultTheme())
	
	// Test that calling the same utility multiple times returns the same rule
	rule1 := Padding(manager, "4")
	rule2 := Padding(manager, "4")
	
	if rule1.String() != rule2.String() {
		t.Errorf("Expected same rule content for duplicate utility calls")
	}
	
	// Test that manager caches rules
	rules := manager.GetRules()
	if len(rules) != 1 {
		t.Errorf("Expected 1 rule cached, got %d", len(rules))
	}
}

func TestColorUtilities(t *testing.T) {
	manager := NewUtilityManager(DefaultTheme())
	
	// Test text color utility
	textRule := TextColor(manager, "blue-500")
	expected := ".text-blue-500{color:#3b82f6}"
	if textRule.String() != expected {
		t.Errorf("Expected %s, got %s", expected, textRule.String())
	}
	
	// Test background color utility
	bgRule := BackgroundColor(manager, "red-500")
	expected = ".bg-red-500{background-color:#ef4444}"
	if bgRule.String() != expected {
		t.Errorf("Expected %s, got %s", expected, bgRule.String())
	}
	
	// Test border color utility
	borderRule := BorderColor(manager, "gray-300")
	expected = ".border-gray-300{border-color:#d1d5db}"
	if borderRule.String() != expected {
		t.Errorf("Expected %s, got %s", expected, borderRule.String())
	}
}

func TestSpacingUtilities(t *testing.T) {
	manager := NewUtilityManager(DefaultTheme())
	
	// Test padding utility
	paddingRule := Padding(manager, "4")
	expected := ".p-4{padding:1rem}"
	if paddingRule.String() != expected {
		t.Errorf("Expected %s, got %s", expected, paddingRule.String())
	}
	
	// Test padding X utility
	pxRule := PaddingX(manager, "2")
	if !strings.Contains(pxRule.String(), "padding-left:0.5rem") {
		t.Errorf("Expected padding-left in px rule, got %s", pxRule.String())
	}
	if !strings.Contains(pxRule.String(), "padding-right:0.5rem") {
		t.Errorf("Expected padding-right in px rule, got %s", pxRule.String())
	}
	
	// Test width utility with special values
	wFullRule := Width(manager, "full")
	expected = ".w-full{width:100%}"
	if wFullRule.String() != expected {
		t.Errorf("Expected %s, got %s", expected, wFullRule.String())
	}
	
	wAutoRule := Width(manager, "auto")
	expected = ".w-auto{width:auto}"
	if wAutoRule.String() != expected {
		t.Errorf("Expected %s, got %s", expected, wAutoRule.String())
	}
}

func TestLayoutUtilities(t *testing.T) {
	// Test new direct API
	flexRule := Flex()
	expected := ".flex{display:flex}"
	if flexRule.String() != expected {
		t.Errorf("Expected %s, got %s", expected, flexRule.String())
	}
	
	blockRule := Block()
	expected = ".block{display:block}"
	if blockRule.String() != expected {
		t.Errorf("Expected %s, got %s", expected, blockRule.String())
	}
	
	// Test flexbox utilities
	justifyCenterRule := JustifyCenter()
	expected = ".justify-center{justify-content:center}"
	if justifyCenterRule.String() != expected {
		t.Errorf("Expected %s, got %s", expected, justifyCenterRule.String())
	}
}

func TestTypographyUtilities(t *testing.T) {
	manager := NewUtilityManager(DefaultTheme())
	
	// Test text size utility
	textLgRule := TextSize(manager, "lg")
	expected := ".text-lg{font-size:1.125rem}"
	if textLgRule.String() != expected {
		t.Errorf("Expected %s, got %s", expected, textLgRule.String())
	}
	
	// Test font weight utility
	fontBoldRule := FontWeight(manager, "bold")
	expected = ".font-bold{font-weight:700}"
	if fontBoldRule.String() != expected {
		t.Errorf("Expected %s, got %s", expected, fontBoldRule.String())
	}
	
	// Test text align utility
	textCenterRule := TextAlign(manager, "center")
	expected = ".text-center{text-align:center}"
	if textCenterRule.String() != expected {
		t.Errorf("Expected %s, got %s", expected, textCenterRule.String())
	}
}

func TestConvenienceFunctions(t *testing.T) {
	// Test that new direct API functions work and generate expected CSS
	bgRule := BgBlue500()
	if !strings.Contains(bgRule.String(), "background-color:#3b82f6") {
		t.Errorf("Expected bg-blue-500 to contain background-color:#3b82f6, got %s", bgRule.String())
	}
	
	textRule := TextWhite()
	if !strings.Contains(textRule.String(), "color:#ffffff") {
		t.Errorf("Expected text-white to contain color:#ffffff, got %s", textRule.String())
	}
	
	pRule := P4()
	if !strings.Contains(pRule.String(), "padding:1rem") {
		t.Errorf("Expected p-4 to contain padding:1rem, got %s", pRule.String())
	}

	// Test arbitrary value support
	customBgRule := Bg("#bada55")
	if !strings.Contains(customBgRule.String(), "background-color:#bada55") {
		t.Errorf("Expected custom bg to contain background-color:#bada55, got %s", customBgRule.String())
	}
}

func TestGenerateUtilityStylesheet(t *testing.T) {
	stylesheet := GenerateUtilityStylesheet()
	
	// Should generate a significant number of utilities
	if len(stylesheet.Items) < 100 {
		t.Errorf("Expected at least 100 utility classes, got %d", len(stylesheet.Items))
	}
	
	// Check that some common utilities are included
	cssOutput := stylesheet.String()
	expectedUtilities := []string{
		".flex{display:flex}",
		".block{display:block}",
		".hidden{display:none}",
		".text-center{text-align:center}",
		".font-bold{font-weight:700}",
	}
	
	for _, expected := range expectedUtilities {
		if !strings.Contains(cssOutput, expected) {
			t.Errorf("Expected utility stylesheet to contain %s", expected)
		}
	}
}

func TestClassName(t *testing.T) {
	// Test class name generation
	tests := []struct {
		prefix   string
		value    string
		expected string
	}{
		{"text", "lg", "text-lg"},
		{"bg", "blue-500", "bg-blue-500"},
		{"p", "4", "p-4"},
		{"m", "", "m"},
	}
	
	for _, test := range tests {
		result := ClassName(test.prefix, test.value)
		if result != test.expected {
			t.Errorf("ClassName(%q, %q) = %q, expected %q", test.prefix, test.value, result, test.expected)
		}
	}
}