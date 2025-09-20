package tailwind

import (
	"fmt"

	"github.com/ahmed-com/typesafe-css/css"
	"github.com/ahmed-com/typesafe-css/cssgen"
)

func textAlignDecl(alignment string) css.Decl {
	switch alignment {
	case "center":
		return cssgen.SetTextAlign(cssgen.TextAlignValCenter)
	case "left":
		return cssgen.SetTextAlign(cssgen.TextAlignValLeft)
	case "right":
		return cssgen.SetTextAlign(cssgen.TextAlignValRight)
	case "justify":
		return cssgen.SetTextAlign(cssgen.TextAlignValJustify)
	case "start":
		return cssgen.SetTextAlign(cssgen.TextAlignValStart)
	case "end":
		return cssgen.SetTextAlign(cssgen.TextAlignValEnd)
	case "match-parent":
		return cssgen.SetTextAlign(cssgen.TextAlignValMatchParent)
	default:
		return css.Set(cssgen.TextAlign, css.Keyword(alignment))
	}
}

func textTransformDecl(transform string) css.Decl {
	switch transform {
	case "uppercase":
		return cssgen.SetTextTransform(cssgen.TextTransformValUppercase)
	case "lowercase":
		return cssgen.SetTextTransform(cssgen.TextTransformValLowercase)
	case "capitalize":
		return cssgen.SetTextTransform(cssgen.TextTransformValCapitalize)
	case "none":
		return cssgen.SetTextTransform(cssgen.TextTransformValNone)
	default:
		return css.Set(cssgen.TextTransform, css.Keyword(transform))
	}
}

func textDecorationDecl(decoration string) css.Decl {
	switch decoration {
	case "underline":
		return cssgen.SetTextDecorationLine(cssgen.TextDecorationLineValUnderline)
	case "overline":
		return cssgen.SetTextDecorationLine(cssgen.TextDecorationLineValOverline)
	case "line-through":
		return cssgen.SetTextDecorationLine(cssgen.TextDecorationLineValLineThrough)
	case "no-underline":
		return cssgen.SetTextDecorationLine(cssgen.TextDecorationLineValNone)
	default:
		return css.Set(cssgen.TextDecorationLine, css.Keyword(decoration))
	}
}

// TextSize creates a text size utility class.
// Example: TextSize(manager, "lg") generates ".text-lg { font-size: 1.125rem; }"
func TextSize(manager *UtilityManager, sizeKey string) css.Rule {
	className := ClassName("text", sizeKey)

	return manager.GetOrCreateRule(className, func() css.Rule {
		var size css.Value
		if token, exists := manager.Theme().FontSizeToken(sizeKey); exists {
			size = token.Value()
		} else {
			// Fallback to treating the key as a direct length value
			size = css.Raw(sizeKey)
		}

		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set(cssgen.FontSize, size),
		)
	})
}

// FontWeight creates a font weight utility class.
func FontWeight(manager *UtilityManager, weight string) css.Rule {
	className := ClassName("font", weight)

	return manager.GetOrCreateRule(className, func() css.Rule {
		var weightValue css.Value

		if token, exists := manager.Theme().FontWeightToken(weight); exists {
			weightValue = token.Value()
		} else {
			// Fallback to treating the weight as a direct value
			weightValue = css.Raw(weight)
		}

		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set(cssgen.FontWeight, weightValue),
		)
	})
}

// TextAlign creates a text alignment utility class.
func TextAlign(manager *UtilityManager, alignment string) css.Rule {
	className := ClassName("text", alignment)

	return manager.GetOrCreateRule(className, func() css.Rule {
		return css.RuleSet(fmt.Sprintf(".%s", className),
			textAlignDecl(alignment),
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
			css.Set(cssgen.FontFamily, fontFamily),
		)
	})
}

// LineHeight creates a line height utility class.
func LineHeight(manager *UtilityManager, height string) css.Rule {
	className := ClassName("leading", height)

	return manager.GetOrCreateRule(className, func() css.Rule {
		var lineHeight css.Value

		// Try to get from theme first
		if token, exists := manager.Theme().LineHeightToken(height); exists {
			lineHeight = token.Value()
		} else {
			// Legacy fallback values
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
				if spacing, exists := manager.Theme().SpacingToken(height); exists {
					lineHeight = spacing.Value()
				} else {
					// Fallback to treating as direct value
					lineHeight = css.Raw(height)
				}
			}
		}

		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set(cssgen.LineHeight, lineHeight),
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
			css.Set(cssgen.LetterSpacing, letterSpacing),
		)
	})
}

// TextTransform creates a text transform utility class.
func TextTransform(manager *UtilityManager, transform string) css.Rule {
	className := transform

	return manager.GetOrCreateRule(className, func() css.Rule {
		return css.RuleSet(fmt.Sprintf(".%s", className),
			textTransformDecl(transform),
		)
	})
}

// TextDecoration creates a text decoration utility class.
func TextDecoration(manager *UtilityManager, decoration string) css.Rule {
	className := decoration

	return manager.GetOrCreateRule(className, func() css.Rule {
		return css.RuleSet(fmt.Sprintf(".%s", className),
			textDecorationDecl(decoration),
		)
	})
}

// Legacy convenience functions for backward compatibility - use manager-based functions for custom themes
