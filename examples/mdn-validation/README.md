# MDN Validation Example

This example demonstrates CSS properties and values that have been validated against the authoritative MDN (Mozilla Developer Network) CSS documentation.

## What This Example Shows

- **MDN-validated CSS properties** ensuring spec compliance
- **Extended property values** beyond basic implementations
- **Modern CSS features** including Grid, Flexbox, and advanced cursors
- **Comprehensive value coverage** with validated enums
- **Real-world layouts** using validated properties
- **Template integration** with validated CSS

## Key Concepts Demonstrated

### 1. Enhanced Display Values

The MDN validation ensures access to all display values:

```go
// Standard display values
cssgen.SetDisplay(cssgen.DisplayValBlock)
cssgen.SetDisplay(cssgen.DisplayValFlex)
cssgen.SetDisplay(cssgen.DisplayValGrid)

// Advanced display values (MDN validated)
cssgen.SetDisplay(cssgen.DisplayValFlowRoot)    // For containing floats
cssgen.SetDisplay(cssgen.DisplayValContents)    // Removes element from layout
cssgen.SetDisplay(cssgen.DisplayValRunIn)       // Run-in display
```

### 2. Extended Cursor Values

MDN validation provides comprehensive cursor support:

```go
// Standard cursors
cssgen.SetCursor(cssgen.CursorValPointer)
cssgen.SetCursor(cssgen.CursorValDefault)

// Extended cursors (MDN validated)
cssgen.SetCursor(cssgen.CursorValColResize)     // Column resize cursor
cssgen.SetCursor(cssgen.CursorValZoomIn)        // Zoom in cursor
cssgen.SetCursor(cssgen.CursorValZoomOut)       // Zoom out cursor
cssgen.SetCursor(cssgen.CursorValGrab)          // Grab cursor
cssgen.SetCursor(cssgen.CursorValGrabbing)      // Grabbing cursor
```

### 3. Advanced Alignment Values

```go
// Standard alignment
cssgen.SetAlignItems(cssgen.AlignItemsValCenter)
cssgen.SetAlignItems(cssgen.AlignItemsValStretch)

// Advanced alignment (MDN validated)
cssgen.SetAlignItems(cssgen.AlignItemsValStart)      // Logical start
cssgen.SetAlignItems(cssgen.AlignItemsValEnd)        // Logical end
cssgen.SetAlignItems(cssgen.AlignItemsValSelfStart)  // Self-start
cssgen.SetAlignItems(cssgen.AlignItemsValSelfEnd)    // Self-end
cssgen.SetAlignItems(cssgen.AlignItemsValNormal)     // Normal alignment
```

### 4. Enhanced Text Alignment

```go
// Standard text alignment
cssgen.SetTextAlign(cssgen.TextAlignValLeft)
cssgen.SetTextAlign(cssgen.TextAlignValCenter)
cssgen.SetTextAlign(cssgen.TextAlignValRight)

// Advanced text alignment (MDN validated)
cssgen.SetTextAlign(cssgen.TextAlignValMatchParent)  // Match parent alignment
cssgen.SetTextAlign(cssgen.TextAlignValStart)        // Logical start
cssgen.SetTextAlign(cssgen.TextAlignValEnd)          // Logical end
```

### 5. Modern Grid Layout

```go
modernCard := css.RuleSet(".modern-card",
    // Grid with advanced properties
    cssgen.SetDisplay(cssgen.DisplayValGrid),
    css.Set("grid-template-areas", css.Raw("\"header header\" \"content sidebar\" \"footer footer\"")),
    css.Set("grid-template-rows", css.Raw("auto 1fr auto")),
    css.Set("grid-template-columns", css.Raw("2fr 1fr")),
    css.Set("gap", css.Rem(1.5)),
    
    // Enhanced visual styling
    css.Set("background", css.Raw("linear-gradient(135deg, #667eea 0%, #764ba2 100%)")),
    css.Set("border-radius", css.Rem(1)),
    css.Set("padding", css.Rem(2)),
    css.Set("box-shadow", css.Raw("0 20px 25px -5px rgba(0, 0, 0, 0.1)")),
)
```

### 6. Interactive Elements

```go
interactiveElements := css.RuleSet(".interactive",
    cssgen.SetCursor(cssgen.CursorValPointer),
    css.Set("transition", css.Raw("all 0.2s ease")),
    css.Set("user-select", css.Raw("none")), // Prevent text selection
)

// Specialized hover states
resizableHover := css.RuleSet(".resizable:hover",
    cssgen.SetCursor(cssgen.CursorValColResize),
)

zoomableHover := css.RuleSet(".zoomable:hover",
    cssgen.SetCursor(cssgen.CursorValZoomIn),
)
```

## MDN Validation Benefits

### 1. Spec Compliance
All properties and values are validated against official MDN documentation:

```go
// These values are guaranteed to be valid per MDN specs
cssgen.SetDisplay(cssgen.DisplayValFlowRoot)    // ✅ Valid MDN value
cssgen.SetCursor(cssgen.CursorValColResize)     // ✅ Valid MDN value
cssgen.SetAlignItems(cssgen.AlignItemsValNormal) // ✅ Valid MDN value
```

### 2. Comprehensive Coverage
MDN validation ensures no valid CSS values are missing:

```go
// Standard values
cssgen.CursorValPointer
cssgen.CursorValDefault

// Extended values (often missed in other implementations)
cssgen.CursorValNEResize     // Northeast resize
cssgen.CursorValNWResize     // Northwest resize
cssgen.CursorValSEResize     // Southeast resize
cssgen.CursorValSWResize     // Southwest resize
cssgen.CursorValEWResize     // East-west resize
cssgen.CursorValNSResize     // North-south resize
```

### 3. Future-Proof
Regular validation against MDN ensures new CSS features are included:

```go
// Modern features automatically included when available
cssgen.SetDisplay(cssgen.DisplayValFlowRoot)    // CSS3 feature
cssgen.SetAlignItems(cssgen.AlignItemsValNormal) // Flexbox enhancement
```

## Real-World Layout Examples

### 1. Modern Card Layout

```go
modernCard := css.RuleSet(".modern-card",
    // Use validated grid display
    cssgen.SetDisplay(cssgen.DisplayValGrid),
    
    // Advanced grid properties
    css.Set("grid-template-columns", css.Raw("repeat(auto-fit, minmax(250px, 1fr))")),
    css.Set("gap", css.Rem(1.5)),
    css.Set("padding", css.Rem(2)),
    
    // Enhanced styling
    css.Set("background", css.Raw("linear-gradient(135deg, #667eea 0%, #764ba2 100%)")),
    css.Set("border-radius", css.Rem(1)),
    css.Set("box-shadow", css.Raw("0 20px 25px -5px rgba(0, 0, 0, 0.1)")),
)
```

### 2. Grid Container with Auto-Fit

```go
gridContainer := css.RuleSet(".grid-container",
    cssgen.SetDisplay(cssgen.DisplayValGrid),
    css.Set("grid-template-columns", css.Raw("repeat(auto-fit, minmax(250px, 1fr))")),
    css.Set("gap", css.Rem(1.5)),
    css.Set("padding", css.Rem(2)),
)
```

### 3. Interactive Elements with Enhanced Cursors

```go
// Base interactive styles
interactiveElements := css.RuleSet(".interactive",
    cssgen.SetCursor(cssgen.CursorValPointer),
    css.Set("transition", css.Raw("all 0.2s ease")),
)

// Specialized cursor interactions
hoverStates := css.RuleSet(".resizable:hover",
    cssgen.SetCursor(cssgen.CursorValColResize),
)

zoomableHover := css.RuleSet(".zoomable:hover",
    cssgen.SetCursor(cssgen.CursorValZoomIn),
)
```

### 4. Layout Wrapper with Contents Display

```go
// Special layout using contents display (removes box from layout)
wrapperRemoval := css.RuleSet(".layout-wrapper",
    cssgen.SetDisplay(cssgen.DisplayValContents),
)
```

### 5. Flow Root for Float Containment

```go
// Flow root for containing floats without clearfix hacks
flowRoot := css.RuleSet(".clearfix",
    cssgen.SetDisplay(cssgen.DisplayValFlowRoot),
)
```

### 6. Enhanced Typography

```go
typography := css.RuleSet(".text-content",
    cssgen.SetTextAlign(cssgen.TextAlignValMatchParent),
    cssgen.SetFontWeight(cssgen.FontWeightValNormal),
    css.Set("line-height", css.Raw("1.6")),
    css.Set("font-variant-numeric", css.Raw("tabular-nums")),
)
```

## HTML Template Integration

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>MDN-Validated CSS Demo</title>
    <style>
        {{.CSS}}
        
        /* Additional base styles */
        body {
            margin: 0;
            padding: 2rem;
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
            background: linear-gradient(135deg, #1e3c72 0%, #2a5298 100%);
            min-height: 100vh;
        }
    </style>
</head>
<body>
    <div class="grid-container">
        <div class="modern-card interactive">
            <h2>MDN-Validated CSS</h2>
            <p>This card uses CSS properties validated against MDN documentation.</p>
            <div class="layout-wrapper">
                <button class="resizable">Resizable Element</button>
                <button class="zoomable">Zoomable Element</button>
            </div>
        </div>
        
        <div class="clearfix text-content">
            <h3>Typography Example</h3>
            <p>This text demonstrates enhanced typography with match-parent alignment.</p>
        </div>
    </div>
</body>
</html>
```

## Validation Process

The MDN validation process ensures:

### 1. Property Coverage
All CSS properties from MDN documentation are included:

```go
// Properties like 'align-items' include all MDN-documented values
cssgen.AlignItemsValNormal     // From CSS Box Alignment spec
cssgen.AlignItemsValStart      // From CSS Logical Properties spec
cssgen.AlignItemsValEnd        // From CSS Logical Properties spec
```

### 2. Value Accuracy
Property values match MDN specifications exactly:

```go
// Cursor values match MDN cursor documentation precisely
cssgen.CursorValEWResize      // east-west-resize
cssgen.CursorValNEResize      // ne-resize
cssgen.CursorValNWSEResize    // nwse-resize
```

### 3. Regular Updates
Validation runs against current MDN data to catch new features:

```bash
# Validation command (when MDN data is available)
go run ./cmd/mdnvalidate
```

## Running the Example

```bash
cd examples/mdn-validation
go run main.go
```

## Expected Output

The example generates:

1. **Modern grid layout** with auto-fitting columns
2. **Interactive elements** with enhanced cursor states
3. **Advanced typography** with logical property values
4. **Complete HTML page** demonstrating all features
5. **Validation confirmation** showing MDN compliance

## Use Cases

This pattern is ideal for:

- **Production applications** requiring spec compliance
- **Design systems** needing comprehensive CSS coverage
- **Modern web applications** using latest CSS features
- **Accessibility-focused projects** leveraging logical properties
- **Cross-browser compatibility** with validated properties

## MDN Validation Features

The validation system provides:

### Extended Cursor Support
```
✅ cursor: col-resize, zoom-in, zoom-out, and many more
✅ All resize directional cursors (n-resize, ne-resize, etc.)
✅ Interactive cursors (grab, grabbing, etc.)
```

### Enhanced Alignment
```
✅ align-items: start, end, self-start, self-end, normal
✅ justify-content: start, end, normal, stretch
✅ Logical property values for international layouts
```

### Modern Display Values
```
✅ display: flow-root (for containing floats)
✅ display: contents (removes element from layout tree)
✅ All CSS3+ display values
```

### Comprehensive Typography
```
✅ text-align: match-parent, start, end
✅ All font-weight numeric and keyword values
✅ Modern typography features
```

## Best Practices

### 1. Leverage MDN-Validated Values
Use the comprehensive value coverage:

```go
// Take advantage of all available cursor states
cssgen.SetCursor(cssgen.CursorValColResize)  // For resizable columns
cssgen.SetCursor(cssgen.CursorValZoomIn)     // For zoomable content
```

### 2. Use Modern Layout Features
Adopt modern CSS with confidence:

```go
// Use flow-root instead of clearfix hacks
cssgen.SetDisplay(cssgen.DisplayValFlowRoot)

// Use contents for layout wrapper removal
cssgen.SetDisplay(cssgen.DisplayValContents)
```

### 3. Embrace Logical Properties
Use logical properties for international layouts:

```go
cssgen.SetTextAlign(cssgen.TextAlignValStart)  // Logical start
cssgen.SetAlignItems(cssgen.AlignItemsValEnd)  // Logical end
```

## Advantages

1. **Spec Compliance**: Guaranteed valid CSS per MDN documentation
2. **Comprehensive Coverage**: No missing CSS values or properties
3. **Modern Features**: Access to latest CSS developments
4. **Cross-Browser Safety**: Only documented, supported features
5. **Future-Proof**: Regular updates with new CSS features
6. **Learning Aid**: Discover CSS capabilities through type system

## Next Steps

- Explore utility-first approach in [examples/tailwind/](../tailwind/)
- See type-safe configuration in [examples/enhanced-tailwind/](../enhanced-tailwind/)
- Check out generated API basics in [examples/generated/](../generated/)
- Read about [MDN Validation](../../cmd/mdnvalidate/README.md) tooling