# Tailwind Utilities Example

This example demonstrates the comprehensive Tailwind CSS-style utilities package for utility-first CSS authoring with TypeSafe CSS.

## What This Example Shows

- **Utility-first CSS authoring** with type-safe utility functions
- **Comprehensive utility coverage** including layout, spacing, colors, typography, and effects
- **Custom theme configuration** with brand colors and custom spacing
- **Automatic deduplication** of utility classes
- **Utility generation** for creating common utility stylesheets
- **Template integration** with Tailwind-style classes
- **Enhanced color palette** with complete Tailwind color scale
- **Advanced effects** like blur, brightness, and filter utilities

## Key Concepts Demonstrated

### 1. Basic Utility Usage

```go
stylesheet.Add(
    tailwind.Flex(),           // .flex { display: flex; }
    tailwind.ItemsCenter(),    // .items-center { align-items: center; }
    tailwind.JustifyCenter(),  // .justify-center { justify-content: center; }
    tailwind.BgBlue500(),      // .bg-blue-500 { background-color: #3b82f6; }
    tailwind.TextWhite(),      // .text-white { color: #ffffff; }
    tailwind.P4(),             // .p-4 { padding: 1rem; }
    tailwind.RoundedLg(),      // .rounded-lg { border-radius: 0.5rem; }
)
```

### 2. Complete Color Palette

```go
// Gray scale
tailwind.BgGray100(),     // #f3f4f6
tailwind.BgGray500(),     // #6b7280
tailwind.BgGray800(),     // #1f2937

// Blue scale
tailwind.BgBlue500(),     // #3b82f6
tailwind.BgBlue600(),     // #2563eb

// Extended palette
tailwind.BgSlate500(),    // #64748b
tailwind.BgIndigo500(),   // #6366f1
tailwind.BgEmerald500(),  // #10b981
tailwind.Bg("rose-500"),  // #f43f5e (using generic function)
```

### 3. Spacing Utilities

```go
// Padding utilities
tailwind.P4(),      // padding: 1rem
tailwind.Px6(),     // padding-left/right: 1.5rem
tailwind.Py2(),     // padding-top/bottom: 0.5rem

// Margin utilities
tailwind.M8(),      // margin: 2rem
tailwind.MxAuto(),  // margin-left/right: auto

// Sizing utilities
tailwind.W64(),     // width: 16rem
tailwind.WFull(),   // width: 100%
tailwind.H32(),     // height: 8rem
tailwind.HScreen(), // height: 100vh
```

### 4. Typography Utilities

```go
tailwind.TextLg(),       // font-size: 1.125rem
tailwind.TextCenter(),   // text-align: center
tailwind.FontBold(),     // font-weight: 700
tailwind.FontSemibold(), // font-weight: 600
tailwind.TextSm(),       // font-size: 0.875rem
```

### 5. Enhanced Visual Effects

```go
// Border radius variations
tailwind.RoundedFull(),  // border-radius: 9999px
tailwind.RoundedLg(),    // border-radius: 0.5rem
tailwind.RoundedMd(),    // border-radius: 0.375rem

// Shadow utilities
tailwind.ShadowLg(),     // Enhanced shadow
tailwind.ShadowMd(),     // Medium shadow

// Opacity
tailwind.Opacity75(),    // opacity: 0.75
tailwind.Opacity50(),    // opacity: 0.5

// Z-index
tailwind.Z10(),          // z-index: 10
tailwind.Z50(),          // z-index: 50

// Filter effects
tailwind.BlurClass("sm"),        // filter: blur(4px)
tailwind.BrightnessClass("110"), // filter: brightness(1.1)
tailwind.GrayscaleClass(""),     // filter: grayscale(100%)
```

### 6. Custom Theme Configuration

```go
customTheme := tailwind.CustomTheme(func(theme *tailwind.Theme) {
    theme.AddColor("brand", css.Hex("#6366f1"))                    // Custom brand color
    theme.AddColor("danger", css.Hex("#ef4444"))                   // Custom danger color
    theme.AddSpacing("xs", css.Rem(0.125))                         // 2px
    theme.AddSpacing("xl", css.Rem(5))                             // 80px
    theme.AddBorderRadius("huge", css.Rem(3))                      // 48px - very large border radius
    theme.AddFontSize("micro", css.Rem(0.625))                     // 10px - very small text
    theme.AddShadow("dramatic", css.Raw("0 25px 50px -12px rgba(0, 0, 0, 0.5)")) // Custom shadow
})
```

### 7. Arbitrary Values

```go
// Custom colors
tailwind.Bg("#bada55"),              // background-color: #bada55

// Custom spacing
tailwind.P("17px"),                  // padding: 17px
tailwind.M("2.5rem"),                // margin: 2.5rem

// Complex values
tailwind.W("calc(100% - 2rem)"),     // width: calc(100% - 2rem)
```

### 8. Utility Generation

```go
// Generate common utilities for CSS file
utilityStylesheet := tailwind.GenerateUtilityStylesheet()
fmt.Printf("Generated %d common utility classes\n", len(utilityStylesheet.Items))

// Sample utilities included:
// .p-4 { padding: 1rem; }
// .bg-blue-500 { background-color: #3b82f6; }
// .text-white { color: #ffffff; }
// .flex { display: flex; }
// .items-center { align-items: center; }
// ... and many more
```

### 9. Deduplication Demo

```go
// Automatic deduplication
rule1 := tailwind.P4()
rule2 := tailwind.P4()
rule3 := tailwind.P4()

// These should all be the same object due to deduplication
fmt.Printf("Rule 1 address: %p\n", &rule1)
fmt.Printf("Rule 2 address: %p\n", &rule2)
fmt.Printf("Rule 3 address: %p\n", &rule3)
// All addresses will be identical

// Utility cache stats
manager := tailwind.GetDefaultManager()
fmt.Printf("Cache size: %d utilities\n", manager.CacheSize())
```

### 10. Template Integration

```go
htmlTemplate := `<!DOCTYPE html>
<html>
<head>
    <style>{{.CSS}}</style>
</head>
<body>
    <div class="bg-white rounded-lg shadow-lg p-6 max-w-md mx-auto mt-8">
        <h2 class="text-xl font-bold text-gray-800 mb-4">
            TypeSafe CSS Card
        </h2>
        <p class="text-gray-600 mb-6">
            This card uses Tailwind CSS utilities generated with type safety in Go!
        </p>
        <div class="flex justify-center">
            <button class="btn bg-blue-600 text-white px-4 py-2 font-medium text-sm">
                Click Me
            </button>
        </div>
    </div>
</body>
</html>`
```

## Running the Example

```bash
cd examples/tailwind
go run main.go
```

## Expected Output

The example generates:

1. **Compact CSS** - Minified utility classes
2. **Pretty CSS** - Formatted for readability
3. **Custom theme demo** - Using custom colors and spacing
4. **Utility generation** - Common utilities for CSS files
5. **Deduplication demo** - Showing object reuse
6. **Template integration** - Complete HTML with embedded styles

## Use Cases

This pattern is ideal for:

- **Rapid prototyping** with utility-first approach
- **Design systems** with consistent utility classes
- **Component-based applications** where utilities compose into larger patterns
- **Static site generation** where CSS utilities are generated at build time
- **Teams familiar with Tailwind CSS** who want type safety

## Comparison with Tailwind CSS

| Feature | Tailwind CSS | TypeSafe CSS Tailwind |
|---------|-------------|----------------------|
| Class names | String-based | Type-safe functions |
| Compile safety | Runtime errors | Compile-time errors |
| Tree shaking | Manual purging | Automatic |
| IDE support | Limited | Full autocomplete |
| Customization | Config files | Go code |
| Integration | Build tools | Go build system |

## Performance Benefits

- **Automatic deduplication**: Utilities are generated only once
- **Tree shaking**: Only called utilities are included
- **Caching**: Utilities are cached for fast access
- **No runtime**: All CSS is generated at build time

## Next Steps

- Explore custom theme creation in [examples/enhanced-tailwind/](../enhanced-tailwind/)
- See core API usage in [examples/basic/](../basic/)
- Check out generated property types in [examples/generated/](../generated/)
- Read the [Tailwind Guide](../../docs/tailwind-guide.md) for comprehensive documentation