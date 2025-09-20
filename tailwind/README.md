# Tailwind CSS Utilities for TypeSafe CSS

This package provides Tailwind CSS-style utility classes for the TypeSafe CSS library, enabling utility-first CSS authoring with full type safety in Go.

## Features

- **Tailwind-inspired utility classes** - Common utilities like `bg-blue-500`, `text-center`, `p-4`, etc.
- **Theme-based design tokens** - Configurable colors, spacing, typography, and breakpoints
- **Automatic deduplication** - Utility classes are generated only once, even when called multiple times
- **Custom themes** - Override default theme values or add your own design tokens
- **Type-safe API** - Build on top of the core TypeSafe CSS foundation
- **Zero dependencies** - Pure Go implementation that extends the core css package

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/ahmed-com/typesafe-css/css"
    "github.com/ahmed-com/typesafe-css/tailwind"
)

func main() {
    var stylesheet css.Stylesheet
    
    // Add Tailwind-style utilities
    stylesheet.Add(
        tailwind.Bg("blue-500"),     // .bg-blue-500 { background-color: #3b82f6; }
        tailwind.Text("white"),      // .text-white { color: #ffffff; }
        tailwind.P("4"),             // .p-4 { padding: 1rem; }
        tailwind.DisplayFlex(),      // .flex { display: flex; }
        tailwind.JustifyCenterClass(), // .justify-center { justify-content: center; }
    )
    
    fmt.Println(stylesheet.String())
}
```

## Available Utilities

### Layout & Display
- `DisplayFlex()`, `DisplayBlock()`, `DisplayGrid()`, `DisplayHidden()`
- `FlexRowClass()`, `FlexColClass()`
- `JustifyStartClass()`, `JustifyCenterClass()`, `JustifyBetweenClass()`
- `ItemsStartClass()`, `ItemsCenterClass()`, `ItemsStretchClass()`

### Spacing
- `P(size)`, `Px(size)`, `Py(size)` - Padding utilities
- `M(size)`, `Mx(size)`, `My(size)` - Margin utilities  
- `W(size)`, `H(size)` - Width and height utilities

### Colors (Complete Tailwind Palette)
- `Bg(color)` - Background color utilities for all 247 Tailwind colors
- `Text(color)` - Text color utilities
- `Border(color)` - Border color utilities
- **Color families**: slate, gray, zinc, neutral, stone, red, orange, amber, yellow, lime, green, emerald, teal, cyan, sky, blue, indigo, violet, purple, fuchsia, pink, rose
- **Special colors**: inherit, current, transparent, black, white

### Typography
- `TextXs()`, `TextSm()`, `TextBase()`, `TextLg()`, `TextXl()` - Font sizes
- `FontThin()`, `FontNormal()`, `FontBold()` - Font weights (9 weights available)
- `TextLeft()`, `TextCenter()`, `TextRight()` - Text alignment
- Line height utilities with 14 different values

### Position & Layout
- `StaticClass()`, `RelativeClass()`, `AbsoluteClass()`, `FixedClass()`
- `Z10()`, `Z50()`, `ZAuto()` - Z-index utilities

### Visual Effects
- **Border Radius**: `RoundedSm()`, `RoundedLg()`, `RoundedFull()` - 9 radius values
- **Box Shadows**: `ShadowSm()`, `ShadowLg()`, `Shadow2xl()` - 8 shadow presets
- **Opacity**: `Opacity0()` through `Opacity100()` - 21 opacity levels
- **Filters**: `BlurClass()`, `BrightnessClass()`, `GrayscaleClass()`, `ContrastClass()`, etc.

## Theme Configuration

### Default Theme

The package includes a comprehensive Tailwind CSS configuration with:

```go
theme := tailwind.DefaultTheme()
// 247 colors: All Tailwind color families with full intensity ranges
// 35 spacing values: 0, px, 0.5, 1, 1.5, 2, 2.5, 3, 4, 5, 6, 8, 10, 12, 16, 20, 24, etc.
// 13 font sizes: xs, sm, base, lg, xl, 2xl, 3xl, 4xl, 5xl, 6xl, 7xl, 8xl, 9xl
// 9 border radius values: none, sm, md, lg, xl, 2xl, 3xl, full
// 8 box shadow presets: sm, md, lg, xl, 2xl, inner, none
// 21 opacity levels: 0, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100
// 9 font weights: thin, extralight, light, normal, medium, semibold, bold, extrabold, black
// Plus many more configuration categories matching Tailwind CSS defaults
```

### Custom Themes

Create custom themes by overriding default values. Now supports all Tailwind configuration categories:

```go
customTheme := tailwind.CustomTheme(func(theme *tailwind.Theme) {
    theme.AddColor("brand", css.Hex("#6366f1"))   // custom brand color
    theme.AddColor("accent", css.Hex("#f59e0b"))  // accent color
    theme.AddColor("danger", css.Hex("#ef4444"))  // danger color

    theme.AddSpacing("xs", css.Rem(0.125)) // 2px
    theme.AddSpacing("xl", css.Rem(5))     // 80px
    theme.AddSpacing("2xl", css.Rem(6))    // 96px

    theme.AddBorderRadius("huge", css.Rem(3)) // 48px
    theme.AddBoxShadow("glow", css.Raw("0 0 20px rgba(99, 102, 241, 0.5)"))
    theme.AddOpacity("15", css.Raw("0.15"))
    theme.AddFontWeight("super", css.Keyword("950"))
})

// Use with a custom manager
manager := tailwind.WithCustomTheme(customTheme)
bgRule := tailwind.BackgroundColor(manager, "brand")  // .bg-brand { background-color: #6366f1; }
radiusRule := tailwind.BorderRadius(manager, "huge")  // .rounded-huge { border-radius: 3rem; }

// Or set as the new default theme
tailwind.SetDefaultTheme(customTheme)
```

## Enhanced Visual Effects

The expanded configuration includes comprehensive visual effect utilities:

```go
// Border radius utilities
stylesheet.Add(
    tailwind.RoundedSm(),     // border-radius: 0.125rem
    tailwind.RoundedLg(),     // border-radius: 0.5rem  
    tailwind.RoundedFull(),   // border-radius: 9999px
    tailwind.Rounded("xl"),   // border-radius: 0.75rem
)

// Box shadow utilities  
stylesheet.Add(
    tailwind.ShadowSm(),      // subtle shadow
    tailwind.ShadowLg(),      // larger shadow
    tailwind.Shadow2xl(),     // very large shadow
    tailwind.ShadowInner(),   // inset shadow
)

// Opacity utilities
stylesheet.Add(
    tailwind.Opacity25(),     // opacity: 0.25
    tailwind.Opacity50(),     // opacity: 0.5
    tailwind.Opacity75(),     // opacity: 0.75
)

// Z-index utilities
stylesheet.Add(
    tailwind.Z10(),           // z-index: 10
    tailwind.Z50(),           // z-index: 50
    tailwind.ZAuto(),         // z-index: auto
)

// Filter effects
stylesheet.Add(
    tailwind.BlurClass("sm"),       // filter: blur(4px)
    tailwind.BrightnessClass("110"), // filter: brightness(1.1)
    tailwind.GrayscaleClass(""),     // filter: grayscale(100%)
    tailwind.ContrastClass("125"),   // filter: contrast(1.25)
)
```

## Complete Color Palette

Now supports all 247 Tailwind CSS colors:

```go
// Use any Tailwind color
stylesheet.Add(
    tailwind.BgSlate500(),    // All slate colors
    tailwind.BgIndigo500(),   // All indigo colors  
    tailwind.BgEmerald500(),  // All emerald colors
    tailwind.BgRose500(),     // All rose colors
    
    // Or use the generic function
    tailwind.Bg("violet-400"),
    tailwind.Bg("amber-200"),
    tailwind.Text("cyan-600"),
)
```

## Utility Management & Deduplication

The package automatically handles deduplication of utility classes:

```go
manager := tailwind.GetDefaultManager()

// These calls will generate the same rule (cached)
rule1 := tailwind.P("4")
rule2 := tailwind.P("4")
rule3 := tailwind.P("4")

// All rules will have identical CSS output
fmt.Println(rule1.String()) // .p-4{padding:1rem}
fmt.Println(rule2.String()) // .p-4{padding:1rem}  
fmt.Println(rule3.String()) // .p-4{padding:1rem}

// Manager caches unique rules
fmt.Printf("Cached rules: %d\n", len(manager.GetRules()))
```

## Generating Common Utilities

Generate a comprehensive set of common utilities:

```go
// Generate 100+ common utility classes
stylesheet := tailwind.GenerateUtilityStylesheet()

// Includes:
// - Display utilities (flex, block, hidden, etc.)
// - Flexbox utilities (justify-center, items-center, etc.)
// - Position utilities (relative, absolute, etc.)
// - Text utilities (text-xs through text-3xl)
// - Font weight utilities (font-thin through font-bold)
// - Spacing utilities for common sizes (p-0 through p-24, etc.)
// - Color utilities for neutral and primary colors

fmt.Printf("Generated %d utility classes\n", len(stylesheet.Items))
```

## Template Integration

Use with Go templates for complete HTML generation:

```go
htmlTemplate := `<!DOCTYPE html>
<html>
<head>
    <style>{{.CSS}}</style>
</head>
<body class="bg-gray-100">
    <div class="flex justify-center items-center">
        <div class="bg-white p-6 rounded shadow">
            <h1 class="text-lg font-bold text-center">Hello World</h1>
            <button class="bg-blue-500 text-white px-4 py-2 rounded">
                Click Me
            </button>
        </div>
    </div>
</body>
</html>`

var stylesheet css.Stylesheet
stylesheet.Add(
    tailwind.Bg("gray-100"),
    tailwind.DisplayFlex(),
    tailwind.JustifyCenterClass(),
    tailwind.ItemsCenterClass(),
    tailwind.Bg("white"),
    tailwind.P("6"),
    // ... other utilities
)

tmpl := template.Must(template.New("page").Parse(htmlTemplate))
data := struct{ CSS template.CSS }{
    CSS: template.CSS(stylesheet.String()),
}
tmpl.Execute(os.Stdout, data)
```

## API Reference

### Core Functions

- `DefaultTheme() *Theme` - Returns the default Tailwind-inspired theme
- `CustomTheme(modifiers ...func(*Theme)) *Theme` - Creates a custom theme by mutating a fresh default
- `NewUtilityManager(theme *Theme) *UtilityManager` - Creates a new utility manager
- `GetDefaultManager() *UtilityManager` - Returns the global default manager
- `SetDefaultTheme(theme *Theme)` - Updates the default theme
- `GenerateUtilityStylesheet() css.Stylesheet` - Generates common utilities

### Utility Categories

#### Colors
- `Text(color string) css.Rule` - Text color utility
- `Bg(color string) css.Rule` - Background color utility  
- `Border(color string) css.Rule` - Border color utility

#### Spacing  
- `P(size string) css.Rule` - Padding utility
- `Px(size string) css.Rule` - Horizontal padding utility
- `Py(size string) css.Rule` - Vertical padding utility
- `M(size string) css.Rule` - Margin utility
- `Mx(size string) css.Rule` - Horizontal margin utility
- `My(size string) css.Rule` - Vertical margin utility
- `W(size string) css.Rule` - Width utility
- `H(size string) css.Rule` - Height utility

#### Layout
- `DisplayFlex() css.Rule` - Flex display utility
- `DisplayBlock() css.Rule` - Block display utility
- `DisplayHidden() css.Rule` - Hidden display utility
- `FlexRowClass() css.Rule` - Flex row direction
- `FlexColClass() css.Rule` - Flex column direction
- `JustifyCenterClass() css.Rule` - Justify content center
- `ItemsCenterClass() css.Rule` - Align items center

#### Typography
- `TextXs() css.Rule` through `Text3xl() css.Rule` - Font sizes
- `FontThin() css.Rule` through `FontBold() css.Rule` - Font weights
- `TextLeft() css.Rule`, `TextCenter() css.Rule`, `TextRight() css.Rule` - Text alignment

## Examples

See [examples/tailwind/main.go](../examples/tailwind/main.go) for a comprehensive demonstration of features including:

- Basic utility usage
- Custom theme configuration  
- Deduplication demonstration
- Utility generation
- Template integration

## Performance

- Utilities are cached and deduplicated automatically
- Thread-safe concurrent usage
- Minimal memory allocation
- Fast CSS generation and serialization

## Comparison with Tailwind CSS

This package provides similar utility patterns to Tailwind CSS but with Go's type safety:

| Tailwind CSS | TypeSafe CSS Tailwind |
|--------------|----------------------|
| `bg-blue-500` | `tailwind.Bg("blue-500")` |
| `text-center` | `tailwind.TextCenter()` |
| `p-4` | `tailwind.P("4")` |
| `flex justify-center` | `tailwind.DisplayFlex()`, `tailwind.JustifyCenterClass()` |

The key advantages:

- **Type safety** - No typos in class names or values
- **IDE support** - Full autocomplete and refactoring
- **Compile-time validation** - Catch errors before runtime
- **Go integration** - Use variables, loops, and functions to generate CSS
- **No build step** - Pure Go, no external tools required
