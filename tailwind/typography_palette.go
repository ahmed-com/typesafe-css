package tailwind

import "github.com/ahmed-com/typesafe-css/css"

// FontSizeScale captures Tailwind's font-size tokens.
type FontSizeScale struct {
	Xs   ValueToken
	Sm   ValueToken
	Base ValueToken
	Lg   ValueToken
	Xl   ValueToken
	X2L  ValueToken
	X3L  ValueToken
	X4L  ValueToken
	X5L  ValueToken
	X6L  ValueToken
	X7L  ValueToken
	X8L  ValueToken
	X9L  ValueToken
}

func (s FontSizeScale) tokens() []ValueToken {
	return []ValueToken{
		s.Xs,
		s.Sm,
		s.Base,
		s.Lg,
		s.Xl,
		s.X2L,
		s.X3L,
		s.X4L,
		s.X5L,
		s.X6L,
		s.X7L,
		s.X8L,
		s.X9L,
	}
}

func buildFontSizeScale() FontSizeScale {
	return FontSizeScale{
		Xs:   newValueToken("xs", css.Rem(0.75)),
		Sm:   newValueToken("sm", css.Rem(0.875)),
		Base: newValueToken("base", css.Rem(1)),
		Lg:   newValueToken("lg", css.Rem(1.125)),
		Xl:   newValueToken("xl", css.Rem(1.25)),
		X2L:  newValueToken("2xl", css.Rem(1.5)),
		X3L:  newValueToken("3xl", css.Rem(1.875)),
		X4L:  newValueToken("4xl", css.Rem(2.25)),
		X5L:  newValueToken("5xl", css.Rem(3)),
		X6L:  newValueToken("6xl", css.Rem(3.75)),
		X7L:  newValueToken("7xl", css.Rem(4.5)),
		X8L:  newValueToken("8xl", css.Rem(6)),
		X9L:  newValueToken("9xl", css.Rem(8)),
	}
}

// FontWeightScale captures Tailwind's font-weight tokens.
type FontWeightScale struct {
	Thin       ValueToken
	ExtraLight ValueToken
	Light      ValueToken
	Normal     ValueToken
	Medium     ValueToken
	SemiBold   ValueToken
	Bold       ValueToken
	ExtraBold  ValueToken
	Black      ValueToken
}

func (s FontWeightScale) tokens() []ValueToken {
	return []ValueToken{
		s.Thin,
		s.ExtraLight,
		s.Light,
		s.Normal,
		s.Medium,
		s.SemiBold,
		s.Bold,
		s.ExtraBold,
		s.Black,
	}
}

func buildFontWeightScale() FontWeightScale {
	return FontWeightScale{
		Thin:       newValueToken("thin", css.Keyword("100")),
		ExtraLight: newValueToken("extralight", css.Keyword("200")),
		Light:      newValueToken("light", css.Keyword("300")),
		Normal:     newValueToken("normal", css.Keyword("400")),
		Medium:     newValueToken("medium", css.Keyword("500")),
		SemiBold:   newValueToken("semibold", css.Keyword("600")),
		Bold:       newValueToken("bold", css.Keyword("700")),
		ExtraBold:  newValueToken("extrabold", css.Keyword("800")),
		Black:      newValueToken("black", css.Keyword("900")),
	}
}

// LineHeightScale captures Tailwind's line-height tokens.
type LineHeightScale struct {
	None    ValueToken
	Tight   ValueToken
	Snug    ValueToken
	Normal  ValueToken
	Relaxed ValueToken
	Loose   ValueToken
	L3      ValueToken
	L4      ValueToken
	L5      ValueToken
	L6      ValueToken
	L7      ValueToken
	L8      ValueToken
	L9      ValueToken
	L10     ValueToken
}

func (s LineHeightScale) tokens() []ValueToken {
	return []ValueToken{
		s.None,
		s.Tight,
		s.Snug,
		s.Normal,
		s.Relaxed,
		s.Loose,
		s.L3,
		s.L4,
		s.L5,
		s.L6,
		s.L7,
		s.L8,
		s.L9,
		s.L10,
	}
}

func buildLineHeightScale() LineHeightScale {
	return LineHeightScale{
		None:    newValueToken("none", css.Keyword("1")),
		Tight:   newValueToken("tight", css.Keyword("1.25")),
		Snug:    newValueToken("snug", css.Keyword("1.375")),
		Normal:  newValueToken("normal", css.Keyword("1.5")),
		Relaxed: newValueToken("relaxed", css.Keyword("1.625")),
		Loose:   newValueToken("loose", css.Keyword("2")),
		L3:      newValueToken("3", css.Raw(".75rem")),
		L4:      newValueToken("4", css.Raw("1rem")),
		L5:      newValueToken("5", css.Raw("1.25rem")),
		L6:      newValueToken("6", css.Raw("1.5rem")),
		L7:      newValueToken("7", css.Raw("1.75rem")),
		L8:      newValueToken("8", css.Raw("2rem")),
		L9:      newValueToken("9", css.Raw("2.25rem")),
		L10:     newValueToken("10", css.Raw("2.5rem")),
	}
}
