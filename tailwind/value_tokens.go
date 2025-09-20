package tailwind

import "github.com/ahmed-com/typesafe-css/css"

// ValueToken represents a named Tailwind token that resolves to a css.Value.
type ValueToken struct {
	suffix string
	value  css.Value
}

func newValueToken(suffix string, value css.Value) ValueToken {
	return ValueToken{suffix: suffix, value: value}
}

// Suffix returns the Tailwind class suffix for this token (e.g. "4", "lg").
func (t ValueToken) Suffix() string { return t.suffix }

// Value returns the CSS value for this token.
func (t ValueToken) Value() css.Value { return t.value }

// WithValue returns a copy of the token with the supplied css.Value.
func (t ValueToken) WithValue(value css.Value) ValueToken {
	t.value = value
	return t
}

// IsZero reports whether the token is unset.
func (t ValueToken) IsZero() bool {
	return t.suffix == "" && t.value == nil
}
