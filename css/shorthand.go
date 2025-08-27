package css

import (
	"fmt"
	"strings"
)

// Shorthand helpers for common CSS patterns

// Box model helpers

// PadXY creates padding with horizontal and vertical values
func PadXY(x, y Value) Value {
	return Raw(fmt.Sprintf("%s %s", y.String(), x.String()))
}

// PadAll creates padding with same value for all sides
func PadAll(v Value) Value {
	return v
}

// PadFull creates padding with individual values for each side (top, right, bottom, left)
func PadFull(top, right, bottom, left Value) Value {
	return Raw(fmt.Sprintf("%s %s %s %s", top.String(), right.String(), bottom.String(), left.String()))
}

// MarginXY creates margin with horizontal and vertical values
func MarginXY(x, y Value) Value {
	return Raw(fmt.Sprintf("%s %s", y.String(), x.String()))
}

// MarginAll creates margin with same value for all sides
func MarginAll(v Value) Value {
	return v
}

// MarginFull creates margin with individual values for each side (top, right, bottom, left)
func MarginFull(top, right, bottom, left Value) Value {
	return Raw(fmt.Sprintf("%s %s %s %s", top.String(), right.String(), bottom.String(), left.String()))
}

// BorderShorthand creates a border with width, style, and color
func BorderShorthand(width Value, style Keyword, color Value) Value {
	return Raw(fmt.Sprintf("%s %s %s", width.String(), style.String(), color.String()))
}

// FontShorthand creates a font declaration with size and family
func FontShorthand(size Value, family Value) Value {
	return Raw(fmt.Sprintf("%s %s", size.String(), family.String()))
}

// BackgroundShorthand creates background with color and image
func BackgroundShorthand(color Value, image Value) Value {
	if image == nil || image.String() == "" {
		return color
	}
	return Raw(fmt.Sprintf("%s %s", color.String(), image.String()))
}

// Flexbox helpers

// FlexCenter creates declarations for centered flex container
func FlexCenter() []Decl {
	return []Decl{
		Set(Display, DisplayFlex),
		Set(JustifyContent, Keyword("center")),
		Set(AlignItems, Keyword("center")),
	}
}

// FlexColumn creates declarations for column flex direction
func FlexColumn() []Decl {
	return []Decl{
		Set(Display, DisplayFlex),
		Set(FlexDirection, Keyword("column")),
	}
}

// FlexRow creates declarations for row flex direction (default, but explicit)
func FlexRow() []Decl {
	return []Decl{
		Set(Display, DisplayFlex),
		Set(FlexDirection, Keyword("row")),
	}
}

// Utility functions

// Multiple creates a space-separated list of values
func Multiple(values ...Value) Value {
	if len(values) == 0 {
		return Raw("")
	}
	
	parts := make([]string, len(values))
	for i, v := range values {
		parts[i] = v.String()
	}
	return Raw(strings.Join(parts, " "))
}

// Comma creates a comma-separated list of values
func Comma(values ...Value) Value {
	if len(values) == 0 {
		return Raw("")
	}
	
	parts := make([]string, len(values))
	for i, v := range values {
		parts[i] = v.String()
	}
	return Raw(strings.Join(parts, ", "))
}

// Common border styles
var (
	BorderSolid  = Keyword("solid")
	BorderDashed = Keyword("dashed")
	BorderDotted = Keyword("dotted")
	BorderNone   = Keyword("none")
)

// Common text alignment
var (
	TextLeft   = Keyword("left")
	TextRight  = Keyword("right")
	TextCenter = Keyword("center")
	TextJustify = Keyword("justify")
)

// Common flexbox values
var (
	JustifyStart    = Keyword("flex-start")
	JustifyEnd      = Keyword("flex-end")
	JustifyCenter   = Keyword("center")
	JustifyBetween  = Keyword("space-between")
	JustifyAround   = Keyword("space-around")
	JustifyEvenly   = Keyword("space-evenly")
	
	AlignStart    = Keyword("flex-start")
	AlignEnd      = Keyword("flex-end")
	AlignCenter   = Keyword("center")
	AlignStretch  = Keyword("stretch")
	AlignBaseline = Keyword("baseline")
)