package tailwind

import (
	"fmt"

	"github.com/ahmed-com/typesafe-css/css"
)

// TextSize creates a text size utility class.
// Example: TextSize(manager, "lg") generates ".text-lg { font-size: 1.125rem; }"
func TextSize(manager *UtilityManager, sizeKey string) css.Rule {
	className := ClassName("text", sizeKey)
	
	return manager.GetOrCreateRule(className, func() css.Rule {
		var size css.Value
		if fontSize, exists := manager.Theme().FontSizes[sizeKey]; exists {
			size = fontSize
		} else {
			// Fallback to treating the key as a direct length value
			size = css.Raw(sizeKey)
		}
		
		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set(css.FontSize, size),
		)
	})
}

// FontWeight creates a font weight utility class.
func FontWeight(manager *UtilityManager, weight string) css.Rule {
	className := ClassName("font", weight)
	
	return manager.GetOrCreateRule(className, func() css.Rule {
		var weightValue css.Value
		
		switch weight {
		case "thin":
			weightValue = css.Keyword("100")
		case "extralight":
			weightValue = css.Keyword("200")
		case "light":
			weightValue = css.Keyword("300")
		case "normal":
			weightValue = css.Keyword("400")
		case "medium":
			weightValue = css.Keyword("500")
		case "semibold":
			weightValue = css.Keyword("600")
		case "bold":
			weightValue = css.Keyword("700")
		case "extrabold":
			weightValue = css.Keyword("800")
		case "black":
			weightValue = css.Keyword("900")
		default:
			// Fallback to treating the weight as a direct value
			weightValue = css.Raw(weight)
		}
		
		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set("font-weight", weightValue),
		)
	})
}

// TextAlign creates a text alignment utility class.
func TextAlign(manager *UtilityManager, alignment string) css.Rule {
	className := ClassName("text", alignment)
	
	return manager.GetOrCreateRule(className, func() css.Rule {
		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set(css.TextAlign, css.Keyword(alignment)),
		)
	})
}

// FontFamily creates a font family utility class.
func FontFamily(manager *UtilityManager, family string) css.Rule {
	className := ClassName("font", family)
	
	return manager.GetOrCreateRule(className, func() css.Rule {
		var fontFamily css.Value
		
		switch family {
		case "sans":
			fontFamily = css.Raw(`ui-sans-serif, system-ui, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji"`)
		case "serif":
			fontFamily = css.Raw(`ui-serif, Georgia, Cambria, "Times New Roman", Times, serif`)
		case "mono":
			fontFamily = css.Raw(`ui-monospace, SFMono-Regular, "SF Mono", Consolas, "Liberation Mono", Menlo, monospace`)
		default:
			// Fallback to treating the family as a direct value
			fontFamily = css.Raw(family)
		}
		
		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set(css.FontFamily, fontFamily),
		)
	})
}

// LineHeight creates a line height utility class.
func LineHeight(manager *UtilityManager, height string) css.Rule {
	className := ClassName("leading", height)
	
	return manager.GetOrCreateRule(className, func() css.Rule {
		var lineHeight css.Value
		
		switch height {
		case "none":
			lineHeight = css.Keyword("1")
		case "tight":
			lineHeight = css.Keyword("1.25")
		case "snug":
			lineHeight = css.Keyword("1.375")
		case "normal":
			lineHeight = css.Keyword("1.5")
		case "relaxed":
			lineHeight = css.Keyword("1.625")
		case "loose":
			lineHeight = css.Keyword("2")
		default:
			// Check if it's a numeric value from theme spacing
			if spacing, exists := manager.Theme().Spacing[height]; exists {
				lineHeight = spacing
			} else {
				// Fallback to treating as direct value
				lineHeight = css.Raw(height)
			}
		}
		
		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set("line-height", lineHeight),
		)
	})
}

// LetterSpacing creates a letter spacing utility class.
func LetterSpacing(manager *UtilityManager, spacing string) css.Rule {
	className := ClassName("tracking", spacing)
	
	return manager.GetOrCreateRule(className, func() css.Rule {
		var letterSpacing css.Value
		
		switch spacing {
		case "tighter":
			letterSpacing = css.Em(-0.05)
		case "tight":
			letterSpacing = css.Em(-0.025)
		case "normal":
			letterSpacing = css.Em(0)
		case "wide":
			letterSpacing = css.Em(0.025)
		case "wider":
			letterSpacing = css.Em(0.05)
		case "widest":
			letterSpacing = css.Em(0.1)
		default:
			// Fallback to treating the spacing as a direct value
			letterSpacing = css.Raw(spacing)
		}
		
		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set("letter-spacing", letterSpacing),
		)
	})
}

// TextTransform creates a text transform utility class.
func TextTransform(manager *UtilityManager, transform string) css.Rule {
	className := transform
	
	return manager.GetOrCreateRule(className, func() css.Rule {
		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set("text-transform", css.Keyword(transform)),
		)
	})
}

// TextDecoration creates a text decoration utility class.
func TextDecoration(manager *UtilityManager, decoration string) css.Rule {
	className := decoration
	
	return manager.GetOrCreateRule(className, func() css.Rule {
		var decorationValue css.Value
		
		switch decoration {
		case "underline":
			decorationValue = css.Keyword("underline")
		case "overline":
			decorationValue = css.Keyword("overline")
		case "line-through":
			decorationValue = css.Keyword("line-through")
		case "no-underline":
			decorationValue = css.Keyword("none")
		default:
			decorationValue = css.Keyword(decoration)
		}
		
		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set("text-decoration", decorationValue),
		)
	})
}

// Convenience functions that use the default manager

// TextXs creates a text-xs utility using the default manager.
func TextXs() css.Rule { return TextSize(defaultManager, "xs") }

// TextSm creates a text-sm utility using the default manager.
func TextSm() css.Rule { return TextSize(defaultManager, "sm") }

// TextBase creates a text-base utility using the default manager.
func TextBase() css.Rule { return TextSize(defaultManager, "base") }

// TextLg creates a text-lg utility using the default manager.
func TextLg() css.Rule { return TextSize(defaultManager, "lg") }

// TextXl creates a text-xl utility using the default manager.
func TextXl() css.Rule { return TextSize(defaultManager, "xl") }

// Text2xl creates a text-2xl utility using the default manager.
func Text2xl() css.Rule { return TextSize(defaultManager, "2xl") }

// Text3xl creates a text-3xl utility using the default manager.
func Text3xl() css.Rule { return TextSize(defaultManager, "3xl") }

// Font weight utilities
func FontThin() css.Rule { return FontWeight(defaultManager, "thin") }
func FontLight() css.Rule { return FontWeight(defaultManager, "light") }
func FontNormal() css.Rule { return FontWeight(defaultManager, "normal") }
func FontMedium() css.Rule { return FontWeight(defaultManager, "medium") }
func FontSemibold() css.Rule { return FontWeight(defaultManager, "semibold") }
func FontBold() css.Rule { return FontWeight(defaultManager, "bold") }

// Text alignment utilities
func TextLeft() css.Rule { return TextAlign(defaultManager, "left") }
func TextCenter() css.Rule { return TextAlign(defaultManager, "center") }
func TextRight() css.Rule { return TextAlign(defaultManager, "right") }
func TextJustify() css.Rule { return TextAlign(defaultManager, "justify") }

// Font family utilities
func FontSans() css.Rule { return FontFamily(defaultManager, "sans") }
func FontSerif() css.Rule { return FontFamily(defaultManager, "serif") }
func FontMono() css.Rule { return FontFamily(defaultManager, "mono") }

// Text transform utilities
func Uppercase() css.Rule { return TextTransform(defaultManager, "uppercase") }
func Lowercase() css.Rule { return TextTransform(defaultManager, "lowercase") }
func Capitalize() css.Rule { return TextTransform(defaultManager, "capitalize") }
func NormalCase() css.Rule { return TextTransform(defaultManager, "none") }

// Text decoration utilities
func Underline() css.Rule { return TextDecoration(defaultManager, "underline") }
func LineThrough() css.Rule { return TextDecoration(defaultManager, "line-through") }
func NoUnderline() css.Rule { return TextDecoration(defaultManager, "no-underline") }