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
		tailwind.Flex(),
		tailwind.Block(),
		tailwind.Hidden(),
		tailwind.JustifyCenter(),
		tailwind.ItemsCenter(),
		tailwind.FlexCol(),

		// Spacing utilities - demonstrate deduplication
		tailwind.P4(),     // padding: 1rem
		tailwind.Px6(),    // padding-left/right: 1.5rem
		tailwind.Py2(),    // padding-top/bottom: 0.5rem
		tailwind.M8(),     // margin: 2rem
		tailwind.MxAuto(), // margin-left/right: auto
		tailwind.W64(),    // width: 16rem
		tailwind.WFull(),  // width: 100%
		tailwind.H32(),    // height: 8rem
		tailwind.HScreen(), // height: 100vh

		// Color utilities
		tailwind.BgBlue500(),    // background-color: #3b82f6
		tailwind.BgGray100(),    // background-color: #f3f4f6
		tailwind.TextWhite(),    // color: #ffffff
		tailwind.TextGray800(),  // color: #1f2937
		tailwind.Bg("#bada55"),  // arbitrary color

		// Typography utilities
		tailwind.TextLg(),          // font-size: 1.125rem
		tailwind.TextCenter(),      // text-align: center
		tailwind.FontBold(),        // font-weight: 700
		tailwind.FontSemibold(),    // font-weight: 600
		tailwind.TextSm(),          // font-size: 0.875rem

		// Position utilities
		tailwind.Relative(),   // position: relative
		tailwind.Absolute(),   // position: absolute
		
		// New enhanced utilities
		// Border radius
		tailwind.RoundedLg(),     // border-radius: 0.5rem
		tailwind.RoundedFull(),   // border-radius: 9999px
		tailwind.Rounded("md"),   // border-radius: 0.375rem
		
		// Box shadows
		tailwind.ShadowLg(),      // box-shadow: 0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1)
		tailwind.ShadowSm(),      // box-shadow: 0 1px 2px 0 rgb(0 0 0 / 0.05)
		tailwind.Shadow("xl"),    // box-shadow: 0 20px 25px -5px rgb(0 0 0 / 0.1), 0 8px 10px -6px rgb(0 0 0 / 0.1)
		
		// Opacity
		tailwind.Opacity75(),     // opacity: 0.75
		tailwind.Opacity50(),     // opacity: 0.5
		
		// Z-index
		tailwind.Z10(),           // z-index: 10
		tailwind.Z50(),           // z-index: 50
		
		// Filter effects
		tailwind.BlurClass("sm"),      // filter: blur(4px)
		tailwind.BrightnessClass("110"), // filter: brightness(1.1)
		tailwind.GrayscaleClass(""),     // filter: grayscale(100%)
		
		// Enhanced colors from full palette
		tailwind.BgSlate500(),    // background-color: #64748b
		tailwind.BgIndigo500(),   // background-color: #6366f1
		tailwind.BgEmerald500(),  // background-color: #10b981
		tailwind.Bg("rose-500"),  // background-color: #f43f5e (using generic function)
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
		BorderRadius: map[string]css.Length{
			"huge": css.Rem(3), // 48px - very large border radius
		},
		BoxShadow: map[string]string{
			"glow": "0 0 20px rgba(99, 102, 241, 0.5)", // Custom glow effect
		},
		Opacity: map[string]string{
			"15": "0.15", // Custom opacity value
		},
	})

	customManager := tailwind.WithCustomTheme(customTheme)
	var customStylesheet css.Stylesheet

	customStylesheet.Add(
		tailwind.BackgroundColor(customManager, "brand"),
		tailwind.TextColor(customManager, "danger"),
		tailwind.Padding(customManager, "xs"),
		tailwind.Margin(customManager, "xl"),
		// New enhanced utilities
		tailwind.BorderRadius(customManager, "huge"),
		tailwind.BoxShadow(customManager, "glow"),
		tailwind.Opacity(customManager, "15"),
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
	rule1 := tailwind.P4()
	rule2 := tailwind.P4() 
	rule3 := tailwind.P4()
	
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
		tailwind.BgWhite(),
		tailwind.TextGray900(),
		tailwind.P6(),
		tailwind.M4(),
		tailwind.W96(),
		tailwind.Flex(),
		tailwind.FlexCol(),
		tailwind.JustifyBetween(),
		tailwind.Relative(),
		css.RuleSet(".card",
			css.Set(css.BorderRadius, css.Px(8)),
			css.Set("box-shadow", css.Raw("0 4px 6px -1px rgba(0, 0, 0, 0.1)")),
		),
		// Button utilities
		tailwind.BgBlue600(),
		tailwind.TextWhite(),
		tailwind.Px4(),
		tailwind.Py2(),
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