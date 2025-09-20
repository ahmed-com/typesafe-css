package tailwind

import (
	"fmt"

	"github.com/ahmed-com/typesafe-css/css"
)

// Padding creates a padding utility class.
// Example: Padding(manager, "4") generates ".p-4 { padding: 1rem; }"
func Padding(manager *UtilityManager, sizeKey string) css.Rule {
	className := ClassName("p", sizeKey)
	
	return manager.GetOrCreateRule(className, func() css.Rule {
		var size css.Value
		if length, exists := manager.Theme().Spacing[sizeKey]; exists {
			size = length
		} else {
			// Fallback to treating the key as a direct length value
			size = css.Raw(sizeKey)
		}
		
		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set(css.Padding, size),
		)
	})
}

// PaddingX creates a horizontal padding utility class.
// Example: PaddingX(manager, "4") generates ".px-4 { padding-left: 1rem; padding-right: 1rem; }"
func PaddingX(manager *UtilityManager, sizeKey string) css.Rule {
	className := ClassName("px", sizeKey)
	
	return manager.GetOrCreateRule(className, func() css.Rule {
		var size css.Value
		if length, exists := manager.Theme().Spacing[sizeKey]; exists {
			size = length
		} else {
			// Fallback to treating the key as a direct length value
			size = css.Raw(sizeKey)
		}
		
		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set("padding-left", size),
			css.Set("padding-right", size),
		)
	})
}

// PaddingY creates a vertical padding utility class.
// Example: PaddingY(manager, "4") generates ".py-4 { padding-top: 1rem; padding-bottom: 1rem; }"
func PaddingY(manager *UtilityManager, sizeKey string) css.Rule {
	className := ClassName("py", sizeKey)
	
	return manager.GetOrCreateRule(className, func() css.Rule {
		var size css.Value
		if length, exists := manager.Theme().Spacing[sizeKey]; exists {
			size = length
		} else {
			// Fallback to treating the key as a direct length value
			size = css.Raw(sizeKey)
		}
		
		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set("padding-top", size),
			css.Set("padding-bottom", size),
		)
	})
}

// Margin creates a margin utility class.
// Example: Margin(manager, "4") generates ".m-4 { margin: 1rem; }"
func Margin(manager *UtilityManager, sizeKey string) css.Rule {
	className := ClassName("m", sizeKey)
	
	return manager.GetOrCreateRule(className, func() css.Rule {
		var size css.Value
		if length, exists := manager.Theme().Spacing[sizeKey]; exists {
			size = length
		} else {
			// Fallback to treating the key as a direct length value
			size = css.Raw(sizeKey)
		}
		
		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set(css.Margin, size),
		)
	})
}

// MarginX creates a horizontal margin utility class.
// Example: MarginX(manager, "4") generates ".mx-4 { margin-left: 1rem; margin-right: 1rem; }"
func MarginX(manager *UtilityManager, sizeKey string) css.Rule {
	className := ClassName("mx", sizeKey)
	
	return manager.GetOrCreateRule(className, func() css.Rule {
		var size css.Value
		if length, exists := manager.Theme().Spacing[sizeKey]; exists {
			size = length
		} else {
			// Fallback to treating the key as a direct length value
			size = css.Raw(sizeKey)
		}
		
		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set("margin-left", size),
			css.Set("margin-right", size),
		)
	})
}

// MarginY creates a vertical margin utility class.
// Example: MarginY(manager, "4") generates ".my-4 { margin-top: 1rem; margin-bottom: 1rem; }"
func MarginY(manager *UtilityManager, sizeKey string) css.Rule {
	className := ClassName("my", sizeKey)
	
	return manager.GetOrCreateRule(className, func() css.Rule {
		var size css.Value
		if length, exists := manager.Theme().Spacing[sizeKey]; exists {
			size = length
		} else {
			// Fallback to treating the key as a direct length value
			size = css.Raw(sizeKey)
		}
		
		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set("margin-top", size),
			css.Set("margin-bottom", size),
		)
	})
}

// Width creates a width utility class.
// Example: Width(manager, "64") generates ".w-64 { width: 16rem; }"
func Width(manager *UtilityManager, sizeKey string) css.Rule {
	className := ClassName("w", sizeKey)
	
	return manager.GetOrCreateRule(className, func() css.Rule {
		var size css.Value
		if length, exists := manager.Theme().Spacing[sizeKey]; exists {
			size = length
		} else {
			// Handle special width values
			switch sizeKey {
			case "full":
				size = css.Percent(100)
			case "auto":
				size = css.Auto
			case "fit":
				size = css.Raw("fit-content")
			case "screen":
				size = css.Raw("100vw")
			default:
				// Fallback to treating the key as a direct length value
				size = css.Raw(sizeKey)
			}
		}
		
		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set(css.Width, size),
		)
	})
}

// Height creates a height utility class.
// Example: Height(manager, "64") generates ".h-64 { height: 16rem; }"
func Height(manager *UtilityManager, sizeKey string) css.Rule {
	className := ClassName("h", sizeKey)
	
	return manager.GetOrCreateRule(className, func() css.Rule {
		var size css.Value
		if length, exists := manager.Theme().Spacing[sizeKey]; exists {
			size = length
		} else {
			// Handle special height values
			switch sizeKey {
			case "full":
				size = css.Percent(100)
			case "auto":
				size = css.Auto
			case "fit":
				size = css.Raw("fit-content")
			case "screen":
				size = css.Raw("100vh")
			default:
				// Fallback to treating the key as a direct length value
				size = css.Raw(sizeKey)
			}
		}
		
		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set(css.Height, size),
		)
	})
}

// Convenience functions that use the default manager

// P creates a padding utility using the default manager.
func P(sizeKey string) css.Rule {
	return Padding(defaultManager, sizeKey)
}

// Px creates a horizontal padding utility using the default manager.
func Px(sizeKey string) css.Rule {
	return PaddingX(defaultManager, sizeKey)
}

// Py creates a vertical padding utility using the default manager.
func Py(sizeKey string) css.Rule {
	return PaddingY(defaultManager, sizeKey)
}

// M creates a margin utility using the default manager.
func M(sizeKey string) css.Rule {
	return Margin(defaultManager, sizeKey)
}

// Mx creates a horizontal margin utility using the default manager.
func Mx(sizeKey string) css.Rule {
	return MarginX(defaultManager, sizeKey)
}

// My creates a vertical margin utility using the default manager.
func My(sizeKey string) css.Rule {
	return MarginY(defaultManager, sizeKey)
}

// W creates a width utility using the default manager.
func W(sizeKey string) css.Rule {
	return Width(defaultManager, sizeKey)
}

// H creates a height utility using the default manager.
func H(sizeKey string) css.Rule {
	return Height(defaultManager, sizeKey)
}

// BorderRadius creates a border radius utility class.
// Example: BorderRadius(manager, "lg") generates ".rounded-lg { border-radius: 0.5rem; }"
func BorderRadius(manager *UtilityManager, radiusKey string) css.Rule {
	className := ClassName("rounded", radiusKey)
	
	return manager.GetOrCreateRule(className, func() css.Rule {
		var radius css.Value
		if borderRadius, exists := manager.Theme().BorderRadius[radiusKey]; exists {
			radius = borderRadius
		} else {
			// Fallback to treating the key as a direct length value
			radius = css.Raw(radiusKey)
		}
		
		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set("border-radius", radius),
		)
	})
}

// Rounded creates a border radius utility using the default manager.
func Rounded(radiusKey string) css.Rule {
	return BorderRadius(defaultManager, radiusKey)
}