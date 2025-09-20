// Package tailwind provides Tailwind CSS-style utility classes for type-safe CSS generation.
//
// This package builds on top of the core css package to provide a Tailwind-inspired
// utility-first CSS authoring experience with full type safety.
//
// Example usage:
//
//	import (
//		"github.com/ahmed-com/typesafe-css/css"
//		"github.com/ahmed-com/typesafe-css/tailwind"
//	)
//
//	func main() {
//		var stylesheet css.Stylesheet
//
//		// Add utility classes to stylesheet
//		stylesheet.Add(
//			tailwind.BgBlue500(),        // .bg-blue-500 { background-color: #3b82f6; }
//			tailwind.TextWhite(),        // .text-white { color: #ffffff; }
//			tailwind.P4(),               // .p-4 { padding: 1rem; }
//			tailwind.Flex(),             // .flex { display: flex; }
//			tailwind.JustifyCenter(),    // .justify-center { justify-content: center; }
//			tailwind.Bg("#bada55"),      // .bg-[#bada55] { background-color: #bada55; }
//		)
//
//		fmt.Println(stylesheet.String())
//	}
//
// The package provides several categories of utilities:
//
// Colors: BgBlue500, TextWhite, etc. for specific color utilities
// Spacing: P4, Px2, Py4, M8, etc. for specific spacing utilities
// Layout: Flex, Block, JustifyCenter, etc. for layout utilities
// Typography: TextLg, FontBold, TextCenter, etc. for typography utilities
//
// All utilities support deduplication - calling the same utility function multiple
// times will only generate the CSS rule once.
//
// For arbitrary values, use the Bg(), Text(), P(), etc. functions with custom values:
//
//	tailwind.Bg("#custom")     - Custom background color
//	tailwind.Text("rgb(255,0,0)") - Custom text color
package tailwind

import (
	"github.com/ahmed-com/typesafe-css/css"
)

// GenerateUtilityStylesheet creates a stylesheet with commonly used utility classes.
// This is a convenience function for generating a basic set of Tailwind-style utilities.
func GenerateUtilityStylesheet() css.Stylesheet {
	var stylesheet css.Stylesheet

	// Add common display utilities
	stylesheet.Add(
		Block(),
		Inline(),
		InlineBlock(),
		Flex(),
		Grid(),
		Hidden(),
	)

	// Add common flexbox utilities
	stylesheet.Add(
		FlexRow(),
		FlexCol(),
		JustifyStart(),
		JustifyEnd(),
		JustifyCenter(),
		JustifyBetween(),
		ItemsStart(),
		ItemsEnd(),
		ItemsCenter(),
		ItemsStretch(),
	)

	// Add common position utilities
	stylesheet.Add(
		Static(),
		Relative(),
		Absolute(),
		Fixed(),
		Sticky(),
	)

	// Add common text size utilities
	stylesheet.Add(
		TextXs(),
		TextSm(),
		TextBase(),
		TextLg(),
		TextXl(),
		Text2xl(),
		Text3xl(),
	)

	// Add common font weight utilities
	stylesheet.Add(
		FontThin(),
		FontLight(),
		FontNormal(),
		FontMedium(),
		FontSemibold(),
		FontBold(),
	)

	// Add common text alignment utilities
	stylesheet.Add(
		TextLeft(),
		TextCenter(),
		TextRight(),
	)

	// Add some common spacing utilities
	stylesheet.Add(
		P0(), P1(), P2(), P3(), P4(), P5(), P6(), P8(), P10(), P12(), P16(), P20(), P24(),
		Px0(), Px1(), Px2(), Px3(), Px4(), Px5(), Px6(), Px8(), Px10(), Px12(), Px16(), Px20(), Px24(),
		Py0(), Py1(), Py2(), Py3(), Py4(), Py5(), Py6(), Py8(), Py10(), Py12(), Py16(), Py20(), Py24(),
		M0(), M1(), M2(), M3(), M4(), M5(), M6(), M8(), M10(), M12(), M16(), M20(), M24(),
		Mx0(), Mx1(), Mx2(), Mx3(), Mx4(), Mx5(), Mx6(), Mx8(), Mx10(), Mx12(), Mx16(), Mx20(), Mx24(), MxAuto(),
		My0(), My1(), My2(), My3(), My4(), My5(), My6(), My8(), My10(), My12(), My16(), My20(), My24(),
	)

	// Add some common color utilities (neutral colors)
	stylesheet.Add(
		TextBlack(), TextWhite(), TextGray100(), TextGray200(), TextGray300(), TextGray500(), TextGray700(), TextGray900(),
		BgBlack(), BgWhite(), BgGray100(), BgGray200(), BgGray300(), BgGray500(), BgGray700(), BgGray900(),
	)

	// Add some common primary/accent colors
	stylesheet.Add(
		TextBlue500(), TextBlue600(), TextBlue700(), TextRed500(), TextGreen500(),
		BgBlue500(), BgBlue600(), BgBlue700(), BgRed500(), BgGreen500(),
	)

	return stylesheet
}

// AddUtilitiesToStylesheet is a helper function to add utility rules to an existing stylesheet.
// This ensures utilities are only added once even if called multiple times.
func AddUtilitiesToStylesheet(stylesheet *css.Stylesheet, utilities ...css.Rule) {
	for _, utility := range utilities {
		stylesheet.Add(utility)
	}
}

// WithCustomTheme returns a new utility manager configured with a custom theme.
// This is useful when you want to use utilities with a different theme than the default.
func WithCustomTheme(theme *Theme) *UtilityManager {
	return NewUtilityManager(theme)
}

// ResetDefaultTheme resets the default theme to the built-in Tailwind-inspired theme.
func ResetDefaultTheme() {
	// TODO: Update to use new typed config system
	// SetDefaultTheme(DefaultTheme())
}
