# TypeSafe CSS

A type-safe CSS authoring library for Go that eliminates stringly-typed CSS generation footguns.

## Features

- **Type-safe CSS APIs** - No more string concatenation or typos in CSS properties
- **Compact and readable output** - Generates clean, valid CSS
- **Template integration** - Works seamlessly with Go's `html/template`
- **Shorthand helpers** - Convenient functions for common CSS patterns
- **Zero runtime dependencies** - Pure Go implementation

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/ahmed-com/typesafe-css/css"
)

func main() {
    // Create type-safe CSS rules
    btn := css.RuleSet(".btn",
        css.Set(css.Display, css.DisplayInline),
        css.Set(css.ColorP, css.Hex("#0af")),
        css.Set(css.Padding, css.PadXY(css.Px(8), css.Px(12))),
        css.Set(css.FontSize, css.Rem(0.875)),
        css.Set(css.BorderRadius, css.Px(4)),
    )

    // Output CSS
    fmt.Println(css.PrettyCSS(btn))
}
```

Output:
```css
.btn {
  display:inline;
  color:#0af;
  padding:12px 8px;
  font-size:0.875rem;
  border-radius:4px;
}
```

## Core Types

### Values
- `css.Keyword` - CSS keywords like "block", "flex", "auto"
- `css.Length` - Lengths with units (px, rem, em, %)
- `css.Color` - Colors (hex, rgb, rgba)
- `css.Raw` - Raw CSS values for complex expressions

### Constructors
- `css.Px(16)` → "16px"
- `css.Rem(1.5)` → "1.5rem"
- `css.Hex("#ff0000")` → "#ff0000"
- `css.RGB(255, 0, 0)` → "rgb(255 0 0)"

### Structure
- `css.Property` - CSS property names
- `css.Decl` - Property-value pairs
- `css.Rule` - Selector + declarations
- `css.AtRule` - @media, @keyframes, etc.
- `css.Stylesheet` - Collection of rules

## Advanced Usage

### Responsive Design
```go
mobile := css.AtRule{
    Name:   "media",
    Params: "(max-width: 640px)",
    Body: []css.Item{
        css.RuleSet(".btn", 
            css.Set(css.Display, css.DisplayBlock),
            css.Set(css.Width, css.Percent(100)),
        ),
    },
}
```

### Flexbox Helpers
```go
container := css.RuleSet(".container",
    css.FlexCenter()..., // display: flex; justify-content: center; align-items: center;
)
```

### Template Integration
```go
htmlTemplate := `<style>{{ . }}</style>`
t := template.Must(template.New("styles").Parse(htmlTemplate))
t.Execute(w, css.PrettyCSS(rules...))
```

## Project Status

This project follows a comprehensive [roadmap](roadmap.md) for building a complete type-safe CSS solution. Current progress is tracked in [todo.md](todo.md).

**Current Status**: ✅ Phase 1 Complete - Core API implemented and tested

### Completed
- ✅ Core value types and constructors
- ✅ Property and rule structures  
- ✅ CSS serialization (compact and pretty-printed)
- ✅ Shorthand helpers for common patterns
- ✅ Template integration
- ✅ Comprehensive tests

### Next Steps
- Code generation pipeline for comprehensive CSS property coverage
- Advanced validation and linting capabilities
- Performance optimizations

## Installation

```bash
go get github.com/ahmed-com/typesafe-css
```

## Examples

See the [examples](examples/) directory for complete working examples.

## License

MIT License - see [LICENSE](LICENSE) for details.