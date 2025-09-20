// This file defines the specific configuration structures for each theme category.
// Each configuration struct provides strongly-typed properties instead of string maps.

package tailwind

// Color-related configurations

// ColorsConfig defines the base color palette with strongly-typed color properties.
type ColorsConfig struct {
	// CSS system colors
	Inherit     ColorValue
	Current     ColorValue
	Transparent ColorValue

	// Neutral colors
	Black ColorValue
	White ColorValue

	// Slate color scale
	Slate50  ColorValue
	Slate100 ColorValue
	Slate200 ColorValue
	Slate300 ColorValue
	Slate400 ColorValue
	Slate500 ColorValue
	Slate600 ColorValue
	Slate700 ColorValue
	Slate800 ColorValue
	Slate900 ColorValue
	Slate950 ColorValue

	// Gray color scale
	Gray50  ColorValue
	Gray100 ColorValue
	Gray200 ColorValue
	Gray300 ColorValue
	Gray400 ColorValue
	Gray500 ColorValue
	Gray600 ColorValue
	Gray700 ColorValue
	Gray800 ColorValue
	Gray900 ColorValue
	Gray950 ColorValue

	// Zinc color scale
	Zinc50  ColorValue
	Zinc100 ColorValue
	Zinc200 ColorValue
	Zinc300 ColorValue
	Zinc400 ColorValue
	Zinc500 ColorValue
	Zinc600 ColorValue
	Zinc700 ColorValue
	Zinc800 ColorValue
	Zinc900 ColorValue
	Zinc950 ColorValue

	// Neutral color scale
	Neutral50  ColorValue
	Neutral100 ColorValue
	Neutral200 ColorValue
	Neutral300 ColorValue
	Neutral400 ColorValue
	Neutral500 ColorValue
	Neutral600 ColorValue
	Neutral700 ColorValue
	Neutral800 ColorValue
	Neutral900 ColorValue
	Neutral950 ColorValue

	// Stone color scale
	Stone50  ColorValue
	Stone100 ColorValue
	Stone200 ColorValue
	Stone300 ColorValue
	Stone400 ColorValue
	Stone500 ColorValue
	Stone600 ColorValue
	Stone700 ColorValue
	Stone800 ColorValue
	Stone900 ColorValue
	Stone950 ColorValue

	// Red color scale
	Red50  ColorValue
	Red100 ColorValue
	Red200 ColorValue
	Red300 ColorValue
	Red400 ColorValue
	Red500 ColorValue
	Red600 ColorValue
	Red700 ColorValue
	Red800 ColorValue
	Red900 ColorValue
	Red950 ColorValue

	// Orange color scale
	Orange50  ColorValue
	Orange100 ColorValue
	Orange200 ColorValue
	Orange300 ColorValue
	Orange400 ColorValue
	Orange500 ColorValue
	Orange600 ColorValue
	Orange700 ColorValue
	Orange800 ColorValue
	Orange900 ColorValue
	Orange950 ColorValue

	// Amber color scale
	Amber50  ColorValue
	Amber100 ColorValue
	Amber200 ColorValue
	Amber300 ColorValue
	Amber400 ColorValue
	Amber500 ColorValue
	Amber600 ColorValue
	Amber700 ColorValue
	Amber800 ColorValue
	Amber900 ColorValue
	Amber950 ColorValue

	// Yellow color scale
	Yellow50  ColorValue
	Yellow100 ColorValue
	Yellow200 ColorValue
	Yellow300 ColorValue
	Yellow400 ColorValue
	Yellow500 ColorValue
	Yellow600 ColorValue
	Yellow700 ColorValue
	Yellow800 ColorValue
	Yellow900 ColorValue
	Yellow950 ColorValue

	// Lime color scale
	Lime50  ColorValue
	Lime100 ColorValue
	Lime200 ColorValue
	Lime300 ColorValue
	Lime400 ColorValue
	Lime500 ColorValue
	Lime600 ColorValue
	Lime700 ColorValue
	Lime800 ColorValue
	Lime900 ColorValue
	Lime950 ColorValue

	// Green color scale
	Green50  ColorValue
	Green100 ColorValue
	Green200 ColorValue
	Green300 ColorValue
	Green400 ColorValue
	Green500 ColorValue
	Green600 ColorValue
	Green700 ColorValue
	Green800 ColorValue
	Green900 ColorValue
	Green950 ColorValue

	// Emerald color scale
	Emerald50  ColorValue
	Emerald100 ColorValue
	Emerald200 ColorValue
	Emerald300 ColorValue
	Emerald400 ColorValue
	Emerald500 ColorValue
	Emerald600 ColorValue
	Emerald700 ColorValue
	Emerald800 ColorValue
	Emerald900 ColorValue
	Emerald950 ColorValue

	// Teal color scale
	Teal50  ColorValue
	Teal100 ColorValue
	Teal200 ColorValue
	Teal300 ColorValue
	Teal400 ColorValue
	Teal500 ColorValue
	Teal600 ColorValue
	Teal700 ColorValue
	Teal800 ColorValue
	Teal900 ColorValue
	Teal950 ColorValue

	// Cyan color scale
	Cyan50  ColorValue
	Cyan100 ColorValue
	Cyan200 ColorValue
	Cyan300 ColorValue
	Cyan400 ColorValue
	Cyan500 ColorValue
	Cyan600 ColorValue
	Cyan700 ColorValue
	Cyan800 ColorValue
	Cyan900 ColorValue
	Cyan950 ColorValue

	// Sky color scale
	Sky50  ColorValue
	Sky100 ColorValue
	Sky200 ColorValue
	Sky300 ColorValue
	Sky400 ColorValue
	Sky500 ColorValue
	Sky600 ColorValue
	Sky700 ColorValue
	Sky800 ColorValue
	Sky900 ColorValue
	Sky950 ColorValue

	// Blue color scale
	Blue50  ColorValue
	Blue100 ColorValue
	Blue200 ColorValue
	Blue300 ColorValue
	Blue400 ColorValue
	Blue500 ColorValue
	Blue600 ColorValue
	Blue700 ColorValue
	Blue800 ColorValue
	Blue900 ColorValue
	Blue950 ColorValue

	// Indigo color scale
	Indigo50  ColorValue
	Indigo100 ColorValue
	Indigo200 ColorValue
	Indigo300 ColorValue
	Indigo400 ColorValue
	Indigo500 ColorValue
	Indigo600 ColorValue
	Indigo700 ColorValue
	Indigo800 ColorValue
	Indigo900 ColorValue
	Indigo950 ColorValue

	// Violet color scale
	Violet50  ColorValue
	Violet100 ColorValue
	Violet200 ColorValue
	Violet300 ColorValue
	Violet400 ColorValue
	Violet500 ColorValue
	Violet600 ColorValue
	Violet700 ColorValue
	Violet800 ColorValue
	Violet900 ColorValue
	Violet950 ColorValue

	// Purple color scale
	Purple50  ColorValue
	Purple100 ColorValue
	Purple200 ColorValue
	Purple300 ColorValue
	Purple400 ColorValue
	Purple500 ColorValue
	Purple600 ColorValue
	Purple700 ColorValue
	Purple800 ColorValue
	Purple900 ColorValue
	Purple950 ColorValue

	// Fuchsia color scale
	Fuchsia50  ColorValue
	Fuchsia100 ColorValue
	Fuchsia200 ColorValue
	Fuchsia300 ColorValue
	Fuchsia400 ColorValue
	Fuchsia500 ColorValue
	Fuchsia600 ColorValue
	Fuchsia700 ColorValue
	Fuchsia800 ColorValue
	Fuchsia900 ColorValue
	Fuchsia950 ColorValue

	// Pink color scale
	Pink50  ColorValue
	Pink100 ColorValue
	Pink200 ColorValue
	Pink300 ColorValue
	Pink400 ColorValue
	Pink500 ColorValue
	Pink600 ColorValue
	Pink700 ColorValue
	Pink800 ColorValue
	Pink900 ColorValue
	Pink950 ColorValue

	// Rose color scale
	Rose50  ColorValue
	Rose100 ColorValue
	Rose200 ColorValue
	Rose300 ColorValue
	Rose400 ColorValue
	Rose500 ColorValue
	Rose600 ColorValue
	Rose700 ColorValue
	Rose800 ColorValue
	Rose900 ColorValue
	Rose950 ColorValue
}

// AccentColorConfig defines accent color values.
type AccentColorConfig struct {
	Auto ColorValue
	// Inherits from Colors
}

// BackgroundColorConfig defines background color values.
type BackgroundColorConfig struct {
	// Inherits from Colors
}

// BorderColorConfig defines border color values.
type BorderColorConfig struct {
	DEFAULT ColorValue // Default border color
	// Inherits from Colors
}

// CaretColorConfig defines caret color values.
type CaretColorConfig struct {
	// Inherits from Colors
}

// DivideColorConfig defines divide color values.
type DivideColorConfig struct {
	// Inherits from BorderColor
}

// FillConfig defines fill color values.
type FillConfig struct {
	None ColorValue
	// Inherits from Colors
}

// GradientColorStopsConfig defines gradient color stop values.
type GradientColorStopsConfig struct {
	// Inherits from Colors
}

// OutlineColorConfig defines outline color values.
type OutlineColorConfig struct {
	// Inherits from Colors
}

// PlaceholderColorConfig defines placeholder color values.
type PlaceholderColorConfig struct {
	// Inherits from Colors
}

// RingColorConfig defines ring color values.
type RingColorConfig struct {
	DEFAULT ColorValue // Default ring color (blue-500)
	// Inherits from Colors
}

// RingOffsetColorConfig defines ring offset color values.
type RingOffsetColorConfig struct {
	// Inherits from Colors
}

// StrokeConfig defines stroke color values.
type StrokeConfig struct {
	None ColorValue
	// Inherits from Colors
}

// TextColorConfig defines text color values.
type TextColorConfig struct {
	// Inherits from Colors
}

// TextDecorationColorConfig defines text decoration color values.
type TextDecorationColorConfig struct {
	// Inherits from Colors
}

// Spacing and sizing configurations

// SpacingConfig defines spacing scale values.
type SpacingConfig struct {
	Px      LengthValue // 1px
	Size0   LengthValue // 0px
	Size0_5 LengthValue // 0.125rem
	Size1   LengthValue // 0.25rem
	Size1_5 LengthValue // 0.375rem
	Size2   LengthValue // 0.5rem
	Size2_5 LengthValue // 0.625rem
	Size3   LengthValue // 0.75rem
	Size3_5 LengthValue // 0.875rem
	Size4   LengthValue // 1rem
	Size5   LengthValue // 1.25rem
	Size6   LengthValue // 1.5rem
	Size7   LengthValue // 1.75rem
	Size8   LengthValue // 2rem
	Size9   LengthValue // 2.25rem
	Size10  LengthValue // 2.5rem
	Size11  LengthValue // 2.75rem
	Size12  LengthValue // 3rem
	Size14  LengthValue // 3.5rem
	Size16  LengthValue // 4rem
	Size20  LengthValue // 5rem
	Size24  LengthValue // 6rem
	Size28  LengthValue // 7rem
	Size32  LengthValue // 8rem
	Size36  LengthValue // 9rem
	Size40  LengthValue // 10rem
	Size44  LengthValue // 11rem
	Size48  LengthValue // 12rem
	Size52  LengthValue // 13rem
	Size56  LengthValue // 14rem
	Size60  LengthValue // 15rem
	Size64  LengthValue // 16rem
	Size72  LengthValue // 18rem
	Size80  LengthValue // 20rem
	Size96  LengthValue // 24rem
}

// BorderRadiusConfig defines border radius values.
type BorderRadiusConfig struct {
	None    LengthValue // 0px
	Sm      LengthValue // 0.125rem
	Base    LengthValue // 0.25rem (default)
	Md      LengthValue // 0.375rem
	Lg      LengthValue // 0.5rem
	Xl      LengthValue // 0.75rem
	Size2xl LengthValue // 1rem
	Size3xl LengthValue // 1.5rem
	Full    LengthValue // 9999px
}

// BorderWidthConfig defines border width values.
type BorderWidthConfig struct {
	DEFAULT LengthValue // 1px
	Size0   LengthValue // 0px
	Size2   LengthValue // 2px
	Size4   LengthValue // 4px
	Size8   LengthValue // 8px
}

// ... (more spacing configs would continue here)

// Typography configurations

// FontFamilyConfig defines font family values.
type FontFamilyConfig struct {
	Sans  []string // ui-sans-serif, system-ui, ...
	Serif []string // ui-serif, Georgia, ...
	Mono  []string // ui-monospace, SFMono-Regular, ...
}

// FontSizeConfig defines font size values with optional line-height.
type FontSizeConfig struct {
	Xs      FontSizeValue // 0.75rem, line-height: 1rem
	Sm      FontSizeValue // 0.875rem, line-height: 1.25rem
	Base    FontSizeValue // 1rem, line-height: 1.5rem
	Lg      FontSizeValue // 1.125rem, line-height: 1.75rem
	Xl      FontSizeValue // 1.25rem, line-height: 1.75rem
	Size2xl FontSizeValue // 1.5rem, line-height: 2rem
	Size3xl FontSizeValue // 1.875rem, line-height: 2.25rem
	Size4xl FontSizeValue // 2.25rem, line-height: 2.5rem
	Size5xl FontSizeValue // 3rem, line-height: 1
	Size6xl FontSizeValue // 3.75rem, line-height: 1
	Size7xl FontSizeValue // 4.5rem, line-height: 1
	Size8xl FontSizeValue // 6rem, line-height: 1
	Size9xl FontSizeValue // 8rem, line-height: 1
}

// FontSizeValue represents a font size with optional line height.
type FontSizeValue struct {
	Size       LengthValue
	LineHeight LengthValue // Optional line height (can be nil interface)
}

// FontWeightConfig defines font weight values using cssgen types.
type FontWeightConfig struct {
	Thin       KeywordValue // 100
	Extralight KeywordValue // 200
	Light      KeywordValue // 300
	Normal     KeywordValue // 400
	Medium     KeywordValue // 500
	Semibold   KeywordValue // 600
	Bold       KeywordValue // 700
	Extrabold  KeywordValue // 800
	Black      KeywordValue // 900
}

// Layout configurations

// AspectRatioConfig defines aspect ratio values.
type AspectRatioConfig struct {
	Auto   KeywordValue // auto
	Square string       // 1 / 1
	Video  string       // 16 / 9
}

// AnimationConfig defines animation values.
type AnimationConfig struct {
	None   KeywordValue // none
	Spin   string       // spin 1s linear infinite
	Ping   string       // ping 1s cubic-bezier(0, 0, 0.2, 1) infinite
	Pulse  string       // pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite
	Bounce string       // bounce 1s infinite
}

// Visual effects configurations

// BlurConfig defines blur values.
type BlurConfig struct {
	Size0   LengthValue // 0
	None    string      // ""
	Sm      LengthValue // 4px
	DEFAULT LengthValue // 8px
	Md      LengthValue // 12px
	Lg      LengthValue // 16px
	Xl      LengthValue // 24px
	Size2xl LengthValue // 40px
	Size3xl LengthValue // 64px
}

// BrightnessConfig defines brightness values.
type BrightnessConfig struct {
	Size0   NumberValue // 0
	Size50  NumberValue // .5
	Size75  NumberValue // .75
	Size90  NumberValue // .9
	Size95  NumberValue // .95
	Size100 NumberValue // 1
	Size105 NumberValue // 1.05
	Size110 NumberValue // 1.1
	Size125 NumberValue // 1.25
	Size150 NumberValue // 1.5
	Size200 NumberValue // 2
}

// OpacityConfig defines opacity values.
type OpacityConfig struct {
	Size0   NumberValue // 0
	Size5   NumberValue // 0.05
	Size10  NumberValue // 0.1
	Size15  NumberValue // 0.15
	Size20  NumberValue // 0.2
	Size25  NumberValue // 0.25
	Size30  NumberValue // 0.3
	Size35  NumberValue // 0.35
	Size40  NumberValue // 0.4
	Size45  NumberValue // 0.45
	Size50  NumberValue // 0.5
	Size55  NumberValue // 0.55
	Size60  NumberValue // 0.6
	Size65  NumberValue // 0.65
	Size70  NumberValue // 0.7
	Size75  NumberValue // 0.75
	Size80  NumberValue // 0.8
	Size85  NumberValue // 0.85
	Size90  NumberValue // 0.9
	Size95  NumberValue // 0.95
	Size100 NumberValue // 1
}

// CursorConfig defines cursor values using cssgen types.
type CursorConfig struct {
	Auto         KeywordValue // auto
	Default      KeywordValue // default
	Pointer      KeywordValue // pointer
	Wait         KeywordValue // wait
	Text         KeywordValue // text
	Move         KeywordValue // move
	Help         KeywordValue // help
	NotAllowed   KeywordValue // not-allowed
	None         KeywordValue // none
	ContextMenu  KeywordValue // context-menu
	Progress     KeywordValue // progress
	Cell         KeywordValue // cell
	Crosshair    KeywordValue // crosshair
	VerticalText KeywordValue // vertical-text
	Alias        KeywordValue // alias
	Copy         KeywordValue // copy
	NoDrop       KeywordValue // no-drop
	Grab         KeywordValue // grab
	Grabbing     KeywordValue // grabbing
	AllScroll    KeywordValue // all-scroll
	ColResize    KeywordValue // col-resize
	RowResize    KeywordValue // row-resize
	NResize      KeywordValue // n-resize
	EResize      KeywordValue // e-resize
	SResize      KeywordValue // s-resize
	WResize      KeywordValue // w-resize
	NeResize     KeywordValue // ne-resize
	NwResize     KeywordValue // nw-resize
	SeResize     KeywordValue // se-resize
	SwResize     KeywordValue // sw-resize
	EwResize     KeywordValue // ew-resize
	NsResize     KeywordValue // ns-resize
	NeswResize   KeywordValue // nesw-resize
	NwseResize   KeywordValue // nwse-resize
	ZoomIn       KeywordValue // zoom-in
	ZoomOut      KeywordValue // zoom-out
}

// Responsive configurations

// ScreensConfig defines responsive breakpoints.
type ScreensConfig struct {
	Sm      LengthValue // 640px
	Md      LengthValue // 768px
	Lg      LengthValue // 1024px
	Xl      LengthValue // 1280px
	Size2xl LengthValue // 1536px
}

// Placeholder configurations for remaining types
// These would be fully implemented with proper types

type DivideWidthConfig struct{}
type GapConfig struct{}
type HeightConfig struct{}
type InlineSizeConfig struct{}
type InsetConfig struct{}
type MarginConfig struct{}
type MaxHeightConfig struct{}
type MaxWidthConfig struct{}
type MinHeightConfig struct{}
type MinWidthConfig struct{}
type OutlineOffsetConfig struct{}
type OutlineWidthConfig struct{}
type PaddingConfig struct{}
type RingOffsetWidthConfig struct{}
type RingWidthConfig struct{}
type ScrollMarginConfig struct{}
type ScrollPaddingConfig struct{}
type SpaceConfig struct{}
type StrokeWidthConfig struct{}
type TextIndentConfig struct{}
type WidthConfig struct{}
type LetterSpacingConfig struct{}
type LineHeightConfig struct{}
type ListStyleTypeConfig struct{}
type TextUnderlineOffsetConfig struct{}
type TextDecorationThicknessConfig struct{}
type ColumnsConfig struct{}
type FlexBasisConfig struct{}
type GridAutoColumnsConfig struct{}
type GridAutoRowsConfig struct{}
type GridColumnConfig struct{}
type GridColumnEndConfig struct{}
type GridColumnStartConfig struct{}
type GridRowConfig struct{}
type GridRowEndConfig struct{}
type GridRowStartConfig struct{}
type GridTemplateColumnsConfig struct{}
type GridTemplateRowsConfig struct{}
type OrderConfig struct{}
type BackdropBlurConfig struct{}
type BackdropBrightnessConfig struct{}
type BackdropContrastConfig struct{}
type BackdropGrayscaleConfig struct{}
type BackdropHueRotateConfig struct{}
type BackdropInvertConfig struct{}
type BackdropOpacityConfig struct{}
type BackdropSaturateConfig struct{}
type BackdropSepiaConfig struct{}
type BoxShadowConfig struct{}
type BoxShadowColorConfig struct{}
type ContrastConfig struct{}
type DropShadowConfig struct{}
type GrayscaleConfig struct{}
type HueRotateConfig struct{}
type InvertConfig struct{}
type SaturateConfig struct{}
type SepiaConfig struct{}
type BackgroundImageConfig struct{}
type BackgroundOpacityConfig struct{}
type BackgroundPositionConfig struct{}
type BackgroundSizeConfig struct{}
type BorderOpacityConfig struct{}
type BorderSpacingConfig struct{}
type DivideOpacityConfig struct{}
type ContainerConfig struct{}
type KeyframesConfig struct{}
type TransitionDelayConfig struct{}
type TransitionDurationConfig struct{}
type TransitionPropertyConfig struct{}
type TransitionTimingFunctionConfig struct{}
type RotateConfig struct{}
type ScaleConfig struct{}
type SkewConfig struct{}
type TransformOriginConfig struct{}
type TranslateConfig struct{}
type SizeConfig struct{}
type ZIndexConfig struct{}
type ContentConfig struct{}
type FlexConfig struct{}
type FlexGrowConfig struct{}
type FlexShrinkConfig struct{}
type LineClampConfig struct{}
type PlaceholderOpacityConfig struct{}
type RingOpacityConfig struct{}
type TextOpacityConfig struct{}
type WillChangeConfig struct{}
type AriaConfig struct{}
type DataConfig struct{}
type SupportsConfig struct{}
