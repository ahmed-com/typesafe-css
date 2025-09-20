package tailwind

import (
	"fmt"

	"github.com/ahmed-com/typesafe-css/css"
	"github.com/ahmed-com/typesafe-css/cssgen"
)

// Opacity creates an opacity utility class.
// Example: Opacity(manager, "50") generates ".opacity-50 { opacity: 0.5; }"
func Opacity(manager *UtilityManager, opacityKey string) css.Rule {
	className := ClassName("opacity", opacityKey)

	return manager.GetOrCreateRule(className, func() css.Rule {
		var opacity css.Value
		if token, exists := manager.Theme().OpacityToken(opacityKey); exists {
			opacity = token.Value()
		} else {
			// Fallback to treating the key as a direct value
			opacity = css.Raw(opacityKey)
		}

		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set(cssgen.Opacity, opacity),
		)
	})
}

// BoxShadow creates a box shadow utility class.
// Example: BoxShadow(manager, "lg") generates ".shadow-lg { box-shadow: 0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1); }"
func BoxShadow(manager *UtilityManager, shadowKey string) css.Rule {
	className := ClassName("shadow", shadowKey)

	return manager.GetOrCreateRule(className, func() css.Rule {
		var shadow css.Value
		if token, exists := manager.Theme().BoxShadowToken(shadowKey); exists {
			shadow = token.Value()
		} else {
			// Fallback to treating the key as a direct value
			shadow = css.Raw(shadowKey)
		}

		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set(cssgen.BoxShadow, shadow),
		)
	})
}

// Blur creates a blur filter utility class.
// Example: Blur(manager, "sm") generates ".blur-sm { filter: blur(4px); }"
func Blur(manager *UtilityManager, blurKey string) css.Rule {
	className := ClassName("blur", blurKey)

	return manager.GetOrCreateRule(className, func() css.Rule {
		if token, exists := manager.Theme().BlurToken(blurKey); exists {
			value := token.Value().String()
			if value == "" {
				return css.RuleSet(fmt.Sprintf(".%s", className),
					cssgen.SetFilter(cssgen.FilterValNone),
				)
			}
			return css.RuleSet(fmt.Sprintf(".%s", className),
				css.Set(cssgen.Filter, css.Raw(fmt.Sprintf("blur(%s)", value))),
			)
		}
		// Fallback to treating the key as a direct value
		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set(cssgen.Filter, css.Raw(fmt.Sprintf("blur(%s)", blurKey))),
		)
	})
}

// Brightness creates a brightness filter utility class.
// Example: Brightness(manager, "110") generates ".brightness-110 { filter: brightness(1.1); }"
func Brightness(manager *UtilityManager, brightnessKey string) css.Rule {
	className := ClassName("brightness", brightnessKey)

	return manager.GetOrCreateRule(className, func() css.Rule {
		var brightness css.Value
		if token, exists := manager.Theme().BrightnessToken(brightnessKey); exists {
			brightness = css.Raw(fmt.Sprintf("brightness(%s)", token.Value().String()))
		} else {
			// Fallback to treating the key as a direct value
			brightness = css.Raw(fmt.Sprintf("brightness(%s)", brightnessKey))
		}

		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set(cssgen.Filter, brightness),
		)
	})
}

// Contrast creates a contrast filter utility class.
// Example: Contrast(manager, "125") generates ".contrast-125 { filter: contrast(1.25); }"
func Contrast(manager *UtilityManager, contrastKey string) css.Rule {
	className := ClassName("contrast", contrastKey)

	return manager.GetOrCreateRule(className, func() css.Rule {
		var contrast css.Value
		if token, exists := manager.Theme().ContrastToken(contrastKey); exists {
			contrast = css.Raw(fmt.Sprintf("contrast(%s)", token.Value().String()))
		} else {
			// Fallback to treating the key as a direct value
			contrast = css.Raw(fmt.Sprintf("contrast(%s)", contrastKey))
		}

		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set(cssgen.Filter, contrast),
		)
	})
}

// Grayscale creates a grayscale filter utility class.
// Example: Grayscale(manager, "") generates ".grayscale { filter: grayscale(100%); }"
func Grayscale(manager *UtilityManager, grayscaleKey string) css.Rule {
	className := ClassName("grayscale", grayscaleKey)

	return manager.GetOrCreateRule(className, func() css.Rule {
		var grayscale css.Value
		if token, exists := manager.Theme().GrayscaleToken(grayscaleKey); exists {
			grayscale = css.Raw(fmt.Sprintf("grayscale(%s)", token.Value().String()))
		} else {
			// Fallback to treating the key as a direct value
			grayscale = css.Raw(fmt.Sprintf("grayscale(%s)", grayscaleKey))
		}

		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set(cssgen.Filter, grayscale),
		)
	})
}

// Invert creates an invert filter utility class.
// Example: Invert(manager, "") generates ".invert { filter: invert(100%); }"
func Invert(manager *UtilityManager, invertKey string) css.Rule {
	className := ClassName("invert", invertKey)

	return manager.GetOrCreateRule(className, func() css.Rule {
		var invert css.Value
		if token, exists := manager.Theme().InvertToken(invertKey); exists {
			invert = css.Raw(fmt.Sprintf("invert(%s)", token.Value().String()))
		} else {
			// Fallback to treating the key as a direct value
			invert = css.Raw(fmt.Sprintf("invert(%s)", invertKey))
		}

		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set(cssgen.Filter, invert),
		)
	})
}

// Saturate creates a saturate filter utility class.
// Example: Saturate(manager, "150") generates ".saturate-150 { filter: saturate(1.5); }"
func Saturate(manager *UtilityManager, saturateKey string) css.Rule {
	className := ClassName("saturate", saturateKey)

	return manager.GetOrCreateRule(className, func() css.Rule {
		var saturate css.Value
		if token, exists := manager.Theme().SaturateToken(saturateKey); exists {
			saturate = css.Raw(fmt.Sprintf("saturate(%s)", token.Value().String()))
		} else {
			// Fallback to treating the key as a direct value
			saturate = css.Raw(fmt.Sprintf("saturate(%s)", saturateKey))
		}

		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set(cssgen.Filter, saturate),
		)
	})
}

// Sepia creates a sepia filter utility class.
// Example: Sepia(manager, "") generates ".sepia { filter: sepia(100%); }"
func Sepia(manager *UtilityManager, sepiaKey string) css.Rule {
	className := ClassName("sepia", sepiaKey)

	return manager.GetOrCreateRule(className, func() css.Rule {
		var sepia css.Value
		if token, exists := manager.Theme().SepiaToken(sepiaKey); exists {
			sepia = css.Raw(fmt.Sprintf("sepia(%s)", token.Value().String()))
		} else {
			// Fallback to treating the key as a direct value
			sepia = css.Raw(fmt.Sprintf("sepia(%s)", sepiaKey))
		}

		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set(cssgen.Filter, sepia),
		)
	})
}

// ZIndex creates a z-index utility class.
// Example: ZIndex(manager, "10") generates ".z-10 { z-index: 10; }"
func ZIndex(manager *UtilityManager, zKey string) css.Rule {
	className := ClassName("z", zKey)

	return manager.GetOrCreateRule(className, func() css.Rule {
		if zKey == "auto" {
			return css.RuleSet(fmt.Sprintf(".%s", className),
				cssgen.SetZIndex(cssgen.ZIndexValAuto),
			)
		}
		var zIndex css.Value
		if token, exists := manager.Theme().ZIndexToken(zKey); exists {
			zIndex = token.Value()
		} else {
			// Fallback to treating the key as a direct value
			zIndex = css.Raw(zKey)
		}
		return css.RuleSet(fmt.Sprintf(".%s", className),
			css.Set(cssgen.ZIndex, zIndex),
		)
	})
}

// Convenience functions that use the default manager

// OpacityClass creates an opacity utility using the default manager.
func OpacityClass(opacityKey string) css.Rule {
	return Opacity(defaultManager, opacityKey)
}

// Shadow creates a box shadow utility using the default manager.
func Shadow(shadowKey string) css.Rule {
	return BoxShadow(defaultManager, shadowKey)
}

// BlurClass creates a blur utility using the default manager.
func BlurClass(blurKey string) css.Rule {
	return Blur(defaultManager, blurKey)
}

// BrightnessClass creates a brightness utility using the default manager.
func BrightnessClass(brightnessKey string) css.Rule {
	return Brightness(defaultManager, brightnessKey)
}

// ContrastClass creates a contrast utility using the default manager.
func ContrastClass(contrastKey string) css.Rule {
	return Contrast(defaultManager, contrastKey)
}

// GrayscaleClass creates a grayscale utility using the default manager.
func GrayscaleClass(grayscaleKey string) css.Rule {
	return Grayscale(defaultManager, grayscaleKey)
}

// InvertClass creates an invert utility using the default manager.
func InvertClass(invertKey string) css.Rule {
	return Invert(defaultManager, invertKey)
}

// SaturateClass creates a saturate utility using the default manager.
func SaturateClass(saturateKey string) css.Rule {
	return Saturate(defaultManager, saturateKey)
}

// SepiaClass creates a sepia utility using the default manager.
func SepiaClass(sepiaKey string) css.Rule {
	return Sepia(defaultManager, sepiaKey)
}

// Z creates a z-index utility using the default manager.
func Z(zKey string) css.Rule {
	return ZIndex(defaultManager, zKey)
}
