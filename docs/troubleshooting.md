# Troubleshooting Guide

Common issues and solutions when working with TypeSafe CSS.

## Table of Contents

- [Installation Issues](#installation-issues)
- [Compilation Errors](#compilation-errors)
- [Runtime Issues](#runtime-issues)
- [Template Integration Problems](#template-integration-problems)
- [Performance Issues](#performance-issues)
- [Generated Code Issues](#generated-code-issues)
- [Tailwind Utilities Issues](#tailwind-utilities-issues)
- [Getting Help](#getting-help)

## Installation Issues

### Module Not Found

**Problem:**
```
go: cannot find module github.com/ahmed-com/typesafe-css
```

**Solution:**
```bash
# Ensure you're using the correct module path
go get github.com/ahmed-com/typesafe-css

# Check your go.mod file
cat go.mod

# Clean module cache if needed
go clean -modcache
go mod tidy
```

### Import Path Issues

**Problem:**
```go
import "github.com/ahmed-com/typesafe-css" // ❌ Wrong
```

**Solution:**
```go
// ✅ Correct - import specific packages
import "github.com/ahmed-com/typesafe-css/css"
import "github.com/ahmed-com/typesafe-css/tailwind"
import "github.com/ahmed-com/typesafe-css/cssgen"
```

## Compilation Errors

### Type Mismatch Errors

**Problem:**
```go
css.Set(css.Width, "100px") // ❌ Type error
```

**Error:**
```
cannot use "100px" (type string) as type css.Value in argument
```

**Solution:**
```go
// ✅ Use typed constructors
css.Set(css.Width, css.Px(100))
css.Set(css.Height, css.Rem(1.5))
css.Set(css.BackgroundColor, css.Hex("#3b82f6"))
```

### Property Name Errors

**Problem:**
```go
css.Set("width", css.Px(100)) // ❌ Type error
```

**Error:**
```
cannot use "width" (type string) as type css.Property in argument
```

**Solution:**
```go
// ✅ Use typed property constants
css.Set(css.Width, css.Px(100))
css.Set(css.Height, css.Px(200))
css.Set(css.ColorP, css.Hex("#000000")) // Note: ColorP for color property
```

### Undefined Function Errors

**Problem:**
```go
tailwind.BackgroundBlue500() // ❌ Undefined function
```

**Error:**
```
undefined: tailwind.BackgroundBlue500
```

**Solution:**
```go
// ✅ Use correct function names
tailwind.BgBlue500()        // Not BackgroundBlue500
tailwind.TextWhite()        // Not TextColorWhite
tailwind.P4()               // Not PaddingFour
tailwind.M8()               // Not MarginEight
```

### Missing Import Errors

**Problem:**
```go
rule := css.RuleSet(".btn", ...) // ❌ Undefined: css
```

**Solution:**
```go
import "github.com/ahmed-com/typesafe-css/css"

rule := css.RuleSet(".btn", 
    css.Set(css.Display, css.DisplayInlineBlock),
)
```

## Runtime Issues

### Empty CSS Output

**Problem:**
```go
var stylesheet css.Stylesheet
fmt.Println(stylesheet.String()) // Prints empty string
```

**Solution:**
```go
var stylesheet css.Stylesheet

// ❌ You must add rules to the stylesheet
// stylesheet.String() returns empty string

// ✅ Add rules first
rule := css.RuleSet(".btn", css.Set(css.ColorP, css.Hex("#000")))
stylesheet.Add(rule)
fmt.Println(stylesheet.String()) // Now prints CSS
```

### Nil Pointer Errors

**Problem:**
```go
var stylesheet *css.Stylesheet
stylesheet.Add(rule) // ❌ Panic: nil pointer dereference
```

**Solution:**
```go
// ✅ Initialize the stylesheet properly
var stylesheet css.Stylesheet  // Value, not pointer
stylesheet.Add(rule)

// Or with pointer
stylesheet := &css.Stylesheet{}
stylesheet.Add(rule)
```

### Invalid CSS Values

**Problem:**
```go
css.Set(css.Width, css.Raw("invalid-value"))
```

**Solution:**
```go
// ✅ Use valid CSS values
css.Set(css.Width, css.Px(100))     // Valid pixel value
css.Set(css.Width, css.Percent(50)) // Valid percentage
css.Set(css.Width, css.Raw("calc(100% - 20px)")) // Valid calc expression
```

## Template Integration Problems

### Template Execution Errors

**Problem:**
```go
tmpl.Execute(os.Stdout, stylesheet.String()) // ❌ Raw string
```

**Error:**
```
template: unsafe content type
```

**Solution:**
```go
// ✅ Use template.CSS for safety
data := struct {
    CSS template.CSS
}{
    CSS: template.CSS(stylesheet.String()),
}
tmpl.Execute(os.Stdout, data)
```

### Template Parsing Errors

**Problem:**
```go
tmpl := template.Must(template.New("page").Parse(`
<style>{{.CSS}}</style>
`))
```

**Error:**
```
template: "page":1: unexpected "<" in command
```

**Solution:**
```go
// ✅ Use proper template syntax
tmpl := template.Must(template.New("page").Parse(`
<style>{{.CSS}}</style>
`))
// Make sure CSS is of type template.CSS in your data
```

### Missing CSS in Output

**Problem:**
CSS doesn't appear in the rendered HTML.

**Solution:**
```go
// ✅ Ensure CSS is properly embedded
type PageData struct {
    Title string
    CSS   template.CSS // Important: use template.CSS type
}

data := PageData{
    Title: "My Page",
    CSS:   template.CSS(stylesheet.String()), // Convert to template.CSS
}
```

## Performance Issues

### Slow CSS Generation

**Problem:**
CSS generation is taking too long.

**Solution:**
```go
// ✅ Cache generated CSS
var cachedCSS string
var cacheOnce sync.Once

func GetCSS() string {
    cacheOnce.Do(func() {
        stylesheet := generateStylesheet()
        cachedCSS = stylesheet.String()
    })
    return cachedCSS
}

// ✅ Pre-generate CSS at startup
func init() {
    globalCSS = generateGlobalCSS().String()
}
```

### Memory Usage Issues

**Problem:**
High memory usage with large stylesheets.

**Solution:**
```go
// ✅ Use deduplication
cache := make(map[string]css.Rule)

func addRuleOnce(stylesheet *css.Stylesheet, rule css.Rule) {
    key := rule.String()
    if _, exists := cache[key]; !exists {
        cache[key] = rule
        stylesheet.Add(rule)
    }
}

// ✅ Generate CSS on demand
func generatePageCSS(pageType string) string {
    // Only generate CSS needed for this page type
    switch pageType {
    case "home":
        return generateHomeCSS().String()
    case "blog":
        return generateBlogCSS().String()
    default:
        return generateDefaultCSS().String()
    }
}
```

## Generated Code Issues

### Missing Generated Properties

**Problem:**
```go
cssgen.SetWidth(cssgen.WidthVal100px) // ❌ Undefined
```

**Error:**
```
undefined: cssgen.WidthVal100px
```

**Solution:**
```go
// ✅ Width doesn't have predefined enum values, use core API
css.Set(css.Width, css.Px(100))

// ✅ Use generated properties only for enum-based properties
cssgen.SetDisplay(cssgen.DisplayValFlex)        // ✅ Has enum values
cssgen.SetPosition(cssgen.PositionValRelative)  // ✅ Has enum values
```

### Code Generation Failures

**Problem:**
```bash
go generate
# Error: failed to generate CSS properties
```

**Solution:**
```bash
# ✅ Ensure spec files are present
ls spec/
# Should contain spec.json and other specification files

# ✅ Run code generation manually
go run ./cmd/cssgen -in ./spec -out ./cssgen -pkg cssgen

# ✅ Check for any syntax errors in generated code
go build ./cssgen
```

## Tailwind Utilities Issues

### Utility Function Not Found

**Problem:**
```go
tailwind.Margin4() // ❌ Undefined function
```

**Solution:**
```go
// ✅ Use correct function names
tailwind.M4()           // Margin all sides
tailwind.Mx4()          // Margin horizontal
tailwind.My4()          // Margin vertical
tailwind.Mt4()          // Margin top
tailwind.Mr4()          // Margin right
tailwind.Mb4()          // Margin bottom
tailwind.Ml4()          // Margin left
```

### Theme Customization Issues

**Problem:**
```go
theme := tailwind.CustomTheme(func(t *tailwind.Theme) {
    t.Colors["custom"] = css.Hex("#custom") // ❌ Wrong approach
})
```

**Solution:**
```go
// ✅ Use proper theme customization methods
theme := tailwind.CustomTheme(func(t *tailwind.Theme) {
    t.AddColor("custom", css.Hex("#custom"))
    t.AddSpacing("custom", css.Rem(1.5))
    t.AddBorderRadius("custom", css.Px(12))
})
```

### Utility Deduplication Not Working

**Problem:**
Multiple identical utility classes in output.

**Solution:**
```go
// ✅ Use the default manager for automatic deduplication
var stylesheet css.Stylesheet
stylesheet.Add(
    tailwind.P4(),  // First occurrence
    tailwind.P4(),  // Automatically deduplicated
    tailwind.P4(),  // Automatically deduplicated
)

// ✅ Check that you're using the same manager instance
manager := tailwind.GetDefaultManager()
rule1 := manager.P("4")
rule2 := manager.P("4") // Same rule instance
```

## Getting Help

### Debug Information

When asking for help, provide:

```go
// Include this information:
fmt.Printf("Go version: %s\n", runtime.Version())
fmt.Printf("GOOS: %s\n", runtime.GOOS)
fmt.Printf("GOARCH: %s\n", runtime.GOARCH)

// Your problematic code:
rule := css.RuleSet(".example",
    css.Set(css.Display, css.DisplayFlex),
    // ... rest of your code
)
fmt.Printf("Generated CSS: %s\n", rule.String())
```

### Minimal Reproduction

Create a minimal example:

```go
package main

import (
    "fmt"
    "github.com/ahmed-com/typesafe-css/css"
)

func main() {
    // Minimal code that reproduces the issue
    rule := css.RuleSet(".test",
        css.Set(css.ColorP, css.Hex("#000000")),
    )
    
    fmt.Println(rule.String())
    // Expected: .test{color:#000000}
    // Actual: [describe what you're seeing]
}
```

### Common Debugging Commands

```bash
# Check module version
go list -m github.com/ahmed-com/typesafe-css

# Verify imports
go mod why github.com/ahmed-com/typesafe-css

# Clean and rebuild
go clean -cache
go mod tidy
go build

# Run tests
go test ./...

# Check for race conditions
go run -race main.go

# Profile memory usage
go run -memprofile=mem.prof main.go
go tool pprof mem.prof
```

### Error Message Patterns

| Error Type | Pattern | Solution |
|------------|---------|----------|
| Import error | `package not found` | Check import path |
| Type error | `cannot use X as type Y` | Use typed constructors |
| Undefined error | `undefined: X` | Check function/type names |
| Template error | `unsafe content type` | Use `template.CSS` |
| Runtime panic | `nil pointer dereference` | Initialize variables |

### Resources

- **Examples**: Check the [examples/](../examples/) directory for working code
- **Documentation**: Review [COOKBOOK.md](../COOKBOOK.md) for patterns
- **API Reference**: Use `go doc` command for function signatures
- **Source Code**: Read the source for implementation details
- **Issues**: Open an issue on GitHub with reproduction steps

### Checklist Before Asking for Help

- [ ] Checked this troubleshooting guide
- [ ] Reviewed the relevant documentation
- [ ] Looked at similar examples in the examples directory
- [ ] Created a minimal reproduction case
- [ ] Included version information and error messages
- [ ] Searched existing GitHub issues

Remember: The type system is designed to help catch errors at compile time. If you're getting runtime errors, it's often because the type safety is being bypassed somewhere. Use the typed constructors and functions provided by TypeSafe CSS for the best experience.