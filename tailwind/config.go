// Package tailwind provides type-safe Tailwind CSS configuration and utilities.
// This file defines the strongly-typed configuration structure that replaces
// string-based maps with nested properties for maximum type safety.

package tailwind

import (
	"fmt"

	"github.com/ahmed-com/typesafe-css/css"
	"github.com/ahmed-com/typesafe-css/cssgen"
)

// Config represents the complete Tailwind CSS configuration with strongly-typed
// nested properties instead of string-based maps.
type Config struct {
	// Content specifies which files to scan for class names
	Content []string

	// Presets allows extending configuration from base presets
	Presets []Config

	// DarkMode configures how dark mode is activated
	DarkMode DarkModeConfig

	// Theme contains all design tokens and their values
	Theme ThemeConfig
}

// DarkModeConfig specifies how dark mode variants are handled.
type DarkModeConfig struct {
	// Strategy determines how dark mode is activated
	Strategy DarkModeStrategy

	// Selector is the CSS selector for manual dark mode (when Strategy is Selector)
	Selector string
}

// DarkModeStrategy defines how dark mode is activated.
type DarkModeStrategy int

const (
	// DarkModeMedia uses CSS media queries for automatic dark mode detection
	DarkModeMedia DarkModeStrategy = iota

	// DarkModeClass uses a CSS class on a parent element
	DarkModeClass

	// DarkModeSelector uses a custom CSS selector
	DarkModeSelector
)

// ThemeConfig contains all design tokens for the Tailwind configuration.
// Each property corresponds to a category in the Tailwind theme configuration.
type ThemeConfig struct {
	// Color properties
	AccentColor         AccentColorConfig
	BackgroundColor     BackgroundColorConfig
	BorderColor         BorderColorConfig
	CaretColor          CaretColorConfig
	Colors              ColorsConfig
	DivideColor         DivideColorConfig
	Fill                FillConfig
	GradientColorStops  GradientColorStopsConfig
	OutlineColor        OutlineColorConfig
	PlaceholderColor    PlaceholderColorConfig
	RingColor           RingColorConfig
	RingOffsetColor     RingOffsetColorConfig
	StrokeColor         StrokeConfig
	TextColor           TextColorConfig
	TextDecorationColor TextDecorationColorConfig

	// Spacing and sizing properties
	Spacing         SpacingConfig
	BorderRadius    BorderRadiusConfig
	BorderWidth     BorderWidthConfig
	DivideWidth     DivideWidthConfig
	Gap             GapConfig
	Height          HeightConfig
	InlineSIze      InlineSizeConfig
	Inset           InsetConfig
	Margin          MarginConfig
	MaxHeight       MaxHeightConfig
	MaxWidth        MaxWidthConfig
	MinHeight       MinHeightConfig
	MinWidth        MinWidthConfig
	OutlineOffset   OutlineOffsetConfig
	OutlineWidth    OutlineWidthConfig
	Padding         PaddingConfig
	RingOffsetWidth RingOffsetWidthConfig
	RingWidth       RingWidthConfig
	ScrollMargin    ScrollMarginConfig
	ScrollPadding   ScrollPaddingConfig
	Space           SpaceConfig
	StrokeWidth     StrokeWidthConfig
	TextIndent      TextIndentConfig
	Width           WidthConfig

	// Typography properties
	FontFamily              FontFamilyConfig
	FontSize                FontSizeConfig
	FontWeight              FontWeightConfig
	LetterSpacing           LetterSpacingConfig
	LineHeight              LineHeightConfig
	ListStyleType           ListStyleTypeConfig
	TextUnderlineOffset     TextUnderlineOffsetConfig
	TextDecorationThickness TextDecorationThicknessConfig

	// Layout properties
	AspectRatio         AspectRatioConfig
	Columns             ColumnsConfig
	FlexBasis           FlexBasisConfig
	GridAutoColumns     GridAutoColumnsConfig
	GridAutoRows        GridAutoRowsConfig
	GridColumn          GridColumnConfig
	GridColumnEnd       GridColumnEndConfig
	GridColumnStart     GridColumnStartConfig
	GridRow             GridRowConfig
	GridRowEnd          GridRowEndConfig
	GridRowStart        GridRowStartConfig
	GridTemplateColumns GridTemplateColumnsConfig
	GridTemplateRows    GridTemplateRowsConfig
	Order               OrderConfig

	// Visual effects properties
	Animation          AnimationConfig
	BackdropBlur       BackdropBlurConfig
	BackdropBrightness BackdropBrightnessConfig
	BackdropContrast   BackdropContrastConfig
	BackdropGrayscale  BackdropGrayscaleConfig
	BackdropHueRotate  BackdropHueRotateConfig
	BackdropInvert     BackdropInvertConfig
	BackdropOpacity    BackdropOpacityConfig
	BackdropSaturate   BackdropSaturateConfig
	BackdropSepia      BackdropSepiaConfig
	Blur               BlurConfig
	Brightness         BrightnessConfig
	BoxShadow          BoxShadowConfig
	BoxShadowColor     BoxShadowColorConfig
	Contrast           ContrastConfig
	DropShadow         DropShadowConfig
	Grayscale          GrayscaleConfig
	HueRotate          HueRotateConfig
	Invert             InvertConfig
	Opacity            OpacityConfig
	Saturate           SaturateConfig
	Sepia              SepiaConfig

	// Background properties
	BackgroundImage    BackgroundImageConfig
	BackgroundOpacity  BackgroundOpacityConfig
	BackgroundPosition BackgroundPositionConfig
	BackgroundSize     BackgroundSizeConfig

	// Border properties
	BorderOpacity BorderOpacityConfig
	BorderSpacing BorderSpacingConfig
	DivideOpacity DivideOpacityConfig

	// Interaction properties
	Cursor CursorConfig

	// Layout control properties
	Container ContainerConfig
	Screens   ScreensConfig

	// Animation properties
	Keyframes                KeyframesConfig
	TransitionDelay          TransitionDelayConfig
	TransitionDuration       TransitionDurationConfig
	TransitionProperty       TransitionPropertyConfig
	TransitionTimingFunction TransitionTimingFunctionConfig

	// Transform properties
	Rotate          RotateConfig
	Scale           ScaleConfig
	Skew            SkewConfig
	TransformOrigin TransformOriginConfig
	Translate       TranslateConfig

	// Responsive properties
	Size SizeConfig

	// Z-index properties
	ZIndex ZIndexConfig

	// Other properties
	Content            ContentConfig
	Flex               FlexConfig
	FlexGrow           FlexGrowConfig
	FlexShrink         FlexShrinkConfig
	LineClamp          LineClampConfig
	PlaceholderOpacity PlaceholderOpacityConfig
	RingOpacity        RingOpacityConfig
	TextOpacity        TextOpacityConfig
	WillChange         WillChangeConfig

	// Extended properties for arbitrary selector support
	Aria     AriaConfig
	Data     DataConfig
	Supports SupportsConfig
}

// Base value types that integrate with cssgen types

// ColorValue represents a color value that can be used in CSS properties.
type ColorValue interface {
	// ToCSSValue converts the color to a css.Value
	ToCSSValue() css.Value

	// String returns the string representation
	String() string
}

// LengthValue represents a length value that can be used in CSS properties.
type LengthValue interface {
	// ToCSSValue converts the length to a css.Value
	ToCSSValue() css.Value

	// String returns the string representation
	String() string
}

// TimeValue represents a time value for animations and transitions.
type TimeValue interface {
	// ToCSSValue converts the time to a css.Value
	ToCSSValue() css.Value

	// String returns the string representation
	String() string
}

// AngleValue represents an angle value for rotations.
type AngleValue interface {
	// ToCSSValue converts the angle to a css.Value
	ToCSSValue() css.Value

	// String returns the string representation
	String() string
}

// NumberValue represents a unitless number value.
type NumberValue interface {
	// ToCSSValue converts the number to a css.Value
	ToCSSValue() css.Value

	// String returns the string representation
	String() string
}

// PercentageValue represents a percentage value.
type PercentageValue interface {
	// ToCSSValue converts the percentage to a css.Value
	ToCSSValue() css.Value

	// String returns the string representation
	String() string
}

// Concrete implementations of common value types

// StaticColor represents a predefined color value.
type StaticColor struct {
	Name  string
	Value css.Color
}

func (c StaticColor) ToCSSValue() css.Value { return c.Value }
func (c StaticColor) String() string        { return c.Name }

// StaticLength represents a predefined length value.
type StaticLength struct {
	Name  string
	Value css.Length
}

func (l StaticLength) ToCSSValue() css.Value { return l.Value }
func (l StaticLength) String() string        { return l.Name }

// StaticTime represents a predefined time value.
type StaticTime struct {
	Name  string
	Value css.Raw // Using Raw for time values since css.Time doesn't exist yet
}

func (t StaticTime) ToCSSValue() css.Value { return t.Value }
func (t StaticTime) String() string        { return t.Name }

// StaticAngle represents a predefined angle value.
type StaticAngle struct {
	Name  string
	Value css.Raw // Using Raw for angle values since css.Angle doesn't exist yet
}

func (a StaticAngle) ToCSSValue() css.Value { return a.Value }
func (a StaticAngle) String() string        { return a.Name }

// StaticNumber represents a predefined number value.
type StaticNumber struct {
	Name  string
	Value css.Raw // Using Raw for number values since css.Number doesn't exist yet
}

func (n StaticNumber) ToCSSValue() css.Value { return n.Value }
func (n StaticNumber) String() string        { return n.Name }

// StaticPercentage represents a predefined percentage value.
type StaticPercentage struct {
	Name  string
	Value css.Length // Percentage is represented as Length
}

func (p StaticPercentage) ToCSSValue() css.Value { return p.Value }
func (p StaticPercentage) String() string        { return p.Name }

// KeywordValue represents a CSS keyword value that integrates with cssgen.
type KeywordValue struct {
	Name    string
	Keyword css.Keyword
}

func (k KeywordValue) ToCSSValue() css.Value { return k.Keyword }
func (k KeywordValue) String() string        { return k.Name }

// ArbitraryValue allows for custom CSS values while maintaining type safety.
type ArbitraryValue struct {
	Value css.Value
}

func (a ArbitraryValue) ToCSSValue() css.Value { return a.Value }
func (a ArbitraryValue) String() string        { return a.Value.String() }

// Helper functions for creating common value types

// ColorFromHex creates a StaticColor from a hex string.
func ColorFromHex(name, hex string) StaticColor {
	return StaticColor{Name: name, Value: css.Hex(hex)}
}

// ColorFromRGB creates a StaticColor from RGB values.
func ColorFromRGB(name string, r, g, b uint8) StaticColor {
	return StaticColor{Name: name, Value: css.RGB(r, g, b)}
}

// ColorFromRGBA creates a StaticColor from RGBA values.
func ColorFromRGBA(name string, r, g, b uint8, a float64) StaticColor {
	return StaticColor{Name: name, Value: css.RGBA(r, g, b, uint8(a*255))}
}

// LengthFromPx creates a StaticLength from pixel value.
func LengthFromPx(name string, px float64) StaticLength {
	return StaticLength{Name: name, Value: css.Px(int(px))}
}

// LengthFromRem creates a StaticLength from rem value.
func LengthFromRem(name string, rem float64) StaticLength {
	return StaticLength{Name: name, Value: css.Rem(rem)}
}

// LengthFromEm creates a StaticLength from em value.
func LengthFromEm(name string, em float64) StaticLength {
	return StaticLength{Name: name, Value: css.Em(em)}
}

// LengthFromPercent creates a StaticLength from percentage value.
func LengthFromPercent(name string, pct float64) StaticLength {
	return StaticLength{Name: name, Value: css.Percent(pct)}
}

// TimeFromMs creates a StaticTime from milliseconds.
func TimeFromMs(name string, ms float64) StaticTime {
	return StaticTime{Name: name, Value: css.Raw(fmt.Sprintf("%.0fms", ms))}
}

// TimeFromS creates a StaticTime from seconds.
func TimeFromS(name string, s float64) StaticTime {
	return StaticTime{Name: name, Value: css.Raw(fmt.Sprintf("%.3fs", s))}
}

// AngleFromDeg creates a StaticAngle from degrees.
func AngleFromDeg(name string, deg float64) StaticAngle {
	return StaticAngle{Name: name, Value: css.Raw(fmt.Sprintf("%.1fdeg", deg))}
}

// NumberFromFloat creates a StaticNumber from a float64.
func NumberFromFloat(name string, val float64) StaticNumber {
	return StaticNumber{Name: name, Value: css.Raw(fmt.Sprintf("%.3f", val))}
}

// PercentageFromFloat creates a StaticPercentage from a float64.
func PercentageFromFloat(name string, pct float64) StaticPercentage {
	return StaticPercentage{Name: name, Value: css.Percent(pct)}
}

// Keyword helpers that integrate with cssgen

// DisplayKeyword creates a KeywordValue from a cssgen DisplayVal.
func DisplayKeyword(name string, val cssgen.DisplayVal) KeywordValue {
	return KeywordValue{Name: name, Keyword: css.Keyword(string(val))}
}

// FlexDirectionKeyword creates a KeywordValue from a cssgen FlexDirectionVal.
func FlexDirectionKeyword(name string, val cssgen.FlexDirectionVal) KeywordValue {
	return KeywordValue{Name: name, Keyword: css.Keyword(string(val))}
}

// JustifyContentKeyword creates a KeywordValue from a cssgen JustifyContentVal.
func JustifyContentKeyword(name string, val cssgen.JustifyContentVal) KeywordValue {
	return KeywordValue{Name: name, Keyword: css.Keyword(string(val))}
}

// AlignItemsKeyword creates a KeywordValue from a cssgen AlignItemsVal.
func AlignItemsKeyword(name string, val cssgen.AlignItemsVal) KeywordValue {
	return KeywordValue{Name: name, Keyword: css.Keyword(string(val))}
}

// PositionKeyword creates a KeywordValue from a cssgen PositionVal.
func PositionKeyword(name string, val cssgen.PositionVal) KeywordValue {
	return KeywordValue{Name: name, Keyword: css.Keyword(string(val))}
}

// FontWeightKeyword creates a KeywordValue from a cssgen FontWeightVal.
func FontWeightKeyword(name string, val cssgen.FontWeightVal) KeywordValue {
	return KeywordValue{Name: name, Keyword: css.Keyword(string(val))}
}

// CursorKeyword creates a KeywordValue from a cssgen CursorVal.
func CursorKeyword(name string, val cssgen.CursorVal) KeywordValue {
	return KeywordValue{Name: name, Keyword: css.Keyword(string(val))}
}
