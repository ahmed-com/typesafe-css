package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/ahmed-com/typesafe-css/css"
	"github.com/ahmed-com/typesafe-css/tailwind"
)

func main() {
	fmt.Println("=== Tailwind CSS Utilities Demo ===")
	fmt.Println()

	// Create a stylesheet with Tailwind-style utilities
	var stylesheet css.Stylesheet

	// Add various utility classes to demonstrate functionality
	stylesheet.Add(
		// Layout utilities
		tailwind.DisplayFlex(),
		tailwind.DisplayBlock(),
		tailwind.DisplayHidden(),
		tailwind.JustifyCenterClass(),
		tailwind.ItemsCenterClass(),
		tailwind.FlexColClass(),

		// Spacing utilities - demonstrate deduplication
		tailwind.P("4"),     // padding: 1rem
		tailwind.Px("6"),    // padding-left/right: 1.5rem
		tailwind.Py("2"),    // padding-top/bottom: 0.5rem
		tailwind.M("8"),     // margin: 2rem
		tailwind.Mx("auto"), // margin-left/right: auto
		tailwind.W("64"),    // width: 16rem
		tailwind.W("full"),  // width: 100%
		tailwind.H("32"),    // height: 8rem
		tailwind.H("screen"), // height: 100vh

		// Color utilities
		tailwind.Bg("blue-500"),    // background-color: #3b82f6
		tailwind.Bg("gray-100"),    // background-color: #f3f4f6
		tailwind.Text("white"),     // color: #ffffff
		tailwind.Text("gray-800"),  // color: #1f2937
		tailwind.Border("gray-300"), // border-color: #d1d5db

		// Typography utilities
		tailwind.TextLg(),          // font-size: 1.125rem
		tailwind.TextCenter(),      // text-align: center
		tailwind.FontBold(),        // font-weight: 700
		tailwind.FontSemibold(),    // font-weight: 600
		tailwind.TextSm(),          // font-size: 0.875rem

		// Position utilities
		tailwind.RelativeClass(),   // position: relative
		tailwind.AbsoluteClass(),   // position: absolute
	)

	// Print compact CSS
	fmt.Println("Generated Tailwind Utilities (Compact):")
	fmt.Println(stylesheet.String())
	fmt.Println()

	// Print pretty-formatted CSS
	fmt.Println("Generated Tailwind Utilities (Pretty):")
	fmt.Println(css.PrettyCSS(stylesheet.Items...))
	fmt.Println()

	// Demonstrate custom theme usage
	fmt.Println("=== Custom Theme Demo ===")
	customTheme := tailwind.CustomTheme(&tailwind.Theme{
		Colors: map[string]css.Color{
			"brand":   css.Hex("#6366f1"), // Custom brand color
			"danger":  css.Hex("#ef4444"), // Custom danger color
		},
		Spacing: map[string]css.Length{
			"xs": css.Rem(0.125), // 2px
			"xl": css.Rem(5),     // 80px
		},
	})

	customManager := tailwind.WithCustomTheme(customTheme)
	var customStylesheet css.Stylesheet

	customStylesheet.Add(
		tailwind.BackgroundColor(customManager, "brand"),
		tailwind.TextColor(customManager, "danger"),
		tailwind.Padding(customManager, "xs"),
		tailwind.Margin(customManager, "xl"),
	)

	fmt.Println("Custom Theme Utilities:")
	fmt.Println(css.PrettyCSS(customStylesheet.Items...))
	fmt.Println()

	// Demonstrate utility generation function
	fmt.Println("=== Utility Generation Demo ===")
	utilityStylesheet := tailwind.GenerateUtilityStylesheet()
	fmt.Printf("Generated %d common utility classes\n", len(utilityStylesheet.Items))
	fmt.Println()

	// Show first few utilities as examples
	fmt.Println("Sample Generated Utilities:")
	for i, item := range utilityStylesheet.Items[:10] {
		fmt.Printf("%d. %s\n", i+1, item.String())
	}
	fmt.Println("... and more")
	fmt.Println()

	// Demonstrate deduplication by showing that the same utility returns the same rule
	fmt.Println("=== Deduplication Demo ===")
	manager := tailwind.GetDefaultManager()
	
	// Call the same utility multiple times
	rule1 := tailwind.P("4")
	rule2 := tailwind.P("4") 
	rule3 := tailwind.P("4")
	
	// These should all be the same object due to deduplication
	fmt.Printf("Rule 1 address: %p\n", &rule1)
	fmt.Printf("Rule 2 address: %p\n", &rule2)
	fmt.Printf("Rule 3 address: %p\n", &rule3)
	fmt.Printf("All rules identical: %t\n", rule1.String() == rule2.String() && rule2.String() == rule3.String())
	fmt.Printf("Manager has %d unique rules cached\n\n", len(manager.GetRules()))

	// Demonstrate template integration
	fmt.Println("=== Template Integration Demo ===")
	
	// Create a practical example with a card component
	var cardStylesheet css.Stylesheet
	cardStylesheet.Add(
		tailwind.Bg("white"),
		tailwind.Text("gray-900"),
		tailwind.P("6"),
		tailwind.M("4"),
		tailwind.W("96"),
		tailwind.DisplayFlex(),
		tailwind.FlexColClass(),
		tailwind.JustifyBetweenClass(),
		tailwind.RelativeClass(),
		css.RuleSet(".card",
			css.Set(css.BorderRadius, css.Px(8)),
			css.Set("box-shadow", css.Raw("0 4px 6px -1px rgba(0, 0, 0, 0.1)")),
		),
		// Button utilities
		tailwind.Bg("blue-600"),
		tailwind.Text("white"),
		tailwind.Px("4"),
		tailwind.Py("2"),
		tailwind.FontMedium(),
		tailwind.TextSm(),
		css.RuleSet(".btn",
			css.Set(css.BorderRadius, css.Px(4)),
			css.Set("transition", css.Raw("background-color 0.2s")),
		),
		css.RuleSet(".btn:hover",
			css.Set(css.BackgroundColor, css.Hex("#2563eb")),
		),
	)

	htmlTemplate := `<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Tailwind CSS Demo</title>
	<style>{{.CSS}}</style>
</head>
<body class="bg-gray-100">
	<div class="flex justify-center items-center" style="min-height: 100vh;">
		<div class="card bg-white text-gray-900 p-6 m-4 w-96 flex flex-col relative">
			<h1 class="text-lg font-bold text-center">Tailwind Demo</h1>
			<p class="text-sm text-gray-800 py-2">
				This card uses Tailwind CSS utilities generated with type safety in Go!
			</p>
			<div class="flex justify-center">
				<button class="btn bg-blue-600 text-white px-4 py-2 font-medium text-sm">
					Click Me
				</button>
			</div>
		</div>
	</div>
</body>
</html>`

	tmpl := template.Must(template.New("demo").Parse(htmlTemplate))
	data := struct {
		CSS template.CSS
	}{
		CSS: template.CSS(cardStylesheet.String()),
	}

	fmt.Println("HTML Template with Tailwind Utilities:")
	tmpl.Execute(os.Stdout, data)
}