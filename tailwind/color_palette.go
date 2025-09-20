package tailwind

import "github.com/ahmed-com/typesafe-css/css"

// ColorToken represents a named Tailwind color entry.
type ColorToken struct {
	suffix string
	value  css.Color
}

func newColorToken(suffix string, value css.Color) ColorToken {
	return ColorToken{suffix: suffix, value: value}
}

// NewColorToken constructs a ColorToken for custom palette extensions.
func NewColorToken(suffix string, value css.Color) ColorToken {
	return newColorToken(suffix, value)
}

// Suffix returns the Tailwind class suffix for this color (e.g. "blue-500").
func (c ColorToken) Suffix() string { return c.suffix }

// Value returns the CSS color value for this token.
func (c ColorToken) Value() css.Color { return c.value }

// WithValue returns a copy of the token with the provided CSS value.
func (c ColorToken) WithValue(value css.Color) ColorToken {
	c.value = value
	return c
}

// IsZero reports whether the token is unset.
func (c ColorToken) IsZero() bool {
	return c.suffix == "" && c.value == ""
}

// ColorScale captures the standard Tailwind 50-950 scale for a color family.
type ColorScale struct {
	C50  ColorToken
	C100 ColorToken
	C200 ColorToken
	C300 ColorToken
	C400 ColorToken
	C500 ColorToken
	C600 ColorToken
	C700 ColorToken
	C800 ColorToken
	C900 ColorToken
	C950 ColorToken
}

func (cs ColorScale) tokens() []ColorToken {
	return []ColorToken{
		cs.C50, cs.C100, cs.C200, cs.C300, cs.C400,
		cs.C500, cs.C600, cs.C700, cs.C800, cs.C900, cs.C950,
	}
}

func newColorScale(prefix string, hexes []string) ColorScale {
	if len(hexes) != 11 {
		panic("tailwind: expected 11 entries for color scale")
	}
	return ColorScale{
		C50:  newColorToken(prefix+"-50", css.Hex(hexes[0])),
		C100: newColorToken(prefix+"-100", css.Hex(hexes[1])),
		C200: newColorToken(prefix+"-200", css.Hex(hexes[2])),
		C300: newColorToken(prefix+"-300", css.Hex(hexes[3])),
		C400: newColorToken(prefix+"-400", css.Hex(hexes[4])),
		C500: newColorToken(prefix+"-500", css.Hex(hexes[5])),
		C600: newColorToken(prefix+"-600", css.Hex(hexes[6])),
		C700: newColorToken(prefix+"-700", css.Hex(hexes[7])),
		C800: newColorToken(prefix+"-800", css.Hex(hexes[8])),
		C900: newColorToken(prefix+"-900", css.Hex(hexes[9])),
		C950: newColorToken(prefix+"-950", css.Hex(hexes[10])),
	}
}

// ColorPalette captures the default Tailwind color tokens in a nested, typed structure.
type ColorPalette struct {
	Inherit     ColorToken
	Current     ColorToken
	Transparent ColorToken
	Black       ColorToken
	White       ColorToken

	Slate   ColorScale
	Gray    ColorScale
	Zinc    ColorScale
	Neutral ColorScale
	Stone   ColorScale
	Red     ColorScale
	Orange  ColorScale
	Amber   ColorScale
	Yellow  ColorScale
	Lime    ColorScale
	Green   ColorScale
	Emerald ColorScale
	Teal    ColorScale
	Cyan    ColorScale
	Sky     ColorScale
	Blue    ColorScale
	Indigo  ColorScale
	Violet  ColorScale
	Purple  ColorScale
	Fuchsia ColorScale
	Pink    ColorScale
	Rose    ColorScale
}

func (p ColorPalette) tokens() []ColorToken {
	tokens := []ColorToken{
		p.Inherit,
		p.Current,
		p.Transparent,
		p.Black,
		p.White,
	}

	families := []ColorScale{
		p.Slate,
		p.Gray,
		p.Zinc,
		p.Neutral,
		p.Stone,
		p.Red,
		p.Orange,
		p.Amber,
		p.Yellow,
		p.Lime,
		p.Green,
		p.Emerald,
		p.Teal,
		p.Cyan,
		p.Sky,
		p.Blue,
		p.Indigo,
		p.Violet,
		p.Purple,
		p.Fuchsia,
		p.Pink,
		p.Rose,
	}

	for _, family := range families {
		tokens = append(tokens, family.tokens()...)
	}

	return tokens
}

func buildDefaultColorPalette() ColorPalette {
	return ColorPalette{
		Inherit:     newColorToken("inherit", css.Color("inherit")),
		Current:     newColorToken("current", css.Color("current")),
		Transparent: newColorToken("transparent", css.Color("transparent")),
		Black:       newColorToken("black", css.Hex("#000000")),
		White:       newColorToken("white", css.Hex("#ffffff")),

		Slate: newColorScale("slate", []string{
			"#f8fafc", "#f1f5f9", "#e2e8f0", "#cbd5e1", "#94a3b8",
			"#64748b", "#475569", "#334155", "#1e293b", "#0f172a", "#020617",
		}),
		Gray: newColorScale("gray", []string{
			"#f9fafb", "#f3f4f6", "#e5e7eb", "#d1d5db", "#9ca3af",
			"#6b7280", "#4b5563", "#374151", "#1f2937", "#111827", "#030712",
		}),
		Zinc: newColorScale("zinc", []string{
			"#fafafa", "#f4f4f5", "#e4e4e7", "#d4d4d8", "#a1a1aa",
			"#71717a", "#52525b", "#3f3f46", "#27272a", "#18181b", "#09090b",
		}),
		Neutral: newColorScale("neutral", []string{
			"#fafafa", "#f5f5f5", "#e5e5e5", "#d4d4d4", "#a3a3a3",
			"#737373", "#525252", "#404040", "#262626", "#171717", "#0a0a0a",
		}),
		Stone: newColorScale("stone", []string{
			"#fafaf9", "#f5f5f4", "#e7e5e4", "#d6d3d1", "#a8a29e",
			"#78716c", "#57534e", "#44403c", "#292524", "#1c1917", "#0c0a09",
		}),
		Red: newColorScale("red", []string{
			"#fef2f2", "#fee2e2", "#fecaca", "#fca5a5", "#f87171",
			"#ef4444", "#dc2626", "#b91c1c", "#991b1b", "#7f1d1d", "#450a0a",
		}),
		Orange: newColorScale("orange", []string{
			"#fff7ed", "#ffedd5", "#fed7aa", "#fdba74", "#fb923c",
			"#f97316", "#ea580c", "#c2410c", "#9a3412", "#7c2d12", "#431407",
		}),
		Amber: newColorScale("amber", []string{
			"#fffbeb", "#fef3c7", "#fde68a", "#fcd34d", "#fbbf24",
			"#f59e0b", "#d97706", "#b45309", "#92400e", "#78350f", "#451a03",
		}),
		Yellow: newColorScale("yellow", []string{
			"#fefce8", "#fef9c3", "#fef08a", "#fde047", "#facc15",
			"#eab308", "#ca8a04", "#a16207", "#854d0e", "#713f12", "#422006",
		}),
		Lime: newColorScale("lime", []string{
			"#f7fee7", "#ecfccb", "#d9f99d", "#bef264", "#a3e635",
			"#84cc16", "#65a30d", "#4d7c0f", "#365314", "#1a2e05", "#0a1a00",
		}),
		Green: newColorScale("green", []string{
			"#f0fdf4", "#dcfce7", "#bbf7d0", "#86efac", "#4ade80",
			"#22c55e", "#16a34a", "#15803d", "#166534", "#14532d", "#052e16",
		}),
		Emerald: newColorScale("emerald", []string{
			"#ecfdf5", "#d1fae5", "#a7f3d0", "#6ee7b7", "#34d399",
			"#10b981", "#059669", "#047857", "#065f46", "#064e3b", "#022c22",
		}),
		Teal: newColorScale("teal", []string{
			"#f0fdfa", "#ccfbf1", "#99f6e4", "#5eead4", "#2dd4bf",
			"#14b8a6", "#0d9488", "#0f766e", "#115e59", "#134e4a", "#042f2e",
		}),
		Cyan: newColorScale("cyan", []string{
			"#ecfeff", "#cffafe", "#a5f3fc", "#67e8f9", "#22d3ee",
			"#06b6d4", "#0891b2", "#0e7490", "#155e75", "#164e63", "#083344",
		}),
		Sky: newColorScale("sky", []string{
			"#f0f9ff", "#e0f2fe", "#bae6fd", "#7dd3fc", "#38bdf8",
			"#0ea5e9", "#0284c7", "#0369a1", "#075985", "#0c4a6e", "#082f49",
		}),
		Blue: newColorScale("blue", []string{
			"#eff6ff", "#dbeafe", "#bfdbfe", "#93c5fd", "#60a5fa",
			"#3b82f6", "#2563eb", "#1d4ed8", "#1e40af", "#1e3a8a", "#172554",
		}),
		Indigo: newColorScale("indigo", []string{
			"#eef2ff", "#e0e7ff", "#c7d2fe", "#a5b4fc", "#818cf8",
			"#6366f1", "#4f46e5", "#4338ca", "#3730a3", "#312e81", "#1e1b4b",
		}),
		Violet: newColorScale("violet", []string{
			"#f5f3ff", "#ede9fe", "#ddd6fe", "#c4b5fd", "#a78bfa",
			"#8b5cf6", "#7c3aed", "#6d28d9", "#5b21b6", "#4c1d95", "#2e1065",
		}),
		Purple: newColorScale("purple", []string{
			"#faf5ff", "#f3e8ff", "#e9d5ff", "#d8b4fe", "#c084fc",
			"#a855f7", "#9333ea", "#7e22ce", "#6b21a8", "#581c87", "#3b0764",
		}),
		Fuchsia: newColorScale("fuchsia", []string{
			"#fdf4ff", "#fae8ff", "#f5d0fe", "#f0abfc", "#e879f9",
			"#d946ef", "#c026d3", "#a21caf", "#86198f", "#701a75", "#4a044e",
		}),
		Pink: newColorScale("pink", []string{
			"#fdf2f8", "#fce7f3", "#fbcfe8", "#f9a8d4", "#f472b6",
			"#ec4899", "#db2777", "#be185d", "#9d174d", "#831843", "#500724",
		}),
		Rose: newColorScale("rose", []string{
			"#fff1f2", "#ffe4e6", "#fecdd3", "#fda4af", "#fb7185",
			"#f43f5e", "#e11d48", "#be123c", "#9f1239", "#881337", "#4c0519",
		}),
	}
}

var defaultColorPalette = buildDefaultColorPalette()

// DefaultColorPalette returns a copy of the built-in Tailwind color palette.
func DefaultColorPalette() ColorPalette {
	return defaultColorPalette
}
