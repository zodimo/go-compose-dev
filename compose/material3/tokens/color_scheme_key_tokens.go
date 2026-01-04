package tokens

// ColorSchemeKeyToken represents a key for a color in the Material 3 color scheme.
// These keys are used to look up specific color roles.
type ColorSchemeKeyToken int

const (
	// ColorSchemeKeyUnspecified is the zero value, representing no specific color key.
	ColorSchemeKeyUnspecified ColorSchemeKeyToken = iota

	ColorSchemeKeyBackground
	ColorSchemeKeyError
	ColorSchemeKeyErrorContainer
	ColorSchemeKeyInverseOnSurface
	ColorSchemeKeyInversePrimary
	ColorSchemeKeyInverseSurface
	ColorSchemeKeyOnBackground
	ColorSchemeKeyOnError
	ColorSchemeKeyOnErrorContainer
	ColorSchemeKeyOnPrimary
	ColorSchemeKeyOnPrimaryContainer
	ColorSchemeKeyOnPrimaryFixed
	ColorSchemeKeyOnPrimaryFixedVariant
	ColorSchemeKeyOnSecondary
	ColorSchemeKeyOnSecondaryContainer
	ColorSchemeKeyOnSecondaryFixed
	ColorSchemeKeyOnSecondaryFixedVariant
	ColorSchemeKeyOnSurface
	ColorSchemeKeyOnSurfaceVariant
	ColorSchemeKeyOnTertiary
	ColorSchemeKeyOnTertiaryContainer
	ColorSchemeKeyOnTertiaryFixed
	ColorSchemeKeyOnTertiaryFixedVariant
	ColorSchemeKeyOutline
	ColorSchemeKeyOutlineVariant
	ColorSchemeKeyPrimary
	ColorSchemeKeyPrimaryContainer
	ColorSchemeKeyPrimaryFixed
	ColorSchemeKeyPrimaryFixedDim
	ColorSchemeKeyScrim
	ColorSchemeKeySecondary
	ColorSchemeKeySecondaryContainer
	ColorSchemeKeySecondaryFixed
	ColorSchemeKeySecondaryFixedDim
	ColorSchemeKeySurface
	ColorSchemeKeySurfaceBright
	ColorSchemeKeySurfaceContainer
	ColorSchemeKeySurfaceContainerHigh
	ColorSchemeKeySurfaceContainerHighest
	ColorSchemeKeySurfaceContainerLow
	ColorSchemeKeySurfaceContainerLowest
	ColorSchemeKeySurfaceDim
	ColorSchemeKeySurfaceTint
	ColorSchemeKeySurfaceVariant
	ColorSchemeKeyTertiary
	ColorSchemeKeyTertiaryContainer
	ColorSchemeKeyTertiaryFixed
	ColorSchemeKeyTertiaryFixedDim
)

func (c ColorSchemeKeyToken) String() string {
	switch c {
	case ColorSchemeKeyBackground:
		return "Background"
	case ColorSchemeKeyError:
		return "Error"
	case ColorSchemeKeyErrorContainer:
		return "ErrorContainer"
	case ColorSchemeKeyInverseOnSurface:
		return "InverseOnSurface"
	case ColorSchemeKeyInversePrimary:
		return "InversePrimary"
	case ColorSchemeKeyInverseSurface:
		return "InverseSurface"
	case ColorSchemeKeyOnBackground:
		return "OnBackground"
	case ColorSchemeKeyOnError:
		return "OnError"
	case ColorSchemeKeyOnErrorContainer:
		return "OnErrorContainer"
	case ColorSchemeKeyOnPrimary:
		return "OnPrimary"
	case ColorSchemeKeyOnPrimaryContainer:
		return "OnPrimaryContainer"
	case ColorSchemeKeyOnPrimaryFixed:
		return "OnPrimaryFixed"
	case ColorSchemeKeyOnPrimaryFixedVariant:
		return "OnPrimaryFixedVariant"
	case ColorSchemeKeyOnSecondary:
		return "OnSecondary"
	case ColorSchemeKeyOnSecondaryContainer:
		return "OnSecondaryContainer"
	case ColorSchemeKeyOnSecondaryFixed:
		return "OnSecondaryFixed"
	case ColorSchemeKeyOnSecondaryFixedVariant:
		return "OnSecondaryFixedVariant"
	case ColorSchemeKeyOnSurface:
		return "OnSurface"
	case ColorSchemeKeyOnSurfaceVariant:
		return "OnSurfaceVariant"
	case ColorSchemeKeyOnTertiary:
		return "OnTertiary"
	case ColorSchemeKeyOnTertiaryContainer:
		return "OnTertiaryContainer"
	case ColorSchemeKeyOnTertiaryFixed:
		return "OnTertiaryFixed"
	case ColorSchemeKeyOnTertiaryFixedVariant:
		return "OnTertiaryFixedVariant"
	case ColorSchemeKeyOutline:
		return "Outline"
	case ColorSchemeKeyOutlineVariant:
		return "OutlineVariant"
	case ColorSchemeKeyPrimary:
		return "Primary"
	case ColorSchemeKeyPrimaryContainer:
		return "PrimaryContainer"
	case ColorSchemeKeyPrimaryFixed:
		return "PrimaryFixed"
	case ColorSchemeKeyPrimaryFixedDim:
		return "PrimaryFixedDim"
	case ColorSchemeKeyScrim:
		return "Scrim"
	case ColorSchemeKeySecondary:
		return "Secondary"
	case ColorSchemeKeySecondaryContainer:
		return "SecondaryContainer"
	case ColorSchemeKeySecondaryFixed:
		return "SecondaryFixed"
	case ColorSchemeKeySecondaryFixedDim:
		return "SecondaryFixedDim"
	case ColorSchemeKeySurface:
		return "Surface"
	case ColorSchemeKeySurfaceBright:
		return "SurfaceBright"
	case ColorSchemeKeySurfaceContainer:
		return "SurfaceContainer"
	case ColorSchemeKeySurfaceContainerHigh:
		return "SurfaceContainerHigh"
	case ColorSchemeKeySurfaceContainerHighest:
		return "SurfaceContainerHighest"
	case ColorSchemeKeySurfaceContainerLow:
		return "SurfaceContainerLow"
	case ColorSchemeKeySurfaceContainerLowest:
		return "SurfaceContainerLowest"
	case ColorSchemeKeySurfaceDim:
		return "SurfaceDim"
	case ColorSchemeKeySurfaceTint:
		return "SurfaceTint"
	case ColorSchemeKeySurfaceVariant:
		return "SurfaceVariant"
	case ColorSchemeKeyTertiary:
		return "Tertiary"
	case ColorSchemeKeyTertiaryContainer:
		return "TertiaryContainer"
	case ColorSchemeKeyTertiaryFixed:
		return "TertiaryFixed"
	case ColorSchemeKeyTertiaryFixedDim:
		return "TertiaryFixedDim"
	default:
		return "Unspecified"
	}
}
