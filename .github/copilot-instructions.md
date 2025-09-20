# TypeSafe CSS - AI Agent Instructions

## Project Overview
TypeSafe CSS is a Go library that provides type-safe CSS authoring APIs to eliminate string-based CSS generation footguns. It features a dual architecture: **core API** (hand-curated, stable) + **generated code** (comprehensive property coverage via codegen).

**Current Status**: The project has completed multiple phases including core API implementation, code generation pipeline, and comprehensive Tailwind CSS utilities. The library is feature-complete with extensive CSS property coverage and utility-first authoring support.

## Core Architecture

### Two-Layer Design
- **`/css/`** - Core authoring API (stable, no dependencies, manually maintained)
  - `value.go` - Value types (`Length`, `Color`, `Keyword`, `Raw`)
  - `property.go` - Properties, declarations, rules, stylesheets
  - `shorthand.go` - Helpers for padding, margin, flexbox patterns
  - `serialize.go` - String output (compact + pretty-printed)
- **`/cssgen/`** - Generated types/constants/setters (DO NOT EDIT manually)
  - `properties_gen.go`, `keywords_gen.go`, `setters_gen.go`
- **`/tailwind/`** - Tailwind CSS-style utility package (utility-first authoring)
  - Complete utility class generation with theme support
- **`/spec/`** - MDN CSS data and specifications (source for code generation)
- **`cmd/cssgen/`** - Code generator tool
- **`cmd/mdnvalidate/`** - MDN validation tool

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

// Tailwind utilities pattern - utility-first CSS
tailwind.Bg("blue-500")     // .bg-blue-500 { background-color: #3b82f6; }
tailwind.P("4")             // .p-4 { padding: 1rem; }
tailwind.DisplayFlex()      // .flex { display: flex; }
```

## Code Generation Workflow

### Critical Commands
```bash
# Regenerate CSS properties/types from spec
go generate

# The above runs: go run ./cmd/cssgen -in ./spec -out ./cssgen -pkg cssgen
```

### Spec-Driven Development
- Spec source: `cmd/cssgen/spec/spec.json` (versioned, pinned CSS property definitions)
- Generator: `cmd/cssgen/main.go` (reads JSON → Go constants/types/setters)
- Generated code includes headers with spec version + timestamp for tracking

### Tailwind CSS Utilities
- **`/tailwind/`** - Comprehensive Tailwind CSS-style utility package
  - `tailwind.go` - Core utility functions and class generation
  - `colors.go`, `spacing.go`, `typography.go` - Design token definitions
  - `theme.go` - Configurable theming system with custom overrides
  - `manager.go` - Utility deduplication and class management
- Utility-first CSS authoring: `tailwind.Bg("blue-500")`, `tailwind.P("4")`
- Custom themes and arbitrary values: `tailwind.Bg("#bada55")`
- Automatic class deduplication and type-safe API

### MDN Validation Tool
- **`cmd/mdnvalidate/`** - Validates CSS specifications against authoritative MDN data
- Ensures generated types/constants match official CSS specifications
- Helps catch missing properties and incorrect keyword values
- Can update spec files with latest MDN data
- Integrated into test suite for continuous validation

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
- MDN validation tests ensure spec accuracy against authoritative sources
- Use `go test ./...` (no test files in examples - they're demos)
- Comprehensive test coverage includes Tailwind utilities and theme validation

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
- `examples/tailwind/` - Tailwind CSS utilities demonstration
- `examples/enhanced-tailwind/` - Advanced Tailwind features and theming
- `examples/mdn-validation/` - MDN validation tool usage
- All examples demonstrate template integration and common CSS patterns

When extending this codebase, always consider: Does this belong in core (stable, minimal) or generated (comprehensive, spec-driven) layer?
