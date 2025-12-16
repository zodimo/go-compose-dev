package theme

import "github.com/zodimo/go-compose/theme/colorrole"

var (
	colorPrimary               = colorDescriptor{colorRole: colorrole.Primary}
	colorOnPrimary             = colorDescriptor{colorRole: colorrole.OnPrimary}
	colorPrimaryContainer      = colorDescriptor{colorRole: colorrole.PrimaryContainer}
	colorOnPrimaryContainer    = colorDescriptor{colorRole: colorrole.OnPrimaryContainer}
	colorPrimaryFixed          = colorDescriptor{colorRole: colorrole.PrimaryFixed}
	colorOnPrimaryFixed        = colorDescriptor{colorRole: colorrole.OnPrimaryFixed}
	colorPrimaryFixedVariant   = colorDescriptor{colorRole: colorrole.PrimaryFixedVariant}
	colorOnPrimaryFixedVariant = colorDescriptor{colorRole: colorrole.OnPrimaryFixedVariant}

	colorSecondary               = colorDescriptor{colorRole: colorrole.Secondary}
	colorOnSecondary             = colorDescriptor{colorRole: colorrole.OnSecondary}
	colorSecondaryContainer      = colorDescriptor{colorRole: colorrole.SecondaryContainer}
	colorOnSecondaryContainer    = colorDescriptor{colorRole: colorrole.OnSecondaryContainer}
	colorSecondaryFixed          = colorDescriptor{colorRole: colorrole.SecondaryFixed}
	colorOnSecondaryFixed        = colorDescriptor{colorRole: colorrole.OnSecondaryFixed}
	colorSecondaryFixedVariant   = colorDescriptor{colorRole: colorrole.SecondaryFixedVariant}
	colorOnSecondaryFixedVariant = colorDescriptor{colorRole: colorrole.OnSecondaryFixedVariant}

	colorTertiary               = colorDescriptor{colorRole: colorrole.Tertiary}
	colorOnTertiary             = colorDescriptor{colorRole: colorrole.OnTertiary}
	colorTertiaryContainer      = colorDescriptor{colorRole: colorrole.TertiaryContainer}
	colorOnTertiaryContainer    = colorDescriptor{colorRole: colorrole.OnTertiaryContainer}
	colorTertiaryFixed          = colorDescriptor{colorRole: colorrole.TertiaryFixed}
	colorOnTertiaryFixed        = colorDescriptor{colorRole: colorrole.OnTertiaryFixed}
	colorTertiaryFixedVariant   = colorDescriptor{colorRole: colorrole.TertiaryFixedVariant}
	colorOnTertiaryFixedVariant = colorDescriptor{colorRole: colorrole.OnTertiaryFixedVariant}

	colorSurface                 = colorDescriptor{colorRole: colorrole.Surface}
	colorOnSurface               = colorDescriptor{colorRole: colorrole.OnSurface}
	colorSurfaceVariant          = colorDescriptor{colorRole: colorrole.SurfaceVariant}
	colorOnSurfaceVariant        = colorDescriptor{colorRole: colorrole.OnSurfaceVariant}
	colorSurfaceDim              = colorDescriptor{colorRole: colorrole.SurfaceDim}
	colorSurfaceBright           = colorDescriptor{colorRole: colorrole.SurfaceBright}
	colorSurfaceContainerLowest  = colorDescriptor{colorRole: colorrole.SurfaceContainerLowest}
	colorSurfaceContainerLow     = colorDescriptor{colorRole: colorrole.SurfaceContainerLow}
	colorSurfaceContainer        = colorDescriptor{colorRole: colorrole.SurfaceContainer}
	colorSurfaceContainerHigh    = colorDescriptor{colorRole: colorrole.SurfaceContainerHigh}
	colorSurfaceContainerHighest = colorDescriptor{colorRole: colorrole.SurfaceContainerHighest}

	colorInverseSurface   = colorDescriptor{colorRole: colorrole.InverseSurface}
	colorOnInverseSurface = colorDescriptor{colorRole: colorrole.OnInverseSurface}
	colorInversePrimary   = colorDescriptor{colorRole: colorrole.InversePrimary}

	colorBackground   = colorDescriptor{colorRole: colorrole.Background}
	colorOnBackground = colorDescriptor{colorRole: colorrole.OnBackground}

	colorOutline        = colorDescriptor{colorRole: colorrole.Outline}
	colorOutlineVariant = colorDescriptor{colorRole: colorrole.OutlineVariant}

	colorError            = colorDescriptor{colorRole: colorrole.Error}
	colorOnError          = colorDescriptor{colorRole: colorrole.OnError}
	colorErrorContainer   = colorDescriptor{colorRole: colorrole.ErrorContainer}
	colorOnErrorContainer = colorDescriptor{colorRole: colorrole.OnErrorContainer}

	colorScrim  = colorDescriptor{colorRole: colorrole.Scrim}
	colorShadow = colorDescriptor{colorRole: colorrole.Shadow}

	// Basic Roles
	colorBasicBg         = colorDescriptor{colorRole: colorrole.BasicBg}
	colorBasicFg         = colorDescriptor{colorRole: colorrole.BasicFg}
	colorBasicContrastBg = colorDescriptor{colorRole: colorrole.BasicContrastBg}
	colorBasicContrastFg = colorDescriptor{colorRole: colorrole.BasicContrastFg}
)

type PrimaryColorRoleDescriptors struct {
	Primary        ColorDescriptor
	OnPrimary      ColorDescriptor
	Container      ColorDescriptor
	OnContainer    ColorDescriptor
	Fixed          ColorDescriptor
	OnFixed        ColorDescriptor
	FixedVariant   ColorDescriptor
	OnFixedVariant ColorDescriptor
}

func newPrimaryColorRoleDescriptors() PrimaryColorRoleDescriptors {
	return PrimaryColorRoleDescriptors{
		Primary:        colorPrimary,
		OnPrimary:      colorOnPrimary,
		Container:      colorPrimaryContainer,
		OnContainer:    colorOnPrimaryContainer,
		Fixed:          colorPrimaryFixed,
		OnFixed:        colorOnPrimaryFixed,
		FixedVariant:   colorPrimaryFixedVariant,
		OnFixedVariant: colorOnPrimaryFixedVariant,
	}
}

type SecondaryColorRoleDescriptors struct {
	Secondary      ColorDescriptor
	OnSecondary    ColorDescriptor
	Container      ColorDescriptor
	OnContainer    ColorDescriptor
	Fixed          ColorDescriptor
	OnFixed        ColorDescriptor
	FixedVariant   ColorDescriptor
	OnFixedVariant ColorDescriptor
}

func newSecondaryColorRoleDescriptors() SecondaryColorRoleDescriptors {
	return SecondaryColorRoleDescriptors{
		Secondary:      colorSecondary,
		OnSecondary:    colorOnSecondary,
		Container:      colorSecondaryContainer,
		OnContainer:    colorOnSecondaryContainer,
		Fixed:          colorSecondaryFixed,
		OnFixed:        colorOnSecondaryFixed,
		FixedVariant:   colorSecondaryFixedVariant,
		OnFixedVariant: colorOnSecondaryFixedVariant,
	}
}

type TertiaryColorRoleDescriptors struct {
	Tertiary       ColorDescriptor
	OnTertiary     ColorDescriptor
	Container      ColorDescriptor
	OnContainer    ColorDescriptor
	Fixed          ColorDescriptor
	OnFixed        ColorDescriptor
	FixedVariant   ColorDescriptor
	OnFixedVariant ColorDescriptor
}

func newTertiaryColorRoleDescriptors() TertiaryColorRoleDescriptors {
	return TertiaryColorRoleDescriptors{
		Tertiary:       colorTertiary,
		OnTertiary:     colorOnTertiary,
		Container:      colorTertiaryContainer,
		OnContainer:    colorOnTertiaryContainer,
		Fixed:          colorTertiaryFixed,
		OnFixed:        colorOnTertiaryFixed,
		FixedVariant:   colorTertiaryFixedVariant,
		OnFixedVariant: colorOnTertiaryFixedVariant,
	}
}

type SurfaceColorRoleDescriptors struct {
	Surface          ColorDescriptor
	OnSurface        ColorDescriptor
	Variant          ColorDescriptor
	OnVariant        ColorDescriptor
	Dim              ColorDescriptor
	Bright           ColorDescriptor
	ContainerLowest  ColorDescriptor
	ContainerLow     ColorDescriptor
	Container        ColorDescriptor
	ContainerHigh    ColorDescriptor
	ContainerHighest ColorDescriptor
}

func newSurfaceColorRoleDescriptors() SurfaceColorRoleDescriptors {
	return SurfaceColorRoleDescriptors{
		Surface:          colorSurface,
		OnSurface:        colorOnSurface,
		Variant:          colorSurfaceVariant,
		OnVariant:        colorOnSurfaceVariant,
		Dim:              colorSurfaceDim,
		Bright:           colorSurfaceBright,
		ContainerLowest:  colorSurfaceContainerLowest,
		ContainerLow:     colorSurfaceContainerLow,
		Container:        colorSurfaceContainer,
		ContainerHigh:    colorSurfaceContainerHigh,
		ContainerHighest: colorSurfaceContainerHighest,
	}
}

type InverseColorRoleDescriptors struct {
	Surface   ColorDescriptor
	OnSurface ColorDescriptor
	Primary   ColorDescriptor
}

func newInverseColorRoleDescriptors() InverseColorRoleDescriptors {
	return InverseColorRoleDescriptors{
		Surface:   colorInverseSurface,
		OnSurface: colorOnInverseSurface,
		Primary:   colorInversePrimary,
	}
}

type BackgroundColorRoleDescriptors struct {
	Background   ColorDescriptor
	OnBackground ColorDescriptor
}

func newBackgroundColorRoleDescriptors() BackgroundColorRoleDescriptors {
	return BackgroundColorRoleDescriptors{
		Background:   colorBackground,
		OnBackground: colorOnBackground,
	}
}

type OutlineColorRoleDescriptors struct {
	Outline        ColorDescriptor
	OutlineVariant ColorDescriptor
}

func newOutlineColorRoleDescriptors() OutlineColorRoleDescriptors {
	return OutlineColorRoleDescriptors{
		Outline:        colorOutline,
		OutlineVariant: colorOutlineVariant,
	}
}

type ErrorColorRoleDescriptors struct {
	Error       ColorDescriptor
	OnError     ColorDescriptor
	Container   ColorDescriptor
	OnContainer ColorDescriptor
}

func newErrorColorRoleDescriptors() ErrorColorRoleDescriptors {
	return ErrorColorRoleDescriptors{
		Error:       colorError,
		OnError:     colorOnError,
		Container:   colorErrorContainer,
		OnContainer: colorOnErrorContainer,
	}
}

type ScrimColorRoleDescriptors struct {
	Scrim  ColorDescriptor
	Shadow ColorDescriptor
}

func newScrimColorRoleDescriptors() ScrimColorRoleDescriptors {
	return ScrimColorRoleDescriptors{
		Scrim:  colorScrim,
		Shadow: colorShadow,
	}
}

// Roles for Original Gioui Matrial Theme Palette
type BasicColorRoleDescriptors struct {
	BasicBg         ColorDescriptor
	BasicFg         ColorDescriptor
	BasicContrastBg ColorDescriptor
	BasicContrastFg ColorDescriptor
}

func newBasicColorRoleDescriptors() BasicColorRoleDescriptors {
	return BasicColorRoleDescriptors{
		BasicBg:         colorBasicBg,
		BasicFg:         colorBasicFg,
		BasicContrastBg: colorBasicContrastBg,
		BasicContrastFg: colorBasicContrastFg,
	}
}

type ColorRolesDescriptors struct {
	PrimaryRoles         PrimaryColorRoleDescriptors
	SecondaryRoles       SecondaryColorRoleDescriptors
	TertiaryRoles        TertiaryColorRoleDescriptors
	SurfaceRoles         SurfaceColorRoleDescriptors
	InverseRoles         InverseColorRoleDescriptors
	BackgroundColorRoles BackgroundColorRoleDescriptors
	OutlineRoles         OutlineColorRoleDescriptors
	ErrorRoles           ErrorColorRoleDescriptors
	ScrimRoles           ScrimColorRoleDescriptors
}

func newM3ColorRolesDescriptors() ColorRolesDescriptors {
	return ColorRolesDescriptors{
		PrimaryRoles:         newPrimaryColorRoleDescriptors(),
		SecondaryRoles:       newSecondaryColorRoleDescriptors(),
		TertiaryRoles:        newTertiaryColorRoleDescriptors(),
		SurfaceRoles:         newSurfaceColorRoleDescriptors(),
		InverseRoles:         newInverseColorRoleDescriptors(),
		BackgroundColorRoles: newBackgroundColorRoleDescriptors(),
		OutlineRoles:         newOutlineColorRoleDescriptors(),
		ErrorRoles:           newErrorColorRoleDescriptors(),
		ScrimRoles:           newScrimColorRoleDescriptors(),
	}
}

type ColorRoleDescriptors struct {
	ColorRolesDescriptors
	BasicColorRoleDescriptors
}

func NewColorRoleDescriptors() ColorRoleDescriptors {
	return ColorRoleDescriptors{
		ColorRolesDescriptors:     newM3ColorRolesDescriptors(),
		BasicColorRoleDescriptors: newBasicColorRoleDescriptors(),
	}
}
