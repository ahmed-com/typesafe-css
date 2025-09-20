package tailwind

import "github.com/ahmed-com/typesafe-css/css"

// SpacingScale captures Tailwind's spacing tokens as strongly typed values.
type SpacingScale struct {
	Px   ValueToken
	P0   ValueToken
	P0_5 ValueToken
	P1   ValueToken
	P1_5 ValueToken
	P2   ValueToken
	P2_5 ValueToken
	P3   ValueToken
	P3_5 ValueToken
	P4   ValueToken
	P5   ValueToken
	P6   ValueToken
	P7   ValueToken
	P8   ValueToken
	P9   ValueToken
	P10  ValueToken
	P11  ValueToken
	P12  ValueToken
	P14  ValueToken
	P16  ValueToken
	P20  ValueToken
	P24  ValueToken
	P28  ValueToken
	P32  ValueToken
	P36  ValueToken
	P40  ValueToken
	P44  ValueToken
	P48  ValueToken
	P52  ValueToken
	P56  ValueToken
	P60  ValueToken
	P64  ValueToken
	P72  ValueToken
	P80  ValueToken
	P96  ValueToken
}

func (s SpacingScale) tokens() []ValueToken {
	return []ValueToken{
		s.Px,
		s.P0,
		s.P0_5,
		s.P1,
		s.P1_5,
		s.P2,
		s.P2_5,
		s.P3,
		s.P3_5,
		s.P4,
		s.P5,
		s.P6,
		s.P7,
		s.P8,
		s.P9,
		s.P10,
		s.P11,
		s.P12,
		s.P14,
		s.P16,
		s.P20,
		s.P24,
		s.P28,
		s.P32,
		s.P36,
		s.P40,
		s.P44,
		s.P48,
		s.P52,
		s.P56,
		s.P60,
		s.P64,
		s.P72,
		s.P80,
		s.P96,
	}
}

func buildSpacingScale() SpacingScale {
	return SpacingScale{
		Px:   newValueToken("px", css.Px(1)),
		P0:   newValueToken("0", css.Px(0)),
		P0_5: newValueToken("0.5", css.Rem(0.125)),
		P1:   newValueToken("1", css.Rem(0.25)),
		P1_5: newValueToken("1.5", css.Rem(0.375)),
		P2:   newValueToken("2", css.Rem(0.5)),
		P2_5: newValueToken("2.5", css.Rem(0.625)),
		P3:   newValueToken("3", css.Rem(0.75)),
		P3_5: newValueToken("3.5", css.Rem(0.875)),
		P4:   newValueToken("4", css.Rem(1)),
		P5:   newValueToken("5", css.Rem(1.25)),
		P6:   newValueToken("6", css.Rem(1.5)),
		P7:   newValueToken("7", css.Rem(1.75)),
		P8:   newValueToken("8", css.Rem(2)),
		P9:   newValueToken("9", css.Rem(2.25)),
		P10:  newValueToken("10", css.Rem(2.5)),
		P11:  newValueToken("11", css.Rem(2.75)),
		P12:  newValueToken("12", css.Rem(3)),
		P14:  newValueToken("14", css.Rem(3.5)),
		P16:  newValueToken("16", css.Rem(4)),
		P20:  newValueToken("20", css.Rem(5)),
		P24:  newValueToken("24", css.Rem(6)),
		P28:  newValueToken("28", css.Rem(7)),
		P32:  newValueToken("32", css.Rem(8)),
		P36:  newValueToken("36", css.Rem(9)),
		P40:  newValueToken("40", css.Rem(10)),
		P44:  newValueToken("44", css.Rem(11)),
		P48:  newValueToken("48", css.Rem(12)),
		P52:  newValueToken("52", css.Rem(13)),
		P56:  newValueToken("56", css.Rem(14)),
		P60:  newValueToken("60", css.Rem(15)),
		P64:  newValueToken("64", css.Rem(16)),
		P72:  newValueToken("72", css.Rem(18)),
		P80:  newValueToken("80", css.Rem(20)),
		P96:  newValueToken("96", css.Rem(24)),
	}
}
