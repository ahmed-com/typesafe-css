package main

import (
	"fmt"
	
	"github.com/ahmed-com/typesafe-css/css"
	"github.com/ahmed-com/typesafe-css/tailwind"
)

func main() {
	fmt.Println("=== Enhanced Tailwind CSS Configuration Test ===")
	fmt.Println()

	// Test complete color palette
	fmt.Println("Colors supported:")
	theme := tailwind.DefaultTheme()
	
	colorSamples := []string{
		"slate-500", "gray-500", "zinc-500", "neutral-500", "stone-500",
		"red-500", "orange-500", "amber-500", "yellow-500", "lime-500",
		"green-500", "emerald-500", "teal-500", "cyan-500", "sky-500",
		"blue-500", "indigo-500", "violet-500", "purple-500", "fuchsia-500",
		"pink-500", "rose-500",
	}
	
	for _, color := range colorSamples {
		if _, exists := theme.Colors[color]; exists {
			fmt.Printf("  ✓ %s\n", color)
		} else {
			fmt.Printf("  ✗ %s\n", color)
		}
	}
	
	fmt.Println()
	fmt.Println("New utility categories:")
	
	// Test new utilities
	var stylesheet css.Stylesheet
	
	// Border radius utilities
	stylesheet.Add(
		tailwind.RoundedSm(),
		tailwind.RoundedLg(),
		tailwind.RoundedFull(),
	)
	
	// Shadow utilities  
	stylesheet.Add(
		tailwind.ShadowSm(),
		tailwind.ShadowLg(),
		tailwind.Shadow2xl(),
	)
	
	// Opacity utilities
	stylesheet.Add(
		tailwind.Opacity50(),
		tailwind.Opacity75(),
		tailwind.Opacity100(),
	)
	
	// Z-index utilities
	stylesheet.Add(
		tailwind.Z10(),
		tailwind.Z50(),
		tailwind.ZAuto(),
	)
	
	// Filter utilities
	stylesheet.Add(
		tailwind.BlurClass("sm"),
		tailwind.BrightnessClass("110"),
		tailwind.GrayscaleClass(""),
	)
	
	// Custom border radius
	stylesheet.Add(
		tailwind.Rounded("md"),
		tailwind.Rounded("xl"),
	)
	
	// Custom shadow
	stylesheet.Add(
		tailwind.Shadow("md"),
		tailwind.Shadow("none"),
	)
	
	fmt.Println("Generated CSS:")
	fmt.Println(stylesheet.String())
	
	fmt.Println()
	fmt.Println("Theme configuration counts:")
	fmt.Printf("  Colors: %d\n", len(theme.Colors))
	fmt.Printf("  Spacing: %d\n", len(theme.Spacing))
	fmt.Printf("  Font sizes: %d\n", len(theme.FontSizes))
	fmt.Printf("  Border radius: %d\n", len(theme.BorderRadius))
	fmt.Printf("  Border width: %d\n", len(theme.BorderWidth))
	fmt.Printf("  Box shadow: %d\n", len(theme.BoxShadow))
	fmt.Printf("  Opacity: %d\n", len(theme.Opacity))
	fmt.Printf("  Font weight: %d\n", len(theme.FontWeight))
	fmt.Printf("  Line height: %d\n", len(theme.LineHeight))
	fmt.Printf("  Z-index: %d\n", len(theme.ZIndex))
	fmt.Printf("  Blur: %d\n", len(theme.Blur))
	fmt.Printf("  Brightness: %d\n", len(theme.Brightness))
	fmt.Printf("  Contrast: %d\n", len(theme.Contrast))
	fmt.Printf("  Grayscale: %d\n", len(theme.Grayscale))
	fmt.Printf("  Invert: %d\n", len(theme.Invert))
	fmt.Printf("  Saturate: %d\n", len(theme.Saturate))
	fmt.Printf("  Sepia: %d\n", len(theme.Sepia))
}