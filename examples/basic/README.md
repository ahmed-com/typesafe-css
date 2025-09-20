# Basic Example

This example demonstrates the core TypeSafe CSS API for building stylesheets with type safety.

## What This Example Shows

- **Basic CSS rule creation** using `css.RuleSet()`
- **Type-safe value constructors** like `css.Px()`, `css.Hex()`, `css.Rem()`
- **Shorthand helpers** like `css.PadXY()` and `css.FlexCenter()`
- **Responsive design** with media queries using `css.AtRule`
- **Template integration** with Go's `html/template` package
- **CSS output formatting** with both compact and pretty-printed styles

## Key Concepts Demonstrated

### 1. Type-Safe Value Construction

```go
css.Set(css.ColorP, css.Hex("#0af"))              // Hex color
css.Set(css.Padding, css.PadXY(css.Px(8), css.Px(12)))  // Shorthand padding
css.Set(css.FontSize, css.Rem(0.875))             // Rem units
css.Set(css.BorderRadius, css.Px(4))              // Pixel units
```

### 2. Shorthand Helpers

```go
css.PadXY(css.Px(8), css.Px(12))                  // padding: 12px 8px
css.BorderShorthand(css.Px(1), css.BorderSolid, css.Hex("#ccc"))  // border: 1px solid #ccc
css.FlexCenter()...                                // display: flex; justify-content: center; align-items: center
```

### 3. Responsive Design with Media Queries

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

### 4. Template Integration

```go
htmlTmpl := `<!DOCTYPE html>
<html>
<head>
<style>
{{ . }}
</style>
</head>
<body>
<div class="container">
  <button class="btn">Click me</button>
</div>
</body>
</html>`

t := template.Must(template.New("page").Parse(htmlTmpl))
err := t.Execute(os.Stdout, css.PrettyCSS(sheet.Items...))
```

## Running the Example

```bash
cd examples/basic
go run main.go
```

## Expected Output

The example generates CSS for:
- A button with inline-block display, padding, colors, and border radius
- A flex container with center alignment
- Responsive behavior that makes the button full-width on mobile

The output includes both compact and pretty-printed CSS, plus a complete HTML template with embedded styles.

## Use Cases

This pattern is ideal for:
- **Component libraries** where you need consistent, reusable styles
- **Design systems** with specific typography and spacing scales
- **Server-side rendered applications** where CSS is generated at build time
- **Static site generators** where styles need to be type-safe and maintainable

## Next Steps

- Try modifying the colors or spacing values
- Add more responsive breakpoints
- Experiment with different shorthand helpers
- See [examples/generated/](../generated/) for using generated property types
- See [examples/tailwind/](../tailwind/) for utility-first CSS approach