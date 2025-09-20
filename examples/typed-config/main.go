package main

import (
	"fmt"

	"github.com/ahmed-com/typesafe-css/tailwind"
)

func main() {
	fmt.Println("🎉 TypeSafe CSS - New Typed Configuration System Demo")
	fmt.Println("==================================================")

	// Create a typed configuration - no more string maps!
	config := tailwind.DefaultConfig()

	fmt.Println("\n📋 Configuration Structure:")
	fmt.Printf("✅ Colors are typed: config.Theme.Colors.Red50 = %s\n", config.Theme.Colors.Red50.String())
	fmt.Printf("✅ Spacing is typed: config.Theme.Spacing.Size4 = %s\n", config.Theme.Spacing.Size4.String())
	fmt.Printf("✅ Font sizes are typed: config.Theme.FontSize.Xl.Size = %s\n", config.Theme.FontSize.Xl.Size.String())

	fmt.Println("\n🔧 Custom Value Creation:")

	// Create custom typed values
	customColor := tailwind.ColorFromHex("brand-primary", "#3b82f6")
	customSpacing := tailwind.LengthFromRem("custom-spacing", 1.5)

	fmt.Printf("✅ Custom color: %s -> %s\n", customColor.Name, customColor.ToCSSValue().String())
	fmt.Printf("✅ Custom spacing: %s -> %s\n", customSpacing.Name, customSpacing.ToCSSValue().String())

	fmt.Println("\n🎯 Type Safety Benefits:")
	fmt.Println("✅ No more string-based map access like theme.Colors[\"red-50\"]")
	fmt.Println("✅ Compile-time safety with config.Theme.Colors.Red50")
	fmt.Println("✅ IDE autocomplete and IntelliSense support")
	fmt.Println("✅ Leverages generated CSS property types from cssgen")
	fmt.Println("✅ Nested named properties instead of flat string keys")

	fmt.Println("\n🚀 The typed configuration system is now ready!")
	fmt.Println("📝 Next steps: Complete utility generator for all config types")
}
