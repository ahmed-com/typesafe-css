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
//			tailwind.Bg("blue-500"),     // .bg-blue-500 { background-color: #3b82f6; }
//			tailwind.Text("white"),      // .text-white { color: #ffffff; }
//			tailwind.P("4"),             // .p-4 { padding: 1rem; }
//			tailwind.DisplayFlex(),      // .flex { display: flex; }
//			tailwind.JustifyCenterClass(), // .justify-center { justify-content: center; }
//		)
//		
//		fmt.Println(stylesheet.String())
//	}
//
// The package provides several categories of utilities:
//
// Colors: Text, Bg, Border functions for color utilities
// Spacing: P, Px, Py, M, Mx, My, W, H functions for spacing utilities
// Layout: DisplayFlex, DisplayBlock, JustifyCenter, etc. for layout utilities
// Typography: TextSm, FontBold, TextCenter, etc. for typography utilities
//
// All utilities support deduplication - calling the same utility function multiple
// times will only generate the CSS rule once.
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
		DisplayBlock(),
		DisplayInline(),
		DisplayInlineBlock(),
		DisplayFlex(),
		DisplayGrid(),
		DisplayHidden(),
	)
	
	// Add common flexbox utilities
	stylesheet.Add(
		FlexRowClass(),
		FlexColClass(),
		JustifyStartClass(),
		JustifyEndClass(),
		JustifyCenterClass(),
		JustifyBetweenClass(),
		ItemsStartClass(),
		ItemsEndClass(),
		ItemsCenterClass(),
		ItemsStretchClass(),
	)
	
	// Add common position utilities
	stylesheet.Add(
		StaticClass(),
		RelativeClass(),
		AbsoluteClass(),
		FixedClass(),
		StickyClass(),
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
	spacingSizes := []string{"0", "1", "2", "3", "4", "5", "6", "8", "10", "12", "16", "20", "24"}
	for _, size := range spacingSizes {
		stylesheet.Add(
			P(size),
			Px(size),
			Py(size),
			M(size),
			Mx(size),
			My(size),
		)
	}
	
	// Add some common color utilities (neutral colors)
	colorKeys := []string{"black", "white", "gray-100", "gray-200", "gray-300", "gray-500", "gray-700", "gray-900"}
	for _, color := range colorKeys {
		stylesheet.Add(
			Text(color),
			Bg(color),
		)
	}
	
	// Add some common primary/accent colors
	primaryColors := []string{"blue-500", "blue-600", "blue-700", "red-500", "green-500"}
	for _, color := range primaryColors {
		stylesheet.Add(
			Text(color),
			Bg(color),
		)
	}
	
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
	SetDefaultTheme(DefaultTheme())
}