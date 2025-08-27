// Package css provides type-safe CSS authoring in Go.
package css

import (
	"fmt"
	"strings"
)

// Value represents any CSS value that can be converted to a string.
type Value interface {
	String() string
}

// Raw represents a raw CSS value string (catch-all for complex values).
type Raw string

func (r Raw) String() string { return string(r) }

// Keyword represents CSS keywords like "block", "flex", "auto".
type Keyword string

func (k Keyword) String() string { return string(k) }

// Length represents CSS length values with units.
type Length string

func (l Length) String() string { return string(l) }

// Color represents CSS color values.
type Color string

func (c Color) String() string { return string(c) }

// Length constructors
func Px(n int) Length {
	return Length(fmt.Sprintf("%dpx", n))
}

func Rem(x float64) Length {
	return Length(trim0(fmt.Sprintf("%.4frem", x)))
}

func Em(x float64) Length {
	return Length(trim0(fmt.Sprintf("%.4fem", x)))
}

func Percent(x float64) Length {
	return Length(trim0(fmt.Sprintf("%.6f%%", x)))
}

// Color constructors
func Hex(hex string) Color {
	return Color(hex)
}

func RGB(r, g, b uint8) Color {
	return Color(fmt.Sprintf("rgb(%d %d %d)", r, g, b))
}

func RGBA(r, g, b, a uint8) Color {
	return Color(fmt.Sprintf("rgba(%d %d %d / %.3f)", r, g, b, float64(a)/255))
}

// CSS custom property reference
func Var(name string) Raw {
	return Raw("var(" + name + ")")
}

// trim0 removes trailing zeros from decimal numbers
func trim0(s string) string {
	if !strings.Contains(s, ".") {
		return s
	}
	
	// Find the position of the decimal point
	dotIndex := strings.Index(s, ".")
	if dotIndex == -1 {
		return s
	}
	
	// Split into number part and unit part
	var numberPart, unitPart string
	for i := dotIndex + 1; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			numberPart = s[:i]
			unitPart = s[i:]
			break
		}
	}
	
	// If we didn't find a unit, the whole thing is the number
	if numberPart == "" {
		numberPart = s
	}
	
	// Trim trailing zeros from the number part
	numberPart = strings.TrimRight(numberPart, "0")
	numberPart = strings.TrimRight(numberPart, ".")
	
	return numberPart + unitPart
}

// Common keywords (hand-curated, minimal set for day-1)
var (
	// Display values
	DisplayBlock  = Keyword("block")
	DisplayFlex   = Keyword("flex")
	DisplayInline = Keyword("inline")
	DisplayNone   = Keyword("none")

	// Position values
	PositionStatic   = Keyword("static")
	PositionRelative = Keyword("relative")
	PositionAbsolute = Keyword("absolute")
	PositionFixed    = Keyword("fixed")
	PositionSticky   = Keyword("sticky")

	// Common values
	Auto   = Keyword("auto")
	None   = Keyword("none")
	Inherit = Keyword("inherit")
	Initial = Keyword("initial")
)