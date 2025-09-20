// Package tailwind provides Tailwind CSS-style utility classes for type-safe CSS generation.
package tailwind

import (
	"github.com/ahmed-com/typesafe-css/css"
)

// Theme defines the design tokens used for generating utility classes.
type Theme struct {
	Colors  map[string]css.Color
	Spacing map[string]css.Length
	FontSizes map[string]css.Length
	Screens map[string]string
}

// DefaultTheme provides a Tailwind-inspired default theme configuration.
func DefaultTheme() *Theme {
	return &Theme{
		Colors: map[string]css.Color{
			// Grayscale
			"black":     css.Hex("#000000"),
			"white":     css.Hex("#ffffff"),
			"gray-50":   css.Hex("#f9fafb"),
			"gray-100":  css.Hex("#f3f4f6"),
			"gray-200":  css.Hex("#e5e7eb"),
			"gray-300":  css.Hex("#d1d5db"),
			"gray-400":  css.Hex("#9ca3af"),
			"gray-500":  css.Hex("#6b7280"),
			"gray-600":  css.Hex("#4b5563"),
			"gray-700":  css.Hex("#374151"),
			"gray-800":  css.Hex("#1f2937"),
			"gray-900":  css.Hex("#111827"),
			"gray-950":  css.Hex("#030712"),
			
			// Blue
			"blue-50":   css.Hex("#eff6ff"),
			"blue-100":  css.Hex("#dbeafe"),
			"blue-200":  css.Hex("#bfdbfe"),
			"blue-300":  css.Hex("#93c5fd"),
			"blue-400":  css.Hex("#60a5fa"),
			"blue-500":  css.Hex("#3b82f6"),
			"blue-600":  css.Hex("#2563eb"),
			"blue-700":  css.Hex("#1d4ed8"),
			"blue-800":  css.Hex("#1e40af"),
			"blue-900":  css.Hex("#1e3a8a"),
			"blue-950":  css.Hex("#172554"),
			
			// Red
			"red-50":    css.Hex("#fef2f2"),
			"red-100":   css.Hex("#fee2e2"),
			"red-200":   css.Hex("#fecaca"),
			"red-300":   css.Hex("#fca5a5"),
			"red-400":   css.Hex("#f87171"),
			"red-500":   css.Hex("#ef4444"),
			"red-600":   css.Hex("#dc2626"),
			"red-700":   css.Hex("#b91c1c"),
			"red-800":   css.Hex("#991b1b"),
			"red-900":   css.Hex("#7f1d1d"),
			"red-950":   css.Hex("#450a0a"),
			
			// Green
			"green-50":  css.Hex("#f0fdf4"),
			"green-100": css.Hex("#dcfce7"),
			"green-200": css.Hex("#bbf7d0"),
			"green-300": css.Hex("#86efac"),
			"green-400": css.Hex("#4ade80"),
			"green-500": css.Hex("#22c55e"),
			"green-600": css.Hex("#16a34a"),
			"green-700": css.Hex("#15803d"),
			"green-800": css.Hex("#166534"),
			"green-900": css.Hex("#14532d"),
			"green-950": css.Hex("#052e16"),
		},
		
		Spacing: map[string]css.Length{
			"0":    css.Px(0),
			"px":   css.Px(1),
			"0.5":  css.Rem(0.125),  // 2px
			"1":    css.Rem(0.25),   // 4px
			"1.5":  css.Rem(0.375),  // 6px
			"2":    css.Rem(0.5),    // 8px
			"2.5":  css.Rem(0.625),  // 10px
			"3":    css.Rem(0.75),   // 12px
			"3.5":  css.Rem(0.875),  // 14px
			"4":    css.Rem(1),      // 16px
			"5":    css.Rem(1.25),   // 20px
			"6":    css.Rem(1.5),    // 24px
			"7":    css.Rem(1.75),   // 28px
			"8":    css.Rem(2),      // 32px
			"9":    css.Rem(2.25),   // 36px
			"10":   css.Rem(2.5),    // 40px
			"11":   css.Rem(2.75),   // 44px
			"12":   css.Rem(3),      // 48px
			"14":   css.Rem(3.5),    // 56px
			"16":   css.Rem(4),      // 64px
			"20":   css.Rem(5),      // 80px
			"24":   css.Rem(6),      // 96px
			"28":   css.Rem(7),      // 112px
			"32":   css.Rem(8),      // 128px
			"36":   css.Rem(9),      // 144px
			"40":   css.Rem(10),     // 160px
			"44":   css.Rem(11),     // 176px
			"48":   css.Rem(12),     // 192px
			"52":   css.Rem(13),     // 208px
			"56":   css.Rem(14),     // 224px
			"60":   css.Rem(15),     // 240px
			"64":   css.Rem(16),     // 256px
			"72":   css.Rem(18),     // 288px
			"80":   css.Rem(20),     // 320px
			"96":   css.Rem(24),     // 384px
		},
		
		FontSizes: map[string]css.Length{
			"xs":   css.Rem(0.75),   // 12px
			"sm":   css.Rem(0.875),  // 14px
			"base": css.Rem(1),      // 16px
			"lg":   css.Rem(1.125),  // 18px
			"xl":   css.Rem(1.25),   // 20px
			"2xl":  css.Rem(1.5),    // 24px
			"3xl":  css.Rem(1.875),  // 30px
			"4xl":  css.Rem(2.25),   // 36px
			"5xl":  css.Rem(3),      // 48px
			"6xl":  css.Rem(3.75),   // 60px
			"7xl":  css.Rem(4.5),    // 72px
			"8xl":  css.Rem(6),      // 96px
			"9xl":  css.Rem(8),      // 128px
		},
		
		Screens: map[string]string{
			"sm":  "(min-width: 640px)",
			"md":  "(min-width: 768px)",
			"lg":  "(min-width: 1024px)",
			"xl":  "(min-width: 1280px)",
			"2xl": "(min-width: 1536px)",
		},
	}
}

// CustomTheme creates a new theme by merging custom values with the default theme.
func CustomTheme(overrides *Theme) *Theme {
	base := DefaultTheme()
	
	if overrides == nil {
		return base
	}
	
	// Merge colors
	if overrides.Colors != nil {
		for k, v := range overrides.Colors {
			base.Colors[k] = v
		}
	}
	
	// Merge spacing
	if overrides.Spacing != nil {
		for k, v := range overrides.Spacing {
			base.Spacing[k] = v
		}
	}
	
	// Merge font sizes
	if overrides.FontSizes != nil {
		for k, v := range overrides.FontSizes {
			base.FontSizes[k] = v
		}
	}
	
	// Merge screens
	if overrides.Screens != nil {
		for k, v := range overrides.Screens {
			base.Screens[k] = v
		}
	}
	
	return base
}