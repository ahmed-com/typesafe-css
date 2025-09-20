# Best Practices Guide

Essential patterns and recommendations for building maintainable, performant CSS with TypeSafe CSS.

## Table of Contents

- [Code Organization](#code-organization)
- [Type Safety Patterns](#type-safety-patterns)
- [Performance Optimization](#performance-optimization)
- [Maintainability](#maintainability)
- [Team Development](#team-development)
- [Testing Strategies](#testing-strategies)
- [Common Pitfalls](#common-pitfalls)

## Code Organization

### 1. Layered Architecture

Organize your CSS generation in clear layers:

```go
// styles/
//   ├── base/          - Reset, typography, base styles
//   ├── components/    - Reusable UI components
//   ├── utilities/     - Utility classes
//   ├── layouts/       - Page layouts and grids
//   └── themes/        - Theme configurations

// base/typography.go
func CreateTypographyBase() css.Stylesheet {
    var stylesheet css.Stylesheet
    
    // Base typography styles
    stylesheet.Add(css.RuleSet("body",
        css.Set(css.FontFamily, css.Raw("system-ui, sans-serif")),
        css.Set(css.LineHeight, css.Raw("1.6")),
        css.Set(css.ColorP, css.Hex("#374151")),
    ))
    
    return stylesheet
}

// components/button.go
func CreateButtonComponent() css.Stylesheet {
    var stylesheet css.Stylesheet
    
    // Component-specific styles
    stylesheet.Add(css.RuleSet(".btn",
        css.Set(css.Display, css.DisplayInlineBlock),
        css.Set(css.Padding, css.PadXY(css.Px(12), css.Px(24))),
        css.Set(css.FontWeight, css.FontWeight600),
        css.Set(css.BorderRadius, css.Px(6)),
        css.Set(css.Cursor, css.CursorPointer),
        css.Set(css.Transition, css.Raw("all 0.2s ease")),
    ))
    
    return stylesheet
}

// themes/default.go
func DefaultTheme() *Theme {
    return &Theme{
        Colors: map[string]css.Color{
            "primary":   css.Hex("#3b82f6"),
            "secondary": css.Hex("#6b7280"),
        },
        Spacing: map[string]css.Length{
            "sm": css.Px(8),
            "md": css.Px(16),
            "lg": css.Px(24),
        },
    }
}
```

### 2. Package Structure

Structure your packages for clarity and reusability:

```go
// pkg/styles/
//   ├── config.go      - Configuration and themes
//   ├── base.go        - Base styles and resets
//   ├── components.go  - UI components
//   ├── utilities.go   - Utility classes
//   └── builder.go     - Main stylesheet builder

// pkg/styles/builder.go
type StylesheetBuilder struct {
    config *Config
    cache  map[string]css.Stylesheet
}

func NewBuilder(config *Config) *StylesheetBuilder {
    return &StylesheetBuilder{
        config: config,
        cache:  make(map[string]css.Stylesheet),
    }
}

func (b *StylesheetBuilder) Build() css.Stylesheet {
    var combined css.Stylesheet
    
    // Build in dependency order
    combined.Add(b.buildBase().Items...)
    combined.Add(b.buildComponents().Items...)
    combined.Add(b.buildUtilities().Items...)
    
    return combined
}
```

### 3. Component-Based Organization

Create self-contained components:

```go
// Component interface for consistency
type Component interface {
    CSS() css.Stylesheet
    Name() string
    Dependencies() []Component
}

// Button component
type Button struct {
    variant string
    size    string
    theme   *Theme
}

func NewButton(variant, size string, theme *Theme) *Button {
    return &Button{
        variant: variant,
        size:    size,
        theme:   theme,
    }
}

func (b *Button) CSS() css.Stylesheet {
    var stylesheet css.Stylesheet
    
    // Base button styles
    stylesheet.Add(b.baseStyles())
    
    // Variant-specific styles
    stylesheet.Add(b.variantStyles())
    
    // Size-specific styles
    stylesheet.Add(b.sizeStyles())
    
    return stylesheet
}

func (b *Button) Name() string {
    return fmt.Sprintf("button-%s-%s", b.variant, b.size)
}

func (b *Button) Dependencies() []Component {
    return []Component{} // No dependencies
}
```

## Type Safety Patterns

### 1. Value Object Pattern

Create strongly typed value objects:

```go
// Color value object
type Color struct {
    value css.Color
    name  string
}

func NewColor(name string, value css.Color) Color {
    return Color{name: name, value: value}
}

func (c Color) CSS() css.Color {
    return c.value
}

func (c Color) String() string {
    return c.name
}

// Length value object
type Length struct {
    value css.Length
    name  string
}

func NewLength(name string, value css.Length) Length {
    return Length{name: name, value: value}
}

func (l Length) CSS() css.Length {
    return l.value
}

// Usage
var (
    PrimaryBlue = NewColor("primary-blue", css.Hex("#3b82f6"))
    BaseSpacing = NewLength("base-spacing", css.Rem(1))
)

// Type-safe usage
css.Set(css.BackgroundColor, PrimaryBlue.CSS())
css.Set(css.Padding, BaseSpacing.CSS())
```

### 2. Builder Pattern for Complex Styles

Use builders for complex styling scenarios:

```go
type StyleBuilder struct {
    declarations []css.Declaration
}

func NewStyleBuilder() *StyleBuilder {
    return &StyleBuilder{
        declarations: make([]css.Declaration, 0),
    }
}

func (b *StyleBuilder) BackgroundColor(color css.Color) *StyleBuilder {
    b.declarations = append(b.declarations, css.Set(css.BackgroundColor, color))
    return b
}

func (b *StyleBuilder) Padding(padding css.Length) *StyleBuilder {
    b.declarations = append(b.declarations, css.Set(css.Padding, padding))
    return b
}

func (b *StyleBuilder) BorderRadius(radius css.Length) *StyleBuilder {
    b.declarations = append(b.declarations, css.Set(css.BorderRadius, radius))
    return b
}

func (b *StyleBuilder) Build(selector string) css.Rule {
    return css.RuleSet(selector, b.declarations...)
}

// Usage
button := NewStyleBuilder().
    BackgroundColor(css.Hex("#3b82f6")).
    Padding(css.Px(16)).
    BorderRadius(css.Px(6)).
    Build(".btn-primary")
```

### 3. Configuration Validation

Validate configurations at compile time:

```go
type Config struct {
    Colors  ColorPalette
    Spacing SpacingScale
    Fonts   FontScale
}

type ColorPalette struct {
    Primary   css.Color
    Secondary css.Color
    Success   css.Color
    Error     css.Color
}

type SpacingScale struct {
    XS css.Length
    SM css.Length
    MD css.Length
    LG css.Length
    XL css.Length
}

// Validation at initialization
func NewConfig() *Config {
    config := &Config{
        Colors: ColorPalette{
            Primary:   css.Hex("#3b82f6"),
            Secondary: css.Hex("#6b7280"),
            Success:   css.Hex("#10b981"),
            Error:     css.Hex("#ef4444"),
        },
        Spacing: SpacingScale{
            XS: css.Px(4),
            SM: css.Px(8),
            MD: css.Px(16),
            LG: css.Px(24),
            XL: css.Px(32),
        },
    }
    
    // Validation would happen here
    // All properties are guaranteed to be set due to struct initialization
    
    return config
}
```

## Performance Optimization

### 1. CSS Deduplication

Implement automatic deduplication:

```go
type StylesheetCache struct {
    rules map[string]css.Rule
    mu    sync.RWMutex
}

func NewStylesheetCache() *StylesheetCache {
    return &StylesheetCache{
        rules: make(map[string]css.Rule),
    }
}

func (c *StylesheetCache) AddRule(rule css.Rule) {
    c.mu.Lock()
    defer c.mu.Unlock()
    
    key := rule.String()
    if _, exists := c.rules[key]; !exists {
        c.rules[key] = rule
    }
}

func (c *StylesheetCache) GetStylesheet() css.Stylesheet {
    c.mu.RLock()
    defer c.mu.RUnlock()
    
    var stylesheet css.Stylesheet
    for _, rule := range c.rules {
        stylesheet.Add(rule)
    }
    
    return stylesheet
}
```

### 2. Lazy Generation

Generate styles only when needed:

```go
type LazyStylesheet struct {
    generator func() css.Stylesheet
    cached    *css.Stylesheet
    mu        sync.Once
}

func NewLazyStylesheet(generator func() css.Stylesheet) *LazyStylesheet {
    return &LazyStylesheet{
        generator: generator,
    }
}

func (l *LazyStylesheet) Get() css.Stylesheet {
    l.mu.Do(func() {
        stylesheet := l.generator()
        l.cached = &stylesheet
    })
    return *l.cached
}

// Usage
var buttonStyles = NewLazyStylesheet(func() css.Stylesheet {
    // Expensive style generation
    return generateButtonStyles()
})

// Only generated when first accessed
styles := buttonStyles.Get()
```

### 3. Build-Time Optimization

Generate critical CSS at build time:

```go
//go:generate go run ./cmd/generate-critical-css

package main

import (
    "os"
    "github.com/ahmed-com/typesafe-css/css"
)

func main() {
    // Generate critical CSS for above-the-fold content
    critical := generateCriticalCSS()
    
    // Generate full CSS for lazy loading
    full := generateFullCSS()
    
    // Write critical CSS inline
    err := os.WriteFile("assets/critical.css", []byte(critical.String()), 0644)
    if err != nil {
        panic(err)
    }
    
    // Write full CSS for lazy loading
    err = os.WriteFile("assets/full.css", []byte(full.String()), 0644)
    if err != nil {
        panic(err)
    }
}

func generateCriticalCSS() css.Stylesheet {
    var stylesheet css.Stylesheet
    
    // Only include styles needed for initial render
    stylesheet.Add(
        createResetStyles(),
        createHeaderStyles(),
        createHeroStyles(),
    )
    
    return stylesheet
}
```

## Maintainability

### 1. Version Your Styles

Track style versions for compatibility:

```go
type StylesheetVersion struct {
    Major int
    Minor int
    Patch int
}

func (v StylesheetVersion) String() string {
    return fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patch)
}

type VersionedStylesheet struct {
    Version    StylesheetVersion
    Stylesheet css.Stylesheet
    Comment    string
}

func NewVersionedStylesheet(version StylesheetVersion, comment string) *VersionedStylesheet {
    return &VersionedStylesheet{
        Version: version,
        Comment: comment,
    }
}

func (v *VersionedStylesheet) AddRule(rule css.Rule) {
    v.Stylesheet.Add(rule)
}

func (v *VersionedStylesheet) String() string {
    header := fmt.Sprintf("/* %s - %s */\n", v.Version.String(), v.Comment)
    return header + v.Stylesheet.String()
}
```

### 2. Documentation in Code

Document your styles with clear comments:

```go
// CreateButtonVariants generates all button style variants.
// Supports: primary, secondary, success, danger, warning
// Sizes: sm, md, lg
// States: default, hover, focus, disabled
func CreateButtonVariants(theme *Theme) css.Stylesheet {
    var stylesheet css.Stylesheet
    
    variants := map[string]css.Color{
        "primary":   theme.Colors.Primary,
        "secondary": theme.Colors.Secondary,
        "success":   theme.Colors.Success,
        "danger":    theme.Colors.Danger,
        "warning":   theme.Colors.Warning,
    }
    
    sizes := map[string]struct {
        padding  css.Length
        fontSize css.Length
    }{
        "sm": {css.Px(8), css.Rem(0.875)},
        "md": {css.Px(12), css.Rem(1)},
        "lg": {css.Px(16), css.Rem(1.125)},
    }
    
    for variant, color := range variants {
        for size, props := range sizes {
            rule := createButtonRule(variant, size, color, props.padding, props.fontSize)
            stylesheet.Add(rule)
        }
    }
    
    return stylesheet
}
```

### 3. Testable Style Generation

Make your styles testable:

```go
func TestButtonStyles(t *testing.T) {
    theme := &Theme{
        Colors: ColorPalette{
            Primary: css.Hex("#3b82f6"),
        },
    }
    
    stylesheet := CreateButtonVariants(theme)
    
    // Test that primary button exists
    cssOutput := stylesheet.String()
    if !strings.Contains(cssOutput, ".btn-primary") {
        t.Error("Primary button class not found")
    }
    
    // Test that primary color is used
    if !strings.Contains(cssOutput, "#3b82f6") {
        t.Error("Primary color not applied")
    }
}

func TestStyleOutput(t *testing.T) {
    rule := css.RuleSet(".test",
        css.Set(css.ColorP, css.Hex("#ff0000")),
        css.Set(css.Padding, css.Px(16)),
    )
    
    expected := ".test{color:#ff0000;padding:16px}"
    actual := rule.String()
    
    if actual != expected {
        t.Errorf("Expected %q, got %q", expected, actual)
    }
}
```

## Team Development

### 1. Style Guidelines

Establish clear team guidelines:

```go
// Style Guidelines:
// 1. Always use type-safe constructors (css.Px, css.Hex, etc.)
// 2. Prefer shorthand helpers where available
// 3. Use meaningful selector names
// 4. Group related styles together
// 5. Comment complex calculations

// ✅ Good
func CreateCard() css.Rule {
    return css.RuleSet(".card",
        // Layout
        css.Set(css.Display, css.DisplayBlock),
        css.Set(css.Width, css.Percent(100)),
        
        // Spacing
        css.Set(css.Padding, css.Px(24)),
        css.Set(css.Margin, css.PadXY(css.Px(16), css.Px(0))),
        
        // Appearance
        css.Set(css.BackgroundColor, css.Hex("#ffffff")),
        css.Set(css.BorderRadius, css.Px(8)),
        css.Set(css.BoxShadow, css.Raw("0 4px 6px -1px rgba(0, 0, 0, 0.1)")),
    )
}

// ❌ Avoid
func CreateBadCard() css.Rule {
    return css.RuleSet(".c",
        css.Set(css.Display, css.Raw("block")), // Use typed values
        css.Set(css.BackgroundColor, css.Raw("#fff")), // Use css.Hex
        css.Set(css.Padding, css.Raw("24px")), // Use css.Px
    )
}
```

### 2. Code Review Checklist

Create a checklist for style reviews:

```markdown
## CSS Code Review Checklist

### Type Safety
- [ ] All values use type-safe constructors
- [ ] No raw strings for standard CSS values
- [ ] Proper use of shorthand helpers

### Performance
- [ ] No duplicate rules
- [ ] Efficient selector usage
- [ ] Appropriate use of CSS variables

### Maintainability
- [ ] Clear, semantic selector names
- [ ] Logical grouping of declarations
- [ ] Appropriate comments for complex styles

### Consistency
- [ ] Follows team naming conventions
- [ ] Uses established theme tokens
- [ ] Consistent spacing and sizing
```

### 3. Shared Components

Create a shared component library:

```go
// pkg/design-system/

// Button component with all variants
type Button struct {
    Variant ButtonVariant
    Size    ButtonSize
    State   ButtonState
}

type ButtonVariant int
const (
    ButtonPrimary ButtonVariant = iota
    ButtonSecondary
    ButtonSuccess
    ButtonDanger
)

type ButtonSize int
const (
    ButtonSmall ButtonSize = iota
    ButtonMedium
    ButtonLarge
)

func (b Button) CSS() css.Rule {
    className := fmt.Sprintf(".btn-%s-%s", b.Variant.String(), b.Size.String())
    
    var declarations []css.Declaration
    declarations = append(declarations, b.baseStyles()...)
    declarations = append(declarations, b.variantStyles()...)
    declarations = append(declarations, b.sizeStyles()...)
    
    return css.RuleSet(className, declarations...)
}
```

## Testing Strategies

### 1. Unit Testing CSS Generation

Test individual style functions:

```go
func TestPrimaryButtonStyle(t *testing.T) {
    button := Button{
        Variant: ButtonPrimary,
        Size:    ButtonMedium,
    }
    
    rule := button.CSS()
    css := rule.String()
    
    // Test selector
    if !strings.Contains(css, ".btn-primary-medium") {
        t.Error("Wrong selector generated")
    }
    
    // Test background color
    if !strings.Contains(css, "background-color:#3b82f6") {
        t.Error("Primary color not applied")
    }
}
```

### 2. Golden File Testing

Test CSS output against golden files:

```go
func TestCSSOutput(t *testing.T) {
    stylesheet := GenerateTestStylesheet()
    actual := css.PrettyCSS(stylesheet.Items...)
    
    goldenFile := "testdata/expected.css"
    
    if *updateGolden {
        err := os.WriteFile(goldenFile, []byte(actual), 0644)
        if err != nil {
            t.Fatal(err)
        }
        return
    }
    
    expected, err := os.ReadFile(goldenFile)
    if err != nil {
        t.Fatal(err)
    }
    
    if actual != string(expected) {
        t.Errorf("CSS output mismatch.\nExpected:\n%s\nActual:\n%s", expected, actual)
    }
}
```

### 3. Integration Testing

Test complete style generation:

```go
func TestCompleteStylesheetGeneration(t *testing.T) {
    config := &Config{
        Theme:     DefaultTheme(),
        Features:  []string{"buttons", "forms", "navigation"},
        Minify:    false,
    }
    
    generator := NewStylesheetGenerator(config)
    stylesheet := generator.Generate()
    
    // Test that all expected components are present
    css := stylesheet.String()
    
    expectedClasses := []string{
        ".btn-primary",
        ".form-input",
        ".nav-link",
    }
    
    for _, class := range expectedClasses {
        if !strings.Contains(css, class) {
            t.Errorf("Expected class %s not found", class)
        }
    }
}
```

## Common Pitfalls

### 1. Avoid String Concatenation

```go
// ❌ Don't do this
color := "#" + colorHex
css.Set(css.BackgroundColor, css.Raw(color))

// ✅ Do this instead
css.Set(css.BackgroundColor, css.Hex(colorHex))
```

### 2. Don't Bypass Type Safety

```go
// ❌ Don't bypass the type system
css.Set(css.Width, css.Raw("100px"))

// ✅ Use typed constructors
css.Set(css.Width, css.Px(100))
```

### 3. Avoid Hardcoded Values

```go
// ❌ Hardcoded values
css.Set(css.Padding, css.Px(16))
css.Set(css.Margin, css.Px(16))

// ✅ Use theme tokens
css.Set(css.Padding, theme.Spacing.MD)
css.Set(css.Margin, theme.Spacing.MD)
```

### 4. Don't Ignore Performance

```go
// ❌ Generating styles in templates
func HandleRequest(w http.ResponseWriter, r *http.Request) {
    // Don't generate styles on every request
    styles := generateStyles()
    // ...
}

// ✅ Pre-generate or cache styles
var cachedStyles = generateStyles()

func HandleRequest(w http.ResponseWriter, r *http.Request) {
    // Use pre-generated styles
    styles := cachedStyles
    // ...
}
```

### 5. Don't Mix Abstraction Levels

```go
// ❌ Mixing core API with utilities randomly
css.RuleSet(".button",
    tailwind.BgBlue500(),              // Utility
    css.Set(css.Padding, css.Px(16)), // Core API
    tailwind.TextWhite(),              // Utility
    css.Set(css.FontWeight, css.FontWeight600), // Core API
)

// ✅ Be consistent within a rule
css.RuleSet(".button",
    // All utilities
    tailwind.BgBlue500(),
    tailwind.TextWhite(),
    tailwind.P4(),
    tailwind.FontSemibold(),
)

// Or all core API
css.RuleSet(".button",
    css.Set(css.BackgroundColor, css.Hex("#3b82f6")),
    css.Set(css.ColorP, css.Hex("#ffffff")),
    css.Set(css.Padding, css.Px(16)),
    css.Set(css.FontWeight, css.FontWeight600),
)
```

Following these best practices will help you build maintainable, performant, and type-safe CSS with TypeSafe CSS. The key is to embrace the type system, organize your code clearly, and leverage the safety guarantees that Go provides.