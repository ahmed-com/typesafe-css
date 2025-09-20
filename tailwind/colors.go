package tailwind

import (
	"fmt"
	"strings"

	"github.com/ahmed-com/typesafe-css/css"
	"github.com/ahmed-com/typesafe-css/cssgen"
)

// TextColorToken creates a text color utility class from a strongly typed color token.
func TextColorToken(manager *UtilityManager, token ColorToken) css.Rule {
	return colorTokenRule(manager, "text", token, cssgen.Color)
}

// TextColor resolves a color key against the theme palette (including custom additions)
// and creates the corresponding text color utility. Unknown keys fall back to raw CSS
// color parsing to maintain backward compatibility with arbitrary inputs.
func TextColor(manager *UtilityManager, colorKey string) css.Rule {
	if token, ok := manager.Theme().ColorByName(colorKey); ok {
		return TextColorToken(manager, token)
	}
	return colorStringRule(manager, "text", colorKey, cssgen.Color)
}

// BackgroundColorToken creates a background color utility from a typed color token.
func BackgroundColorToken(manager *UtilityManager, token ColorToken) css.Rule {
	return colorTokenRule(manager, "bg", token, cssgen.BackgroundColor)
}

// BackgroundColor resolves a color key against the theme palette and creates a
// background-color utility class.
func BackgroundColor(manager *UtilityManager, colorKey string) css.Rule {
	if token, ok := manager.Theme().ColorByName(colorKey); ok {
		return BackgroundColorToken(manager, token)
	}
	return colorStringRule(manager, "bg", colorKey, cssgen.BackgroundColor)
}

// BorderColorToken creates a border color utility from a typed color token.
func BorderColorToken(manager *UtilityManager, token ColorToken) css.Rule {
	return colorTokenRule(manager, "border", token, cssgen.BorderColor)
}

// BorderColor resolves a color key against the palette and creates a border-color utility.
func BorderColor(manager *UtilityManager, colorKey string) css.Rule {
	if token, ok := manager.Theme().ColorByName(colorKey); ok {
		return BorderColorToken(manager, token)
	}
	return colorStringRule(manager, "border", colorKey, cssgen.BorderColor)
}

// Convenience functions for the default manager using the typed palette.

// TextToken creates a text color utility using the default manager and a typed token.
func TextToken(token ColorToken) css.Rule {
	return TextColorToken(defaultManager, token)
}

// BgColor creates a background color utility using the default manager and a typed token.
func BgColor(token ColorToken) css.Rule {
	return BackgroundColorToken(defaultManager, token)
}

// BorderToken creates a border color utility using the default manager and a typed token.
func BorderToken(token ColorToken) css.Rule {
	return BorderColorToken(defaultManager, token)
}

// Existing string-based convenience wrappers remain for backward compatibility.

// Text creates a text color utility using the default manager and a string key.
func Text(colorKey string) css.Rule {
	return TextColor(defaultManager, colorKey)
}

// Border creates a border color utility using the default manager and a string key.
func Border(colorKey string) css.Rule {
	return BorderColor(defaultManager, colorKey)
}

func colorTokenRule(manager *UtilityManager, prefix string, token ColorToken, property css.Property) css.Rule {
	className := ClassName(prefix, token.Suffix())
	return manager.GetOrCreateRule(className, func() css.Rule {
		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set(property, token.Value()),
		)
	})
}

func colorStringRule(manager *UtilityManager, prefix, colorKey string, property css.Property) css.Rule {
	className := ClassName(prefix, colorKey)
	value := cssColorFromString(colorKey)
	return manager.GetOrCreateRule(className, func() css.Rule {
		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set(property, value),
		)
	})
}

func cssColorFromString(value string) css.Color {
	if strings.HasPrefix(value, "#") {
		return css.Hex(value)
	}
	return css.Color(value)
}
