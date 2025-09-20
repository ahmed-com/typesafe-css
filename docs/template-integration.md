# Template Integration Guide

Learn how to integrate TypeSafe CSS with Go's template system and popular web frameworks.

## Table of Contents

- [Basic Template Integration](#basic-template-integration)
- [Advanced Template Patterns](#advanced-template-patterns)
- [Web Framework Integration](#web-framework-integration)
- [Static Site Generation](#static-site-generation)
- [Performance Optimization](#performance-optimization)
- [Best Practices](#best-practices)

## Basic Template Integration

### HTML Templates

Use TypeSafe CSS with Go's built-in `html/template` package:

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
    
    // Create template
    tmpl := template.Must(template.New("page").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <style>{{.CSS}}</style>
</head>
<body>
    <div class="bg-blue-500 text-white p-4 rounded-lg">
        <h1>{{.Heading}}</h1>
        <p>{{.Content}}</p>
    </div>
</body>
</html>
    `))
    
    // Execute template with data
    data := struct {
        Title   string
        Heading string
        Content string
        CSS     template.CSS
    }{
        Title:   "TypeSafe CSS Demo",
        Heading: "Welcome",
        Content: "This page uses type-safe CSS!",
        CSS:     template.CSS(stylesheet.String()),
    }
    
    tmpl.Execute(os.Stdout, data)
}
```

### Text Templates

For non-HTML output, use `text/template`:

```go
import "text/template"

func generateCSS() string {
    var stylesheet css.Stylesheet
    stylesheet.Add(
        tailwind.BgBlue500(),
        tailwind.TextWhite(),
        tailwind.P4(),
    )
    return stylesheet.String()
}

func main() {
    tmpl := template.Must(template.New("css").Parse(`
/* Generated CSS */
{{.CSS}}

/* Custom additions */
.custom-class {
    /* Add custom styles here */
}
    `))
    
    data := struct {
        CSS string
    }{
        CSS: generateCSS(),
    }
    
    tmpl.Execute(os.Stdout, data)
}
```

## Advanced Template Patterns

### Component-Based Templates

Create reusable components with their own styles:

```go
type Component struct {
    Name   string
    Styles css.Stylesheet
}

func (c *Component) CSS() template.CSS {
    return template.CSS(c.Styles.String())
}

func CreateButton() *Component {
    var styles css.Stylesheet
    styles.Add(
        tailwind.BgBlue500(),
        tailwind.HoverBgBlue600(),
        tailwind.TextWhite(),
        tailwind.FontSemibold(),
        tailwind.Py2(),
        tailwind.Px4(),
        tailwind.RoundedMd(),
        tailwind.CursorPointer(),
    )
    
    return &Component{
        Name:   "button",
        Styles: styles,
    }
}

func CreateCard() *Component {
    var styles css.Stylesheet
    styles.Add(
        tailwind.BgWhite(),
        tailwind.ShadowLg(),
        tailwind.RoundedLg(),
        tailwind.P6(),
        tailwind.M4(),
    )
    
    return &Component{
        Name:   "card",
        Styles: styles,
    }
}

func main() {
    button := CreateButton()
    card := CreateCard()
    
    tmpl := template.Must(template.New("page").Parse(`
<!DOCTYPE html>
<html>
<head>
    <style>
        {{.Button.CSS}}
        {{.Card.CSS}}
    </style>
</head>
<body>
    <div class="card">
        <h2>Component Example</h2>
        <button class="button">Click Me</button>
    </div>
</body>
</html>
    `))
    
    data := struct {
        Button *Component
        Card   *Component
    }{
        Button: button,
        Card:   card,
    }
    
    tmpl.Execute(os.Stdout, data)
}
```

### Conditional Styles

Apply styles based on template conditions:

```go
type PageData struct {
    Title    string
    IsDark   bool
    IsLarge  bool
    Content  string
}

func (p *PageData) CSS() template.CSS {
    var stylesheet css.Stylesheet
    
    // Base styles
    stylesheet.Add(
        tailwind.P4(),
        tailwind.FontSans(),
    )
    
    // Conditional styles
    if p.IsDark {
        stylesheet.Add(
            tailwind.BgGray900(),
            tailwind.TextWhite(),
        )
    } else {
        stylesheet.Add(
            tailwind.BgWhite(),
            tailwind.TextGray900(),
        )
    }
    
    if p.IsLarge {
        stylesheet.Add(
            tailwind.TextXl(),
            tailwind.P8(),
        )
    } else {
        stylesheet.Add(
            tailwind.TextBase(),
            tailwind.P4(),
        )
    }
    
    return template.CSS(stylesheet.String())
}

func main() {
    data := &PageData{
        Title:   "Conditional Styles",
        IsDark:  true,
        IsLarge: false,
        Content: "This page adapts its styles based on conditions.",
    }
    
    tmpl := template.Must(template.New("page").Parse(`
<!DOCTYPE html>
<html>
<head>
    <style>{{.CSS}}</style>
</head>
<body>
    <div class="container">
        <h1>{{.Title}}</h1>
        <p>{{.Content}}</p>
    </div>
</body>
</html>
    `))
    
    tmpl.Execute(os.Stdout, data)
}
```

### Responsive Templates

Create responsive designs with media queries:

```go
func CreateResponsiveStyles() css.Stylesheet {
    var stylesheet css.Stylesheet
    
    // Base styles (mobile first)
    stylesheet.Add(css.RuleSet(".container",
        css.Set(css.Padding, css.Px(16)),
        css.Set(css.MaxWidth, css.Percent(100)),
    ))
    
    stylesheet.Add(css.RuleSet(".grid",
        css.Set(css.Display, css.DisplayGrid),
        css.Set(css.GridTemplateColumns, css.Raw("1fr")),
        css.Set(css.Gap, css.Px(16)),
    ))
    
    // Tablet styles
    stylesheet.Add(css.AtRule{
        Name:   "media",
        Params: "(min-width: 768px)",
        Body: []css.Item{
            css.RuleSet(".container",
                css.Set(css.Padding, css.Px(32)),
                css.Set(css.MaxWidth, css.Px(768)),
                css.Set(css.Margin, css.Raw("0 auto")),
            ),
            css.RuleSet(".grid",
                css.Set(css.GridTemplateColumns, css.Raw("repeat(2, 1fr)")),
                css.Set(css.Gap, css.Px(24)),
            ),
        },
    })
    
    // Desktop styles
    stylesheet.Add(css.AtRule{
        Name:   "media",
        Params: "(min-width: 1024px)",
        Body: []css.Item{
            css.RuleSet(".container",
                css.Set(css.MaxWidth, css.Px(1024)),
                css.Set(css.Padding, css.Px(48)),
            ),
            css.RuleSet(".grid",
                css.Set(css.GridTemplateColumns, css.Raw("repeat(3, 1fr)")),
                css.Set(css.Gap, css.Px(32)),
            ),
        },
    })
    
    return stylesheet
}
```

## Web Framework Integration

### Gin Framework

```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/ahmed-com/typesafe-css/css"
    "github.com/ahmed-com/typesafe-css/tailwind"
)

func setupRoutes() *gin.Engine {
    r := gin.Default()
    
    // Generate CSS once at startup
    stylesheet := generateGlobalCSS()
    cssContent := stylesheet.String()
    
    // Serve CSS endpoint
    r.GET("/styles.css", func(c *gin.Context) {
        c.Header("Content-Type", "text/css; charset=utf-8")
        c.Header("Cache-Control", "public, max-age=31536000") // Cache for 1 year
        c.String(200, cssContent)
    })
    
    // Load HTML templates
    r.LoadHTMLGlob("templates/*")
    
    // Home page
    r.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", gin.H{
            "title": "TypeSafe CSS App",
            "user":  "John Doe",
        })
    })
    
    // Dynamic page with custom styles
    r.GET("/profile/:id", func(c *gin.Context) {
        userID := c.Param("id")
        
        // Generate user-specific styles
        userStyles := generateUserStyles(userID)
        
        c.HTML(200, "profile.html", gin.H{
            "title":      "User Profile",
            "userID":     userID,
            "customCSS":  template.CSS(userStyles.String()),
        })
    })
    
    return r
}

func generateGlobalCSS() css.Stylesheet {
    var stylesheet css.Stylesheet
    
    // Add global utility classes
    stylesheet.Add(
        tailwind.GenerateUtilityStylesheet().Items...,
    )
    
    // Add custom component styles
    stylesheet.Add(
        createHeaderStyles().Items...,
        createFooterStyles().Items...,
        createNavigationStyles().Items...,
    )
    
    return stylesheet
}

func generateUserStyles(userID string) css.Stylesheet {
    var stylesheet css.Stylesheet
    
    // Generate user-specific theme
    if userID == "premium" {
        stylesheet.Add(
            tailwind.BgGold500(),
            tailwind.TextBlack(),
        )
    } else {
        stylesheet.Add(
            tailwind.BgGray100(),
            tailwind.TextGray800(),
        )
    }
    
    return stylesheet
}
```

### Echo Framework

```go
package main

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/ahmed-com/typesafe-css/css"
    "github.com/ahmed-com/typesafe-css/tailwind"
)

type TemplateRenderer struct {
    templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
    e := echo.New()
    
    // Setup template renderer
    renderer := &TemplateRenderer{
        templates: template.Must(template.ParseGlob("templates/*.html")),
    }
    e.Renderer = renderer
    
    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    
    // Generate and cache CSS
    globalCSS := generateAppCSS()
    
    // CSS endpoint
    e.GET("/app.css", func(c echo.Context) error {
        c.Response().Header().Set("Content-Type", "text/css")
        return c.String(200, globalCSS.String())
    })
    
    // Routes
    e.GET("/", func(c echo.Context) error {
        return c.Render(200, "index.html", map[string]interface{}{
            "title": "Echo + TypeSafe CSS",
        })
    })
    
    e.Logger.Fatal(e.Start(":8080"))
}
```

### Fiber Framework

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html"
    "github.com/ahmed-com/typesafe-css/css"
    "github.com/ahmed-com/typesafe-css/tailwind"
)

func main() {
    // Create a new template engine
    engine := html.New("./templates", ".html")
    
    // Create Fiber app with template engine
    app := fiber.New(fiber.Config{
        Views: engine,
    })
    
    // Generate CSS
    appCSS := generateFiberCSS()
    
    // Serve CSS
    app.Get("/styles.css", func(c *fiber.Ctx) error {
        c.Set("Content-Type", "text/css")
        return c.SendString(appCSS.String())
    })
    
    // Routes
    app.Get("/", func(c *fiber.Ctx) error {
        return c.Render("index", fiber.Map{
            "Title": "Fiber + TypeSafe CSS",
        })
    })
    
    app.Listen(":3000")
}

func generateFiberCSS() css.Stylesheet {
    var stylesheet css.Stylesheet
    
    // Fiber-specific optimizations
    stylesheet.Add(
        // Fast loading utilities
        tailwind.P4(),
        tailwind.M4(),
        tailwind.Flex(),
        tailwind.Grid(),
        
        // Performance-critical styles
        css.RuleSet("*",
            css.Set(css.BoxSizing, css.Raw("border-box")),
        ),
    )
    
    return stylesheet
}
```

## Static Site Generation

### Basic Static Generation

```go
package main

import (
    "os"
    "path/filepath"
    "text/template"
    "github.com/ahmed-com/typesafe-css/css"
    "github.com/ahmed-com/typesafe-css/tailwind"
)

type Page struct {
    Title    string
    Content  string
    Styles   css.Stylesheet
    Output   string
}

func generateSite() error {
    // Define pages
    pages := []Page{
        {
            Title:   "Home",
            Content: "Welcome to our site!",
            Styles:  createHomeStyles(),
            Output:  "public/index.html",
        },
        {
            Title:   "About",
            Content: "Learn more about us.",
            Styles:  createAboutStyles(),
            Output:  "public/about.html",
        },
    }
    
    // Load template
    tmpl := template.Must(template.ParseFiles("templates/page.html"))
    
    // Generate pages
    for _, page := range pages {
        // Ensure directory exists
        os.MkdirAll(filepath.Dir(page.Output), 0755)
        
        // Create output file
        file, err := os.Create(page.Output)
        if err != nil {
            return err
        }
        defer file.Close()
        
        // Execute template
        data := struct {
            Title string
            Content string
            CSS   template.CSS
        }{
            Title:   page.Title,
            Content: page.Content,
            CSS:     template.CSS(page.Styles.String()),
        }
        
        err = tmpl.Execute(file, data)
        if err != nil {
            return err
        }
    }
    
    // Generate global CSS file
    return generateGlobalCSS()
}

func generateGlobalCSS() error {
    utilities := tailwind.GenerateUtilityStylesheet()
    
    file, err := os.Create("public/styles.css")
    if err != nil {
        return err
    }
    defer file.Close()
    
    _, err = file.WriteString(utilities.String())
    return err
}
```

### Hugo Integration

```go
// layouts/_default/baseof.html template helper

{{ define "css" }}
{{ $cssContent := partial "generate-css.html" . }}
<style>{{ $cssContent | safeCSS }}</style>
{{ end }}
```

```go
// layouts/partials/generate-css.html

{{/* This would call a Go function to generate CSS */}}
{{ return (generateTypeCSS .) }}
```

## Performance Optimization

### CSS Caching

```go
type CSSCache struct {
    cache map[string]string
    mu    sync.RWMutex
}

func NewCSSCache() *CSSCache {
    return &CSSCache{
        cache: make(map[string]string),
    }
}

func (c *CSSCache) Get(key string) (string, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    css, exists := c.cache[key]
    return css, exists
}

func (c *CSSCache) Set(key, css string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.cache[key] = css
}

func (c *CSSCache) GetOrGenerate(key string, generator func() css.Stylesheet) string {
    if cached, exists := c.Get(key); exists {
        return cached
    }
    
    stylesheet := generator()
    generated := stylesheet.String()
    c.Set(key, generated)
    return generated
}

// Usage in web application
var cssCache = NewCSSCache()

func handlePage(w http.ResponseWriter, r *http.Request) {
    pageCSS := cssCache.GetOrGenerate("page-styles", func() css.Stylesheet {
        var stylesheet css.Stylesheet
        // Generate styles...
        return stylesheet
    })
    
    // Use pageCSS in template...
}
```

### Minification

```go
func MinifyCSS(stylesheet css.Stylesheet) string {
    // Use compact string representation
    return stylesheet.String() // Already minified
}

func MinifyAndCompress(stylesheet css.Stylesheet) ([]byte, error) {
    minified := stylesheet.String()
    
    var buf bytes.Buffer
    gzipWriter := gzip.NewWriter(&buf)
    
    _, err := gzipWriter.Write([]byte(minified))
    if err != nil {
        return nil, err
    }
    
    err = gzipWriter.Close()
    if err != nil {
        return nil, err
    }
    
    return buf.Bytes(), nil
}
```

### Build-Time Generation

```go
//go:generate go run generate-css.go

package main

import (
    "os"
    "github.com/ahmed-com/typesafe-css/css"
    "github.com/ahmed-com/typesafe-css/tailwind"
)

func main() {
    // Generate CSS at build time
    utilities := tailwind.GenerateUtilityStylesheet()
    components := generateComponentCSS()
    
    // Combine all CSS
    var combined css.Stylesheet
    combined.Add(utilities.Items...)
    combined.Add(components.Items...)
    
    // Write to file
    css := combined.String()
    err := os.WriteFile("assets/styles.css", []byte(css), 0644)
    if err != nil {
        panic(err)
    }
    
    // Generate hash for cache busting
    hash := generateHash(css)
    err = os.WriteFile("assets/styles.hash", []byte(hash), 0644)
    if err != nil {
        panic(err)
    }
}
```

## Best Practices

### 1. Separate Concerns

```go
// CSS generation
func GenerateCSS() css.Stylesheet {
    var stylesheet css.Stylesheet
    // Generate styles...
    return stylesheet
}

// Template data
type PageData struct {
    Title   string
    Content string
}

// Template execution
func RenderPage(data PageData) error {
    styles := GenerateCSS()
    
    templateData := struct {
        PageData
        CSS template.CSS
    }{
        PageData: data,
        CSS:      template.CSS(styles.String()),
    }
    
    return tmpl.Execute(os.Stdout, templateData)
}
```

### 2. Use CSS Variables for Theming

```go
func GenerateThemeCSS(theme string) css.Stylesheet {
    var stylesheet css.Stylesheet
    
    var colors map[string]css.Color
    switch theme {
    case "dark":
        colors = map[string]css.Color{
            "--bg-primary":   css.Hex("#1f2937"),
            "--text-primary": css.Hex("#f9fafb"),
        }
    default:
        colors = map[string]css.Color{
            "--bg-primary":   css.Hex("#ffffff"),
            "--text-primary": css.Hex("#111827"),
        }
    }
    
    var declarations []css.Declaration
    for prop, color := range colors {
        declarations = append(declarations, css.Set(prop, color))
    }
    
    stylesheet.Add(css.RuleSet(":root", declarations...))
    return stylesheet
}
```

### 3. Component-Based Organization

```go
type StyleComponent interface {
    CSS() css.Stylesheet
    Name() string
}

type Button struct {
    variant string
}

func (b Button) CSS() css.Stylesheet {
    var stylesheet css.Stylesheet
    // Generate button styles based on variant
    return stylesheet
}

func (b Button) Name() string {
    return "button"
}

func CombineComponents(components ...StyleComponent) css.Stylesheet {
    var combined css.Stylesheet
    for _, component := range components {
        combined.Add(component.CSS().Items...)
    }
    return combined
}
```

### 4. Error Handling

```go
func SafeTemplateExecution(tmpl *template.Template, data interface{}) error {
    var buf bytes.Buffer
    
    // Execute to buffer first to catch errors
    err := tmpl.Execute(&buf, data)
    if err != nil {
        return fmt.Errorf("template execution failed: %w", err)
    }
    
    // Write to actual output only if successful
    _, err = os.Stdout.Write(buf.Bytes())
    return err
}
```

This guide covers the essential patterns for integrating TypeSafe CSS with Go templates and web frameworks. The type-safe approach ensures your styles are maintainable and error-free across your entire application stack.