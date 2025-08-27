package css

import (
	"strings"
)

// String implementations for CSS serialization

func (d Decl) String() string {
	return string(d.Property) + ":" + d.Value.String()
}

func (r Rule) String() string {
	if len(r.Decls) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString(r.Selector)
	b.WriteString("{")
	
	for i, decl := range r.Decls {
		if i > 0 {
			b.WriteString(";")
		}
		b.WriteString(decl.String())
	}
	
	b.WriteString("}")
	return b.String()
}

func (a AtRule) String() string {
	var b strings.Builder
	b.WriteString("@")
	b.WriteString(a.Name)
	
	if a.Params != "" {
		b.WriteString(" ")
		b.WriteString(a.Params)
	}
	
	if len(a.Body) > 0 {
		b.WriteString("{")
		for _, item := range a.Body {
			b.WriteString(item.String())
		}
		b.WriteString("}")
	} else {
		b.WriteString(";")
	}
	
	return b.String()
}

func (s Stylesheet) String() string {
	var b strings.Builder
	for _, item := range s.Items {
		b.WriteString(item.String())
	}
	return b.String()
}

// CSS creates a stylesheet string from items
func CSS(items ...Item) string {
	var b strings.Builder
	for _, item := range items {
		b.WriteString(item.String())
	}
	return b.String()
}

// Minified versions (optional for later)
// For now, we keep it simple and readable

// PrettyCSS creates a formatted CSS string with proper indentation
func PrettyCSS(items ...Item) string {
	var b strings.Builder
	for i, item := range items {
		if i > 0 {
			b.WriteString("\n\n")
		}
		b.WriteString(prettyItem(item, 0))
	}
	return b.String()
}

func prettyItem(item Item, indent int) string {
	switch v := item.(type) {
	case Rule:
		return prettyRule(v, indent)
	case AtRule:
		return prettyAtRule(v, indent)
	default:
		return item.String()
	}
}

func prettyRule(r Rule, indent int) string {
	if len(r.Decls) == 0 {
		return ""
	}

	var b strings.Builder
	indentStr := strings.Repeat("  ", indent)
	
	b.WriteString(indentStr)
	b.WriteString(r.Selector)
	b.WriteString(" {\n")
	
	for _, decl := range r.Decls {
		b.WriteString(indentStr)
		b.WriteString("  ")
		b.WriteString(decl.String())
		b.WriteString(";\n")
	}
	
	b.WriteString(indentStr)
	b.WriteString("}")
	return b.String()
}

func prettyAtRule(a AtRule, indent int) string {
	var b strings.Builder
	indentStr := strings.Repeat("  ", indent)
	
	b.WriteString(indentStr)
	b.WriteString("@")
	b.WriteString(a.Name)
	
	if a.Params != "" {
		b.WriteString(" ")
		b.WriteString(a.Params)
	}
	
	if len(a.Body) > 0 {
		b.WriteString(" {\n")
		for _, item := range a.Body {
			b.WriteString(prettyItem(item, indent+1))
			b.WriteString("\n")
		}
		b.WriteString(indentStr)
		b.WriteString("}")
	} else {
		b.WriteString(";")
	}
	
	return b.String()
}