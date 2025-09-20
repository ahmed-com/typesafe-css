package tailwind

import (
	"github.com/ahmed-com/typesafe-css/css"
	"github.com/ahmed-com/typesafe-css/cssgen"
)

// BorderRadiusScale captures Tailwind's border-radius tokens.
type BorderRadiusScale struct {
	None    ValueToken
	Sm      ValueToken
	Default ValueToken
	Md      ValueToken
	Lg      ValueToken
	Xl      ValueToken
	X2L     ValueToken
	X3L     ValueToken
	Full    ValueToken
}

func (s BorderRadiusScale) tokens() []ValueToken {
	return []ValueToken{
		s.None,
		s.Sm,
		s.Default,
		s.Md,
		s.Lg,
		s.Xl,
		s.X2L,
		s.X3L,
		s.Full,
	}
}

func buildBorderRadiusScale() BorderRadiusScale {
	return BorderRadiusScale{
		None:    newValueToken("none", css.Px(0)),
		Sm:      newValueToken("sm", css.Rem(0.125)),
		Default: newValueToken("", css.Rem(0.25)),
		Md:      newValueToken("md", css.Rem(0.375)),
		Lg:      newValueToken("lg", css.Rem(0.5)),
		Xl:      newValueToken("xl", css.Rem(0.75)),
		X2L:     newValueToken("2xl", css.Rem(1)),
		X3L:     newValueToken("3xl", css.Rem(1.5)),
		Full:    newValueToken("full", css.Px(9999)),
	}
}

// BorderWidthScale captures Tailwind's border-width tokens.
type BorderWidthScale struct {
	Default ValueToken
	W0      ValueToken
	W2      ValueToken
	W4      ValueToken
	W8      ValueToken
}

func (s BorderWidthScale) tokens() []ValueToken {
	return []ValueToken{
		s.Default,
		s.W0,
		s.W2,
		s.W4,
		s.W8,
	}
}

func buildBorderWidthScale() BorderWidthScale {
	return BorderWidthScale{
		Default: newValueToken("", css.Px(1)),
		W0:      newValueToken("0", css.Px(0)),
		W2:      newValueToken("2", css.Px(2)),
		W4:      newValueToken("4", css.Px(4)),
		W8:      newValueToken("8", css.Px(8)),
	}
}

// BoxShadowScale captures Tailwind's box-shadow tokens.
type BoxShadowScale struct {
	Sm      ValueToken
	Default ValueToken
	Md      ValueToken
	Lg      ValueToken
	Xl      ValueToken
	X2L     ValueToken
	Inner   ValueToken
	None    ValueToken
}

func (s BoxShadowScale) tokens() []ValueToken {
	return []ValueToken{
		s.Sm,
		s.Default,
		s.Md,
		s.Lg,
		s.Xl,
		s.X2L,
		s.Inner,
		s.None,
	}
}

func buildBoxShadowScale() BoxShadowScale {
	return BoxShadowScale{
		Sm:      newValueToken("sm", css.Raw("0 1px 2px 0 rgb(0 0 0 / 0.05)")),
		Default: newValueToken("", css.Raw("0 1px 3px 0 rgb(0 0 0 / 0.1), 0 1px 2px -1px rgb(0 0 0 / 0.1)")),
		Md:      newValueToken("md", css.Raw("0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1)")),
		Lg:      newValueToken("lg", css.Raw("0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1)")),
		Xl:      newValueToken("xl", css.Raw("0 20px 25px -5px rgb(0 0 0 / 0.1), 0 8px 10px -6px rgb(0 0 0 / 0.1)")),
		X2L:     newValueToken("2xl", css.Raw("0 25px 50px -12px rgb(0 0 0 / 0.25)")),
		Inner:   newValueToken("inner", css.Raw("inset 0 2px 4px 0 rgb(0 0 0 / 0.05)")),
		None:    newValueToken("none", css.Keyword(string(cssgen.BoxShadowValNone))),
	}
}

// OpacityScale captures Tailwind's opacity tokens.
type OpacityScale struct {
	P0   ValueToken
	P5   ValueToken
	P10  ValueToken
	P15  ValueToken
	P20  ValueToken
	P25  ValueToken
	P30  ValueToken
	P35  ValueToken
	P40  ValueToken
	P45  ValueToken
	P50  ValueToken
	P55  ValueToken
	P60  ValueToken
	P65  ValueToken
	P70  ValueToken
	P75  ValueToken
	P80  ValueToken
	P85  ValueToken
	P90  ValueToken
	P95  ValueToken
	P100 ValueToken
}

func (s OpacityScale) tokens() []ValueToken {
	return []ValueToken{
		s.P0,
		s.P5,
		s.P10,
		s.P15,
		s.P20,
		s.P25,
		s.P30,
		s.P35,
		s.P40,
		s.P45,
		s.P50,
		s.P55,
		s.P60,
		s.P65,
		s.P70,
		s.P75,
		s.P80,
		s.P85,
		s.P90,
		s.P95,
		s.P100,
	}
}

func buildOpacityScale() OpacityScale {
	return OpacityScale{
		P0:   newValueToken("0", css.Raw("0")),
		P5:   newValueToken("5", css.Raw("0.05")),
		P10:  newValueToken("10", css.Raw("0.1")),
		P15:  newValueToken("15", css.Raw("0.15")),
		P20:  newValueToken("20", css.Raw("0.2")),
		P25:  newValueToken("25", css.Raw("0.25")),
		P30:  newValueToken("30", css.Raw("0.3")),
		P35:  newValueToken("35", css.Raw("0.35")),
		P40:  newValueToken("40", css.Raw("0.4")),
		P45:  newValueToken("45", css.Raw("0.45")),
		P50:  newValueToken("50", css.Raw("0.5")),
		P55:  newValueToken("55", css.Raw("0.55")),
		P60:  newValueToken("60", css.Raw("0.6")),
		P65:  newValueToken("65", css.Raw("0.65")),
		P70:  newValueToken("70", css.Raw("0.7")),
		P75:  newValueToken("75", css.Raw("0.75")),
		P80:  newValueToken("80", css.Raw("0.8")),
		P85:  newValueToken("85", css.Raw("0.85")),
		P90:  newValueToken("90", css.Raw("0.9")),
		P95:  newValueToken("95", css.Raw("0.95")),
		P100: newValueToken("100", css.Raw("1")),
	}
}

// ZIndexScale captures Tailwind's z-index tokens.
type ZIndexScale struct {
	Auto ValueToken
	Z0   ValueToken
	Z10  ValueToken
	Z20  ValueToken
	Z30  ValueToken
	Z40  ValueToken
	Z50  ValueToken
}

func (s ZIndexScale) tokens() []ValueToken {
	return []ValueToken{
		s.Auto,
		s.Z0,
		s.Z10,
		s.Z20,
		s.Z30,
		s.Z40,
		s.Z50,
	}
}

func buildZIndexScale() ZIndexScale {
	return ZIndexScale{
		Auto: newValueToken("auto", css.Keyword(string(cssgen.ZIndexValAuto))),
		Z0:   newValueToken("0", css.Raw("0")),
		Z10:  newValueToken("10", css.Raw("10")),
		Z20:  newValueToken("20", css.Raw("20")),
		Z30:  newValueToken("30", css.Raw("30")),
		Z40:  newValueToken("40", css.Raw("40")),
		Z50:  newValueToken("50", css.Raw("50")),
	}
}

// BlurScale captures Tailwind's blur tokens.
type BlurScale struct {
	Zero    ValueToken
	None    ValueToken
	Sm      ValueToken
	Default ValueToken
	Md      ValueToken
	Lg      ValueToken
	Xl      ValueToken
	X2L     ValueToken
	X3L     ValueToken
}

func (s BlurScale) tokens() []ValueToken {
	return []ValueToken{
		s.Zero,
		s.None,
		s.Sm,
		s.Default,
		s.Md,
		s.Lg,
		s.Xl,
		s.X2L,
		s.X3L,
	}
}

func buildBlurScale() BlurScale {
	return BlurScale{
		Zero:    newValueToken("0", css.Raw("0")),
		None:    newValueToken("none", css.Raw("")),
		Sm:      newValueToken("sm", css.Raw("4px")),
		Default: newValueToken("", css.Raw("8px")),
		Md:      newValueToken("md", css.Raw("12px")),
		Lg:      newValueToken("lg", css.Raw("16px")),
		Xl:      newValueToken("xl", css.Raw("24px")),
		X2L:     newValueToken("2xl", css.Raw("40px")),
		X3L:     newValueToken("3xl", css.Raw("64px")),
	}
}

// BrightnessScale captures Tailwind's brightness tokens.
type BrightnessScale struct {
	P0   ValueToken
	P50  ValueToken
	P75  ValueToken
	P90  ValueToken
	P95  ValueToken
	P100 ValueToken
	P105 ValueToken
	P110 ValueToken
	P125 ValueToken
	P150 ValueToken
	P200 ValueToken
}

func (s BrightnessScale) tokens() []ValueToken {
	return []ValueToken{
		s.P0,
		s.P50,
		s.P75,
		s.P90,
		s.P95,
		s.P100,
		s.P105,
		s.P110,
		s.P125,
		s.P150,
		s.P200,
	}
}

func buildBrightnessScale() BrightnessScale {
	return BrightnessScale{
		P0:   newValueToken("0", css.Raw("0")),
		P50:  newValueToken("50", css.Raw(".5")),
		P75:  newValueToken("75", css.Raw(".75")),
		P90:  newValueToken("90", css.Raw(".9")),
		P95:  newValueToken("95", css.Raw(".95")),
		P100: newValueToken("100", css.Raw("1")),
		P105: newValueToken("105", css.Raw("1.05")),
		P110: newValueToken("110", css.Raw("1.1")),
		P125: newValueToken("125", css.Raw("1.25")),
		P150: newValueToken("150", css.Raw("1.5")),
		P200: newValueToken("200", css.Raw("2")),
	}
}

// ContrastScale captures Tailwind's contrast tokens.
type ContrastScale struct {
	P0   ValueToken
	P50  ValueToken
	P75  ValueToken
	P100 ValueToken
	P125 ValueToken
	P150 ValueToken
	P200 ValueToken
}

func (s ContrastScale) tokens() []ValueToken {
	return []ValueToken{
		s.P0,
		s.P50,
		s.P75,
		s.P100,
		s.P125,
		s.P150,
		s.P200,
	}
}

func buildContrastScale() ContrastScale {
	return ContrastScale{
		P0:   newValueToken("0", css.Raw("0")),
		P50:  newValueToken("50", css.Raw(".5")),
		P75:  newValueToken("75", css.Raw(".75")),
		P100: newValueToken("100", css.Raw("1")),
		P125: newValueToken("125", css.Raw("1.25")),
		P150: newValueToken("150", css.Raw("1.5")),
		P200: newValueToken("200", css.Raw("2")),
	}
}

// SaturateScale captures Tailwind's saturate tokens.
type SaturateScale struct {
	P0   ValueToken
	P50  ValueToken
	P100 ValueToken
	P150 ValueToken
	P200 ValueToken
}

func (s SaturateScale) tokens() []ValueToken {
	return []ValueToken{
		s.P0,
		s.P50,
		s.P100,
		s.P150,
		s.P200,
	}
}

func buildSaturateScale() SaturateScale {
	return SaturateScale{
		P0:   newValueToken("0", css.Raw("0")),
		P50:  newValueToken("50", css.Raw(".5")),
		P100: newValueToken("100", css.Raw("1")),
		P150: newValueToken("150", css.Raw("1.5")),
		P200: newValueToken("200", css.Raw("2")),
	}
}

// PercentageScale captures percentage-based filter tokens (grayscale, invert, sepia).
type PercentageScale struct {
	Zero    ValueToken
	Default ValueToken
}

func (s PercentageScale) tokens() []ValueToken {
	return []ValueToken{s.Zero, s.Default}
}

func buildGrayscaleScale() PercentageScale {
	return PercentageScale{
		Zero:    newValueToken("0", css.Raw("0")),
		Default: newValueToken("", css.Raw("100%")),
	}
}

func buildInvertScale() PercentageScale {
	return PercentageScale{
		Zero:    newValueToken("0", css.Raw("0")),
		Default: newValueToken("", css.Raw("100%")),
	}
}

func buildSepiaScale() PercentageScale {
	return PercentageScale{
		Zero:    newValueToken("0", css.Raw("0")),
		Default: newValueToken("", css.Raw("100%")),
	}
}
