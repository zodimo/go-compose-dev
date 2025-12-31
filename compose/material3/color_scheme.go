package material3

type ColorScheme struct {
	Primary             *ColorSet
	PrimaryContainer    *ColorSet
	PrimaryFixed        *ColorSet
	PrimaryFixedVariant *ColorSet

	Secondary             *ColorSet
	SecondaryContainer    *ColorSet
	SecondaryFixed        *ColorSet
	SecondaryFixedVariant *ColorSet

	Tertiary             *ColorSet
	TertiaryContainer    *ColorSet
	TertiaryFixed        *ColorSet
	TertiaryFixedVariant *ColorSet

	Surface                 *ColorSet
	SurfaceVariant          *ColorSet
	SurfaceDim              Color
	SurfaceBright           Color
	SurfaceContainerLowest  Color
	SurfaceContainerLow     Color
	SurfaceContainer        Color
	SurfaceContainerHigh    Color
	SurfaceContainerHighest Color

	InverseSurface *ColorSet
	InversePrimary Color

	Background *ColorSet

	Outline        Color
	OutlineVariant Color

	Error          *ColorSet
	ErrorContainer *ColorSet

	Scrim  Color
	Shadow Color
}

func (c *ColorScheme) Copy(options ...ColorSchemeOption) *ColorScheme {
	copy := *c
	for _, option := range options {
		option(&copy)
	}
	return &copy
}
