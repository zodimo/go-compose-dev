package tokens

// ColorSchemeKeyToken represents a key for a color in the Material 3 color scheme.
// These keys are used to look up specific color roles.
type ColorSchemeTokenKey int

const (
	// ColorSchemeKeyUnspecified is the zero value, representing no specific color key.
	ColorSchemeTokenKeyUnspecified ColorSchemeTokenKey = iota
	ColorSchemeTokenKeyBackground
	ColorSchemeTokenKeyError
	ColorSchemeTokenKeyErrorContainer
	ColorSchemeTokenKeyInverseOnSurface
	ColorSchemeTokenKeyInversePrimary
	ColorSchemeTokenKeyInverseSurface
	ColorSchemeTokenKeyOnBackground
	ColorSchemeTokenKeyOnError
	ColorSchemeTokenKeyOnErrorContainer
	ColorSchemeTokenKeyOnPrimary
	ColorSchemeTokenKeyOnPrimaryContainer
	ColorSchemeTokenKeyOnPrimaryFixed
	ColorSchemeTokenKeyOnPrimaryFixedVariant
	ColorSchemeTokenKeyOnSecondary
	ColorSchemeTokenKeyOnSecondaryContainer
	ColorSchemeTokenKeyOnSecondaryFixed
	ColorSchemeTokenKeyOnSecondaryFixedVariant
	ColorSchemeTokenKeyOnSurface
	ColorSchemeTokenKeyOnSurfaceVariant
	ColorSchemeTokenKeyOnTertiary
	ColorSchemeTokenKeyOnTertiaryContainer
	ColorSchemeTokenKeyOnTertiaryFixed
	ColorSchemeTokenKeyOnTertiaryFixedVariant
	ColorSchemeTokenKeyOutline
	ColorSchemeTokenKeyOutlineVariant
	ColorSchemeTokenKeyPrimary
	ColorSchemeTokenKeyPrimaryContainer
	ColorSchemeTokenKeyPrimaryFixed
	ColorSchemeTokenKeyPrimaryFixedDim
	ColorSchemeTokenKeyScrim
	ColorSchemeTokenKeySecondary
	ColorSchemeTokenKeySecondaryContainer
	ColorSchemeTokenKeySecondaryFixed
	ColorSchemeTokenKeySecondaryFixedDim
	ColorSchemeTokenKeySurface
	ColorSchemeTokenKeySurfaceBright
	ColorSchemeTokenKeySurfaceContainer
	ColorSchemeTokenKeySurfaceContainerHigh
	ColorSchemeTokenKeySurfaceContainerHighest
	ColorSchemeTokenKeySurfaceContainerLow
	ColorSchemeTokenKeySurfaceContainerLowest
	ColorSchemeTokenKeySurfaceDim
	ColorSchemeTokenKeySurfaceTint
	ColorSchemeTokenKeySurfaceVariant
	ColorSchemeTokenKeyTertiary
	ColorSchemeTokenKeyTertiaryContainer
	ColorSchemeTokenKeyTertiaryFixed
	ColorSchemeTokenKeyTertiaryFixedDim
)

func (c ColorSchemeTokenKey) String() string {
	switch c {
	case ColorSchemeTokenKeyBackground:
		return "Background"
	case ColorSchemeTokenKeyError:
		return "Error"
	case ColorSchemeTokenKeyErrorContainer:
		return "ErrorContainer"
	case ColorSchemeTokenKeyInverseOnSurface:
		return "InverseOnSurface"
	case ColorSchemeTokenKeyInversePrimary:
		return "InversePrimary"
	case ColorSchemeTokenKeyInverseSurface:
		return "InverseSurface"
	case ColorSchemeTokenKeyOnBackground:
		return "OnBackground"
	case ColorSchemeTokenKeyOnError:
		return "OnError"
	case ColorSchemeTokenKeyOnErrorContainer:
		return "OnErrorContainer"
	case ColorSchemeTokenKeyOnPrimary:
		return "OnPrimary"
	case ColorSchemeTokenKeyOnPrimaryContainer:
		return "OnPrimaryContainer"
	case ColorSchemeTokenKeyOnPrimaryFixed:
		return "OnPrimaryFixed"
	case ColorSchemeTokenKeyOnPrimaryFixedVariant:
		return "OnPrimaryFixedVariant"
	case ColorSchemeTokenKeyOnSecondary:
		return "OnSecondary"
	case ColorSchemeTokenKeyOnSecondaryContainer:
		return "OnSecondaryContainer"
	case ColorSchemeTokenKeyOnSecondaryFixed:
		return "OnSecondaryFixed"
	case ColorSchemeTokenKeyOnSecondaryFixedVariant:
		return "OnSecondaryFixedVariant"
	case ColorSchemeTokenKeyOnSurface:
		return "OnSurface"
	case ColorSchemeTokenKeyOnSurfaceVariant:
		return "OnSurfaceVariant"
	case ColorSchemeTokenKeyOnTertiary:
		return "OnTertiary"
	case ColorSchemeTokenKeyOnTertiaryContainer:
		return "OnTertiaryContainer"
	case ColorSchemeTokenKeyOnTertiaryFixed:
		return "OnTertiaryFixed"
	case ColorSchemeTokenKeyOnTertiaryFixedVariant:
		return "OnTertiaryFixedVariant"
	case ColorSchemeTokenKeyOutline:
		return "Outline"
	case ColorSchemeTokenKeyOutlineVariant:
		return "OutlineVariant"
	case ColorSchemeTokenKeyPrimary:
		return "Primary"
	case ColorSchemeTokenKeyPrimaryContainer:
		return "PrimaryContainer"
	case ColorSchemeTokenKeyPrimaryFixed:
		return "PrimaryFixed"
	case ColorSchemeTokenKeyPrimaryFixedDim:
		return "PrimaryFixedDim"
	case ColorSchemeTokenKeyScrim:
		return "Scrim"
	case ColorSchemeTokenKeySecondary:
		return "Secondary"
	case ColorSchemeTokenKeySecondaryContainer:
		return "SecondaryContainer"
	case ColorSchemeTokenKeySecondaryFixed:
		return "SecondaryFixed"
	case ColorSchemeTokenKeySecondaryFixedDim:
		return "SecondaryFixedDim"
	case ColorSchemeTokenKeySurface:
		return "Surface"
	case ColorSchemeTokenKeySurfaceBright:
		return "SurfaceBright"
	case ColorSchemeTokenKeySurfaceContainer:
		return "SurfaceContainer"
	case ColorSchemeTokenKeySurfaceContainerHigh:
		return "SurfaceContainerHigh"
	case ColorSchemeTokenKeySurfaceContainerHighest:
		return "SurfaceContainerHighest"
	case ColorSchemeTokenKeySurfaceContainerLow:
		return "SurfaceContainerLow"
	case ColorSchemeTokenKeySurfaceContainerLowest:
		return "SurfaceContainerLowest"
	case ColorSchemeTokenKeySurfaceDim:
		return "SurfaceDim"
	case ColorSchemeTokenKeySurfaceTint:
		return "SurfaceTint"
	case ColorSchemeTokenKeySurfaceVariant:
		return "SurfaceVariant"
	case ColorSchemeTokenKeyTertiary:
		return "Tertiary"
	case ColorSchemeTokenKeyTertiaryContainer:
		return "TertiaryContainer"
	case ColorSchemeTokenKeyTertiaryFixed:
		return "TertiaryFixed"
	case ColorSchemeTokenKeyTertiaryFixedDim:
		return "TertiaryFixedDim"
	default:
		return "Unspecified"
	}
}

var ColorSchemeKeyTokens = ColorSchemeKeyTokensData{
	Background:              ColorSchemeTokenKeyBackground,
	Error:                   ColorSchemeTokenKeyError,
	ErrorContainer:          ColorSchemeTokenKeyErrorContainer,
	InverseOnSurface:        ColorSchemeTokenKeyInverseOnSurface,
	InversePrimary:          ColorSchemeTokenKeyInversePrimary,
	InverseSurface:          ColorSchemeTokenKeyInverseSurface,
	OnBackground:            ColorSchemeTokenKeyOnBackground,
	OnError:                 ColorSchemeTokenKeyOnError,
	OnErrorContainer:        ColorSchemeTokenKeyOnErrorContainer,
	OnPrimary:               ColorSchemeTokenKeyOnPrimary,
	OnPrimaryContainer:      ColorSchemeTokenKeyOnPrimaryContainer,
	OnPrimaryFixed:          ColorSchemeTokenKeyOnPrimaryFixed,
	OnPrimaryFixedVariant:   ColorSchemeTokenKeyOnPrimaryFixedVariant,
	OnSecondary:             ColorSchemeTokenKeyOnSecondary,
	OnSecondaryContainer:    ColorSchemeTokenKeyOnSecondaryContainer,
	OnSecondaryFixed:        ColorSchemeTokenKeyOnSecondaryFixed,
	OnSecondaryFixedVariant: ColorSchemeTokenKeyOnSecondaryFixedVariant,
	OnSurface:               ColorSchemeTokenKeyOnSurface,
	OnSurfaceVariant:        ColorSchemeTokenKeyOnSurfaceVariant,
	OnTertiary:              ColorSchemeTokenKeyOnTertiary,
	OnTertiaryContainer:     ColorSchemeTokenKeyOnTertiaryContainer,
	OnTertiaryFixed:         ColorSchemeTokenKeyOnTertiaryFixed,
	OnTertiaryFixedVariant:  ColorSchemeTokenKeyOnTertiaryFixedVariant,
	Outline:                 ColorSchemeTokenKeyOutline,
	OutlineVariant:          ColorSchemeTokenKeyOutlineVariant,
	Primary:                 ColorSchemeTokenKeyPrimary,
	PrimaryContainer:        ColorSchemeTokenKeyPrimaryContainer,
	PrimaryFixed:            ColorSchemeTokenKeyPrimaryFixed,
	PrimaryFixedDim:         ColorSchemeTokenKeyPrimaryFixedDim,
	Scrim:                   ColorSchemeTokenKeyScrim,
	Secondary:               ColorSchemeTokenKeySecondary,
	SecondaryContainer:      ColorSchemeTokenKeySecondaryContainer,
	SecondaryFixed:          ColorSchemeTokenKeySecondaryFixed,
	SecondaryFixedDim:       ColorSchemeTokenKeySecondaryFixedDim,
	Surface:                 ColorSchemeTokenKeySurface,
	SurfaceBright:           ColorSchemeTokenKeySurfaceBright,
	SurfaceContainer:        ColorSchemeTokenKeySurfaceContainer,
	SurfaceContainerHigh:    ColorSchemeTokenKeySurfaceContainerHigh,
	SurfaceContainerHighest: ColorSchemeTokenKeySurfaceContainerHighest,
	SurfaceContainerLow:     ColorSchemeTokenKeySurfaceContainerLow,
	SurfaceContainerLowest:  ColorSchemeTokenKeySurfaceContainerLowest,
	SurfaceDim:              ColorSchemeTokenKeySurfaceDim,
	SurfaceTint:             ColorSchemeTokenKeySurfaceTint,
	SurfaceVariant:          ColorSchemeTokenKeySurfaceVariant,
	Tertiary:                ColorSchemeTokenKeyTertiary,
	TertiaryContainer:       ColorSchemeTokenKeyTertiaryContainer,
	TertiaryFixed:           ColorSchemeTokenKeyTertiaryFixed,
	TertiaryFixedDim:        ColorSchemeTokenKeyTertiaryFixedDim,
}

type ColorSchemeKeyTokensData struct {
	Background              ColorSchemeTokenKey
	Error                   ColorSchemeTokenKey
	ErrorContainer          ColorSchemeTokenKey
	InverseOnSurface        ColorSchemeTokenKey
	InversePrimary          ColorSchemeTokenKey
	InverseSurface          ColorSchemeTokenKey
	OnBackground            ColorSchemeTokenKey
	OnError                 ColorSchemeTokenKey
	OnErrorContainer        ColorSchemeTokenKey
	OnPrimary               ColorSchemeTokenKey
	OnPrimaryContainer      ColorSchemeTokenKey
	OnPrimaryFixed          ColorSchemeTokenKey
	OnPrimaryFixedVariant   ColorSchemeTokenKey
	OnSecondary             ColorSchemeTokenKey
	OnSecondaryContainer    ColorSchemeTokenKey
	OnSecondaryFixed        ColorSchemeTokenKey
	OnSecondaryFixedVariant ColorSchemeTokenKey
	OnSurface               ColorSchemeTokenKey
	OnSurfaceVariant        ColorSchemeTokenKey
	OnTertiary              ColorSchemeTokenKey
	OnTertiaryContainer     ColorSchemeTokenKey
	OnTertiaryFixed         ColorSchemeTokenKey
	OnTertiaryFixedVariant  ColorSchemeTokenKey
	Outline                 ColorSchemeTokenKey
	OutlineVariant          ColorSchemeTokenKey
	Primary                 ColorSchemeTokenKey
	PrimaryContainer        ColorSchemeTokenKey
	PrimaryFixed            ColorSchemeTokenKey
	PrimaryFixedDim         ColorSchemeTokenKey
	Scrim                   ColorSchemeTokenKey
	Secondary               ColorSchemeTokenKey
	SecondaryContainer      ColorSchemeTokenKey
	SecondaryFixed          ColorSchemeTokenKey
	SecondaryFixedDim       ColorSchemeTokenKey
	Surface                 ColorSchemeTokenKey
	SurfaceBright           ColorSchemeTokenKey
	SurfaceContainer        ColorSchemeTokenKey
	SurfaceContainerHigh    ColorSchemeTokenKey
	SurfaceContainerHighest ColorSchemeTokenKey
	SurfaceContainerLow     ColorSchemeTokenKey
	SurfaceContainerLowest  ColorSchemeTokenKey
	SurfaceDim              ColorSchemeTokenKey
	SurfaceTint             ColorSchemeTokenKey
	SurfaceVariant          ColorSchemeTokenKey
	Tertiary                ColorSchemeTokenKey
	TertiaryContainer       ColorSchemeTokenKey
	TertiaryFixed           ColorSchemeTokenKey
	TertiaryFixedDim        ColorSchemeTokenKey
}
