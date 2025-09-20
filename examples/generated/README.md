# Generated API Example

This example demonstrates using the generated type-safe CSS properties and values alongside the core API.

## What This Example Shows

- **Generated type-safe properties** from the `cssgen` package
- **Mixing generated and core APIs** for comprehensive CSS authoring
- **Type safety benefits** of generated enums and setters
- **Template integration** with mixed API usage
- **Hover states and pseudo-selectors** using generated properties

## Key Concepts Demonstrated

### 1. Generated Type-Safe Properties

The `cssgen` package provides type-safe properties generated from CSS specifications:

```go
// Type-safe display property with enum values
cssgen.SetDisplay(cssgen.DisplayValFlex),
cssgen.SetDisplay(cssgen.DisplayValInlineBlock),

// Type-safe positioning
cssgen.SetPosition(cssgen.PositionValRelative),

// Type-safe flex properties
cssgen.SetFlexDirection(cssgen.FlexDirectionValColumn),
cssgen.SetJustifyContent(cssgen.JustifyContentValSpaceBetween),
cssgen.SetAlignItems(cssgen.AlignItemsValCenter),

// Type-safe cursor values
cssgen.SetCursor(cssgen.CursorValPointer),

// Type-safe text properties
cssgen.SetTextAlign(cssgen.TextAlignValCenter),
cssgen.SetFontWeight(cssgen.FontWeightValBold),
```

### 2. Benefits of Generated API

#### Compile-time Safety
```go
// ❌ This won't compile - prevents typos
cssgen.SetDisplay(cssgen.DisplayValFlx)  // Typo: "Flx" instead of "Flex"

// ✅ This is type-safe
cssgen.SetDisplay(cssgen.DisplayValFlex)
```

#### IDE Autocomplete
The generated enums provide full IDE autocomplete support for all valid CSS values.

#### Spec Compliance
Generated values are automatically updated from CSS specifications, ensuring accuracy.

### 3. Mixed API Usage

Combine generated properties with core API for comprehensive styling:

```go
css.RuleSet(".card",
    // Generated type-safe properties
    cssgen.SetDisplay(cssgen.DisplayValFlex),
    cssgen.SetFlexDirection(cssgen.FlexDirectionValColumn),
    cssgen.SetJustifyContent(cssgen.JustifyContentValSpaceBetween),
    cssgen.SetAlignItems(cssgen.AlignItemsValCenter),
    cssgen.SetPosition(cssgen.PositionValRelative),
    
    // Core API for values not yet generated
    css.Set(css.Width, css.Px(300)),
    css.Set(css.Height, css.Px(200)),
    css.Set(css.Padding, css.Px(16)),
    css.Set(css.BackgroundColor, css.Hex("#f0f0f0")),
    css.Set(css.BorderRadius, css.Px(8)),
)
```

### 4. Button Component Example

```go
buttonRule := css.RuleSet(".btn",
    // Generated properties for type safety
    cssgen.SetDisplay(cssgen.DisplayValInlineBlock),
    cssgen.SetCursor(cssgen.CursorValPointer),
    cssgen.SetTextAlign(cssgen.TextAlignValCenter),
    cssgen.SetFontWeight(cssgen.FontWeightValBold),
    
    // Core API for custom values
    css.Set(css.Padding, css.Raw("8px 16px")),
    css.Set(css.BackgroundColor, css.Hex("#007bff")),
    css.Set(css.ColorP, css.Hex("#ffffff")),
    css.Set(css.BorderRadius, css.Px(4)),
    css.Set(css.Border, css.Raw("none")),
)
```

### 5. Hover States

```go
buttonHoverRule := css.RuleSet(".btn:hover",
    css.Set(css.BackgroundColor, css.Hex("#0056b3")),
    cssgen.SetCursor(cssgen.CursorValPointer),
)
```

### 6. Generated vs Core API Comparison

| Feature | Generated API | Core API |
|---------|--------------|----------|
| Type safety | Full enum safety | Value constructor safety |
| Property names | Compile-time checked | Compile-time checked |
| Property values | Enum-based (limited) | Constructor-based (flexible) |
| IDE support | Full autocomplete | Partial autocomplete |
| Flexibility | Spec-defined values | Any valid CSS value |
| Use case | Standard properties | Custom/complex values |

### 7. When to Use Each API

**Use Generated API when:**
- Working with standard CSS properties
- You want maximum type safety
- You need IDE autocomplete for property values
- You're learning CSS and want guidance on valid values

**Use Core API when:**
- You need custom values (calculations, variables, etc.)
- Working with newer CSS features not yet in specs
- You need shorthand helpers
- You want maximum flexibility

**Combine Both when:**
- Building comprehensive stylesheets
- You want type safety where available, flexibility where needed
- Working with teams of varying CSS experience levels

## Code Structure

The example creates:

1. **Card component** using flexbox with generated properties
2. **Button component** with mixed APIs for comprehensive styling
3. **Hover effects** showing pseudo-selector support
4. **Complete stylesheet** combining all components
5. **HTML template** demonstrating real-world usage

## Generated Property Categories

The `cssgen` package includes:

### Display & Layout
```go
cssgen.DisplayValBlock
cssgen.DisplayValFlex
cssgen.DisplayValGrid
cssgen.DisplayValInline
cssgen.DisplayValInlineBlock
cssgen.DisplayValNone
```

### Positioning
```go
cssgen.PositionValStatic
cssgen.PositionValRelative
cssgen.PositionValAbsolute
cssgen.PositionValFixed
cssgen.PositionValSticky
```

### Flexbox
```go
cssgen.FlexDirectionValRow
cssgen.FlexDirectionValColumn
cssgen.JustifyContentValCenter
cssgen.JustifyContentValSpaceBetween
cssgen.AlignItemsValCenter
cssgen.AlignItemsValStretch
```

### Typography
```go
cssgen.FontWeightValNormal
cssgen.FontWeightValBold
cssgen.TextAlignValLeft
cssgen.TextAlignValCenter
cssgen.TextAlignValRight
```

### Interaction
```go
cssgen.CursorValPointer
cssgen.CursorValDefault
cssgen.CursorValHelp
cssgen.CursorValWait
```

## Running the Example

```bash
cd examples/generated
go run main.go
```

## Expected Output

The example generates:

1. **Generated CSS** showing type-safe properties in action
2. **HTML template** with embedded styles
3. **Card and button components** demonstrating real-world usage

## Use Cases

This pattern is ideal for:

- **Large applications** where type safety is crucial
- **Team environments** where CSS expertise varies
- **Design systems** requiring consistent property usage
- **Applications** that need both safety and flexibility
- **Learning environments** where guidance on valid CSS values helps

## Best Practices

### 1. Start with Generated API
Begin with generated properties for maximum safety:

```go
// Start with type-safe generated properties
css.RuleSet(".component",
    cssgen.SetDisplay(cssgen.DisplayValFlex),
    cssgen.SetPosition(cssgen.PositionValRelative),
)
```

### 2. Add Core API for Custom Values
Use core API for values not available in generated enums:

```go
css.RuleSet(".component",
    // Generated properties
    cssgen.SetDisplay(cssgen.DisplayValFlex),
    
    // Core API for custom values
    css.Set(css.BackgroundColor, css.Hex("#custom")),
    css.Set(css.Padding, css.PadXY(css.Px(8), css.Px(16))),
)
```

### 3. Organize by API Type
Group generated and core properties for clarity:

```go
css.RuleSet(".component",
    // Generated properties (type-safe enums)
    cssgen.SetDisplay(cssgen.DisplayValFlex),
    cssgen.SetJustifyContent(cssgen.JustifyContentValCenter),
    cssgen.SetAlignItems(cssgen.AlignItemsValCenter),
    
    // Core properties (custom values)
    css.Set(css.Width, css.Px(300)),
    css.Set(css.Height, css.Px(200)),
    css.Set(css.Padding, css.Px(16)),
)
```

## Advantages

1. **Maximum Type Safety**: Generated enums prevent invalid property values
2. **IDE Support**: Full autocomplete for property names and values
3. **Spec Compliance**: Always up-to-date with CSS specifications
4. **Learning Aid**: Discover valid CSS values through autocomplete
5. **Flexibility**: Core API available for custom requirements
6. **Migration Path**: Easy to adopt incrementally

## Next Steps

- Explore utility-first approach in [examples/tailwind/](../tailwind/)
- See core API fundamentals in [examples/basic/](../basic/)
- Check out MDN-validated properties in [examples/mdn-validation/](../mdn-validation/)
- Read the [Core API Guide](../../docs/core-api.md) for comprehensive documentation