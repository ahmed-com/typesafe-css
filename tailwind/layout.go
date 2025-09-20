package tailwind

import (
	"github.com/ahmed-com/typesafe-css/css"
)

// Layout utilities for display properties

// Block creates a "block" display utility.
func Block(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("block", func() css.Rule {
		return css.RuleSet(".block",
			css.Set(css.Display, css.DisplayBlock),
		)
	})
}

// Inline creates an "inline" display utility.
func Inline(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("inline", func() css.Rule {
		return css.RuleSet(".inline",
			css.Set(css.Display, css.DisplayInline),
		)
	})
}

// InlineBlock creates an "inline-block" display utility.
func InlineBlock(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("inline-block", func() css.Rule {
		return css.RuleSet(".inline-block",
			css.Set(css.Display, css.Keyword("inline-block")),
		)
	})
}

// Flex creates a "flex" display utility.
func Flex(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("flex", func() css.Rule {
		return css.RuleSet(".flex",
			css.Set(css.Display, css.DisplayFlex),
		)
	})
}

// InlineFlex creates an "inline-flex" display utility.
func InlineFlex(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("inline-flex", func() css.Rule {
		return css.RuleSet(".inline-flex",
			css.Set(css.Display, css.Keyword("inline-flex")),
		)
	})
}

// Grid creates a "grid" display utility.
func Grid(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("grid", func() css.Rule {
		return css.RuleSet(".grid",
			css.Set(css.Display, css.Keyword("grid")),
		)
	})
}

// InlineGrid creates an "inline-grid" display utility.
func InlineGrid(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("inline-grid", func() css.Rule {
		return css.RuleSet(".inline-grid",
			css.Set(css.Display, css.Keyword("inline-grid")),
		)
	})
}

// Hidden creates a "hidden" display utility.
func Hidden(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("hidden", func() css.Rule {
		return css.RuleSet(".hidden",
			css.Set(css.Display, css.DisplayNone),
		)
	})
}

// Flexbox utilities

// FlexRow creates a "flex-row" utility.
func FlexRow(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("flex-row", func() css.Rule {
		return css.RuleSet(".flex-row",
			css.Set(css.FlexDirection, css.Keyword("row")),
		)
	})
}

// FlexCol creates a "flex-col" utility.
func FlexCol(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("flex-col", func() css.Rule {
		return css.RuleSet(".flex-col",
			css.Set(css.FlexDirection, css.Keyword("column")),
		)
	})
}

// JustifyStart creates a "justify-start" utility.
func JustifyStart(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("justify-start", func() css.Rule {
		return css.RuleSet(".justify-start",
			css.Set(css.JustifyContent, css.Keyword("flex-start")),
		)
	})
}

// JustifyEnd creates a "justify-end" utility.
func JustifyEnd(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("justify-end", func() css.Rule {
		return css.RuleSet(".justify-end",
			css.Set(css.JustifyContent, css.Keyword("flex-end")),
		)
	})
}

// JustifyCenter creates a "justify-center" utility.
func JustifyCenter(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("justify-center", func() css.Rule {
		return css.RuleSet(".justify-center",
			css.Set(css.JustifyContent, css.Keyword("center")),
		)
	})
}

// JustifyBetween creates a "justify-between" utility.
func JustifyBetween(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("justify-between", func() css.Rule {
		return css.RuleSet(".justify-between",
			css.Set(css.JustifyContent, css.Keyword("space-between")),
		)
	})
}

// JustifyAround creates a "justify-around" utility.
func JustifyAround(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("justify-around", func() css.Rule {
		return css.RuleSet(".justify-around",
			css.Set(css.JustifyContent, css.Keyword("space-around")),
		)
	})
}

// JustifyEvenly creates a "justify-evenly" utility.
func JustifyEvenly(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("justify-evenly", func() css.Rule {
		return css.RuleSet(".justify-evenly",
			css.Set(css.JustifyContent, css.Keyword("space-evenly")),
		)
	})
}

// ItemsStart creates an "items-start" utility.
func ItemsStart(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("items-start", func() css.Rule {
		return css.RuleSet(".items-start",
			css.Set(css.AlignItems, css.Keyword("flex-start")),
		)
	})
}

// ItemsEnd creates an "items-end" utility.
func ItemsEnd(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("items-end", func() css.Rule {
		return css.RuleSet(".items-end",
			css.Set(css.AlignItems, css.Keyword("flex-end")),
		)
	})
}

// ItemsCenter creates an "items-center" utility.
func ItemsCenter(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("items-center", func() css.Rule {
		return css.RuleSet(".items-center",
			css.Set(css.AlignItems, css.Keyword("center")),
		)
	})
}

// ItemsBaseline creates an "items-baseline" utility.
func ItemsBaseline(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("items-baseline", func() css.Rule {
		return css.RuleSet(".items-baseline",
			css.Set(css.AlignItems, css.Keyword("baseline")),
		)
	})
}

// ItemsStretch creates an "items-stretch" utility.
func ItemsStretch(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("items-stretch", func() css.Rule {
		return css.RuleSet(".items-stretch",
			css.Set(css.AlignItems, css.Keyword("stretch")),
		)
	})
}

// Position utilities

// Static creates a "static" position utility.
func Static(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("static", func() css.Rule {
		return css.RuleSet(".static",
			css.Set(css.Position, css.PositionStatic),
		)
	})
}

// Fixed creates a "fixed" position utility.
func Fixed(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("fixed", func() css.Rule {
		return css.RuleSet(".fixed",
			css.Set(css.Position, css.PositionFixed),
		)
	})
}

// Absolute creates an "absolute" position utility.
func Absolute(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("absolute", func() css.Rule {
		return css.RuleSet(".absolute",
			css.Set(css.Position, css.PositionAbsolute),
		)
	})
}

// Relative creates a "relative" position utility.
func Relative(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("relative", func() css.Rule {
		return css.RuleSet(".relative",
			css.Set(css.Position, css.PositionRelative),
		)
	})
}

// Sticky creates a "sticky" position utility.
func Sticky(manager *UtilityManager) css.Rule {
	return manager.GetOrCreateRule("sticky", func() css.Rule {
		return css.RuleSet(".sticky",
			css.Set(css.Position, css.PositionSticky),
		)
	})
}

// Convenience functions that use the default manager

// Display utilities
func DisplayBlock() css.Rule { return Block(defaultManager) }
func DisplayInline() css.Rule { return Inline(defaultManager) }
func DisplayInlineBlock() css.Rule { return InlineBlock(defaultManager) }
func DisplayFlex() css.Rule { return Flex(defaultManager) }
func DisplayInlineFlex() css.Rule { return InlineFlex(defaultManager) }
func DisplayGrid() css.Rule { return Grid(defaultManager) }
func DisplayInlineGrid() css.Rule { return InlineGrid(defaultManager) }
func DisplayHidden() css.Rule { return Hidden(defaultManager) }

// Flexbox utilities  
func FlexRowClass() css.Rule { return FlexRow(defaultManager) }
func FlexColClass() css.Rule { return FlexCol(defaultManager) }
func JustifyStartClass() css.Rule { return JustifyStart(defaultManager) }
func JustifyEndClass() css.Rule { return JustifyEnd(defaultManager) }
func JustifyCenterClass() css.Rule { return JustifyCenter(defaultManager) }
func JustifyBetweenClass() css.Rule { return JustifyBetween(defaultManager) }
func JustifyAroundClass() css.Rule { return JustifyAround(defaultManager) }
func JustifyEvenlyClass() css.Rule { return JustifyEvenly(defaultManager) }
func ItemsStartClass() css.Rule { return ItemsStart(defaultManager) }
func ItemsEndClass() css.Rule { return ItemsEnd(defaultManager) }
func ItemsCenterClass() css.Rule { return ItemsCenter(defaultManager) }
func ItemsBaselineClass() css.Rule { return ItemsBaseline(defaultManager) }
func ItemsStretchClass() css.Rule { return ItemsStretch(defaultManager) }

// Position utilities
func StaticClass() css.Rule { return Static(defaultManager) }
func FixedClass() css.Rule { return Fixed(defaultManager) }
func AbsoluteClass() css.Rule { return Absolute(defaultManager) }
func RelativeClass() css.Rule { return Relative(defaultManager) }
func StickyClass() css.Rule { return Sticky(defaultManager) }