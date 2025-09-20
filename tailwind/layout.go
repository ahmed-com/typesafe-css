package tailwind

import (
	"github.com/ahmed-com/typesafe-css/css"
)

// Legacy API - these functions are deprecated in favor of the direct functions in generated.go
// They are kept for backward compatibility with custom managers

// Legacy display utilities
func BlockWithManager(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("block", func() css.Rule {
		return css.RuleSet(".block", css.Set(css.Display, css.DisplayBlock))
	})
}

func FlexWithManager(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("flex", func() css.Rule {
		return css.RuleSet(".flex", css.Set(css.Display, css.DisplayFlex))
	})
}

// Convenience functions for backward compatibility
func DisplayBlock() css.Rule { return Block() }
func DisplayFlex() css.Rule { return Flex() }
func FlexRowClass() css.Rule { return FlexRow() }
func FlexColClass() css.Rule { return FlexCol() }
func JustifyStartClass() css.Rule { return JustifyStart() }
func JustifyEndClass() css.Rule { return JustifyEnd() }
func JustifyCenterClass() css.Rule { return JustifyCenter() }
func JustifyBetweenClass() css.Rule { return JustifyBetween() }
func JustifyAroundClass() css.Rule { return JustifyAround() }
func JustifyEvenlyClass() css.Rule { return JustifyEvenly() }
func ItemsStartClass() css.Rule { return ItemsStart() }
func ItemsEndClass() css.Rule { return ItemsEnd() }
func ItemsCenterClass() css.Rule { return ItemsCenter() }
func ItemsBaselineClass() css.Rule { return ItemsBaseline() }
func ItemsStretchClass() css.Rule { return ItemsStretch() }
func StaticClass() css.Rule { return Static() }
func FixedClass() css.Rule { return Fixed() }
func AbsoluteClass() css.Rule { return Absolute() }
func RelativeClass() css.Rule { return Relative() }
func StickyClass() css.Rule { return Sticky() }