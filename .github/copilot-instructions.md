# TypeSafe CSS - AI Agent Instructions

## Project Overview
TypeSafe CSS is a Go library that provides type-safe CSS authoring APIs to eliminate string-based CSS generation footguns. It features a dual architecture: **core API** (hand-curated, stable) + **generated code** (comprehensive property coverage via codegen).

## Core Architecture

### Two-Layer Design
- **`/css/`** - Core authoring API (stable, no dependencies, manually maintained)
  - `value.go` - Value types (`Length`, `Color`, `Keyword`, `Raw`)
  - `property.go` - Properties, declarations, rules, stylesheets
  - `shorthand.go` - Helpers for padding, margin, flexbox patterns
  - `serialize.go` - String output (compact + pretty-printed)
- **`/cssgen/`** - Generated types/constants/setters (DO NOT EDIT manually)
  - `properties_gen.go`, `keywords_gen.go`, `setters_gen.go`

### Key Patterns
```go
// Core API pattern - always available
css.RuleSet(".btn", 
    css.Set(css.Display, css.DisplayFlex),
    css.Set(css.Padding, css.PadXY(css.Px(8), css.Px(12))),
)

// Generated API pattern - type-safe enums
cssgen.SetDisplay(cssgen.DisplayValFlex)
cssgen.SetPosition(cssgen.PositionValRelative)
```

## Code Generation Workflow

### Critical Commands
```bash
# Regenerate CSS properties/types from spec
go generate

# The above runs: go run ./cmd/cssgen -in ./cmd/cssgen/spec/spec.json -out ./cssgen -pkg cssgen
```

### Spec-Driven Development
- Spec source: `cmd/cssgen/spec/spec.json` (versioned, pinned CSS property definitions)
- Generator: `cmd/cssgen/main.go` (reads JSON → Go constants/types/setters)
- Generated code includes headers with spec version + timestamp for tracking

## Development Conventions

### Value Construction
- **Always use constructors**: `css.Px(16)` not `css.Length("16px")`
- **Raw for complex values**: `css.Raw("calc(100% - 8px)")` for functions
- **Shorthand helpers**: `css.PadXY()`, `css.FlexCenter()` for common patterns

### Serialization Strategy
- `String()` methods produce compact CSS (no extra whitespace)
- `css.PrettyCSS()` for human-readable formatting
- Template integration via `template.CSS()` for safe HTML embedding

### Testing Patterns
- Golden tests for `String()` output in `css/css_test.go`
- Test both core API and generated API interactions
- Use `go test ./...` (no test files in examples - they're demos)

## Extension Points

### Adding Properties
1. **Core properties**: Add to `css/property.go` constants manually
2. **Comprehensive coverage**: Update `cmd/cssgen/spec/spec.json` + run `go generate`
3. **Custom shorthands**: Add to `css/shorthand.go` with canonicalization

### Spec Updates
- Modify `cmd/cssgen/spec/spec.json` following existing structure
- Include `keywords` array for finite enums (generates typed setters)
- Empty `keywords` means freeform values (generic `css.Set()` only)

## Integration Patterns

### Template Usage
```go
// Safe CSS embedding in templates
tmpl := template.Must(template.New("page").Parse(`<style>{{.CSS}}</style>`))
data := struct{ CSS template.CSS }{CSS: template.CSS(stylesheet.String())}
```

### Responsive/At-Rules
```go
css.AtRule{
    Name: "media", 
    Params: "(max-width: 640px)",
    Body: []css.Item{css.RuleSet(".btn", ...)},
}
```

## Critical Constraints
- **Never edit `/cssgen/` manually** - always regenerate via `go generate`
- **Preserve string canonicalization** - shorthands must collapse properly (`PadXY` → `"12px 8px"`)
- **Maintain zero runtime deps** - core `/css/` package is self-contained
- **Type safety first** - prefer generated setters over generic `css.Set()` when available

## Examples Reference
- `examples/basic/` - Core API usage patterns  
- `examples/generated/` - Mixed core + generated API usage
- Both demonstrate template integration and common CSS patterns

When extending this codebase, always consider: Does this belong in core (stable, minimal) or generated (comprehensive, spec-driven) layer?
