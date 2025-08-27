package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/ahmed-com/typesafe-css/css"
)

func main() {
	// Create some CSS rules using the type-safe API
	btn := css.RuleSet(".btn",
		css.Set(css.Display, css.DisplayInline),
		css.Set(css.ColorP, css.Hex("#0af")),
		css.Set(css.Padding, css.PadXY(css.Px(8), css.Px(12))),
		css.Set(css.FontSize, css.Rem(0.875)),
		css.Set(css.BackgroundColor, css.Hex("#f0f0f0")),
		css.Set(css.BorderRadius, css.Px(4)),
		css.Set(css.Border, css.BorderShorthand(css.Px(1), css.BorderSolid, css.Hex("#ccc"))),
	)

	// Create a responsive rule
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

	// Create a flex container
	container := css.RuleSet(".container",
		css.FlexCenter()...,
	)

	// Create a stylesheet and add rules
	var sheet css.Stylesheet
	sheet.Add(btn, container, mobile)

	// Output compact CSS
	fmt.Println("=== Compact CSS ===")
	fmt.Println(css.CSS(sheet.Items...))
	fmt.Println()

	// Output pretty-formatted CSS
	fmt.Println("=== Pretty CSS ===")
	fmt.Println(css.PrettyCSS(sheet.Items...))
	fmt.Println()

	// Demonstrate template integration
	fmt.Println("=== Template Integration ===")
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
	if err != nil {
		fmt.Printf("Template error: %v\n", err)
	}
}