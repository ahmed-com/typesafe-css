# Enhanced Tailwind Example

This example demonstrates the advanced typed configuration system for Tailwind CSS utilities with enhanced type safety and compile-time guarantees.

## What This Example Shows

- **Typed configuration system** with structured theme definitions
- **Compile-time type safety** for design tokens
- **Enhanced utility generation** with typed configuration
- **Type-safe value access** without string-based lookups
- **Advanced configuration patterns** for design systems
- **IDE autocomplete support** for all configuration values

## Key Concepts Demonstrated

### 1. Typed Configuration Structure

Instead of string-based maps, the enhanced system provides structured, typed configuration:

```go
// Old approach (string-based)
// theme.Colors["red-50"]  // Runtime error if key doesn't exist

// New approach (typed)
config := tailwind.DefaultConfig()
redColor := config.Theme.Colors.Red50  // Compile-time safety
```

### 2. Type-Safe Design Tokens

All design tokens are strongly typed with compile-time guarantees:

```go
config := tailwind.DefaultConfig()

// Type-safe color access
red50 := config.Theme.Colors.Red50        // Returns *StaticColor
blue500 := config.Theme.Colors.Blue500    // Returns *StaticColor

// Type-safe spacing access  
size1 := config.Theme.Spacing.Size1       // Returns *StaticLength
size4 := config.Theme.Spacing.Size4       // Returns *StaticLength

// Type-safe font size access
textSm := config.Theme.FontSize.Sm.Size    // Returns *StaticLength
textXl := config.Theme.FontSize.Xl.Size    // Returns *StaticLength
```

### 3. Value Type Conversion

Typed values provide safe conversion to CSS values:

```go
config := tailwind.DefaultConfig()

// Type-safe color conversion
redColor := config.Theme.Colors.Red500
cssValue := redColor.ToCSSValue()  // Returns css.Color
fmt.Printf("CSS Value: %s\n", cssValue.String())  // "#ef4444"

// Type-safe spacing conversion
mediumSpacing := config.Theme.Spacing.Size4
cssLength := mediumSpacing.ToCSSValue()  // Returns css.Length
fmt.Printf("CSS Length: %s\n", cssLength.String())  // "1rem"
```

### 4. Custom Value Creation

Create custom typed values with the same safety guarantees:

```go
// Create custom color values
customColor := tailwind.ColorFromHex("brand-primary", "#3b82f6")
customSpacing := tailwind.LengthFromRem("custom-spacing", 1.5)

// Use in CSS generation
cssColor := customColor.ToCSSValue()    // css.Hex("#3b82f6")
cssLength := customSpacing.ToCSSValue() // css.Rem(1.5)
```

### 5. Enhanced Utility Generation

The typed configuration enables more sophisticated utility generation:

```go
config := tailwind.DefaultConfig()
generator := tailwind.NewUtilityGenerator(config)

// Generate utilities with typed configuration
stylesheet := generator.GenerateUtilities()

fmt.Printf("Generated %d utility rules\n", len(stylesheet.Items))
```

### 6. Configuration Inspection

Inspect configuration structure at runtime:

```go
config := tailwind.DefaultConfig()

// Inspect available colors
fmt.Println("Available colors:")
fmt.Printf("  Red 50: %s\n", config.Theme.Colors.Red50.String())
fmt.Printf("  Blue 500: %s\n", config.Theme.Colors.Blue500.String())

// Inspect spacing scale
fmt.Println("Spacing scale:")
fmt.Printf("  Size 1: %s\n", config.Theme.Spacing.Size1.String())
fmt.Printf("  Size 4: %s\n", config.Theme.Spacing.Size4.String())
fmt.Printf("  Size 16: %s\n", config.Theme.Spacing.Size16.String())

// Inspect typography
fmt.Println("Typography:")
fmt.Printf("  Text SM: %s\n", config.Theme.FontSize.Sm.Size.String())
fmt.Printf("  Text LG: %s\n", config.Theme.FontSize.Lg.Size.String())
fmt.Printf("  Text XL: %s\n", config.Theme.FontSize.Xl.Size.String())
```

## Configuration Structure

### Color System

```go
type ColorScale struct {
    // Grayscale
    Black   *StaticColor
    White   *StaticColor
    Gray50  *StaticColor
    Gray100 *StaticColor
    // ... through Gray900
    
    // Color scales (each with 50-900 variants)
    Red50   *StaticColor
    Red100  *StaticColor
    // ... through Red900
    
    Blue50  *StaticColor
    Blue100 *StaticColor
    // ... through Blue900
    
    // Extended palette
    Slate50  *StaticColor
    Zinc50   *StaticColor
    Neutral50 *StaticColor
    Stone50  *StaticColor
    // ... etc
}
```

### Spacing System

```go
type SpacingScale struct {
    Size0   *StaticLength  // 0
    Size1   *StaticLength  // 0.25rem
    Size2   *StaticLength  // 0.5rem
    Size3   *StaticLength  // 0.75rem
    Size4   *StaticLength  // 1rem
    Size5   *StaticLength  // 1.25rem
    Size6   *StaticLength  // 1.5rem
    Size8   *StaticLength  // 2rem
    Size10  *StaticLength  // 2.5rem
    Size12  *StaticLength  // 3rem
    Size16  *StaticLength  // 4rem
    Size20  *StaticLength  // 5rem
    Size24  *StaticLength  // 6rem
    Size32  *StaticLength  // 8rem
    Size40  *StaticLength  // 10rem
    Size48  *StaticLength  // 12rem
    Size56  *StaticLength  // 14rem
    Size64  *StaticLength  // 16rem
}
```

### Typography System

```go
type FontSizeScale struct {
    Xs   FontSizeConfig  // 0.75rem
    Sm   FontSizeConfig  // 0.875rem
    Base FontSizeConfig  // 1rem
    Lg   FontSizeConfig  // 1.125rem
    Xl   FontSizeConfig  // 1.25rem
    Xl2  FontSizeConfig  // 1.5rem
    Xl3  FontSizeConfig  // 1.875rem
    Xl4  FontSizeConfig  // 2.25rem
    Xl5  FontSizeConfig  // 3rem
    Xl6  FontSizeConfig  // 3.75rem
    Xl7  FontSizeConfig  // 4.5rem
    Xl8  FontSizeConfig  // 6rem
    Xl9  FontSizeConfig  // 8rem
}

type FontSizeConfig struct {
    Size       *StaticLength
    LineHeight *StaticLength
}
```

## Type Safety Benefits

### 1. Compile-Time Error Prevention

```go
// ❌ Old approach - runtime errors possible
// color := theme.Colors["red-50"]  // Typo in key name
// if color == nil { /* handle error */ }

// ✅ New approach - compile-time safety
config := tailwind.DefaultConfig()
color := config.Theme.Colors.Red50  // Guaranteed to exist
```

### 2. IDE Autocomplete

The typed system provides full IDE support:

```go
config := tailwind.DefaultConfig()

// IDE shows all available colors
config.Theme.Colors.  // Autocomplete shows: Red50, Red100, Blue500, etc.

// IDE shows all available spacing
config.Theme.Spacing. // Autocomplete shows: Size1, Size2, Size4, etc.
```

### 3. Refactoring Safety

Go's refactoring tools work perfectly with the typed system:

```go
// Renaming fields is safe across the entire codebase
config.Theme.Colors.Red500  // Can be safely renamed with Go tools
```

## Advanced Patterns

### 1. Configuration Composition

```go
func CreateDesignSystem() *tailwind.Config {
    config := tailwind.DefaultConfig()
    
    // Configuration is composable and type-safe
    primaryColor := config.Theme.Colors.Blue500
    accentColor := config.Theme.Colors.Indigo500
    spacing := config.Theme.Spacing.Size4
    
    // Use in custom components
    return config
}
```

### 2. Theme Validation

```go
func ValidateTheme(config *tailwind.Config) error {
    // All theme values are guaranteed to exist
    if config.Theme.Colors.Red50 == nil {
        return errors.New("red-50 not configured") // This check is unnecessary now
    }
    
    // Type system prevents invalid configurations
    return nil
}
```

### 3. Utility Generation with Types

```go
func GenerateCustomUtilities(config *tailwind.Config) css.Stylesheet {
    var stylesheet css.Stylesheet
    
    // Type-safe utility generation
    colors := []struct {
        name  string
        color *tailwind.StaticColor
    }{
        {"primary", config.Theme.Colors.Blue500},
        {"secondary", config.Theme.Colors.Gray500},
        {"accent", config.Theme.Colors.Indigo500},
    }
    
    for _, c := range colors {
        rule := css.RuleSet(fmt.Sprintf(".bg-%s", c.name),
            css.Set(css.BackgroundColor, c.color.ToCSSValue()),
        )
        stylesheet.Add(rule)
    }
    
    return stylesheet
}
```

## Running the Example

```bash
cd examples/enhanced-tailwind
go run main.go
```

## Expected Output

The example demonstrates:

1. **Configuration structure** showing typed access to design tokens
2. **Type safety verification** confirming compile-time guarantees
3. **Value conversion** showing CSS value generation
4. **Utility generation** with the enhanced system
5. **Performance metrics** showing generated utility count

## Comparison: String-Based vs Typed

| Feature | String-Based | Typed Configuration |
|---------|-------------|-------------------|
| Error detection | Runtime | Compile-time |
| IDE support | Limited | Full autocomplete |
| Refactoring | Manual | Automatic |
| Type safety | None | Complete |
| Performance | Map lookups | Direct access |
| Learning curve | Higher | Lower (guided by types) |

## Migration Path

### From String-Based
```go
// Old approach
color := theme.Colors["red-500"]
spacing := theme.Spacing["4"]

// New approach
config := tailwind.DefaultConfig()
color := config.Theme.Colors.Red500
spacing := config.Theme.Spacing.Size4
```

### From Manual Configuration
```go
// Old approach
colors := map[string]css.Color{
    "red-500": css.Hex("#ef4444"),
    "blue-500": css.Hex("#3b82f6"),
}

// New approach
config := tailwind.DefaultConfig()
red := config.Theme.Colors.Red500.ToCSSValue()
blue := config.Theme.Colors.Blue500.ToCSSValue()
```

## Future Enhancements

The typed configuration system enables:

1. **Design system validation** at compile time
2. **Automatic documentation** generation from types
3. **Theme migration tools** with type safety
4. **IDE plugins** leveraging type information
5. **Performance optimizations** through static analysis

## Best Practices

### 1. Use Default Configuration
Start with the comprehensive default configuration:

```go
config := tailwind.DefaultConfig()
// All standard Tailwind values are available
```

### 2. Extend Thoughtfully
Add custom values when the defaults don't meet your needs:

```go
// Add custom brand colors
brandColor := tailwind.ColorFromHex("brand", "#custom")
```

### 3. Leverage Type Safety
Let the type system guide your usage:

```go
// IDE autocomplete shows all available options
config.Theme.Colors. // Shows all color options
```

## Next Steps

- Explore utility patterns in [examples/tailwind/](../tailwind/)
- See core API usage in [examples/basic/](../basic/)
- Check out generated properties in [examples/generated/](../generated/)
- Read the comprehensive [Tailwind Guide](../../docs/tailwind-guide.md)