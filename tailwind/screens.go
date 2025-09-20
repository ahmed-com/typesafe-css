package tailwind

// Screens defines named responsive breakpoints.
type Screens struct {
	Sm  string
	Md  string
	Lg  string
	Xl  string
	X2L string
}

func buildScreens() Screens {
	return Screens{
		Sm:  "(min-width: 640px)",
		Md:  "(min-width: 768px)",
		Lg:  "(min-width: 1024px)",
		Xl:  "(min-width: 1280px)",
		X2L: "(min-width: 1536px)",
	}
}

func (s Screens) tokens() map[string]string {
	return map[string]string{
		"sm":  s.Sm,
		"md":  s.Md,
		"lg":  s.Lg,
		"xl":  s.Xl,
		"2xl": s.X2L,
	}
}
