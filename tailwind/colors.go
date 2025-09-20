package tailwind

import (
	"fmt"

	"github.com/ahmed-com/typesafe-css/css"
)

// TextColor creates a text color utility class.
// Example: TextColor(manager, "blue-500") generates ".text-blue-500 { color: #3b82f6; }"
func TextColor(manager *UtilityManager, colorKey string) css.Rule {
	className := ClassName("text", colorKey)
	
	return manager.GetOrCreateRule(className, func() css.Rule {
		color, exists := manager.Theme().Colors[colorKey]
		if !exists {
			// Fallback to treating the key as a direct color value
			color = css.Hex(colorKey)
		}
		
		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set(css.ColorP, color),
		)
	})
}

// BackgroundColor creates a background color utility class.
// Example: BackgroundColor(manager, "red-500") generates ".bg-red-500 { background-color: #ef4444; }"
func BackgroundColor(manager *UtilityManager, colorKey string) css.Rule {
	className := ClassName("bg", colorKey)
	
	return manager.GetOrCreateRule(className, func() css.Rule {
		color, exists := manager.Theme().Colors[colorKey]
		if !exists {
			// Fallback to treating the key as a direct color value
			color = css.Hex(colorKey)
		}
		
		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set(css.BackgroundColor, color),
		)
	})
}

// BorderColor creates a border color utility class.
// Example: BorderColor(manager, "gray-300") generates ".border-gray-300 { border-color: #d1d5db; }"
func BorderColor(manager *UtilityManager, colorKey string) css.Rule {
	className := ClassName("border", colorKey)
	
	return manager.GetOrCreateRule(className, func() css.Rule {
		color, exists := manager.Theme().Colors[colorKey]
		if !exists {
			// Fallback to treating the key as a direct color value
			color = css.Hex(colorKey)
		}
		
		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set("border-color", color),
		)
	})
}

// Convenience functions that use the default manager

// Text creates a text color utility using the default manager.
func Text(colorKey string) css.Rule {
	return TextColor(defaultManager, colorKey)
}

// Border creates a border color utility using the default manager.
func Border(colorKey string) css.Rule {
	return BorderColor(defaultManager, colorKey)
}