# Tailwind Utilities Guide

A comprehensive guide to using Tailwind CSS-style utilities with TypeSafe CSS.

## Table of Contents

- [Introduction](#introduction)
- [Basic Usage](#basic-usage)
- [Available Utilities](#available-utilities)
- [Custom Themes](#custom-themes)
- [Utility Generation](#utility-generation)
- [Performance & Deduplication](#performance--deduplication)
- [Best Practices](#best-practices)
- [Migration from Tailwind CSS](#migration-from-tailwind-css)

## Introduction

The TypeSafe CSS Tailwind package provides utility-first CSS authoring with full type safety. Instead of writing custom CSS, you compose designs from small, reusable utility classes.

### Why Tailwind Utilities?

- **Rapid Development**: Compose designs directly in your markup
- **Consistency**: Use a predefined design system
- **Type Safety**: Prevent utility naming errors at compile time
- **Tree Shaking**: Only used utilities are included in output
- **Customization**: Full control over design tokens

## Basic Usage

### Import and Setup

```go
import (
    "github.com/ahmed-com/typesafe-css/css"
    "github.com/ahmed-com/typesafe-css/tailwind"
)

func main() {
    var stylesheet css.Stylesheet
    
    // Add utility classes
    stylesheet.Add(
        tailwind.BgBlue500(),
        tailwind.TextWhite(),
        tailwind.P4(),
        tailwind.RoundedLg(),
    )
    
    fmt.Println(stylesheet.String())
}
```

### Simple Example

Create a button with utilities:

```go
func CreateButtonUtilities() []css.Rule {
    return []css.Rule{
        tailwind.BgBlue500(),     // background-color: #3b82f6
        tailwind.HoverBgBlue600(), // hover:bg-blue-600
        tailwind.TextWhite(),     // color: #ffffff
        tailwind.FontSemibold(),  // font-weight: 600
        tailwind.Py2(),           // padding-top/bottom: 0.5rem
        tailwind.Px4(),           // padding-left/right: 1rem
        tailwind.RoundedMd(),     // border-radius: 0.375rem
        tailwind.CursorPointer(), // cursor: pointer
    }
}
```

HTML usage:
```html
<button class="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-4 rounded-md cursor-pointer">
    Click me
</button>
```

## Available Utilities

### Layout & Display

```go
// Display utilities
tailwind.Block()        // display: block
tailwind.Flex()         // display: flex
tailwind.Grid()         // display: grid
tailwind.Hidden()       // display: none
tailwind.Inline()       // display: inline
tailwind.InlineBlock()  // display: inline-block
tailwind.InlineFlex()   // display: inline-flex

// Flexbox utilities
tailwind.FlexRow()      // flex-direction: row
tailwind.FlexCol()      // flex-direction: column
tailwind.FlexWrap()     // flex-wrap: wrap
tailwind.FlexNoWrap()   // flex-wrap: nowrap

// Justify content
tailwind.JustifyStart()    // justify-content: flex-start
tailwind.JustifyCenter()   // justify-content: center
tailwind.JustifyEnd()      // justify-content: flex-end
tailwind.JustifyBetween()  // justify-content: space-between
tailwind.JustifyAround()   // justify-content: space-around

// Align items
tailwind.ItemsStart()    // align-items: flex-start
tailwind.ItemsCenter()   // align-items: center
tailwind.ItemsEnd()      // align-items: flex-end
tailwind.ItemsStretch()  // align-items: stretch
```

### Spacing

```go
// Padding utilities
tailwind.P0()     // padding: 0
tailwind.P1()     // padding: 0.25rem
tailwind.P2()     // padding: 0.5rem
tailwind.P3()     // padding: 0.75rem
tailwind.P4()     // padding: 1rem
tailwind.P5()     // padding: 1.25rem
tailwind.P6()     // padding: 1.5rem
tailwind.P8()     // padding: 2rem
tailwind.P10()    // padding: 2.5rem
tailwind.P12()    // padding: 3rem

// Directional padding
tailwind.Px2()    // padding-left/right: 0.5rem
tailwind.Py2()    // padding-top/bottom: 0.5rem
tailwind.Pt4()    // padding-top: 1rem
tailwind.Pr4()    // padding-right: 1rem
tailwind.Pb4()    // padding-bottom: 1rem
tailwind.Pl4()    // padding-left: 1rem

// Margin utilities (similar pattern)
tailwind.M0()     // margin: 0
tailwind.M4()     // margin: 1rem
tailwind.MxAuto() // margin-left/right: auto
tailwind.My4()    // margin-top/bottom: 1rem
tailwind.Mt2()    // margin-top: 0.5rem
// ... etc

// Arbitrary spacing
tailwind.P("17px")    // padding: 17px
tailwind.M("2.5rem")  // margin: 2.5rem
```

### Sizing

```go
// Width utilities
tailwind.W0()       // width: 0
tailwind.W1()       // width: 0.25rem
tailwind.W4()       // width: 1rem
tailwind.W8()       // width: 2rem
tailwind.W16()      // width: 4rem
tailwind.W32()      // width: 8rem
tailwind.W64()      // width: 16rem
tailwind.WFull()    // width: 100%
tailwind.WScreen()  // width: 100vw
tailwind.WAuto()    // width: auto

// Height utilities
tailwind.H0()       // height: 0
tailwind.H4()       // height: 1rem
tailwind.H8()       // height: 2rem
tailwind.H16()      // height: 4rem
tailwind.H32()      // height: 8rem
tailwind.H64()      // height: 16rem
tailwind.HFull()    // height: 100%
tailwind.HScreen()  // height: 100vh
tailwind.HAuto()    // height: auto

// Arbitrary sizing
tailwind.W("50px")     // width: 50px
tailwind.H("75vh")     // height: 75vh
```

### Colors

The complete Tailwind color palette is available:

```go
// Background colors
tailwind.BgTransparent()  // background-color: transparent
tailwind.BgWhite()        // background-color: #ffffff
tailwind.BgBlack()        // background-color: #000000

// Gray scale
tailwind.BgGray50()       // background-color: #f9fafb
tailwind.BgGray100()      // background-color: #f3f4f6
tailwind.BgGray200()      // background-color: #e5e7eb
tailwind.BgGray300()      // background-color: #d1d5db
tailwind.BgGray400()      // background-color: #9ca3af
tailwind.BgGray500()      // background-color: #6b7280
tailwind.BgGray600()      // background-color: #4b5563
tailwind.BgGray700()      // background-color: #374151
tailwind.BgGray800()      // background-color: #1f2937
tailwind.BgGray900()      // background-color: #111827

// Blue scale
tailwind.BgBlue50()       // background-color: #eff6ff
tailwind.BgBlue100()      // background-color: #dbeafe
tailwind.BgBlue200()      // background-color: #bfdbfe
tailwind.BgBlue300()      // background-color: #93c5fd
tailwind.BgBlue400()      // background-color: #60a5fa
tailwind.BgBlue500()      // background-color: #3b82f6
tailwind.BgBlue600()      // background-color: #2563eb
tailwind.BgBlue700()      // background-color: #1d4ed8
tailwind.BgBlue800()      // background-color: #1e40af
tailwind.BgBlue900()      // background-color: #1e3a8a

// Text colors (same pattern)
tailwind.TextWhite()      // color: #ffffff
tailwind.TextBlack()      // color: #000000
tailwind.TextGray500()    // color: #6b7280
tailwind.TextBlue500()    // color: #3b82f6
// ... etc

// Additional color families
tailwind.BgRed500()       // Red
tailwind.BgGreen500()     // Green  
tailwind.BgYellow500()    // Yellow
tailwind.BgIndigo500()    // Indigo
tailwind.BgPurple500()    // Purple
tailwind.BgPink500()      // Pink
tailwind.BgEmerald500()   // Emerald
tailwind.BgTeal500()      // Teal
tailwind.BgCyan500()      // Cyan
tailwind.BgSky500()       // Sky
tailwind.BgViolet500()    // Violet
tailwind.BgFuchsia500()   // Fuchsia
tailwind.BgRose500()      // Rose
tailwind.BgSlate500()     // Slate
tailwind.BgZinc500()      // Zinc
tailwind.BgNeutral500()   // Neutral
tailwind.BgStone500()     // Stone

// Arbitrary colors
tailwind.Bg("#bada55")    // background-color: #bada55
tailwind.Text("rgb(255, 0, 0)")  // color: rgb(255, 0, 0)
```

### Typography

```go
// Font size
tailwind.TextXs()     // font-size: 0.75rem
tailwind.TextSm()     // font-size: 0.875rem
tailwind.TextBase()   // font-size: 1rem
tailwind.TextLg()     // font-size: 1.125rem
tailwind.TextXl()     // font-size: 1.25rem
tailwind.Text2xl()    // font-size: 1.5rem
tailwind.Text3xl()    // font-size: 1.875rem
tailwind.Text4xl()    // font-size: 2.25rem

// Font weight
tailwind.FontThin()       // font-weight: 100
tailwind.FontExtralight() // font-weight: 200
tailwind.FontLight()      // font-weight: 300
tailwind.FontNormal()     // font-weight: 400
tailwind.FontMedium()     // font-weight: 500
tailwind.FontSemibold()   // font-weight: 600
tailwind.FontBold()       // font-weight: 700
tailwind.FontExtrabold()  // font-weight: 800
tailwind.FontBlack()      // font-weight: 900

// Text alignment
tailwind.TextLeft()       // text-align: left
tailwind.TextCenter()     // text-align: center
tailwind.TextRight()      // text-align: right
tailwind.TextJustify()    // text-align: justify

// Text decoration
tailwind.Underline()      // text-decoration: underline
tailwind.LineThrough()    // text-decoration: line-through
tailwind.NoUnderline()    // text-decoration: none
```

### Borders

```go
// Border width
tailwind.Border()     // border-width: 1px
tailwind.Border0()    // border-width: 0
tailwind.Border2()    // border-width: 2px
tailwind.Border4()    // border-width: 4px
tailwind.Border8()    // border-width: 8px

// Border radius
tailwind.RoundedNone()  // border-radius: 0
tailwind.RoundedSm()    // border-radius: 0.125rem
tailwind.Rounded()      // border-radius: 0.25rem
tailwind.RoundedMd()    // border-radius: 0.375rem
tailwind.RoundedLg()    // border-radius: 0.5rem
tailwind.RoundedXl()    // border-radius: 0.75rem
tailwind.Rounded2xl()   // border-radius: 1rem
tailwind.Rounded3xl()   // border-radius: 1.5rem
tailwind.RoundedFull()  // border-radius: 9999px

// Border colors
tailwind.BorderTransparent()  // border-color: transparent
tailwind.BorderGray200()      // border-color: #e5e7eb
tailwind.BorderGray300()      // border-color: #d1d5db
tailwind.BorderBlue500()      // border-color: #3b82f6
// ... etc (follows same pattern as background colors)
```

### Shadows & Effects

```go
// Box shadow
tailwind.ShadowSm()     // box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05)
tailwind.Shadow()       // box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06)
tailwind.ShadowMd()     // box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06)
tailwind.ShadowLg()     // box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05)
tailwind.ShadowXl()     // box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04)
tailwind.Shadow2xl()    // box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25)
tailwind.ShadowNone()   // box-shadow: none

// Opacity
tailwind.Opacity0()     // opacity: 0
tailwind.Opacity25()    // opacity: 0.25
tailwind.Opacity50()    // opacity: 0.5
tailwind.Opacity75()    // opacity: 0.75
tailwind.Opacity100()   // opacity: 1

// Filter effects
tailwind.BlurSm()       // filter: blur(4px)
tailwind.BlurClass("") // filter: blur(8px)
tailwind.BrightnessClass("110") // filter: brightness(1.1)
tailwind.GrayscaleClass("")     // filter: grayscale(100%)
```

### Positioning

```go
// Position
tailwind.Static()     // position: static
tailwind.Fixed()      // position: fixed
tailwind.Absolute()   // position: absolute
tailwind.Relative()   // position: relative
tailwind.Sticky()     // position: sticky

// Z-index
tailwind.Z0()     // z-index: 0
tailwind.Z10()    // z-index: 10
tailwind.Z20()    // z-index: 20
tailwind.Z30()    // z-index: 30
tailwind.Z40()    // z-index: 40
tailwind.Z50()    // z-index: 50
tailwind.ZAuto()  // z-index: auto

// Top, right, bottom, left
tailwind.Top0()    // top: 0
tailwind.Right0()  // right: 0
tailwind.Bottom0() // bottom: 0
tailwind.Left0()   // left: 0
```

## Custom Themes

### Creating Custom Themes

Customize the default theme to match your design system:

```go
func CreateBrandTheme() *tailwind.Theme {
    return tailwind.CustomTheme(func(theme *tailwind.Theme) {
        // Brand colors
        theme.AddColor("primary", css.Hex("#6366f1"))
        theme.AddColor("secondary", css.Hex("#8b5cf6"))
        theme.AddColor("accent", css.Hex("#06b6d4"))
        theme.AddColor("danger", css.Hex("#ef4444"))
        theme.AddColor("success", css.Hex("#10b981"))
        theme.AddColor("warning", css.Hex("#f59e0b"))
        
        // Custom spacing scale
        theme.AddSpacing("xs", css.Rem(0.125))   // 2px
        theme.AddSpacing("sm", css.Rem(0.375))   // 6px
        theme.AddSpacing("md", css.Rem(0.875))   // 14px
        theme.AddSpacing("lg", css.Rem(1.75))    // 28px
        theme.AddSpacing("xl", css.Rem(3.5))     // 56px
        theme.AddSpacing("2xl", css.Rem(7))      // 112px
        
        // Custom border radius
        theme.AddBorderRadius("xs", css.Px(2))
        theme.AddBorderRadius("sm", css.Px(4))
        theme.AddBorderRadius("md", css.Px(8))
        theme.AddBorderRadius("lg", css.Px(16))
        theme.AddBorderRadius("xl", css.Px(32))
        theme.AddBorderRadius("pill", css.Px(999))
        
        // Custom font sizes
        theme.AddFontSize("xs", css.Rem(0.625))    // 10px
        theme.AddFontSize("sm", css.Rem(0.75))     // 12px
        theme.AddFontSize("base", css.Rem(0.875))  // 14px
        theme.AddFontSize("lg", css.Rem(1.125))    // 18px
        theme.AddFontSize("xl", css.Rem(1.375))    // 22px
        theme.AddFontSize("2xl", css.Rem(1.75))    // 28px
        theme.AddFontSize("3xl", css.Rem(2.25))    // 36px
    })
}
```

### Using Custom Themes

```go
func main() {
    // Set global custom theme
    customTheme := CreateBrandTheme()
    tailwind.SetDefaultTheme(customTheme)
    
    // Or create a custom utility manager
    manager := tailwind.NewUtilityManager(customTheme)
    
    var stylesheet css.Stylesheet
    stylesheet.Add(
        // Use custom colors
        manager.Bg("primary"),     // Uses your custom primary color
        manager.Text("secondary"), // Uses your custom secondary color
        
        // Use custom spacing
        manager.P("lg"),           // Uses your custom lg spacing
        manager.M("xl"),           // Uses your custom xl spacing
        
        // Mix with standard utilities
        tailwind.FontBold(),
        tailwind.RoundedFull(),
    )
}
```

### Theme Overrides

Override specific parts of the default theme:

```go
func CreateMinimalOverride() *tailwind.Theme {
    return tailwind.CustomTheme(func(theme *tailwind.Theme) {
        // Just override brand colors, keep everything else default
        theme.AddColor("brand", css.Hex("#3b82f6"))
        theme.AddColor("brand-light", css.Hex("#60a5fa"))
        theme.AddColor("brand-dark", css.Hex("#1d4ed8"))
    })
}
```

## Utility Generation

### Generating Common Utilities

Generate a stylesheet with commonly used utilities:

```go
func GenerateUtilityCSS() css.Stylesheet {
    // This generates ~140+ common utility classes
    return tailwind.GenerateUtilityStylesheet()
}

func main() {
    utilityCSS := GenerateUtilityCSS()
    
    // Write to file for static serving
    err := os.WriteFile("public/utilities.css", []byte(utilityCSS.String()), 0644)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Generated %d utility classes\n", len(utilityCSS.Items))
}
```

### Custom Utility Generation

Generate utilities for specific use cases:

```go
func GenerateButtonUtilities() []css.Rule {
    var utilities []css.Rule
    
    // Generate button size utilities
    sizes := map[string]struct{ py, px css.Length }{
        "sm": {css.Px(4), css.Px(8)},
        "md": {css.Px(8), css.Px(16)},
        "lg": {css.Px(12), css.Px(24)},
        "xl": {css.Px(16), css.Px(32)},
    }
    
    for size, padding := range sizes {
        utilities = append(utilities, css.RuleSet(fmt.Sprintf(".btn-%s", size),
            css.Set(css.PaddingTop, padding.py),
            css.Set(css.PaddingBottom, padding.py),
            css.Set(css.PaddingLeft, padding.px),
            css.Set(css.PaddingRight, padding.px),
        ))
    }
    
    // Generate button variant utilities
    variants := map[string]struct{ bg, text css.Color }{
        "primary":   {css.Hex("#3b82f6"), css.Hex("#ffffff")},
        "secondary": {css.Hex("#6b7280"), css.Hex("#ffffff")},
        "success":   {css.Hex("#10b981"), css.Hex("#ffffff")},
        "danger":    {css.Hex("#ef4444"), css.Hex("#ffffff")},
    }
    
    for variant, colors := range variants {
        utilities = append(utilities, css.RuleSet(fmt.Sprintf(".btn-%s", variant),
            css.Set(css.BackgroundColor, colors.bg),
            css.Set(css.ColorP, colors.text),
        ))
    }
    
    return utilities
}
```

## Performance & Deduplication

### Automatic Deduplication

The Tailwind package automatically deduplicates utilities:

```go
func DemonstrateDeduplication() {
    var stylesheet css.Stylesheet
    
    // These will only generate one CSS rule
    stylesheet.Add(
        tailwind.P4(),  // First occurrence
        tailwind.P4(),  // Deduplicated automatically
        tailwind.P4(),  // Deduplicated automatically
    )
    
    // Result: Only one .p-4 { padding: 1rem; } rule
    fmt.Printf("Rules: %d\n", len(stylesheet.Items)) // Output: 1
}
```

### Utility Caching

Utilities are cached for performance:

```go
// First call generates the utility
rule1 := tailwind.P4()

// Subsequent calls return the cached utility (very fast)
rule2 := tailwind.P4()
rule3 := tailwind.P4()

// All rules point to the same cached object
fmt.Printf("Same object: %t\n", &rule1 == &rule2) // Output: true
```

### Tree Shaking

Only called utilities are included in the final CSS:

```go
// Only these utilities will be in the final CSS
var stylesheet css.Stylesheet
stylesheet.Add(
    tailwind.BgBlue500(),  // Included
    tailwind.TextWhite(),  // Included
    tailwind.P4(),         // Included
)

// These utilities are NOT included (not called)
// tailwind.BgRed500()
// tailwind.TextBlack()
// tailwind.P8()
```

## Best Practices

### 1. Organize Utilities by Component

```go
func CreateCardUtilities() []css.Rule {
    return []css.Rule{
        // Layout
        tailwind.BgWhite(),
        tailwind.RoundedLg(),
        tailwind.ShadowMd(),
        
        // Spacing
        tailwind.P6(),
        tailwind.M4(),
        
        // Typography
        tailwind.TextGray800(),
    }
}

func CreateButtonUtilities() []css.Rule {
    return []css.Rule{
        // Base
        tailwind.InlineBlock(),
        tailwind.FontSemibold(),
        tailwind.CursorPointer(),
        
        // Spacing
        tailwind.Py2(),
        tailwind.Px4(),
        
        // Visual
        tailwind.RoundedMd(),
        tailwind.BgBlue500(),
        tailwind.TextWhite(),
    }
}
```

### 2. Use Semantic Grouping

```go
func CreateLayoutUtilities() []css.Rule {
    return []css.Rule{
        // Container
        tailwind.MaxW7xl(),
        tailwind.MxAuto(),
        tailwind.Px4(),
        
        // Grid
        tailwind.Grid(),
        tailwind.GridCols12(),
        tailwind.Gap6(),
    }
}

func CreateTypographyUtilities() []css.Rule {
    return []css.Rule{
        // Headings
        tailwind.Text4xl(),
        tailwind.FontBold(),
        tailwind.TextGray900(),
        
        // Body text
        tailwind.TextBase(),
        tailwind.TextGray600(),
        tailwind.LeadingRelaxed(),
    }
}
```

### 3. Create Utility Bundles

```go
func FlexCenter() []css.Rule {
    return []css.Rule{
        tailwind.Flex(),
        tailwind.ItemsCenter(),
        tailwind.JustifyCenter(),
    }
}

func CardBase() []css.Rule {
    return []css.Rule{
        tailwind.BgWhite(),
        tailwind.RoundedLg(),
        tailwind.ShadowMd(),
        tailwind.P6(),
    }
}

func ButtonPrimary() []css.Rule {
    return []css.Rule{
        tailwind.BgBlue500(),
        tailwind.HoverBgBlue600(),
        tailwind.TextWhite(),
        tailwind.FontSemibold(),
        tailwind.Py2(),
        tailwind.Px4(),
        tailwind.RoundedMd(),
        tailwind.CursorPointer(),
    }
}
```

### 4. Responsive Design Patterns

```go
func CreateResponsiveGrid() []css.Item {
    return []css.Item{
        // Mobile first
        css.RuleSet(".grid",
            css.Set(css.Display, css.DisplayGrid),
            css.Set(css.GridTemplateColumns, css.Raw("1fr")),
            css.Set(css.Gap, css.Px(16)),
        ),
        
        // Tablet
        css.AtRule{
            Name:   "media",
            Params: "(min-width: 768px)",
            Body: []css.Item{
                css.RuleSet(".grid",
                    css.Set(css.GridTemplateColumns, css.Raw("repeat(2, 1fr)")),
                    css.Set(css.Gap, css.Px(24)),
                ),
            },
        },
        
        // Desktop
        css.AtRule{
            Name:   "media",
            Params: "(min-width: 1024px)",
            Body: []css.Item{
                css.RuleSet(".grid",
                    css.Set(css.GridTemplateColumns, css.Raw("repeat(3, 1fr)")),
                    css.Set(css.Gap, css.Px(32)),
                ),
            },
        },
    }
}
```

## Migration from Tailwind CSS

### Class Name Mapping

Most Tailwind classes have equivalent functions:

```
Tailwind CSS         →  TypeSafe CSS
=======================================
bg-blue-500         →  tailwind.BgBlue500()
text-white          →  tailwind.TextWhite()
p-4                 →  tailwind.P4()
rounded-lg          →  tailwind.RoundedLg()
flex                →  tailwind.Flex()
items-center        →  tailwind.ItemsCenter()
justify-center      →  tailwind.JustifyCenter()
font-semibold       →  tailwind.FontSemibold()
hover:bg-blue-600   →  tailwind.HoverBgBlue600()
```

### Converting Existing HTML

Before (Tailwind CSS):
```html
<div class="bg-white rounded-lg shadow-md p-6 m-4">
    <h2 class="text-xl font-bold text-gray-900 mb-4">Card Title</h2>
    <p class="text-gray-600">Card content goes here.</p>
    <button class="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-4 rounded-md mt-4">
        Click me
    </button>
</div>
```

After (TypeSafe CSS):
```go
var stylesheet css.Stylesheet
stylesheet.Add(
    // Card styles
    tailwind.BgWhite(),
    tailwind.RoundedLg(),
    tailwind.ShadowMd(),
    tailwind.P6(),
    tailwind.M4(),
    
    // Title styles
    tailwind.TextXl(),
    tailwind.FontBold(),
    tailwind.TextGray900(),
    tailwind.Mb4(),
    
    // Content styles
    tailwind.TextGray600(),
    
    // Button styles
    tailwind.BgBlue500(),
    tailwind.HoverBgBlue600(),
    tailwind.TextWhite(),
    tailwind.FontSemibold(),
    tailwind.Py2(),
    tailwind.Px4(),
    tailwind.RoundedMd(),
    tailwind.Mt4(),
)
```

### Advanced Features

#### Arbitrary Values
Both support arbitrary values:

```go
// Tailwind CSS: bg-[#bada55]
tailwind.Bg("#bada55")

// Tailwind CSS: p-[17px]
tailwind.P("17px")

// Tailwind CSS: text-[clamp(1rem,4vw,2rem)]
tailwind.Text("clamp(1rem, 4vw, 2rem)")
```

#### Custom Themes
TypeSafe CSS provides more structured theme customization:

```go
// More structured than Tailwind's theme config
customTheme := tailwind.CustomTheme(func(theme *tailwind.Theme) {
    theme.AddColor("brand", css.Hex("#6366f1"))
    theme.AddSpacing("custom", css.Rem(1.5))
    theme.AddBorderRadius("custom", css.Px(12))
})
```

### Benefits of Migration

1. **Compile-time Safety**: Catch utility naming errors at compile time
2. **Type Safety**: Full Go type system benefits
3. **Tree Shaking**: Automatic dead code elimination
4. **IDE Support**: Full autocomplete and IntelliSense
5. **Refactoring**: Safe renaming and refactoring with Go tools
6. **Integration**: Seamless integration with Go toolchain

The TypeSafe CSS Tailwind package provides all the benefits of utility-first CSS with the added safety and tooling of Go's type system.