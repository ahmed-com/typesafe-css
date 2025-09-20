// Package tailwind provides Tailwind CSS-style utility classes for type-safe CSS generation.
package tailwind

import "github.com/ahmed-com/typesafe-css/css"

// Theme defines the design tokens used for generating utility classes.
type Theme struct {
	Colors       ColorPalette
	colorIndex   map[string]ColorToken
	customColors map[string]ColorToken

	Spacing       SpacingScale
	spacingIndex  map[string]ValueToken
	customSpacing map[string]ValueToken

	FontSizes       FontSizeScale
	fontSizeIndex   map[string]ValueToken
	customFontSizes map[string]ValueToken

	Screens Screens

	BorderRadius       BorderRadiusScale
	borderRadiusIndex  map[string]ValueToken
	customBorderRadius map[string]ValueToken

	BorderWidth       BorderWidthScale
	borderWidthIndex  map[string]ValueToken
	customBorderWidth map[string]ValueToken

	BoxShadow       BoxShadowScale
	boxShadowIndex  map[string]ValueToken
	customBoxShadow map[string]ValueToken

	Opacity       OpacityScale
	opacityIndex  map[string]ValueToken
	customOpacity map[string]ValueToken

	FontWeight       FontWeightScale
	fontWeightIndex  map[string]ValueToken
	customFontWeight map[string]ValueToken

	LineHeight       LineHeightScale
	lineHeightIndex  map[string]ValueToken
	customLineHeight map[string]ValueToken

	ZIndex       ZIndexScale
	zIndexIndex  map[string]ValueToken
	customZIndex map[string]ValueToken

	Blur       BlurScale
	blurIndex  map[string]ValueToken
	customBlur map[string]ValueToken

	Brightness       BrightnessScale
	brightnessIndex  map[string]ValueToken
	customBrightness map[string]ValueToken

	Contrast       ContrastScale
	contrastIndex  map[string]ValueToken
	customContrast map[string]ValueToken

	Grayscale       PercentageScale
	grayscaleIndex  map[string]ValueToken
	customGrayscale map[string]ValueToken

	Invert       PercentageScale
	invertIndex  map[string]ValueToken
	customInvert map[string]ValueToken

	Saturate       SaturateScale
	saturateIndex  map[string]ValueToken
	customSaturate map[string]ValueToken

	Sepia       PercentageScale
	sepiaIndex  map[string]ValueToken
	customSepia map[string]ValueToken
}

// Rebuild recomputes derived indices after mutating theme data.
func (t *Theme) Rebuild() {
	t.rebuildColorIndex()
	t.rebuildSpacingIndex()
	t.rebuildFontSizeIndex()
	t.rebuildBorderRadiusIndex()
	t.rebuildBorderWidthIndex()
	t.rebuildBoxShadowIndex()
	t.rebuildOpacityIndex()
	t.rebuildFontWeightIndex()
	t.rebuildLineHeightIndex()
	t.rebuildZIndexIndex()
	t.rebuildBlurIndex()
	t.rebuildBrightnessIndex()
	t.rebuildContrastIndex()
	t.rebuildGrayscaleIndex()
	t.rebuildInvertIndex()
	t.rebuildSaturateIndex()
	t.rebuildSepiaIndex()
}

func makeValueIndex(base []ValueToken, custom map[string]ValueToken) map[string]ValueToken {
	index := make(map[string]ValueToken, len(base)+len(custom))
	for _, token := range base {
		index[token.Suffix()] = token
	}
	for name, token := range custom {
		index[name] = token
	}
	return index
}

func (t *Theme) rebuildColorIndex() {
	index := make(map[string]ColorToken)
	for _, token := range t.Colors.tokens() {
		index[token.Suffix()] = token
	}
	for name, token := range t.customColors {
		index[name] = token
	}
	t.colorIndex = index
}

func (t *Theme) rebuildSpacingIndex() {
	if t.customSpacing == nil {
		t.customSpacing = make(map[string]ValueToken)
	}
	t.spacingIndex = makeValueIndex(t.Spacing.tokens(), t.customSpacing)
}

func (t *Theme) rebuildFontSizeIndex() {
	if t.customFontSizes == nil {
		t.customFontSizes = make(map[string]ValueToken)
	}
	t.fontSizeIndex = makeValueIndex(t.FontSizes.tokens(), t.customFontSizes)
}

func (t *Theme) rebuildBorderRadiusIndex() {
	if t.customBorderRadius == nil {
		t.customBorderRadius = make(map[string]ValueToken)
	}
	t.borderRadiusIndex = makeValueIndex(t.BorderRadius.tokens(), t.customBorderRadius)
}

func (t *Theme) rebuildBorderWidthIndex() {
	if t.customBorderWidth == nil {
		t.customBorderWidth = make(map[string]ValueToken)
	}
	t.borderWidthIndex = makeValueIndex(t.BorderWidth.tokens(), t.customBorderWidth)
}

func (t *Theme) rebuildBoxShadowIndex() {
	if t.customBoxShadow == nil {
		t.customBoxShadow = make(map[string]ValueToken)
	}
	t.boxShadowIndex = makeValueIndex(t.BoxShadow.tokens(), t.customBoxShadow)
}

func (t *Theme) rebuildOpacityIndex() {
	if t.customOpacity == nil {
		t.customOpacity = make(map[string]ValueToken)
	}
	t.opacityIndex = makeValueIndex(t.Opacity.tokens(), t.customOpacity)
}

func (t *Theme) rebuildFontWeightIndex() {
	if t.customFontWeight == nil {
		t.customFontWeight = make(map[string]ValueToken)
	}
	t.fontWeightIndex = makeValueIndex(t.FontWeight.tokens(), t.customFontWeight)
}

func (t *Theme) rebuildLineHeightIndex() {
	if t.customLineHeight == nil {
		t.customLineHeight = make(map[string]ValueToken)
	}
	t.lineHeightIndex = makeValueIndex(t.LineHeight.tokens(), t.customLineHeight)
}

func (t *Theme) rebuildZIndexIndex() {
	if t.customZIndex == nil {
		t.customZIndex = make(map[string]ValueToken)
	}
	t.zIndexIndex = makeValueIndex(t.ZIndex.tokens(), t.customZIndex)
}

func (t *Theme) rebuildBlurIndex() {
	if t.customBlur == nil {
		t.customBlur = make(map[string]ValueToken)
	}
	t.blurIndex = makeValueIndex(t.Blur.tokens(), t.customBlur)
}

func (t *Theme) rebuildBrightnessIndex() {
	if t.customBrightness == nil {
		t.customBrightness = make(map[string]ValueToken)
	}
	t.brightnessIndex = makeValueIndex(t.Brightness.tokens(), t.customBrightness)
}

func (t *Theme) rebuildContrastIndex() {
	if t.customContrast == nil {
		t.customContrast = make(map[string]ValueToken)
	}
	t.contrastIndex = makeValueIndex(t.Contrast.tokens(), t.customContrast)
}

func (t *Theme) rebuildGrayscaleIndex() {
	if t.customGrayscale == nil {
		t.customGrayscale = make(map[string]ValueToken)
	}
	t.grayscaleIndex = makeValueIndex(t.Grayscale.tokens(), t.customGrayscale)
}

func (t *Theme) rebuildInvertIndex() {
	if t.customInvert == nil {
		t.customInvert = make(map[string]ValueToken)
	}
	t.invertIndex = makeValueIndex(t.Invert.tokens(), t.customInvert)
}

func (t *Theme) rebuildSaturateIndex() {
	if t.customSaturate == nil {
		t.customSaturate = make(map[string]ValueToken)
	}
	t.saturateIndex = makeValueIndex(t.Saturate.tokens(), t.customSaturate)
}

func (t *Theme) rebuildSepiaIndex() {
	if t.customSepia == nil {
		t.customSepia = make(map[string]ValueToken)
	}
	t.sepiaIndex = makeValueIndex(t.Sepia.tokens(), t.customSepia)
}

// ColorByName looks up a color token by its Tailwind suffix.
func (t *Theme) ColorByName(name string) (ColorToken, bool) {
	if t.colorIndex == nil {
		t.rebuildColorIndex()
	}
	token, ok := t.colorIndex[name]
	return token, ok
}

// AddColor registers or overrides a custom color token on the theme.
func (t *Theme) AddColor(name string, value css.Color) {
	if t.customColors == nil {
		t.customColors = make(map[string]ColorToken)
	}
	t.customColors[name] = newColorToken(name, value)
	t.rebuildColorIndex()
}

// SetColorToken registers or overrides a color token using the token's suffix.
func (t *Theme) SetColorToken(token ColorToken) {
	if token.Suffix() == "" {
		return
	}
	t.AddColor(token.Suffix(), token.Value())
}

// ColorCount returns the number of registered color tokens, including custom entries.
func (t *Theme) ColorCount() int {
	if t.colorIndex == nil {
		t.rebuildColorIndex()
	}
	return len(t.colorIndex)
}

// SpacingToken returns the spacing token for the given suffix.
func (t *Theme) SpacingToken(name string) (ValueToken, bool) {
	if t.spacingIndex == nil {
		t.rebuildSpacingIndex()
	}
	token, ok := t.spacingIndex[name]
	return token, ok
}

// AddSpacing registers or overrides a spacing token.
func (t *Theme) AddSpacing(name string, value css.Value) {
	if t.customSpacing == nil {
		t.customSpacing = make(map[string]ValueToken)
	}
	t.customSpacing[name] = newValueToken(name, value)
	t.rebuildSpacingIndex()
}

// FontSizeToken returns the font-size token for the given suffix.
func (t *Theme) FontSizeToken(name string) (ValueToken, bool) {
	if t.fontSizeIndex == nil {
		t.rebuildFontSizeIndex()
	}
	token, ok := t.fontSizeIndex[name]
	return token, ok
}

// AddFontSize registers or overrides a font-size token.
func (t *Theme) AddFontSize(name string, value css.Value) {
	if t.customFontSizes == nil {
		t.customFontSizes = make(map[string]ValueToken)
	}
	t.customFontSizes[name] = newValueToken(name, value)
	t.rebuildFontSizeIndex()
}

// BorderRadiusToken returns the border-radius token for the given suffix.
func (t *Theme) BorderRadiusToken(name string) (ValueToken, bool) {
	if t.borderRadiusIndex == nil {
		t.rebuildBorderRadiusIndex()
	}
	token, ok := t.borderRadiusIndex[name]
	return token, ok
}

// AddBorderRadius registers or overrides a border-radius token.
func (t *Theme) AddBorderRadius(name string, value css.Value) {
	if t.customBorderRadius == nil {
		t.customBorderRadius = make(map[string]ValueToken)
	}
	t.customBorderRadius[name] = newValueToken(name, value)
	t.rebuildBorderRadiusIndex()
}

// BorderWidthToken returns the border-width token for the given suffix.
func (t *Theme) BorderWidthToken(name string) (ValueToken, bool) {
	if t.borderWidthIndex == nil {
		t.rebuildBorderWidthIndex()
	}
	token, ok := t.borderWidthIndex[name]
	return token, ok
}

// AddBorderWidth registers or overrides a border-width token.
func (t *Theme) AddBorderWidth(name string, value css.Value) {
	if t.customBorderWidth == nil {
		t.customBorderWidth = make(map[string]ValueToken)
	}
	t.customBorderWidth[name] = newValueToken(name, value)
	t.rebuildBorderWidthIndex()
}

// BoxShadowToken returns the box-shadow token for the given suffix.
func (t *Theme) BoxShadowToken(name string) (ValueToken, bool) {
	if t.boxShadowIndex == nil {
		t.rebuildBoxShadowIndex()
	}
	token, ok := t.boxShadowIndex[name]
	return token, ok
}

// AddBoxShadow registers or overrides a box-shadow token.
func (t *Theme) AddBoxShadow(name string, value css.Value) {
	if t.customBoxShadow == nil {
		t.customBoxShadow = make(map[string]ValueToken)
	}
	t.customBoxShadow[name] = newValueToken(name, value)
	t.rebuildBoxShadowIndex()
}

// OpacityToken returns the opacity token for the given suffix.
func (t *Theme) OpacityToken(name string) (ValueToken, bool) {
	if t.opacityIndex == nil {
		t.rebuildOpacityIndex()
	}
	token, ok := t.opacityIndex[name]
	return token, ok
}

// AddOpacity registers or overrides an opacity token.
func (t *Theme) AddOpacity(name string, value css.Value) {
	if t.customOpacity == nil {
		t.customOpacity = make(map[string]ValueToken)
	}
	t.customOpacity[name] = newValueToken(name, value)
	t.rebuildOpacityIndex()
}

// FontWeightToken returns the font-weight token for the given suffix.
func (t *Theme) FontWeightToken(name string) (ValueToken, bool) {
	if t.fontWeightIndex == nil {
		t.rebuildFontWeightIndex()
	}
	token, ok := t.fontWeightIndex[name]
	return token, ok
}

// AddFontWeight registers or overrides a font-weight token.
func (t *Theme) AddFontWeight(name string, value css.Value) {
	if t.customFontWeight == nil {
		t.customFontWeight = make(map[string]ValueToken)
	}
	t.customFontWeight[name] = newValueToken(name, value)
	t.rebuildFontWeightIndex()
}

// LineHeightToken returns the line-height token for the given suffix.
func (t *Theme) LineHeightToken(name string) (ValueToken, bool) {
	if t.lineHeightIndex == nil {
		t.rebuildLineHeightIndex()
	}
	token, ok := t.lineHeightIndex[name]
	return token, ok
}

// AddLineHeight registers or overrides a line-height token.
func (t *Theme) AddLineHeight(name string, value css.Value) {
	if t.customLineHeight == nil {
		t.customLineHeight = make(map[string]ValueToken)
	}
	t.customLineHeight[name] = newValueToken(name, value)
	t.rebuildLineHeightIndex()
}

// ZIndexToken returns the z-index token for the given suffix.
func (t *Theme) ZIndexToken(name string) (ValueToken, bool) {
	if t.zIndexIndex == nil {
		t.rebuildZIndexIndex()
	}
	token, ok := t.zIndexIndex[name]
	return token, ok
}

// AddZIndex registers or overrides a z-index token.
func (t *Theme) AddZIndex(name string, value css.Value) {
	if t.customZIndex == nil {
		t.customZIndex = make(map[string]ValueToken)
	}
	t.customZIndex[name] = newValueToken(name, value)
	t.rebuildZIndexIndex()
}

// BlurToken returns the blur token for the given suffix.
func (t *Theme) BlurToken(name string) (ValueToken, bool) {
	if t.blurIndex == nil {
		t.rebuildBlurIndex()
	}
	token, ok := t.blurIndex[name]
	return token, ok
}

// AddBlur registers or overrides a blur token.
func (t *Theme) AddBlur(name string, value css.Value) {
	if t.customBlur == nil {
		t.customBlur = make(map[string]ValueToken)
	}
	t.customBlur[name] = newValueToken(name, value)
	t.rebuildBlurIndex()
}

// BrightnessToken returns the brightness token for the given suffix.
func (t *Theme) BrightnessToken(name string) (ValueToken, bool) {
	if t.brightnessIndex == nil {
		t.rebuildBrightnessIndex()
	}
	token, ok := t.brightnessIndex[name]
	return token, ok
}

// AddBrightness registers or overrides a brightness token.
func (t *Theme) AddBrightness(name string, value css.Value) {
	if t.customBrightness == nil {
		t.customBrightness = make(map[string]ValueToken)
	}
	t.customBrightness[name] = newValueToken(name, value)
	t.rebuildBrightnessIndex()
}

// ContrastToken returns the contrast token for the given suffix.
func (t *Theme) ContrastToken(name string) (ValueToken, bool) {
	if t.contrastIndex == nil {
		t.rebuildContrastIndex()
	}
	token, ok := t.contrastIndex[name]
	return token, ok
}

// AddContrast registers or overrides a contrast token.
func (t *Theme) AddContrast(name string, value css.Value) {
	if t.customContrast == nil {
		t.customContrast = make(map[string]ValueToken)
	}
	t.customContrast[name] = newValueToken(name, value)
	t.rebuildContrastIndex()
}

// GrayscaleToken returns the grayscale token for the given suffix.
func (t *Theme) GrayscaleToken(name string) (ValueToken, bool) {
	if t.grayscaleIndex == nil {
		t.rebuildGrayscaleIndex()
	}
	token, ok := t.grayscaleIndex[name]
	return token, ok
}

// AddGrayscale registers or overrides a grayscale token.
func (t *Theme) AddGrayscale(name string, value css.Value) {
	if t.customGrayscale == nil {
		t.customGrayscale = make(map[string]ValueToken)
	}
	t.customGrayscale[name] = newValueToken(name, value)
	t.rebuildGrayscaleIndex()
}

// InvertToken returns the invert token for the given suffix.
func (t *Theme) InvertToken(name string) (ValueToken, bool) {
	if t.invertIndex == nil {
		t.rebuildInvertIndex()
	}
	token, ok := t.invertIndex[name]
	return token, ok
}

// AddInvert registers or overrides an invert token.
func (t *Theme) AddInvert(name string, value css.Value) {
	if t.customInvert == nil {
		t.customInvert = make(map[string]ValueToken)
	}
	t.customInvert[name] = newValueToken(name, value)
	t.rebuildInvertIndex()
}

// SaturateToken returns the saturate token for the given suffix.
func (t *Theme) SaturateToken(name string) (ValueToken, bool) {
	if t.saturateIndex == nil {
		t.rebuildSaturateIndex()
	}
	token, ok := t.saturateIndex[name]
	return token, ok
}

// AddSaturate registers or overrides a saturate token.
func (t *Theme) AddSaturate(name string, value css.Value) {
	if t.customSaturate == nil {
		t.customSaturate = make(map[string]ValueToken)
	}
	t.customSaturate[name] = newValueToken(name, value)
	t.rebuildSaturateIndex()
}

// SepiaToken returns the sepia token for the given suffix.
func (t *Theme) SepiaToken(name string) (ValueToken, bool) {
	if t.sepiaIndex == nil {
		t.rebuildSepiaIndex()
	}
	token, ok := t.sepiaIndex[name]
	return token, ok
}

// AddSepia registers or overrides a sepia token.
func (t *Theme) AddSepia(name string, value css.Value) {
	if t.customSepia == nil {
		t.customSepia = make(map[string]ValueToken)
	}
	t.customSepia[name] = newValueToken(name, value)
	t.rebuildSepiaIndex()
}

// Count helper methods for inspecting theme configuration sizes.
func (t *Theme) SpacingCount() int {
	if t.spacingIndex == nil {
		t.rebuildSpacingIndex()
	}
	return len(t.spacingIndex)
}

func (t *Theme) FontSizeCount() int {
	if t.fontSizeIndex == nil {
		t.rebuildFontSizeIndex()
	}
	return len(t.fontSizeIndex)
}

func (t *Theme) BorderRadiusCount() int {
	if t.borderRadiusIndex == nil {
		t.rebuildBorderRadiusIndex()
	}
	return len(t.borderRadiusIndex)
}

func (t *Theme) BorderWidthCount() int {
	if t.borderWidthIndex == nil {
		t.rebuildBorderWidthIndex()
	}
	return len(t.borderWidthIndex)
}

func (t *Theme) BoxShadowCount() int {
	if t.boxShadowIndex == nil {
		t.rebuildBoxShadowIndex()
	}
	return len(t.boxShadowIndex)
}

func (t *Theme) OpacityCount() int {
	if t.opacityIndex == nil {
		t.rebuildOpacityIndex()
	}
	return len(t.opacityIndex)
}

func (t *Theme) FontWeightCount() int {
	if t.fontWeightIndex == nil {
		t.rebuildFontWeightIndex()
	}
	return len(t.fontWeightIndex)
}

func (t *Theme) LineHeightCount() int {
	if t.lineHeightIndex == nil {
		t.rebuildLineHeightIndex()
	}
	return len(t.lineHeightIndex)
}

func (t *Theme) ZIndexCount() int {
	if t.zIndexIndex == nil {
		t.rebuildZIndexIndex()
	}
	return len(t.zIndexIndex)
}

func (t *Theme) BlurCount() int {
	if t.blurIndex == nil {
		t.rebuildBlurIndex()
	}
	return len(t.blurIndex)
}

func (t *Theme) BrightnessCount() int {
	if t.brightnessIndex == nil {
		t.rebuildBrightnessIndex()
	}
	return len(t.brightnessIndex)
}

func (t *Theme) ContrastCount() int {
	if t.contrastIndex == nil {
		t.rebuildContrastIndex()
	}
	return len(t.contrastIndex)
}

func (t *Theme) GrayscaleCount() int {
	if t.grayscaleIndex == nil {
		t.rebuildGrayscaleIndex()
	}
	return len(t.grayscaleIndex)
}

func (t *Theme) InvertCount() int {
	if t.invertIndex == nil {
		t.rebuildInvertIndex()
	}
	return len(t.invertIndex)
}

func (t *Theme) SaturateCount() int {
	if t.saturateIndex == nil {
		t.rebuildSaturateIndex()
	}
	return len(t.saturateIndex)
}

func (t *Theme) SepiaCount() int {
	if t.sepiaIndex == nil {
		t.rebuildSepiaIndex()
	}
	return len(t.sepiaIndex)
}

// CustomTheme returns a theme derived from the default theme with the supplied
// mutation callbacks applied. Each modifier receives a fresh copy of the
// default theme, allowing callers to tweak nested fields without manual
// merging logic.
func CustomTheme(modifiers ...func(*Theme)) *Theme {
	// Note: This function now uses the typed DefaultTheme from defaults.go
	// TODO: This function should be updated to work with the new typed Config system
	return nil // Placeholder - needs update for typed system
}
