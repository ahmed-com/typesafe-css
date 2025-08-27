package css

import (
	"strings"
	"testing"
)

func TestValues(t *testing.T) {
	tests := []struct {
		name     string
		value    Value
		expected string
	}{
		{"Keyword", Keyword("block"), "block"},
		{"Raw", Raw("url('/image.png')"), "url('/image.png')"},
		{"Px", Px(16), "16px"},
		{"Rem", Rem(1.5), "1.5rem"},
		{"Em", Em(2.0), "2em"},
		{"Percent", Percent(50.0), "50%"},
		{"Hex color", Hex("#ff0000"), "#ff0000"},
		{"RGB color", RGB(255, 0, 0), "rgb(255 0 0)"},
		{"RGBA color", RGBA(255, 0, 0, 128), "rgba(255 0 0 / 0.502)"},
		{"CSS var", Var("--primary"), "var(--primary)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.value.String(); got != tt.expected {
				t.Errorf("Value.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestTrim0(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1.5000rem", "1.5rem"},
		{"2.0000em", "2em"},
		{"50.000000%", "50%"},
		{"1.2500px", "1.25px"},
		{"16px", "16px"}, // no decimal
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := trim0(tt.input); got != tt.expected {
				t.Errorf("trim0(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestDeclaration(t *testing.T) {
	decl := Set(Display, DisplayBlock)
	expected := "display:block"
	if got := decl.String(); got != expected {
		t.Errorf("Decl.String() = %v, want %v", got, expected)
	}
}

func TestRule(t *testing.T) {
	rule := RuleSet(".test",
		Set(Display, DisplayBlock),
		Set(ColorP, Hex("#000")),
	)
	
	result := rule.String()
	
	// Check that it contains the selector and properties
	if !strings.Contains(result, ".test") {
		t.Errorf("Rule should contain selector '.test', got: %v", result)
	}
	if !strings.Contains(result, "display:block") {
		t.Errorf("Rule should contain 'display:block', got: %v", result)
	}
	if !strings.Contains(result, "color:#000") {
		t.Errorf("Rule should contain 'color:#000', got: %v", result)
	}
	if !strings.HasSuffix(result, "}") {
		t.Errorf("Rule should end with '}', got: %v", result)
	}
}

func TestAtRule(t *testing.T) {
	atRule := AtRule{
		Name:   "media",
		Params: "(max-width: 640px)",
		Body: []Item{
			RuleSet(".btn", Set(Display, DisplayBlock)),
		},
	}
	
	result := atRule.String()
	
	if !strings.Contains(result, "@media") {
		t.Errorf("AtRule should contain '@media', got: %v", result)
	}
	if !strings.Contains(result, "(max-width: 640px)") {
		t.Errorf("AtRule should contain params, got: %v", result)
	}
	if !strings.Contains(result, ".btn") {
		t.Errorf("AtRule should contain nested rule, got: %v", result)
	}
}

func TestShorthands(t *testing.T) {
	tests := []struct {
		name     string
		value    Value
		contains []string // parts that should be in the output
	}{
		{
			"PadXY",
			PadXY(Px(8), Px(12)),
			[]string{"12px", "8px"},
		},
		{
			"BorderShorthand",
			BorderShorthand(Px(1), BorderSolid, Hex("#ccc")),
			[]string{"1px", "solid", "#ccc"},
		},
		{
			"Multiple values",
			Multiple(Px(10), Px(20), Px(30)),
			[]string{"10px", "20px", "30px"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.value.String()
			for _, part := range tt.contains {
				if !strings.Contains(result, part) {
					t.Errorf("Value should contain '%v', got: %v", part, result)
				}
			}
		})
	}
}

func TestFlexHelpers(t *testing.T) {
	centerDecls := FlexCenter()
	
	if len(centerDecls) != 3 {
		t.Errorf("FlexCenter should return 3 declarations, got %d", len(centerDecls))
	}
	
	// Check that it includes display: flex
	foundDisplayFlex := false
	for _, decl := range centerDecls {
		if decl.Property == Display && decl.Value.String() == "flex" {
			foundDisplayFlex = true
			break
		}
	}
	if !foundDisplayFlex {
		t.Errorf("FlexCenter should include display: flex")
	}
}

func TestStylesheet(t *testing.T) {
	var sheet Stylesheet
	rule1 := RuleSet(".test1", Set(Display, DisplayBlock))
	rule2 := RuleSet(".test2", Set(ColorP, Hex("#000")))
	
	sheet.Add(rule1, rule2)
	
	if len(sheet.Items) != 2 {
		t.Errorf("Stylesheet should have 2 items, got %d", len(sheet.Items))
	}
	
	result := sheet.String()
	if !strings.Contains(result, ".test1") {
		t.Errorf("Stylesheet should contain '.test1', got: %v", result)
	}
	if !strings.Contains(result, ".test2") {
		t.Errorf("Stylesheet should contain '.test2', got: %v", result)
	}
}