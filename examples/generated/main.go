package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/ahmed-com/typesafe-css/css"
	"github.com/ahmed-com/typesafe-css/cssgen"
)

func main() {
	// Example using generated type-safe CSS properties and values
	rule := css.RuleSet(".card",
		// Type-safe display property
		cssgen.SetDisplay(cssgen.DisplayValFlex),
		
		// Type-safe flex properties  
		cssgen.SetFlexDirection(cssgen.FlexDirectionValColumn),
		cssgen.SetJustifyContent(cssgen.JustifyContentValSpaceBetween),
		cssgen.SetAlignItems(cssgen.AlignItemsValCenter),
		
		// Type-safe positioning
		cssgen.SetPosition(cssgen.PositionValRelative),
		
		// Manual properties (not generated yet)
		css.Set(css.Width, css.Px(300)),
		css.Set(css.Height, css.Px(200)),
		css.Set(css.Padding, css.Px(16)),
		css.Set(css.BackgroundColor, css.Hex("#f0f0f0")),
		css.Set(css.BorderRadius, css.Px(8)),
	)

	// Create a button with generated and manual styles
	buttonRule := css.RuleSet(".btn",
		cssgen.SetDisplay(cssgen.DisplayValInlineBlock),
		cssgen.SetCursor(cssgen.CursorValPointer),
		cssgen.SetTextAlign(cssgen.TextAlignValCenter),
		cssgen.SetFontWeight(cssgen.FontWeightValBold),
		
		css.Set(css.Padding, css.Raw("8px 16px")),
		css.Set(css.BackgroundColor, css.Hex("#007bff")),
		css.Set(css.ColorP, css.Hex("#ffffff")),
		css.Set(css.BorderRadius, css.Px(4)),
		css.Set(css.Border, css.Raw("none")),
	)

	// Hover state using manual selector
	buttonHoverRule := css.RuleSet(".btn:hover",
		css.Set(css.BackgroundColor, css.Hex("#0056b3")),
		cssgen.SetCursor(cssgen.CursorValPointer),
	)

	// Create stylesheet
	var stylesheet css.Stylesheet
	stylesheet.Add(rule, buttonRule, buttonHoverRule)

	// Output CSS
	fmt.Println("Generated CSS:")
	fmt.Println(stylesheet.String())
	fmt.Println()

	// Demonstrate template usage
	tmpl := template.Must(template.New("demo").Parse(`
<!DOCTYPE html>
<html>
<head>
	<style>{{.CSS}}</style>
</head>
<body>
	<div class="card">
		<h2>Card Title</h2>
		<p>This card uses type-safe CSS properties generated from the spec.</p>
		<button class="btn">Click Me</button>
	</div>
</body>
</html>
`))

	data := struct {
		CSS template.CSS
	}{
		CSS: template.CSS(stylesheet.String()),
	}

	fmt.Println("HTML Template with embedded CSS:")
	tmpl.Execute(os.Stdout, data)
}