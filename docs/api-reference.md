# API Reference

Complete reference for all TypeSafe CSS packages and APIs.

## Table of Contents

- [Core Package (css)](#core-package-css)
- [Generated Package (cssgen)](#generated-package-cssgen)
- [Tailwind Package (tailwind)](#tailwind-package-tailwind)
- [Type Definitions](#type-definitions)
- [Function Reference](#function-reference)
- [Constants](#constants)

## Core Package (css)

### Import
```go
import "github.com/ahmed-com/typesafe-css/css"
```

### Value Constructors

#### Length Values
```go
func Px(value float64) Length            // Pixel units
func Rem(value float64) Length           // Rem units  
func Em(value float64) Length            // Em units
func Percent(value float64) Length       // Percentage units
func Vh(value float64) Length            // Viewport height units
func Vw(value float64) Length            // Viewport width units
```

**Example:**
```go
css.Px(16)        // "16px"
css.Rem(1.5)      // "1.5rem"
css.Percent(100)  // "100%"
css.Vh(50)        // "50vh"
```

#### Color Values
```go
func Hex(value string) Color                    // Hex color
func RGB(r, g, b int) Color                     // RGB color
func RGBA(r, g, b int, a float64) Color         // RGBA color
func HSL(h, s, l int) Color                     // HSL color
func HSLA(h, s, l int, a float64) Color         // HSLA color
```

**Example:**
```go
css.Hex("#3b82f6")           // "#3b82f6"
css.RGB(59, 130, 246)        // "rgb(59, 130, 246)"
css.RGBA(59, 130, 246, 0.5)  // "rgba(59, 130, 246, 0.5)"
css.HSL(217, 91, 60)         // "hsl(217, 91%, 60%)"
```

#### Other Values
```go
func Raw(value string) Value       // Raw CSS value
func Var(name string) Value        // CSS custom property
func Keyword(value string) Value   // CSS keyword
```

**Example:**
```go
css.Raw("calc(100% - 20px)")       // "calc(100% - 20px)"
css.Var("--color-primary")         // "var(--color-primary)"
css.Keyword("auto")                // "auto"
```

### Structure Types

#### Declaration
```go
func Set(property Property, value Value) Declaration
```

**Example:**
```go
css.Set(css.Width, css.Px(100))
css.Set(css.ColorP, css.Hex("#000000"))
```

#### Rule
```go
func RuleSet(selector string, declarations ...Declaration) Rule
```

**Example:**
```go
css.RuleSet(".btn",
    css.Set(css.Display, css.DisplayInlineBlock),
    css.Set(css.Padding, css.Px(16)),
    css.Set(css.BackgroundColor, css.Hex("#3b82f6")),
)
```

#### AtRule
```go
type AtRule struct {
    Name   string
    Params string
    Body   []Item
}
```

**Example:**
```go
css.AtRule{
    Name:   "media",
    Params: "(max-width: 768px)",
    Body: []css.Item{
        css.RuleSet(".btn",
            css.Set(css.Display, css.DisplayBlock),
        ),
    },
}
```

#### Stylesheet
```go
type Stylesheet struct {
    Items []Item
}

func (s *Stylesheet) Add(items ...Item)        // Add items to stylesheet
func (s *Stylesheet) String() string           // Generate compact CSS
```

**Example:**
```go
var stylesheet css.Stylesheet
stylesheet.Add(rule1, rule2, atRule)
css := stylesheet.String()
```

### Shorthand Helpers

#### Padding/Margin
```go
func PadAll(value Length) Value                 // Padding all sides
func PadXY(x, y Length) Value                   // Padding horizontal/vertical
func PadTRBL(t, r, b, l Length) Value          // Padding top/right/bottom/left

func MarAll(value Length) Value                 // Margin all sides
func MarXY(x, y Length) Value                   // Margin horizontal/vertical
func MarTRBL(t, r, b, l Length) Value          // Margin top/right/bottom/left
```

**Example:**
```go
css.Set(css.Padding, css.PadXY(css.Px(8), css.Px(16)))  // "16px 8px"
css.Set(css.Margin, css.MarAll(css.Px(12)))             // "12px"
```

#### Border
```go
func BorderShorthand(width Length, style Value, color Color) Value
```

**Example:**
```go
css.Set(css.Border, css.BorderShorthand(
    css.Px(1), 
    css.BorderSolid, 
    css.Hex("#e5e7eb")
)) // "1px solid #e5e7eb"
```

#### Flexbox
```go
func FlexCenter() []Declaration                 // Center with flexbox
```

**Example:**
```go
css.RuleSet(".container",
    css.FlexCenter()..., // display: flex; justify-content: center; align-items: center;
)
```

### Utility Functions

#### CSS Output
```go
func CSS(items ...Item) string                  // Compact CSS output
func PrettyCSS(items ...Item) string           // Pretty-printed CSS output
```

**Example:**
```go
compact := css.CSS(rule1, rule2)               // Minified
pretty := css.PrettyCSS(rule1, rule2)          // Formatted
```

### Property Constants

#### Layout & Display
```go
const (
    Display               Property = "display"
    Position              Property = "position"
    Top                   Property = "top"
    Right                 Property = "right"
    Bottom                Property = "bottom"
    Left                  Property = "left"
    ZIndex                Property = "z-index"
    Float                 Property = "float"
    Clear                 Property = "clear"
)
```

#### Box Model
```go
const (
    Width                 Property = "width"
    Height                Property = "height"
    MaxWidth              Property = "max-width"
    MaxHeight             Property = "max-height"
    MinWidth              Property = "min-width"
    MinHeight             Property = "min-height"
    Padding               Property = "padding"
    PaddingTop            Property = "padding-top"
    PaddingRight          Property = "padding-right"
    PaddingBottom         Property = "padding-bottom"
    PaddingLeft           Property = "padding-left"
    Margin                Property = "margin"
    MarginTop             Property = "margin-top"
    MarginRight           Property = "margin-right"
    MarginBottom          Property = "margin-bottom"
    MarginLeft            Property = "margin-left"
    BoxSizing             Property = "box-sizing"
)
```

#### Borders & Backgrounds
```go
const (
    Border                Property = "border"
    BorderTop             Property = "border-top"
    BorderRight           Property = "border-right"
    BorderBottom          Property = "border-bottom"
    BorderLeft            Property = "border-left"
    BorderWidth           Property = "border-width"
    BorderStyle           Property = "border-style"
    BorderColor           Property = "border-color"
    BorderRadius          Property = "border-radius"
    BackgroundColor       Property = "background-color"
    BackgroundImage       Property = "background-image"
    BackgroundSize        Property = "background-size"
    BackgroundPosition    Property = "background-position"
    BackgroundRepeat      Property = "background-repeat"
)
```

#### Typography
```go
const (
    ColorP                Property = "color"        // Note: ColorP to avoid conflict
    FontFamily            Property = "font-family"
    FontSize              Property = "font-size"
    FontWeight            Property = "font-weight"
    FontStyle             Property = "font-style"
    LineHeight            Property = "line-height"
    TextAlign             Property = "text-align"
    TextDecoration        Property = "text-decoration"
    TextTransform         Property = "text-transform"
    LetterSpacing         Property = "letter-spacing"
    WordSpacing           Property = "word-spacing"
)
```

#### Flexbox & Grid
```go
const (
    FlexDirection         Property = "flex-direction"
    FlexWrap              Property = "flex-wrap"
    JustifyContent        Property = "justify-content"
    AlignItems            Property = "align-items"
    AlignContent          Property = "align-content"
    FlexGrow              Property = "flex-grow"
    FlexShrink            Property = "flex-shrink"
    FlexBasis             Property = "flex-basis"
    GridTemplateColumns   Property = "grid-template-columns"
    GridTemplateRows      Property = "grid-template-rows"
    GridGap               Property = "grid-gap"
    Gap                   Property = "gap"
)
```

### Value Constants

#### Display Values
```go
const (
    DisplayBlock          Value = "block"
    DisplayInline         Value = "inline"
    DisplayInlineBlock    Value = "inline-block"
    DisplayFlex           Value = "flex"
    DisplayInlineFlex     Value = "inline-flex"
    DisplayGrid           Value = "grid"
    DisplayInlineGrid     Value = "inline-grid"
    DisplayNone           Value = "none"
)
```

#### Position Values
```go
const (
    PositionStatic        Value = "static"
    PositionRelative      Value = "relative"
    PositionAbsolute      Value = "absolute"
    PositionFixed         Value = "fixed"
    PositionSticky        Value = "sticky"
)
```

#### Text Alignment
```go
const (
    TextAlignLeft         Value = "left"
    TextAlignRight        Value = "right"
    TextAlignCenter       Value = "center"
    TextAlignJustify      Value = "justify"
)
```

#### Font Weights
```go
const (
    FontWeight100         Value = "100"
    FontWeight200         Value = "200"
    FontWeight300         Value = "300"
    FontWeight400         Value = "400"
    FontWeight500         Value = "500"
    FontWeight600         Value = "600"
    FontWeight700         Value = "700"
    FontWeight800         Value = "800"
    FontWeight900         Value = "900"
    FontWeightNormal      Value = "normal"
    FontWeightBold        Value = "bold"
)
```

#### Cursor Values
```go
const (
    CursorDefault         Value = "default"
    CursorPointer         Value = "pointer"
    CursorHelp            Value = "help"
    CursorWait            Value = "wait"
    CursorText            Value = "text"
    CursorMove            Value = "move"
    CursorNotAllowed      Value = "not-allowed"
)
```

#### Border Styles
```go
const (
    BorderSolid           Value = "solid"
    BorderDashed          Value = "dashed"
    BorderDotted          Value = "dotted"
    BorderDouble          Value = "double"
    BorderGroove          Value = "groove"
    BorderRidge           Value = "ridge"
    BorderInset           Value = "inset"
    BorderOutset          Value = "outset"
    BorderNone            Value = "none"
)
```

#### Flexbox Values
```go
const (
    FlexDirectionRow      Value = "row"
    FlexDirectionColumn   Value = "column"
    FlexDirectionRowReverse Value = "row-reverse"
    FlexDirectionColumnReverse Value = "column-reverse"
    
    JustifyContentFlexStart Value = "flex-start"
    JustifyContentFlexEnd   Value = "flex-end"
    JustifyContentCenter    Value = "center"
    JustifyContentSpaceBetween Value = "space-between"
    JustifyContentSpaceAround  Value = "space-around"
    JustifyContentSpaceEvenly  Value = "space-evenly"
    
    AlignItemsFlexStart     Value = "flex-start"
    AlignItemsFlexEnd       Value = "flex-end"
    AlignItemsCenter        Value = "center"
    AlignItemsBaseline      Value = "baseline"
    AlignItemsStretch       Value = "stretch"
)
```

## Generated Package (cssgen)

### Import
```go
import "github.com/ahmed-com/typesafe-css/cssgen"
```

### Generated Setters

#### Display
```go
func SetDisplay(value DisplayValue) Declaration

// Available values:
cssgen.DisplayValBlock
cssgen.DisplayValInline
cssgen.DisplayValInlineBlock
cssgen.DisplayValFlex
cssgen.DisplayValInlineFlex
cssgen.DisplayValGrid
cssgen.DisplayValInlineGrid
cssgen.DisplayValNone
cssgen.DisplayValContents
cssgen.DisplayValFlowRoot
// ... and more
```

#### Position
```go
func SetPosition(value PositionValue) Declaration

// Available values:
cssgen.PositionValStatic
cssgen.PositionValRelative
cssgen.PositionValAbsolute
cssgen.PositionValFixed
cssgen.PositionValSticky
```

#### Text Alignment
```go
func SetTextAlign(value TextAlignValue) Declaration

// Available values:
cssgen.TextAlignValLeft
cssgen.TextAlignValRight
cssgen.TextAlignValCenter
cssgen.TextAlignValJustify
cssgen.TextAlignValStart
cssgen.TextAlignValEnd
cssgen.TextAlignValMatchParent
```

#### Font Weight
```go
func SetFontWeight(value FontWeightValue) Declaration

// Available values:
cssgen.FontWeightValNormal
cssgen.FontWeightValBold
cssgen.FontWeightVal100
cssgen.FontWeightVal200
// ... through FontWeightVal900
```

#### Cursor
```go
func SetCursor(value CursorValue) Declaration

// Available values:
cssgen.CursorValDefault
cssgen.CursorValPointer
cssgen.CursorValHelp
cssgen.CursorValWait
cssgen.CursorValText
cssgen.CursorValMove
cssgen.CursorValNotAllowed
cssgen.CursorValColResize
cssgen.CursorValRowResize
cssgen.CursorValZoomIn
cssgen.CursorValZoomOut
// ... and many more
```

#### Flexbox
```go
func SetFlexDirection(value FlexDirectionValue) Declaration
func SetJustifyContent(value JustifyContentValue) Declaration
func SetAlignItems(value AlignItemsValue) Declaration

// FlexDirection values:
cssgen.FlexDirectionValRow
cssgen.FlexDirectionValColumn
cssgen.FlexDirectionValRowReverse
cssgen.FlexDirectionValColumnReverse

// JustifyContent values:
cssgen.JustifyContentValFlexStart
cssgen.JustifyContentValFlexEnd
cssgen.JustifyContentValCenter
cssgen.JustifyContentValSpaceBetween
cssgen.JustifyContentValSpaceAround
cssgen.JustifyContentValSpaceEvenly

// AlignItems values:
cssgen.AlignItemsValFlexStart
cssgen.AlignItemsValFlexEnd
cssgen.AlignItemsValCenter
cssgen.AlignItemsValBaseline
cssgen.AlignItemsValStretch
cssgen.AlignItemsValStart
cssgen.AlignItemsValEnd
cssgen.AlignItemsValSelfStart
cssgen.AlignItemsValSelfEnd
```

## Tailwind Package (tailwind)

### Import
```go
import "github.com/ahmed-com/typesafe-css/tailwind"
```

### Layout Utilities

#### Display
```go
func Block() Rule                    // display: block
func Flex() Rule                     // display: flex
func Grid() Rule                     // display: grid
func Hidden() Rule                   // display: none
func Inline() Rule                   // display: inline
func InlineBlock() Rule              // display: inline-block
func InlineFlex() Rule               // display: inline-flex
```

#### Flexbox
```go
func FlexRow() Rule                  // flex-direction: row
func FlexCol() Rule                  // flex-direction: column
func FlexWrap() Rule                 // flex-wrap: wrap
func FlexNoWrap() Rule               // flex-wrap: nowrap
func JustifyStart() Rule             // justify-content: flex-start
func JustifyCenter() Rule            // justify-content: center
func JustifyEnd() Rule               // justify-content: flex-end
func JustifyBetween() Rule           // justify-content: space-between
func ItemsStart() Rule               // align-items: flex-start
func ItemsCenter() Rule              // align-items: center
func ItemsEnd() Rule                 // align-items: flex-end
func ItemsStretch() Rule             // align-items: stretch
```

### Spacing Utilities

#### Padding
```go
func P0() Rule                       // padding: 0
func P1() Rule                       // padding: 0.25rem
func P2() Rule                       // padding: 0.5rem
func P3() Rule                       // padding: 0.75rem
func P4() Rule                       // padding: 1rem
func P5() Rule                       // padding: 1.25rem
func P6() Rule                       // padding: 1.5rem
func P8() Rule                       // padding: 2rem
func P10() Rule                      // padding: 2.5rem
func P12() Rule                      // padding: 3rem

// Directional padding
func Px1() Rule                      // padding-left/right: 0.25rem
func Py1() Rule                      // padding-top/bottom: 0.25rem
func Pt1() Rule                      // padding-top: 0.25rem
func Pr1() Rule                      // padding-right: 0.25rem
func Pb1() Rule                      // padding-bottom: 0.25rem
func Pl1() Rule                      // padding-left: 0.25rem
// ... (similar pattern for all sizes)

// Arbitrary values
func P(value string) Rule            // padding: {value}
```

#### Margin
```go
func M0() Rule                       // margin: 0
func M1() Rule                       // margin: 0.25rem
func M2() Rule                       // margin: 0.5rem
// ... (similar pattern to padding)

func MxAuto() Rule                   // margin-left/right: auto
func MyAuto() Rule                   // margin-top/bottom: auto

// Arbitrary values
func M(value string) Rule            // margin: {value}
```

### Sizing Utilities

#### Width
```go
func W0() Rule                       // width: 0
func W1() Rule                       // width: 0.25rem
func W2() Rule                       // width: 0.5rem
// ... continuing pattern
func W64() Rule                      // width: 16rem
func WFull() Rule                    // width: 100%
func WScreen() Rule                  // width: 100vw
func WAuto() Rule                    // width: auto

// Arbitrary values
func W(value string) Rule            // width: {value}
```

#### Height
```go
func H0() Rule                       // height: 0
func H1() Rule                       // height: 0.25rem
// ... (similar pattern to width)
func HFull() Rule                    // height: 100%
func HScreen() Rule                  // height: 100vh
func HAuto() Rule                    // height: auto

// Arbitrary values
func H(value string) Rule            // height: {value}
```

### Color Utilities

#### Background Colors
```go
// Grayscale
func BgTransparent() Rule            // background-color: transparent
func BgWhite() Rule                  // background-color: #ffffff
func BgBlack() Rule                  // background-color: #000000
func BgGray50() Rule                 // background-color: #f9fafb
func BgGray100() Rule                // background-color: #f3f4f6
// ... through BgGray900

// Color scales (each color has 50-900 variants)
func BgRed50() Rule                  // background-color: #fef2f2
func BgRed500() Rule                 // background-color: #ef4444
func BgBlue50() Rule                 // background-color: #eff6ff
func BgBlue500() Rule                // background-color: #3b82f6
// ... (red, orange, amber, yellow, lime, green, emerald, teal, cyan, sky, blue, indigo, violet, purple, fuchsia, pink, rose)

// Extended palette
func BgSlate50() Rule                // Slate color scale
func BgZinc50() Rule                 // Zinc color scale
func BgNeutral50() Rule              // Neutral color scale
func BgStone50() Rule                // Stone color scale

// Arbitrary values
func Bg(value string) Rule           // background-color: {value}
```

#### Text Colors
```go
// Same pattern as background colors
func TextWhite() Rule                // color: #ffffff
func TextBlack() Rule                // color: #000000
func TextGray500() Rule              // color: #6b7280
func TextBlue500() Rule              // color: #3b82f6
// ... (all color variants)

// Arbitrary values
func Text(value string) Rule         // color: {value}
```

### Typography Utilities

#### Font Size
```go
func TextXs() Rule                   // font-size: 0.75rem
func TextSm() Rule                   // font-size: 0.875rem
func TextBase() Rule                 // font-size: 1rem
func TextLg() Rule                   // font-size: 1.125rem
func TextXl() Rule                   // font-size: 1.25rem
func Text2xl() Rule                  // font-size: 1.5rem
func Text3xl() Rule                  // font-size: 1.875rem
func Text4xl() Rule                  // font-size: 2.25rem
func Text5xl() Rule                  // font-size: 3rem
func Text6xl() Rule                  // font-size: 3.75rem
func Text7xl() Rule                  // font-size: 4.5rem
func Text8xl() Rule                  // font-size: 6rem
func Text9xl() Rule                  // font-size: 8rem
```

#### Font Weight
```go
func FontThin() Rule                 // font-weight: 100
func FontExtralight() Rule           // font-weight: 200
func FontLight() Rule                // font-weight: 300
func FontNormal() Rule               // font-weight: 400
func FontMedium() Rule               // font-weight: 500
func FontSemibold() Rule             // font-weight: 600
func FontBold() Rule                 // font-weight: 700
func FontExtrabold() Rule            // font-weight: 800
func FontBlack() Rule                // font-weight: 900
```

#### Text Alignment
```go
func TextLeft() Rule                 // text-align: left
func TextCenter() Rule               // text-align: center
func TextRight() Rule                // text-align: right
func TextJustify() Rule              // text-align: justify
```

### Border Utilities

#### Border Width
```go
func Border() Rule                   // border-width: 1px
func Border0() Rule                  // border-width: 0
func Border2() Rule                  // border-width: 2px
func Border4() Rule                  // border-width: 4px
func Border8() Rule                  // border-width: 8px
```

#### Border Radius
```go
func RoundedNone() Rule              // border-radius: 0
func RoundedSm() Rule                // border-radius: 0.125rem
func Rounded() Rule                  // border-radius: 0.25rem
func RoundedMd() Rule                // border-radius: 0.375rem
func RoundedLg() Rule                // border-radius: 0.5rem
func RoundedXl() Rule                // border-radius: 0.75rem
func Rounded2xl() Rule               // border-radius: 1rem
func Rounded3xl() Rule               // border-radius: 1.5rem
func RoundedFull() Rule              // border-radius: 9999px
```

### Effects Utilities

#### Shadow
```go
func ShadowSm() Rule                 // Small shadow
func Shadow() Rule                   // Default shadow
func ShadowMd() Rule                 // Medium shadow
func ShadowLg() Rule                 // Large shadow
func ShadowXl() Rule                 // Extra large shadow
func Shadow2xl() Rule                // 2X large shadow
func ShadowNone() Rule               // No shadow
```

#### Opacity
```go
func Opacity0() Rule                 // opacity: 0
func Opacity25() Rule                // opacity: 0.25
func Opacity50() Rule                // opacity: 0.5
func Opacity75() Rule                // opacity: 0.75
func Opacity100() Rule               // opacity: 1
```

### Theme Management

#### Default Theme
```go
func DefaultTheme() *Theme           // Returns default Tailwind theme
func SetDefaultTheme(theme *Theme)   // Sets global default theme
```

#### Custom Themes
```go
func CustomTheme(modifiers ...func(*Theme)) *Theme

// Theme modification functions
func (t *Theme) AddColor(name string, value css.Color)
func (t *Theme) AddSpacing(name string, value css.Length)
func (t *Theme) AddBorderRadius(name string, value css.Length)
func (t *Theme) AddFontSize(name string, value css.Length)
func (t *Theme) AddShadow(name string, value css.Value)
```

#### Utility Management
```go
func NewUtilityManager(theme *Theme) *UtilityManager
func GetDefaultManager() *UtilityManager

func GenerateUtilityStylesheet() css.Stylesheet  // Generate common utilities
```

## Type Definitions

### Core Types
```go
type Value interface {
    String() string
}

type Property string

type Declaration struct {
    Property Property
    Value    Value
}

type Rule interface {
    String() string
}

type Item interface {
    String() string
}

type Length interface {
    Value
}

type Color interface {
    Value
}
```

### Tailwind Types
```go
type Theme struct {
    Colors        map[string]ColorToken
    Spacing       map[string]ValueToken
    BorderRadius  map[string]ValueToken
    FontSize      map[string]ValueToken
    // ... other theme properties
}

type UtilityManager struct {
    theme *Theme
    cache map[string]css.Rule
}
```

This API reference covers the essential functions and types. For the most up-to-date information, use `go doc` or check the source code directly.