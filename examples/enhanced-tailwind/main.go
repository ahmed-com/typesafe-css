package main

import (
	"fmt"
	"reflect"

	"github.com/ahmed-com/typesafe-css/tailwind"
)

func main() {
	fmt.Println("=== Enhanced Tailwind CSS Configuration Test ===")
	fmt.Println()

	// Test typed configuration structure
	theme := tailwind.DefaultTheme()
	fmt.Println("✅ Typed Configuration Structure:")
	fmt.Printf("  Theme has %d main configuration categories\n", reflect.TypeOf(theme).NumField())

	// Sample some colors from the typed config
	fmt.Println()
	fmt.Println("✅ Color Configuration Examples:")
	fmt.Printf("  Red 500: %s\n", theme.Colors.Red500.String())
	fmt.Printf("  Blue 500: %s\n", theme.Colors.Blue500.String())
	fmt.Printf("  Green 500: %s\n", theme.Colors.Green500.String())
	fmt.Printf("  Transparent: %s\n", theme.Colors.Transparent.String())

	// Sample spacing values
	fmt.Println()
	fmt.Println("✅ Spacing Configuration Examples:")
	fmt.Printf("  Size 4: %s\n", theme.Spacing.Size4.String())
	fmt.Printf("  Size 8: %s\n", theme.Spacing.Size8.String())
	fmt.Printf("  Size 16: %s\n", theme.Spacing.Size16.String())

	// Sample font sizes - showing the Size component
	fmt.Println()
	fmt.Println("✅ Typography Configuration Examples:")
	fmt.Printf("  Text SM size: %s\n", theme.FontSize.Sm.Size.String())
	fmt.Printf("  Text LG size: %s\n", theme.FontSize.Lg.Size.String())
	fmt.Printf("  Text XL size: %s\n", theme.FontSize.Xl.Size.String())

	fmt.Println()
	fmt.Println("✅ Generated CSS from Utility System:")

	// Test utility generation with the typed config
	generator := tailwind.NewUtilityGenerator(tailwind.DefaultConfig())
	generatedStyles := generator.GenerateUtilities()

	fmt.Printf("✅ Complete utility system generated %d total utility rules\n", len(generatedStyles.Items))

	fmt.Println()
	fmt.Println("✅ Type Safety Demonstration:")

	// Show that we now have compile-time safety
	config := tailwind.DefaultConfig()

	// This is now type-safe - no string lookups!
	redColor := config.Theme.Colors.Red500
	mediumSpacing := config.Theme.Spacing.Size4

	fmt.Printf("  Red color value: %v\n", redColor.ToCSSValue())
	fmt.Printf("  Medium spacing value: %v\n", mediumSpacing.ToCSSValue())

	// Show some generated utilities
	fmt.Println()
	fmt.Println("✅ Sample Generated Utility Rules:")
	for i, item := range generatedStyles.Items[:min(5, len(generatedStyles.Items))] {
		fmt.Printf("  %d. %s\n", i+1, item.String())
	}
	if len(generatedStyles.Items) > 5 {
		fmt.Printf("  ... and %d more\n", len(generatedStyles.Items)-5)
	}

	fmt.Println()
	fmt.Println("🎉 Enhanced type-safe Tailwind CSS system is working!")
}

// min function for compatibility
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
