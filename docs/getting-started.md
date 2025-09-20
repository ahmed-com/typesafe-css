# Getting Started with TypeSafe CSS

Get up and running with TypeSafe CSS in under 5 minutes.

## Installation

```bash
go get github.com/ahmed-com/typesafe-css
```

## Your First Stylesheet

Create a new Go file and import the CSS package:

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
        css.Set(css.Padding, css.PadXY(css.Px(12), css.Px(20))),
        css.Set(css.BackgroundColor, css.Hex("#3b82f6")),
        css.Set(css.ColorP, css.Hex("#ffffff")),
        css.Set(css.BorderRadius, css.Px(6)),
        css.Set(css.Border, css.Raw("none")),
        css.Set(css.Cursor, css.CursorPointer),
    )
    
    // Print the CSS
    fmt.Println(css.PrettyCSS(button))
}
```

Run it:
```bash
go run main.go
```

Output:
```css
.btn {
  display: inline-block;
  padding: 20px 12px;
  background-color: #3b82f6;
  color: #ffffff;
  border-radius: 6px;
  border: none;
  cursor: pointer;
}
```

## Key Concepts

### 1. Type Safety
TypeSafe CSS prevents common CSS errors through Go's type system:

```go
// ❌ This won't compile - type safety prevents errors
css.Set(css.Width, "not-a-valid-length")

// ✅ This is type-safe and will work
css.Set(css.Width, css.Px(100))
```

### 2. Value Constructors
Use typed constructors for CSS values:

```go
css.Px(16)           // "16px"
css.Rem(1.5)         // "1.5rem"
css.Percent(50)      // "50%"
css.Hex("#ff0000")   // "#ff0000"
css.RGB(255, 0, 0)   // "rgb(255, 0, 0)"
css.Raw("calc(100% - 20px)") // For complex expressions
```

### 3. Shorthand Helpers
Common patterns have convenient helpers:

```go
// Padding shortcuts
css.PadXY(css.Px(8), css.Px(16))  // "16px 8px"
css.PadAll(css.Px(12))            // "12px"

// Flexbox helpers
css.FlexCenter()...  // display: flex; justify-content: center; align-items: center;

// Border shortcuts
css.BorderShorthand(css.Px(1), css.BorderSolid, css.Hex("#ccc"))
```

## Adding More Styles

Create multiple rules and combine them:

```go
func main() {
    // Multiple rules
    button := css.RuleSet(".btn", /* ... */)
    card := css.RuleSet(".card", /* ... */)
    
    // Create a stylesheet
    var stylesheet css.Stylesheet
    stylesheet.Add(button, card)
    
    // Output all CSS
    fmt.Println(stylesheet.String())
}
```

## Responsive Design

Add media queries for responsive design:

```go
// Desktop styles
desktop := css.RuleSet(".container",
    css.Set(css.MaxWidth, css.Px(1200)),
    css.Set(css.Margin, css.Raw("0 auto")),
)

// Mobile styles
mobile := css.AtRule{
    Name:   "media",
    Params: "(max-width: 768px)",
    Body: []css.Item{
        css.RuleSet(".container",
            css.Set(css.Padding, css.Px(16)),
        ),
    },
}

var stylesheet css.Stylesheet
stylesheet.Add(desktop, mobile)
```

## Using with HTML Templates

Integrate with Go's `html/template`:

```go
package main

import (
    "html/template"
    "os"
    "github.com/ahmed-com/typesafe-css/css"
)

func main() {
    // Generate CSS
    button := css.RuleSet(".btn", /* styles here */)
    
    // HTML template
    tmpl := template.Must(template.New("page").Parse(`
<!DOCTYPE html>
<html>
<head>
    <style>{{.CSS}}</style>
</head>
<body>
    <button class="btn">Click me</button>
</body>
</html>
    `))
    
    // Execute template
    data := struct {
        CSS template.CSS
    }{
        CSS: template.CSS(css.PrettyCSS(button)),
    }
    
    tmpl.Execute(os.Stdout, data)
}
```

## Using Tailwind Utilities

For utility-first CSS, use the Tailwind package:

```go
import "github.com/ahmed-com/typesafe-css/tailwind"

func main() {
    var stylesheet css.Stylesheet
    
    // Add utility classes
    stylesheet.Add(
        tailwind.BgBlue500(),    // .bg-blue-500 { background-color: #3b82f6; }
        tailwind.TextWhite(),    // .text-white { color: #ffffff; }
        tailwind.P4(),           // .p-4 { padding: 1rem; }
        tailwind.RoundedLg(),    // .rounded-lg { border-radius: 0.5rem; }
    )
    
    fmt.Println(stylesheet.String())
}
```

## Next Steps

Now that you have the basics, explore:

- [Tailwind Utilities Guide](tailwind-guide.md) - Learn the utility-first approach
- [Core API Guide](core-api.md) - Deep dive into the core API
- [Template Integration](template-integration.md) - Advanced template patterns
- [COOKBOOK.md](../COOKBOOK.md) - Practical recipes and examples

## Common Patterns

### Component Styling
```go
func CreateButton(variant string) css.Rule {
    base := []css.Declaration{
        css.Set(css.Display, css.DisplayInlineBlock),
        css.Set(css.Padding, css.PadXY(css.Px(12), css.Px(20))),
        css.Set(css.BorderRadius, css.Px(6)),
        css.Set(css.Border, css.Raw("none")),
        css.Set(css.Cursor, css.CursorPointer),
    }
    
    switch variant {
    case "primary":
        base = append(base,
            css.Set(css.BackgroundColor, css.Hex("#3b82f6")),
            css.Set(css.ColorP, css.Hex("#ffffff")),
        )
    case "secondary":
        base = append(base,
            css.Set(css.BackgroundColor, css.Hex("#f3f4f6")),
            css.Set(css.ColorP, css.Hex("#374151")),
        )
    }
    
    return css.RuleSet(fmt.Sprintf(".btn-%s", variant), base...)
}
```

### Form Styling
```go
func CreateFormStyles() []css.Rule {
    return []css.Rule{
        css.RuleSet(".form-input",
            css.Set(css.Display, css.DisplayBlock),
            css.Set(css.Width, css.Percent(100)),
            css.Set(css.Padding, css.Px(12)),
            css.Set(css.Border, css.BorderShorthand(css.Px(1), css.BorderSolid, css.Hex("#d1d5db"))),
            css.Set(css.BorderRadius, css.Px(4)),
        ),
        css.RuleSet(".form-input:focus",
            css.Set(css.Outline, css.Raw("none")),
            css.Set(css.BorderColor, css.Hex("#3b82f6")),
            css.Set(css.BoxShadow, css.Raw("0 0 0 3px rgba(59, 130, 246, 0.1)")),
        ),
    }
}
```

You're now ready to build type-safe, maintainable CSS with Go!