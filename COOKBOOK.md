# TypeSafe CSS Cookbook

A comprehensive guide to using TypeSafe CSS for building maintainable, type-safe stylesheets in Go applications.

## Table of Contents

- [Quick Start](#quick-start)
- [Core API Patterns](#core-api-patterns)
- [Tailwind Utilities](#tailwind-utilities)
- [Common Use Cases](#common-use-cases)
- [Integration Patterns](#integration-patterns)
- [Advanced Techniques](#advanced-techniques)
- [Performance Tips](#performance-tips)
- [Troubleshooting](#troubleshooting)

## Quick Start

### Basic CSS Generation

```go
package main

import (
    "fmt"
    "github.com/ahmed-com/typesafe-css/css"
)

func main() {
    // Create a simple button style
    button := css.RuleSet(".btn",
        css.Set(css.Display, css.DisplayInlineBlock),
        css.Set(css.Padding, css.PadXY(css.Px(12), css.Px(24))),
        css.Set(css.BackgroundColor, css.Hex("#3b82f6")),
        css.Set(css.ColorP, css.Hex("#ffffff")),
        css.Set(css.BorderRadius, css.Px(6)),
        css.Set(css.Border, css.Raw("none")),
        css.Set(css.Cursor, css.CursorPointer),
    )

    fmt.Println(css.PrettyCSS(button))
}
```

**Output:**
```css
.btn {
  display: inline-block;
  padding: 24px 12px;
  background-color: #3b82f6;
  color: #ffffff;
  border-radius: 6px;
  border: none;
  cursor: pointer;
}
```

### Tailwind-Style Utilities

```go
package main

import (
    "fmt"
    "github.com/ahmed-com/typesafe-css/css"
    "github.com/ahmed-com/typesafe-css/tailwind"
)

func main() {
    var stylesheet css.Stylesheet
    
    // Add utility classes
    stylesheet.Add(
        tailwind.BgBlue500(),      // .bg-blue-500 { background-color: #3b82f6; }
        tailwind.TextWhite(),      // .text-white { color: #ffffff; }
        tailwind.P6(),             // .p-6 { padding: 1.5rem; }
        tailwind.RoundedLg(),      // .rounded-lg { border-radius: 0.5rem; }
        tailwind.CursorPointer(),  // .cursor-pointer { cursor: pointer; }
    )

    fmt.Println(stylesheet.String())
}
```

## Core API Patterns

### 1. Building Component Styles

Create reusable component styles using the core API:

```go
func CreateCard() css.Rule {
    return css.RuleSet(".card",
        css.Set(css.BackgroundColor, css.Hex("#ffffff")),
        css.Set(css.BorderRadius, css.Px(8)),
        css.Set(css.Padding, css.Px(24)),
        css.Set(css.BoxShadow, css.Raw("0 4px 6px -1px rgba(0, 0, 0, 0.1)")),
        css.Set(css.Border, css.BorderShorthand(css.Px(1), css.BorderSolid, css.Hex("#e5e7eb"))),
    )
}

func CreateCardHeader() css.Rule {
    return css.RuleSet(".card-header",
        css.Set(css.FontSize, css.Rem(1.25)),
        css.Set(css.FontWeight, css.FontWeight700),
        css.Set(css.MarginBottom, css.Px(16)),
        css.Set(css.ColorP, css.Hex("#111827")),
    )
}

func CreateCardBody() css.Rule {
    return css.RuleSet(".card-body",
        css.Set(css.ColorP, css.Hex("#6b7280")),
        css.Set(css.LineHeight, css.Raw("1.6")),
    )
}
```

### 2. Responsive Design

Create responsive layouts using media queries:

```go
func CreateResponsiveGrid() []css.Item {
    return []css.Item{
        // Base grid
        css.RuleSet(".grid",
            css.Set(css.Display, css.DisplayGrid),
            css.Set(css.GridTemplateColumns, css.Raw("1fr")),
            css.Set(css.Gap, css.Px(16)),
        ),
        
        // Tablet breakpoint
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
        
        // Desktop breakpoint
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

### 3. CSS Variables and Custom Properties

Use CSS variables for theming:

```go
func CreateThemeVariables() css.Rule {
    return css.RuleSet(":root",
        css.Set("--color-primary", css.Hex("#3b82f6")),
        css.Set("--color-secondary", css.Hex("#6b7280")),
        css.Set("--color-success", css.Hex("#10b981")),
        css.Set("--color-danger", css.Hex("#ef4444")),
        css.Set("--spacing-sm", css.Px(8)),
        css.Set("--spacing-md", css.Px(16)),
        css.Set("--spacing-lg", css.Px(24)),
        css.Set("--border-radius", css.Px(6)),
    )
}

func CreateThemedButton() css.Rule {
    return css.RuleSet(".btn-primary",
        css.Set(css.BackgroundColor, css.Var("--color-primary")),
        css.Set(css.ColorP, css.Hex("#ffffff")),
        css.Set(css.Padding, css.Raw("var(--spacing-sm) var(--spacing-md)")),
        css.Set(css.BorderRadius, css.Var("--border-radius")),
        css.Set(css.Border, css.Raw("none")),
        css.Set(css.Cursor, css.CursorPointer),
    )
}
```

## Tailwind Utilities

### 1. Basic Utility Usage

The Tailwind package provides utility-first CSS authoring:

```go
func CreateUtilityCard() []css.Rule {
    return []css.Rule{
        tailwind.BgWhite(),
        tailwind.RoundedLg(),
        tailwind.P6(),
        tailwind.ShadowLg(),
        tailwind.BorderClass(""),
        tailwind.BorderGray200(),
    }
}

func CreateUtilityButton() []css.Rule {
    return []css.Rule{
        tailwind.BgBlue500(),
        tailwind.TextWhite(),
        tailwind.FontSemibold(),
        tailwind.Py2(),
        tailwind.Px4(),
        tailwind.RoundedMd(),
        tailwind.CursorPointer(),
        // Hover state
        tailwind.HoverBgBlue600(),
    }
}
```

### 2. Custom Theme Configuration

Create and use custom themes:

```go
func CreateCustomTheme() *tailwind.Theme {
    return tailwind.CustomTheme(func(theme *tailwind.Theme) {
        // Add custom colors
        theme.AddColor("brand", css.Hex("#6366f1"))
        theme.AddColor("accent", css.Hex("#ec4899"))
        
        // Add custom spacing
        theme.AddSpacing("xs", css.Rem(0.125))  // 2px
        theme.AddSpacing("xl", css.Rem(5))      // 80px
        
        // Add custom border radius
        theme.AddBorderRadius("huge", css.Rem(3)) // 48px
    })
}

func UseCustomTheme() {
    customTheme := CreateCustomTheme()
    manager := tailwind.NewUtilityManager(customTheme)
    
    // Use custom values
    brandBg := manager.Bg("brand")        // Uses custom brand color
    extraLargePadding := manager.P("xl")  // Uses custom XL spacing
}
```

### 3. Arbitrary Values

Use arbitrary values for one-off customizations:

```go
func CreateArbitraryStyles() []css.Rule {
    return []css.Rule{
        tailwind.Bg("#bada55"),              // Custom hex color
        tailwind.P("17px"),                  // Custom padding
        tailwind.Text("clamp(1rem, 4vw, 2rem)"), // Responsive text
        tailwind.W("calc(100% - 2rem)"),     // Calculated width
    }
}
```

### 4. Utility Deduplication

The Tailwind package automatically deduplicates utilities:

```go
func DemonstrateDedupe() {
    var stylesheet css.Stylesheet
    
    // These will all generate the same CSS rule only once
    stylesheet.Add(
        tailwind.P4(),  // First occurrence
        tailwind.P4(),  // Deduplicated
        tailwind.P4(),  // Deduplicated
    )
    
    // Result: Only one .p-4 { padding: 1rem; } rule
    fmt.Println(stylesheet.String())
}
```

## Common Use Cases

### 1. Form Styling

Create a complete form with type-safe styles:

```go
func CreateFormStyles() css.Stylesheet {
    var stylesheet css.Stylesheet
    
    // Form container
    stylesheet.Add(css.RuleSet(".form",
        css.Set(css.MaxWidth, css.Px(480)),
        css.Set(css.Margin, css.Raw("0 auto")),
        css.Set(css.Padding, css.Px(24)),
    ))
    
    // Form groups
    stylesheet.Add(css.RuleSet(".form-group",
        css.Set(css.MarginBottom, css.Px(20)),
    ))
    
    // Labels
    stylesheet.Add(css.RuleSet(".form-label",
        css.Set(css.Display, css.DisplayBlock),
        css.Set(css.FontWeight, css.FontWeight600),
        css.Set(css.MarginBottom, css.Px(8)),
        css.Set(css.ColorP, css.Hex("#374151")),
    ))
    
    // Input fields
    stylesheet.Add(css.RuleSet(".form-input",
        css.Set(css.Display, css.DisplayBlock),
        css.Set(css.Width, css.Percent(100)),
        css.Set(css.Padding, css.PadXY(css.Px(12), css.Px(16))),
        css.Set(css.Border, css.BorderShorthand(css.Px(1), css.BorderSolid, css.Hex("#d1d5db"))),
        css.Set(css.BorderRadius, css.Px(6)),
        css.Set(css.FontSize, css.Rem(1)),
        css.Set(css.LineHeight, css.Raw("1.5")),
    ))
    
    // Focus states
    stylesheet.Add(css.RuleSet(".form-input:focus",
        css.Set(css.Outline, css.Raw("none")),
        css.Set(css.BorderColor, css.Hex("#3b82f6")),
        css.Set(css.BoxShadow, css.Raw("0 0 0 3px rgba(59, 130, 246, 0.1)")),
    ))
    
    return stylesheet
}
```

### 2. Navigation Bar

Build a responsive navigation bar:

```go
func CreateNavigation() css.Stylesheet {
    var stylesheet css.Stylesheet
    
    // Navigation container
    stylesheet.Add(css.RuleSet(".nav",
        css.Set(css.Display, css.DisplayFlex),
        css.Set(css.JustifyContent, css.JustifyContentSpaceBetween),
        css.Set(css.AlignItems, css.AlignItemsCenter),
        css.Set(css.Padding, css.PadXY(css.Px(16), css.Px(24))),
        css.Set(css.BackgroundColor, css.Hex("#ffffff")),
        css.Set(css.BoxShadow, css.Raw("0 1px 3px 0 rgba(0, 0, 0, 0.1)")),
    ))
    
    // Logo
    stylesheet.Add(css.RuleSet(".nav-logo",
        css.Set(css.FontSize, css.Rem(1.5)),
        css.Set(css.FontWeight, css.FontWeight700),
        css.Set(css.ColorP, css.Hex("#111827")),
    ))
    
    // Navigation links
    stylesheet.Add(css.RuleSet(".nav-links",
        css.Set(css.Display, css.DisplayFlex),
        css.Set(css.Gap, css.Px(32)),
        css.Set(css.ListStyle, css.Raw("none")),
        css.Set(css.Margin, css.Raw("0")),
        css.Set(css.Padding, css.Raw("0")),
    ))
    
    stylesheet.Add(css.RuleSet(".nav-link",
        css.Set(css.ColorP, css.Hex("#6b7280")),
        css.Set(css.TextDecoration, css.Raw("none")),
        css.Set(css.FontWeight, css.FontWeight500),
        css.Set(css.Transition, css.Raw("color 0.2s ease")),
    ))
    
    stylesheet.Add(css.RuleSet(".nav-link:hover",
        css.Set(css.ColorP, css.Hex("#3b82f6")),
    ))
    
    // Mobile responsive
    stylesheet.Add(css.AtRule{
        Name:   "media",
        Params: "(max-width: 768px)",
        Body: []css.Item{
            css.RuleSet(".nav-links",
                css.Set(css.Display, css.DisplayNone),
            ),
        },
    })
    
    return stylesheet
}
```

### 3. Card Grid Layout

Create a responsive card grid:

```go
func CreateCardGrid() css.Stylesheet {
    var stylesheet css.Stylesheet
    
    // Grid container
    stylesheet.Add(css.RuleSet(".card-grid",
        css.Set(css.Display, css.DisplayGrid),
        css.Set(css.GridTemplateColumns, css.Raw("repeat(auto-fit, minmax(300px, 1fr))")),
        css.Set(css.Gap, css.Px(24)),
        css.Set(css.Padding, css.Px(24)),
    ))
    
    // Individual cards
    stylesheet.Add(css.RuleSet(".card",
        css.Set(css.BackgroundColor, css.Hex("#ffffff")),
        css.Set(css.BorderRadius, css.Px(12)),
        css.Set(css.Padding, css.Px(24)),
        css.Set(css.BoxShadow, css.Raw("0 4px 6px -1px rgba(0, 0, 0, 0.1)")),
        css.Set(css.Border, css.BorderShorthand(css.Px(1), css.BorderSolid, css.Hex("#e5e7eb"))),
        css.Set(css.Transition, css.Raw("transform 0.2s ease, box-shadow 0.2s ease")),
    ))
    
    // Hover effects
    stylesheet.Add(css.RuleSet(".card:hover",
        css.Set(css.Transform, css.Raw("translateY(-4px)")),
        css.Set(css.BoxShadow, css.Raw("0 10px 25px -5px rgba(0, 0, 0, 0.1)")),
    ))
    
    return stylesheet
}
```

### 4. Button Variants

Create a complete button system:

```go
func CreateButtonSystem() css.Stylesheet {
    var stylesheet css.Stylesheet
    
    // Base button styles
    baseButton := css.RuleSet(".btn",
        css.Set(css.Display, css.DisplayInlineBlock),
        css.Set(css.Padding, css.PadXY(css.Px(12), css.Px(24))),
        css.Set(css.FontSize, css.Rem(1)),
        css.Set(css.FontWeight, css.FontWeight600),
        css.Set(css.BorderRadius, css.Px(6)),
        css.Set(css.Border, css.Raw("none")),
        css.Set(css.Cursor, css.CursorPointer),
        css.Set(css.TextDecoration, css.Raw("none")),
        css.Set(css.Transition, css.Raw("all 0.2s ease")),
        css.Set(css.TextAlign, css.TextAlignCenter),
    )
    
    // Primary button
    primaryButton := css.RuleSet(".btn-primary",
        css.Set(css.BackgroundColor, css.Hex("#3b82f6")),
        css.Set(css.ColorP, css.Hex("#ffffff")),
    )
    
    primaryButtonHover := css.RuleSet(".btn-primary:hover",
        css.Set(css.BackgroundColor, css.Hex("#2563eb")),
        css.Set(css.Transform, css.Raw("translateY(-1px)")),
        css.Set(css.BoxShadow, css.Raw("0 4px 12px rgba(59, 130, 246, 0.4)")),
    )
    
    // Secondary button
    secondaryButton := css.RuleSet(".btn-secondary",
        css.Set(css.BackgroundColor, css.Hex("#f3f4f6")),
        css.Set(css.ColorP, css.Hex("#374151")),
    )
    
    secondaryButtonHover := css.RuleSet(".btn-secondary:hover",
        css.Set(css.BackgroundColor, css.Hex("#e5e7eb")),
    )
    
    // Danger button
    dangerButton := css.RuleSet(".btn-danger",
        css.Set(css.BackgroundColor, css.Hex("#ef4444")),
        css.Set(css.ColorP, css.Hex("#ffffff")),
    )
    
    dangerButtonHover := css.RuleSet(".btn-danger:hover",
        css.Set(css.BackgroundColor, css.Hex("#dc2626")),
    )
    
    stylesheet.Add(
        baseButton,
        primaryButton, primaryButtonHover,
        secondaryButton, secondaryButtonHover,
        dangerButton, dangerButtonHover,
    )
    
    return stylesheet
}
```

### 5. Dark Mode Support

Implement dark mode with CSS variables:

```go
func CreateDarkModeStyles() css.Stylesheet {
    var stylesheet css.Stylesheet
    
    // Light mode variables (default)
    lightMode := css.RuleSet(":root",
        css.Set("--bg-primary", css.Hex("#ffffff")),
        css.Set("--bg-secondary", css.Hex("#f9fafb")),
        css.Set("--text-primary", css.Hex("#111827")),
        css.Set("--text-secondary", css.Hex("#6b7280")),
        css.Set("--border-color", css.Hex("#e5e7eb")),
    )
    
    // Dark mode variables
    darkMode := css.RuleSet("[data-theme=\"dark\"]",
        css.Set("--bg-primary", css.Hex("#1f2937")),
        css.Set("--bg-secondary", css.Hex("#111827")),
        css.Set("--text-primary", css.Hex("#f9fafb")),
        css.Set("--text-secondary", css.Hex("#d1d5db")),
        css.Set("--border-color", css.Hex("#374151")),
    )
    
    // Components using variables
    card := css.RuleSet(".card",
        css.Set(css.BackgroundColor, css.Var("--bg-primary")),
        css.Set(css.ColorP, css.Var("--text-primary")),
        css.Set(css.BorderColor, css.Var("--border-color")),
        css.Set(css.Transition, css.Raw("all 0.2s ease")),
    )
    
    stylesheet.Add(lightMode, darkMode, card)
    return stylesheet
}
```

## Integration Patterns

### 1. HTML Template Integration

Use with Go's `html/template` package:

```go
package main

import (
    "html/template"
    "os"
    "github.com/ahmed-com/typesafe-css/css"
    "github.com/ahmed-com/typesafe-css/tailwind"
)

func main() {
    // Generate CSS
    var stylesheet css.Stylesheet
    stylesheet.Add(
        tailwind.BgBlue500(),
        tailwind.TextWhite(),
        tailwind.P4(),
        tailwind.RoundedLg(),
    )
    
    // HTML template
    tmpl := template.Must(template.New("page").Parse(`
<!DOCTYPE html>
<html>
<head>
    <style>{{.CSS}}</style>
</head>
<body>
    <div class="bg-blue-500 text-white p-4 rounded-lg">
        Hello, TypeSafe CSS!
    </div>
</body>
</html>
    `))
    
    // Execute template
    data := struct {
        CSS template.CSS
    }{
        CSS: template.CSS(stylesheet.String()),
    }
    
    tmpl.Execute(os.Stdout, data)
}
```

### 2. Web Framework Integration

Example with a popular Go web framework:

```go
// Example with Gin framework
func setupRoutes() {
    r := gin.Default()
    
    // Generate CSS once at startup
    stylesheet := generateCSS()
    cssContent := stylesheet.String()
    
    // Serve CSS
    r.GET("/styles.css", func(c *gin.Context) {
        c.Header("Content-Type", "text/css")
        c.String(200, cssContent)
    })
    
    // Serve pages
    r.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", gin.H{
            "title": "TypeSafe CSS App",
        })
    })
}

func generateCSS() css.Stylesheet {
    var stylesheet css.Stylesheet
    
    // Add your styles here
    stylesheet.Add(
        CreateFormStyles().Items...,
        CreateNavigation().Items...,
        CreateButtonSystem().Items...,
    )
    
    return stylesheet
}
```

### 3. Static Site Generation

Generate CSS files for static sites:

```go
func generateStaticCSS() error {
    var stylesheet css.Stylesheet
    
    // Add all your styles
    stylesheet.Add(
        CreateFormStyles().Items...,
        CreateNavigation().Items...,
        CreateCardGrid().Items...,
    )
    
    // Write to file
    cssContent := css.PrettyCSS(stylesheet.Items...)
    return os.WriteFile("public/styles.css", []byte(cssContent), 0644)
}
```

### 4. Build-Time CSS Generation

Integrate with Go build process:

```go
//go:generate go run generate-css.go

package main

import (
    "fmt"
    "os"
    "github.com/ahmed-com/typesafe-css/css"
    "github.com/ahmed-com/typesafe-css/tailwind"
)

func main() {
    // Generate utility CSS
    utilityCSS := tailwind.GenerateUtilityStylesheet()
    
    // Generate component CSS
    componentCSS := generateComponentCSS()
    
    // Combine and write
    var combined css.Stylesheet
    combined.Add(utilityCSS.Items...)
    combined.Add(componentCSS.Items...)
    
    err := os.WriteFile("dist/styles.css", []byte(combined.String()), 0644)
    if err != nil {
        fmt.Printf("Error generating CSS: %v\n", err)
        os.Exit(1)
    }
    
    fmt.Println("CSS generated successfully!")
}
```

## Advanced Techniques

### 1. CSS-in-Go Components

Create reusable style components:

```go
type ButtonProps struct {
    Variant string
    Size    string
    Color   string
}

func Button(props ButtonProps) css.Rule {
    declarations := []css.Declaration{
        css.Set(css.Display, css.DisplayInlineBlock),
        css.Set(css.Border, css.Raw("none")),
        css.Set(css.Cursor, css.CursorPointer),
        css.Set(css.FontWeight, css.FontWeight600),
        css.Set(css.Transition, css.Raw("all 0.2s ease")),
    }
    
    // Add variant styles
    switch props.Variant {
    case "primary":
        declarations = append(declarations,
            css.Set(css.BackgroundColor, css.Hex("#3b82f6")),
            css.Set(css.ColorP, css.Hex("#ffffff")),
        )
    case "secondary":
        declarations = append(declarations,
            css.Set(css.BackgroundColor, css.Hex("#f3f4f6")),
            css.Set(css.ColorP, css.Hex("#374151")),
        )
    }
    
    // Add size styles
    switch props.Size {
    case "sm":
        declarations = append(declarations,
            css.Set(css.Padding, css.PadXY(css.Px(8), css.Px(16))),
            css.Set(css.FontSize, css.Rem(0.875)),
        )
    case "lg":
        declarations = append(declarations,
            css.Set(css.Padding, css.PadXY(css.Px(16), css.Px(32))),
            css.Set(css.FontSize, css.Rem(1.125)),
        )
    default: // medium
        declarations = append(declarations,
            css.Set(css.Padding, css.PadXY(css.Px(12), css.Px(24))),
            css.Set(css.FontSize, css.Rem(1)),
        )
    }
    
    className := fmt.Sprintf(".btn-%s-%s", props.Variant, props.Size)
    return css.RuleSet(className, declarations...)
}
```

### 2. Conditional Styles

Apply styles based on conditions:

```go
func ConditionalStyles(isDark bool, isMobile bool) []css.Rule {
    var rules []css.Rule
    
    if isDark {
        rules = append(rules, css.RuleSet(".container",
            css.Set(css.BackgroundColor, css.Hex("#1f2937")),
            css.Set(css.ColorP, css.Hex("#f9fafb")),
        ))
    } else {
        rules = append(rules, css.RuleSet(".container",
            css.Set(css.BackgroundColor, css.Hex("#ffffff")),
            css.Set(css.ColorP, css.Hex("#111827")),
        ))
    }
    
    if isMobile {
        rules = append(rules, css.RuleSet(".container",
            css.Set(css.Padding, css.Px(16)),
        ))
    } else {
        rules = append(rules, css.RuleSet(".container",
            css.Set(css.Padding, css.Px(32)),
        ))
    }
    
    return rules
}
```

### 3. Style Composition

Compose styles from multiple sources:

```go
func ComposeStyles(baseStyles, themeStyles, responsiveStyles []css.Declaration) css.Rule {
    var allDeclarations []css.Declaration
    
    allDeclarations = append(allDeclarations, baseStyles...)
    allDeclarations = append(allDeclarations, themeStyles...)
    allDeclarations = append(allDeclarations, responsiveStyles...)
    
    return css.RuleSet(".composed", allDeclarations...)
}

// Usage
func CreateComposedButton() css.Rule {
    base := []css.Declaration{
        css.Set(css.Display, css.DisplayInlineBlock),
        css.Set(css.Cursor, css.CursorPointer),
    }
    
    theme := []css.Declaration{
        css.Set(css.BackgroundColor, css.Var("--color-primary")),
        css.Set(css.ColorP, css.Var("--color-on-primary")),
    }
    
    responsive := []css.Declaration{
        css.Set(css.Padding, css.PadXY(css.Px(8), css.Px(16))),
        css.Set(css.FontSize, css.Rem(0.875)),
    }
    
    return ComposeStyles(base, theme, responsive)
}
```

## Performance Tips

### 1. CSS Deduplication

Automatically deduplicate repeated styles:

```go
func DeduplicateStylesheet(stylesheet css.Stylesheet) css.Stylesheet {
    seen := make(map[string]bool)
    var deduplicated css.Stylesheet
    
    for _, item := range stylesheet.Items {
        cssString := item.String()
        if !seen[cssString] {
            seen[cssString] = true
            deduplicated.Add(item)
        }
    }
    
    return deduplicated
}
```

### 2. Lazy Style Generation

Generate styles only when needed:

```go
type StyleCache struct {
    cache map[string]css.Rule
    mu    sync.RWMutex
}

func (sc *StyleCache) GetStyle(key string, generator func() css.Rule) css.Rule {
    sc.mu.RLock()
    if rule, exists := sc.cache[key]; exists {
        sc.mu.RUnlock()
        return rule
    }
    sc.mu.RUnlock()
    
    sc.mu.Lock()
    defer sc.mu.Unlock()
    
    // Double check after acquiring write lock
    if rule, exists := sc.cache[key]; exists {
        return rule
    }
    
    rule := generator()
    if sc.cache == nil {
        sc.cache = make(map[string]css.Rule)
    }
    sc.cache[key] = rule
    return rule
}
```

### 3. Minified CSS Output

Generate minified CSS for production:

```go
func MinifyCSS(items ...css.Item) string {
    // Use the compact String() method instead of PrettyCSS
    var result strings.Builder
    for _, item := range items {
        result.WriteString(item.String())
    }
    return result.String()
}
```

## Troubleshooting

### Common Issues and Solutions

#### 1. Missing Imports
```go
// ❌ Wrong
import "github.com/ahmed-com/typesafe-css"

// ✅ Correct
import "github.com/ahmed-com/typesafe-css/css"
import "github.com/ahmed-com/typesafe-css/tailwind"
```

#### 2. Type Errors
```go
// ❌ Wrong - using strings directly
css.Set(css.Width, "100px")

// ✅ Correct - using typed constructors
css.Set(css.Width, css.Px(100))
```

#### 3. Template Integration Issues
```go
// ❌ Wrong - not using template.CSS
tmpl.Execute(w, stylesheet.String())

// ✅ Correct - using template.CSS for safety
tmpl.Execute(w, template.CSS(stylesheet.String()))
```

#### 4. Tailwind Utility Naming
```go
// ❌ Wrong - incorrect function names
tailwind.BackgroundBlue500()
tailwind.PaddingFour()

// ✅ Correct - actual function names
tailwind.BgBlue500()
tailwind.P4()
```

### Debugging Tips

1. **Use PrettyCSS for debugging**: When debugging, use `css.PrettyCSS()` to get readable output
2. **Print intermediate results**: Use `fmt.Println()` to inspect generated CSS at each step
3. **Validate with browser dev tools**: Copy generated CSS into browser dev tools to test
4. **Check for duplicate styles**: Use deduplication helpers to identify repeated styles

### Getting Help

- Check the [examples](examples/) directory for working code samples
- Review the [API documentation](https://pkg.go.dev/github.com/ahmed-com/typesafe-css)
- Look at the [test files](css/css_test.go) for usage patterns
- Open an issue on GitHub for bugs or feature requests

---

This cookbook covers the most common patterns and use cases for TypeSafe CSS. For more advanced usage and the latest features, refer to the official documentation and examples in the repository.