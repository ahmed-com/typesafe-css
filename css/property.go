package css

// Property represents a CSS property name.
type Property string

// Decl represents a CSS declaration (property: value).
type Decl struct {
	Property Property
	Value    Value
}

// Rule represents a CSS rule with selector and declarations.
type Rule struct {
	Selector string
	Decls    []Decl
}

// AtRule represents CSS at-rules like @media, @keyframes, etc.
type AtRule struct {
	Name   string // e.g., "media", "keyframes"
	Params string // e.g., "(max-width: 640px)"
	Body   []Item // nested rules/declarations
}

// Item represents any CSS item (Rule or AtRule).
type Item interface {
	String() string
}

// Stylesheet represents a collection of CSS items.
type Stylesheet struct {
	Items []Item
}

// Set creates a CSS declaration with the given property and value.
func Set(p Property, v Value) Decl {
	return Decl{Property: p, Value: v}
}

// RuleSet creates a CSS rule with the given selector and declarations.
func RuleSet(selector string, decls ...Decl) Rule {
	return Rule{
		Selector: selector,
		Decls:    append([]Decl(nil), decls...), // defensive copy
	}
}

// Add appends items to the stylesheet.
func (s *Stylesheet) Add(items ...Item) {
	s.Items = append(s.Items, items...)
}

// Common properties (hand-curated, minimal set for day-1)
const (
	Display    Property = "display"
	ColorP     Property = "color"
	Position   Property = "position"
	Padding    Property = "padding"
	Margin     Property = "margin"
	FontSize   Property = "font-size"
	FontFamily Property = "font-family"
	Width      Property = "width"
	Height     Property = "height"
	Background Property = "background"
	BackgroundColor Property = "background-color"
	Border     Property = "border"
	BorderRadius Property = "border-radius"
	TextAlign  Property = "text-align"
	FlexDirection Property = "flex-direction"
	JustifyContent Property = "justify-content"
	AlignItems Property = "align-items"
)