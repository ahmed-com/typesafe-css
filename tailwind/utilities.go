// This file implements utility class generation that works with the typed configuration
// instead of string-based maps, providing full type safety throughout.

package tailwind

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ahmed-com/typesafe-css/css"
	"github.com/ahmed-com/typesafe-css/cssgen"
)

// UtilityGenerator generates CSS utility classes from the typed configuration.
type UtilityGenerator struct {
	config Config
}

// NewUtilityGenerator creates a new utility generator with the given configuration.
func NewUtilityGenerator(config Config) *UtilityGenerator {
	return &UtilityGenerator{config: config}
}

// NewDefaultUtilityGenerator creates a utility generator with the default configuration.
func NewDefaultUtilityGenerator() *UtilityGenerator {
	return NewUtilityGenerator(DefaultConfig())
}

// GenerateUtilities generates all utility classes for a given category.
func (g *UtilityGenerator) GenerateUtilities() css.Stylesheet {
	var stylesheet css.Stylesheet

	// Generate color utilities
	g.generateColorUtilities(&stylesheet)

	// Generate spacing utilities
	g.generateSpacingUtilities(&stylesheet)

	// Generate typography utilities
	g.generateTypographyUtilities(&stylesheet)

	// Generate layout utilities
	g.generateLayoutUtilities(&stylesheet)

	// Generate effect utilities
	g.generateEffectUtilities(&stylesheet)

	// Generate responsive utilities
	g.generateResponsiveUtilities(&stylesheet)

	return stylesheet
}

// Color utility generation

func (g *UtilityGenerator) generateColorUtilities(stylesheet *css.Stylesheet) {
	// Background color utilities
	g.generateBackgroundColorUtilities(stylesheet)

	// Text color utilities
	g.generateTextColorUtilities(stylesheet)

	// Border color utilities
	g.generateBorderColorUtilities(stylesheet)

	// Ring color utilities
	g.generateRingColorUtilities(stylesheet)
}

func (g *UtilityGenerator) generateBackgroundColorUtilities(stylesheet *css.Stylesheet) {
	colors := g.config.Theme.Colors
	v := reflect.ValueOf(colors)
	t := reflect.TypeOf(colors)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if colorValue, ok := field.Interface().(ColorValue); ok {
			className := fmt.Sprintf("bg-%s", kebabCase(fieldType.Name))
			rule := css.RuleSet(
				"."+className,
				cssgen.SetBackgroundColor(cssgen.BackgroundColorValColorBase),
			)
			// Override with actual color value
			rule = css.RuleSet(
				"."+className,
				css.Set(cssgen.BackgroundColor, colorValue.ToCSSValue()),
			)
			stylesheet.Add(rule)
		}
	}
}

func (g *UtilityGenerator) generateTextColorUtilities(stylesheet *css.Stylesheet) {
	colors := g.config.Theme.Colors
	v := reflect.ValueOf(colors)
	t := reflect.TypeOf(colors)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if colorValue, ok := field.Interface().(ColorValue); ok {
			className := fmt.Sprintf("text-%s", kebabCase(fieldType.Name))
			rule := css.RuleSet(
				"."+className,
				css.Set(cssgen.Color, colorValue.ToCSSValue()),
			)
			stylesheet.Add(rule)
		}
	}
}

func (g *UtilityGenerator) generateBorderColorUtilities(stylesheet *css.Stylesheet) {
	colors := g.config.Theme.Colors
	v := reflect.ValueOf(colors)
	t := reflect.TypeOf(colors)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if colorValue, ok := field.Interface().(ColorValue); ok {
			className := fmt.Sprintf("border-%s", kebabCase(fieldType.Name))
			rule := css.RuleSet(
				"."+className,
				css.Set(cssgen.BorderColor, colorValue.ToCSSValue()),
			)
			stylesheet.Add(rule)
		}
	}
}

func (g *UtilityGenerator) generateRingColorUtilities(stylesheet *css.Stylesheet) {
	colors := g.config.Theme.Colors
	v := reflect.ValueOf(colors)
	t := reflect.TypeOf(colors)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if colorValue, ok := field.Interface().(ColorValue); ok {
			className := fmt.Sprintf("ring-%s", kebabCase(fieldType.Name))
			rule := css.RuleSet(
				"."+className,
				css.Set(css.Property("--tw-ring-color"), colorValue.ToCSSValue()),
			)
			stylesheet.Add(rule)
		}
	}
}

// Spacing utility generation

func (g *UtilityGenerator) generateSpacingUtilities(stylesheet *css.Stylesheet) {
	// Padding utilities
	g.generatePaddingUtilities(stylesheet)

	// Margin utilities
	g.generateMarginUtilities(stylesheet)

	// Gap utilities
	g.generateGapUtilities(stylesheet)
}

func (g *UtilityGenerator) generatePaddingUtilities(stylesheet *css.Stylesheet) {
	spacing := g.config.Theme.Spacing
	v := reflect.ValueOf(spacing)
	t := reflect.TypeOf(spacing)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if lengthValue, ok := field.Interface().(LengthValue); ok {
			suffix := getSpacingSuffix(fieldType.Name)

			// All padding: p-*
			stylesheet.Add(css.RuleSet(
				fmt.Sprintf(".p-%s", suffix),
				css.Set(cssgen.Padding, lengthValue.ToCSSValue()),
			))

			// Horizontal padding: px-*
			stylesheet.Add(css.RuleSet(
				fmt.Sprintf(".px-%s", suffix),
				css.Set(cssgen.PaddingLeft, lengthValue.ToCSSValue()),
				css.Set(cssgen.PaddingRight, lengthValue.ToCSSValue()),
			))

			// Vertical padding: py-*
			stylesheet.Add(css.RuleSet(
				fmt.Sprintf(".py-%s", suffix),
				css.Set(cssgen.PaddingTop, lengthValue.ToCSSValue()),
				css.Set(cssgen.PaddingBottom, lengthValue.ToCSSValue()),
			))

			// Individual sides
			stylesheet.Add(css.RuleSet(
				fmt.Sprintf(".pt-%s", suffix),
				css.Set(cssgen.PaddingTop, lengthValue.ToCSSValue()),
			))
			stylesheet.Add(css.RuleSet(
				fmt.Sprintf(".pr-%s", suffix),
				css.Set(cssgen.PaddingRight, lengthValue.ToCSSValue()),
			))
			stylesheet.Add(css.RuleSet(
				fmt.Sprintf(".pb-%s", suffix),
				css.Set(cssgen.PaddingBottom, lengthValue.ToCSSValue()),
			))
			stylesheet.Add(css.RuleSet(
				fmt.Sprintf(".pl-%s", suffix),
				css.Set(cssgen.PaddingLeft, lengthValue.ToCSSValue()),
			))
		}
	}
}

func (g *UtilityGenerator) generateMarginUtilities(stylesheet *css.Stylesheet) {
	spacing := g.config.Theme.Spacing
	v := reflect.ValueOf(spacing)
	t := reflect.TypeOf(spacing)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if lengthValue, ok := field.Interface().(LengthValue); ok {
			suffix := getSpacingSuffix(fieldType.Name)

			// All margin: m-*
			stylesheet.Add(css.RuleSet(
				fmt.Sprintf(".m-%s", suffix),
				css.Set(cssgen.Margin, lengthValue.ToCSSValue()),
			))

			// Horizontal margin: mx-*
			stylesheet.Add(css.RuleSet(
				fmt.Sprintf(".mx-%s", suffix),
				css.Set(cssgen.MarginLeft, lengthValue.ToCSSValue()),
				css.Set(cssgen.MarginRight, lengthValue.ToCSSValue()),
			))

			// Vertical margin: my-*
			stylesheet.Add(css.RuleSet(
				fmt.Sprintf(".my-%s", suffix),
				css.Set(cssgen.MarginTop, lengthValue.ToCSSValue()),
				css.Set(cssgen.MarginBottom, lengthValue.ToCSSValue()),
			))

			// Individual sides
			stylesheet.Add(css.RuleSet(
				fmt.Sprintf(".mt-%s", suffix),
				css.Set(cssgen.MarginTop, lengthValue.ToCSSValue()),
			))
			stylesheet.Add(css.RuleSet(
				fmt.Sprintf(".mr-%s", suffix),
				css.Set(cssgen.MarginRight, lengthValue.ToCSSValue()),
			))
			stylesheet.Add(css.RuleSet(
				fmt.Sprintf(".mb-%s", suffix),
				css.Set(cssgen.MarginBottom, lengthValue.ToCSSValue()),
			))
			stylesheet.Add(css.RuleSet(
				fmt.Sprintf(".ml-%s", suffix),
				css.Set(cssgen.MarginLeft, lengthValue.ToCSSValue()),
			))
		}
	}
}

func (g *UtilityGenerator) generateGapUtilities(stylesheet *css.Stylesheet) {
	spacing := g.config.Theme.Spacing
	v := reflect.ValueOf(spacing)
	t := reflect.TypeOf(spacing)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if lengthValue, ok := field.Interface().(LengthValue); ok {
			suffix := getSpacingSuffix(fieldType.Name)

			// Gap: gap-*
			stylesheet.Add(css.RuleSet(
				fmt.Sprintf(".gap-%s", suffix),
				css.Set(cssgen.Gap, lengthValue.ToCSSValue()),
			))

			// Column gap: gap-x-*
			stylesheet.Add(css.RuleSet(
				fmt.Sprintf(".gap-x-%s", suffix),
				css.Set(cssgen.ColumnGap, lengthValue.ToCSSValue()),
			))

			// Row gap: gap-y-*
			stylesheet.Add(css.RuleSet(
				fmt.Sprintf(".gap-y-%s", suffix),
				css.Set(cssgen.RowGap, lengthValue.ToCSSValue()),
			))
		}
	}
}

// Typography utility generation

func (g *UtilityGenerator) generateTypographyUtilities(stylesheet *css.Stylesheet) {
	// Font size utilities
	g.generateFontSizeUtilities(stylesheet)

	// Font weight utilities
	g.generateFontWeightUtilities(stylesheet)

	// Font family utilities
	g.generateFontFamilyUtilities(stylesheet)
}

func (g *UtilityGenerator) generateFontSizeUtilities(stylesheet *css.Stylesheet) {
	fontSizes := g.config.Theme.FontSize
	v := reflect.ValueOf(fontSizes)
	t := reflect.TypeOf(fontSizes)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if fontSizeValue, ok := field.Interface().(FontSizeValue); ok {
			suffix := getFontSizeSuffix(fieldType.Name)
			className := fmt.Sprintf("text-%s", suffix)

			var decls []css.Decl
			decls = append(decls, css.Set(cssgen.FontSize, fontSizeValue.Size.ToCSSValue()))

			if fontSizeValue.LineHeight != nil {
				decls = append(decls, css.Set(cssgen.LineHeight, fontSizeValue.LineHeight.ToCSSValue()))
			}

			rule := css.RuleSet("."+className, decls...)
			stylesheet.Add(rule)
		}
	}
}

func (g *UtilityGenerator) generateFontWeightUtilities(stylesheet *css.Stylesheet) {
	fontWeights := g.config.Theme.FontWeight
	v := reflect.ValueOf(fontWeights)
	t := reflect.TypeOf(fontWeights)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if keywordValue, ok := field.Interface().(KeywordValue); ok {
			suffix := kebabCase(fieldType.Name)
			className := fmt.Sprintf("font-%s", suffix)

			rule := css.RuleSet(
				"."+className,
				css.Set(cssgen.FontWeight, keywordValue.ToCSSValue()),
			)
			stylesheet.Add(rule)
		}
	}
}

func (g *UtilityGenerator) generateFontFamilyUtilities(stylesheet *css.Stylesheet) {
	fontFamilies := g.config.Theme.FontFamily

	// Font sans
	rule := css.RuleSet(
		".font-sans",
		css.Set(cssgen.FontFamily, css.Raw(strings.Join(fontFamilies.Sans, ", "))),
	)
	stylesheet.Add(rule)

	// Font serif
	rule = css.RuleSet(
		".font-serif",
		css.Set(cssgen.FontFamily, css.Raw(strings.Join(fontFamilies.Serif, ", "))),
	)
	stylesheet.Add(rule)

	// Font mono
	rule = css.RuleSet(
		".font-mono",
		css.Set(cssgen.FontFamily, css.Raw(strings.Join(fontFamilies.Mono, ", "))),
	)
	stylesheet.Add(rule)
}

// Layout utility generation

func (g *UtilityGenerator) generateLayoutUtilities(stylesheet *css.Stylesheet) {
	// Display utilities
	stylesheet.Add(css.RuleSet(".block", cssgen.SetDisplay(cssgen.DisplayValBlock)))
	stylesheet.Add(css.RuleSet(".inline", cssgen.SetDisplay(cssgen.DisplayValInline)))
	stylesheet.Add(css.RuleSet(".inline-block", cssgen.SetDisplay(cssgen.DisplayValInlineBlock)))
	stylesheet.Add(css.RuleSet(".flex", cssgen.SetDisplay(cssgen.DisplayValFlex)))
	stylesheet.Add(css.RuleSet(".inline-flex", cssgen.SetDisplay(cssgen.DisplayValInlineFlex)))
	stylesheet.Add(css.RuleSet(".grid", cssgen.SetDisplay(cssgen.DisplayValGrid)))
	stylesheet.Add(css.RuleSet(".inline-grid", cssgen.SetDisplay(cssgen.DisplayValInlineGrid)))
	stylesheet.Add(css.RuleSet(".contents", cssgen.SetDisplay(cssgen.DisplayValContents)))
	stylesheet.Add(css.RuleSet(".list-item", cssgen.SetDisplay(cssgen.DisplayValListItem)))
	stylesheet.Add(css.RuleSet(".hidden", cssgen.SetDisplay(cssgen.DisplayValNone)))

	// Position utilities
	stylesheet.Add(css.RuleSet(".static", cssgen.SetPosition(cssgen.PositionValStatic)))
	stylesheet.Add(css.RuleSet(".fixed", cssgen.SetPosition(cssgen.PositionValFixed)))
	stylesheet.Add(css.RuleSet(".absolute", cssgen.SetPosition(cssgen.PositionValAbsolute)))
	stylesheet.Add(css.RuleSet(".relative", cssgen.SetPosition(cssgen.PositionValRelative)))
	stylesheet.Add(css.RuleSet(".sticky", cssgen.SetPosition(cssgen.PositionValSticky)))

	// Flexbox utilities
	stylesheet.Add(css.RuleSet(".flex-row", cssgen.SetFlexDirection(cssgen.FlexDirectionValRow)))
	stylesheet.Add(css.RuleSet(".flex-row-reverse", cssgen.SetFlexDirection(cssgen.FlexDirectionValRowReverse)))
	stylesheet.Add(css.RuleSet(".flex-col", cssgen.SetFlexDirection(cssgen.FlexDirectionValColumn)))
	stylesheet.Add(css.RuleSet(".flex-col-reverse", cssgen.SetFlexDirection(cssgen.FlexDirectionValColumnReverse)))

	stylesheet.Add(css.RuleSet(".flex-wrap", cssgen.SetFlexWrap(cssgen.FlexWrapValWrap)))
	stylesheet.Add(css.RuleSet(".flex-wrap-reverse", cssgen.SetFlexWrap(cssgen.FlexWrapValWrapReverse)))
	stylesheet.Add(css.RuleSet(".flex-nowrap", cssgen.SetFlexWrap(cssgen.FlexWrapValNowrap)))

	stylesheet.Add(css.RuleSet(".justify-start", cssgen.SetJustifyContent(cssgen.JustifyContentValFlexStart)))
	stylesheet.Add(css.RuleSet(".justify-end", cssgen.SetJustifyContent(cssgen.JustifyContentValFlexEnd)))
	stylesheet.Add(css.RuleSet(".justify-center", cssgen.SetJustifyContent(cssgen.JustifyContentValCenter)))
	stylesheet.Add(css.RuleSet(".justify-between", cssgen.SetJustifyContent(cssgen.JustifyContentValSpaceBetween)))
	stylesheet.Add(css.RuleSet(".justify-around", cssgen.SetJustifyContent(cssgen.JustifyContentValSpaceAround)))
	stylesheet.Add(css.RuleSet(".justify-evenly", cssgen.SetJustifyContent(cssgen.JustifyContentValSpaceEvenly)))

	stylesheet.Add(css.RuleSet(".items-start", cssgen.SetAlignItems(cssgen.AlignItemsValFlexStart)))
	stylesheet.Add(css.RuleSet(".items-end", cssgen.SetAlignItems(cssgen.AlignItemsValFlexEnd)))
	stylesheet.Add(css.RuleSet(".items-center", cssgen.SetAlignItems(cssgen.AlignItemsValCenter)))
	stylesheet.Add(css.RuleSet(".items-baseline", cssgen.SetAlignItems(cssgen.AlignItemsValBaseline)))
	stylesheet.Add(css.RuleSet(".items-stretch", cssgen.SetAlignItems(cssgen.AlignItemsValStretch)))
}

// Effect utility generation

func (g *UtilityGenerator) generateEffectUtilities(stylesheet *css.Stylesheet) {
	// Opacity utilities
	g.generateOpacityUtilities(stylesheet)

	// Blur utilities
	g.generateBlurUtilities(stylesheet)

	// Brightness utilities
	g.generateBrightnessUtilities(stylesheet)
}

func (g *UtilityGenerator) generateOpacityUtilities(stylesheet *css.Stylesheet) {
	opacity := g.config.Theme.Opacity
	v := reflect.ValueOf(opacity)
	t := reflect.TypeOf(opacity)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if numberValue, ok := field.Interface().(NumberValue); ok {
			suffix := getOpacitySuffix(fieldType.Name)
			className := fmt.Sprintf("opacity-%s", suffix)

			rule := css.RuleSet(
				"."+className,
				css.Set(cssgen.Opacity, numberValue.ToCSSValue()),
			)
			stylesheet.Add(rule)
		}
	}
}

func (g *UtilityGenerator) generateBlurUtilities(stylesheet *css.Stylesheet) {
	blur := g.config.Theme.Blur
	v := reflect.ValueOf(blur)
	t := reflect.TypeOf(blur)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		suffix := getBlurSuffix(fieldType.Name)
		className := fmt.Sprintf("blur-%s", suffix)

		if lengthValue, ok := field.Interface().(LengthValue); ok {
			rule := css.RuleSet(
				"."+className,
				css.Set(cssgen.Filter, css.Raw(fmt.Sprintf("blur(%s)", lengthValue.ToCSSValue().String()))),
			)
			stylesheet.Add(rule)
		} else if strValue, ok := field.Interface().(string); ok {
			if strValue == "" { // none case
				rule := css.RuleSet(
					"."+className,
					css.Set(cssgen.Filter, css.Keyword("blur(0)")),
				)
				stylesheet.Add(rule)
			}
		}
	}
}

func (g *UtilityGenerator) generateBrightnessUtilities(stylesheet *css.Stylesheet) {
	brightness := g.config.Theme.Brightness
	v := reflect.ValueOf(brightness)
	t := reflect.TypeOf(brightness)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if numberValue, ok := field.Interface().(NumberValue); ok {
			suffix := getBrightnessSuffix(fieldType.Name)
			className := fmt.Sprintf("brightness-%s", suffix)

			rule := css.RuleSet(
				"."+className,
				css.Set(cssgen.Filter, css.Raw(fmt.Sprintf("brightness(%s)", numberValue.ToCSSValue().String()))),
			)
			stylesheet.Add(rule)
		}
	}
}

// Responsive utility generation

func (g *UtilityGenerator) generateResponsiveUtilities(stylesheet *css.Stylesheet) {
	screens := g.config.Theme.Screens
	v := reflect.ValueOf(screens)
	t := reflect.TypeOf(screens)

	// For each breakpoint, we would generate responsive variants
	// This is a simplified example - full implementation would be more complex
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if lengthValue, ok := field.Interface().(LengthValue); ok {
			screenName := getScreenName(fieldType.Name)

			// Example: @media (min-width: 640px) { .sm\:block { display: block; } }
			mediaRule := css.AtRule{
				Name:   "media",
				Params: fmt.Sprintf("(min-width: %s)", lengthValue.ToCSSValue().String()),
				Body: []css.Item{
					css.RuleSet(fmt.Sprintf(".%s\\:block", screenName), cssgen.SetDisplay(cssgen.DisplayValBlock)),
					css.RuleSet(fmt.Sprintf(".%s\\:flex", screenName), cssgen.SetDisplay(cssgen.DisplayValFlex)),
					css.RuleSet(fmt.Sprintf(".%s\\:hidden", screenName), cssgen.SetDisplay(cssgen.DisplayValNone)),
				},
			}
			stylesheet.Add(mediaRule)
		}
	}
}

// Helper functions for naming conversions

func kebabCase(s string) string {
	// Convert CamelCase to kebab-case
	// This is a simplified implementation
	result := ""
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result += "-"
		}
		result += string(r)
	}
	return result
}

func getSpacingSuffix(fieldName string) string {
	// Convert field names like "Size4" to "4", "Px" to "px", etc.
	if fieldName == "Size0" {
		return "0"
	}
	if fieldName == "Size0_5" {
		return "0.5"
	}
	if fieldName == "Size1_5" {
		return "1.5"
	}
	if fieldName == "Size2_5" {
		return "2.5"
	}
	if fieldName == "Size3_5" {
		return "3.5"
	}
	if strings.HasPrefix(fieldName, "Size") {
		return fieldName[4:] // Remove "Size" prefix
	}
	if fieldName == "Px" {
		return "px"
	}
	return fieldName
}

func getFontSizeSuffix(fieldName string) string {
	// Handle Size* field names like Size2xl -> 2xl
	if strings.HasPrefix(fieldName, "Size") {
		return strings.ToLower(fieldName[4:]) // Remove "Size" and convert to lowercase
	}
	return kebabCase(fieldName)
}

func getOpacitySuffix(fieldName string) string {
	// Handle Size* field names like Size50 -> 50
	if strings.HasPrefix(fieldName, "Size") {
		return fieldName[4:] // Remove "Size" prefix
	}
	return fieldName
}

func getBlurSuffix(fieldName string) string {
	// Handle Size* field names like Size2xl -> 2xl
	if strings.HasPrefix(fieldName, "Size") {
		return strings.ToLower(fieldName[4:]) // Remove "Size" and convert to lowercase
	}
	if fieldName == "DEFAULT" {
		return ""
	}
	return kebabCase(fieldName)
}

func getBrightnessSuffix(fieldName string) string {
	// Handle Size* field names like Size100 -> 100
	if strings.HasPrefix(fieldName, "Size") {
		return fieldName[4:] // Remove "Size" prefix
	}
	return fieldName
}

func getScreenName(fieldName string) string {
	// Handle Size* field names like Size2xl -> 2xl
	if strings.HasPrefix(fieldName, "Size") {
		return strings.ToLower(fieldName[4:]) // Remove "Size" and convert to lowercase
	}
	return kebabCase(fieldName)
}
