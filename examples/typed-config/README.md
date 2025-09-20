# Typed Configuration Example

This example demonstrates the new typed configuration system for Tailwind CSS utilities with enhanced type safety and structured design tokens.

## What This Example Shows

- **Typed configuration system** with structured theme access
- **Compile-time type safety** for design tokens
- **Structured design system** with named properties instead of string keys
- **Type-safe value creation** and conversion
- **Enhanced developer experience** with IDE autocomplete
- **Migration path** from string-based configuration

## Key Concepts Demonstrated

### 1. Typed Configuration Access

Instead of string-based map lookups, use structured property access:

```go
// Old approach (string-based, error-prone)
// color := theme.Colors["red-50"]  // Runtime error if key doesn't exist

// New approach (typed, compile-time safe)
config := tailwind.DefaultConfig()
color := config.Theme.Colors.Red50  // Guaranteed to exist at compile time
```

### 2. Structured Theme Properties

All theme properties are strongly typed with clear naming:

```go
config := tailwind.DefaultConfig()

// Color access with full type safety
red50 := config.Theme.Colors.Red50      // *StaticColor
blue500 := config.Theme.Colors.Blue500   // *StaticColor

// Spacing access with type guarantees
size1 := config.Theme.Spacing.Size1     // *StaticLength (0.25rem)
size4 := config.Theme.Spacing.Size4     // *StaticLength (1rem)

// Typography access with structured properties
textSm := config.Theme.FontSize.Sm.Size   // *StaticLength (0.875rem)
textXl := config.Theme.FontSize.Xl.Size   // *StaticLength (1.25rem)
```

### 3. Type-Safe Value Conversion

Typed values provide safe conversion to CSS values:

```go
config := tailwind.DefaultConfig()

// Convert color to CSS value
redColor := config.Theme.Colors.Red50
cssColor := redColor.ToCSSValue()  // Returns css.Color
fmt.Printf("CSS Color: %s\n", cssColor.String())  // "#fef2f2"

// Convert spacing to CSS value
spacing := config.Theme.Spacing.Size4
cssLength := spacing.ToCSSValue()  // Returns css.Length
fmt.Printf("CSS Length: %s\n", cssLength.String())  // "1rem"
```

### 4. Custom Value Creation

Create custom typed values with the same safety guarantees:

```go
// Create custom color with type safety
customColor := tailwind.ColorFromHex("brand-primary", "#3b82f6")
fmt.Printf("Custom color: %s -> %s\n", 
    customColor.Name, 
    customColor.ToCSSValue().String())  // "brand-primary -> #3b82f6"

// Create custom spacing with type safety
customSpacing := tailwind.LengthFromRem("custom-spacing", 1.5)
fmt.Printf("Custom spacing: %s -> %s\n", 
    customSpacing.Name, 
    customSpacing.ToCSSValue().String())  // "custom-spacing -> 1.5rem"
```

## Configuration Structure

### Color System Structure

```go
type ColorScale struct {
    // Grayscale
    Black   *StaticColor  // #000000
    White   *StaticColor  // #ffffff
    
    // Gray scale (50-900)
    Gray50  *StaticColor  // #f9fafb
    Gray100 *StaticColor  // #f3f4f6
    Gray200 *StaticColor  // #e5e7eb
    // ... through Gray900
    
    // Color families (each 50-900)
    Red50   *StaticColor  // #fef2f2
    Red100  *StaticColor  // #fee2e2
    // ... through Red900
    
    Blue50  *StaticColor  // #eff6ff
    Blue100 *StaticColor  // #dbeafe
    // ... through Blue900
    
    // Extended palette
    Slate50   *StaticColor
    Zinc50    *StaticColor
    Neutral50 *StaticColor
    Stone50   *StaticColor
    // ... etc
}
```

### Spacing System Structure

```go
type SpacingScale struct {
    Size0   *StaticLength  // 0
    Size1   *StaticLength  // 0.25rem (4px)
    Size2   *StaticLength  // 0.5rem (8px)
    Size3   *StaticLength  // 0.75rem (12px)
    Size4   *StaticLength  // 1rem (16px)
    Size5   *StaticLength  // 1.25rem (20px)
    Size6   *StaticLength  // 1.5rem (24px)
    Size8   *StaticLength  // 2rem (32px)
    Size10  *StaticLength  // 2.5rem (40px)
    Size12  *StaticLength  // 3rem (48px)
    Size16  *StaticLength  // 4rem (64px)
    Size20  *StaticLength  // 5rem (80px)
    Size24  *StaticLength  // 6rem (96px)
    Size32  *StaticLength  // 8rem (128px)
    Size40  *StaticLength  // 10rem (160px)
    Size48  *StaticLength  // 12rem (192px)
    Size56  *StaticLength  // 14rem (224px)
    Size64  *StaticLength  // 16rem (256px)
}
```

### Typography System Structure

```go
type FontSizeScale struct {
    Xs   FontSizeConfig  // 0.75rem / 1rem
    Sm   FontSizeConfig  // 0.875rem / 1.25rem
    Base FontSizeConfig  // 1rem / 1.5rem
    Lg   FontSizeConfig  // 1.125rem / 1.75rem
    Xl   FontSizeConfig  // 1.25rem / 1.75rem
    Xl2  FontSizeConfig  // 1.5rem / 2rem
    Xl3  FontSizeConfig  // 1.875rem / 2.25rem
    Xl4  FontSizeConfig  // 2.25rem / 2.5rem
    Xl5  FontSizeConfig  // 3rem / 1
    Xl6  FontSizeConfig  // 3.75rem / 1
    Xl7  FontSizeConfig  // 4.5rem / 1
    Xl8  FontSizeConfig  // 6rem / 1
    Xl9  FontSizeConfig  // 8rem / 1
}

type FontSizeConfig struct {
    Size       *StaticLength  // Font size
    LineHeight *StaticLength  // Line height
}
```

## Type Safety Benefits

### 1. Compile-Time Error Prevention

```go
// ❌ Old approach - potential runtime errors
// color := theme.Colors["red-500"]  // Typo in key name causes nil
// if color == nil {
//     // Handle missing color error at runtime
// }

// ✅ New approach - compile-time safety
config := tailwind.DefaultConfig()
color := config.Theme.Colors.Red500  // Typo would be caught at compile time
// color is guaranteed to be non-nil
```

### 2. IDE Autocomplete and IntelliSense

The typed system provides comprehensive IDE support:

```go
config := tailwind.DefaultConfig()

// IDE autocomplete shows all available colors
config.Theme.Colors.  
// Autocomplete displays: Red50, Red100, Red200... Blue50, Blue100... etc.

// IDE autocomplete shows all spacing options
config.Theme.Spacing.  
// Autocomplete displays: Size0, Size1, Size2, Size3, Size4... etc.

// IDE autocomplete shows font size options
config.Theme.FontSize.  
// Autocomplete displays: Xs, Sm, Base, Lg, Xl, Xl2... etc.
```

### 3. Refactoring Safety

Go's refactoring tools work seamlessly:

```go
// Safe renaming across entire codebase
config.Theme.Colors.Red500  // Can be safely renamed with IDE refactoring tools
```

## Demonstration Output

The example shows:

### Configuration Structure
```
📋 Configuration Structure:
✅ Colors are typed: config.Theme.Colors.Red50 = red-50
✅ Spacing is typed: config.Theme.Spacing.Size4 = size-4
✅ Font sizes are typed: config.Theme.FontSize.Xl.Size = xl-size
```

### Custom Value Creation
```
🔧 Custom Value Creation:
✅ Custom color: brand-primary -> #3b82f6
✅ Custom spacing: custom-spacing -> 1.5rem
```

### Type Safety Benefits
```
🎯 Type Safety Benefits:
✅ No more string-based map access like theme.Colors["red-50"]
✅ Compile-time safety with config.Theme.Colors.Red50
✅ IDE autocomplete and IntelliSense support
✅ Leverages generated CSS property types from cssgen
✅ Nested named properties instead of flat string keys
```

## Value Types

### StaticColor Type

```go
type StaticColor struct {
    Name  string    // Human-readable name
    Value css.Color // CSS color value
}

func (c *StaticColor) String() string {
    return c.Name
}

func (c *StaticColor) ToCSSValue() css.Color {
    return c.Value
}
```

### StaticLength Type

```go
type StaticLength struct {
    Name  string     // Human-readable name  
    Value css.Length // CSS length value
}

func (l *StaticLength) String() string {
    return l.Name
}

func (l *StaticLength) ToCSSValue() css.Length {
    return l.Value
}
```

## Advanced Usage Patterns

### 1. Design System Configuration

```go
func ConfigureDesignSystem() *tailwind.Config {
    config := tailwind.DefaultConfig()
    
    // Type-safe access to design tokens
    primaryColor := config.Theme.Colors.Blue500
    secondaryColor := config.Theme.Colors.Gray500
    accentColor := config.Theme.Colors.Indigo500
    
    baseSpacing := config.Theme.Spacing.Size4
    largeSpacing := config.Theme.Spacing.Size8
    
    // Use in component creation
    return config
}
```

### 2. Component-Based Styling

```go
func CreateTypedComponents(config *tailwind.Config) {
    // Type-safe component styling
    buttonPrimary := css.RuleSet(".btn-primary",
        css.Set(css.BackgroundColor, config.Theme.Colors.Blue500.ToCSSValue()),
        css.Set(css.ColorP, config.Theme.Colors.White.ToCSSValue()),
        css.Set(css.Padding, config.Theme.Spacing.Size2.ToCSSValue()),
    )
    
    cardBase := css.RuleSet(".card",
        css.Set(css.BackgroundColor, config.Theme.Colors.White.ToCSSValue()),
        css.Set(css.Padding, config.Theme.Spacing.Size6.ToCSSValue()),
        css.Set(css.BorderRadius, config.Theme.Spacing.Size2.ToCSSValue()),
    )
}
```

### 3. Theme Validation

```go
func ValidateDesignTokens(config *tailwind.Config) bool {
    // All properties are guaranteed to exist due to type system
    // No need for nil checks or error handling
    
    primaryExists := config.Theme.Colors.Blue500 != nil    // Always true
    spacingExists := config.Theme.Spacing.Size4 != nil     // Always true
    fontExists := config.Theme.FontSize.Base.Size != nil   // Always true
    
    return primaryExists && spacingExists && fontExists // Always true
}
```

## Migration Guide

### From String-Based Configuration

```go
// Old approach (string-based)
type OldTheme struct {
    Colors map[string]css.Color
    Spacing map[string]css.Length
}

oldTheme := OldTheme{
    Colors: map[string]css.Color{
        "red-500": css.Hex("#ef4444"),
        "blue-500": css.Hex("#3b82f6"),
    },
}

// Accessing values (error-prone)
red := oldTheme.Colors["red-500"]  // Could be nil
if red == nil {
    // Handle error
}

// New approach (typed)
config := tailwind.DefaultConfig()
red := config.Theme.Colors.Red500.ToCSSValue()  // Guaranteed safe
```

### Benefits of Migration

1. **Compile-Time Safety**: Catch configuration errors during compilation
2. **Better IDE Support**: Full autocomplete and navigation
3. **Refactoring Safety**: IDE tools can safely rename and refactor
4. **Self-Documenting**: Types serve as documentation
5. **Performance**: Direct property access vs map lookups
6. **Maintainability**: Structured configuration is easier to maintain

## Running the Example

```bash
cd examples/typed-config
go run main.go
```

## Expected Output

```
🎉 TypeSafe CSS - New Typed Configuration System Demo
==================================================

📋 Configuration Structure:
✅ Colors are typed: config.Theme.Colors.Red50 = red-50
✅ Spacing is typed: config.Theme.Spacing.Size4 = size-4
✅ Font sizes are typed: config.Theme.FontSize.Xl.Size = xl-size

🔧 Custom Value Creation:
✅ Custom color: brand-primary -> #3b82f6
✅ Custom spacing: custom-spacing -> 1.5rem

🎯 Type Safety Benefits:
✅ No more string-based map access like theme.Colors["red-50"]
✅ Compile-time safety with config.Theme.Colors.Red50
✅ IDE autocomplete and IntelliSense support
✅ Leverages generated CSS property types from cssgen
✅ Nested named properties instead of flat string keys

🚀 The typed configuration system is now ready!
📝 Next steps: Complete utility generator for all config types
```

## Use Cases

This pattern is ideal for:

- **Large applications** requiring configuration safety
- **Design systems** with structured theme definitions
- **Team development** where configuration consistency matters
- **Type-driven development** leveraging Go's type system
- **IDE-heavy workflows** benefiting from autocomplete and navigation

## Future Enhancements

The typed configuration enables:

1. **Automatic documentation** generation from type definitions
2. **Design system validation** at compile time
3. **Theme migration tools** with type safety
4. **Enhanced IDE plugins** leveraging type information
5. **Performance optimizations** through static analysis

## Next Steps

- Explore enhanced Tailwind features in [examples/enhanced-tailwind/](../enhanced-tailwind/)
- See utility generation in [examples/tailwind/](../tailwind/)
- Check out core API patterns in [examples/basic/](../basic/)
- Read the comprehensive [Tailwind Guide](../../docs/tailwind-guide.md)