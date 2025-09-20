// This file provides the default Tailwind configuration implementation
// with all the standard values from the Tailwind CSS framework.

package tailwind

import (
	"github.com/ahmed-com/typesafe-css/css"
	"github.com/ahmed-com/typesafe-css/cssgen"
)

// DefaultConfig returns a complete Tailwind configuration with all default values.
func DefaultConfig() Config {
	return Config{
		DarkMode: DarkModeConfig{
			Strategy: DarkModeMedia,
		},
		Theme: DefaultTheme(),
	}
}

// DefaultTheme returns the default Tailwind theme configuration.
func DefaultTheme() ThemeConfig {
	return ThemeConfig{
		Colors:          DefaultColors(),
		AccentColor:     DefaultAccentColor(),
		BackgroundColor: DefaultBackgroundColor(),
		BorderColor:     DefaultBorderColor(),
		CaretColor:      DefaultCaretColor(),
		Spacing:         DefaultSpacing(),
		BorderRadius:    DefaultBorderRadius(),
		BorderWidth:     DefaultBorderWidth(),
		FontFamily:      DefaultFontFamily(),
		FontSize:        DefaultFontSize(),
		FontWeight:      DefaultFontWeight(),
		AspectRatio:     DefaultAspectRatio(),
		Animation:       DefaultAnimation(),
		Blur:            DefaultBlur(),
		Brightness:      DefaultBrightness(),
		Opacity:         DefaultOpacity(),
		Cursor:          DefaultCursor(),
		Screens:         DefaultScreens(),
	}
}

// Default color configuration

// DefaultColors returns the default Tailwind color palette.
func DefaultColors() ColorsConfig {
	return ColorsConfig{
		// CSS system colors
		Inherit:     ColorFromKeyword("inherit", "inherit"),
		Current:     ColorFromKeyword("current", "currentColor"),
		Transparent: ColorFromKeyword("transparent", "transparent"),

		// Base colors
		Black: ColorFromHex("black", "#000000"),
		White: ColorFromHex("white", "#ffffff"),

		// Slate color scale
		Slate50:  ColorFromHex("slate-50", "#f8fafc"),
		Slate100: ColorFromHex("slate-100", "#f1f5f9"),
		Slate200: ColorFromHex("slate-200", "#e2e8f0"),
		Slate300: ColorFromHex("slate-300", "#cbd5e1"),
		Slate400: ColorFromHex("slate-400", "#94a3b8"),
		Slate500: ColorFromHex("slate-500", "#64748b"),
		Slate600: ColorFromHex("slate-600", "#475569"),
		Slate700: ColorFromHex("slate-700", "#334155"),
		Slate800: ColorFromHex("slate-800", "#1e293b"),
		Slate900: ColorFromHex("slate-900", "#0f172a"),
		Slate950: ColorFromHex("slate-950", "#020617"),

		// Gray color scale
		Gray50:  ColorFromHex("gray-50", "#f9fafb"),
		Gray100: ColorFromHex("gray-100", "#f3f4f6"),
		Gray200: ColorFromHex("gray-200", "#e5e7eb"),
		Gray300: ColorFromHex("gray-300", "#d1d5db"),
		Gray400: ColorFromHex("gray-400", "#9ca3af"),
		Gray500: ColorFromHex("gray-500", "#6b7280"),
		Gray600: ColorFromHex("gray-600", "#4b5563"),
		Gray700: ColorFromHex("gray-700", "#374151"),
		Gray800: ColorFromHex("gray-800", "#1f2937"),
		Gray900: ColorFromHex("gray-900", "#111827"),
		Gray950: ColorFromHex("gray-950", "#030712"),

		// Zinc color scale
		Zinc50:  ColorFromHex("zinc-50", "#fafafa"),
		Zinc100: ColorFromHex("zinc-100", "#f4f4f5"),
		Zinc200: ColorFromHex("zinc-200", "#e4e4e7"),
		Zinc300: ColorFromHex("zinc-300", "#d4d4d8"),
		Zinc400: ColorFromHex("zinc-400", "#a1a1aa"),
		Zinc500: ColorFromHex("zinc-500", "#71717a"),
		Zinc600: ColorFromHex("zinc-600", "#52525b"),
		Zinc700: ColorFromHex("zinc-700", "#3f3f46"),
		Zinc800: ColorFromHex("zinc-800", "#27272a"),
		Zinc900: ColorFromHex("zinc-900", "#18181b"),
		Zinc950: ColorFromHex("zinc-950", "#09090b"),

		// Neutral color scale
		Neutral50:  ColorFromHex("neutral-50", "#fafafa"),
		Neutral100: ColorFromHex("neutral-100", "#f5f5f5"),
		Neutral200: ColorFromHex("neutral-200", "#e5e5e5"),
		Neutral300: ColorFromHex("neutral-300", "#d4d4d4"),
		Neutral400: ColorFromHex("neutral-400", "#a3a3a3"),
		Neutral500: ColorFromHex("neutral-500", "#737373"),
		Neutral600: ColorFromHex("neutral-600", "#525252"),
		Neutral700: ColorFromHex("neutral-700", "#404040"),
		Neutral800: ColorFromHex("neutral-800", "#262626"),
		Neutral900: ColorFromHex("neutral-900", "#171717"),
		Neutral950: ColorFromHex("neutral-950", "#0a0a0a"),

		// Stone color scale
		Stone50:  ColorFromHex("stone-50", "#fafaf9"),
		Stone100: ColorFromHex("stone-100", "#f5f5f4"),
		Stone200: ColorFromHex("stone-200", "#e7e5e4"),
		Stone300: ColorFromHex("stone-300", "#d6d3d1"),
		Stone400: ColorFromHex("stone-400", "#a8a29e"),
		Stone500: ColorFromHex("stone-500", "#78716c"),
		Stone600: ColorFromHex("stone-600", "#57534e"),
		Stone700: ColorFromHex("stone-700", "#44403c"),
		Stone800: ColorFromHex("stone-800", "#292524"),
		Stone900: ColorFromHex("stone-900", "#1c1917"),
		Stone950: ColorFromHex("stone-950", "#0c0a09"),

		// Red color scale
		Red50:  ColorFromHex("red-50", "#fef2f2"),
		Red100: ColorFromHex("red-100", "#fee2e2"),
		Red200: ColorFromHex("red-200", "#fecaca"),
		Red300: ColorFromHex("red-300", "#fca5a5"),
		Red400: ColorFromHex("red-400", "#f87171"),
		Red500: ColorFromHex("red-500", "#ef4444"),
		Red600: ColorFromHex("red-600", "#dc2626"),
		Red700: ColorFromHex("red-700", "#b91c1c"),
		Red800: ColorFromHex("red-800", "#991b1b"),
		Red900: ColorFromHex("red-900", "#7f1d1d"),
		Red950: ColorFromHex("red-950", "#450a0a"),

		// Orange color scale
		Orange50:  ColorFromHex("orange-50", "#fff7ed"),
		Orange100: ColorFromHex("orange-100", "#ffedd5"),
		Orange200: ColorFromHex("orange-200", "#fed7aa"),
		Orange300: ColorFromHex("orange-300", "#fdba74"),
		Orange400: ColorFromHex("orange-400", "#fb923c"),
		Orange500: ColorFromHex("orange-500", "#f97316"),
		Orange600: ColorFromHex("orange-600", "#ea580c"),
		Orange700: ColorFromHex("orange-700", "#c2410c"),
		Orange800: ColorFromHex("orange-800", "#9a3412"),
		Orange900: ColorFromHex("orange-900", "#7c2d12"),
		Orange950: ColorFromHex("orange-950", "#431407"),

		// Amber color scale
		Amber50:  ColorFromHex("amber-50", "#fffbeb"),
		Amber100: ColorFromHex("amber-100", "#fef3c7"),
		Amber200: ColorFromHex("amber-200", "#fde68a"),
		Amber300: ColorFromHex("amber-300", "#fcd34d"),
		Amber400: ColorFromHex("amber-400", "#fbbf24"),
		Amber500: ColorFromHex("amber-500", "#f59e0b"),
		Amber600: ColorFromHex("amber-600", "#d97706"),
		Amber700: ColorFromHex("amber-700", "#b45309"),
		Amber800: ColorFromHex("amber-800", "#92400e"),
		Amber900: ColorFromHex("amber-900", "#78350f"),
		Amber950: ColorFromHex("amber-950", "#451a03"),

		// Yellow color scale
		Yellow50:  ColorFromHex("yellow-50", "#fefce8"),
		Yellow100: ColorFromHex("yellow-100", "#fef9c3"),
		Yellow200: ColorFromHex("yellow-200", "#fef08a"),
		Yellow300: ColorFromHex("yellow-300", "#fde047"),
		Yellow400: ColorFromHex("yellow-400", "#facc15"),
		Yellow500: ColorFromHex("yellow-500", "#eab308"),
		Yellow600: ColorFromHex("yellow-600", "#ca8a04"),
		Yellow700: ColorFromHex("yellow-700", "#a16207"),
		Yellow800: ColorFromHex("yellow-800", "#854d0e"),
		Yellow900: ColorFromHex("yellow-900", "#713f12"),
		Yellow950: ColorFromHex("yellow-950", "#422006"),

		// Lime color scale
		Lime50:  ColorFromHex("lime-50", "#f7fee7"),
		Lime100: ColorFromHex("lime-100", "#ecfccb"),
		Lime200: ColorFromHex("lime-200", "#d9f99d"),
		Lime300: ColorFromHex("lime-300", "#bef264"),
		Lime400: ColorFromHex("lime-400", "#a3e635"),
		Lime500: ColorFromHex("lime-500", "#84cc16"),
		Lime600: ColorFromHex("lime-600", "#65a30d"),
		Lime700: ColorFromHex("lime-700", "#4d7c0f"),
		Lime800: ColorFromHex("lime-800", "#3f6212"),
		Lime900: ColorFromHex("lime-900", "#365314"),
		Lime950: ColorFromHex("lime-950", "#1a2e05"),

		// Green color scale
		Green50:  ColorFromHex("green-50", "#f0fdf4"),
		Green100: ColorFromHex("green-100", "#dcfce7"),
		Green200: ColorFromHex("green-200", "#bbf7d0"),
		Green300: ColorFromHex("green-300", "#86efac"),
		Green400: ColorFromHex("green-400", "#4ade80"),
		Green500: ColorFromHex("green-500", "#22c55e"),
		Green600: ColorFromHex("green-600", "#16a34a"),
		Green700: ColorFromHex("green-700", "#15803d"),
		Green800: ColorFromHex("green-800", "#166534"),
		Green900: ColorFromHex("green-900", "#14532d"),
		Green950: ColorFromHex("green-950", "#052e16"),

		// Emerald color scale
		Emerald50:  ColorFromHex("emerald-50", "#ecfdf5"),
		Emerald100: ColorFromHex("emerald-100", "#d1fae5"),
		Emerald200: ColorFromHex("emerald-200", "#a7f3d0"),
		Emerald300: ColorFromHex("emerald-300", "#6ee7b7"),
		Emerald400: ColorFromHex("emerald-400", "#34d399"),
		Emerald500: ColorFromHex("emerald-500", "#10b981"),
		Emerald600: ColorFromHex("emerald-600", "#059669"),
		Emerald700: ColorFromHex("emerald-700", "#047857"),
		Emerald800: ColorFromHex("emerald-800", "#065f46"),
		Emerald900: ColorFromHex("emerald-900", "#064e3b"),
		Emerald950: ColorFromHex("emerald-950", "#022c22"),

		// Teal color scale
		Teal50:  ColorFromHex("teal-50", "#f0fdfa"),
		Teal100: ColorFromHex("teal-100", "#ccfbf1"),
		Teal200: ColorFromHex("teal-200", "#99f6e4"),
		Teal300: ColorFromHex("teal-300", "#5eead4"),
		Teal400: ColorFromHex("teal-400", "#2dd4bf"),
		Teal500: ColorFromHex("teal-500", "#14b8a6"),
		Teal600: ColorFromHex("teal-600", "#0d9488"),
		Teal700: ColorFromHex("teal-700", "#0f766e"),
		Teal800: ColorFromHex("teal-800", "#115e59"),
		Teal900: ColorFromHex("teal-900", "#134e4a"),
		Teal950: ColorFromHex("teal-950", "#042f2e"),

		// Cyan color scale
		Cyan50:  ColorFromHex("cyan-50", "#ecfeff"),
		Cyan100: ColorFromHex("cyan-100", "#cffafe"),
		Cyan200: ColorFromHex("cyan-200", "#a5f3fc"),
		Cyan300: ColorFromHex("cyan-300", "#67e8f9"),
		Cyan400: ColorFromHex("cyan-400", "#22d3ee"),
		Cyan500: ColorFromHex("cyan-500", "#06b6d4"),
		Cyan600: ColorFromHex("cyan-600", "#0891b2"),
		Cyan700: ColorFromHex("cyan-700", "#0e7490"),
		Cyan800: ColorFromHex("cyan-800", "#155e75"),
		Cyan900: ColorFromHex("cyan-900", "#164e63"),
		Cyan950: ColorFromHex("cyan-950", "#083344"),

		// Sky color scale
		Sky50:  ColorFromHex("sky-50", "#f0f9ff"),
		Sky100: ColorFromHex("sky-100", "#e0f2fe"),
		Sky200: ColorFromHex("sky-200", "#bae6fd"),
		Sky300: ColorFromHex("sky-300", "#7dd3fc"),
		Sky400: ColorFromHex("sky-400", "#38bdf8"),
		Sky500: ColorFromHex("sky-500", "#0ea5e9"),
		Sky600: ColorFromHex("sky-600", "#0284c7"),
		Sky700: ColorFromHex("sky-700", "#0369a1"),
		Sky800: ColorFromHex("sky-800", "#075985"),
		Sky900: ColorFromHex("sky-900", "#0c4a6e"),
		Sky950: ColorFromHex("sky-950", "#082f49"),

		// Blue color scale
		Blue50:  ColorFromHex("blue-50", "#eff6ff"),
		Blue100: ColorFromHex("blue-100", "#dbeafe"),
		Blue200: ColorFromHex("blue-200", "#bfdbfe"),
		Blue300: ColorFromHex("blue-300", "#93c5fd"),
		Blue400: ColorFromHex("blue-400", "#60a5fa"),
		Blue500: ColorFromHex("blue-500", "#3b82f6"),
		Blue600: ColorFromHex("blue-600", "#2563eb"),
		Blue700: ColorFromHex("blue-700", "#1d4ed8"),
		Blue800: ColorFromHex("blue-800", "#1e40af"),
		Blue900: ColorFromHex("blue-900", "#1e3a8a"),
		Blue950: ColorFromHex("blue-950", "#172554"),

		// Indigo color scale
		Indigo50:  ColorFromHex("indigo-50", "#eef2ff"),
		Indigo100: ColorFromHex("indigo-100", "#e0e7ff"),
		Indigo200: ColorFromHex("indigo-200", "#c7d2fe"),
		Indigo300: ColorFromHex("indigo-300", "#a5b4fc"),
		Indigo400: ColorFromHex("indigo-400", "#818cf8"),
		Indigo500: ColorFromHex("indigo-500", "#6366f1"),
		Indigo600: ColorFromHex("indigo-600", "#4f46e5"),
		Indigo700: ColorFromHex("indigo-700", "#4338ca"),
		Indigo800: ColorFromHex("indigo-800", "#3730a3"),
		Indigo900: ColorFromHex("indigo-900", "#312e81"),
		Indigo950: ColorFromHex("indigo-950", "#1e1b4b"),

		// Violet color scale
		Violet50:  ColorFromHex("violet-50", "#f5f3ff"),
		Violet100: ColorFromHex("violet-100", "#ede9fe"),
		Violet200: ColorFromHex("violet-200", "#ddd6fe"),
		Violet300: ColorFromHex("violet-300", "#c4b5fd"),
		Violet400: ColorFromHex("violet-400", "#a78bfa"),
		Violet500: ColorFromHex("violet-500", "#8b5cf6"),
		Violet600: ColorFromHex("violet-600", "#7c3aed"),
		Violet700: ColorFromHex("violet-700", "#6d28d9"),
		Violet800: ColorFromHex("violet-800", "#5b21b6"),
		Violet900: ColorFromHex("violet-900", "#4c1d95"),
		Violet950: ColorFromHex("violet-950", "#2e1065"),

		// Purple color scale
		Purple50:  ColorFromHex("purple-50", "#faf5ff"),
		Purple100: ColorFromHex("purple-100", "#f3e8ff"),
		Purple200: ColorFromHex("purple-200", "#e9d5ff"),
		Purple300: ColorFromHex("purple-300", "#d8b4fe"),
		Purple400: ColorFromHex("purple-400", "#c084fc"),
		Purple500: ColorFromHex("purple-500", "#a855f7"),
		Purple600: ColorFromHex("purple-600", "#9333ea"),
		Purple700: ColorFromHex("purple-700", "#7e22ce"),
		Purple800: ColorFromHex("purple-800", "#6b21a8"),
		Purple900: ColorFromHex("purple-900", "#581c87"),
		Purple950: ColorFromHex("purple-950", "#3b0764"),

		// Fuchsia color scale
		Fuchsia50:  ColorFromHex("fuchsia-50", "#fdf4ff"),
		Fuchsia100: ColorFromHex("fuchsia-100", "#fae8ff"),
		Fuchsia200: ColorFromHex("fuchsia-200", "#f5d0fe"),
		Fuchsia300: ColorFromHex("fuchsia-300", "#f0abfc"),
		Fuchsia400: ColorFromHex("fuchsia-400", "#e879f9"),
		Fuchsia500: ColorFromHex("fuchsia-500", "#d946ef"),
		Fuchsia600: ColorFromHex("fuchsia-600", "#c026d3"),
		Fuchsia700: ColorFromHex("fuchsia-700", "#a21caf"),
		Fuchsia800: ColorFromHex("fuchsia-800", "#86198f"),
		Fuchsia900: ColorFromHex("fuchsia-900", "#701a75"),
		Fuchsia950: ColorFromHex("fuchsia-950", "#4a044e"),

		// Pink color scale
		Pink50:  ColorFromHex("pink-50", "#fdf2f8"),
		Pink100: ColorFromHex("pink-100", "#fce7f3"),
		Pink200: ColorFromHex("pink-200", "#fbcfe8"),
		Pink300: ColorFromHex("pink-300", "#f9a8d4"),
		Pink400: ColorFromHex("pink-400", "#f472b6"),
		Pink500: ColorFromHex("pink-500", "#ec4899"),
		Pink600: ColorFromHex("pink-600", "#db2777"),
		Pink700: ColorFromHex("pink-700", "#be185d"),
		Pink800: ColorFromHex("pink-800", "#9d174d"),
		Pink900: ColorFromHex("pink-900", "#831843"),
		Pink950: ColorFromHex("pink-950", "#500724"),

		// Rose color scale
		Rose50:  ColorFromHex("rose-50", "#fff1f2"),
		Rose100: ColorFromHex("rose-100", "#ffe4e6"),
		Rose200: ColorFromHex("rose-200", "#fecdd3"),
		Rose300: ColorFromHex("rose-300", "#fda4af"),
		Rose400: ColorFromHex("rose-400", "#fb7185"),
		Rose500: ColorFromHex("rose-500", "#f43f5e"),
		Rose600: ColorFromHex("rose-600", "#e11d48"),
		Rose700: ColorFromHex("rose-700", "#be123c"),
		Rose800: ColorFromHex("rose-800", "#9f1239"),
		Rose900: ColorFromHex("rose-900", "#881337"),
		Rose950: ColorFromHex("rose-950", "#4c0519"),
	}
}

// Helper function to create color from keyword
func ColorFromKeyword(name, keyword string) StaticColor {
	return StaticColor{Name: name, Value: css.Color(keyword)}
}

// Default accent color configuration
func DefaultAccentColor() AccentColorConfig {
	return AccentColorConfig{
		Auto: KeywordValue{Name: "auto", Keyword: css.Keyword("auto")},
	}
}

// Default background color configuration
func DefaultBackgroundColor() BackgroundColorConfig {
	return BackgroundColorConfig{}
}

// Default border color configuration
func DefaultBorderColor() BorderColorConfig {
	return BorderColorConfig{
		DEFAULT: ColorFromHex("DEFAULT", "#e5e7eb"), // gray-200
	}
}

// Default caret color configuration
func DefaultCaretColor() CaretColorConfig {
	return CaretColorConfig{}
}

// Default spacing configuration
func DefaultSpacing() SpacingConfig {
	return SpacingConfig{
		Px:      LengthFromPx("px", 1),
		Size0:   LengthFromPx("0", 0),
		Size0_5: LengthFromRem("0.5", 0.125),
		Size1:   LengthFromRem("1", 0.25),
		Size1_5: LengthFromRem("1.5", 0.375),
		Size2:   LengthFromRem("2", 0.5),
		Size2_5: LengthFromRem("2.5", 0.625),
		Size3:   LengthFromRem("3", 0.75),
		Size3_5: LengthFromRem("3.5", 0.875),
		Size4:   LengthFromRem("4", 1),
		Size5:   LengthFromRem("5", 1.25),
		Size6:   LengthFromRem("6", 1.5),
		Size7:   LengthFromRem("7", 1.75),
		Size8:   LengthFromRem("8", 2),
		Size9:   LengthFromRem("9", 2.25),
		Size10:  LengthFromRem("10", 2.5),
		Size11:  LengthFromRem("11", 2.75),
		Size12:  LengthFromRem("12", 3),
		Size14:  LengthFromRem("14", 3.5),
		Size16:  LengthFromRem("16", 4),
		Size20:  LengthFromRem("20", 5),
		Size24:  LengthFromRem("24", 6),
		Size28:  LengthFromRem("28", 7),
		Size32:  LengthFromRem("32", 8),
		Size36:  LengthFromRem("36", 9),
		Size40:  LengthFromRem("40", 10),
		Size44:  LengthFromRem("44", 11),
		Size48:  LengthFromRem("48", 12),
		Size52:  LengthFromRem("52", 13),
		Size56:  LengthFromRem("56", 14),
		Size60:  LengthFromRem("60", 15),
		Size64:  LengthFromRem("64", 16),
		Size72:  LengthFromRem("72", 18),
		Size80:  LengthFromRem("80", 20),
		Size96:  LengthFromRem("96", 24),
	}
}

// Default border radius configuration
func DefaultBorderRadius() BorderRadiusConfig {
	return BorderRadiusConfig{
		None:    LengthFromPx("none", 0),
		Sm:      LengthFromRem("sm", 0.125),
		Base:    LengthFromRem("base", 0.25),
		Md:      LengthFromRem("md", 0.375),
		Lg:      LengthFromRem("lg", 0.5),
		Xl:      LengthFromRem("xl", 0.75),
		Size2xl: LengthFromRem("2xl", 1),
		Size3xl: LengthFromRem("3xl", 1.5),
		Full:    LengthFromPx("full", 9999),
	}
}

// Default border width configuration
func DefaultBorderWidth() BorderWidthConfig {
	return BorderWidthConfig{
		DEFAULT: LengthFromPx("DEFAULT", 1),
		Size0:   LengthFromPx("0", 0),
		Size2:   LengthFromPx("2", 2),
		Size4:   LengthFromPx("4", 4),
		Size8:   LengthFromPx("8", 8),
	}
}

// Default font family configuration
func DefaultFontFamily() FontFamilyConfig {
	return FontFamilyConfig{
		Sans: []string{
			"ui-sans-serif",
			"system-ui",
			"sans-serif",
			"\"Apple Color Emoji\"",
			"\"Segoe UI Emoji\"",
			"\"Segoe UI Symbol\"",
			"\"Noto Color Emoji\"",
		},
		Serif: []string{
			"ui-serif",
			"Georgia",
			"Cambria",
			"\"Times New Roman\"",
			"Times",
			"serif",
		},
		Mono: []string{
			"ui-monospace",
			"SFMono-Regular",
			"Menlo",
			"Monaco",
			"Consolas",
			"\"Liberation Mono\"",
			"\"Courier New\"",
			"monospace",
		},
	}
}

// Default font size configuration
func DefaultFontSize() FontSizeConfig {
	return FontSizeConfig{
		Xs: FontSizeValue{
			Size:       LengthFromRem("xs", 0.75),
			LineHeight: StaticLength{Name: "xs-lh", Value: css.Rem(1)},
		},
		Sm: FontSizeValue{
			Size:       LengthFromRem("sm", 0.875),
			LineHeight: StaticLength{Name: "sm-lh", Value: css.Rem(1.25)},
		},
		Base: FontSizeValue{
			Size:       LengthFromRem("base", 1),
			LineHeight: StaticLength{Name: "base-lh", Value: css.Rem(1.5)},
		},
		Lg: FontSizeValue{
			Size:       LengthFromRem("lg", 1.125),
			LineHeight: StaticLength{Name: "lg-lh", Value: css.Rem(1.75)},
		},
		Xl: FontSizeValue{
			Size:       LengthFromRem("xl", 1.25),
			LineHeight: StaticLength{Name: "xl-lh", Value: css.Rem(1.75)},
		},
		Size2xl: FontSizeValue{
			Size:       LengthFromRem("2xl", 1.5),
			LineHeight: StaticLength{Name: "2xl-lh", Value: css.Rem(2)},
		},
		Size3xl: FontSizeValue{
			Size:       LengthFromRem("3xl", 1.875),
			LineHeight: StaticLength{Name: "3xl-lh", Value: css.Rem(2.25)},
		},
		Size4xl: FontSizeValue{
			Size:       LengthFromRem("4xl", 2.25),
			LineHeight: StaticLength{Name: "4xl-lh", Value: css.Rem(2.5)},
		},
		Size5xl: FontSizeValue{
			Size:       LengthFromRem("5xl", 3),
			LineHeight: StaticNumber{Name: "5xl-lh", Value: css.Raw("1")},
		},
		Size6xl: FontSizeValue{
			Size:       LengthFromRem("6xl", 3.75),
			LineHeight: StaticNumber{Name: "6xl-lh", Value: css.Raw("1")},
		},
		Size7xl: FontSizeValue{
			Size:       LengthFromRem("7xl", 4.5),
			LineHeight: StaticNumber{Name: "7xl-lh", Value: css.Raw("1")},
		},
		Size8xl: FontSizeValue{
			Size:       LengthFromRem("8xl", 6),
			LineHeight: StaticNumber{Name: "8xl-lh", Value: css.Raw("1")},
		},
		Size9xl: FontSizeValue{
			Size:       LengthFromRem("9xl", 8),
			LineHeight: StaticNumber{Name: "9xl-lh", Value: css.Raw("1")},
		},
	}
}

// Default font weight configuration
func DefaultFontWeight() FontWeightConfig {
	return FontWeightConfig{
		Thin:       FontWeightKeyword("thin", cssgen.FontWeightValNormal), // Will be replaced with actual value
		Extralight: FontWeightKeyword("extralight", cssgen.FontWeightValNormal),
		Light:      FontWeightKeyword("light", cssgen.FontWeightValNormal),
		Normal:     FontWeightKeyword("normal", cssgen.FontWeightValNormal),
		Medium:     FontWeightKeyword("medium", cssgen.FontWeightValNormal),
		Semibold:   FontWeightKeyword("semibold", cssgen.FontWeightValNormal),
		Bold:       FontWeightKeyword("bold", cssgen.FontWeightValBold),
		Extrabold:  FontWeightKeyword("extrabold", cssgen.FontWeightValNormal),
		Black:      FontWeightKeyword("black", cssgen.FontWeightValNormal),
	}
}

// Default aspect ratio configuration
func DefaultAspectRatio() AspectRatioConfig {
	return AspectRatioConfig{
		Auto:   KeywordValue{Name: "auto", Keyword: css.Keyword("auto")},
		Square: "1 / 1",
		Video:  "16 / 9",
	}
}

// Default animation configuration
func DefaultAnimation() AnimationConfig {
	return AnimationConfig{
		None:   KeywordValue{Name: "none", Keyword: css.Keyword("none")},
		Spin:   "spin 1s linear infinite",
		Ping:   "ping 1s cubic-bezier(0, 0, 0.2, 1) infinite",
		Pulse:  "pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite",
		Bounce: "bounce 1s infinite",
	}
}

// Default blur configuration
func DefaultBlur() BlurConfig {
	return BlurConfig{
		Size0:   LengthFromPx("0", 0),
		None:    "",
		Sm:      LengthFromPx("sm", 4),
		DEFAULT: LengthFromPx("DEFAULT", 8),
		Md:      LengthFromPx("md", 12),
		Lg:      LengthFromPx("lg", 16),
		Xl:      LengthFromPx("xl", 24),
		Size2xl: LengthFromPx("2xl", 40),
		Size3xl: LengthFromPx("3xl", 64),
	}
}

// Default brightness configuration
func DefaultBrightness() BrightnessConfig {
	return BrightnessConfig{
		Size0:   NumberFromFloat("0", 0),
		Size50:  NumberFromFloat("50", 0.5),
		Size75:  NumberFromFloat("75", 0.75),
		Size90:  NumberFromFloat("90", 0.9),
		Size95:  NumberFromFloat("95", 0.95),
		Size100: NumberFromFloat("100", 1),
		Size105: NumberFromFloat("105", 1.05),
		Size110: NumberFromFloat("110", 1.1),
		Size125: NumberFromFloat("125", 1.25),
		Size150: NumberFromFloat("150", 1.5),
		Size200: NumberFromFloat("200", 2),
	}
}

// Default opacity configuration
func DefaultOpacity() OpacityConfig {
	return OpacityConfig{
		Size0:   NumberFromFloat("0", 0),
		Size5:   NumberFromFloat("5", 0.05),
		Size10:  NumberFromFloat("10", 0.1),
		Size15:  NumberFromFloat("15", 0.15),
		Size20:  NumberFromFloat("20", 0.2),
		Size25:  NumberFromFloat("25", 0.25),
		Size30:  NumberFromFloat("30", 0.3),
		Size35:  NumberFromFloat("35", 0.35),
		Size40:  NumberFromFloat("40", 0.4),
		Size45:  NumberFromFloat("45", 0.45),
		Size50:  NumberFromFloat("50", 0.5),
		Size55:  NumberFromFloat("55", 0.55),
		Size60:  NumberFromFloat("60", 0.6),
		Size65:  NumberFromFloat("65", 0.65),
		Size70:  NumberFromFloat("70", 0.7),
		Size75:  NumberFromFloat("75", 0.75),
		Size80:  NumberFromFloat("80", 0.8),
		Size85:  NumberFromFloat("85", 0.85),
		Size90:  NumberFromFloat("90", 0.9),
		Size95:  NumberFromFloat("95", 0.95),
		Size100: NumberFromFloat("100", 1),
	}
}

// Default cursor configuration
func DefaultCursor() CursorConfig {
	return CursorConfig{
		Auto:         CursorKeyword("auto", cssgen.CursorValAuto),
		Default:      CursorKeyword("default", cssgen.CursorValDefault),
		Pointer:      CursorKeyword("pointer", cssgen.CursorValPointer),
		Wait:         CursorKeyword("wait", cssgen.CursorValWait),
		Text:         CursorKeyword("text", cssgen.CursorValText),
		Move:         CursorKeyword("move", cssgen.CursorValMove),
		Help:         CursorKeyword("help", cssgen.CursorValHelp),
		NotAllowed:   CursorKeyword("not-allowed", cssgen.CursorValNotAllowed),
		None:         CursorKeyword("none", cssgen.CursorValNone),
		ContextMenu:  CursorKeyword("context-menu", cssgen.CursorValContextMenu),
		Progress:     CursorKeyword("progress", cssgen.CursorValProgress),
		Cell:         CursorKeyword("cell", cssgen.CursorValCell),
		Crosshair:    CursorKeyword("crosshair", cssgen.CursorValCrosshair),
		VerticalText: CursorKeyword("vertical-text", cssgen.CursorValVerticalText),
		Alias:        CursorKeyword("alias", cssgen.CursorValAlias),
		Copy:         CursorKeyword("copy", cssgen.CursorValCopy),
		NoDrop:       CursorKeyword("no-drop", cssgen.CursorValNoDrop),
		Grab:         CursorKeyword("grab", cssgen.CursorValGrab),
		Grabbing:     CursorKeyword("grabbing", cssgen.CursorValGrabbing),
		AllScroll:    CursorKeyword("all-scroll", cssgen.CursorValAllScroll),
		ColResize:    CursorKeyword("col-resize", cssgen.CursorValColResize),
		RowResize:    CursorKeyword("row-resize", cssgen.CursorValRowResize),
		NResize:      CursorKeyword("n-resize", cssgen.CursorValNResize),
		EResize:      CursorKeyword("e-resize", cssgen.CursorValEResize),
		SResize:      CursorKeyword("s-resize", cssgen.CursorValSResize),
		WResize:      CursorKeyword("w-resize", cssgen.CursorValWResize),
		NeResize:     CursorKeyword("ne-resize", cssgen.CursorValNeResize),
		NwResize:     CursorKeyword("nw-resize", cssgen.CursorValNwResize),
		SeResize:     CursorKeyword("se-resize", cssgen.CursorValSeResize),
		SwResize:     CursorKeyword("sw-resize", cssgen.CursorValSwResize),
		EwResize:     CursorKeyword("ew-resize", cssgen.CursorValEwResize),
		NsResize:     CursorKeyword("ns-resize", cssgen.CursorValNsResize),
		NeswResize:   CursorKeyword("nesw-resize", cssgen.CursorValNeswResize),
		NwseResize:   CursorKeyword("nwse-resize", cssgen.CursorValNwseResize),
		ZoomIn:       CursorKeyword("zoom-in", cssgen.CursorValZoomIn),
		ZoomOut:      CursorKeyword("zoom-out", cssgen.CursorValZoomOut),
	}
}

// Default screens configuration
func DefaultScreens() ScreensConfig {
	return ScreensConfig{
		Sm:      LengthFromPx("sm", 640),
		Md:      LengthFromPx("md", 768),
		Lg:      LengthFromPx("lg", 1024),
		Xl:      LengthFromPx("xl", 1280),
		Size2xl: LengthFromPx("2xl", 1536),
	}
}
