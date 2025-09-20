// Package tailwind provides Tailwind CSS-style utility classes for type-safe CSS generation.
package tailwind

import (
	"github.com/ahmed-com/typesafe-css/css"
)

// Theme defines the design tokens used for generating utility classes.
type Theme struct {
	Colors        map[string]css.Color
	Spacing       map[string]css.Length
	FontSizes     map[string]css.Length
	Screens       map[string]string
	BorderRadius  map[string]css.Length
	BorderWidth   map[string]css.Length
	BoxShadow     map[string]string
	Opacity       map[string]string
	FontWeight    map[string]string
	LineHeight    map[string]string
	ZIndex        map[string]string
	Blur          map[string]string
	Brightness    map[string]string
	Contrast      map[string]string
	Grayscale     map[string]string
	Invert        map[string]string
	Saturate      map[string]string
	Sepia         map[string]string
}

// DefaultTheme provides a Tailwind-inspired default theme configuration.
func DefaultTheme() *Theme {
	return &Theme{
		Colors: map[string]css.Color{
			// Core colors
			"inherit":     css.Color("inherit"),
			"current":     css.Color("current"),
			"transparent": css.Color("transparent"),
			"black":       css.Hex("#000000"),
			"white":       css.Hex("#ffffff"),
			
			// Slate
			"slate-50":  css.Hex("#f8fafc"),
			"slate-100": css.Hex("#f1f5f9"),
			"slate-200": css.Hex("#e2e8f0"),
			"slate-300": css.Hex("#cbd5e1"),
			"slate-400": css.Hex("#94a3b8"),
			"slate-500": css.Hex("#64748b"),
			"slate-600": css.Hex("#475569"),
			"slate-700": css.Hex("#334155"),
			"slate-800": css.Hex("#1e293b"),
			"slate-900": css.Hex("#0f172a"),
			"slate-950": css.Hex("#020617"),
			
			// Gray
			"gray-50":  css.Hex("#f9fafb"),
			"gray-100": css.Hex("#f3f4f6"),
			"gray-200": css.Hex("#e5e7eb"),
			"gray-300": css.Hex("#d1d5db"),
			"gray-400": css.Hex("#9ca3af"),
			"gray-500": css.Hex("#6b7280"),
			"gray-600": css.Hex("#4b5563"),
			"gray-700": css.Hex("#374151"),
			"gray-800": css.Hex("#1f2937"),
			"gray-900": css.Hex("#111827"),
			"gray-950": css.Hex("#030712"),
			
			// Zinc
			"zinc-50":  css.Hex("#fafafa"),
			"zinc-100": css.Hex("#f4f4f5"),
			"zinc-200": css.Hex("#e4e4e7"),
			"zinc-300": css.Hex("#d4d4d8"),
			"zinc-400": css.Hex("#a1a1aa"),
			"zinc-500": css.Hex("#71717a"),
			"zinc-600": css.Hex("#52525b"),
			"zinc-700": css.Hex("#3f3f46"),
			"zinc-800": css.Hex("#27272a"),
			"zinc-900": css.Hex("#18181b"),
			"zinc-950": css.Hex("#09090b"),
			
			// Neutral
			"neutral-50":  css.Hex("#fafafa"),
			"neutral-100": css.Hex("#f5f5f5"),
			"neutral-200": css.Hex("#e5e5e5"),
			"neutral-300": css.Hex("#d4d4d4"),
			"neutral-400": css.Hex("#a3a3a3"),
			"neutral-500": css.Hex("#737373"),
			"neutral-600": css.Hex("#525252"),
			"neutral-700": css.Hex("#404040"),
			"neutral-800": css.Hex("#262626"),
			"neutral-900": css.Hex("#171717"),
			"neutral-950": css.Hex("#0a0a0a"),
			
			// Stone
			"stone-50":  css.Hex("#fafaf9"),
			"stone-100": css.Hex("#f5f5f4"),
			"stone-200": css.Hex("#e7e5e4"),
			"stone-300": css.Hex("#d6d3d1"),
			"stone-400": css.Hex("#a8a29e"),
			"stone-500": css.Hex("#78716c"),
			"stone-600": css.Hex("#57534e"),
			"stone-700": css.Hex("#44403c"),
			"stone-800": css.Hex("#292524"),
			"stone-900": css.Hex("#1c1917"),
			"stone-950": css.Hex("#0c0a09"),
			
			// Red
			"red-50":  css.Hex("#fef2f2"),
			"red-100": css.Hex("#fee2e2"),
			"red-200": css.Hex("#fecaca"),
			"red-300": css.Hex("#fca5a5"),
			"red-400": css.Hex("#f87171"),
			"red-500": css.Hex("#ef4444"),
			"red-600": css.Hex("#dc2626"),
			"red-700": css.Hex("#b91c1c"),
			"red-800": css.Hex("#991b1b"),
			"red-900": css.Hex("#7f1d1d"),
			"red-950": css.Hex("#450a0a"),
			
			// Orange
			"orange-50":  css.Hex("#fff7ed"),
			"orange-100": css.Hex("#ffedd5"),
			"orange-200": css.Hex("#fed7aa"),
			"orange-300": css.Hex("#fdba74"),
			"orange-400": css.Hex("#fb923c"),
			"orange-500": css.Hex("#f97316"),
			"orange-600": css.Hex("#ea580c"),
			"orange-700": css.Hex("#c2410c"),
			"orange-800": css.Hex("#9a3412"),
			"orange-900": css.Hex("#7c2d12"),
			"orange-950": css.Hex("#431407"),
			
			// Amber
			"amber-50":  css.Hex("#fffbeb"),
			"amber-100": css.Hex("#fef3c7"),
			"amber-200": css.Hex("#fde68a"),
			"amber-300": css.Hex("#fcd34d"),
			"amber-400": css.Hex("#fbbf24"),
			"amber-500": css.Hex("#f59e0b"),
			"amber-600": css.Hex("#d97706"),
			"amber-700": css.Hex("#b45309"),
			"amber-800": css.Hex("#92400e"),
			"amber-900": css.Hex("#78350f"),
			"amber-950": css.Hex("#451a03"),
			
			// Yellow
			"yellow-50":  css.Hex("#fefce8"),
			"yellow-100": css.Hex("#fef9c3"),
			"yellow-200": css.Hex("#fef08a"),
			"yellow-300": css.Hex("#fde047"),
			"yellow-400": css.Hex("#facc15"),
			"yellow-500": css.Hex("#eab308"),
			"yellow-600": css.Hex("#ca8a04"),
			"yellow-700": css.Hex("#a16207"),
			"yellow-800": css.Hex("#854d0e"),
			"yellow-900": css.Hex("#713f12"),
			"yellow-950": css.Hex("#422006"),
			
			// Lime
			"lime-50":  css.Hex("#f7fee7"),
			"lime-100": css.Hex("#ecfccb"),
			"lime-200": css.Hex("#d9f99d"),
			"lime-300": css.Hex("#bef264"),
			"lime-400": css.Hex("#a3e635"),
			"lime-500": css.Hex("#84cc16"),
			"lime-600": css.Hex("#65a30d"),
			"lime-700": css.Hex("#4d7c0f"),
			"lime-800": css.Hex("#365314"),
			"lime-900": css.Hex("#1a2e05"),
			"lime-950": css.Hex("#0a1a00"),
			
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
			
			// Emerald
			"emerald-50":  css.Hex("#ecfdf5"),
			"emerald-100": css.Hex("#d1fae5"),
			"emerald-200": css.Hex("#a7f3d0"),
			"emerald-300": css.Hex("#6ee7b7"),
			"emerald-400": css.Hex("#34d399"),
			"emerald-500": css.Hex("#10b981"),
			"emerald-600": css.Hex("#059669"),
			"emerald-700": css.Hex("#047857"),
			"emerald-800": css.Hex("#065f46"),
			"emerald-900": css.Hex("#064e3b"),
			"emerald-950": css.Hex("#022c22"),
			
			// Teal
			"teal-50":  css.Hex("#f0fdfa"),
			"teal-100": css.Hex("#ccfbf1"),
			"teal-200": css.Hex("#99f6e4"),
			"teal-300": css.Hex("#5eead4"),
			"teal-400": css.Hex("#2dd4bf"),
			"teal-500": css.Hex("#14b8a6"),
			"teal-600": css.Hex("#0d9488"),
			"teal-700": css.Hex("#0f766e"),
			"teal-800": css.Hex("#115e59"),
			"teal-900": css.Hex("#134e4a"),
			"teal-950": css.Hex("#042f2e"),
			
			// Cyan
			"cyan-50":  css.Hex("#ecfeff"),
			"cyan-100": css.Hex("#cffafe"),
			"cyan-200": css.Hex("#a5f3fc"),
			"cyan-300": css.Hex("#67e8f9"),
			"cyan-400": css.Hex("#22d3ee"),
			"cyan-500": css.Hex("#06b6d4"),
			"cyan-600": css.Hex("#0891b2"),
			"cyan-700": css.Hex("#0e7490"),
			"cyan-800": css.Hex("#155e75"),
			"cyan-900": css.Hex("#164e63"),
			"cyan-950": css.Hex("#083344"),
			
			// Sky
			"sky-50":  css.Hex("#f0f9ff"),
			"sky-100": css.Hex("#e0f2fe"),
			"sky-200": css.Hex("#bae6fd"),
			"sky-300": css.Hex("#7dd3fc"),
			"sky-400": css.Hex("#38bdf8"),
			"sky-500": css.Hex("#0ea5e9"),
			"sky-600": css.Hex("#0284c7"),
			"sky-700": css.Hex("#0369a1"),
			"sky-800": css.Hex("#075985"),
			"sky-900": css.Hex("#0c4a6e"),
			"sky-950": css.Hex("#082f49"),
			
			// Blue
			"blue-50":  css.Hex("#eff6ff"),
			"blue-100": css.Hex("#dbeafe"),
			"blue-200": css.Hex("#bfdbfe"),
			"blue-300": css.Hex("#93c5fd"),
			"blue-400": css.Hex("#60a5fa"),
			"blue-500": css.Hex("#3b82f6"),
			"blue-600": css.Hex("#2563eb"),
			"blue-700": css.Hex("#1d4ed8"),
			"blue-800": css.Hex("#1e40af"),
			"blue-900": css.Hex("#1e3a8a"),
			"blue-950": css.Hex("#172554"),
			
			// Indigo
			"indigo-50":  css.Hex("#eef2ff"),
			"indigo-100": css.Hex("#e0e7ff"),
			"indigo-200": css.Hex("#c7d2fe"),
			"indigo-300": css.Hex("#a5b4fc"),
			"indigo-400": css.Hex("#818cf8"),
			"indigo-500": css.Hex("#6366f1"),
			"indigo-600": css.Hex("#4f46e5"),
			"indigo-700": css.Hex("#4338ca"),
			"indigo-800": css.Hex("#3730a3"),
			"indigo-900": css.Hex("#312e81"),
			"indigo-950": css.Hex("#1e1b4b"),
			
			// Violet
			"violet-50":  css.Hex("#f5f3ff"),
			"violet-100": css.Hex("#ede9fe"),
			"violet-200": css.Hex("#ddd6fe"),
			"violet-300": css.Hex("#c4b5fd"),
			"violet-400": css.Hex("#a78bfa"),
			"violet-500": css.Hex("#8b5cf6"),
			"violet-600": css.Hex("#7c3aed"),
			"violet-700": css.Hex("#6d28d9"),
			"violet-800": css.Hex("#5b21b6"),
			"violet-900": css.Hex("#4c1d95"),
			"violet-950": css.Hex("#2e1065"),
			
			// Purple
			"purple-50":  css.Hex("#faf5ff"),
			"purple-100": css.Hex("#f3e8ff"),
			"purple-200": css.Hex("#e9d5ff"),
			"purple-300": css.Hex("#d8b4fe"),
			"purple-400": css.Hex("#c084fc"),
			"purple-500": css.Hex("#a855f7"),
			"purple-600": css.Hex("#9333ea"),
			"purple-700": css.Hex("#7e22ce"),
			"purple-800": css.Hex("#6b21a8"),
			"purple-900": css.Hex("#581c87"),
			"purple-950": css.Hex("#3b0764"),
			
			// Fuchsia
			"fuchsia-50":  css.Hex("#fdf4ff"),
			"fuchsia-100": css.Hex("#fae8ff"),
			"fuchsia-200": css.Hex("#f5d0fe"),
			"fuchsia-300": css.Hex("#f0abfc"),
			"fuchsia-400": css.Hex("#e879f9"),
			"fuchsia-500": css.Hex("#d946ef"),
			"fuchsia-600": css.Hex("#c026d3"),
			"fuchsia-700": css.Hex("#a21caf"),
			"fuchsia-800": css.Hex("#86198f"),
			"fuchsia-900": css.Hex("#701a75"),
			"fuchsia-950": css.Hex("#4a044e"),
			
			// Pink
			"pink-50":  css.Hex("#fdf2f8"),
			"pink-100": css.Hex("#fce7f3"),
			"pink-200": css.Hex("#fbcfe8"),
			"pink-300": css.Hex("#f9a8d4"),
			"pink-400": css.Hex("#f472b6"),
			"pink-500": css.Hex("#ec4899"),
			"pink-600": css.Hex("#db2777"),
			"pink-700": css.Hex("#be185d"),
			"pink-800": css.Hex("#9d174d"),
			"pink-900": css.Hex("#831843"),
			"pink-950": css.Hex("#500724"),
			
			// Rose
			"rose-50":  css.Hex("#fff1f2"),
			"rose-100": css.Hex("#ffe4e6"),
			"rose-200": css.Hex("#fecdd3"),
			"rose-300": css.Hex("#fda4af"),
			"rose-400": css.Hex("#fb7185"),
			"rose-500": css.Hex("#f43f5e"),
			"rose-600": css.Hex("#e11d48"),
			"rose-700": css.Hex("#be123c"),
			"rose-800": css.Hex("#9f1239"),
			"rose-900": css.Hex("#881337"),
			"rose-950": css.Hex("#4c0519"),
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
		
		BorderRadius: map[string]css.Length{
			"none": css.Px(0),
			"sm":   css.Rem(0.125),   // 2px
			"":     css.Rem(0.25),    // 4px (default)
			"md":   css.Rem(0.375),   // 6px
			"lg":   css.Rem(0.5),     // 8px
			"xl":   css.Rem(0.75),    // 12px
			"2xl":  css.Rem(1),       // 16px
			"3xl":  css.Rem(1.5),     // 24px
			"full": css.Px(9999),
		},
		
		BorderWidth: map[string]css.Length{
			"":  css.Px(1), // default
			"0": css.Px(0),
			"2": css.Px(2),
			"4": css.Px(4),
			"8": css.Px(8),
		},
		
		BoxShadow: map[string]string{
			"sm":    "0 1px 2px 0 rgb(0 0 0 / 0.05)",
			"":      "0 1px 3px 0 rgb(0 0 0 / 0.1), 0 1px 2px -1px rgb(0 0 0 / 0.1)", // default
			"md":    "0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1)",
			"lg":    "0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1)",
			"xl":    "0 20px 25px -5px rgb(0 0 0 / 0.1), 0 8px 10px -6px rgb(0 0 0 / 0.1)",
			"2xl":   "0 25px 50px -12px rgb(0 0 0 / 0.25)",
			"inner": "inset 0 2px 4px 0 rgb(0 0 0 / 0.05)",
			"none":  "none",
		},
		
		Opacity: map[string]string{
			"0":   "0",
			"5":   "0.05",
			"10":  "0.1",
			"15":  "0.15",
			"20":  "0.2",
			"25":  "0.25",
			"30":  "0.3",
			"35":  "0.35",
			"40":  "0.4",
			"45":  "0.45",
			"50":  "0.5",
			"55":  "0.55",
			"60":  "0.6",
			"65":  "0.65",
			"70":  "0.7",
			"75":  "0.75",
			"80":  "0.8",
			"85":  "0.85",
			"90":  "0.9",
			"95":  "0.95",
			"100": "1",
		},
		
		FontWeight: map[string]string{
			"thin":       "100",
			"extralight": "200",
			"light":      "300",
			"normal":     "400",
			"medium":     "500",
			"semibold":   "600",
			"bold":       "700",
			"extrabold":  "800",
			"black":      "900",
		},
		
		LineHeight: map[string]string{
			"none":    "1",
			"tight":   "1.25",
			"snug":    "1.375",
			"normal":  "1.5",
			"relaxed": "1.625",
			"loose":   "2",
			"3":       ".75rem",   // 12px
			"4":       "1rem",     // 16px
			"5":       "1.25rem",  // 20px
			"6":       "1.5rem",   // 24px
			"7":       "1.75rem",  // 28px
			"8":       "2rem",     // 32px
			"9":       "2.25rem",  // 36px
			"10":      "2.5rem",   // 40px
		},
		
		ZIndex: map[string]string{
			"auto": "auto",
			"0":    "0",
			"10":   "10",
			"20":   "20",
			"30":   "30",
			"40":   "40",
			"50":   "50",
		},
		
		Blur: map[string]string{
			"0":    "0",
			"none": "",
			"sm":   "4px",
			"":     "8px", // default
			"md":   "12px",
			"lg":   "16px",
			"xl":   "24px",
			"2xl":  "40px",
			"3xl":  "64px",
		},
		
		Brightness: map[string]string{
			"0":   "0",
			"50":  ".5",
			"75":  ".75",
			"90":  ".9",
			"95":  ".95",
			"100": "1",
			"105": "1.05",
			"110": "1.1",
			"125": "1.25",
			"150": "1.5",
			"200": "2",
		},
		
		Contrast: map[string]string{
			"0":   "0",
			"50":  ".5",
			"75":  ".75",
			"100": "1",
			"125": "1.25",
			"150": "1.5",
			"200": "2",
		},
		
		Grayscale: map[string]string{
			"0":   "0",
			"":    "100%", // default
		},
		
		Invert: map[string]string{
			"0": "0",
			"": "100%", // default
		},
		
		Saturate: map[string]string{
			"0":   "0",
			"50":  ".5",
			"100": "1",
			"150": "1.5",
			"200": "2",
		},
		
		Sepia: map[string]string{
			"0": "0",
			"": "100%", // default
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
	
	// Merge border radius
	if overrides.BorderRadius != nil {
		for k, v := range overrides.BorderRadius {
			base.BorderRadius[k] = v
		}
	}
	
	// Merge border width
	if overrides.BorderWidth != nil {
		for k, v := range overrides.BorderWidth {
			base.BorderWidth[k] = v
		}
	}
	
	// Merge box shadow
	if overrides.BoxShadow != nil {
		for k, v := range overrides.BoxShadow {
			base.BoxShadow[k] = v
		}
	}
	
	// Merge opacity
	if overrides.Opacity != nil {
		for k, v := range overrides.Opacity {
			base.Opacity[k] = v
		}
	}
	
	// Merge font weight
	if overrides.FontWeight != nil {
		for k, v := range overrides.FontWeight {
			base.FontWeight[k] = v
		}
	}
	
	// Merge line height
	if overrides.LineHeight != nil {
		for k, v := range overrides.LineHeight {
			base.LineHeight[k] = v
		}
	}
	
	// Merge z-index
	if overrides.ZIndex != nil {
		for k, v := range overrides.ZIndex {
			base.ZIndex[k] = v
		}
	}
	
	// Merge blur
	if overrides.Blur != nil {
		for k, v := range overrides.Blur {
			base.Blur[k] = v
		}
	}
	
	// Merge brightness
	if overrides.Brightness != nil {
		for k, v := range overrides.Brightness {
			base.Brightness[k] = v
		}
	}
	
	// Merge contrast
	if overrides.Contrast != nil {
		for k, v := range overrides.Contrast {
			base.Contrast[k] = v
		}
	}
	
	// Merge grayscale
	if overrides.Grayscale != nil {
		for k, v := range overrides.Grayscale {
			base.Grayscale[k] = v
		}
	}
	
	// Merge invert
	if overrides.Invert != nil {
		for k, v := range overrides.Invert {
			base.Invert[k] = v
		}
	}
	
	// Merge saturate
	if overrides.Saturate != nil {
		for k, v := range overrides.Saturate {
			base.Saturate[k] = v
		}
	}
	
	// Merge sepia
	if overrides.Sepia != nil {
		for k, v := range overrides.Sepia {
			base.Sepia[k] = v
		}
	}
	
	return base
}