// Package main demonstrates MDN-validated CSS properties in action.
package main

import (
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/ahmed-com/typesafe-css/css"
	"github.com/ahmed-com/typesafe-css/cssgen"
)

func main() {
	// Modern CSS layout using MDN-validated properties
	modernCard := css.RuleSet(".modern-card",
		// New display values from MDN validation
		cssgen.SetDisplay(cssgen.DisplayValFlex),
		cssgen.SetFlexDirection(cssgen.FlexDirectionValColumn),
		
		// Enhanced alignment with new MDN values
		cssgen.SetJustifyContent(cssgen.JustifyContentValSpaceBetween),
		cssgen.SetAlignItems(cssgen.AlignItemsValStretch),
		
		// New text alignment option
		cssgen.SetTextAlign(cssgen.TextAlignValStart),
		
		// Modern overflow handling
		cssgen.SetOverflow(cssgen.OverflowValClip),
		
		// Additional manual properties
		css.Set("padding", css.Px(20)),
		css.Set("border-radius", css.Px(8)),
		css.Set("box-shadow", css.Raw("0 4px 6px rgba(0, 0, 0, 0.1)")),
	)

	// CSS Grid container with new display value
	gridContainer := css.RuleSet(".grid-container",
		cssgen.SetDisplay(cssgen.DisplayValGrid),
		css.Set("grid-template-columns", css.Raw("repeat(auto-fit, minmax(250px, 1fr))")),
		css.Set("gap", css.Rem(1.5)),
		css.Set("padding", css.Rem(2)),
	)

	// Interactive elements with expanded cursor options
	interactiveElements := css.RuleSet(".interactive",
		cssgen.SetCursor(cssgen.CursorValPointer),
		css.Set("transition", css.Raw("all 0.2s ease")),
	)

	// Hover states with more cursor types
	hoverStates := css.RuleSet(".resizable:hover",
		cssgen.SetCursor(cssgen.CursorValColResize),
	)

	zoomableHover := css.RuleSet(".zoomable:hover",
		cssgen.SetCursor(cssgen.CursorValZoomIn),
	)

	// Special layout using contents display (removes box from layout)
	wrapperRemoval := css.RuleSet(".layout-wrapper",
		cssgen.SetDisplay(cssgen.DisplayValContents),
	)

	// Flow root for containing floats
	flowRoot := css.RuleSet(".clearfix",
		cssgen.SetDisplay(cssgen.DisplayValFlowRoot),
	)

	// Typography with enhanced text alignment
	typography := css.RuleSet(".text-content",
		cssgen.SetTextAlign(cssgen.TextAlignValMatchParent),
		cssgen.SetFontWeight(cssgen.FontWeightVal400),
		css.Set("line-height", css.Raw("1.6")),
	)

	// Create stylesheet
	var stylesheet css.Stylesheet
	stylesheet.Add(
		modernCard,
		gridContainer,
		interactiveElements,
		hoverStates,
		zoomableHover,
		wrapperRemoval,
		flowRoot,
		typography,
	)

	// HTML template
	htmlTemplate := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>MDN-Validated CSS Demo</title>
    <style>
        {{.CSS}}
        
        /* Additional base styles */
        body {
            font-family: system-ui, -apple-system, sans-serif;
            margin: 0;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
        }
        
        .modern-card {
            background: white;
            margin: 1rem;
            min-height: 200px;
        }
        
        .resizable {
            background: #f0f0f0;
            padding: 1rem;
            border: 2px dashed #ccc;
            margin: 0.5rem;
        }
        
        .zoomable {
            background: #e6f3ff;
            padding: 1rem;
            border: 1px solid #0066cc;
            margin: 0.5rem;
        }
    </style>
</head>
<body>
    <div class="grid-container">
        <div class="modern-card interactive">
            <h2>MDN-Validated CSS Properties</h2>
            <div class="text-content">
                <p>This card demonstrates new CSS properties added through MDN validation:</p>
                <ul>
                    <li><strong>display: flex</strong> - Modern flexbox layout</li>
                    <li><strong>justify-content: space-between</strong> - Enhanced alignment</li>
                    <li><strong>text-align: start</strong> - Logical alignment</li>
                    <li><strong>overflow: clip</strong> - Modern overflow handling</li>
                </ul>
            </div>
        </div>

        <div class="modern-card">
            <h3>Interactive Elements</h3>
            <div class="resizable interactive">
                <p>Hover over me - I show a column resize cursor!</p>
            </div>
            <div class="zoomable interactive">
                <p>Hover over me - I show a zoom-in cursor!</p>
            </div>
        </div>

        <div class="modern-card">
            <h3>Special Display Types</h3>
            <div class="layout-wrapper">
                <div class="clearfix">
                    <p>This content uses <code>display: contents</code> for the wrapper and <code>display: flow-root</code> for clearfix.</p>
                </div>
            </div>
        </div>
    </div>
</body>
</html>`

	// Generate HTML
	tmpl, err := template.New("demo").Parse(htmlTemplate)
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		CSS string
	}{
		CSS: stylesheet.String(),
	}

	file, err := os.Create("mdn_demo.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, data); err != nil {
		log.Fatal(err)
	}

	fmt.Println("MDN-validated CSS demo generated: mdn_demo.html")
	fmt.Println("\nGenerated CSS includes these new MDN-validated properties:")
	fmt.Println("- display: contents, flow-root")
	fmt.Println("- text-align: start, match-parent")
	fmt.Println("- overflow: clip")
	fmt.Println("- cursor: col-resize, zoom-in, zoom-out, and many more")
	fmt.Println("- align-items: start, end, self-start, self-end, normal")
	fmt.Println("- justify-content: start, end, normal, stretch")
}