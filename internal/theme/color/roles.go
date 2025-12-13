package color

// WIP for future
type M3ColorRole int

const (
	Primary M3ColorRole = iota
	OnPrimary
	PrimaryContainer
	OnPrimaryContainer
	PrimaryFixed
	OnPrimaryFixed
	PrimaryFixedVariant
	OnPrimaryFixedVariant

	Secondary
	OnSecondary
	SecondaryContainer
	OnSecondaryContainer
	SecondaryFixed
	OnSecondaryFixed
	SecondaryFixedVariant
	OnSecondaryFixedVariant

	Tertiary
	OnTertiary
	TertiaryContainer
	OnTertiaryContainer
	TertiaryFixed
	OnTertiaryFixed
	TertiaryFixedVariant
	OnTertiaryFixedVariant

	Surface
	OnSurface
	SurfaceVariant
	OnSurfaceVariant

	SurfaceDim
	SurfaceBright
	SurfaceContainerLowest
	SurfaceContainerLow
	SurfaceContainer
	SurfaceContainerHigh
	SurfaceContainerHighest

	InverseSurface
	OnInverseSurface
	InversePrimary

	Background
	OnBackground

	Outline
	OutlineVariant

	Error
	OnError
	ErrorContainer
	OnErrorContainer

	Scrim
	Shadow
)
